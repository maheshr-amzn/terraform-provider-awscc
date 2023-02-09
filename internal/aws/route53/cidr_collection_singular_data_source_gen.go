// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package route53

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_route53_cidr_collection", cidrCollectionDataSource)
}

// cidrCollectionDataSource returns the Terraform awscc_route53_cidr_collection data source.
// This Terraform data source corresponds to the CloudFormation AWS::Route53::CidrCollection resource.
func cidrCollectionDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon resource name (ARN) to uniquely identify the AWS resource.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon resource name (ARN) to uniquely identify the AWS resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "UUID of the CIDR collection.",
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "UUID of the CIDR collection.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Locations
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A complex type that contains information about the list of CIDR locations.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "CidrList": {
		//	        "description": "A list of CIDR blocks.",
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "type": "string"
		//	        },
		//	        "type": "array",
		//	        "uniqueItems": true
		//	      },
		//	      "LocationName": {
		//	        "description": "The name of the location that is associated with the CIDR collection.",
		//	        "maxLength": 16,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "LocationName",
		//	      "CidrList"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"locations": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: CidrList
					"cidr_list": schema.SetAttribute{ /*START ATTRIBUTE*/
						ElementType: types.StringType,
						Description: "A list of CIDR blocks.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: LocationName
					"location_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The name of the location that is associated with the CIDR collection.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A complex type that contains information about the list of CIDR locations.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A unique name for the CIDR collection.",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "^[0-9A-Za-z_\\-]+$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A unique name for the CIDR collection.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Route53::CidrCollection",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53::CidrCollection").WithTerraformTypeName("awscc_route53_cidr_collection")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":           "Arn",
		"cidr_list":     "CidrList",
		"id":            "Id",
		"location_name": "LocationName",
		"locations":     "Locations",
		"name":          "Name",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
