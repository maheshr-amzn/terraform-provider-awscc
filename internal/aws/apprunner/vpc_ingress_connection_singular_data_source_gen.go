// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package apprunner

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_apprunner_vpc_ingress_connection", vpcIngressConnectionDataSource)
}

// vpcIngressConnectionDataSource returns the Terraform awscc_apprunner_vpc_ingress_connection data source.
// This Terraform data source corresponds to the CloudFormation AWS::AppRunner::VpcIngressConnection resource.
func vpcIngressConnectionDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: DomainName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Domain name associated with the VPC Ingress Connection.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": "[A-Za-z0-9*.-]{1,255}",
		//	  "type": "string"
		//	}
		"domain_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Domain name associated with the VPC Ingress Connection.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IngressVpcConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The configuration of customer?s VPC and related VPC endpoint",
		//	  "properties": {
		//	    "VpcEndpointId": {
		//	      "description": "The ID of the VPC endpoint that your App Runner service connects to.",
		//	      "type": "string"
		//	    },
		//	    "VpcId": {
		//	      "description": "The ID of the VPC that the VPC endpoint is used in.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "VpcId",
		//	    "VpcEndpointId"
		//	  ],
		//	  "type": "object"
		//	}
		"ingress_vpc_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: VpcEndpointId
				"vpc_endpoint_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ID of the VPC endpoint that your App Runner service connects to.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: VpcId
				"vpc_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ID of the VPC that the VPC endpoint is used in.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The configuration of customer?s VPC and related VPC endpoint",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ServiceArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the service.",
		//	  "maxLength": 1011,
		//	  "minLength": 1,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"service_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the service.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The current status of the VpcIngressConnection.",
		//	  "enum": [
		//	    "AVAILABLE",
		//	    "PENDING_CREATION",
		//	    "PENDING_UPDATE",
		//	    "PENDING_DELETION",
		//	    "FAILED_CREATION",
		//	    "FAILED_UPDATE",
		//	    "FAILED_DELETION",
		//	    "DELETED"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The current status of the VpcIngressConnection.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
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
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: VpcIngressConnectionArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the VpcIngressConnection.",
		//	  "maxLength": 1011,
		//	  "minLength": 1,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"vpc_ingress_connection_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the VpcIngressConnection.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VpcIngressConnectionName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The customer-provided Vpc Ingress Connection name.",
		//	  "maxLength": 40,
		//	  "minLength": 4,
		//	  "pattern": "[A-Za-z0-9][A-Za-z0-9\\-_]{3,39}",
		//	  "type": "string"
		//	}
		"vpc_ingress_connection_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The customer-provided Vpc Ingress Connection name.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::AppRunner::VpcIngressConnection",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::AppRunner::VpcIngressConnection").WithTerraformTypeName("awscc_apprunner_vpc_ingress_connection")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"domain_name":                 "DomainName",
		"ingress_vpc_configuration":   "IngressVpcConfiguration",
		"key":                         "Key",
		"service_arn":                 "ServiceArn",
		"status":                      "Status",
		"tags":                        "Tags",
		"value":                       "Value",
		"vpc_endpoint_id":             "VpcEndpointId",
		"vpc_id":                      "VpcId",
		"vpc_ingress_connection_arn":  "VpcIngressConnectionArn",
		"vpc_ingress_connection_name": "VpcIngressConnectionName",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
