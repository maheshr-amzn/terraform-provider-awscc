// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package eks

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_eks_nodegroup", nodegroupDataSource)
}

// nodegroupDataSource returns the Terraform awscc_eks_nodegroup data source.
// This Terraform data source corresponds to the CloudFormation AWS::EKS::Nodegroup resource.
func nodegroupDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AmiType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The AMI type for your node group.",
		//	  "type": "string"
		//	}
		"ami_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The AMI type for your node group.",
			Computed:    true,
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
		// Property: CapacityType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The capacity type of your managed node group.",
		//	  "type": "string"
		//	}
		"capacity_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The capacity type of your managed node group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ClusterName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name of the cluster to create the node group in.",
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"cluster_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name of the cluster to create the node group in.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DiskSize
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The root device disk size (in GiB) for your node group instances.",
		//	  "type": "integer"
		//	}
		"disk_size": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The root device disk size (in GiB) for your node group instances.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ForceUpdateEnabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": false,
		//	  "description": "Force the update if the existing node group's pods are unable to be drained due to a pod disruption budget issue.",
		//	  "type": "boolean"
		//	}
		"force_update_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Force the update if the existing node group's pods are unable to be drained due to a pod disruption budget issue.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: InstanceTypes
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specify the instance types for a node group.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"instance_types": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "Specify the instance types for a node group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Labels
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The Kubernetes labels to be applied to the nodes in the node group when they are created.",
		//	  "patternProperties": {
		//	    "": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"labels":            // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The Kubernetes labels to be applied to the nodes in the node group when they are created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LaunchTemplate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "An object representing a node group's launch template specification.",
		//	  "properties": {
		//	    "Id": {
		//	      "minLength": 1,
		//	      "type": "string"
		//	    },
		//	    "Name": {
		//	      "minLength": 1,
		//	      "type": "string"
		//	    },
		//	    "Version": {
		//	      "minLength": 1,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"launch_template": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Id
				"id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Name
				"name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Version
				"version": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "An object representing a node group's launch template specification.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: NodeRole
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the IAM role to associate with your node group.",
		//	  "type": "string"
		//	}
		"node_role": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the IAM role to associate with your node group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: NodegroupName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique name to give your node group.",
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"nodegroup_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The unique name to give your node group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ReleaseVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The AMI version of the Amazon EKS-optimized AMI to use with your node group.",
		//	  "type": "string"
		//	}
		"release_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The AMI version of the Amazon EKS-optimized AMI to use with your node group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RemoteAccess
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The remote access (SSH) configuration to use with your node group.",
		//	  "properties": {
		//	    "Ec2SshKey": {
		//	      "type": "string"
		//	    },
		//	    "SourceSecurityGroups": {
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    }
		//	  },
		//	  "required": [
		//	    "Ec2SshKey"
		//	  ],
		//	  "type": "object"
		//	}
		"remote_access": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Ec2SshKey
				"ec_2_ssh_key": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: SourceSecurityGroups
				"source_security_groups": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The remote access (SSH) configuration to use with your node group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ScalingConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The scaling configuration details for the Auto Scaling group that is created for your node group.",
		//	  "properties": {
		//	    "DesiredSize": {
		//	      "minimum": 0,
		//	      "type": "integer"
		//	    },
		//	    "MaxSize": {
		//	      "minimum": 1,
		//	      "type": "integer"
		//	    },
		//	    "MinSize": {
		//	      "minimum": 0,
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"scaling_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DesiredSize
				"desired_size": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: MaxSize
				"max_size": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: MinSize
				"min_size": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The scaling configuration details for the Auto Scaling group that is created for your node group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Subnets
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The subnets to use for the Auto Scaling group that is created for your node group.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"subnets": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The subnets to use for the Auto Scaling group that is created for your node group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The metadata, as key-value pairs, to apply to the node group to assist with categorization and organization. Follows same schema as Labels for consistency.",
		//	  "patternProperties": {
		//	    "": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"tags":              // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The metadata, as key-value pairs, to apply to the node group to assist with categorization and organization. Follows same schema as Labels for consistency.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Taints
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Kubernetes taints to be applied to the nodes in the node group when they are created.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "An object representing a Taint specification for AWS EKS Nodegroup.",
		//	    "properties": {
		//	      "Effect": {
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Key": {
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"taints": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Effect
					"effect": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
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
			Description: "The Kubernetes taints to be applied to the nodes in the node group when they are created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: UpdateConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The node group update configuration.",
		//	  "properties": {
		//	    "MaxUnavailable": {
		//	      "description": "The maximum number of nodes unavailable at once during a version update. Nodes will be updated in parallel. This value or maxUnavailablePercentage is required to have a value.The maximum number is 100. ",
		//	      "minimum": 1,
		//	      "type": "number"
		//	    },
		//	    "MaxUnavailablePercentage": {
		//	      "description": "The maximum percentage of nodes unavailable during a version update. This percentage of nodes will be updated in parallel, up to 100 nodes at once. This value or maxUnavailable is required to have a value.",
		//	      "maximum": 100,
		//	      "minimum": 1,
		//	      "type": "number"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"update_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: MaxUnavailable
				"max_unavailable": schema.Float64Attribute{ /*START ATTRIBUTE*/
					Description: "The maximum number of nodes unavailable at once during a version update. Nodes will be updated in parallel. This value or maxUnavailablePercentage is required to have a value.The maximum number is 100. ",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: MaxUnavailablePercentage
				"max_unavailable_percentage": schema.Float64Attribute{ /*START ATTRIBUTE*/
					Description: "The maximum percentage of nodes unavailable during a version update. This percentage of nodes will be updated in parallel, up to 100 nodes at once. This value or maxUnavailable is required to have a value.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The node group update configuration.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Version
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Kubernetes version to use for your managed nodes.",
		//	  "type": "string"
		//	}
		"version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Kubernetes version to use for your managed nodes.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::EKS::Nodegroup",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EKS::Nodegroup").WithTerraformTypeName("awscc_eks_nodegroup")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"ami_type":                   "AmiType",
		"arn":                        "Arn",
		"capacity_type":              "CapacityType",
		"cluster_name":               "ClusterName",
		"desired_size":               "DesiredSize",
		"disk_size":                  "DiskSize",
		"ec_2_ssh_key":               "Ec2SshKey",
		"effect":                     "Effect",
		"force_update_enabled":       "ForceUpdateEnabled",
		"id":                         "Id",
		"instance_types":             "InstanceTypes",
		"key":                        "Key",
		"labels":                     "Labels",
		"launch_template":            "LaunchTemplate",
		"max_size":                   "MaxSize",
		"max_unavailable":            "MaxUnavailable",
		"max_unavailable_percentage": "MaxUnavailablePercentage",
		"min_size":                   "MinSize",
		"name":                       "Name",
		"node_role":                  "NodeRole",
		"nodegroup_name":             "NodegroupName",
		"release_version":            "ReleaseVersion",
		"remote_access":              "RemoteAccess",
		"scaling_config":             "ScalingConfig",
		"source_security_groups":     "SourceSecurityGroups",
		"subnets":                    "Subnets",
		"tags":                       "Tags",
		"taints":                     "Taints",
		"update_config":              "UpdateConfig",
		"value":                      "Value",
		"version":                    "Version",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
