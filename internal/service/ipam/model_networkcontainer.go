package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework-validators/boolvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/ipam"
	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/flex"
)

type NetworkcontainerModel struct {
	Ref                              types.String `tfsdk:"ref"`
	Authority                        types.Bool   `tfsdk:"authority"`
	AutoCreateReversezone            types.Bool   `tfsdk:"auto_create_reversezone"`
	Bootfile                         types.String `tfsdk:"bootfile"`
	Bootserver                       types.String `tfsdk:"bootserver"`
	CloudInfo                        types.Object `tfsdk:"cloud_info"`
	Comment                          types.String `tfsdk:"comment"`
	DdnsDomainname                   types.String `tfsdk:"ddns_domainname"`
	DdnsGenerateHostname             types.Bool   `tfsdk:"ddns_generate_hostname"`
	DdnsServerAlwaysUpdates          types.Bool   `tfsdk:"ddns_server_always_updates"`
	DdnsTtl                          types.Int64  `tfsdk:"ddns_ttl"`
	DdnsUpdateFixedAddresses         types.Bool   `tfsdk:"ddns_update_fixed_addresses"`
	DdnsUseOption81                  types.Bool   `tfsdk:"ddns_use_option81"`
	DeleteReason                     types.String `tfsdk:"delete_reason"`
	DenyBootp                        types.Bool   `tfsdk:"deny_bootp"`
	DiscoverNowStatus                types.String `tfsdk:"discover_now_status"`
	DiscoveryBasicPollSettings       types.Object `tfsdk:"discovery_basic_poll_settings"`
	DiscoveryBlackoutSetting         types.Object `tfsdk:"discovery_blackout_setting"`
	DiscoveryEngineType              types.String `tfsdk:"discovery_engine_type"`
	DiscoveryMember                  types.String `tfsdk:"discovery_member"`
	EmailList                        types.List   `tfsdk:"email_list"`
	EnableDdns                       types.Bool   `tfsdk:"enable_ddns"`
	EnableDhcpThresholds             types.Bool   `tfsdk:"enable_dhcp_thresholds"`
	EnableDiscovery                  types.Bool   `tfsdk:"enable_discovery"`
	EnableEmailWarnings              types.Bool   `tfsdk:"enable_email_warnings"`
	EnableImmediateDiscovery         types.Bool   `tfsdk:"enable_immediate_discovery"`
	EnablePxeLeaseTime               types.Bool   `tfsdk:"enable_pxe_lease_time"`
	EnableSnmpWarnings               types.Bool   `tfsdk:"enable_snmp_warnings"`
	EndpointSources                  types.List   `tfsdk:"endpoint_sources"`
	ExtAttrs                         types.Map    `tfsdk:"extattrs"`
	ExtAttrsAll                      types.Map    `tfsdk:"extattrs_all"`
	FederatedRealms                  types.List   `tfsdk:"federated_realms"`
	HighWaterMark                    types.Int64  `tfsdk:"high_water_mark"`
	HighWaterMarkReset               types.Int64  `tfsdk:"high_water_mark_reset"`
	IgnoreDhcpOptionListRequest      types.Bool   `tfsdk:"ignore_dhcp_option_list_request"`
	IgnoreId                         types.String `tfsdk:"ignore_id"`
	IgnoreMacAddresses               types.List   `tfsdk:"ignore_mac_addresses"`
	IpamEmailAddresses               types.List   `tfsdk:"ipam_email_addresses"`
	IpamThresholdSettings            types.Object `tfsdk:"ipam_threshold_settings"`
	IpamTrapSettings                 types.Object `tfsdk:"ipam_trap_settings"`
	LastRirRegistrationUpdateSent    types.Int64  `tfsdk:"last_rir_registration_update_sent"`
	LastRirRegistrationUpdateStatus  types.String `tfsdk:"last_rir_registration_update_status"`
	LeaseScavengeTime                types.Int64  `tfsdk:"lease_scavenge_time"`
	LogicFilterRules                 types.List   `tfsdk:"logic_filter_rules"`
	LowWaterMark                     types.Int64  `tfsdk:"low_water_mark"`
	LowWaterMarkReset                types.Int64  `tfsdk:"low_water_mark_reset"`
	MgmPrivate                       types.Bool   `tfsdk:"mgm_private"`
	MgmPrivateOverridable            types.Bool   `tfsdk:"mgm_private_overridable"`
	MsAdUserData                     types.Object `tfsdk:"ms_ad_user_data"`
	Network                          types.String `tfsdk:"network"`
	FuncCall                         types.Object `tfsdk:"func_call"`
	NetworkContainer                 types.String `tfsdk:"network_container"`
	NetworkView                      types.String `tfsdk:"network_view"`
	Nextserver                       types.String `tfsdk:"nextserver"`
	Options                          types.List   `tfsdk:"options"`
	PortControlBlackoutSetting       types.Object `tfsdk:"port_control_blackout_setting"`
	PxeLeaseTime                     types.Int64  `tfsdk:"pxe_lease_time"`
	RecycleLeases                    types.Bool   `tfsdk:"recycle_leases"`
	RemoveSubnets                    types.Bool   `tfsdk:"remove_subnets"`
	RestartIfNeeded                  types.Bool   `tfsdk:"restart_if_needed"`
	Rir                              types.String `tfsdk:"rir"`
	RirOrganization                  types.String `tfsdk:"rir_organization"`
	RirRegistrationAction            types.String `tfsdk:"rir_registration_action"`
	RirRegistrationStatus            types.String `tfsdk:"rir_registration_status"`
	SamePortControlDiscoveryBlackout types.Bool   `tfsdk:"same_port_control_discovery_blackout"`
	SendRirRequest                   types.Bool   `tfsdk:"send_rir_request"`
	SubscribeSettings                types.Object `tfsdk:"subscribe_settings"`
	Unmanaged                        types.Bool   `tfsdk:"unmanaged"`
	UpdateDnsOnLeaseRenewal          types.Bool   `tfsdk:"update_dns_on_lease_renewal"`
	UseAuthority                     types.Bool   `tfsdk:"use_authority"`
	UseBlackoutSetting               types.Bool   `tfsdk:"use_blackout_setting"`
	UseBootfile                      types.Bool   `tfsdk:"use_bootfile"`
	UseBootserver                    types.Bool   `tfsdk:"use_bootserver"`
	UseDdnsDomainname                types.Bool   `tfsdk:"use_ddns_domainname"`
	UseDdnsGenerateHostname          types.Bool   `tfsdk:"use_ddns_generate_hostname"`
	UseDdnsTtl                       types.Bool   `tfsdk:"use_ddns_ttl"`
	UseDdnsUpdateFixedAddresses      types.Bool   `tfsdk:"use_ddns_update_fixed_addresses"`
	UseDdnsUseOption81               types.Bool   `tfsdk:"use_ddns_use_option81"`
	UseDenyBootp                     types.Bool   `tfsdk:"use_deny_bootp"`
	UseDiscoveryBasicPollingSettings types.Bool   `tfsdk:"use_discovery_basic_polling_settings"`
	UseEmailList                     types.Bool   `tfsdk:"use_email_list"`
	UseEnableDdns                    types.Bool   `tfsdk:"use_enable_ddns"`
	UseEnableDhcpThresholds          types.Bool   `tfsdk:"use_enable_dhcp_thresholds"`
	UseEnableDiscovery               types.Bool   `tfsdk:"use_enable_discovery"`
	UseIgnoreDhcpOptionListRequest   types.Bool   `tfsdk:"use_ignore_dhcp_option_list_request"`
	UseIgnoreId                      types.Bool   `tfsdk:"use_ignore_id"`
	UseIpamEmailAddresses            types.Bool   `tfsdk:"use_ipam_email_addresses"`
	UseIpamThresholdSettings         types.Bool   `tfsdk:"use_ipam_threshold_settings"`
	UseIpamTrapSettings              types.Bool   `tfsdk:"use_ipam_trap_settings"`
	UseLeaseScavengeTime             types.Bool   `tfsdk:"use_lease_scavenge_time"`
	UseLogicFilterRules              types.Bool   `tfsdk:"use_logic_filter_rules"`
	UseMgmPrivate                    types.Bool   `tfsdk:"use_mgm_private"`
	UseNextserver                    types.Bool   `tfsdk:"use_nextserver"`
	UseOptions                       types.Bool   `tfsdk:"use_options"`
	UsePxeLeaseTime                  types.Bool   `tfsdk:"use_pxe_lease_time"`
	UseRecycleLeases                 types.Bool   `tfsdk:"use_recycle_leases"`
	UseSubscribeSettings             types.Bool   `tfsdk:"use_subscribe_settings"`
	UseUpdateDnsOnLeaseRenewal       types.Bool   `tfsdk:"use_update_dns_on_lease_renewal"`
	UseZoneAssociations              types.Bool   `tfsdk:"use_zone_associations"`
	Utilization                      types.Int64  `tfsdk:"utilization"`
	ZoneAssociations                 types.List   `tfsdk:"zone_associations"`
}

var NetworkcontainerAttrTypes = map[string]attr.Type{
	"ref":                                  types.StringType,
	"authority":                            types.BoolType,
	"auto_create_reversezone":              types.BoolType,
	"bootfile":                             types.StringType,
	"bootserver":                           types.StringType,
	"cloud_info":                           types.ObjectType{AttrTypes: NetworkcontainerCloudInfoAttrTypes},
	"comment":                              types.StringType,
	"ddns_domainname":                      types.StringType,
	"ddns_generate_hostname":               types.BoolType,
	"ddns_server_always_updates":           types.BoolType,
	"ddns_ttl":                             types.Int64Type,
	"ddns_update_fixed_addresses":          types.BoolType,
	"ddns_use_option81":                    types.BoolType,
	"delete_reason":                        types.StringType,
	"deny_bootp":                           types.BoolType,
	"discover_now_status":                  types.StringType,
	"discovery_basic_poll_settings":        types.ObjectType{AttrTypes: NetworkcontainerDiscoveryBasicPollSettingsAttrTypes},
	"discovery_blackout_setting":           types.ObjectType{AttrTypes: NetworkcontainerDiscoveryBlackoutSettingAttrTypes},
	"discovery_engine_type":                types.StringType,
	"discovery_member":                     types.StringType,
	"email_list":                           types.ListType{ElemType: types.StringType},
	"enable_ddns":                          types.BoolType,
	"enable_dhcp_thresholds":               types.BoolType,
	"enable_discovery":                     types.BoolType,
	"enable_email_warnings":                types.BoolType,
	"enable_immediate_discovery":           types.BoolType,
	"enable_pxe_lease_time":                types.BoolType,
	"enable_snmp_warnings":                 types.BoolType,
	"endpoint_sources":                     types.ListType{ElemType: types.StringType},
	"extattrs":                             types.MapType{ElemType: types.StringType},
	"extattrs_all":                         types.MapType{ElemType: types.StringType},
	"federated_realms":                     types.ListType{ElemType: types.ObjectType{AttrTypes: NetworkcontainerFederatedRealmsAttrTypes}},
	"high_water_mark":                      types.Int64Type,
	"high_water_mark_reset":                types.Int64Type,
	"ignore_dhcp_option_list_request":      types.BoolType,
	"ignore_id":                            types.StringType,
	"ignore_mac_addresses":                 types.ListType{ElemType: types.StringType},
	"ipam_email_addresses":                 types.ListType{ElemType: types.StringType},
	"ipam_threshold_settings":              types.ObjectType{AttrTypes: NetworkcontainerIpamThresholdSettingsAttrTypes},
	"ipam_trap_settings":                   types.ObjectType{AttrTypes: NetworkcontainerIpamTrapSettingsAttrTypes},
	"last_rir_registration_update_sent":    types.Int64Type,
	"last_rir_registration_update_status":  types.StringType,
	"lease_scavenge_time":                  types.Int64Type,
	"logic_filter_rules":                   types.ListType{ElemType: types.ObjectType{AttrTypes: NetworkcontainerLogicFilterRulesAttrTypes}},
	"low_water_mark":                       types.Int64Type,
	"low_water_mark_reset":                 types.Int64Type,
	"mgm_private":                          types.BoolType,
	"mgm_private_overridable":              types.BoolType,
	"ms_ad_user_data":                      types.ObjectType{AttrTypes: NetworkcontainerMsAdUserDataAttrTypes},
	"network":                              types.StringType,
	"func_call":                            types.ObjectType{AttrTypes: FuncCallAttrTypes},
	"network_container":                    types.StringType,
	"network_view":                         types.StringType,
	"nextserver":                           types.StringType,
	"options":                              types.ListType{ElemType: types.ObjectType{AttrTypes: NetworkcontainerOptionsAttrTypes}},
	"port_control_blackout_setting":        types.ObjectType{AttrTypes: NetworkcontainerPortControlBlackoutSettingAttrTypes},
	"pxe_lease_time":                       types.Int64Type,
	"recycle_leases":                       types.BoolType,
	"remove_subnets":                       types.BoolType,
	"restart_if_needed":                    types.BoolType,
	"rir":                                  types.StringType,
	"rir_organization":                     types.StringType,
	"rir_registration_action":              types.StringType,
	"rir_registration_status":              types.StringType,
	"same_port_control_discovery_blackout": types.BoolType,
	"send_rir_request":                     types.BoolType,
	"subscribe_settings":                   types.ObjectType{AttrTypes: NetworkcontainerSubscribeSettingsAttrTypes},
	"unmanaged":                            types.BoolType,
	"update_dns_on_lease_renewal":          types.BoolType,
	"use_authority":                        types.BoolType,
	"use_blackout_setting":                 types.BoolType,
	"use_bootfile":                         types.BoolType,
	"use_bootserver":                       types.BoolType,
	"use_ddns_domainname":                  types.BoolType,
	"use_ddns_generate_hostname":           types.BoolType,
	"use_ddns_ttl":                         types.BoolType,
	"use_ddns_update_fixed_addresses":      types.BoolType,
	"use_ddns_use_option81":                types.BoolType,
	"use_deny_bootp":                       types.BoolType,
	"use_discovery_basic_polling_settings": types.BoolType,
	"use_email_list":                       types.BoolType,
	"use_enable_ddns":                      types.BoolType,
	"use_enable_dhcp_thresholds":           types.BoolType,
	"use_enable_discovery":                 types.BoolType,
	"use_ignore_dhcp_option_list_request":  types.BoolType,
	"use_ignore_id":                        types.BoolType,
	"use_ipam_email_addresses":             types.BoolType,
	"use_ipam_threshold_settings":          types.BoolType,
	"use_ipam_trap_settings":               types.BoolType,
	"use_lease_scavenge_time":              types.BoolType,
	"use_logic_filter_rules":               types.BoolType,
	"use_mgm_private":                      types.BoolType,
	"use_nextserver":                       types.BoolType,
	"use_options":                          types.BoolType,
	"use_pxe_lease_time":                   types.BoolType,
	"use_recycle_leases":                   types.BoolType,
	"use_subscribe_settings":               types.BoolType,
	"use_update_dns_on_lease_renewal":      types.BoolType,
	"use_zone_associations":                types.BoolType,
	"utilization":                          types.Int64Type,
	"zone_associations":                    types.ListType{ElemType: types.ObjectType{AttrTypes: NetworkcontainerZoneAssociationsAttrTypes}},
}

var NetworkcontainerResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"authority": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Authority for the DHCP network container.",
		Default:             booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_authority")),
		},
	},
	"auto_create_reversezone": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "This flag controls whether reverse zones are automatically created when the network is added.",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"bootfile": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The boot server IPv4 Address or name in FQDN format for the network container. You can specify the name and/or IP address of the boot server that the host needs to boot.",
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_bootfile")),
		},
		Default:  stringdefault.StaticString(""),
		Computed: true,
	},
	"bootserver": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The bootserver address for the network container. You can specify the name and/or IP address of the boot server that the host needs to boot. The boot server IPv4 Address or name in FQDN format.",
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_bootserver")),
		},
		Computed: true,
	},
	"cloud_info": schema.SingleNestedAttribute{
		Attributes:          NetworkcontainerCloudInfoResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Structure containing all cloud API related information for this object.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Comment for the network container; maximum 256 characters.",
		Computed:            true,
	},
	"ddns_domainname": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The dynamic DNS domain name the appliance uses specifically for DDNS updates for this network container.",
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_ddns_domainname")),
		},
		Default:  stringdefault.StaticString(""),
		Computed: true,
	},
	"ddns_generate_hostname": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If this field is set to True, the DHCP server generates a hostname and updates DNS with it when the DHCP client request does not contain a hostname.",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_ddns_generate_hostname")),
		},
	},
	"ddns_server_always_updates": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "This field controls whether the DHCP server is allowed to update DNS, regardless of the DHCP client requests. Note that changes for this field take effect only if ddns_use_option81 is True.",
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("ddns_use_option81")),
			boolvalidator.AlsoRequires(path.MatchRoot("use_ddns_use_option81")),
		},
	},
	"ddns_ttl": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The DNS update Time to Live (TTL) value of a DHCP network container object. The TTL is a 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached.",
		Computed:            true,
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_ddns_ttl")),
		},
	},
	"ddns_update_fixed_addresses": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "By default, the DHCP server does not update DNS when it allocates a fixed address to a client. You can configure the DHCP server to update the A and PTR records of a client with a fixed address. When this feature is enabled and the DHCP server adds A and PTR records for a fixed address, the DHCP server never discards the records.",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_ddns_update_fixed_addresses")),
		},
	},
	"ddns_use_option81": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "The support for DHCP Option 81 at the network container level.",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_ddns_use_option81")),
		},
	},
	"delete_reason": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The reason for deleting the RIR registration request.",
	},
	"deny_bootp": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If set to True, BOOTP settings are disabled and BOOTP requests will be denied.",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_deny_bootp")),
		},
	},
	"discover_now_status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Discover now status for this network container.",
		Default:             stringdefault.StaticString("NONE"),
	},
	"discovery_basic_poll_settings": schema.SingleNestedAttribute{
		Attributes: NetworkcontainerDiscoveryBasicPollSettingsResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Validators: []validator.Object{
			objectvalidator.AlsoRequires(path.MatchRoot("use_discovery_basic_polling_settings")),
		},
		MarkdownDescription: "The discovery basic poll settings for this network container.",
	},
	"discovery_blackout_setting": schema.SingleNestedAttribute{
		Attributes: NetworkcontainerDiscoveryBlackoutSettingResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Validators: []validator.Object{
			objectvalidator.AlsoRequires(path.MatchRoot("use_blackout_setting")),
		},
		MarkdownDescription: "The discovery blackout setting for this network container.",
	},
	"discovery_engine_type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The network discovery engine type.",
		Default:             stringdefault.StaticString("NONE"),
	},
	"discovery_member": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The member that will run discovery for this network container.",
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_enable_discovery")),
		},
		Computed: true,
	},
	"email_list": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The e-mail lists to which the appliance sends DHCP threshold alarm e-mail messages.",
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
			listvalidator.AlsoRequires(path.MatchRoot("use_email_list")),
		},
	},
	"enable_ddns": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "The dynamic DNS updates flag of a DHCP network container object. If set to True, the DHCP server sends DDNS updates to DNS servers in the same Grid, and to external DNS servers.",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_enable_ddns")),
		},
	},
	"enable_dhcp_thresholds": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if DHCP thresholds are enabled for the network container.",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_enable_dhcp_thresholds")),
		},
	},
	"enable_discovery": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether a discovery is enabled or not for this network container. When this is set to False, the network container discovery is disabled.",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_enable_discovery")),
		},
	},
	"enable_email_warnings": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if DHCP threshold warnings are sent through email.",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"enable_immediate_discovery": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the discovery for the network container should be immediately enabled.",
	},
	"enable_pxe_lease_time": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Set this to True if you want the DHCP server to use a different lease time for PXE clients.",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"enable_snmp_warnings": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if DHCP threshold warnings are send through SNMP.",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"endpoint_sources": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The endpoints that provides data for the DHCP Network Container object.",
	},
	"extattrs": schema.MapAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object.",
		ElementType:         types.StringType,
		Default:             mapdefault.StaticValue(types.MapNull(types.StringType)),
	},
	"extattrs_all": schema.MapAttribute{
		Computed:            true,
		MarkdownDescription: "Extensible attributes associated with the object , including default attributes.",
		ElementType:         types.StringType,
	},
	"federated_realms": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NetworkcontainerFederatedRealmsResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "This field contains the federated realms associated to this network container.",
	},
	"high_water_mark": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The percentage of DHCP network container usage threshold above which network container usage is not expected and may warrant your attention. When the high watermark is reached, the Infoblox appliance generates a syslog message and sends a warning (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100.",
		Computed:            true,
	},
	"high_water_mark_reset": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The percentage of DHCP network container usage below which the corresponding SNMP trap is reset. A number that specifies the percentage of allocated addresses. The range is from 1 to 100. The high watermark reset value must be lower than the high watermark value.",
		Computed:            true,
	},
	"ignore_dhcp_option_list_request": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If this field is set to False, the appliance returns all DHCP options the client is eligible to receive, rather than only the list of options the client has requested.",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_ignore_dhcp_option_list_request")),
		},
	},
	"ignore_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Indicates whether the appliance will ignore DHCP client IDs or MAC addresses.",
		Computed:            true,
		Default:             stringdefault.StaticString("NONE"),
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_bootfile")),
			stringvalidator.OneOf("CLIENT", "MACADDR", "NONE"),
		},
	},
	"ignore_mac_addresses": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "A list of MAC addresses the appliance will ignore.",
	},
	"ipam_email_addresses": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The e-mail lists to which the appliance sends IPAM threshold alarm e-mail messages.",
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_ipam_email_addresses")),
		},
	},
	"ipam_threshold_settings": schema.SingleNestedAttribute{
		Attributes: NetworkcontainerIpamThresholdSettingsResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Validators: []validator.Object{
			objectvalidator.AlsoRequires(path.MatchRoot("use_ipam_threshold_settings")),
		},
		MarkdownDescription: "The IPAM threshold settings for this network container.",
	},
	"ipam_trap_settings": schema.SingleNestedAttribute{
		Attributes: NetworkcontainerIpamTrapSettingsResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Validators: []validator.Object{
			objectvalidator.AlsoRequires(path.MatchRoot("use_ipam_trap_settings")),
		},
		MarkdownDescription: "The IPAM trap settings for this network container.",
	},
	"last_rir_registration_update_sent": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The timestamp when the last RIR registration update was sent.",
	},
	"last_rir_registration_update_status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Last RIR registration update status.",
	},
	"lease_scavenge_time": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "An integer that specifies the period of time (in seconds) that frees and backs up leases remained in the database before they are automatically deleted. To disable lease scavenging, set the parameter to -1. The minimum positive value must be greater than 86400 seconds (1 day).",
		Computed:            true,
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_lease_scavenge_time")),
			int64validator.Any(
				int64validator.OneOf(-1),
				int64validator.Between(86400, 2147472000),
			),
		},
	},
	"logic_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NetworkcontainerLogicFilterRulesResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "This field contains the logic filters to be applied on the this network container. This list corresponds to the match rules that are written to the dhcpd configuration file.",
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_logic_filter_rules")),
		},
	},
	"low_water_mark": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The percentage of DHCP network container usage below which the Infoblox appliance generates a syslog message and sends a warning (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100.",
		Computed:            true,
		Validators: []validator.Int64{
			int64validator.Any(
				int64validator.Between(0, 100),
			),
		},
	},
	"low_water_mark_reset": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The percentage of DHCP network container usage threshold below which network container usage is not expected and may warrant your attention. When the low watermark is crossed, the Infoblox appliance generates a syslog message and sends a warning (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100. The low watermark reset value must be higher than the low watermark value.",
		Computed:            true,
		Default:             int64default.StaticInt64(10),
		Validators: []validator.Int64{
			int64validator.Any(
				int64validator.Between(1, 100),
			),
		},
	},
	"mgm_private": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "This field controls whether this object is synchronized with the Multi-Grid Master. If this field is set to True, objects are not synchronized.",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_mgm_private")),
		},
	},
	"mgm_private_overridable": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "This field is assumed to be True unless filled by any conforming objects, such as Network, IPv6 Network, Network Container, IPv6 Network Container, and Network View. This value is set to False if mgm_private is set to True in the parent object.",
		Default:             booldefault.StaticBool(true),
	},
	"ms_ad_user_data": schema.SingleNestedAttribute{
		Attributes: NetworkcontainerMsAdUserDataResourceSchemaAttributes,
		Optional:   true,
	},
	"network": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The IPv4 Address of the record.",
		Computed:            true,
	},
	"func_call": schema.SingleNestedAttribute{
		Computed:            true,
		Attributes:          FuncCallResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Function call to be executed.",
	},
	"network_container": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The network container to which this network belongs, if any.",
	},
	"network_view": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the network view in which this network resides.",
		Computed:            true,
		Default:             stringdefault.StaticString("default"),
	},
	"nextserver": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name in FQDN and/or IPv4 Address of the next server that the host needs to boot.",
		Validators: []validator.String{
			stringvalidator.AlsoRequires(path.MatchRoot("use_nextserver")),
		},
		Computed: true,
	},
	"options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NetworkcontainerOptionsResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object.",
		Computed:            true,
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_options")),
		},
	},
	"port_control_blackout_setting": schema.SingleNestedAttribute{
		Attributes: NetworkcontainerPortControlBlackoutSettingResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Validators: []validator.Object{
			objectvalidator.AlsoRequires(path.MatchRoot("use_blackout_setting")),
		},
		MarkdownDescription: "The port control blackout setting for this network container.",
	},
	"pxe_lease_time": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The PXE lease time value of a DHCP Network container object. Some hosts use PXE (Preboot Execution Environment) to boot remotely from a server. To better manage your IP resources, set a different lease time for PXE boot requests. You can configure the DHCP server to allocate an IP address with a shorter lease time to hosts that send PXE boot requests, so IP addresses are not leased longer than necessary. A 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached.",
		Computed:            true,
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRoot("use_pxe_lease_time")),
			int64validator.Any(
				int64validator.Between(0, 4294967295),
			),
		},
	},
	"recycle_leases": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If the field is set to True, the leases are kept in the Recycle Bin until one week after expiration. Otherwise, the leases are permanently deleted.",
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_recycle_leases")),
		},
	},
	"remove_subnets": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Remove subnets delete option. Determines whether all child objects should be removed alongside with the network container or child objects should be assigned to another parental container. By default child objects are deleted with the network container.",
	},
	"restart_if_needed": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Restarts the member service.",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"rir": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The registry (RIR) that allocated the network container address space.",
		Default:             stringdefault.StaticString("NONE"),
	},
	"rir_organization": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The RIR organization assoicated with the network container.",
		Computed:            true,
	},
	"rir_registration_action": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The RIR registration action.",
		Computed:            true,
	},
	"rir_registration_status": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The registration status of the network container in RIR.",
		Computed:            true,
		Default:             stringdefault.StaticString("NOT_REGISTERED"),
	},
	"same_port_control_discovery_blackout": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If the field is set to True, the discovery blackout setting will be used for port control blackout setting.",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_blackout_setting")),
		},
	},
	"send_rir_request": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether to send the RIR registration request.",
	},
	"subscribe_settings": schema.SingleNestedAttribute{
		Attributes: NetworkcontainerSubscribeSettingsResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Validators: []validator.Object{
			objectvalidator.AlsoRequires(path.MatchRoot("use_subscribe_settings")),
		},
	},
	"unmanaged": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether the network container is unmanaged or not.",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"update_dns_on_lease_renewal": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "This field controls whether the DHCP server updates DNS when a DHCP lease is renewed.",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRoot("use_update_dns_on_lease_renewal")),
		},
	},
	"use_authority": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: authority",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_blackout_setting": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: discovery_blackout_setting , port_control_blackout_setting, same_port_control_discovery_blackout",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_bootfile": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: bootfile",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_bootserver": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: bootserver",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_ddns_domainname": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ddns_domainname",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_ddns_generate_hostname": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ddns_generate_hostname",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_ddns_ttl": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ddns_ttl",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_ddns_update_fixed_addresses": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ddns_update_fixed_addresses",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_ddns_use_option81": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ddns_use_option81",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_deny_bootp": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: deny_bootp",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_discovery_basic_polling_settings": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: discovery_basic_poll_settings",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_email_list": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: email_list",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_enable_ddns": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: enable_ddns",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_enable_dhcp_thresholds": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: enable_dhcp_thresholds",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_enable_discovery": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: discovery_member , enable_discovery",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_ignore_dhcp_option_list_request": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ignore_dhcp_option_list_request",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_ignore_id": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ignore_id",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_ipam_email_addresses": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ipam_email_addresses",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_ipam_threshold_settings": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ipam_threshold_settings",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_ipam_trap_settings": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: ipam_trap_settings",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_lease_scavenge_time": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: lease_scavenge_time",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_logic_filter_rules": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: logic_filter_rules",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_mgm_private": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: mgm_private",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_nextserver": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: nextserver",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_options": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: options",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_pxe_lease_time": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: pxe_lease_time",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_recycle_leases": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: recycle_leases",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_subscribe_settings": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: subscribe_settings",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_update_dns_on_lease_renewal": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: update_dns_on_lease_renewal",
		Computed:            true,
		Default:             booldefault.StaticBool(false),
	},
	"use_zone_associations": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: zone_associations",
		Computed:            true,
		Default:             booldefault.StaticBool(true),
	},
	"utilization": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The network container utilization in percentage.",
	},
	"zone_associations": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NetworkcontainerZoneAssociationsResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "The list of zones associated with this network.",
		Validators: []validator.List{
			listvalidator.AlsoRequires(path.MatchRoot("use_zone_associations")),
		},
	},
}

func (m *NetworkcontainerModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *ipam.Networkcontainer {
	if m == nil {
		return nil
	}
	to := &ipam.Networkcontainer{
		Ref:                              flex.ExpandStringPointer(m.Ref),
		Authority:                        flex.ExpandBoolPointer(m.Authority),
		Bootfile:                         flex.ExpandStringPointer(m.Bootfile),
		Bootserver:                       flex.ExpandStringPointer(m.Bootserver),
		CloudInfo:                        ExpandNetworkcontainerCloudInfo(ctx, m.CloudInfo, diags),
		Comment:                          flex.ExpandStringPointer(m.Comment),
		DdnsDomainname:                   flex.ExpandStringPointer(m.DdnsDomainname),
		DdnsGenerateHostname:             flex.ExpandBoolPointer(m.DdnsGenerateHostname),
		DdnsServerAlwaysUpdates:          flex.ExpandBoolPointer(m.DdnsServerAlwaysUpdates),
		DdnsTtl:                          flex.ExpandInt64Pointer(m.DdnsTtl),
		DdnsUpdateFixedAddresses:         flex.ExpandBoolPointer(m.DdnsUpdateFixedAddresses),
		DdnsUseOption81:                  flex.ExpandBoolPointer(m.DdnsUseOption81),
		DeleteReason:                     flex.ExpandStringPointer(m.DeleteReason),
		DenyBootp:                        flex.ExpandBoolPointer(m.DenyBootp),
		DiscoveryBasicPollSettings:       ExpandNetworkcontainerDiscoveryBasicPollSettings(ctx, m.DiscoveryBasicPollSettings, diags),
		DiscoveryBlackoutSetting:         ExpandNetworkcontainerDiscoveryBlackoutSetting(ctx, m.DiscoveryBlackoutSetting, diags),
		DiscoveryMember:                  flex.ExpandStringPointer(m.DiscoveryMember),
		EmailList:                        flex.ExpandFrameworkListString(ctx, m.EmailList, diags),
		EnableDdns:                       flex.ExpandBoolPointer(m.EnableDdns),
		EnableDhcpThresholds:             flex.ExpandBoolPointer(m.EnableDhcpThresholds),
		EnableDiscovery:                  flex.ExpandBoolPointer(m.EnableDiscovery),
		EnableEmailWarnings:              flex.ExpandBoolPointer(m.EnableEmailWarnings),
		EnableImmediateDiscovery:         flex.ExpandBoolPointer(m.EnableImmediateDiscovery),
		EnablePxeLeaseTime:               flex.ExpandBoolPointer(m.EnablePxeLeaseTime),
		EnableSnmpWarnings:               flex.ExpandBoolPointer(m.EnableSnmpWarnings),
		ExtAttrs:                         ExpandExtAttr(ctx, m.ExtAttrs, diags),
		FederatedRealms:                  flex.ExpandFrameworkListNestedBlock(ctx, m.FederatedRealms, diags, ExpandNetworkcontainerFederatedRealms),
		HighWaterMark:                    flex.ExpandInt64Pointer(m.HighWaterMark),
		HighWaterMarkReset:               flex.ExpandInt64Pointer(m.HighWaterMarkReset),
		IgnoreDhcpOptionListRequest:      flex.ExpandBoolPointer(m.IgnoreDhcpOptionListRequest),
		IgnoreId:                         flex.ExpandStringPointer(m.IgnoreId),
		IgnoreMacAddresses:               flex.ExpandFrameworkListString(ctx, m.IgnoreMacAddresses, diags),
		IpamEmailAddresses:               flex.ExpandFrameworkListString(ctx, m.IpamEmailAddresses, diags),
		IpamThresholdSettings:            ExpandNetworkcontainerIpamThresholdSettings(ctx, m.IpamThresholdSettings, diags),
		IpamTrapSettings:                 ExpandNetworkcontainerIpamTrapSettings(ctx, m.IpamTrapSettings, diags),
		LeaseScavengeTime:                flex.ExpandInt64Pointer(m.LeaseScavengeTime),
		LogicFilterRules:                 flex.ExpandFrameworkListNestedBlock(ctx, m.LogicFilterRules, diags, ExpandNetworkcontainerLogicFilterRules),
		LowWaterMark:                     flex.ExpandInt64Pointer(m.LowWaterMark),
		LowWaterMarkReset:                flex.ExpandInt64Pointer(m.LowWaterMarkReset),
		MgmPrivate:                       flex.ExpandBoolPointer(m.MgmPrivate),
		MsAdUserData:                     ExpandNetworkcontainerMsAdUserData(ctx, m.MsAdUserData, diags),
		Network:                          ExpandNetworkcontainerNetwork(m.Network),
		FuncCall:                         ExpandFuncCall(ctx, m.FuncCall, diags),
		Nextserver:                       flex.ExpandStringPointer(m.Nextserver),
		Options:                          flex.ExpandFrameworkListNestedBlock(ctx, m.Options, diags, ExpandNetworkcontainerOptions),
		PortControlBlackoutSetting:       ExpandNetworkcontainerPortControlBlackoutSetting(ctx, m.PortControlBlackoutSetting, diags),
		PxeLeaseTime:                     flex.ExpandInt64Pointer(m.PxeLeaseTime),
		RecycleLeases:                    flex.ExpandBoolPointer(m.RecycleLeases),
		RemoveSubnets:                    flex.ExpandBoolPointer(m.RemoveSubnets),
		RestartIfNeeded:                  flex.ExpandBoolPointer(m.RestartIfNeeded),
		RirOrganization:                  flex.ExpandStringPointer(m.RirOrganization),
		RirRegistrationAction:            flex.ExpandStringPointer(m.RirRegistrationAction),
		RirRegistrationStatus:            flex.ExpandStringPointer(m.RirRegistrationStatus),
		SamePortControlDiscoveryBlackout: flex.ExpandBoolPointer(m.SamePortControlDiscoveryBlackout),
		SendRirRequest:                   flex.ExpandBoolPointer(m.SendRirRequest),
		SubscribeSettings:                ExpandNetworkcontainerSubscribeSettings(ctx, m.SubscribeSettings, diags),
		Unmanaged:                        flex.ExpandBoolPointer(m.Unmanaged),
		UpdateDnsOnLeaseRenewal:          flex.ExpandBoolPointer(m.UpdateDnsOnLeaseRenewal),
		UseAuthority:                     flex.ExpandBoolPointer(m.UseAuthority),
		UseBlackoutSetting:               flex.ExpandBoolPointer(m.UseBlackoutSetting),
		UseBootfile:                      flex.ExpandBoolPointer(m.UseBootfile),
		UseBootserver:                    flex.ExpandBoolPointer(m.UseBootserver),
		UseDdnsDomainname:                flex.ExpandBoolPointer(m.UseDdnsDomainname),
		UseDdnsGenerateHostname:          flex.ExpandBoolPointer(m.UseDdnsGenerateHostname),
		UseDdnsTtl:                       flex.ExpandBoolPointer(m.UseDdnsTtl),
		UseDdnsUpdateFixedAddresses:      flex.ExpandBoolPointer(m.UseDdnsUpdateFixedAddresses),
		UseDdnsUseOption81:               flex.ExpandBoolPointer(m.UseDdnsUseOption81),
		UseDenyBootp:                     flex.ExpandBoolPointer(m.UseDenyBootp),
		UseDiscoveryBasicPollingSettings: flex.ExpandBoolPointer(m.UseDiscoveryBasicPollingSettings),
		UseEmailList:                     flex.ExpandBoolPointer(m.UseEmailList),
		UseEnableDdns:                    flex.ExpandBoolPointer(m.UseEnableDdns),
		UseEnableDhcpThresholds:          flex.ExpandBoolPointer(m.UseEnableDhcpThresholds),
		UseEnableDiscovery:               flex.ExpandBoolPointer(m.UseEnableDiscovery),
		UseIgnoreDhcpOptionListRequest:   flex.ExpandBoolPointer(m.UseIgnoreDhcpOptionListRequest),
		UseIgnoreId:                      flex.ExpandBoolPointer(m.UseIgnoreId),
		UseIpamEmailAddresses:            flex.ExpandBoolPointer(m.UseIpamEmailAddresses),
		UseIpamThresholdSettings:         flex.ExpandBoolPointer(m.UseIpamThresholdSettings),
		UseIpamTrapSettings:              flex.ExpandBoolPointer(m.UseIpamTrapSettings),
		UseLeaseScavengeTime:             flex.ExpandBoolPointer(m.UseLeaseScavengeTime),
		UseLogicFilterRules:              flex.ExpandBoolPointer(m.UseLogicFilterRules),
		UseMgmPrivate:                    flex.ExpandBoolPointer(m.UseMgmPrivate),
		UseNextserver:                    flex.ExpandBoolPointer(m.UseNextserver),
		UseOptions:                       flex.ExpandBoolPointer(m.UseOptions),
		UsePxeLeaseTime:                  flex.ExpandBoolPointer(m.UsePxeLeaseTime),
		UseRecycleLeases:                 flex.ExpandBoolPointer(m.UseRecycleLeases),
		UseSubscribeSettings:             flex.ExpandBoolPointer(m.UseSubscribeSettings),
		UseUpdateDnsOnLeaseRenewal:       flex.ExpandBoolPointer(m.UseUpdateDnsOnLeaseRenewal),
		UseZoneAssociations:              flex.ExpandBoolPointer(m.UseZoneAssociations),
		ZoneAssociations:                 flex.ExpandFrameworkListNestedBlock(ctx, m.ZoneAssociations, diags, ExpandNetworkcontainerZoneAssociations),
	}
	if isCreate {
		to.NetworkContainer = flex.ExpandStringPointer(m.NetworkContainer)
		to.NetworkView = flex.ExpandStringPointer(m.NetworkView)
		to.Network = ExpandNetworkcontainerNetwork(m.Network)
	}
	return to
}

func FlattenNetworkcontainer(ctx context.Context, from *ipam.Networkcontainer, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkcontainerAttrTypes)
	}
	m := NetworkcontainerModel{}
	m.Flatten(ctx, from, diags)
	m.ExtAttrs = m.ExtAttrsAll
	t, d := types.ObjectValueFrom(ctx, NetworkcontainerAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NetworkcontainerModel) Flatten(ctx context.Context, from *ipam.Networkcontainer, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NetworkcontainerModel{}
	}
	m.Ref = flex.FlattenStringPointer(from.Ref)
	m.Authority = types.BoolPointerValue(from.Authority)
	m.Bootfile = flex.FlattenStringPointer(from.Bootfile)
	m.Bootserver = flex.FlattenStringPointer(from.Bootserver)
	m.CloudInfo = FlattenNetworkcontainerCloudInfo(ctx, from.CloudInfo, diags)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.DdnsDomainname = flex.FlattenStringPointer(from.DdnsDomainname)
	m.DdnsGenerateHostname = types.BoolPointerValue(from.DdnsGenerateHostname)
	m.DdnsServerAlwaysUpdates = types.BoolPointerValue(from.DdnsServerAlwaysUpdates)
	m.DdnsTtl = flex.FlattenInt64Pointer(from.DdnsTtl)
	m.DdnsUpdateFixedAddresses = types.BoolPointerValue(from.DdnsUpdateFixedAddresses)
	m.DdnsUseOption81 = types.BoolPointerValue(from.DdnsUseOption81)
	m.DenyBootp = types.BoolPointerValue(from.DenyBootp)
	m.DiscoverNowStatus = flex.FlattenStringPointer(from.DiscoverNowStatus)
	m.DiscoveryBasicPollSettings = FlattenNetworkcontainerDiscoveryBasicPollSettings(ctx, from.DiscoveryBasicPollSettings, diags)
	m.DiscoveryBlackoutSetting = FlattenNetworkcontainerDiscoveryBlackoutSetting(ctx, from.DiscoveryBlackoutSetting, diags)
	m.DiscoveryEngineType = flex.FlattenStringPointer(from.DiscoveryEngineType)
	m.DiscoveryMember = flex.FlattenStringPointer(from.DiscoveryMember)
	m.EmailList = flex.FlattenFrameworkListString(ctx, from.EmailList, diags)
	m.EnableDdns = types.BoolPointerValue(from.EnableDdns)
	m.EnableDhcpThresholds = types.BoolPointerValue(from.EnableDhcpThresholds)
	m.EnableDiscovery = types.BoolPointerValue(from.EnableDiscovery)
	m.EnableEmailWarnings = types.BoolPointerValue(from.EnableEmailWarnings)
	m.EnableImmediateDiscovery = types.BoolPointerValue(from.EnableImmediateDiscovery)
	m.EnablePxeLeaseTime = types.BoolPointerValue(from.EnablePxeLeaseTime)
	m.EnableSnmpWarnings = types.BoolPointerValue(from.EnableSnmpWarnings)
	m.EndpointSources = flex.FlattenFrameworkListString(ctx, from.EndpointSources, diags)
	m.ExtAttrsAll = FlattenExtAttr(ctx, *from.ExtAttrs, diags)
	m.FederatedRealms = flex.FlattenFrameworkListNestedBlock(ctx, from.FederatedRealms, NetworkcontainerFederatedRealmsAttrTypes, diags, FlattenNetworkcontainerFederatedRealms)
	m.HighWaterMark = flex.FlattenInt64Pointer(from.HighWaterMark)
	m.HighWaterMarkReset = flex.FlattenInt64Pointer(from.HighWaterMarkReset)
	m.IgnoreDhcpOptionListRequest = types.BoolPointerValue(from.IgnoreDhcpOptionListRequest)
	m.IgnoreId = flex.FlattenStringPointer(from.IgnoreId)
	m.IgnoreMacAddresses = flex.FlattenFrameworkListString(ctx, from.IgnoreMacAddresses, diags)
	m.IpamEmailAddresses = flex.FlattenFrameworkListString(ctx, from.IpamEmailAddresses, diags)
	m.IpamThresholdSettings = FlattenNetworkcontainerIpamThresholdSettings(ctx, from.IpamThresholdSettings, diags)
	m.IpamTrapSettings = FlattenNetworkcontainerIpamTrapSettings(ctx, from.IpamTrapSettings, diags)
	m.LastRirRegistrationUpdateSent = flex.FlattenInt64Pointer(from.LastRirRegistrationUpdateSent)
	m.LastRirRegistrationUpdateStatus = flex.FlattenStringPointer(from.LastRirRegistrationUpdateStatus)
	m.LeaseScavengeTime = flex.FlattenInt64Pointer(from.LeaseScavengeTime)
	m.LogicFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.LogicFilterRules, NetworkcontainerLogicFilterRulesAttrTypes, diags, FlattenNetworkcontainerLogicFilterRules)
	m.LowWaterMark = flex.FlattenInt64Pointer(from.LowWaterMark)
	m.LowWaterMarkReset = flex.FlattenInt64Pointer(from.LowWaterMarkReset)
	m.MgmPrivate = types.BoolPointerValue(from.MgmPrivate)
	m.MgmPrivateOverridable = types.BoolPointerValue(from.MgmPrivateOverridable)
	m.MsAdUserData = FlattenNetworkcontainerMsAdUserData(ctx, from.MsAdUserData, diags)
	m.Network = FlattenNetworkcontainerNetwork(from.Network)
	if m.FuncCall.IsNull() || m.FuncCall.IsUnknown() {
		m.FuncCall = FlattenFuncCall(ctx, from.FuncCall, diags)
	}
	m.NetworkContainer = flex.FlattenStringPointer(from.NetworkContainer)
	m.NetworkView = flex.FlattenStringPointer(from.NetworkView)
	m.Nextserver = flex.FlattenStringPointer(from.Nextserver)
	m.Options = flex.FlattenFrameworkListNestedBlock(ctx, from.Options, NetworkcontainerOptionsAttrTypes, diags, FlattenNetworkcontainerOptions)
	m.PortControlBlackoutSetting = FlattenNetworkcontainerPortControlBlackoutSetting(ctx, from.PortControlBlackoutSetting, diags)
	m.PxeLeaseTime = flex.FlattenInt64Pointer(from.PxeLeaseTime)
	m.RecycleLeases = types.BoolPointerValue(from.RecycleLeases)
	m.RemoveSubnets = types.BoolPointerValue(from.RemoveSubnets)
	m.Rir = flex.FlattenStringPointer(from.Rir)
	m.RirOrganization = flex.FlattenStringPointer(from.RirOrganization)
	m.RirRegistrationAction = flex.FlattenStringPointer(from.RirRegistrationAction)
	m.RirRegistrationStatus = flex.FlattenStringPointer(from.RirRegistrationStatus)
	m.SamePortControlDiscoveryBlackout = types.BoolPointerValue(from.SamePortControlDiscoveryBlackout)
	m.SendRirRequest = types.BoolPointerValue(from.SendRirRequest)
	m.SubscribeSettings = FlattenNetworkcontainerSubscribeSettings(ctx, from.SubscribeSettings, diags)
	m.Unmanaged = types.BoolPointerValue(from.Unmanaged)
	m.UpdateDnsOnLeaseRenewal = types.BoolPointerValue(from.UpdateDnsOnLeaseRenewal)
	m.UseAuthority = types.BoolPointerValue(from.UseAuthority)
	m.UseBlackoutSetting = types.BoolPointerValue(from.UseBlackoutSetting)
	m.UseBootfile = types.BoolPointerValue(from.UseBootfile)
	m.UseBootserver = types.BoolPointerValue(from.UseBootserver)
	m.UseDdnsDomainname = types.BoolPointerValue(from.UseDdnsDomainname)
	m.UseDdnsGenerateHostname = types.BoolPointerValue(from.UseDdnsGenerateHostname)
	m.UseDdnsTtl = types.BoolPointerValue(from.UseDdnsTtl)
	m.UseDdnsUpdateFixedAddresses = types.BoolPointerValue(from.UseDdnsUpdateFixedAddresses)
	m.UseDdnsUseOption81 = types.BoolPointerValue(from.UseDdnsUseOption81)
	m.UseDenyBootp = types.BoolPointerValue(from.UseDenyBootp)
	m.UseDiscoveryBasicPollingSettings = types.BoolPointerValue(from.UseDiscoveryBasicPollingSettings)
	m.UseEmailList = types.BoolPointerValue(from.UseEmailList)
	m.UseEnableDdns = types.BoolPointerValue(from.UseEnableDdns)
	m.UseEnableDhcpThresholds = types.BoolPointerValue(from.UseEnableDhcpThresholds)
	m.UseEnableDiscovery = types.BoolPointerValue(from.UseEnableDiscovery)
	m.UseIgnoreDhcpOptionListRequest = types.BoolPointerValue(from.UseIgnoreDhcpOptionListRequest)
	m.UseIgnoreId = types.BoolPointerValue(from.UseIgnoreId)
	m.UseIpamEmailAddresses = types.BoolPointerValue(from.UseIpamEmailAddresses)
	m.UseIpamThresholdSettings = types.BoolPointerValue(from.UseIpamThresholdSettings)
	m.UseIpamTrapSettings = types.BoolPointerValue(from.UseIpamTrapSettings)
	m.UseLeaseScavengeTime = types.BoolPointerValue(from.UseLeaseScavengeTime)
	m.UseLogicFilterRules = types.BoolPointerValue(from.UseLogicFilterRules)
	m.UseMgmPrivate = types.BoolPointerValue(from.UseMgmPrivate)
	m.UseNextserver = types.BoolPointerValue(from.UseNextserver)
	m.UseOptions = types.BoolPointerValue(from.UseOptions)
	m.UsePxeLeaseTime = types.BoolPointerValue(from.UsePxeLeaseTime)
	m.UseRecycleLeases = types.BoolPointerValue(from.UseRecycleLeases)
	m.UseSubscribeSettings = types.BoolPointerValue(from.UseSubscribeSettings)
	m.UseUpdateDnsOnLeaseRenewal = types.BoolPointerValue(from.UseUpdateDnsOnLeaseRenewal)
	m.UseZoneAssociations = types.BoolPointerValue(from.UseZoneAssociations)
	m.Utilization = flex.FlattenInt64Pointer(from.Utilization)
	m.ZoneAssociations = flex.FlattenFrameworkListNestedBlock(ctx, from.ZoneAssociations, NetworkcontainerZoneAssociationsAttrTypes, diags, FlattenNetworkcontainerZoneAssociations)
}

func ExpandNetworkcontainerNetwork(str types.String) *ipam.NetworkcontainerNetwork {
	if str.IsNull() {
		return &ipam.NetworkcontainerNetwork{}
	}
	var m ipam.NetworkcontainerNetwork
	m.String = flex.ExpandStringPointer(str)

	return &m
}

func FlattenNetworkcontainerNetwork(from *ipam.NetworkcontainerNetwork) types.String {
	if from.String == nil {
		return types.StringNull()
	}
	m := flex.FlattenStringPointer(from.String)
	return m
}
