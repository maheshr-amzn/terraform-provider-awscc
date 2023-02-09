// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package glue

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"regexp"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_glue_schema", schemaResource)
}

// schemaResource returns the Terraform awscc_glue_schema resource.
// This Terraform resource corresponds to the CloudFormation AWS::Glue::Schema resource.
func schemaResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Amazon Resource Name for the Schema.",
		//	  "pattern": "arn:(aws|aws-us-gov|aws-cn):glue:.*",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Amazon Resource Name for the Schema.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CheckpointVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Specify checkpoint version for update. This is only required to update the Compatibility.",
		//	  "properties": {
		//	    "IsLatest": {
		//	      "description": "Indicates if the latest version needs to be updated.",
		//	      "type": "boolean"
		//	    },
		//	    "VersionNumber": {
		//	      "description": "Indicates the version number in the schema to update.",
		//	      "maximum": 100000,
		//	      "minimum": 1,
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"checkpoint_version": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: IsLatest
				"is_latest": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Indicates if the latest version needs to be updated.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: VersionNumber
				"version_number": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "Indicates the version number in the schema to update.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Int64{ /*START VALIDATORS*/
						int64validator.Between(1, 100000),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Specify checkpoint version for update. This is only required to update the Compatibility.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Compatibility
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Compatibility setting for the schema.",
		//	  "enum": [
		//	    "NONE",
		//	    "DISABLED",
		//	    "BACKWARD",
		//	    "BACKWARD_ALL",
		//	    "FORWARD",
		//	    "FORWARD_ALL",
		//	    "FULL",
		//	    "FULL_ALL"
		//	  ],
		//	  "type": "string"
		//	}
		"compatibility": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Compatibility setting for the schema.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"NONE",
					"DISABLED",
					"BACKWARD",
					"BACKWARD_ALL",
					"FORWARD",
					"FORWARD_ALL",
					"FULL",
					"FULL_ALL",
				),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: DataFormat
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Data format name to use for the schema. Accepted values: 'AVRO', 'JSON', 'PROTOBUF'",
		//	  "enum": [
		//	    "AVRO",
		//	    "JSON",
		//	    "PROTOBUF"
		//	  ],
		//	  "type": "string"
		//	}
		"data_format": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Data format name to use for the schema. Accepted values: 'AVRO', 'JSON', 'PROTOBUF'",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"AVRO",
					"JSON",
					"PROTOBUF",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description of the schema. If description is not provided, there will not be any default value for this.",
		//	  "maxLength": 1000,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description of the schema. If description is not provided, there will not be any default value for this.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(0, 1000),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: InitialSchemaVersionId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Represents the version ID associated with the initial schema version.",
		//	  "pattern": "[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}",
		//	  "type": "string"
		//	}
		"initial_schema_version_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Represents the version ID associated with the initial schema version.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name of the schema.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name of the schema.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 255),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Registry
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Identifier for the registry which the schema is part of.",
		//	  "properties": {
		//	    "Arn": {
		//	      "description": "Amazon Resource Name for the Registry.",
		//	      "pattern": "arn:(aws|aws-us-gov|aws-cn):glue:.*",
		//	      "type": "string"
		//	    },
		//	    "Name": {
		//	      "description": "Name of the registry in which the schema will be created.",
		//	      "maxLength": 255,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"registry": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Arn
				"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Amazon Resource Name for the Registry.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.RegexMatches(regexp.MustCompile("arn:(aws|aws-us-gov|aws-cn):glue:.*"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Name
				"name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Name of the registry in which the schema will be created.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 255),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Identifier for the registry which the schema is part of.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SchemaDefinition
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Definition for the initial schema version in plain-text.",
		//	  "maxLength": 170000,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"schema_definition": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Definition for the initial schema version in plain-text.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 170000),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// SchemaDefinition is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "List of tags to tag the schema",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "description": "A key to identify the tag.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "Corresponding tag value for the key.",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 10,
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "A key to identify the tag.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Corresponding tag value for the key.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "List of tags to tag the schema",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(0, 10),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
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
		Description: "This resource represents a schema of Glue Schema Registry.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Glue::Schema").WithTerraformTypeName("awscc_glue_schema")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                       "Arn",
		"checkpoint_version":        "CheckpointVersion",
		"compatibility":             "Compatibility",
		"data_format":               "DataFormat",
		"description":               "Description",
		"initial_schema_version_id": "InitialSchemaVersionId",
		"is_latest":                 "IsLatest",
		"key":                       "Key",
		"name":                      "Name",
		"registry":                  "Registry",
		"schema_definition":         "SchemaDefinition",
		"tags":                      "Tags",
		"value":                     "Value",
		"version_number":            "VersionNumber",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/SchemaDefinition",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
