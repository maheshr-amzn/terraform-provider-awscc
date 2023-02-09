// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_ec2_subnet_network_acl_association", subnetNetworkAclAssociationResource)
}

// subnetNetworkAclAssociationResource returns the Terraform awscc_ec2_subnet_network_acl_association resource.
// This Terraform resource corresponds to the CloudFormation AWS::EC2::SubnetNetworkAclAssociation resource.
func subnetNetworkAclAssociationResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AssociationId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"association_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: NetworkAclId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the network ACL",
		//	  "type": "string"
		//	}
		"network_acl_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the network ACL",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SubnetId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the subnet",
		//	  "type": "string"
		//	}
		"subnet_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the subnet",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Resource Type definition for AWS::EC2::SubnetNetworkAclAssociation",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::SubnetNetworkAclAssociation").WithTerraformTypeName("awscc_ec2_subnet_network_acl_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"association_id": "AssociationId",
		"network_acl_id": "NetworkAclId",
		"subnet_id":      "SubnetId",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
