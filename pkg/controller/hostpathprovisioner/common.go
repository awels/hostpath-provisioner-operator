/*
Copyright 2019 The hostpath provisioner operator Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package hostpathprovisioner

const (
	// OperatorImageDefault is the default value of the operator container image name.
	OperatorImageDefault = "hostpath-provisioner-operator"
	// ProvisionerImageDefault is the default value of the provisioner container image name.
	ProvisionerImageDefault = "hostpath-provisioner"
	// CsiProvisionerImageDefault is the default value of the hostpath provisioner csi container image name.
	CsiProvisionerImageDefault = "hostpath-provisioner-csi"
	// CsiAttacherImageDefault is the default value of the sig storage csi attacher side car container image name.
	CsiAttacherImageDefault = "k8s.gcr.io/sig-storage/csi-attacher:v3.2.1"
	// CsiExternalHealthMonitorControllerImageDefault is the default value of the sig storage csi health monitor controller side car container image name.
	CsiExternalHealthMonitorControllerImageDefault = "k8s.gcr.io/sig-storage/csi-external-health-monitor-controller:v0.3.0"
	// CsiNodeDriverRegistrationImageDefault is the default value of the sig storage csi node registration side car container image name.
	CsiNodeDriverRegistrationImageDefault = "k8s.gcr.io/sig-storage/csi-node-driver-registrar:v2.2.0"
	// LivenessProbeImageDefault is the default value of the liveness probe side car container image name.
	LivenessProbeImageDefault = "k8s.gcr.io/sig-storage/livenessprobe:v2.3.0"
	// CsiSigStorageProvisionerImageDefault is the default value of the sig storage csi provisioner side car container image name.
	CsiSigStorageProvisionerImageDefault = "k8s.gcr.io/sig-storage/csi-provisioner:v2.2.1"

	provisionerImageEnvVarName                     = "PROVISIONER_IMAGE"
	csiProvisionerImageEnvVarName                  = "CSI_PROVISIONER_IMAGE"
	externalHealthMonitorControllerImageEnvVarName = "EXTERNAL_HEALTH_MON_IMAGE"
	nodeDriverRegistrarImageEnvVarName             = "NODE_DRIVER_REG_IMAGE"
	livenessProbeImageEnvVarName                   = "LIVENESS_PROVE_IMAGE"
	csiAttacherImageEnvVarName                     = "ATTACHER_IMAGE"
	csiSigStorageProvisionerImageEnvVarName        = "CSI_SIG_STORAGE_PROVISIONER_IMAGE"
	verbosityEnvVarName                            = "VERBOSITY"

	// OperatorServiceAccountName is the name of Service Account used to run the operator.
	OperatorServiceAccountName = "hostpath-provisioner-operator"
	// ProvisionerServiceAccountName is the name of Service Account used to run the controller.
	ProvisionerServiceAccountName = "hostpath-provisioner-admin"
	attacherName                  = "hostpath-provisioner-attacher"
	healthCheckName               = "hostpath-provisioner-health-check"
	// MultiPurposeHostPathProvisionerName is the name used for the DaemonSet, ClusterRole/Binding, SCC and k8s-app label value.
	MultiPurposeHostPathProvisionerName = "hostpath-provisioner"

	createResourceStart   = "CreateResourceStart"
	createResourceFailed  = "CreateResourceFailed"
	createResourceSuccess = "CreateResourceSuccess"

	updateResourceStart   = "UpdateResourceStart"
	updateResourceFailed  = "UpdateResourceFailed"
	updateResourceSuccess = "UpdateResourceSuccess"

	createMessageStart     = "Started creation of resource %T %s"
	createMessageFailed    = "Failed to create resource %s, %v"
	createMessageSucceeded = "Successfully created resource %T %s"

	updateMessageStart     = "Started update of resource %T %s"
	updateMessageFailed    = "Failed to update resource %s, %v"
	updateMessageSucceeded = "Successfully updated resource %T %s"

	provisionerHealthy        = "ProvisionerHealthy"
	provisionerHealthyMessage = "Provisioner Healthy"

	watchNameSpace = "WatchNameSpace"

	deployStarted        = "DeployStarted"
	deployStartedMessage = "Started Deployment"

	upgradeStarted = "UpgradeStarted"

	reconcileFailed = "Reconcile Failed"
)
