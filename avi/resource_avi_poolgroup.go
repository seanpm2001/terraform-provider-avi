// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourcePoolGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cloud_config_cksum": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"cloud_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"created_by": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"deactivate_primary_pool_on_down": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"deployment_policy_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"fail_action": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceFailActionSchema(),
		},
		"implicit_priority_labels": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"markers": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceRoleFilterMatchLabelSchema(),
		},
		"members": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourcePoolGroupMemberSchema(),
		},
		"min_servers": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateInteger,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"priority_labels_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"service_metadata": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviPoolGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviPoolGroupCreate,
		Read:   ResourceAviPoolGroupRead,
		Update: resourceAviPoolGroupUpdate,
		Delete: resourceAviPoolGroupDelete,
		Schema: ResourcePoolGroupSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourcePoolGroupImporter,
		},
	}
}

func ResourcePoolGroupImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourcePoolGroupSchema()
	return ResourceImporter(d, m, "poolgroup", s)
}

func ResourceAviPoolGroupRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePoolGroupSchema()
	err := APIRead(d, meta, "poolgroup", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviPoolGroupCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePoolGroupSchema()
	err := APICreateOrUpdate(d, meta, "poolgroup", s)
	if err == nil {
		err = ResourceAviPoolGroupRead(d, meta)
	}
	return err
}

func resourceAviPoolGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePoolGroupSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "poolgroup", s)
	if err == nil {
		err = ResourceAviPoolGroupRead(d, meta)
	}
	return err
}

func resourceAviPoolGroupDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "poolgroup")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
