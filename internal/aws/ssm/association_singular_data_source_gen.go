// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ssm_association", associationDataSource)
}

// associationDataSource returns the Terraform awscc_ssm_association data source.
// This Terraform data source corresponds to the CloudFormation AWS::SSM::Association resource.
func associationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ApplyOnlyAtCronInterval
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "boolean"
		//	}
		"apply_only_at_cron_interval": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: AssociationId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Unique identifier of the association.",
		//	  "examples": [
		//	    "88df7b09-95e8-48c4-a3cb-08c2c20d5110",
		//	    "203dd0ec-0055-4bf0-a872-707f72ef06aa"
		//	  ],
		//	  "pattern": "[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}",
		//	  "type": "string"
		//	}
		"association_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Unique identifier of the association.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AssociationName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the association.",
		//	  "pattern": "^[a-zA-Z0-9_\\-.]{3,128}$",
		//	  "type": "string"
		//	}
		"association_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the association.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AutomationTargetParameterName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 50,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"automation_target_parameter_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: CalendarNames
		// CloudFormation resource type schema:
		//
		//	{
		//	  "examples": [
		//	    [
		//	      "calendar1",
		//	      "calendar2"
		//	    ],
		//	    [
		//	      "calendar3"
		//	    ]
		//	  ],
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"calendar_names": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ComplianceSeverity
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "CRITICAL",
		//	    "HIGH",
		//	    "MEDIUM",
		//	    "LOW",
		//	    "UNSPECIFIED"
		//	  ],
		//	  "type": "string"
		//	}
		"compliance_severity": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: DocumentVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The version of the SSM document to associate with the target.",
		//	  "pattern": "([$]LATEST|[$]DEFAULT|^[1-9][0-9]*$)",
		//	  "type": "string"
		//	}
		"document_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The version of the SSM document to associate with the target.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: InstanceId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the instance that the SSM document is associated with.",
		//	  "examples": [
		//	    "i-0e60836d21cf313c4",
		//	    "mi-0532c22e49636ee13"
		//	  ],
		//	  "pattern": "(^i-(\\w{8}|\\w{17})$)|(^mi-\\w{17}$)",
		//	  "type": "string"
		//	}
		"instance_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the instance that the SSM document is associated with.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MaxConcurrency
		// CloudFormation resource type schema:
		//
		//	{
		//	  "examples": [
		//	    "1%",
		//	    "10%",
		//	    "50%",
		//	    "1"
		//	  ],
		//	  "pattern": "^([1-9][0-9]{0,6}|[1-9][0-9]%|[1-9]%|100%)$",
		//	  "type": "string"
		//	}
		"max_concurrency": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: MaxErrors
		// CloudFormation resource type schema:
		//
		//	{
		//	  "examples": [
		//	    "1%",
		//	    "10%",
		//	    "50%",
		//	    "1"
		//	  ],
		//	  "pattern": "^([1-9][0-9]{0,6}|[0]|[1-9][0-9]%|[0-9]%|100%)$",
		//	  "type": "string"
		//	}
		"max_errors": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the SSM document.",
		//	  "examples": [
		//	    "AWS-GatherSoftwareInventory",
		//	    "MyCustomSSMDocument"
		//	  ],
		//	  "pattern": "^[a-zA-Z0-9_\\-.:/]{3,200}$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the SSM document.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: OutputLocation
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "S3Location": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "OutputS3BucketName": {
		//	          "maxLength": 63,
		//	          "minLength": 3,
		//	          "type": "string"
		//	        },
		//	        "OutputS3KeyPrefix": {
		//	          "maxLength": 1024,
		//	          "type": "string"
		//	        },
		//	        "OutputS3Region": {
		//	          "maxLength": 20,
		//	          "minLength": 3,
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"output_location": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: S3Location
				"s3_location": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: OutputS3BucketName
						"output_s3_bucket_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: OutputS3KeyPrefix
						"output_s3_key_prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: OutputS3Region
						"output_s3_region": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Parameters
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Parameter values that the SSM document uses at runtime.",
		//	  "patternProperties": {
		//	    "": {
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"parameters":        // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.ListType{ElemType: types.StringType},
			Description: "Parameter values that the SSM document uses at runtime.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ScheduleExpression
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A Cron or Rate expression that specifies when the association is applied to the target.",
		//	  "examples": [
		//	    "cron(0 0 */1 * * ? *)",
		//	    "cron(0 16 ? * TUE *)",
		//	    "rate(30 minutes)",
		//	    "rate(7 days)"
		//	  ],
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"schedule_expression": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A Cron or Rate expression that specifies when the association is applied to the target.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ScheduleOffset
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maximum": 6,
		//	  "minimum": 1,
		//	  "type": "integer"
		//	}
		"schedule_offset": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: SyncCompliance
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "AUTO",
		//	    "MANUAL"
		//	  ],
		//	  "type": "string"
		//	}
		"sync_compliance": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Targets
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The targets that the SSM document sends commands to.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "pattern": "^[\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]{1,128}$|resource-groups:Name",
		//	        "type": "string"
		//	      },
		//	      "Values": {
		//	        "items": {
		//	          "type": "string"
		//	        },
		//	        "maxItems": 50,
		//	        "minItems": 0,
		//	        "type": "array"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Values"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 5,
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"targets": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Values
					"values": schema.ListAttribute{ /*START ATTRIBUTE*/
						ElementType: types.StringType,
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The targets that the SSM document sends commands to.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: WaitForSuccessTimeoutSeconds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maximum": 172800,
		//	  "minimum": 15,
		//	  "type": "integer"
		//	}
		"wait_for_success_timeout_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::SSM::Association",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SSM::Association").WithTerraformTypeName("awscc_ssm_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"apply_only_at_cron_interval":      "ApplyOnlyAtCronInterval",
		"association_id":                   "AssociationId",
		"association_name":                 "AssociationName",
		"automation_target_parameter_name": "AutomationTargetParameterName",
		"calendar_names":                   "CalendarNames",
		"compliance_severity":              "ComplianceSeverity",
		"document_version":                 "DocumentVersion",
		"instance_id":                      "InstanceId",
		"key":                              "Key",
		"max_concurrency":                  "MaxConcurrency",
		"max_errors":                       "MaxErrors",
		"name":                             "Name",
		"output_location":                  "OutputLocation",
		"output_s3_bucket_name":            "OutputS3BucketName",
		"output_s3_key_prefix":             "OutputS3KeyPrefix",
		"output_s3_region":                 "OutputS3Region",
		"parameters":                       "Parameters",
		"s3_location":                      "S3Location",
		"schedule_expression":              "ScheduleExpression",
		"schedule_offset":                  "ScheduleOffset",
		"sync_compliance":                  "SyncCompliance",
		"targets":                          "Targets",
		"values":                           "Values",
		"wait_for_success_timeout_seconds": "WaitForSuccessTimeoutSeconds",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
