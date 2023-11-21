package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/vmware/alb-sdk/go/clients"
)

func TestAVIUserBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIUserConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIUserExists("avi_user.testUser"),
					resource.TestCheckResourceAttr(
						"avi_user.testUser", "name", "test-tf-user"),
					resource.TestCheckResourceAttr(
						"avi_user.testUser", "is_superuser", "true"),
					resource.TestCheckResourceAttr(
						"avi_user.testUser", "local", "true"),
				),
			},
			{
				Config: testAccAVIUserupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIUserExists("avi_user.testUser"),
					resource.TestCheckResourceAttr(
						"avi_user.testUser", "name", "test-tf-user"),
					resource.TestCheckResourceAttr(
						"avi_user.testUser", "is_superuser", "true"),
					resource.TestCheckResourceAttr(
						"avi_user.testUser", "local", "true"),
				),
			},
			{
				ResourceName:      "avi_user.testUser",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVIUserConfig,
			},
		},
	})

}

func testAccCheckAVIUserExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI User ID is set")
		}
		url := strings.SplitN(rs.Primary.ID, "/api", 2)[1]
		uuid := strings.Split(url, "#")[0]
		path := "api" + uuid
		err := conn.Get(path, &obj)
		if err != nil {
			return err
		}
		return nil
	}

}

func testAccCheckAVIUserDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_user" {
			continue
		}
		url := strings.SplitN(rs.Primary.ID, "/api", 2)[1]
		uuid := strings.Split(url, "#")[0]
		path := "api" + uuid
		err := conn.Get(path, &obj)
		if err != nil {
			if strings.Contains(err.Error(), "404") {
				return nil
			}
			return err
		}
		if len(obj.(map[string]interface{})) > 0 {
			return fmt.Errorf("AVI User still exists")
		}
	}
	return nil
}

const testAccAVIUserConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
data "avi_role" "default-system-admin-role" {
    name= "System-Admin"
}
data "avi_useraccountprofile" "default-user-account-profile" {
    name= "Default-User-Account-Profile"
}
resource "avi_user" "testUser" {
	access {
	role_ref = "${data.avi_tenant.default_tenant.id}/role/${element(split("/", data.avi_role.default-system-admin-role.id),5)}"
	tenant_ref = data.avi_tenant.default_tenant.id
	all_tenants = false
}
	password = "pbkdf2_sha256$100000$vwZd950E3jSj$tC/x4hJBolHm2Ki4uVNbMW59ZQcC95/p5UZUWjmTuFs="
	username = "test-tf-user"
	name = "test-tf-user"
	full_name = "System Administrator"
	email = ""
	is_superuser = true
	default_tenant_ref = data.avi_tenant.default_tenant.id
	local = true
	user_profile_ref = data.avi_useraccountprofile.default-user-account-profile.id
}
`

const testAccAVIUserupdatedConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
data "avi_role" "default-system-admin-role" {
    name= "System-Admin"
}
data "avi_useraccountprofile" "default-user-account-profile" {
    name= "Default-User-Account-Profile"
}
resource "avi_user" "testUser" {
	access {
	role_ref = "${data.avi_tenant.default_tenant.id}/role/${element(split("/", data.avi_role.default-system-admin-role.id),5)}"
	tenant_ref = data.avi_tenant.default_tenant.id
	all_tenants = false
}
	password = "pbkdf2_sha256$100000$vwZd950E3jSj$tC/x4hJBolHm2Ki4uVNbMW59ZQcC95/p5UZUWjmTuFs="
	username = "test-tf-user"
	name = "test-tf-user"
	full_name = "System Administrator"
	email = "testaviuser@testaviuser23.com"
	is_superuser = true
	default_tenant_ref = data.avi_tenant.default_tenant.id
	local = true
	user_profile_ref = data.avi_useraccountprofile.default-user-account-profile.id
}
`
