// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package networkmanager

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_networkmanager_connect_attachment", connectAttachmentDataSource)
}

// connectAttachmentDataSource returns the Terraform awscc_networkmanager_connect_attachment data source.
// This Terraform data source corresponds to the CloudFormation AWS::NetworkManager::ConnectAttachment resource.
func connectAttachmentDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AttachmentId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the attachment.",
		//	  "type": "string"
		//	}
		"attachment_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the attachment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AttachmentPolicyRuleNumber
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The policy rule number associated with the attachment.",
		//	  "type": "integer"
		//	}
		"attachment_policy_rule_number": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The policy rule number associated with the attachment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AttachmentType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of attachment.",
		//	  "type": "string"
		//	}
		"attachment_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of attachment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CoreNetworkArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of a core network for the VPC attachment.",
		//	  "type": "string"
		//	}
		"core_network_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of a core network for the VPC attachment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CoreNetworkId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ID of the CoreNetwork that the attachment will be attached to.",
		//	  "type": "string"
		//	}
		"core_network_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "ID of the CoreNetwork that the attachment will be attached to.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Creation time of the attachment.",
		//	  "type": "string"
		//	}
		"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Creation time of the attachment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EdgeLocation
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Edge location of the attachment.",
		//	  "type": "string"
		//	}
		"edge_location": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Edge location of the attachment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Options
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Protocol options for connect attachment",
		//	  "properties": {
		//	    "Protocol": {
		//	      "description": "Tunnel protocol for connect attachment",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Protocol
				"protocol": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Tunnel protocol for connect attachment",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Protocol options for connect attachment",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: OwnerAccountId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the attachment account owner.",
		//	  "type": "string"
		//	}
		"owner_account_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the attachment account owner.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ProposedSegmentChange
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The attachment to move from one segment to another.",
		//	  "properties": {
		//	    "AttachmentPolicyRuleNumber": {
		//	      "description": "New policy rule number of the attachment",
		//	      "type": "integer"
		//	    },
		//	    "SegmentName": {
		//	      "description": "Proposed segment name",
		//	      "type": "string"
		//	    },
		//	    "Tags": {
		//	      "description": "Proposed tags for the Segment.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "A key-value pair to associate with a resource.",
		//	        "properties": {
		//	          "Key": {
		//	            "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	            "type": "string"
		//	          },
		//	          "Value": {
		//	            "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "Key",
		//	          "Value"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "type": "array"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"proposed_segment_change": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AttachmentPolicyRuleNumber
				"attachment_policy_rule_number": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "New policy rule number of the attachment",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: SegmentName
				"segment_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Proposed segment name",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Tags
				"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Key
							"key": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Value
							"value": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "Proposed tags for the Segment.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The attachment to move from one segment to another.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResourceArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The attachment resource ARN.",
		//	  "type": "string"
		//	}
		"resource_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The attachment resource ARN.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SegmentName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the segment attachment.",
		//	  "type": "string"
		//	}
		"segment_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the segment attachment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: State
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "State of the attachment.",
		//	  "type": "string"
		//	}
		"state": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "State of the attachment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Tags for the attachment.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Tags for the attachment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TransportAttachmentId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Id of transport attachment",
		//	  "type": "string"
		//	}
		"transport_attachment_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Id of transport attachment",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: UpdatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Last update time of the attachment.",
		//	  "type": "string"
		//	}
		"updated_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Last update time of the attachment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::NetworkManager::ConnectAttachment",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::NetworkManager::ConnectAttachment").WithTerraformTypeName("awscc_networkmanager_connect_attachment")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"attachment_id":                 "AttachmentId",
		"attachment_policy_rule_number": "AttachmentPolicyRuleNumber",
		"attachment_type":               "AttachmentType",
		"core_network_arn":              "CoreNetworkArn",
		"core_network_id":               "CoreNetworkId",
		"created_at":                    "CreatedAt",
		"edge_location":                 "EdgeLocation",
		"key":                           "Key",
		"options":                       "Options",
		"owner_account_id":              "OwnerAccountId",
		"proposed_segment_change":       "ProposedSegmentChange",
		"protocol":                      "Protocol",
		"resource_arn":                  "ResourceArn",
		"segment_name":                  "SegmentName",
		"state":                         "State",
		"tags":                          "Tags",
		"transport_attachment_id":       "TransportAttachmentId",
		"updated_at":                    "UpdatedAt",
		"value":                         "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
