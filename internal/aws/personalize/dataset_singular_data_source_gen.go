// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package personalize

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_personalize_dataset", datasetDataSource)
}

// datasetDataSource returns the Terraform awscc_personalize_dataset data source.
// This Terraform data source corresponds to the CloudFormation AWS::Personalize::Dataset resource.
func datasetDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: DatasetArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the dataset",
		//	  "maxLength": 256,
		//	  "pattern": "arn:([a-z\\d-]+):personalize:.*:.*:.+",
		//	  "type": "string"
		//	}
		"dataset_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the dataset",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DatasetGroupArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the dataset group to add the dataset to",
		//	  "maxLength": 256,
		//	  "pattern": "arn:([a-z\\d-]+):personalize:.*:.*:.+",
		//	  "type": "string"
		//	}
		"dataset_group_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the dataset group to add the dataset to",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DatasetImportJob
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Initial DatasetImportJob for the created dataset",
		//	  "properties": {
		//	    "DataSource": {
		//	      "additionalProperties": false,
		//	      "description": "The Amazon S3 bucket that contains the training data to import.",
		//	      "properties": {
		//	        "DataLocation": {
		//	          "description": "The path to the Amazon S3 bucket where the data that you want to upload to your dataset is stored.",
		//	          "maxLength": 256,
		//	          "pattern": "(s3|http|https)://.+",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "DatasetArn": {
		//	      "description": "The ARN of the dataset that receives the imported data",
		//	      "maxLength": 256,
		//	      "pattern": "arn:([a-z\\d-]+):personalize:.*:.*:.+",
		//	      "type": "string"
		//	    },
		//	    "DatasetImportJobArn": {
		//	      "description": "The ARN of the dataset import job",
		//	      "maxLength": 256,
		//	      "pattern": "arn:([a-z\\d-]+):personalize:.*:.*:.+",
		//	      "type": "string"
		//	    },
		//	    "JobName": {
		//	      "description": "The name for the dataset import job.",
		//	      "maxLength": 63,
		//	      "minLength": 1,
		//	      "pattern": "^[a-zA-Z0-9][a-zA-Z0-9\\-_]*",
		//	      "type": "string"
		//	    },
		//	    "RoleArn": {
		//	      "description": "The ARN of the IAM role that has permissions to read from the Amazon S3 data source.",
		//	      "maxLength": 256,
		//	      "pattern": "arn:([a-z\\d-]+):iam::\\d{12}:role/?[a-zA-Z_0-9+=,.@\\-_/]+",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"dataset_import_job": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DataSource
				"data_source": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: DataLocation
						"data_location": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The path to the Amazon S3 bucket where the data that you want to upload to your dataset is stored.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The Amazon S3 bucket that contains the training data to import.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: DatasetArn
				"dataset_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ARN of the dataset that receives the imported data",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: DatasetImportJobArn
				"dataset_import_job_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ARN of the dataset import job",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: JobName
				"job_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The name for the dataset import job.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: RoleArn
				"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ARN of the IAM role that has permissions to read from the Amazon S3 data source.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Initial DatasetImportJob for the created dataset",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DatasetType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of dataset",
		//	  "enum": [
		//	    "Interactions",
		//	    "Items",
		//	    "Users"
		//	  ],
		//	  "maxLength": 256,
		//	  "type": "string"
		//	}
		"dataset_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of dataset",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name for the dataset",
		//	  "maxLength": 63,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9][a-zA-Z0-9\\-_]*",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name for the dataset",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SchemaArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the schema to associate with the dataset. The schema defines the dataset fields.",
		//	  "maxLength": 256,
		//	  "pattern": "arn:([a-z\\d-]+):personalize:.*:.*:.+",
		//	  "type": "string"
		//	}
		"schema_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the schema to associate with the dataset. The schema defines the dataset fields.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Personalize::Dataset",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Personalize::Dataset").WithTerraformTypeName("awscc_personalize_dataset")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"data_location":          "DataLocation",
		"data_source":            "DataSource",
		"dataset_arn":            "DatasetArn",
		"dataset_group_arn":      "DatasetGroupArn",
		"dataset_import_job":     "DatasetImportJob",
		"dataset_import_job_arn": "DatasetImportJobArn",
		"dataset_type":           "DatasetType",
		"job_name":               "JobName",
		"name":                   "Name",
		"role_arn":               "RoleArn",
		"schema_arn":             "SchemaArn",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
