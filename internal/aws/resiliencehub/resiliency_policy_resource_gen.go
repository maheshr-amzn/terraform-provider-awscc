// Code generated by generators/resource/main.go; DO NOT EDIT.

package resiliencehub

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_resiliencehub_resiliency_policy", resiliencyPolicyResourceType)
}

// resiliencyPolicyResourceType returns the Terraform awscc_resiliencehub_resiliency_policy resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::ResilienceHub::ResiliencyPolicy resource type.
func resiliencyPolicyResourceType(ctx context.Context) (provider.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"data_location_constraint": {
			// Property: DataLocationConstraint
			// CloudFormation resource type schema:
			// {
			//   "description": "Data Location Constraint of the Policy.",
			//   "enum": [
			//     "AnyLocation",
			//     "SameContinent",
			//     "SameCountry"
			//   ],
			//   "type": "string"
			// }
			Description: "Data Location Constraint of the Policy.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"AnyLocation",
					"SameContinent",
					"SameCountry",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"policy": {
			// Property: Policy
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "patternProperties": {
			//     "": {
			//       "additionalProperties": false,
			//       "description": "Failure Policy.",
			//       "properties": {
			//         "RpoInSecs": {
			//           "description": "RPO in seconds.",
			//           "type": "integer"
			//         },
			//         "RtoInSecs": {
			//           "description": "RTO in seconds.",
			//           "type": "integer"
			//         }
			//       },
			//       "required": [
			//         "RtoInSecs",
			//         "RpoInSecs"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			// Pattern: ""
			Attributes: tfsdk.MapNestedAttributes(
				map[string]tfsdk.Attribute{
					"rpo_in_secs": {
						// Property: RpoInSecs
						Description: "RPO in seconds.",
						Type:        types.Int64Type,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"rto_in_secs": {
						// Property: RtoInSecs
						Description: "RTO in seconds.",
						Type:        types.Int64Type,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Required: true,
		},
		"policy_arn": {
			// Property: PolicyArn
			// CloudFormation resource type schema:
			// {
			//   "description": "Amazon Resource Name (ARN) of the Resiliency Policy.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Amazon Resource Name (ARN) of the Resiliency Policy.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"policy_description": {
			// Property: PolicyDescription
			// CloudFormation resource type schema:
			// {
			//   "description": "Description of Resiliency Policy.",
			//   "maxLength": 500,
			//   "type": "string"
			// }
			Description: "Description of Resiliency Policy.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenAtMost(500),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"policy_name": {
			// Property: PolicyName
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of Resiliency Policy.",
			//   "pattern": "^[A-Za-z0-9][A-Za-z0-9_\\-]{1,59}$",
			//   "type": "string"
			// }
			Description: "Name of Resiliency Policy.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringMatch(regexp.MustCompile("^[A-Za-z0-9][A-Za-z0-9_\\-]{1,59}$"), ""),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "patternProperties": {
			//     "": {
			//       "maxLength": 256,
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"tier": {
			// Property: Tier
			// CloudFormation resource type schema:
			// {
			//   "description": "Resiliency Policy Tier.",
			//   "enum": [
			//     "MissionCritical",
			//     "Critical",
			//     "Important",
			//     "CoreServices",
			//     "NonCritical"
			//   ],
			//   "type": "string"
			// }
			Description: "Resiliency Policy Tier.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"MissionCritical",
					"Critical",
					"Important",
					"CoreServices",
					"NonCritical",
				}),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			resource.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Type Definition for Resiliency Policy.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ResilienceHub::ResiliencyPolicy").WithTerraformTypeName("awscc_resiliencehub_resiliency_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"data_location_constraint": "DataLocationConstraint",
		"policy":                   "Policy",
		"policy_arn":               "PolicyArn",
		"policy_description":       "PolicyDescription",
		"policy_name":              "PolicyName",
		"rpo_in_secs":              "RpoInSecs",
		"rto_in_secs":              "RtoInSecs",
		"tags":                     "Tags",
		"tier":                     "Tier",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
