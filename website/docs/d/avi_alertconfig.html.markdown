<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_alertconfig"
sidebar_current: "docs-avi-datasource-alertconfig"
description: |-
  Get information of Avi AlertConfig.
---

# avi_alertconfig

This data source is used to to get avi_alertconfig objects.

## Example Usage

```hcl
data "avi_alertconfig" "foo_alertconfig" {
    uuid = "alertconfig-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search AlertConfig by name.
* `uuid` - (Optional) Search AlertConfig by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `action_group_ref` - The alert config will trigger the selected alert action, which can send notifications and execute a controlscript. It is a reference to an object of type actiongroupconfig. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `alert_rule` - List of filters matching on events or client logs used for triggering alerts. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `autoscale_alert` - This alert config applies to auto scale alerts. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `category` - Determines whether an alert is raised immediately when event occurs (realtime) or after specified number of events occurs within rolling time window. Enum options - REALTIME, ROLLINGWINDOW, WATERMARK. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `description` - A custom description field. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `enabled` - Enable or disable this alert config from generating new alerts. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `expiry_time` - An alert is expired and deleted after the expiry time has elapsed. The original event triggering the alert remains in the event's log. Allowed values are 1-31536000. Unit is sec. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `name` - Name of the alert configuration. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `obj_uuid` - Uuid of the resource for which alert was raised. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `object_type` - The object type to which the alert config is associated with. Valid object types are - virtual service, pool, service engine. Enum options - VIRTUALSERVICE, POOL, HEALTHMONITOR, NETWORKPROFILE, APPLICATIONPROFILE, HTTPPOLICYSET, DNSPOLICY, SECURITYPOLICY, IPADDRGROUP, STRINGGROUP, SSLPROFILE, SSLKEYANDCERTIFICATE, NETWORKSECURITYPOLICY, APPLICATIONPERSISTENCEPROFILE, ANALYTICSPROFILE, VSDATASCRIPTSET, TENANT, PKIPROFILE, AUTHPROFILE, CLOUD... Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `recommendation` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `rolling_window` - Only if the number of events is reached or exceeded within the time window will an alert be generated. Allowed values are 1-31536000. Unit is sec. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `source` - Signifies system events or the type of client logsused in this alert configuration. Enum options - CONN_LOGS, APP_LOGS, EVENT_LOGS, METRICS. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `summary` - Summary of reason why alert is generated. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `threshold` - An alert is created only when the number of events meets or exceeds this number within the chosen time frame. Allowed values are 1-65536. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `throttle` - Alerts are suppressed (throttled) for this duration of time since the last alert was raised for this alert config. Allowed values are 0-31536000. Unit is sec. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `uuid` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

