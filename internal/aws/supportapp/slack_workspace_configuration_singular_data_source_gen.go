// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package supportapp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_supportapp_slack_workspace_configuration", slackWorkspaceConfigurationDataSource)
}

// slackWorkspaceConfigurationDataSource returns the Terraform awscc_supportapp_slack_workspace_configuration data source.
// This Terraform data source corresponds to the CloudFormation AWS::SupportApp::SlackWorkspaceConfiguration resource.
func slackWorkspaceConfigurationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: TeamId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The team ID in Slack, which uniquely identifies a workspace.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "^\\S+$",
		//	  "type": "string"
		//	}
		"team_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The team ID in Slack, which uniquely identifies a workspace.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VersionId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An identifier used to update an existing Slack workspace configuration in AWS CloudFormation.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "^[0-9]+$",
		//	  "type": "string"
		//	}
		"version_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An identifier used to update an existing Slack workspace configuration in AWS CloudFormation.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::SupportApp::SlackWorkspaceConfiguration",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SupportApp::SlackWorkspaceConfiguration").WithTerraformTypeName("awscc_supportapp_slack_workspace_configuration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"team_id":    "TeamId",
		"version_id": "VersionId",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
