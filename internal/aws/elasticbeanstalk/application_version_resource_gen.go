// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package elasticbeanstalk

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_elasticbeanstalk_application_version", applicationVersionResource)
}

// applicationVersionResource returns the Terraform awscc_elasticbeanstalk_application_version resource.
// This Terraform resource corresponds to the CloudFormation AWS::ElasticBeanstalk::ApplicationVersion resource.
func applicationVersionResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ApplicationName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the Elastic Beanstalk application that is associated with this application version. ",
		//	  "type": "string"
		//	}
		"application_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the Elastic Beanstalk application that is associated with this application version. ",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description of this application version.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description of this application version.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SourceBundle
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The Amazon S3 bucket and key that identify the location of the source bundle for this version. ",
		//	  "properties": {
		//	    "S3Bucket": {
		//	      "description": "The Amazon S3 bucket where the data is located.",
		//	      "type": "string"
		//	    },
		//	    "S3Key": {
		//	      "description": "The Amazon S3 key where the data is located.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "S3Bucket",
		//	    "S3Key"
		//	  ],
		//	  "type": "object"
		//	}
		"source_bundle": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: S3Bucket
				"s3_bucket": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The Amazon S3 bucket where the data is located.",
					Required:    true,
				}, /*END ATTRIBUTE*/
				// Property: S3Key
				"s3_key": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The Amazon S3 key where the data is located.",
					Required:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The Amazon S3 bucket and key that identify the location of the source bundle for this version. ",
			Required:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	schema := schema.Schema{
		Description: "Resource Type definition for AWS::ElasticBeanstalk::ApplicationVersion",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ElasticBeanstalk::ApplicationVersion").WithTerraformTypeName("awscc_elasticbeanstalk_application_version")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"application_name": "ApplicationName",
		"description":      "Description",
		"id":               "Id",
		"s3_bucket":        "S3Bucket",
		"s3_key":           "S3Key",
		"source_bundle":    "SourceBundle",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
