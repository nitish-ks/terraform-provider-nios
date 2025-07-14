package ipam

import (
	"context"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-nios/internal/flex"
)

// specialDhcpOptions defines the list of special DHCP options that support use_option flag
var specialDhcpOptions = []string{
	"routers",
	"router-templates",
	"domain-name-servers",
	"domain-name",
	"broadcast-address",
	"broadcast-address-offset",
	"dhcp-lease-time",
	"dhcp6.name-servers",
}

// isSpecialDhcpOption checks if the given option name is a special DHCP option
func isSpecialDhcpOption(name string) bool {
	return slices.Contains(specialDhcpOptions, name)
}

type NetworkcontainerOptionsModel struct {
	Name        types.String `tfsdk:"name"`
	Num         types.Int64  `tfsdk:"num"`
	VendorClass types.String `tfsdk:"vendor_class"`
	Value       types.String `tfsdk:"value"`
	UseOption   types.Bool   `tfsdk:"use_option"`
}

var NetworkcontainerOptionsAttrTypes = map[string]attr.Type{
	"name":         types.StringType,
	"num":          types.Int64Type,
	"vendor_class": types.StringType,
	"value":        types.StringType,
	"use_option":   types.BoolType,
}

var NetworkcontainerOptionsResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Name of the DHCP option.",
		Computed:            true,
	},
	"num": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The code of the DHCP option.",
		Computed:            true,
	},
	"vendor_class": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the space this DHCP option is associated to.",
		Computed:            true,
	},
	"value": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Value of the DHCP option",
		Computed:            true,
	},
	"use_option": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Only applies to special options that are displayed separately from other options and have a use flag. These options are: * routers * router-templates * domain-name-servers * domain-name * broadcast-address * broadcast-address-offset * dhcp-lease-time * dhcp6.name-servers",
		Computed:            true,
	},
}

func ExpandNetworkcontainerOptions(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.NetworkcontainerOptions {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkcontainerOptionsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NetworkcontainerOptionsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.NetworkcontainerOptions {
	if m == nil {
		return nil
	}

	to := &ipam.NetworkcontainerOptions{
		Name:        flex.ExpandStringPointer(m.Name),
		Num:         flex.ExpandInt64Pointer(m.Num),
		VendorClass: flex.ExpandStringPointer(m.VendorClass),
		Value:       flex.ExpandStringPointer(m.Value),
	}

	// Only set UseOption for special DHCP options that support it
	if !m.Name.IsNull() && !m.Name.IsUnknown() {
		optionName := m.Name.ValueString()
		if isSpecialDhcpOption(optionName) {
			// For special options, include the use_option flag
			to.UseOption = flex.ExpandBoolPointer(m.UseOption)
		}
		// For non-special options, don't include UseOption at all to avoid API errors
	}

	return to
}

func FlattenNetworkcontainerOptions(ctx context.Context, from *ipam.NetworkcontainerOptions, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkcontainerOptionsAttrTypes)
	}
	m := NetworkcontainerOptionsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkcontainerOptionsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NetworkcontainerOptionsModel) Flatten(ctx context.Context, from *ipam.NetworkcontainerOptions, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NetworkcontainerOptionsModel{}
	}
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Num = flex.FlattenInt64Pointer(from.Num)
	m.VendorClass = flex.FlattenStringPointer(from.VendorClass)
	m.Value = flex.FlattenStringPointer(from.Value)

	// Handle use_option field based on option type to ensure state consistency
	if from.Name != nil && isSpecialDhcpOption(*from.Name) {
		// For special options, respect the API response
		m.UseOption = types.BoolPointerValue(from.UseOption)
	} else {
		// For non-special options, always set to false since we don't send use_option to API
		// This prevents state inconsistencies when transitioning between option types
		m.UseOption = types.BoolValue(false)
	}
}
