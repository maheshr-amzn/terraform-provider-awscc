---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_xray_resource_policy Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::XRay::ResourcePolicy
---

# awscc_xray_resource_policy (Data Source)

Data Source schema for AWS::XRay::ResourcePolicy



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `bypass_policy_lockout_check` (Boolean) A flag to indicate whether to bypass the resource policy lockout safety check
- `policy_document` (String) The resource policy document, which can be up to 5kb in size.
- `policy_name` (String) The name of the resource policy. Must be unique within a specific AWS account.


