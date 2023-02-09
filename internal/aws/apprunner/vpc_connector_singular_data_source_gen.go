// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package apprunner

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_apprunner_vpc_connector", vpcConnectorDataSource)
}

// vpcConnectorDataSource returns the Terraform awscc_apprunner_vpc_connector data source.
// This Terraform data source corresponds to the CloudFormation AWS::AppRunner::VpcConnector resource.
func vpcConnectorDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: SecurityGroups
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of IDs of security groups that App Runner should use for access to AWS resources under the specified subnets. If not specified, App Runner uses the default security group of the Amazon VPC. The default security group allows all outbound traffic.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"security_groups": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A list of IDs of security groups that App Runner should use for access to AWS resources under the specified subnets. If not specified, App Runner uses the default security group of the Amazon VPC. The default security group allows all outbound traffic.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Subnets
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of IDs of subnets that App Runner should use when it associates your service with a custom Amazon VPC. Specify IDs of subnets of a single Amazon VPC. App Runner determines the Amazon VPC from the subnets you specify.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "minItems": 1,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"subnets": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A list of IDs of subnets that App Runner should use when it associates your service with a custom Amazon VPC. Specify IDs of subnets of a single Amazon VPC. App Runner determines the Amazon VPC from the subnets you specify.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of metadata items that you can associate with your VPC connector resource. A tag is a key-value pair.",
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
			Description: "A list of metadata items that you can associate with your VPC connector resource. A tag is a key-value pair.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VpcConnectorArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of this VPC connector.",
		//	  "maxLength": 1011,
		//	  "minLength": 44,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"vpc_connector_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of this VPC connector.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VpcConnectorName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A name for the VPC connector. If you don't specify a name, AWS CloudFormation generates a name for your VPC connector.",
		//	  "maxLength": 40,
		//	  "minLength": 4,
		//	  "pattern": "^[A-Za-z0-9][A-Za-z0-9-\\\\_]{3,39}$",
		//	  "type": "string"
		//	}
		"vpc_connector_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A name for the VPC connector. If you don't specify a name, AWS CloudFormation generates a name for your VPC connector.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VpcConnectorRevision
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The revision of this VPC connector. It's unique among all the active connectors (\"Status\": \"ACTIVE\") that share the same Name.",
		//	  "type": "integer"
		//	}
		"vpc_connector_revision": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The revision of this VPC connector. It's unique among all the active connectors (\"Status\": \"ACTIVE\") that share the same Name.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::AppRunner::VpcConnector",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::AppRunner::VpcConnector").WithTerraformTypeName("awscc_apprunner_vpc_connector")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"key":                    "Key",
		"security_groups":        "SecurityGroups",
		"subnets":                "Subnets",
		"tags":                   "Tags",
		"value":                  "Value",
		"vpc_connector_arn":      "VpcConnectorArn",
		"vpc_connector_name":     "VpcConnectorName",
		"vpc_connector_revision": "VpcConnectorRevision",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
