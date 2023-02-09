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
	registry.AddDataSourceFactory("awscc_ec2_host", hostDataSource)
}

// hostDataSource returns the Terraform awscc_ec2_host data source.
// This Terraform data source corresponds to the CloudFormation AWS::EC2::Host resource.
func hostDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AutoPlacement
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID.",
		//	  "type": "string"
		//	}
		"auto_placement": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AvailabilityZone
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Availability Zone in which to allocate the Dedicated Host.",
		//	  "type": "string"
		//	}
		"availability_zone": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Availability Zone in which to allocate the Dedicated Host.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: HostId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Id of the host created.",
		//	  "type": "string"
		//	}
		"host_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Id of the host created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: HostRecovery
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether to enable or disable host recovery for the Dedicated Host. Host recovery is disabled by default.",
		//	  "type": "string"
		//	}
		"host_recovery": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether to enable or disable host recovery for the Dedicated Host. Host recovery is disabled by default.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: InstanceFamily
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the instance family to be supported by the Dedicated Hosts. If you specify an instance family, the Dedicated Hosts support multiple instance types within that instance family.",
		//	  "type": "string"
		//	}
		"instance_family": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the instance family to be supported by the Dedicated Hosts. If you specify an instance family, the Dedicated Hosts support multiple instance types within that instance family.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: InstanceType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the instance type to be supported by the Dedicated Hosts. If you specify an instance type, the Dedicated Hosts support instances of the specified instance type only.",
		//	  "type": "string"
		//	}
		"instance_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the instance type to be supported by the Dedicated Hosts. If you specify an instance type, the Dedicated Hosts support instances of the specified instance type only.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: OutpostArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the Amazon Web Services Outpost on which to allocate the Dedicated Host.",
		//	  "type": "string"
		//	}
		"outpost_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the Amazon Web Services Outpost on which to allocate the Dedicated Host.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::EC2::Host",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::Host").WithTerraformTypeName("awscc_ec2_host")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"auto_placement":    "AutoPlacement",
		"availability_zone": "AvailabilityZone",
		"host_id":           "HostId",
		"host_recovery":     "HostRecovery",
		"instance_family":   "InstanceFamily",
		"instance_type":     "InstanceType",
		"outpost_arn":       "OutpostArn",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
