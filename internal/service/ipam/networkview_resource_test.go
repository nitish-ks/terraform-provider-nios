package ipam_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

var readableAttributesForNetworkview = "associated_dns_views,associated_members,cloud_info,comment,ddns_dns_view,ddns_zone_primaries,extattrs,internal_forward_zones,is_default,mgm_private,ms_ad_user_data,name,remote_forward_zones,remote_reverse_zones"

func TestAccNetworkviewResource_basic(t *testing.T) {
	var resourceName = "nios_ipam_networkview.test"
	var v ipam.Networkview

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkviewBasicConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkviewResource_disappears(t *testing.T) {
	resourceName := "nios_ipam_networkview.test"
	var v ipam.Networkview

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNetworkviewDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkviewBasicConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					testAccCheckNetworkviewDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccNetworkviewResource_Ref(t *testing.T) {
	var resourceName = "nios_ipam_networkview.test_ref"
	var v ipam.Networkview

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkviewRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkviewRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkviewResource_CloudInfo(t *testing.T) {
	var resourceName = "nios_ipam_networkview.test_cloud_info"
	var v ipam.Networkview

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkviewCloudInfo("CLOUD_INFO_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_info", "CLOUD_INFO_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkviewCloudInfo("CLOUD_INFO_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_info", "CLOUD_INFO_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkviewResource_Comment(t *testing.T) {
	var resourceName = "nios_ipam_networkview.test_comment"
	var v ipam.Networkview

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkviewComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkviewComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkviewResource_DdnsDnsView(t *testing.T) {
	var resourceName = "nios_ipam_networkview.test_ddns_dns_view"
	var v ipam.Networkview

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkviewDdnsDnsView("DDNS_DNS_VIEW_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_dns_view", "DDNS_DNS_VIEW_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkviewDdnsDnsView("DDNS_DNS_VIEW_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_dns_view", "DDNS_DNS_VIEW_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkviewResource_DdnsZonePrimaries(t *testing.T) {
	var resourceName = "nios_ipam_networkview.test_ddns_zone_primaries"
	var v ipam.Networkview

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkviewDdnsZonePrimaries("DDNS_ZONE_PRIMARIES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_zone_primaries", "DDNS_ZONE_PRIMARIES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkviewDdnsZonePrimaries("DDNS_ZONE_PRIMARIES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_zone_primaries", "DDNS_ZONE_PRIMARIES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkviewResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_ipam_networkview.test_extattrs"
	var v ipam.Networkview

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkviewExtAttrs("EXT_ATTRS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkviewExtAttrs("EXT_ATTRS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkviewResource_FederatedRealms(t *testing.T) {
	var resourceName = "nios_ipam_networkview.test_federated_realms"
	var v ipam.Networkview

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkviewFederatedRealms("FEDERATED_REALMS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "federated_realms", "FEDERATED_REALMS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkviewFederatedRealms("FEDERATED_REALMS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "federated_realms", "FEDERATED_REALMS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkviewResource_InternalForwardZones(t *testing.T) {
	var resourceName = "nios_ipam_networkview.test_internal_forward_zones"
	var v ipam.Networkview

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkviewInternalForwardZones("INTERNAL_FORWARD_ZONES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "internal_forward_zones", "INTERNAL_FORWARD_ZONES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkviewInternalForwardZones("INTERNAL_FORWARD_ZONES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "internal_forward_zones", "INTERNAL_FORWARD_ZONES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkviewResource_MgmPrivate(t *testing.T) {
	var resourceName = "nios_ipam_networkview.test_mgm_private"
	var v ipam.Networkview

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkviewMgmPrivate("MGM_PRIVATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mgm_private", "MGM_PRIVATE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkviewMgmPrivate("MGM_PRIVATE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "mgm_private", "MGM_PRIVATE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkviewResource_MsAdUserData(t *testing.T) {
	var resourceName = "nios_ipam_networkview.test_ms_ad_user_data"
	var v ipam.Networkview

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkviewMsAdUserData("MS_AD_USER_DATA_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_ad_user_data", "MS_AD_USER_DATA_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkviewMsAdUserData("MS_AD_USER_DATA_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_ad_user_data", "MS_AD_USER_DATA_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkviewResource_Name(t *testing.T) {
	var resourceName = "nios_ipam_networkview.test_name"
	var v ipam.Networkview

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkviewName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkviewName("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkviewResource_RemoteForwardZones(t *testing.T) {
	var resourceName = "nios_ipam_networkview.test_remote_forward_zones"
	var v ipam.Networkview

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkviewRemoteForwardZones("REMOTE_FORWARD_ZONES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "remote_forward_zones", "REMOTE_FORWARD_ZONES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkviewRemoteForwardZones("REMOTE_FORWARD_ZONES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "remote_forward_zones", "REMOTE_FORWARD_ZONES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccNetworkviewResource_RemoteReverseZones(t *testing.T) {
	var resourceName = "nios_ipam_networkview.test_remote_reverse_zones"
	var v ipam.Networkview

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccNetworkviewRemoteReverseZones("REMOTE_REVERSE_ZONES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "remote_reverse_zones", "REMOTE_REVERSE_ZONES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccNetworkviewRemoteReverseZones("REMOTE_REVERSE_ZONES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkviewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "remote_reverse_zones", "REMOTE_REVERSE_ZONES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckNetworkviewExists(ctx context.Context, resourceName string, v *ipam.Networkview) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.IPAMAPI.
			NetworkviewAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForNetworkview).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetNetworkviewResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetNetworkviewResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckNetworkviewDestroy(ctx context.Context, v *ipam.Networkview) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.IPAMAPI.
			NetworkviewAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForNetworkview).
			Execute()
		if err != nil {
			if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
				// resource was deleted
				return nil
			}
			return err
		}
		return errors.New("expected to be deleted")
	}
}

func testAccCheckNetworkviewDisappears(ctx context.Context, v *ipam.Networkview) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.IPAMAPI.
			NetworkviewAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccNetworkviewBasicConfig(string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "nios_ipam_networkview" "test" {
}
`)
}

func testAccNetworkviewRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networkview" "test_ref" {
    ref = %q
}
`, ref)
}

func testAccNetworkviewCloudInfo(cloudInfo string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networkview" "test_cloud_info" {
    cloud_info = %q
}
`, cloudInfo)
}

func testAccNetworkviewComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networkview" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccNetworkviewDdnsDnsView(ddnsDnsView string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networkview" "test_ddns_dns_view" {
    ddns_dns_view = %q
}
`, ddnsDnsView)
}

func testAccNetworkviewDdnsZonePrimaries(ddnsZonePrimaries string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networkview" "test_ddns_zone_primaries" {
    ddns_zone_primaries = %q
}
`, ddnsZonePrimaries)
}

func testAccNetworkviewExtAttrs(extAttrs string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networkview" "test_extattrs" {
    extattrs = %q
}
`, extAttrs)
}

func testAccNetworkviewFederatedRealms(federatedRealms string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networkview" "test_federated_realms" {
    federated_realms = %q
}
`, federatedRealms)
}

func testAccNetworkviewInternalForwardZones(internalForwardZones string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networkview" "test_internal_forward_zones" {
    internal_forward_zones = %q
}
`, internalForwardZones)
}

func testAccNetworkviewMgmPrivate(mgmPrivate string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networkview" "test_mgm_private" {
    mgm_private = %q
}
`, mgmPrivate)
}

func testAccNetworkviewMsAdUserData(msAdUserData string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networkview" "test_ms_ad_user_data" {
    ms_ad_user_data = %q
}
`, msAdUserData)
}

func testAccNetworkviewName(name string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networkview" "test_name" {
    name = %q
}
`, name)
}

func testAccNetworkviewRemoteForwardZones(remoteForwardZones string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networkview" "test_remote_forward_zones" {
    remote_forward_zones = %q
}
`, remoteForwardZones)
}

func testAccNetworkviewRemoteReverseZones(remoteReverseZones string) string {
	return fmt.Sprintf(`
resource "nios_ipam_networkview" "test_remote_reverse_zones" {
    remote_reverse_zones = %q
}
`, remoteReverseZones)
}
