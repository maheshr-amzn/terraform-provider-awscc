// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_apigateway_usage_plan_key", usagePlanKeyDataSource)
}

// usagePlanKeyDataSource returns the Terraform awscc_apigateway_usage_plan_key data source.
// This Terraform data source corresponds to the CloudFormation AWS::ApiGateway::UsagePlanKey resource.
func usagePlanKeyDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An autogenerated ID which is a combination of the ID of the key and ID of the usage plan combined with a : such as 123abcdef:abc123.",
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An autogenerated ID which is a combination of the ID of the key and ID of the usage plan combined with a : such as 123abcdef:abc123.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: KeyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the usage plan key.",
		//	  "type": "string"
		//	}
		"key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the usage plan key.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: KeyType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of usage plan key. Currently, the only valid key type is API_KEY.",
		//	  "enum": [
		//	    "API_KEY"
		//	  ],
		//	  "type": "string"
		//	}
		"key_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of usage plan key. Currently, the only valid key type is API_KEY.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: UsagePlanId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the usage plan.",
		//	  "type": "string"
		//	}
		"usage_plan_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the usage plan.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::ApiGateway::UsagePlanKey",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApiGateway::UsagePlanKey").WithTerraformTypeName("awscc_apigateway_usage_plan_key")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"id":            "Id",
		"key_id":        "KeyId",
		"key_type":      "KeyType",
		"usage_plan_id": "UsagePlanId",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
