// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package route53recoveryreadiness

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_route53recoveryreadiness_readiness_check", readinessCheckDataSource)
}

// readinessCheckDataSource returns the Terraform awscc_route53recoveryreadiness_readiness_check data source.
// This Terraform data source corresponds to the CloudFormation AWS::Route53RecoveryReadiness::ReadinessCheck resource.
func readinessCheckDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ReadinessCheckArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the readiness check.",
		//	  "maxLength": 256,
		//	  "type": "string"
		//	}
		"readiness_check_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the readiness check.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ReadinessCheckName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name of the ReadinessCheck to create.",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "[a-zA-Z0-9_]+",
		//	  "type": "string"
		//	}
		"readiness_check_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name of the ReadinessCheck to create.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResourceSetName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the resource set to check.",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "[a-zA-Z0-9_]+",
		//	  "type": "string"
		//	}
		"resource_set_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the resource set to check.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A collection of tags associated with a resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A collection of tags associated with a resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Route53RecoveryReadiness::ReadinessCheck",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53RecoveryReadiness::ReadinessCheck").WithTerraformTypeName("awscc_route53recoveryreadiness_readiness_check")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"key":                  "Key",
		"readiness_check_arn":  "ReadinessCheckArn",
		"readiness_check_name": "ReadinessCheckName",
		"resource_set_name":    "ResourceSetName",
		"tags":                 "Tags",
		"value":                "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
