// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package directoryservice

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_directoryservice_simple_ad", simpleADResource)
}

// simpleADResource returns the Terraform awscc_directoryservice_simple_ad resource.
// This Terraform resource corresponds to the CloudFormation AWS::DirectoryService::SimpleAD resource.
func simpleADResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Alias
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The alias for a directory.",
		//	  "type": "string"
		//	}
		"alias": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The alias for a directory.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CreateAlias
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the configuration set.",
		//	  "type": "boolean"
		//	}
		"create_alias": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the configuration set.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
				boolplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Description for the directory.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Description for the directory.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DirectoryId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique identifier for a directory.",
		//	  "type": "string"
		//	}
		"directory_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The unique identifier for a directory.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DnsIpAddresses
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The IP addresses of the DNS servers for the directory, such as [ \"172.31.3.154\", \"172.31.63.203\" ].",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"dns_ip_addresses": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The IP addresses of the DNS servers for the directory, such as [ \"172.31.3.154\", \"172.31.63.203\" ].",
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EnableSso
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Whether to enable single sign-on for a Simple Active Directory in AWS.",
		//	  "type": "boolean"
		//	}
		"enable_sso": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Whether to enable single sign-on for a Simple Active Directory in AWS.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The fully qualified domain name for the AWS Managed Simple AD directory.",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The fully qualified domain name for the AWS Managed Simple AD directory.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Password
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The password for the default administrative user named Admin.",
		//	  "type": "string"
		//	}
		"password": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The password for the default administrative user named Admin.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// Password is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: ShortName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The NetBIOS name for your domain.",
		//	  "type": "string"
		//	}
		"short_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The NetBIOS name for your domain.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Size
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The size of the directory.",
		//	  "type": "string"
		//	}
		"size": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The size of the directory.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VpcSettings
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "$comment": "Contains VPC information",
		//	  "description": "VPC settings of the Simple AD directory server in AWS.",
		//	  "properties": {
		//	    "SubnetIds": {
		//	      "description": "The identifiers of the subnets for the directory servers. The two subnets must be in different Availability Zones. AWS Directory Service specifies a directory server and a DNS server in each of these subnets.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "VpcId": {
		//	      "description": "The identifier of the VPC in which to create the directory.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "VpcId",
		//	    "SubnetIds"
		//	  ],
		//	  "type": "object"
		//	}
		"vpc_settings": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: SubnetIds
				"subnet_ids": schema.SetAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "The identifiers of the subnets for the directory servers. The two subnets must be in different Availability Zones. AWS Directory Service specifies a directory server and a DNS server in each of these subnets.",
					Required:    true,
				}, /*END ATTRIBUTE*/
				// Property: VpcId
				"vpc_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The identifier of the VPC in which to create the directory.",
					Required:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "VPC settings of the Simple AD directory server in AWS.",
			Required:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.RequiresReplace(),
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
		Description: "Resource Type definition for AWS::DirectoryService::SimpleAD",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::DirectoryService::SimpleAD").WithTerraformTypeName("awscc_directoryservice_simple_ad")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"alias":            "Alias",
		"create_alias":     "CreateAlias",
		"description":      "Description",
		"directory_id":     "DirectoryId",
		"dns_ip_addresses": "DnsIpAddresses",
		"enable_sso":       "EnableSso",
		"name":             "Name",
		"password":         "Password",
		"short_name":       "ShortName",
		"size":             "Size",
		"subnet_ids":       "SubnetIds",
		"vpc_id":           "VpcId",
		"vpc_settings":     "VpcSettings",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Password",
	})
	opts = opts.WithCreateTimeoutInMinutes(60).WithDeleteTimeoutInMinutes(100)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
