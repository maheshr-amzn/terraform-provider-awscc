// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package robomaker

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"regexp"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_robomaker_simulation_application_version", simulationApplicationVersionResource)
}

// simulationApplicationVersionResource returns the Terraform awscc_robomaker_simulation_application_version resource.
// This Terraform resource corresponds to the CloudFormation AWS::RoboMaker::SimulationApplicationVersion resource.
func simulationApplicationVersionResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Application
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "arn:[\\w+=/,.@-]+:[\\w+=/,.@-]+:[\\w+=/,.@-]*:[0-9]*:[\\w+=,.@-]+(/[\\w+=,.@-]+)*",
		//	  "type": "string"
		//	}
		"application": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("arn:[\\w+=/,.@-]+:[\\w+=/,.@-]+:[\\w+=/,.@-]*:[0-9]*:[\\w+=,.@-]+(/[\\w+=,.@-]+)*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ApplicationVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"application_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "arn:[\\w+=/,.@-]+:[\\w+=/,.@-]+:[\\w+=/,.@-]*:[0-9]*:[\\w+=,.@-]+(/[\\w+=,.@-]+)*",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CurrentRevisionId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The revision ID of robot application.",
		//	  "maxLength": 40,
		//	  "minLength": 1,
		//	  "pattern": "[a-zA-Z0-9_.\\-]*",
		//	  "type": "string"
		//	}
		"current_revision_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The revision ID of robot application.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 40),
				stringvalidator.RegexMatches(regexp.MustCompile("[a-zA-Z0-9_.\\-]*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
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
		Description: "AWS::RoboMaker::SimulationApplicationVersion resource creates an AWS RoboMaker SimulationApplicationVersion. This helps you control which code your simulation uses.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::RoboMaker::SimulationApplicationVersion").WithTerraformTypeName("awscc_robomaker_simulation_application_version")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"application":         "Application",
		"application_version": "ApplicationVersion",
		"arn":                 "Arn",
		"current_revision_id": "CurrentRevisionId",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
