// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package mediapackage

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_mediapackage_channel", channelDataSource)
}

// channelDataSource returns the Terraform awscc_mediapackage_channel data source.
// This Terraform data source corresponds to the CloudFormation AWS::MediaPackage::Channel resource.
func channelDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) assigned to the Channel.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) assigned to the Channel.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A short text description of the Channel.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A short text description of the Channel.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EgressAccessLogs
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The configuration parameters for egress access logging.",
		//	  "properties": {
		//	    "LogGroupName": {
		//	      "description": "Sets a custom AWS CloudWatch log group name for access logs. If a log group name isn't specified, the defaults are used: /aws/MediaPackage/EgressAccessLogs for egress access logs and /aws/MediaPackage/IngressAccessLogs for ingress access logs.",
		//	      "maxLength": 256,
		//	      "minLength": 1,
		//	      "pattern": "",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"egress_access_logs": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: LogGroupName
				"log_group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Sets a custom AWS CloudWatch log group name for access logs. If a log group name isn't specified, the defaults are used: /aws/MediaPackage/EgressAccessLogs for egress access logs and /aws/MediaPackage/IngressAccessLogs for ingress access logs.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The configuration parameters for egress access logging.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: HlsIngest
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "An HTTP Live Streaming (HLS) ingest resource configuration.",
		//	  "properties": {
		//	    "ingestEndpoints": {
		//	      "description": "A list of endpoints to which the source stream should be sent.",
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "An endpoint for ingesting source content for a Channel.",
		//	        "properties": {
		//	          "Id": {
		//	            "description": "The system generated unique identifier for the IngestEndpoint",
		//	            "type": "string"
		//	          },
		//	          "Password": {
		//	            "description": "The system generated password for ingest authentication.",
		//	            "type": "string"
		//	          },
		//	          "Url": {
		//	            "description": "The ingest URL to which the source stream should be sent.",
		//	            "type": "string"
		//	          },
		//	          "Username": {
		//	            "description": "The system generated username for ingest authentication.",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "Id",
		//	          "Username",
		//	          "Password",
		//	          "Url"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "type": "array"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"hls_ingest": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ingestEndpoints
				"ingest_endpoints": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Id
							"id": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The system generated unique identifier for the IngestEndpoint",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Password
							"password": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The system generated password for ingest authentication.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Url
							"url": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The ingest URL to which the source stream should be sent.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Username
							"username": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The system generated username for ingest authentication.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "A list of endpoints to which the source stream should be sent.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "An HTTP Live Streaming (HLS) ingest resource configuration.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the Channel.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the Channel.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IngressAccessLogs
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The configuration parameters for egress access logging.",
		//	  "properties": {
		//	    "LogGroupName": {
		//	      "description": "Sets a custom AWS CloudWatch log group name for access logs. If a log group name isn't specified, the defaults are used: /aws/MediaPackage/EgressAccessLogs for egress access logs and /aws/MediaPackage/IngressAccessLogs for ingress access logs.",
		//	      "maxLength": 256,
		//	      "minLength": 1,
		//	      "pattern": "",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"ingress_access_logs": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: LogGroupName
				"log_group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Sets a custom AWS CloudWatch log group name for access logs. If a log group name isn't specified, the defaults are used: /aws/MediaPackage/EgressAccessLogs for egress access logs and /aws/MediaPackage/IngressAccessLogs for ingress access logs.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The configuration parameters for egress access logging.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A collection of tags associated with a resource",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A collection of tags associated with a resource",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::MediaPackage::Channel",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::MediaPackage::Channel").WithTerraformTypeName("awscc_mediapackage_channel")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                 "Arn",
		"description":         "Description",
		"egress_access_logs":  "EgressAccessLogs",
		"hls_ingest":          "HlsIngest",
		"id":                  "Id",
		"ingest_endpoints":    "ingestEndpoints",
		"ingress_access_logs": "IngressAccessLogs",
		"key":                 "Key",
		"log_group_name":      "LogGroupName",
		"password":            "Password",
		"tags":                "Tags",
		"url":                 "Url",
		"username":            "Username",
		"value":               "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
