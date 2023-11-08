package v1alpha1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_MCOObjectReference = map[string]string{
	"":     "MCOObjectReference holds information about an object the MCO either owns or modifies in some way",
	"name": "name is the object name. Must be a lowercase RFC-1123 hostname (https://tools.ietf.org/html/rfc1123) It may consist of only alphanumeric characters, hyphens (-) and periods (.) and must be at most 253 characters in length.",
}

func (MCOObjectReference) SwaggerDoc() map[string]string {
	return map_MCOObjectReference
}

var map_MachineConfigNode = map[string]string{
	"":       "MachineConfigNode describes the health of the Machines on the system Compatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.",
	"spec":   "spec describes the configuration of the machine config node.",
	"status": "status describes the last observed state of this machine config node.",
}

func (MachineConfigNode) SwaggerDoc() map[string]string {
	return map_MachineConfigNode
}

var map_MachineConfigNodeList = map[string]string{
	"": "MachineConfigNodeList describes all of the MachinesStates on the system\n\nCompatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.",
}

func (MachineConfigNodeList) SwaggerDoc() map[string]string {
	return map_MachineConfigNodeList
}

var map_MachineConfigNodeSpec = map[string]string{
	"":              "MachineConfigNodeSpec describes the MachineConfigNode we are managing.",
	"node":          "node contains a reference to the node for this machine config node.",
	"pool":          "pool contains a reference to the machine config pool that this machine config node's referenced node belongs to.",
	"configVersion": "configVersion holds the desired config version for the node targeted by this machine config node resource. The desired version represents the machine config the node will attempt to update to. This gets set before the machine config operator validates the new machine config against the current machine config.",
}

func (MachineConfigNodeSpec) SwaggerDoc() map[string]string {
	return map_MachineConfigNodeSpec
}

var map_MachineConfigNodeSpecMachineConfigVersion = map[string]string{
	"":        "MachineConfigNodeSpecMachineConfigVersion holds the desired config version for the current observed machine config node. When Current is not equal to Desired; the MachineConfigOperator is in an upgrade phase and the machine config node will take account of upgrade related events. Otherwise they will be ignored given that certain operations happen both during the MCO's upgrade mode and the daily operations mode.",
	"desired": "desired is the name of the machine config that the the node should be upgraded to. This value is set when the machine config pool generates a new version of its rendered configuration. When this value is changed, the machine config daemon starts the node upgrade process. This value gets set in the machine config node spec once the machine config has been targeted for upgrade and before it is validated. Must be a lowercase RFC-1123 hostname (https://tools.ietf.org/html/rfc1123) It may consist of only alphanumeric characters, hyphens (-) and periods (.) and must be at most 253 characters in length.",
}

func (MachineConfigNodeSpecMachineConfigVersion) SwaggerDoc() map[string]string {
	return map_MachineConfigNodeSpecMachineConfigVersion
}

var map_MachineConfigNodeStatus = map[string]string{
	"":                   "MachineConfigNodeStatus holds the reported information on a particular machine config node.",
	"conditions":         "conditions represent the observations of a machine config node's current state.",
	"observedGeneration": "observedGeneration represents the generation observed by the controller. This field is updated when the controller observes a change to the desiredConfig in the configVersion of the machine config node spec.",
	"configVersion":      "configVersion describes the current and desired machine config for this node. The current version represents the current machine config for the node and is updated after a successful update. The desired version represents the machine config the node will attempt to update to. This desired machine config has been compared to the current machine config and has been validated by the machine config operator as one that is valid and that exists.",
}

func (MachineConfigNodeStatus) SwaggerDoc() map[string]string {
	return map_MachineConfigNodeStatus
}

var map_MachineConfigNodeStatusMachineConfigVersion = map[string]string{
	"":        "MachineConfigNodeStatusMachineConfigVersion holds the current and desired config versions as last updated in the MCN status. When the current and desired versions are not matched, the machine config pool is processing an upgrade and the machine config node will monitor the upgrade process. When the current and desired versions do not match, the machine config node will ignore these events given that certain operations happen both during the MCO's upgrade mode and the daily operations mode.",
	"current": "current is the name of the machine config currently in use on the node. This value is updated once the machine config daemon has completed the update of the configuration for the node. This value should match the desired version unless an upgrade is in progress. Must be a lowercase RFC-1123 hostname (https://tools.ietf.org/html/rfc1123) It may consist of only alphanumeric characters, hyphens (-) and periods (.) and must be at most 253 characters in length.",
	"desired": "desired is the MachineConfig the node wants to upgrade to. This value gets set in the machine config node status once the machine config has been validated against the current machine config. Must be a lowercase RFC-1123 hostname (https://tools.ietf.org/html/rfc1123) It may consist of only alphanumeric characters, hyphens (-) and periods (.) and must be at most 253 characters in length.",
}

func (MachineConfigNodeStatusMachineConfigVersion) SwaggerDoc() map[string]string {
	return map_MachineConfigNodeStatusMachineConfigVersion
}

// AUTO-GENERATED FUNCTIONS END HERE
