// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_quicksight_dashboard", dashboardDataSource)
}

// dashboardDataSource returns the Terraform awscc_quicksight_dashboard data source.
// This Terraform data source corresponds to the CloudFormation AWS::QuickSight::Dashboard resource.
func dashboardDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the resource.\u003c/p\u003e",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "<p>The Amazon Resource Name (ARN) of the resource.</p>",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AwsAccountId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 12,
		//	  "minLength": 12,
		//	  "pattern": "^[0-9]{12}$",
		//	  "type": "string"
		//	}
		"aws_account_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: CreatedTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eThe time that this dataset was created.\u003c/p\u003e",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"created_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "<p>The time that this dataset was created.</p>",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DashboardId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 2048,
		//	  "minLength": 1,
		//	  "pattern": "[\\w\\-]+",
		//	  "type": "string"
		//	}
		"dashboard_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: DashboardPublishOptions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "\u003cp\u003eDashboard publish options.\u003c/p\u003e",
		//	  "properties": {
		//	    "AdHocFilteringOption": {
		//	      "additionalProperties": false,
		//	      "description": "\u003cp\u003eAd hoc (one-time) filtering option.\u003c/p\u003e",
		//	      "properties": {
		//	        "AvailabilityStatus": {
		//	          "enum": [
		//	            "ENABLED",
		//	            "DISABLED"
		//	          ],
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "ExportToCSVOption": {
		//	      "additionalProperties": false,
		//	      "description": "\u003cp\u003eExport to .csv option.\u003c/p\u003e",
		//	      "properties": {
		//	        "AvailabilityStatus": {
		//	          "enum": [
		//	            "ENABLED",
		//	            "DISABLED"
		//	          ],
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "SheetControlsOption": {
		//	      "additionalProperties": false,
		//	      "description": "\u003cp\u003eSheet controls option.\u003c/p\u003e",
		//	      "properties": {
		//	        "VisibilityState": {
		//	          "enum": [
		//	            "EXPANDED",
		//	            "COLLAPSED"
		//	          ],
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"dashboard_publish_options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AdHocFilteringOption
				"ad_hoc_filtering_option": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: AvailabilityStatus
						"availability_status": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "<p>Ad hoc (one-time) filtering option.</p>",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ExportToCSVOption
				"export_to_csv_option": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: AvailabilityStatus
						"availability_status": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "<p>Export to .csv option.</p>",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: SheetControlsOption
				"sheet_controls_option": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: VisibilityState
						"visibility_state": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "<p>Sheet controls option.</p>",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "<p>Dashboard publish options.</p>",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LastPublishedTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eThe last time that this dataset was published.\u003c/p\u003e",
		//	  "format": "string",
		//	  "type": "string"
		//	}
		"last_published_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "<p>The last time that this dataset was published.</p>",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LastUpdatedTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eThe last time that this dataset was updated.\u003c/p\u003e",
		//	  "format": "string",
		//	  "type": "string"
		//	}
		"last_updated_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "<p>The last time that this dataset was updated.</p>",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eThe display name of the dashboard.\u003c/p\u003e",
		//	  "maxLength": 2048,
		//	  "minLength": 1,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "<p>The display name of the dashboard.</p>",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Parameters
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "\u003cp\u003eA list of QuickSight parameters and the list's override values.\u003c/p\u003e",
		//	  "properties": {
		//	    "DateTimeParameters": {
		//	      "description": "\u003cp\u003eDate-time parameters.\u003c/p\u003e",
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "\u003cp\u003eA date-time parameter.\u003c/p\u003e",
		//	        "properties": {
		//	          "Name": {
		//	            "description": "\u003cp\u003eA display name for the date-time parameter.\u003c/p\u003e",
		//	            "pattern": ".*\\S.*",
		//	            "type": "string"
		//	          },
		//	          "Values": {
		//	            "description": "\u003cp\u003eThe values for the date-time parameter.\u003c/p\u003e",
		//	            "items": {
		//	              "type": "string"
		//	            },
		//	            "type": "array"
		//	          }
		//	        },
		//	        "required": [
		//	          "Name",
		//	          "Values"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "maxItems": 100,
		//	      "minItems": 0,
		//	      "type": "array"
		//	    },
		//	    "DecimalParameters": {
		//	      "description": "\u003cp\u003eDecimal parameters.\u003c/p\u003e",
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "\u003cp\u003eA decimal parameter.\u003c/p\u003e",
		//	        "properties": {
		//	          "Name": {
		//	            "description": "\u003cp\u003eA display name for the decimal parameter.\u003c/p\u003e",
		//	            "pattern": ".*\\S.*",
		//	            "type": "string"
		//	          },
		//	          "Values": {
		//	            "description": "\u003cp\u003eThe values for the decimal parameter.\u003c/p\u003e",
		//	            "items": {
		//	              "type": "number"
		//	            },
		//	            "type": "array"
		//	          }
		//	        },
		//	        "required": [
		//	          "Name",
		//	          "Values"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "maxItems": 100,
		//	      "minItems": 0,
		//	      "type": "array"
		//	    },
		//	    "IntegerParameters": {
		//	      "description": "\u003cp\u003eInteger parameters.\u003c/p\u003e",
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "\u003cp\u003eAn integer parameter.\u003c/p\u003e",
		//	        "properties": {
		//	          "Name": {
		//	            "description": "\u003cp\u003eThe name of the integer parameter.\u003c/p\u003e",
		//	            "pattern": ".*\\S.*",
		//	            "type": "string"
		//	          },
		//	          "Values": {
		//	            "description": "\u003cp\u003eThe values for the integer parameter.\u003c/p\u003e",
		//	            "items": {
		//	              "type": "number"
		//	            },
		//	            "type": "array"
		//	          }
		//	        },
		//	        "required": [
		//	          "Name",
		//	          "Values"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "maxItems": 100,
		//	      "minItems": 0,
		//	      "type": "array"
		//	    },
		//	    "StringParameters": {
		//	      "description": "\u003cp\u003eString parameters.\u003c/p\u003e",
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "\u003cp\u003eA string parameter.\u003c/p\u003e",
		//	        "properties": {
		//	          "Name": {
		//	            "description": "\u003cp\u003eA display name for a string parameter.\u003c/p\u003e",
		//	            "pattern": ".*\\S.*",
		//	            "type": "string"
		//	          },
		//	          "Values": {
		//	            "description": "\u003cp\u003eThe values of a string parameter.\u003c/p\u003e",
		//	            "items": {
		//	              "type": "string"
		//	            },
		//	            "type": "array"
		//	          }
		//	        },
		//	        "required": [
		//	          "Name",
		//	          "Values"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "maxItems": 100,
		//	      "minItems": 0,
		//	      "type": "array"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"parameters": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DateTimeParameters
				"date_time_parameters": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Name
							"name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "<p>A display name for the date-time parameter.</p>",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Values
							"values": schema.ListAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Description: "<p>The values for the date-time parameter.</p>",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "<p>Date-time parameters.</p>",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: DecimalParameters
				"decimal_parameters": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Name
							"name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "<p>A display name for the decimal parameter.</p>",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Values
							"values": schema.ListAttribute{ /*START ATTRIBUTE*/
								ElementType: types.Float64Type,
								Description: "<p>The values for the decimal parameter.</p>",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "<p>Decimal parameters.</p>",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: IntegerParameters
				"integer_parameters": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Name
							"name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "<p>The name of the integer parameter.</p>",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Values
							"values": schema.ListAttribute{ /*START ATTRIBUTE*/
								ElementType: types.Float64Type,
								Description: "<p>The values for the integer parameter.</p>",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "<p>Integer parameters.</p>",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: StringParameters
				"string_parameters": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Name
							"name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "<p>A display name for a string parameter.</p>",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Values
							"values": schema.ListAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Description: "<p>The values of a string parameter.</p>",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "<p>String parameters.</p>",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "<p>A list of QuickSight parameters and the list's override values.</p>",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Permissions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eA structure that contains the permissions of the dashboard. You can use this structure\n            for granting permissions by providing a list of IAM action information for each\n            principal ARN. \u003c/p\u003e\n\n        \u003cp\u003eTo specify no permissions, omit the permissions list.\u003c/p\u003e",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "\u003cp\u003ePermission for the resource.\u003c/p\u003e",
		//	    "properties": {
		//	      "Actions": {
		//	        "description": "\u003cp\u003eThe IAM action to grant or revoke permissions on.\u003c/p\u003e",
		//	        "items": {
		//	          "type": "string"
		//	        },
		//	        "maxItems": 16,
		//	        "minItems": 1,
		//	        "type": "array"
		//	      },
		//	      "Principal": {
		//	        "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the principal. This can be one of the\n            following:\u003c/p\u003e\n        \u003cul\u003e\n            \u003cli\u003e\n                \u003cp\u003eThe ARN of an Amazon QuickSight user or group associated with a data source or dataset. (This is common.)\u003c/p\u003e\n            \u003c/li\u003e\n            \u003cli\u003e\n                \u003cp\u003eThe ARN of an Amazon QuickSight user, group, or namespace associated with an analysis, dashboard, template, or theme. (This is common.)\u003c/p\u003e\n            \u003c/li\u003e\n            \u003cli\u003e\n                \u003cp\u003eThe ARN of an AWS account root: This is an IAM ARN rather than a QuickSight\n                    ARN. Use this option only to share resources (templates) across AWS accounts.\n                    (This is less common.) \u003c/p\u003e\n            \u003c/li\u003e\n         \u003c/ul\u003e",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Actions",
		//	      "Principal"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 64,
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"permissions": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Actions
					"actions": schema.ListAttribute{ /*START ATTRIBUTE*/
						ElementType: types.StringType,
						Description: "<p>The IAM action to grant or revoke permissions on.</p>",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Principal
					"principal": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "<p>The Amazon Resource Name (ARN) of the principal. This can be one of the\n            following:</p>\n        <ul>\n            <li>\n                <p>The ARN of an Amazon QuickSight user or group associated with a data source or dataset. (This is common.)</p>\n            </li>\n            <li>\n                <p>The ARN of an Amazon QuickSight user, group, or namespace associated with an analysis, dashboard, template, or theme. (This is common.)</p>\n            </li>\n            <li>\n                <p>The ARN of an AWS account root: This is an IAM ARN rather than a QuickSight\n                    ARN. Use this option only to share resources (templates) across AWS accounts.\n                    (This is less common.) </p>\n            </li>\n         </ul>",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "<p>A structure that contains the permissions of the dashboard. You can use this structure\n            for granting permissions by providing a list of IAM action information for each\n            principal ARN. </p>\n\n        <p>To specify no permissions, omit the permissions list.</p>",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SourceEntity
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "\u003cp\u003eDashboard source entity.\u003c/p\u003e",
		//	  "properties": {
		//	    "SourceTemplate": {
		//	      "additionalProperties": false,
		//	      "description": "\u003cp\u003eDashboard source template.\u003c/p\u003e",
		//	      "properties": {
		//	        "Arn": {
		//	          "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the resource.\u003c/p\u003e",
		//	          "type": "string"
		//	        },
		//	        "DataSetReferences": {
		//	          "description": "\u003cp\u003eDataset references.\u003c/p\u003e",
		//	          "items": {
		//	            "additionalProperties": false,
		//	            "description": "\u003cp\u003eDataset reference.\u003c/p\u003e",
		//	            "properties": {
		//	              "DataSetArn": {
		//	                "description": "\u003cp\u003eDataset Amazon Resource Name (ARN).\u003c/p\u003e",
		//	                "type": "string"
		//	              },
		//	              "DataSetPlaceholder": {
		//	                "description": "\u003cp\u003eDataset placeholder.\u003c/p\u003e",
		//	                "pattern": ".*\\S.*",
		//	                "type": "string"
		//	              }
		//	            },
		//	            "required": [
		//	              "DataSetArn",
		//	              "DataSetPlaceholder"
		//	            ],
		//	            "type": "object"
		//	          },
		//	          "minItems": 1,
		//	          "type": "array"
		//	        }
		//	      },
		//	      "required": [
		//	        "Arn",
		//	        "DataSetReferences"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"source_entity": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: SourceTemplate
				"source_template": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Arn
						"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "<p>The Amazon Resource Name (ARN) of the resource.</p>",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: DataSetReferences
						"data_set_references": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
							NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: DataSetArn
									"data_set_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "<p>Dataset Amazon Resource Name (ARN).</p>",
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: DataSetPlaceholder
									"data_set_placeholder": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "<p>Dataset placeholder.</p>",
										Computed:    true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
							}, /*END NESTED OBJECT*/
							Description: "<p>Dataset references.</p>",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "<p>Dashboard source template.</p>",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "<p>Dashboard source entity.</p>",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eContains a map of the key-value pairs for the resource tag or tags assigned to the\n            dashboard.\u003c/p\u003e",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "\u003cp\u003eThe key or keys of the key-value pairs for the resource tag or tags assigned to the\n            resource.\u003c/p\u003e",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "\u003cp\u003eTag key.\u003c/p\u003e",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "\u003cp\u003eTag value.\u003c/p\u003e",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 200,
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "<p>Tag key.</p>",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "<p>Tag value.</p>",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "<p>Contains a map of the key-value pairs for the resource tag or tags assigned to the\n            dashboard.</p>",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ThemeArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the theme that is being used for this dashboard. If\n            you add a value for this field, it overrides the value that is used in the source\n            entity. The theme ARN must exist in the same AWS account where you create the\n            dashboard.\u003c/p\u003e",
		//	  "type": "string"
		//	}
		"theme_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "<p>The Amazon Resource Name (ARN) of the theme that is being used for this dashboard. If\n            you add a value for this field, it overrides the value that is used in the source\n            entity. The theme ARN must exist in the same AWS account where you create the\n            dashboard.</p>",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Version
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "\u003cp\u003eDashboard version.\u003c/p\u003e",
		//	  "properties": {
		//	    "Arn": {
		//	      "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the resource.\u003c/p\u003e",
		//	      "type": "string"
		//	    },
		//	    "CreatedTime": {
		//	      "description": "\u003cp\u003eThe time that this dashboard version was created.\u003c/p\u003e",
		//	      "format": "string",
		//	      "type": "string"
		//	    },
		//	    "DataSetArns": {
		//	      "description": "\u003cp\u003eThe Amazon Resource Numbers (ARNs) for the datasets that are associated with this\n            version of the dashboard.\u003c/p\u003e",
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "maxItems": 100,
		//	      "minItems": 0,
		//	      "type": "array"
		//	    },
		//	    "Description": {
		//	      "description": "\u003cp\u003eDescription.\u003c/p\u003e",
		//	      "maxLength": 512,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    },
		//	    "Errors": {
		//	      "description": "\u003cp\u003eErrors associated with this dashboard version.\u003c/p\u003e",
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "\u003cp\u003eDashboard error.\u003c/p\u003e",
		//	        "properties": {
		//	          "Message": {
		//	            "description": "\u003cp\u003eMessage.\u003c/p\u003e",
		//	            "pattern": ".*\\S.*",
		//	            "type": "string"
		//	          },
		//	          "Type": {
		//	            "enum": [
		//	              "ACCESS_DENIED",
		//	              "SOURCE_NOT_FOUND",
		//	              "DATA_SET_NOT_FOUND",
		//	              "INTERNAL_FAILURE",
		//	              "PARAMETER_VALUE_INCOMPATIBLE",
		//	              "PARAMETER_TYPE_INVALID",
		//	              "PARAMETER_NOT_FOUND",
		//	              "COLUMN_TYPE_MISMATCH",
		//	              "COLUMN_GEOGRAPHIC_ROLE_MISMATCH",
		//	              "COLUMN_REPLACEMENT_MISSING"
		//	            ],
		//	            "type": "string"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "minItems": 1,
		//	      "type": "array"
		//	    },
		//	    "Sheets": {
		//	      "description": "\u003cp\u003eA list of the associated sheets with the unique identifier and name of each sheet.\u003c/p\u003e",
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "\u003cp\u003eA \u003ci\u003esheet\u003c/i\u003e, which is an object that contains a set of visuals that\n            are viewed together on one page in the Amazon QuickSight console. Every analysis and dashboard\n            contains at least one sheet. Each sheet contains at least one visualization widget, for\n            example a chart, pivot table, or narrative insight. Sheets can be associated with other\n            components, such as controls, filters, and so on.\u003c/p\u003e",
		//	        "properties": {
		//	          "Name": {
		//	            "description": "\u003cp\u003eThe name of a sheet. This name is displayed on the sheet's tab in the QuickSight\n            console.\u003c/p\u003e",
		//	            "pattern": ".*\\S.*",
		//	            "type": "string"
		//	          },
		//	          "SheetId": {
		//	            "description": "\u003cp\u003eThe unique identifier associated with a sheet.\u003c/p\u003e",
		//	            "maxLength": 2048,
		//	            "minLength": 1,
		//	            "pattern": "[\\w\\-]+",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "maxItems": 20,
		//	      "minItems": 0,
		//	      "type": "array"
		//	    },
		//	    "SourceEntityArn": {
		//	      "description": "\u003cp\u003eSource entity ARN.\u003c/p\u003e",
		//	      "type": "string"
		//	    },
		//	    "Status": {
		//	      "enum": [
		//	        "CREATION_IN_PROGRESS",
		//	        "CREATION_SUCCESSFUL",
		//	        "CREATION_FAILED",
		//	        "UPDATE_IN_PROGRESS",
		//	        "UPDATE_SUCCESSFUL",
		//	        "UPDATE_FAILED",
		//	        "DELETED"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "ThemeArn": {
		//	      "description": "\u003cp\u003eThe ARN of the theme associated with a version of the dashboard.\u003c/p\u003e",
		//	      "type": "string"
		//	    },
		//	    "VersionNumber": {
		//	      "description": "\u003cp\u003eVersion number for this version of the dashboard.\u003c/p\u003e",
		//	      "minimum": 1,
		//	      "type": "number"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"version": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Arn
				"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "<p>The Amazon Resource Name (ARN) of the resource.</p>",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: CreatedTime
				"created_time": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "<p>The time that this dashboard version was created.</p>",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: DataSetArns
				"data_set_arns": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "<p>The Amazon Resource Numbers (ARNs) for the datasets that are associated with this\n            version of the dashboard.</p>",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Description
				"description": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "<p>Description.</p>",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Errors
				"errors": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Message
							"message": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "<p>Message.</p>",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Type
							"type": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "<p>Errors associated with this dashboard version.</p>",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Sheets
				"sheets": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Name
							"name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "<p>The name of a sheet. This name is displayed on the sheet's tab in the QuickSight\n            console.</p>",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: SheetId
							"sheet_id": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "<p>The unique identifier associated with a sheet.</p>",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "<p>A list of the associated sheets with the unique identifier and name of each sheet.</p>",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: SourceEntityArn
				"source_entity_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "<p>Source entity ARN.</p>",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Status
				"status": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: ThemeArn
				"theme_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "<p>The ARN of the theme associated with a version of the dashboard.</p>",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: VersionNumber
				"version_number": schema.Float64Attribute{ /*START ATTRIBUTE*/
					Description: "<p>Version number for this version of the dashboard.</p>",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "<p>Dashboard version.</p>",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VersionDescription
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eA description for the first version of the dashboard being created.\u003c/p\u003e",
		//	  "maxLength": 512,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"version_description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "<p>A description for the first version of the dashboard being created.</p>",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::QuickSight::Dashboard",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::QuickSight::Dashboard").WithTerraformTypeName("awscc_quicksight_dashboard")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"actions":                   "Actions",
		"ad_hoc_filtering_option":   "AdHocFilteringOption",
		"arn":                       "Arn",
		"availability_status":       "AvailabilityStatus",
		"aws_account_id":            "AwsAccountId",
		"created_time":              "CreatedTime",
		"dashboard_id":              "DashboardId",
		"dashboard_publish_options": "DashboardPublishOptions",
		"data_set_arn":              "DataSetArn",
		"data_set_arns":             "DataSetArns",
		"data_set_placeholder":      "DataSetPlaceholder",
		"data_set_references":       "DataSetReferences",
		"date_time_parameters":      "DateTimeParameters",
		"decimal_parameters":        "DecimalParameters",
		"description":               "Description",
		"errors":                    "Errors",
		"export_to_csv_option":      "ExportToCSVOption",
		"integer_parameters":        "IntegerParameters",
		"key":                       "Key",
		"last_published_time":       "LastPublishedTime",
		"last_updated_time":         "LastUpdatedTime",
		"message":                   "Message",
		"name":                      "Name",
		"parameters":                "Parameters",
		"permissions":               "Permissions",
		"principal":                 "Principal",
		"sheet_controls_option":     "SheetControlsOption",
		"sheet_id":                  "SheetId",
		"sheets":                    "Sheets",
		"source_entity":             "SourceEntity",
		"source_entity_arn":         "SourceEntityArn",
		"source_template":           "SourceTemplate",
		"status":                    "Status",
		"string_parameters":         "StringParameters",
		"tags":                      "Tags",
		"theme_arn":                 "ThemeArn",
		"type":                      "Type",
		"value":                     "Value",
		"values":                    "Values",
		"version":                   "Version",
		"version_description":       "VersionDescription",
		"version_number":            "VersionNumber",
		"visibility_state":          "VisibilityState",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
