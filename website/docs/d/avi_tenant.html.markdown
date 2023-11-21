<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_tenant"
sidebar_current: "docs-avi-datasource-tenant"
description: |-
  Get information of Avi Tenant.
---

# avi_tenant

This data source is used to to get avi_tenant objects.

## Example Usage

```hcl
data "avi_tenant" "foo_tenant" {
    uuid = "tenant-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search Tenant by name.
* `uuid` - (Optional) Search Tenant by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `attrs` - Key/value tenant attributes. Field introduced in 30.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `config_settings` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `created_by` - Creator of this tenant. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `description` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `enforce_label_group` - The referred label groups are enforced on the tenant if this is set to true.if this is set to false, the label groups are suggested for the tenant. Field introduced in 20.1.5. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `label_group_refs` - The label_groups to be enforced on the tenant. This is strictly enforced only if enforce_label_group is set to true. It is a reference to an object of type labelgroup. Field introduced in 20.1.5. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `local` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `name` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `uuid` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

