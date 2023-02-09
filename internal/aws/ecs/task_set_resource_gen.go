// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package ecs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_ecs_task_set", taskSetResource)
}

// taskSetResource returns the Terraform awscc_ecs_task_set resource.
// This Terraform resource corresponds to the CloudFormation AWS::ECS::TaskSet resource.
func taskSetResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Cluster
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The short name or full Amazon Resource Name (ARN) of the cluster that hosts the service to create the task set in.",
		//	  "type": "string"
		//	}
		"cluster": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The short name or full Amazon Resource Name (ARN) of the cluster that hosts the service to create the task set in.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ExternalId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An optional non-unique tag that identifies this task set in external systems. If the task set is associated with a service discovery registry, the tasks in this task set will have the ECS_TASK_SET_EXTERNAL_ID AWS Cloud Map attribute set to the provided value. ",
		//	  "type": "string"
		//	}
		"external_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An optional non-unique tag that identifies this task set in external systems. If the task set is associated with a service discovery registry, the tasks in this task set will have the ECS_TASK_SET_EXTERNAL_ID AWS Cloud Map attribute set to the provided value. ",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the task set.",
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the task set.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LaunchType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The launch type that new tasks in the task set will use. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html in the Amazon Elastic Container Service Developer Guide. ",
		//	  "enum": [
		//	    "EC2",
		//	    "FARGATE"
		//	  ],
		//	  "type": "string"
		//	}
		"launch_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The launch type that new tasks in the task set will use. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html in the Amazon Elastic Container Service Developer Guide. ",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"EC2",
					"FARGATE",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LoadBalancers
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A load balancer object representing the load balancer to use with the task set. The supported load balancer types are either an Application Load Balancer or a Network Load Balancer. ",
		//	    "properties": {
		//	      "ContainerName": {
		//	        "description": "The name of the container (as it appears in a container definition) to associate with the load balancer.",
		//	        "type": "string"
		//	      },
		//	      "ContainerPort": {
		//	        "description": "The port on the container to associate with the load balancer. This port must correspond to a containerPort in the task definition the tasks in the service are using. For tasks that use the EC2 launch type, the container instance they are launched on must allow ingress traffic on the hostPort of the port mapping.",
		//	        "type": "integer"
		//	      },
		//	      "LoadBalancerName": {
		//	        "description": "The name of the load balancer to associate with the Amazon ECS service or task set. A load balancer name is only specified when using a Classic Load Balancer. If you are using an Application Load Balancer or a Network Load Balancer this should be omitted.",
		//	        "type": "string"
		//	      },
		//	      "TargetGroupArn": {
		//	        "description": "The full Amazon Resource Name (ARN) of the Elastic Load Balancing target group or groups associated with a service or task set. A target group ARN is only specified when using an Application Load Balancer or Network Load Balancer. If you are using a Classic Load Balancer this should be omitted. For services using the ECS deployment controller, you can specify one or multiple target groups. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/register-multiple-targetgroups.html in the Amazon Elastic Container Service Developer Guide. For services using the CODE_DEPLOY deployment controller, you are required to define two target groups for the load balancer. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-type-bluegreen.html in the Amazon Elastic Container Service Developer Guide. If your service's task definition uses the awsvpc network mode (which is required for the Fargate launch type), you must choose ip as the target type, not instance, when creating your target groups because tasks that use the awsvpc network mode are associated with an elastic network interface, not an Amazon EC2 instance.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"load_balancers": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: ContainerName
					"container_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The name of the container (as it appears in a container definition) to associate with the load balancer.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: ContainerPort
					"container_port": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Description: "The port on the container to associate with the load balancer. This port must correspond to a containerPort in the task definition the tasks in the service are using. For tasks that use the EC2 launch type, the container instance they are launched on must allow ingress traffic on the hostPort of the port mapping.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
							int64planmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: LoadBalancerName
					"load_balancer_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The name of the load balancer to associate with the Amazon ECS service or task set. A load balancer name is only specified when using a Classic Load Balancer. If you are using an Application Load Balancer or a Network Load Balancer this should be omitted.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: TargetGroupArn
					"target_group_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The full Amazon Resource Name (ARN) of the Elastic Load Balancing target group or groups associated with a service or task set. A target group ARN is only specified when using an Application Load Balancer or Network Load Balancer. If you are using a Classic Load Balancer this should be omitted. For services using the ECS deployment controller, you can specify one or multiple target groups. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/register-multiple-targetgroups.html in the Amazon Elastic Container Service Developer Guide. For services using the CODE_DEPLOY deployment controller, you are required to define two target groups for the load balancer. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-type-bluegreen.html in the Amazon Elastic Container Service Developer Guide. If your service's task definition uses the awsvpc network mode (which is required for the Fargate launch type), you must choose ip as the target type, not instance, when creating your target groups because tasks that use the awsvpc network mode are associated with an elastic network interface, not an Amazon EC2 instance.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
				listplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: NetworkConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "An object representing the network configuration for a task or service.",
		//	  "properties": {
		//	    "AwsVpcConfiguration": {
		//	      "additionalProperties": false,
		//	      "description": "The VPC subnets and security groups associated with a task. All specified subnets and security groups must be from the same VPC.",
		//	      "properties": {
		//	        "AssignPublicIp": {
		//	          "description": "Whether the task's elastic network interface receives a public IP address. The default value is DISABLED.",
		//	          "enum": [
		//	            "DISABLED",
		//	            "ENABLED"
		//	          ],
		//	          "type": "string"
		//	        },
		//	        "SecurityGroups": {
		//	          "description": "The security groups associated with the task or service. If you do not specify a security group, the default security group for the VPC is used. There is a limit of 5 security groups that can be specified per AwsVpcConfiguration.",
		//	          "items": {
		//	            "type": "string"
		//	          },
		//	          "maxItems": 5,
		//	          "type": "array"
		//	        },
		//	        "Subnets": {
		//	          "description": "The subnets associated with the task or service. There is a limit of 16 subnets that can be specified per AwsVpcConfiguration.",
		//	          "items": {
		//	            "type": "string"
		//	          },
		//	          "maxItems": 16,
		//	          "type": "array"
		//	        }
		//	      },
		//	      "required": [
		//	        "Subnets"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"network_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AwsVpcConfiguration
				"aws_vpc_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: AssignPublicIp
						"assign_public_ip": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Whether the task's elastic network interface receives a public IP address. The default value is DISABLED.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.OneOf(
									"DISABLED",
									"ENABLED",
								),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: SecurityGroups
						"security_groups": schema.ListAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "The security groups associated with the task or service. If you do not specify a security group, the default security group for the VPC is used. There is a limit of 5 security groups that can be specified per AwsVpcConfiguration.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.List{ /*START VALIDATORS*/
								listvalidator.SizeAtMost(5),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
								listplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Subnets
						"subnets": schema.ListAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "The subnets associated with the task or service. There is a limit of 16 subnets that can be specified per AwsVpcConfiguration.",
							Required:    true,
							Validators: []validator.List{ /*START VALIDATORS*/
								listvalidator.SizeAtMost(16),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The VPC subnets and security groups associated with a task. All specified subnets and security groups must be from the same VPC.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "An object representing the network configuration for a task or service.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PlatformVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The platform version that the tasks in the task set should use. A platform version is specified only for tasks using the Fargate launch type. If one isn't specified, the LATEST platform version is used by default.",
		//	  "type": "string"
		//	}
		"platform_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The platform version that the tasks in the task set should use. A platform version is specified only for tasks using the Fargate launch type. If one isn't specified, the LATEST platform version is used by default.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Scale
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "A floating-point percentage of the desired number of tasks to place and keep running in the task set.",
		//	  "properties": {
		//	    "Unit": {
		//	      "description": "The unit of measure for the scale value.",
		//	      "enum": [
		//	        "PERCENT"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "Value": {
		//	      "description": "The value, specified as a percent total of a service's desiredCount, to scale the task set. Accepted values are numbers between 0 and 100.",
		//	      "maximum": 100,
		//	      "minimum": 0,
		//	      "type": "number"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"scale": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Unit
				"unit": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The unit of measure for the scale value.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"PERCENT",
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Value
				"value": schema.Float64Attribute{ /*START ATTRIBUTE*/
					Description: "The value, specified as a percent total of a service's desiredCount, to scale the task set. Accepted values are numbers between 0 and 100.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Float64{ /*START VALIDATORS*/
						float64validator.Between(0.000000, 100.000000),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
						float64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "A floating-point percentage of the desired number of tasks to place and keep running in the task set.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Service
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The short name or full Amazon Resource Name (ARN) of the service to create the task set in.",
		//	  "type": "string"
		//	}
		"service": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The short name or full Amazon Resource Name (ARN) of the service to create the task set in.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ServiceRegistries
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The details of the service discovery registries to assign to this task set. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-discovery.html.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "ContainerName": {
		//	        "description": "The container name value, already specified in the task definition, to be used for your service discovery service. If the task definition that your service task specifies uses the bridge or host network mode, you must specify a containerName and containerPort combination from the task definition. If the task definition that your service task specifies uses the awsvpc network mode and a type SRV DNS record is used, you must specify either a containerName and containerPort combination or a port value, but not both.",
		//	        "type": "string"
		//	      },
		//	      "ContainerPort": {
		//	        "description": "The port value, already specified in the task definition, to be used for your service discovery service. If the task definition your service task specifies uses the bridge or host network mode, you must specify a containerName and containerPort combination from the task definition. If the task definition your service task specifies uses the awsvpc network mode and a type SRV DNS record is used, you must specify either a containerName and containerPort combination or a port value, but not both.",
		//	        "type": "integer"
		//	      },
		//	      "Port": {
		//	        "description": "The port value used if your service discovery service specified an SRV record. This field may be used if both the awsvpc network mode and SRV records are used.",
		//	        "type": "integer"
		//	      },
		//	      "RegistryArn": {
		//	        "description": "The Amazon Resource Name (ARN) of the service registry. The currently supported service registry is AWS Cloud Map. For more information, see https://docs.aws.amazon.com/cloud-map/latest/api/API_CreateService.html",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"service_registries": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: ContainerName
					"container_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The container name value, already specified in the task definition, to be used for your service discovery service. If the task definition that your service task specifies uses the bridge or host network mode, you must specify a containerName and containerPort combination from the task definition. If the task definition that your service task specifies uses the awsvpc network mode and a type SRV DNS record is used, you must specify either a containerName and containerPort combination or a port value, but not both.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: ContainerPort
					"container_port": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Description: "The port value, already specified in the task definition, to be used for your service discovery service. If the task definition your service task specifies uses the bridge or host network mode, you must specify a containerName and containerPort combination from the task definition. If the task definition your service task specifies uses the awsvpc network mode and a type SRV DNS record is used, you must specify either a containerName and containerPort combination or a port value, but not both.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
							int64planmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Port
					"port": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Description: "The port value used if your service discovery service specified an SRV record. This field may be used if both the awsvpc network mode and SRV records are used.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
							int64planmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: RegistryArn
					"registry_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The Amazon Resource Name (ARN) of the service registry. The currently supported service registry is AWS Cloud Map. For more information, see https://docs.aws.amazon.com/cloud-map/latest/api/API_CreateService.html",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The details of the service discovery registries to assign to this task set. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-discovery.html.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
				listplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TaskDefinition
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The short name or full Amazon Resource Name (ARN) of the task definition for the tasks in the task set to use.",
		//	  "type": "string"
		//	}
		"task_definition": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The short name or full Amazon Resource Name (ARN) of the task definition for the tasks in the task set to use.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	schema := schema.Schema{
		Description: "Create a task set in the specified cluster and service. This is used when a service uses the EXTERNAL deployment controller type. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.htmlin the Amazon Elastic Container Service Developer Guide.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ECS::TaskSet").WithTerraformTypeName("awscc_ecs_task_set")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"assign_public_ip":      "AssignPublicIp",
		"aws_vpc_configuration": "AwsVpcConfiguration",
		"cluster":               "Cluster",
		"container_name":        "ContainerName",
		"container_port":        "ContainerPort",
		"external_id":           "ExternalId",
		"id":                    "Id",
		"launch_type":           "LaunchType",
		"load_balancer_name":    "LoadBalancerName",
		"load_balancers":        "LoadBalancers",
		"network_configuration": "NetworkConfiguration",
		"platform_version":      "PlatformVersion",
		"port":                  "Port",
		"registry_arn":          "RegistryArn",
		"scale":                 "Scale",
		"security_groups":       "SecurityGroups",
		"service":               "Service",
		"service_registries":    "ServiceRegistries",
		"subnets":               "Subnets",
		"target_group_arn":      "TargetGroupArn",
		"task_definition":       "TaskDefinition",
		"unit":                  "Unit",
		"value":                 "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
