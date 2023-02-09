// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package route53resolver

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_route53resolver_resolver_query_logging_config", resolverQueryLoggingConfigDataSource)
}

// resolverQueryLoggingConfigDataSource returns the Terraform awscc_route53resolver_resolver_query_logging_config data source.
// This Terraform data source corresponds to the CloudFormation AWS::Route53Resolver::ResolverQueryLoggingConfig resource.
func resolverQueryLoggingConfigDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Arn",
		//	  "maxLength": 600,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Arn",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AssociationCount
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Count",
		//	  "type": "integer"
		//	}
		"association_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Count",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CreationTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Rfc3339TimeString",
		//	  "maxLength": 40,
		//	  "minLength": 20,
		//	  "type": "string"
		//	}
		"creation_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Rfc3339TimeString",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CreatorRequestId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The id of the creator request.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"creator_request_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The id of the creator request.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DestinationArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "destination arn",
		//	  "maxLength": 600,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"destination_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "destination arn",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ResourceId",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "ResourceId",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ResolverQueryLogConfigName",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "ResolverQueryLogConfigName",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: OwnerId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "AccountId",
		//	  "maxLength": 32,
		//	  "minLength": 12,
		//	  "type": "string"
		//	}
		"owner_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "AccountId",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ShareStatus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ShareStatus, possible values are NOT_SHARED, SHARED_WITH_ME, SHARED_BY_ME.",
		//	  "enum": [
		//	    "NOT_SHARED",
		//	    "SHARED_WITH_ME",
		//	    "SHARED_BY_ME"
		//	  ],
		//	  "type": "string"
		//	}
		"share_status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "ShareStatus, possible values are NOT_SHARED, SHARED_WITH_ME, SHARED_BY_ME.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ResolverQueryLogConfigStatus, possible values are CREATING, CREATED, DELETED AND FAILED.",
		//	  "enum": [
		//	    "CREATING",
		//	    "CREATED",
		//	    "DELETING",
		//	    "FAILED"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "ResolverQueryLogConfigStatus, possible values are CREATING, CREATED, DELETED AND FAILED.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Route53Resolver::ResolverQueryLoggingConfig",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53Resolver::ResolverQueryLoggingConfig").WithTerraformTypeName("awscc_route53resolver_resolver_query_logging_config")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                "Arn",
		"association_count":  "AssociationCount",
		"creation_time":      "CreationTime",
		"creator_request_id": "CreatorRequestId",
		"destination_arn":    "DestinationArn",
		"id":                 "Id",
		"name":               "Name",
		"owner_id":           "OwnerId",
		"share_status":       "ShareStatus",
		"status":             "Status",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
