package managedclusters

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgentPoolMode string

const (
	AgentPoolModeSystem AgentPoolMode = "System"
	AgentPoolModeUser   AgentPoolMode = "User"
)

func PossibleValuesForAgentPoolMode() []string {
	return []string{
		string(AgentPoolModeSystem),
		string(AgentPoolModeUser),
	}
}

func parseAgentPoolMode(input string) (*AgentPoolMode, error) {
	vals := map[string]AgentPoolMode{
		"system": AgentPoolModeSystem,
		"user":   AgentPoolModeUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AgentPoolMode(input)
	return &out, nil
}

type AgentPoolType string

const (
	AgentPoolTypeAvailabilitySet         AgentPoolType = "AvailabilitySet"
	AgentPoolTypeVirtualMachineScaleSets AgentPoolType = "VirtualMachineScaleSets"
)

func PossibleValuesForAgentPoolType() []string {
	return []string{
		string(AgentPoolTypeAvailabilitySet),
		string(AgentPoolTypeVirtualMachineScaleSets),
	}
}

func parseAgentPoolType(input string) (*AgentPoolType, error) {
	vals := map[string]AgentPoolType{
		"availabilityset":         AgentPoolTypeAvailabilitySet,
		"virtualmachinescalesets": AgentPoolTypeVirtualMachineScaleSets,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AgentPoolType(input)
	return &out, nil
}

type BackendPoolType string

const (
	BackendPoolTypeNodeIP              BackendPoolType = "NodeIP"
	BackendPoolTypeNodeIPConfiguration BackendPoolType = "NodeIPConfiguration"
)

func PossibleValuesForBackendPoolType() []string {
	return []string{
		string(BackendPoolTypeNodeIP),
		string(BackendPoolTypeNodeIPConfiguration),
	}
}

func parseBackendPoolType(input string) (*BackendPoolType, error) {
	vals := map[string]BackendPoolType{
		"nodeip":              BackendPoolTypeNodeIP,
		"nodeipconfiguration": BackendPoolTypeNodeIPConfiguration,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BackendPoolType(input)
	return &out, nil
}

type Code string

const (
	CodeRunning Code = "Running"
	CodeStopped Code = "Stopped"
)

func PossibleValuesForCode() []string {
	return []string{
		string(CodeRunning),
		string(CodeStopped),
	}
}

func parseCode(input string) (*Code, error) {
	vals := map[string]Code{
		"running": CodeRunning,
		"stopped": CodeStopped,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Code(input)
	return &out, nil
}

type Expander string

const (
	ExpanderLeastNegativewaste Expander = "least-waste"
	ExpanderMostNegativepods   Expander = "most-pods"
	ExpanderPriority           Expander = "priority"
	ExpanderRandom             Expander = "random"
)

func PossibleValuesForExpander() []string {
	return []string{
		string(ExpanderLeastNegativewaste),
		string(ExpanderMostNegativepods),
		string(ExpanderPriority),
		string(ExpanderRandom),
	}
}

func parseExpander(input string) (*Expander, error) {
	vals := map[string]Expander{
		"least-waste": ExpanderLeastNegativewaste,
		"most-pods":   ExpanderMostNegativepods,
		"priority":    ExpanderPriority,
		"random":      ExpanderRandom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Expander(input)
	return &out, nil
}

type Format string

const (
	FormatAzure Format = "azure"
	FormatExec  Format = "exec"
)

func PossibleValuesForFormat() []string {
	return []string{
		string(FormatAzure),
		string(FormatExec),
	}
}

func parseFormat(input string) (*Format, error) {
	vals := map[string]Format{
		"azure": FormatAzure,
		"exec":  FormatExec,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Format(input)
	return &out, nil
}

type GPUInstanceProfile string

const (
	GPUInstanceProfileMIGFourg  GPUInstanceProfile = "MIG4g"
	GPUInstanceProfileMIGOneg   GPUInstanceProfile = "MIG1g"
	GPUInstanceProfileMIGSeveng GPUInstanceProfile = "MIG7g"
	GPUInstanceProfileMIGThreeg GPUInstanceProfile = "MIG3g"
	GPUInstanceProfileMIGTwog   GPUInstanceProfile = "MIG2g"
)

func PossibleValuesForGPUInstanceProfile() []string {
	return []string{
		string(GPUInstanceProfileMIGFourg),
		string(GPUInstanceProfileMIGOneg),
		string(GPUInstanceProfileMIGSeveng),
		string(GPUInstanceProfileMIGThreeg),
		string(GPUInstanceProfileMIGTwog),
	}
}

func parseGPUInstanceProfile(input string) (*GPUInstanceProfile, error) {
	vals := map[string]GPUInstanceProfile{
		"mig4g": GPUInstanceProfileMIGFourg,
		"mig1g": GPUInstanceProfileMIGOneg,
		"mig7g": GPUInstanceProfileMIGSeveng,
		"mig3g": GPUInstanceProfileMIGThreeg,
		"mig2g": GPUInstanceProfileMIGTwog,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GPUInstanceProfile(input)
	return &out, nil
}

type IPFamily string

const (
	IPFamilyIPvFour IPFamily = "IPv4"
	IPFamilyIPvSix  IPFamily = "IPv6"
)

func PossibleValuesForIPFamily() []string {
	return []string{
		string(IPFamilyIPvFour),
		string(IPFamilyIPvSix),
	}
}

func parseIPFamily(input string) (*IPFamily, error) {
	vals := map[string]IPFamily{
		"ipv4": IPFamilyIPvFour,
		"ipv6": IPFamilyIPvSix,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IPFamily(input)
	return &out, nil
}

type IstioIngressGatewayMode string

const (
	IstioIngressGatewayModeExternal IstioIngressGatewayMode = "External"
	IstioIngressGatewayModeInternal IstioIngressGatewayMode = "Internal"
)

func PossibleValuesForIstioIngressGatewayMode() []string {
	return []string{
		string(IstioIngressGatewayModeExternal),
		string(IstioIngressGatewayModeInternal),
	}
}

func parseIstioIngressGatewayMode(input string) (*IstioIngressGatewayMode, error) {
	vals := map[string]IstioIngressGatewayMode{
		"external": IstioIngressGatewayModeExternal,
		"internal": IstioIngressGatewayModeInternal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IstioIngressGatewayMode(input)
	return &out, nil
}

type KeyVaultNetworkAccessTypes string

const (
	KeyVaultNetworkAccessTypesPrivate KeyVaultNetworkAccessTypes = "Private"
	KeyVaultNetworkAccessTypesPublic  KeyVaultNetworkAccessTypes = "Public"
)

func PossibleValuesForKeyVaultNetworkAccessTypes() []string {
	return []string{
		string(KeyVaultNetworkAccessTypesPrivate),
		string(KeyVaultNetworkAccessTypesPublic),
	}
}

func parseKeyVaultNetworkAccessTypes(input string) (*KeyVaultNetworkAccessTypes, error) {
	vals := map[string]KeyVaultNetworkAccessTypes{
		"private": KeyVaultNetworkAccessTypesPrivate,
		"public":  KeyVaultNetworkAccessTypesPublic,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KeyVaultNetworkAccessTypes(input)
	return &out, nil
}

type KubeletDiskType string

const (
	KubeletDiskTypeOS        KubeletDiskType = "OS"
	KubeletDiskTypeTemporary KubeletDiskType = "Temporary"
)

func PossibleValuesForKubeletDiskType() []string {
	return []string{
		string(KubeletDiskTypeOS),
		string(KubeletDiskTypeTemporary),
	}
}

func parseKubeletDiskType(input string) (*KubeletDiskType, error) {
	vals := map[string]KubeletDiskType{
		"os":        KubeletDiskTypeOS,
		"temporary": KubeletDiskTypeTemporary,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KubeletDiskType(input)
	return &out, nil
}

type KubernetesSupportPlan string

const (
	KubernetesSupportPlanAKSLongTermSupport KubernetesSupportPlan = "AKSLongTermSupport"
	KubernetesSupportPlanKubernetesOfficial KubernetesSupportPlan = "KubernetesOfficial"
)

func PossibleValuesForKubernetesSupportPlan() []string {
	return []string{
		string(KubernetesSupportPlanAKSLongTermSupport),
		string(KubernetesSupportPlanKubernetesOfficial),
	}
}

func parseKubernetesSupportPlan(input string) (*KubernetesSupportPlan, error) {
	vals := map[string]KubernetesSupportPlan{
		"akslongtermsupport": KubernetesSupportPlanAKSLongTermSupport,
		"kubernetesofficial": KubernetesSupportPlanKubernetesOfficial,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KubernetesSupportPlan(input)
	return &out, nil
}

type LicenseType string

const (
	LicenseTypeNone          LicenseType = "None"
	LicenseTypeWindowsServer LicenseType = "Windows_Server"
)

func PossibleValuesForLicenseType() []string {
	return []string{
		string(LicenseTypeNone),
		string(LicenseTypeWindowsServer),
	}
}

func parseLicenseType(input string) (*LicenseType, error) {
	vals := map[string]LicenseType{
		"none":           LicenseTypeNone,
		"windows_server": LicenseTypeWindowsServer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LicenseType(input)
	return &out, nil
}

type LoadBalancerSku string

const (
	LoadBalancerSkuBasic    LoadBalancerSku = "basic"
	LoadBalancerSkuStandard LoadBalancerSku = "standard"
)

func PossibleValuesForLoadBalancerSku() []string {
	return []string{
		string(LoadBalancerSkuBasic),
		string(LoadBalancerSkuStandard),
	}
}

func parseLoadBalancerSku(input string) (*LoadBalancerSku, error) {
	vals := map[string]LoadBalancerSku{
		"basic":    LoadBalancerSkuBasic,
		"standard": LoadBalancerSkuStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LoadBalancerSku(input)
	return &out, nil
}

type ManagedClusterPodIdentityProvisioningState string

const (
	ManagedClusterPodIdentityProvisioningStateAssigned  ManagedClusterPodIdentityProvisioningState = "Assigned"
	ManagedClusterPodIdentityProvisioningStateCanceled  ManagedClusterPodIdentityProvisioningState = "Canceled"
	ManagedClusterPodIdentityProvisioningStateDeleting  ManagedClusterPodIdentityProvisioningState = "Deleting"
	ManagedClusterPodIdentityProvisioningStateFailed    ManagedClusterPodIdentityProvisioningState = "Failed"
	ManagedClusterPodIdentityProvisioningStateSucceeded ManagedClusterPodIdentityProvisioningState = "Succeeded"
	ManagedClusterPodIdentityProvisioningStateUpdating  ManagedClusterPodIdentityProvisioningState = "Updating"
)

func PossibleValuesForManagedClusterPodIdentityProvisioningState() []string {
	return []string{
		string(ManagedClusterPodIdentityProvisioningStateAssigned),
		string(ManagedClusterPodIdentityProvisioningStateCanceled),
		string(ManagedClusterPodIdentityProvisioningStateDeleting),
		string(ManagedClusterPodIdentityProvisioningStateFailed),
		string(ManagedClusterPodIdentityProvisioningStateSucceeded),
		string(ManagedClusterPodIdentityProvisioningStateUpdating),
	}
}

func parseManagedClusterPodIdentityProvisioningState(input string) (*ManagedClusterPodIdentityProvisioningState, error) {
	vals := map[string]ManagedClusterPodIdentityProvisioningState{
		"assigned":  ManagedClusterPodIdentityProvisioningStateAssigned,
		"canceled":  ManagedClusterPodIdentityProvisioningStateCanceled,
		"deleting":  ManagedClusterPodIdentityProvisioningStateDeleting,
		"failed":    ManagedClusterPodIdentityProvisioningStateFailed,
		"succeeded": ManagedClusterPodIdentityProvisioningStateSucceeded,
		"updating":  ManagedClusterPodIdentityProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedClusterPodIdentityProvisioningState(input)
	return &out, nil
}

type ManagedClusterSKUName string

const (
	ManagedClusterSKUNameBase ManagedClusterSKUName = "Base"
)

func PossibleValuesForManagedClusterSKUName() []string {
	return []string{
		string(ManagedClusterSKUNameBase),
	}
}

func parseManagedClusterSKUName(input string) (*ManagedClusterSKUName, error) {
	vals := map[string]ManagedClusterSKUName{
		"base": ManagedClusterSKUNameBase,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedClusterSKUName(input)
	return &out, nil
}

type ManagedClusterSKUTier string

const (
	ManagedClusterSKUTierFree     ManagedClusterSKUTier = "Free"
	ManagedClusterSKUTierPremium  ManagedClusterSKUTier = "Premium"
	ManagedClusterSKUTierStandard ManagedClusterSKUTier = "Standard"
)

func PossibleValuesForManagedClusterSKUTier() []string {
	return []string{
		string(ManagedClusterSKUTierFree),
		string(ManagedClusterSKUTierPremium),
		string(ManagedClusterSKUTierStandard),
	}
}

func parseManagedClusterSKUTier(input string) (*ManagedClusterSKUTier, error) {
	vals := map[string]ManagedClusterSKUTier{
		"free":     ManagedClusterSKUTierFree,
		"premium":  ManagedClusterSKUTierPremium,
		"standard": ManagedClusterSKUTierStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedClusterSKUTier(input)
	return &out, nil
}

type NetworkDataplane string

const (
	NetworkDataplaneAzure  NetworkDataplane = "azure"
	NetworkDataplaneCilium NetworkDataplane = "cilium"
)

func PossibleValuesForNetworkDataplane() []string {
	return []string{
		string(NetworkDataplaneAzure),
		string(NetworkDataplaneCilium),
	}
}

func parseNetworkDataplane(input string) (*NetworkDataplane, error) {
	vals := map[string]NetworkDataplane{
		"azure":  NetworkDataplaneAzure,
		"cilium": NetworkDataplaneCilium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkDataplane(input)
	return &out, nil
}

type NetworkMode string

const (
	NetworkModeBridge      NetworkMode = "bridge"
	NetworkModeTransparent NetworkMode = "transparent"
)

func PossibleValuesForNetworkMode() []string {
	return []string{
		string(NetworkModeBridge),
		string(NetworkModeTransparent),
	}
}

func parseNetworkMode(input string) (*NetworkMode, error) {
	vals := map[string]NetworkMode{
		"bridge":      NetworkModeBridge,
		"transparent": NetworkModeTransparent,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkMode(input)
	return &out, nil
}

type NetworkPlugin string

const (
	NetworkPluginAzure   NetworkPlugin = "azure"
	NetworkPluginKubenet NetworkPlugin = "kubenet"
	NetworkPluginNone    NetworkPlugin = "none"
)

func PossibleValuesForNetworkPlugin() []string {
	return []string{
		string(NetworkPluginAzure),
		string(NetworkPluginKubenet),
		string(NetworkPluginNone),
	}
}

func parseNetworkPlugin(input string) (*NetworkPlugin, error) {
	vals := map[string]NetworkPlugin{
		"azure":   NetworkPluginAzure,
		"kubenet": NetworkPluginKubenet,
		"none":    NetworkPluginNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkPlugin(input)
	return &out, nil
}

type NetworkPluginMode string

const (
	NetworkPluginModeOverlay NetworkPluginMode = "overlay"
)

func PossibleValuesForNetworkPluginMode() []string {
	return []string{
		string(NetworkPluginModeOverlay),
	}
}

func parseNetworkPluginMode(input string) (*NetworkPluginMode, error) {
	vals := map[string]NetworkPluginMode{
		"overlay": NetworkPluginModeOverlay,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkPluginMode(input)
	return &out, nil
}

type NetworkPolicy string

const (
	NetworkPolicyAzure  NetworkPolicy = "azure"
	NetworkPolicyCalico NetworkPolicy = "calico"
	NetworkPolicyCilium NetworkPolicy = "cilium"
)

func PossibleValuesForNetworkPolicy() []string {
	return []string{
		string(NetworkPolicyAzure),
		string(NetworkPolicyCalico),
		string(NetworkPolicyCilium),
	}
}

func parseNetworkPolicy(input string) (*NetworkPolicy, error) {
	vals := map[string]NetworkPolicy{
		"azure":  NetworkPolicyAzure,
		"calico": NetworkPolicyCalico,
		"cilium": NetworkPolicyCilium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkPolicy(input)
	return &out, nil
}

type NodeOSUpgradeChannel string

const (
	NodeOSUpgradeChannelNodeImage NodeOSUpgradeChannel = "NodeImage"
	NodeOSUpgradeChannelNone      NodeOSUpgradeChannel = "None"
	NodeOSUpgradeChannelUnmanaged NodeOSUpgradeChannel = "Unmanaged"
)

func PossibleValuesForNodeOSUpgradeChannel() []string {
	return []string{
		string(NodeOSUpgradeChannelNodeImage),
		string(NodeOSUpgradeChannelNone),
		string(NodeOSUpgradeChannelUnmanaged),
	}
}

func parseNodeOSUpgradeChannel(input string) (*NodeOSUpgradeChannel, error) {
	vals := map[string]NodeOSUpgradeChannel{
		"nodeimage": NodeOSUpgradeChannelNodeImage,
		"none":      NodeOSUpgradeChannelNone,
		"unmanaged": NodeOSUpgradeChannelUnmanaged,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NodeOSUpgradeChannel(input)
	return &out, nil
}

type OSDiskType string

const (
	OSDiskTypeEphemeral OSDiskType = "Ephemeral"
	OSDiskTypeManaged   OSDiskType = "Managed"
)

func PossibleValuesForOSDiskType() []string {
	return []string{
		string(OSDiskTypeEphemeral),
		string(OSDiskTypeManaged),
	}
}

func parseOSDiskType(input string) (*OSDiskType, error) {
	vals := map[string]OSDiskType{
		"ephemeral": OSDiskTypeEphemeral,
		"managed":   OSDiskTypeManaged,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OSDiskType(input)
	return &out, nil
}

type OSSKU string

const (
	OSSKUAzureLinux            OSSKU = "AzureLinux"
	OSSKUCBLMariner            OSSKU = "CBLMariner"
	OSSKUUbuntu                OSSKU = "Ubuntu"
	OSSKUWindowsTwoZeroOneNine OSSKU = "Windows2019"
	OSSKUWindowsTwoZeroTwoTwo  OSSKU = "Windows2022"
)

func PossibleValuesForOSSKU() []string {
	return []string{
		string(OSSKUAzureLinux),
		string(OSSKUCBLMariner),
		string(OSSKUUbuntu),
		string(OSSKUWindowsTwoZeroOneNine),
		string(OSSKUWindowsTwoZeroTwoTwo),
	}
}

func parseOSSKU(input string) (*OSSKU, error) {
	vals := map[string]OSSKU{
		"azurelinux":  OSSKUAzureLinux,
		"cblmariner":  OSSKUCBLMariner,
		"ubuntu":      OSSKUUbuntu,
		"windows2019": OSSKUWindowsTwoZeroOneNine,
		"windows2022": OSSKUWindowsTwoZeroTwoTwo,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OSSKU(input)
	return &out, nil
}

type OSType string

const (
	OSTypeLinux   OSType = "Linux"
	OSTypeWindows OSType = "Windows"
)

func PossibleValuesForOSType() []string {
	return []string{
		string(OSTypeLinux),
		string(OSTypeWindows),
	}
}

func parseOSType(input string) (*OSType, error) {
	vals := map[string]OSType{
		"linux":   OSTypeLinux,
		"windows": OSTypeWindows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OSType(input)
	return &out, nil
}

type OutboundType string

const (
	OutboundTypeLoadBalancer           OutboundType = "loadBalancer"
	OutboundTypeManagedNATGateway      OutboundType = "managedNATGateway"
	OutboundTypeUserAssignedNATGateway OutboundType = "userAssignedNATGateway"
	OutboundTypeUserDefinedRouting     OutboundType = "userDefinedRouting"
)

func PossibleValuesForOutboundType() []string {
	return []string{
		string(OutboundTypeLoadBalancer),
		string(OutboundTypeManagedNATGateway),
		string(OutboundTypeUserAssignedNATGateway),
		string(OutboundTypeUserDefinedRouting),
	}
}

func parseOutboundType(input string) (*OutboundType, error) {
	vals := map[string]OutboundType{
		"loadbalancer":           OutboundTypeLoadBalancer,
		"managednatgateway":      OutboundTypeManagedNATGateway,
		"userassignednatgateway": OutboundTypeUserAssignedNATGateway,
		"userdefinedrouting":     OutboundTypeUserDefinedRouting,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OutboundType(input)
	return &out, nil
}

type Protocol string

const (
	ProtocolTCP Protocol = "TCP"
	ProtocolUDP Protocol = "UDP"
)

func PossibleValuesForProtocol() []string {
	return []string{
		string(ProtocolTCP),
		string(ProtocolUDP),
	}
}

func parseProtocol(input string) (*Protocol, error) {
	vals := map[string]Protocol{
		"tcp": ProtocolTCP,
		"udp": ProtocolUDP,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Protocol(input)
	return &out, nil
}

type PublicNetworkAccess string

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = "Disabled"
	PublicNetworkAccessEnabled  PublicNetworkAccess = "Enabled"
)

func PossibleValuesForPublicNetworkAccess() []string {
	return []string{
		string(PublicNetworkAccessDisabled),
		string(PublicNetworkAccessEnabled),
	}
}

func parsePublicNetworkAccess(input string) (*PublicNetworkAccess, error) {
	vals := map[string]PublicNetworkAccess{
		"disabled": PublicNetworkAccessDisabled,
		"enabled":  PublicNetworkAccessEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PublicNetworkAccess(input)
	return &out, nil
}

type ResourceIdentityType string

const (
	ResourceIdentityTypeNone           ResourceIdentityType = "None"
	ResourceIdentityTypeSystemAssigned ResourceIdentityType = "SystemAssigned"
	ResourceIdentityTypeUserAssigned   ResourceIdentityType = "UserAssigned"
)

func PossibleValuesForResourceIdentityType() []string {
	return []string{
		string(ResourceIdentityTypeNone),
		string(ResourceIdentityTypeSystemAssigned),
		string(ResourceIdentityTypeUserAssigned),
	}
}

func parseResourceIdentityType(input string) (*ResourceIdentityType, error) {
	vals := map[string]ResourceIdentityType{
		"none":           ResourceIdentityTypeNone,
		"systemassigned": ResourceIdentityTypeSystemAssigned,
		"userassigned":   ResourceIdentityTypeUserAssigned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResourceIdentityType(input)
	return &out, nil
}

type ScaleDownMode string

const (
	ScaleDownModeDeallocate ScaleDownMode = "Deallocate"
	ScaleDownModeDelete     ScaleDownMode = "Delete"
)

func PossibleValuesForScaleDownMode() []string {
	return []string{
		string(ScaleDownModeDeallocate),
		string(ScaleDownModeDelete),
	}
}

func parseScaleDownMode(input string) (*ScaleDownMode, error) {
	vals := map[string]ScaleDownMode{
		"deallocate": ScaleDownModeDeallocate,
		"delete":     ScaleDownModeDelete,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScaleDownMode(input)
	return &out, nil
}

type ScaleSetEvictionPolicy string

const (
	ScaleSetEvictionPolicyDeallocate ScaleSetEvictionPolicy = "Deallocate"
	ScaleSetEvictionPolicyDelete     ScaleSetEvictionPolicy = "Delete"
)

func PossibleValuesForScaleSetEvictionPolicy() []string {
	return []string{
		string(ScaleSetEvictionPolicyDeallocate),
		string(ScaleSetEvictionPolicyDelete),
	}
}

func parseScaleSetEvictionPolicy(input string) (*ScaleSetEvictionPolicy, error) {
	vals := map[string]ScaleSetEvictionPolicy{
		"deallocate": ScaleSetEvictionPolicyDeallocate,
		"delete":     ScaleSetEvictionPolicyDelete,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScaleSetEvictionPolicy(input)
	return &out, nil
}

type ScaleSetPriority string

const (
	ScaleSetPriorityRegular ScaleSetPriority = "Regular"
	ScaleSetPrioritySpot    ScaleSetPriority = "Spot"
)

func PossibleValuesForScaleSetPriority() []string {
	return []string{
		string(ScaleSetPriorityRegular),
		string(ScaleSetPrioritySpot),
	}
}

func parseScaleSetPriority(input string) (*ScaleSetPriority, error) {
	vals := map[string]ScaleSetPriority{
		"regular": ScaleSetPriorityRegular,
		"spot":    ScaleSetPrioritySpot,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScaleSetPriority(input)
	return &out, nil
}

type ServiceMeshMode string

const (
	ServiceMeshModeDisabled ServiceMeshMode = "Disabled"
	ServiceMeshModeIstio    ServiceMeshMode = "Istio"
)

func PossibleValuesForServiceMeshMode() []string {
	return []string{
		string(ServiceMeshModeDisabled),
		string(ServiceMeshModeIstio),
	}
}

func parseServiceMeshMode(input string) (*ServiceMeshMode, error) {
	vals := map[string]ServiceMeshMode{
		"disabled": ServiceMeshModeDisabled,
		"istio":    ServiceMeshModeIstio,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServiceMeshMode(input)
	return &out, nil
}

type UpgradeChannel string

const (
	UpgradeChannelNodeNegativeimage UpgradeChannel = "node-image"
	UpgradeChannelNone              UpgradeChannel = "none"
	UpgradeChannelPatch             UpgradeChannel = "patch"
	UpgradeChannelRapid             UpgradeChannel = "rapid"
	UpgradeChannelStable            UpgradeChannel = "stable"
)

func PossibleValuesForUpgradeChannel() []string {
	return []string{
		string(UpgradeChannelNodeNegativeimage),
		string(UpgradeChannelNone),
		string(UpgradeChannelPatch),
		string(UpgradeChannelRapid),
		string(UpgradeChannelStable),
	}
}

func parseUpgradeChannel(input string) (*UpgradeChannel, error) {
	vals := map[string]UpgradeChannel{
		"node-image": UpgradeChannelNodeNegativeimage,
		"none":       UpgradeChannelNone,
		"patch":      UpgradeChannelPatch,
		"rapid":      UpgradeChannelRapid,
		"stable":     UpgradeChannelStable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UpgradeChannel(input)
	return &out, nil
}

type WorkloadRuntime string

const (
	WorkloadRuntimeOCIContainer WorkloadRuntime = "OCIContainer"
	WorkloadRuntimeWasmWasi     WorkloadRuntime = "WasmWasi"
)

func PossibleValuesForWorkloadRuntime() []string {
	return []string{
		string(WorkloadRuntimeOCIContainer),
		string(WorkloadRuntimeWasmWasi),
	}
}

func parseWorkloadRuntime(input string) (*WorkloadRuntime, error) {
	vals := map[string]WorkloadRuntime{
		"ocicontainer": WorkloadRuntimeOCIContainer,
		"wasmwasi":     WorkloadRuntimeWasmWasi,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkloadRuntime(input)
	return &out, nil
}
