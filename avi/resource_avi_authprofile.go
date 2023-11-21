// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceAuthProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"http": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceAuthProfileHTTPClientParamsSchema(),
		},
		"jwt_profile_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"ldap": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceLdapAuthSettingsSchema(),
		},
		"markers": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceRoleFilterMatchLabelSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"oauth_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceOAuthProfileSchema(),
		},
		"pa_agent_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"saml": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceSamlSettingsSchema(),
		},
		"tacacs_plus": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceTacacsPlusAuthSettingsSchema(),
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"type": {
			Type:     schema.TypeString,
			Required: true,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviAuthProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviAuthProfileCreate,
		Read:   ResourceAviAuthProfileRead,
		Update: resourceAviAuthProfileUpdate,
		Delete: resourceAviAuthProfileDelete,
		Schema: ResourceAuthProfileSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceAuthProfileImporter,
		},
	}
}

func ResourceAuthProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceAuthProfileSchema()
	return ResourceImporter(d, m, "authprofile", s)
}

func ResourceAviAuthProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAuthProfileSchema()
	err := APIRead(d, meta, "authprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviAuthProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAuthProfileSchema()
	err := APICreateOrUpdate(d, meta, "authprofile", s)
	if err == nil {
		err = ResourceAviAuthProfileRead(d, meta)
	}
	return err
}

func resourceAviAuthProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAuthProfileSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "authprofile", s)
	if err == nil {
		err = ResourceAviAuthProfileRead(d, meta)
	}
	return err
}

func resourceAviAuthProfileDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "authprofile")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
