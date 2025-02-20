// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iotfleetwise

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_iotfleetwise_campaign", campaignDataSource)
}

// campaignDataSource returns the Terraform awscc_iotfleetwise_campaign data source.
// This Terraform data source corresponds to the CloudFormation AWS::IoTFleetWise::Campaign resource.
func campaignDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Action
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "APPROVE",
		//	    "SUSPEND",
		//	    "RESUME",
		//	    "UPDATE"
		//	  ],
		//	  "type": "string"
		//	}
		"action": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: CollectionScheme
		// CloudFormation resource type schema:
		//
		//	{
		//	  "properties": {
		//	    "ConditionBasedCollectionScheme": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "ConditionLanguageVersion": {
		//	          "minimum": 1,
		//	          "type": "integer"
		//	        },
		//	        "Expression": {
		//	          "maxLength": 2048,
		//	          "minLength": 1,
		//	          "type": "string"
		//	        },
		//	        "MinimumTriggerIntervalMs": {
		//	          "maximum": 4294967295,
		//	          "minimum": 0,
		//	          "type": "number"
		//	        },
		//	        "TriggerMode": {
		//	          "enum": [
		//	            "ALWAYS",
		//	            "RISING_EDGE"
		//	          ],
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "Expression"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "TimeBasedCollectionScheme": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "PeriodMs": {
		//	          "maximum": 60000,
		//	          "minimum": 10000,
		//	          "type": "number"
		//	        }
		//	      },
		//	      "required": [
		//	        "PeriodMs"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"collection_scheme": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ConditionBasedCollectionScheme
				"condition_based_collection_scheme": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: ConditionLanguageVersion
						"condition_language_version": schema.Int64Attribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: Expression
						"expression": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: MinimumTriggerIntervalMs
						"minimum_trigger_interval_ms": schema.Float64Attribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: TriggerMode
						"trigger_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: TimeBasedCollectionScheme
				"time_based_collection_scheme": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: PeriodMs
						"period_ms": schema.Float64Attribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Compression
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "OFF",
		//	  "enum": [
		//	    "OFF",
		//	    "SNAPPY"
		//	  ],
		//	  "type": "string"
		//	}
		"compression": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: CreationTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"creation_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: DataExtraDimensions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "maxLength": 150,
		//	    "minLength": 1,
		//	    "pattern": "^[a-zA-Z0-9_.]+$",
		//	    "type": "string"
		//	  },
		//	  "maxItems": 5,
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"data_extra_dimensions": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 2048,
		//	  "minLength": 1,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: DiagnosticsMode
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "OFF",
		//	  "enum": [
		//	    "OFF",
		//	    "SEND_ACTIVE_DTCS"
		//	  ],
		//	  "type": "string"
		//	}
		"diagnostics_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ExpiryTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "253402214400",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"expiry_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: LastModificationTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"last_modification_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 100,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z\\d\\-_:]+$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: PostTriggerCollectionDuration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": 0,
		//	  "maximum": 4294967295,
		//	  "minimum": 0,
		//	  "type": "number"
		//	}
		"post_trigger_collection_duration": schema.Float64Attribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Priority
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": 0,
		//	  "minimum": 0,
		//	  "type": "integer"
		//	}
		"priority": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: SignalCatalogArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"signal_catalog_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: SignalsToCollect
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "MaxSampleCount": {
		//	        "maximum": 4294967295,
		//	        "minimum": 1,
		//	        "type": "number"
		//	      },
		//	      "MinimumSamplingIntervalMs": {
		//	        "maximum": 4294967295,
		//	        "minimum": 0,
		//	        "type": "number"
		//	      },
		//	      "Name": {
		//	        "maxLength": 150,
		//	        "minLength": 1,
		//	        "pattern": "^[\\w|*|-]+(\\.[\\w|*|-]+)*$",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Name"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 1000,
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"signals_to_collect": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: MaxSampleCount
					"max_sample_count": schema.Float64Attribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: MinimumSamplingIntervalMs
					"minimum_sampling_interval_ms": schema.Float64Attribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Name
					"name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: SpoolingMode
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "OFF",
		//	  "enum": [
		//	    "OFF",
		//	    "TO_DISK"
		//	  ],
		//	  "type": "string"
		//	}
		"spooling_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: StartTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "0",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"start_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "CREATING",
		//	    "WAITING_FOR_APPROVAL",
		//	    "RUNNING",
		//	    "SUSPENDED"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
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
		//	  "minItems": 0,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
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
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: TargetArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"target_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::IoTFleetWise::Campaign",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTFleetWise::Campaign").WithTerraformTypeName("awscc_iotfleetwise_campaign")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"action":                            "Action",
		"arn":                               "Arn",
		"collection_scheme":                 "CollectionScheme",
		"compression":                       "Compression",
		"condition_based_collection_scheme": "ConditionBasedCollectionScheme",
		"condition_language_version":        "ConditionLanguageVersion",
		"creation_time":                     "CreationTime",
		"data_extra_dimensions":             "DataExtraDimensions",
		"description":                       "Description",
		"diagnostics_mode":                  "DiagnosticsMode",
		"expiry_time":                       "ExpiryTime",
		"expression":                        "Expression",
		"key":                               "Key",
		"last_modification_time":            "LastModificationTime",
		"max_sample_count":                  "MaxSampleCount",
		"minimum_sampling_interval_ms":      "MinimumSamplingIntervalMs",
		"minimum_trigger_interval_ms":       "MinimumTriggerIntervalMs",
		"name":                              "Name",
		"period_ms":                         "PeriodMs",
		"post_trigger_collection_duration":  "PostTriggerCollectionDuration",
		"priority":                          "Priority",
		"signal_catalog_arn":                "SignalCatalogArn",
		"signals_to_collect":                "SignalsToCollect",
		"spooling_mode":                     "SpoolingMode",
		"start_time":                        "StartTime",
		"status":                            "Status",
		"tags":                              "Tags",
		"target_arn":                        "TargetArn",
		"time_based_collection_scheme":      "TimeBasedCollectionScheme",
		"trigger_mode":                      "TriggerMode",
		"value":                             "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
