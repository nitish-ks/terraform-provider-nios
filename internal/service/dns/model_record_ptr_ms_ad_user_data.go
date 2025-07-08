package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dns"

	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/flex"
)

type RecordPtrMsAdUserDataModel struct {
	ActiveUsersCount types.Int64 `tfsdk:"active_users_count"`
}

var RecordPtrMsAdUserDataAttrTypes = map[string]attr.Type{
	"active_users_count": types.Int64Type,
}

var RecordPtrMsAdUserDataResourceSchemaAttributes = map[string]schema.Attribute{
	"active_users_count": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The number of active users.",
	},
}

func ExpandRecordPtrMsAdUserData(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns.RecordPtrMsAdUserData {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RecordPtrMsAdUserDataModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *RecordPtrMsAdUserDataModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns.RecordPtrMsAdUserData {
	if m == nil {
		return nil
	}
	to := &dns.RecordPtrMsAdUserData{}
	return to
}

func FlattenRecordPtrMsAdUserData(ctx context.Context, from *dns.RecordPtrMsAdUserData, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordPtrMsAdUserDataAttrTypes)
	}
	m := RecordPtrMsAdUserDataModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RecordPtrMsAdUserDataAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RecordPtrMsAdUserDataModel) Flatten(ctx context.Context, from *dns.RecordPtrMsAdUserData, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RecordPtrMsAdUserDataModel{}
	}
	m.ActiveUsersCount = flex.FlattenInt64Pointer(from.ActiveUsersCount)
}
