// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_lightsail_bucket", bucketDataSource)
}

// bucketDataSource returns the Terraform awscc_lightsail_bucket data source.
// This Terraform data source corresponds to the CloudFormation AWS::Lightsail::Bucket resource.
func bucketDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AbleToUpdateBundle
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether the bundle that is currently applied to a bucket can be changed to another bundle. You can update a bucket's bundle only one time within a monthly AWS billing cycle.",
		//	  "type": "boolean"
		//	}
		"able_to_update_bundle": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether the bundle that is currently applied to a bucket can be changed to another bundle. You can update a bucket's bundle only one time within a monthly AWS billing cycle.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AccessRules
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "An object that sets the public accessibility of objects in the specified bucket.",
		//	  "properties": {
		//	    "AllowPublicOverrides": {
		//	      "description": "A Boolean value that indicates whether the access control list (ACL) permissions that are applied to individual objects override the getObject option that is currently specified.",
		//	      "type": "boolean"
		//	    },
		//	    "GetObject": {
		//	      "description": "Specifies the anonymous access to all objects in a bucket.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"access_rules": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AllowPublicOverrides
				"allow_public_overrides": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "A Boolean value that indicates whether the access control list (ACL) permissions that are applied to individual objects override the getObject option that is currently specified.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: GetObject
				"get_object": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Specifies the anonymous access to all objects in a bucket.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "An object that sets the public accessibility of objects in the specified bucket.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: BucketArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"bucket_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: BucketName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name for the bucket.",
		//	  "maxLength": 54,
		//	  "minLength": 3,
		//	  "pattern": "^[a-z0-9][a-z0-9-]{1,52}[a-z0-9]$",
		//	  "type": "string"
		//	}
		"bucket_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name for the bucket.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: BundleId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the bundle to use for the bucket.",
		//	  "type": "string"
		//	}
		"bundle_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the bundle to use for the bucket.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ObjectVersioning
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies whether to enable or disable versioning of objects in the bucket.",
		//	  "type": "boolean"
		//	}
		"object_versioning": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies whether to enable or disable versioning of objects in the bucket.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ReadOnlyAccessAccounts
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of strings to specify the AWS account IDs that can access the bucket.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"read_only_access_accounts": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "An array of strings to specify the AWS account IDs that can access the bucket.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResourcesReceivingAccess
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The names of the Lightsail resources for which to set bucket access.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"resources_receiving_access": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The names of the Lightsail resources for which to set bucket access.",
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
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key"
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
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Url
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The URL of the bucket.",
		//	  "type": "string"
		//	}
		"url": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The URL of the bucket.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Lightsail::Bucket",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Lightsail::Bucket").WithTerraformTypeName("awscc_lightsail_bucket")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"able_to_update_bundle":      "AbleToUpdateBundle",
		"access_rules":               "AccessRules",
		"allow_public_overrides":     "AllowPublicOverrides",
		"bucket_arn":                 "BucketArn",
		"bucket_name":                "BucketName",
		"bundle_id":                  "BundleId",
		"get_object":                 "GetObject",
		"key":                        "Key",
		"object_versioning":          "ObjectVersioning",
		"read_only_access_accounts":  "ReadOnlyAccessAccounts",
		"resources_receiving_access": "ResourcesReceivingAccess",
		"tags":                       "Tags",
		"url":                        "Url",
		"value":                      "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
