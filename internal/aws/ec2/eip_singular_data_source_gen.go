// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ec2_eip", eIPDataSource)
}

// eIPDataSource returns the Terraform awscc_ec2_eip data source.
// This Terraform data source corresponds to the CloudFormation AWS::EC2::EIP resource.
func eIPDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AllocationId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Allocation ID of the EIP generated by resource.",
		//	  "type": "string"
		//	}
		"allocation_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Allocation ID of the EIP generated by resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Domain
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether the Elastic IP address is for use with instances in a VPC or instance in EC2-Classic.",
		//	  "type": "string"
		//	}
		"domain": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether the Elastic IP address is for use with instances in a VPC or instance in EC2-Classic.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: InstanceId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the instance.",
		//	  "type": "string"
		//	}
		"instance_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the instance.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: NetworkBorderGroup
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A unique set of Availability Zones, Local Zones, or Wavelength Zones from which Amazon Web Services advertises IP addresses.",
		//	  "type": "string"
		//	}
		"network_border_group": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A unique set of Availability Zones, Local Zones, or Wavelength Zones from which Amazon Web Services advertises IP addresses.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PublicIp
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The PublicIP of the EIP generated by resource.",
		//	  "type": "string"
		//	}
		"public_ip": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The PublicIP of the EIP generated by resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PublicIpv4Pool
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of an address pool that you own. Use this parameter to let Amazon EC2 select an address from the address pool.",
		//	  "type": "string"
		//	}
		"public_ipv_4_pool": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of an address pool that you own. Use this parameter to let Amazon EC2 select an address from the address pool.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Any tags assigned to the EIP.",
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
		//	  "type": "array",
		//	  "uniqueItems": false
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
			Description: "Any tags assigned to the EIP.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TransferAddress
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The PublicIP of the EIP generated by resource through transfer from another account",
		//	  "type": "string"
		//	}
		"transfer_address": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The PublicIP of the EIP generated by resource through transfer from another account",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::EC2::EIP",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::EIP").WithTerraformTypeName("awscc_ec2_eip")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"allocation_id":        "AllocationId",
		"domain":               "Domain",
		"instance_id":          "InstanceId",
		"key":                  "Key",
		"network_border_group": "NetworkBorderGroup",
		"public_ip":            "PublicIp",
		"public_ipv_4_pool":    "PublicIpv4Pool",
		"tags":                 "Tags",
		"transfer_address":     "TransferAddress",
		"value":                "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
