// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package transfer

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"regexp"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_transfer_certificate", certificateResource)
}

// certificateResource returns the Terraform awscc_transfer_certificate resource.
// This Terraform resource corresponds to the CloudFormation AWS::Transfer::Certificate resource.
func certificateResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ActiveDate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the active date for the certificate.",
		//	  "type": "string"
		//	}
		"active_date": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the active date for the certificate.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the unique Amazon Resource Name (ARN) for the agreement.",
		//	  "maxLength": 1600,
		//	  "minLength": 20,
		//	  "pattern": "arn:.*",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the unique Amazon Resource Name (ARN) for the agreement.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Certificate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the certificate body to be imported.",
		//	  "maxLength": 16384,
		//	  "minLength": 1,
		//	  "pattern": "^[\t\n\r -ÿ]*",
		//	  "type": "string"
		//	}
		"certificate": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the certificate body to be imported.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 16384),
				stringvalidator.RegexMatches(regexp.MustCompile("^[\t\n\r -ÿ]*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CertificateChain
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the certificate chain to be imported.",
		//	  "maxLength": 2097152,
		//	  "minLength": 1,
		//	  "pattern": "^[\t\n\r -ÿ]*",
		//	  "type": "string"
		//	}
		"certificate_chain": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the certificate chain to be imported.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 2097152),
				stringvalidator.RegexMatches(regexp.MustCompile("^[\t\n\r -ÿ]*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CertificateId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A unique identifier for the certificate.",
		//	  "maxLength": 22,
		//	  "minLength": 22,
		//	  "pattern": "^cert-([0-9a-f]{17})$",
		//	  "type": "string"
		//	}
		"certificate_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A unique identifier for the certificate.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A textual description for the certificate.",
		//	  "maxLength": 200,
		//	  "minLength": 1,
		//	  "pattern": "^[\\w\\- ]*$",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A textual description for the certificate.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 200),
				stringvalidator.RegexMatches(regexp.MustCompile("^[\\w\\- ]*$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: InactiveDate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the inactive date for the certificate.",
		//	  "type": "string"
		//	}
		"inactive_date": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the inactive date for the certificate.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: NotAfterDate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the not after date for the certificate.",
		//	  "type": "string"
		//	}
		"not_after_date": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the not after date for the certificate.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: NotBeforeDate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the not before date for the certificate.",
		//	  "type": "string"
		//	}
		"not_before_date": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the not before date for the certificate.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PrivateKey
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the private key for the certificate.",
		//	  "maxLength": 16384,
		//	  "minLength": 1,
		//	  "pattern": "^[\t\n\r -ÿ]*",
		//	  "type": "string"
		//	}
		"private_key": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the private key for the certificate.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 16384),
				stringvalidator.RegexMatches(regexp.MustCompile("^[\t\n\r -ÿ]*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// PrivateKey is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Serial
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies Certificate's serial.",
		//	  "maxLength": 48,
		//	  "minLength": 0,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"serial": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies Certificate's serial.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A status description for the certificate.",
		//	  "enum": [
		//	    "ACTIVE",
		//	    "PENDING",
		//	    "INACTIVE"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A status description for the certificate.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Key-value pairs that can be used to group and search for certificates. Tags are metadata attached to certificates for any purpose.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
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
		//	  "maxItems": 50,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Key-value pairs that can be used to group and search for certificates. Tags are metadata attached to certificates for any purpose.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.SizeAtMost(50),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Describing the type of certificate. With or without a private key.",
		//	  "enum": [
		//	    "CERTIFICATE",
		//	    "CERTIFICATE_WITH_PRIVATE_KEY"
		//	  ],
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Describing the type of certificate. With or without a private key.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Usage
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the usage type for the certificate.",
		//	  "enum": [
		//	    "SIGNING",
		//	    "ENCRYPTION"
		//	  ],
		//	  "type": "string"
		//	}
		"usage": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the usage type for the certificate.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"SIGNING",
					"ENCRYPTION",
				),
			}, /*END VALIDATORS*/
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
		Description: "Resource Type definition for AWS::Transfer::Certificate",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Transfer::Certificate").WithTerraformTypeName("awscc_transfer_certificate")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"active_date":       "ActiveDate",
		"arn":               "Arn",
		"certificate":       "Certificate",
		"certificate_chain": "CertificateChain",
		"certificate_id":    "CertificateId",
		"description":       "Description",
		"inactive_date":     "InactiveDate",
		"key":               "Key",
		"not_after_date":    "NotAfterDate",
		"not_before_date":   "NotBeforeDate",
		"private_key":       "PrivateKey",
		"serial":            "Serial",
		"status":            "Status",
		"tags":              "Tags",
		"type":              "Type",
		"usage":             "Usage",
		"value":             "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/PrivateKey",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
