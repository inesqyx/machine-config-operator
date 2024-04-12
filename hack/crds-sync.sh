#!/usr/bin/env bash

set -euo pipefail

# This is the MCO's API directory in openshift/api, every CRD living here can be directly be copied over.
cp vendor/github.com/openshift/api/machineconfiguration/v1/zz_generated.crd-manifests/*.crd.yaml install/.
cp vendor/github.com/openshift/api/machineconfiguration/v1alpha1/zz_generated.crd-manifests/*.crd.yaml install/.

#  These are MCO CRDs that live in other parts of the openshift/api, so the copies need to be more specific
CRDS_MAPPING=( 
   "operator/v1/zz_generated.crd-manifests/0000_80_machine-config_01_machineconfigurations-Default.crd.yaml:0000_80_machine-config_01_machineconfigurations-Default.crd.yaml"
   "operator/v1/zz_generated.crd-manifests/0000_80_machine-config_01_machineconfigurations-TechPreviewNoUpgrade.crd.yaml:0000_80_machine-config_01_machineconfigurations-TechPreviewNoUpgrade.crd.yaml"
   "config/v1alpha1/zz_generated.crd-manifests/0000_10_config-operator_01_clusterimagepolicies-TechPreviewNoUpgrade.crd.yaml:0000_10_config-operator_01_clusterimagepolicies-TechPreviewNoUpgrade.crd.yaml"
)

for crd in "${CRDS_MAPPING[@]}" ; do
    SRC="${crd%%:*}"
    DES="${crd##*:}"
    cp "vendor/github.com/openshift/api/machineconfiguration/$SRC" "install/$DES"
done

#this one goes in manifests rather than install, but should it? 
cp "vendor/github.com/openshift/api/config/v1alpha1/zz_generated.crd-manifests/0000_10_config-operator_01_clusterimagepolicies-TechPreviewNoUpgrade.crd.yaml" "install/0000_10_config-operator_01_clusterimagepolicy-TechPreviewNoUpgrade.crd.yaml"
cp "vendor/github.com/openshift/api/machineconfiguration/v1/zz_generated.crd-manifests/0000_80_machine-config_01_controllerconfigs-Default.crd.yaml" "manifests/controllerconfig.crd.yaml"
cp "vendor/github.com/openshift/api/machineconfiguration/v1alpha1/zz_generated.crd-manifests/0000_80_machine-config_01_machineconfignodes-TechPreviewNoUpgrade.crd.yaml" "manifests/0000_80_machine-config_01_machineconfignode-TechPreviewNoUpgrade.crd.yaml" 
#cp "vendor/github.com/openshift/api/operator/v1/0000_80_machine-config-operator_01_config.crd.yaml" "install/0000_80_machine-config-operator_01_config.crd.yaml" 
cp "vendor/github.com/openshift/api/operator/v1/zz_generated.crd-manifests/0000_80_machine-config_01_machineconfigurations-Default.crd.yaml" "install/0000_80_machine-config_01_config.crd.yaml"
cp "vendor/github.com/openshift/api/machineconfiguration/v1alpha1/zz_generated.crd-manifests/0000_80_machine-config_01_machineosbuilds-TechPreviewNoUpgrade.crd.yaml" "manifests/0000_80_machine-config_01_machineosbuild-TechPreviewNoUpgrade.crd.yaml" 
cp "vendor/github.com/openshift/api/machineconfiguration/v1alpha1/zz_generated.crd-manifests/0000_80_machine-config_01_machineosconfigs-TechPreviewNoUpgrade.crd.yaml" "manifests/0000_80_machine-config_01_machineosconfig-TechPreviewNoUpgrade.crd.yaml" 




