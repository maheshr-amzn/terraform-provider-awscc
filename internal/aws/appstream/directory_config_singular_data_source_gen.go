// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package appstream

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_appstream_directory_config", directoryConfigDataSource)
}

// directoryConfigDataSource returns the Terraform awscc_appstream_directory_config data source.
// This Terraform data source corresponds to the CloudFormation AWS::AppStream::DirectoryConfig resource.
func directoryConfigDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CertificateBasedAuthProperties
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "CertificateAuthorityArn": {
		//	      "type": "string"
		//	    },
		//	    "Status": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"certificate_based_auth_properties": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: CertificateAuthorityArn
				"certificate_authority_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Status
				"status": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: DirectoryName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"directory_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: OrganizationalUnitDistinguishedNames
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"organizational_unit_distinguished_names": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ServiceAccountCredentials
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "AccountName": {
		//	      "type": "string"
		//	    },
		//	    "AccountPassword": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "AccountName",
		//	    "AccountPassword"
		//	  ],
		//	  "type": "object"
		//	}
		"service_account_credentials": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AccountName
				"account_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: AccountPassword
				"account_password": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::AppStream::DirectoryConfig",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::AppStream::DirectoryConfig").WithTerraformTypeName("awscc_appstream_directory_config")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account_name":                            "AccountName",
		"account_password":                        "AccountPassword",
		"certificate_authority_arn":               "CertificateAuthorityArn",
		"certificate_based_auth_properties":       "CertificateBasedAuthProperties",
		"directory_name":                          "DirectoryName",
		"organizational_unit_distinguished_names": "OrganizationalUnitDistinguishedNames",
		"service_account_credentials":             "ServiceAccountCredentials",
		"status":                                  "Status",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
