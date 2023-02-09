// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package rds

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_rds_db_proxy_endpoint", dBProxyEndpointDataSource)
}

// dBProxyEndpointDataSource returns the Terraform awscc_rds_db_proxy_endpoint data source.
// This Terraform data source corresponds to the CloudFormation AWS::RDS::DBProxyEndpoint resource.
func dBProxyEndpointDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: DBProxyEndpointArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) for the DB proxy endpoint.",
		//	  "pattern": "arn:aws[A-Za-z0-9-]{0,64}:rds:[A-Za-z0-9-]{1,64}:[0-9]{12}:.*",
		//	  "type": "string"
		//	}
		"db_proxy_endpoint_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) for the DB proxy endpoint.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DBProxyEndpointName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier for the DB proxy endpoint. This name must be unique for all DB proxy endpoints owned by your AWS account in the specified AWS Region.",
		//	  "maxLength": 64,
		//	  "pattern": "[0-z]*",
		//	  "type": "string"
		//	}
		"db_proxy_endpoint_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier for the DB proxy endpoint. This name must be unique for all DB proxy endpoints owned by your AWS account in the specified AWS Region.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DBProxyName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier for the proxy. This name must be unique for all proxies owned by your AWS account in the specified AWS Region.",
		//	  "maxLength": 64,
		//	  "pattern": "[0-z]*",
		//	  "type": "string"
		//	}
		"db_proxy_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier for the proxy. This name must be unique for all proxies owned by your AWS account in the specified AWS Region.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Endpoint
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The endpoint that you can use to connect to the DB proxy. You include the endpoint value in the connection string for a database client application.",
		//	  "maxLength": 256,
		//	  "type": "string"
		//	}
		"endpoint": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The endpoint that you can use to connect to the DB proxy. You include the endpoint value in the connection string for a database client application.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IsDefault
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A value that indicates whether this endpoint is the default endpoint for the associated DB proxy. Default DB proxy endpoints always have read/write capability. Other endpoints that you associate with the DB proxy can be either read/write or read-only.",
		//	  "type": "boolean"
		//	}
		"is_default": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "A value that indicates whether this endpoint is the default endpoint for the associated DB proxy. Default DB proxy endpoints always have read/write capability. Other endpoints that you associate with the DB proxy can be either read/write or read-only.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An optional set of key-value pairs to associate arbitrary data of your choosing with the DB proxy endpoint.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "pattern": "(\\w|\\d|\\s|\\\\|-|\\.:=+-)*",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 128,
		//	        "pattern": "(\\w|\\d|\\s|\\\\|-|\\.:=+-)*",
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
			Description: "An optional set of key-value pairs to associate arbitrary data of your choosing with the DB proxy endpoint.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TargetRole
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A value that indicates whether the DB proxy endpoint can be used for read/write or read-only operations.",
		//	  "enum": [
		//	    "READ_WRITE",
		//	    "READ_ONLY"
		//	  ],
		//	  "type": "string"
		//	}
		"target_role": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A value that indicates whether the DB proxy endpoint can be used for read/write or read-only operations.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VpcId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "VPC ID to associate with the new DB proxy endpoint.",
		//	  "type": "string"
		//	}
		"vpc_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "VPC ID to associate with the new DB proxy endpoint.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VpcSecurityGroupIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "VPC security group IDs to associate with the new DB proxy endpoint.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"vpc_security_group_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "VPC security group IDs to associate with the new DB proxy endpoint.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VpcSubnetIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "VPC subnet IDs to associate with the new DB proxy endpoint.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "minItems": 2,
		//	  "type": "array"
		//	}
		"vpc_subnet_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "VPC subnet IDs to associate with the new DB proxy endpoint.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::RDS::DBProxyEndpoint",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::RDS::DBProxyEndpoint").WithTerraformTypeName("awscc_rds_db_proxy_endpoint")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"db_proxy_endpoint_arn":  "DBProxyEndpointArn",
		"db_proxy_endpoint_name": "DBProxyEndpointName",
		"db_proxy_name":          "DBProxyName",
		"endpoint":               "Endpoint",
		"is_default":             "IsDefault",
		"key":                    "Key",
		"tags":                   "Tags",
		"target_role":            "TargetRole",
		"value":                  "Value",
		"vpc_id":                 "VpcId",
		"vpc_security_group_ids": "VpcSecurityGroupIds",
		"vpc_subnet_ids":         "VpcSubnetIds",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
