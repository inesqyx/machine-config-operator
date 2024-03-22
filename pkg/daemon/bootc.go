package daemon

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	rpmostreeclient "github.com/coreos/rpmostree-client-go/pkg/client"
	pivotutils "github.com/openshift/machine-config-operator/pkg/daemon/pivot/utils"
	"k8s.io/klog/v2"
)

// Synchronously invoke rpm-ostree, writing its stdout to our stdout,
// and gathering stderr into a buffer which will be returned in err
// in case of error.
func runBootc(args ...string) error {
	return runCmdSync("bootc", args...)
}

func (r *RpmOstreeClient) Initialize() error {

	// Commands like update and rebase need the pull secrets to pull images and manifests,
	// make sure we get access to them when we Initialize

	err := useMergedPullSecrets()
	if err != nil {
		klog.Errorf("error while linking rpm-ostree pull secrets %v", err)
	}

	return nil
}

// GetBootedDeployment returns the current deployment found
func (r *RpmOstreeClient) GetBootedAndStagedDeployment() (booted, staged *rpmostreeclient.Deployment, err error) {
	status, err := r.client.QueryStatus()
	if err != nil {
		return nil, nil, err
	}

	booted, err = status.GetBootedDeployment()
	staged = status.GetStagedDeployment()

	return
}

// GetStatus returns multi-line human-readable text describing system status
func (r *RpmOstreeClient) GetStatus() (string, error) {
	output, err := runGetOut("bootc", "status")
	if err != nil {
		return "", err
	}

	return string(output), nil
}

// GetBootedOSImageURL returns the image URL as well as the OSTree version(for logging) and the ostree commit (for comparisons)
// Returns the empty string if the host doesn't have a custom origin that matches pivot://
// (This could be the case for e.g. FCOS, or a future RHCOS which comes not-pivoted by default)

// replace references w bootc
func (r *RpmOstreeClient) GetBootedOSImageURL() (string, string, string, error) {
	bootedDeployment, _, err := r.GetBootedAndStagedDeployment()
	if err != nil {
		return "", "", "", err
	}

	// the canonical image URL is stored in the custom origin field.
	osImageURL := ""
	if len(bootedDeployment.CustomOrigin) > 0 {
		if strings.HasPrefix(bootedDeployment.CustomOrigin[0], "pivot://") {
			osImageURL = bootedDeployment.CustomOrigin[0][len("pivot://"):]
		}
	}

	// we have container images now, make sure we can parse those too
	if bootedDeployment.ContainerImageReference != "" {
		// right now remove ostree remote, and transport from container image reference
		ostreeImageReference, err := bootedDeployment.RequireContainerImage()
		if err != nil {
			return "", "", "", err
		}
		osImageURL = ostreeImageReference.Imgref.Image
	}

	baseChecksum := bootedDeployment.GetBaseChecksum()
	return osImageURL, bootedDeployment.Version, baseChecksum, nil
}

// stay the same
func podmanInspect(imgURL string) (imgdata *imageInspection, err error) {
	// Pull the container image if not already available
	var authArgs []string
	if _, err := os.Stat(ostreeAuthFile); err == nil {
		authArgs = append(authArgs, "--authfile", ostreeAuthFile)
	}
	args := []string{"pull", "-q"}
	args = append(args, authArgs...)
	args = append(args, imgURL)
	_, err = pivotutils.RunExt(numRetriesNetCommands, "podman", args...)
	if err != nil {
		return
	}

	inspectArgs := []string{"inspect", "--type=image"}
	inspectArgs = append(inspectArgs, fmt.Sprintf("%s", imgURL))
	var output []byte
	output, err = runGetOut("podman", inspectArgs...)
	if err != nil {
		return
	}
	var imagedataArray []imageInspection
	err = json.Unmarshal(output, &imagedataArray)
	if err != nil {
		err = fmt.Errorf("unmarshaling podman inspect: %w", err)
		return
	}
	imgdata = &imagedataArray[0]
	return

}

// RebaseLayered rebases system or errors if already rebased
func Switch(imgURL string) (err error) {
	// Try to re-link the merged pull secrets if they exist, since it could have been populated without a daemon reboot
	useMergedPullSecrets()
	klog.Infof("Executing rebase to %s", imgURL)
	return runBootc("switch", "ostree-unverified-registry:"+imgURL)
}

// linkOstreeAuthFile gives the rpm-ostree client access to secrets in the file located at `path` by symlinking so that
// rpm-ostree can use those secrets to pull images. This can be called multiple times to overwrite an older link.
func linkOstreeAuthFile(path string) error {
	if _, err := os.Lstat(ostreeAuthFile); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			if err := os.MkdirAll("/run/ostree", 0o544); err != nil {
				return err
			}
		}
	} else {
		// Remove older symlink if it exists since it needs to be overwritten
		if err := os.Remove(ostreeAuthFile); err != nil {
			return err
		}
	}

	klog.Infof("Linking ostree authfile to %s", path)
	err := os.Symlink(path, ostreeAuthFile)
	return err
}

// useMergedSecrets gives the rpm-ostree client access to secrets for the internal registry and the global pull
// secret. It does this by symlinking the merged secrets file into /run/ostree. If it fails to find the
// merged secrets, it will use the default pull secret file instead.
func useMergedPullSecrets() error {

	// check if merged secret file exists
	if _, err := os.Stat(imageRegistryAuthFile); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			klog.Errorf("Merged secret file does not exist; defaulting to cluster pull secret")
			return linkOstreeAuthFile(kubeletAuthFile)
		}
	}
	// Check that merged secret file is valid JSON
	if file, err := os.ReadFile(imageRegistryAuthFile); err != nil {
		klog.Errorf("Merged secret file could not be read; defaulting to cluster pull secret %v", err)
		return linkOstreeAuthFile(kubeletAuthFile)
	} else if !json.Valid(file) {
		klog.Errorf("Merged secret file could not be validated; defaulting to cluster pull secret %v", err)
		return linkOstreeAuthFile(kubeletAuthFile)
	}

	// Attempt to link the merged secrets file
	return linkOstreeAuthFile(imageRegistryAuthFile)
}
