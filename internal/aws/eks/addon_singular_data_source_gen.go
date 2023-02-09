// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package eks

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_eks_addon", addonDataSource)
}

// addonDataSource returns the Terraform awscc_eks_addon data source.
// This Terraform data source corresponds to the CloudFormation AWS::EKS::Addon resource.
func addonDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AddonName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name of Addon",
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"addon_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name of Addon",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AddonVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Version of Addon",
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"addon_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Version of Addon",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Amazon Resource Name (ARN) of the add-on",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Amazon Resource Name (ARN) of the add-on",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ClusterName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name of Cluster",
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"cluster_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name of Cluster",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ConfigurationValues
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The configuration values to use with the add-on",
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"configuration_values": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The configuration values to use with the add-on",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PreserveOnDelete
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "PreserveOnDelete parameter value",
		//	  "type": "boolean"
		//	}
		"preserve_on_delete": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "PreserveOnDelete parameter value",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResolveConflicts
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Resolve parameter value conflicts",
		//	  "enum": [
		//	    "NONE",
		//	    "OVERWRITE",
		//	    "PRESERVE"
		//	  ],
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"resolve_conflicts": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Resolve parameter value conflicts",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ServiceAccountRoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "IAM role to bind to the add-on's service account",
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"service_account_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "IAM role to bind to the add-on's service account",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 127,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 255,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::EKS::Addon",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EKS::Addon").WithTerraformTypeName("awscc_eks_addon")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"addon_name":               "AddonName",
		"addon_version":            "AddonVersion",
		"arn":                      "Arn",
		"cluster_name":             "ClusterName",
		"configuration_values":     "ConfigurationValues",
		"key":                      "Key",
		"preserve_on_delete":       "PreserveOnDelete",
		"resolve_conflicts":        "ResolveConflicts",
		"service_account_role_arn": "ServiceAccountRoleArn",
		"tags":                     "Tags",
		"value":                    "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
