// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package route53recoveryreadiness

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_route53recoveryreadiness_resource_set", resourceSetDataSource)
}

// resourceSetDataSource returns the Terraform awscc_route53recoveryreadiness_resource_set data source.
// This Terraform data source corresponds to the CloudFormation AWS::Route53RecoveryReadiness::ResourceSet resource.
func resourceSetDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ResourceSetArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the resource set.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"resource_set_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the resource set.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResourceSetName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the resource set to create.",
		//	  "type": "string"
		//	}
		"resource_set_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the resource set to create.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResourceSetType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The resource type of the resources in the resource set. Enter one of the following values for resource type: \n\nAWS: :AutoScaling: :AutoScalingGroup, AWS: :CloudWatch: :Alarm, AWS: :EC2: :CustomerGateway, AWS: :DynamoDB: :Table, AWS: :EC2: :Volume, AWS: :ElasticLoadBalancing: :LoadBalancer, AWS: :ElasticLoadBalancingV2: :LoadBalancer, AWS: :MSK: :Cluster, AWS: :RDS: :DBCluster, AWS: :Route53: :HealthCheck, AWS: :SQS: :Queue, AWS: :SNS: :Topic, AWS: :SNS: :Subscription, AWS: :EC2: :VPC, AWS: :EC2: :VPNConnection, AWS: :EC2: :VPNGateway, AWS::Route53RecoveryReadiness::DNSTargetResource",
		//	  "type": "string"
		//	}
		"resource_set_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The resource type of the resources in the resource set. Enter one of the following values for resource type: \n\nAWS: :AutoScaling: :AutoScalingGroup, AWS: :CloudWatch: :Alarm, AWS: :EC2: :CustomerGateway, AWS: :DynamoDB: :Table, AWS: :EC2: :Volume, AWS: :ElasticLoadBalancing: :LoadBalancer, AWS: :ElasticLoadBalancingV2: :LoadBalancer, AWS: :MSK: :Cluster, AWS: :RDS: :DBCluster, AWS: :Route53: :HealthCheck, AWS: :SQS: :Queue, AWS: :SNS: :Topic, AWS: :SNS: :Subscription, AWS: :EC2: :VPC, AWS: :EC2: :VPNConnection, AWS: :EC2: :VPNGateway, AWS::Route53RecoveryReadiness::DNSTargetResource",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Resources
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of resource objects in the resource set.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "The resource element of a ResourceSet",
		//	    "properties": {
		//	      "ComponentId": {
		//	        "description": "The component identifier of the resource, generated when DNS target resource is used.",
		//	        "type": "string"
		//	      },
		//	      "DnsTargetResource": {
		//	        "additionalProperties": false,
		//	        "description": "A component for DNS/routing control readiness checks.",
		//	        "properties": {
		//	          "DomainName": {
		//	            "description": "The domain name that acts as an ingress point to a portion of the customer application.",
		//	            "type": "string"
		//	          },
		//	          "HostedZoneArn": {
		//	            "description": "The hosted zone Amazon Resource Name (ARN) that contains the DNS record with the provided name of the target resource.",
		//	            "type": "string"
		//	          },
		//	          "RecordSetId": {
		//	            "description": "The Route 53 record set ID that will uniquely identify a DNS record, given a name and a type.",
		//	            "type": "string"
		//	          },
		//	          "RecordType": {
		//	            "description": "The type of DNS record of the target resource.",
		//	            "type": "string"
		//	          },
		//	          "TargetResource": {
		//	            "additionalProperties": false,
		//	            "description": "The target resource that the Route 53 record points to.",
		//	            "oneOf": [
		//	              {
		//	                "required": [
		//	                  "NLBResource"
		//	                ]
		//	              },
		//	              {
		//	                "required": [
		//	                  "R53Resource"
		//	                ]
		//	              }
		//	            ],
		//	            "properties": {
		//	              "NLBResource": {
		//	                "additionalProperties": false,
		//	                "description": "The Network Load Balancer resource that a DNS target resource points to.",
		//	                "properties": {
		//	                  "Arn": {
		//	                    "description": "A Network Load Balancer resource Amazon Resource Name (ARN).",
		//	                    "type": "string"
		//	                  }
		//	                },
		//	                "type": "object"
		//	              },
		//	              "R53Resource": {
		//	                "additionalProperties": false,
		//	                "description": "The Route 53 resource that a DNS target resource record points to.",
		//	                "properties": {
		//	                  "DomainName": {
		//	                    "description": "The DNS target domain name.",
		//	                    "type": "string"
		//	                  },
		//	                  "RecordSetId": {
		//	                    "description": "The Resource Record set id.",
		//	                    "type": "string"
		//	                  }
		//	                },
		//	                "type": "object"
		//	              }
		//	            },
		//	            "type": "object"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "ReadinessScopes": {
		//	        "description": "A list of recovery group Amazon Resource Names (ARNs) and cell ARNs that this resource is contained within.",
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "maxItems": 5,
		//	          "type": "string"
		//	        },
		//	        "type": "array"
		//	      },
		//	      "ResourceArn": {
		//	        "description": "The Amazon Resource Name (ARN) of the AWS resource.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "maxItems": 6,
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"resources": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: ComponentId
					"component_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The component identifier of the resource, generated when DNS target resource is used.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: DnsTargetResource
					"dns_target_resource": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: DomainName
							"domain_name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The domain name that acts as an ingress point to a portion of the customer application.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: HostedZoneArn
							"hosted_zone_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The hosted zone Amazon Resource Name (ARN) that contains the DNS record with the provided name of the target resource.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: RecordSetId
							"record_set_id": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The Route 53 record set ID that will uniquely identify a DNS record, given a name and a type.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: RecordType
							"record_type": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The type of DNS record of the target resource.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: TargetResource
							"target_resource": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: NLBResource
									"nlb_resource": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
										Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
											// Property: Arn
											"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
												Description: "A Network Load Balancer resource Amazon Resource Name (ARN).",
												Computed:    true,
											}, /*END ATTRIBUTE*/
										}, /*END SCHEMA*/
										Description: "The Network Load Balancer resource that a DNS target resource points to.",
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: R53Resource
									"r53_resource": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
										Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
											// Property: DomainName
											"domain_name": schema.StringAttribute{ /*START ATTRIBUTE*/
												Description: "The DNS target domain name.",
												Computed:    true,
											}, /*END ATTRIBUTE*/
											// Property: RecordSetId
											"record_set_id": schema.StringAttribute{ /*START ATTRIBUTE*/
												Description: "The Resource Record set id.",
												Computed:    true,
											}, /*END ATTRIBUTE*/
										}, /*END SCHEMA*/
										Description: "The Route 53 resource that a DNS target resource record points to.",
										Computed:    true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Description: "The target resource that the Route 53 record points to.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "A component for DNS/routing control readiness checks.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: ReadinessScopes
					"readiness_scopes": schema.ListAttribute{ /*START ATTRIBUTE*/
						ElementType: types.StringType,
						Description: "A list of recovery group Amazon Resource Names (ARNs) and cell ARNs that this resource is contained within.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: ResourceArn
					"resource_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The Amazon Resource Name (ARN) of the AWS resource.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of resource objects in the resource set.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A tag to associate with the parameters for a resource set.",
		//	  "insertionOrder": false,
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
		//	  "type": "array"
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
			Description: "A tag to associate with the parameters for a resource set.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Route53RecoveryReadiness::ResourceSet",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53RecoveryReadiness::ResourceSet").WithTerraformTypeName("awscc_route53recoveryreadiness_resource_set")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                 "Arn",
		"component_id":        "ComponentId",
		"dns_target_resource": "DnsTargetResource",
		"domain_name":         "DomainName",
		"hosted_zone_arn":     "HostedZoneArn",
		"key":                 "Key",
		"nlb_resource":        "NLBResource",
		"r53_resource":        "R53Resource",
		"readiness_scopes":    "ReadinessScopes",
		"record_set_id":       "RecordSetId",
		"record_type":         "RecordType",
		"resource_arn":        "ResourceArn",
		"resource_set_arn":    "ResourceSetArn",
		"resource_set_name":   "ResourceSetName",
		"resource_set_type":   "ResourceSetType",
		"resources":           "Resources",
		"tags":                "Tags",
		"target_resource":     "TargetResource",
		"value":               "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
