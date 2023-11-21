package avi

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/vmware/alb-sdk/go/clients"
)

func TestAVIApplicationProfileBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAVIApplicationProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIApplicationProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIApplicationProfileExists("avi_applicationprofile.testApplicationProfile"),
					resource.TestCheckResourceAttr(
						"avi_applicationprofile.testApplicationProfile", "name", "test-System-Secure-HTTP-abc"),
					resource.TestCheckResourceAttr(
						"avi_applicationprofile.testApplicationProfile", "preserve_client_ip", "false"),
					resource.TestCheckResourceAttr(
						"avi_applicationprofile.testApplicationProfile", "preserve_client_port", "false"),
				),
			},
			{
				Config: testAccAVIApplicationProfileupdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAVIApplicationProfileExists("avi_applicationprofile.testApplicationProfile"),
					resource.TestCheckResourceAttr(
						"avi_applicationprofile.testApplicationProfile", "name", "test-System-Secure-HTTP-updated"),
					resource.TestCheckResourceAttr(
						"avi_applicationprofile.testApplicationProfile", "preserve_client_ip", "false"),
					resource.TestCheckResourceAttr(
						"avi_applicationprofile.testApplicationProfile", "preserve_client_port", "false"),
				),
			},
			{
				ResourceName:      "avi_applicationprofile.testApplicationProfile",
				ImportState:       true,
				ImportStateVerify: false,
				Config:            testAccAVIApplicationProfileConfig,
			},
		},
	})

}

func testAccCheckAVIApplicationProfileExists(resourcename string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := testAccProvider.Meta().(*clients.AviClient).AviSession
		var obj interface{}
		rs, ok := s.RootModule().Resources[resourcename]
		if !ok {
			return fmt.Errorf("Not found: %s", resourcename)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("No AVI ApplicationProfile ID is set")
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

func testAccCheckAVIApplicationProfileDestroy(s *terraform.State) error {
	conn := testAccProvider.Meta().(*clients.AviClient).AviSession
	var obj interface{}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "avi_applicationprofile" {
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
			return fmt.Errorf("AVI ApplicationProfile still exists")
		}
	}
	return nil
}

const testAccAVIApplicationProfileConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
data "avi_stringgroup" "system_compressiblestringgroup" {
    name= "System-Compressible-Content-Types"
}
data "avi_stringgroup" "system_cacheablestringgroup" {
    name= "System-Cacheable-Resource-Types"
}
data "avi_stringgroup" "system_mobilestringgroup" {
    name = "System-Devices-Mobile"
}
resource "avi_applicationprofile" "testApplicationProfile" {
	name = "test-System-Secure-HTTP-abc"
	type = "APPLICATION_PROFILE_TYPE_HTTP"
	tenant_ref = data.avi_tenant.default_tenant.id
	preserve_client_ip = false
	http_profile {
		max_rps_uri = "0"
		keepalive_header = false
		max_rps_cip_uri = "0"
		x_forwarded_proto_enabled = true
		connection_multiplexing_enabled = true
		websockets_enabled = true
		enable_request_body_buffering = false
		hsts_enabled = true
		compression_profile {
			type = "AUTO_COMPRESSION"
			compressible_content_ref = data.avi_stringgroup.system_compressiblestringgroup.id
			compression = false
			remove_accept_encoding_header = true
			mobile_str_ref = data.avi_stringgroup.system_mobilestringgroup.id
		}
		xff_enabled = true
		disable_keepalive_posts_msie6 = true
		keepalive_timeout = "30000"
		ssl_client_certificate_mode = "SSL_CLIENT_CERTIFICATE_NONE"
		http_to_https = true
		respond_with_100_continue = true
		max_bad_rps_cip_uri = "0"
		client_body_timeout = "30000"
		httponly_enabled = true
		hsts_max_age = "365"
		max_bad_rps_cip = "0"
		server_side_redirect_to_https = true
		client_max_header_size = "12"
		client_max_request_size = "48"
		cache_config {
			min_object_size = "100"
			query_cacheable = false
			xcache_header = true
			age_header = true
			enabled = false
			default_expire = "600"
			max_cache_size = "0"
			heuristic_expire = false
			date_header = true
			aggressive = false
			max_object_size = "4194304"
			mime_types_group_refs = [data.avi_stringgroup.system_cacheablestringgroup.id]
		}
		http2_profile {
			http2_initial_window_size = "64"
			max_http2_concurrent_streams_per_connection = "128"
			max_http2_control_frames_per_connection = "1000"
			max_http2_empty_data_frames_per_connection = "1000"
			max_http2_header_field_size = "4096"
			max_http2_queued_frames_to_client_per_connection = "1000"
			max_http2_requests_per_connection = "1000"
		}
		max_rps_unknown_uri = "0"
		post_accept_timeout = "30000"
		client_header_timeout = "10000"
		secure_cookie_enabled = true
		max_response_headers_size = "48"
		xff_alternate_name = "X-Forwarded-For"
		max_rps_cip = "0"
		enable_fire_and_forget = false
		max_rps_unknown_cip = "0"
		allow_dots_in_header_name = false
		max_bad_rps_uri = "0"
		use_app_keepalive_timeout = false
	}
	preserve_client_port = false
}
`

const testAccAVIApplicationProfileupdatedConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
data "avi_stringgroup" "system_compressiblestringgroup" {
    name= "System-Compressible-Content-Types"
}
data "avi_stringgroup" "system_cacheablestringgroup" {
    name= "System-Cacheable-Resource-Types"
}
data "avi_stringgroup" "system_mobilestringgroup" {
    name = "System-Devices-Mobile"
}
resource "avi_applicationprofile" "testApplicationProfile" {
	name = "test-System-Secure-HTTP-updated"
	type = "APPLICATION_PROFILE_TYPE_HTTP"
	tenant_ref = data.avi_tenant.default_tenant.id
	preserve_client_ip = false
	http_profile {
		max_rps_uri = "0"
		keepalive_header = false
		max_rps_cip_uri = "0"
		x_forwarded_proto_enabled = true
		connection_multiplexing_enabled = true
		websockets_enabled = true
		enable_request_body_buffering = false
		hsts_enabled = true
		compression_profile {
			type = "AUTO_COMPRESSION"
			compressible_content_ref = data.avi_stringgroup.system_compressiblestringgroup.id
			compression = false
			remove_accept_encoding_header = true
			mobile_str_ref = data.avi_stringgroup.system_mobilestringgroup.id
		}
		xff_enabled = true
		disable_keepalive_posts_msie6 = true
		keepalive_timeout = "30000"
		ssl_client_certificate_mode = "SSL_CLIENT_CERTIFICATE_NONE"
		http_to_https = true
		respond_with_100_continue = true
		max_bad_rps_cip_uri = "0"
		client_body_timeout = "30000"
		httponly_enabled = true
		hsts_max_age = "365"
		max_bad_rps_cip = "0"
		server_side_redirect_to_https = true
		client_max_header_size = "12"
		client_max_request_size = "48"
		cache_config {
			min_object_size = "100"
			query_cacheable = false
			xcache_header = true
			age_header = true
			enabled = false
			default_expire = "600"
			max_cache_size = "0"
			heuristic_expire = false
			date_header = true
			aggressive = false
			max_object_size = "4194304"
			mime_types_group_refs = [data.avi_stringgroup.system_cacheablestringgroup.id]
		}
		http2_profile {
			http2_initial_window_size = "64"
			max_http2_concurrent_streams_per_connection = "128"
			max_http2_control_frames_per_connection = "1000"
			max_http2_empty_data_frames_per_connection = "1000"
			max_http2_header_field_size = "4096"
			max_http2_queued_frames_to_client_per_connection = "1000"
			max_http2_requests_per_connection = "1000"
		}
		max_rps_unknown_uri = "0"
		post_accept_timeout = "30000"
		client_header_timeout = "10000"
		secure_cookie_enabled = true
		max_response_headers_size = "48"
		xff_alternate_name = "X-Forwarded-For"
		max_rps_cip = "0"
		enable_fire_and_forget = false
		max_rps_unknown_cip = "0"
		allow_dots_in_header_name = false
		max_bad_rps_uri = "0"
		use_app_keepalive_timeout = false
	}
	preserve_client_port = false
}
`
