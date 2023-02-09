// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ec2_network_insights_access_scope", networkInsightsAccessScopeDataSource)
}

// networkInsightsAccessScopeDataSource returns the Terraform awscc_ec2_network_insights_access_scope data source.
// This Terraform data source corresponds to the CloudFormation AWS::EC2::NetworkInsightsAccessScope resource.
func networkInsightsAccessScopeDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CreatedDate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"created_date": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ExcludePaths
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Destination": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "PacketHeaderStatement": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "DestinationAddresses": {
		//	                "insertionOrder": true,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              },
		//	              "DestinationPorts": {
		//	                "insertionOrder": true,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              },
		//	              "DestinationPrefixLists": {
		//	                "insertionOrder": true,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              },
		//	              "Protocols": {
		//	                "insertionOrder": true,
		//	                "items": {
		//	                  "enum": [
		//	                    "tcp",
		//	                    "udp"
		//	                  ],
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              },
		//	              "SourceAddresses": {
		//	                "insertionOrder": true,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              },
		//	              "SourcePorts": {
		//	                "insertionOrder": true,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              },
		//	              "SourcePrefixLists": {
		//	                "insertionOrder": true,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              }
		//	            },
		//	            "type": "object"
		//	          },
		//	          "ResourceStatement": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "ResourceTypes": {
		//	                "insertionOrder": true,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              },
		//	              "Resources": {
		//	                "insertionOrder": true,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              }
		//	            },
		//	            "type": "object"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "Source": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "PacketHeaderStatement": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "DestinationAddresses": {
		//	                "insertionOrder": true,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              },
		//	              "DestinationPorts": {
		//	                "insertionOrder": true,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              },
		//	              "DestinationPrefixLists": {
		//	                "insertionOrder": true,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              },
		//	              "Protocols": {
		//	                "insertionOrder": true,
		//	                "items": {
		//	                  "enum": [
		//	                    "tcp",
		//	                    "udp"
		//	                  ],
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              },
		//	              "SourceAddresses": {
		//	                "insertionOrder": true,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              },
		//	              "SourcePorts": {
		//	                "insertionOrder": true,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              },
		//	              "SourcePrefixLists": {
		//	                "insertionOrder": true,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              }
		//	            },
		//	            "type": "object"
		//	          },
		//	          "ResourceStatement": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "ResourceTypes": {
		//	                "insertionOrder": true,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              },
		//	              "Resources": {
		//	                "insertionOrder": true,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              }
		//	            },
		//	            "type": "object"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "ThroughResources": {
		//	        "insertionOrder": true,
		//	        "items": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "ResourceStatement": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "ResourceTypes": {
		//	                  "insertionOrder": true,
		//	                  "items": {
		//	                    "type": "string"
		//	                  },
		//	                  "type": "array"
		//	                },
		//	                "Resources": {
		//	                  "insertionOrder": true,
		//	                  "items": {
		//	                    "type": "string"
		//	                  },
		//	                  "type": "array"
		//	                }
		//	              },
		//	              "type": "object"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "type": "array"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"exclude_paths": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Destination
					"destination": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: PacketHeaderStatement
							"packet_header_statement": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: DestinationAddresses
									"destination_addresses": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: DestinationPorts
									"destination_ports": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: DestinationPrefixLists
									"destination_prefix_lists": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: Protocols
									"protocols": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: SourceAddresses
									"source_addresses": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: SourcePorts
									"source_ports": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: SourcePrefixLists
									"source_prefix_lists": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Computed:    true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: ResourceStatement
							"resource_statement": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: ResourceTypes
									"resource_types": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: Resources
									"resources": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Computed:    true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Computed: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Source
					"source": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: PacketHeaderStatement
							"packet_header_statement": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: DestinationAddresses
									"destination_addresses": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: DestinationPorts
									"destination_ports": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: DestinationPrefixLists
									"destination_prefix_lists": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: Protocols
									"protocols": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: SourceAddresses
									"source_addresses": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: SourcePorts
									"source_ports": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: SourcePrefixLists
									"source_prefix_lists": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Computed:    true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: ResourceStatement
							"resource_statement": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: ResourceTypes
									"resource_types": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: Resources
									"resources": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Computed:    true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Computed: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: ThroughResources
					"through_resources": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: ResourceStatement
								"resource_statement": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: ResourceTypes
										"resource_types": schema.ListAttribute{ /*START ATTRIBUTE*/
											ElementType: types.StringType,
											Computed:    true,
										}, /*END ATTRIBUTE*/
										// Property: Resources
										"resources": schema.ListAttribute{ /*START ATTRIBUTE*/
											ElementType: types.StringType,
											Computed:    true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Computed: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: MatchPaths
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Destination": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "PacketHeaderStatement": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "DestinationAddresses": {
		//	                "insertionOrder": true,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              },
		//	              "DestinationPorts": {
		//	                "insertionOrder": true,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              },
		//	              "DestinationPrefixLists": {
		//	                "insertionOrder": true,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              },
		//	              "Protocols": {
		//	                "insertionOrder": true,
		//	                "items": {
		//	                  "enum": [
		//	                    "tcp",
		//	                    "udp"
		//	                  ],
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              },
		//	              "SourceAddresses": {
		//	                "insertionOrder": true,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              },
		//	              "SourcePorts": {
		//	                "insertionOrder": true,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              },
		//	              "SourcePrefixLists": {
		//	                "insertionOrder": true,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              }
		//	            },
		//	            "type": "object"
		//	          },
		//	          "ResourceStatement": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "ResourceTypes": {
		//	                "insertionOrder": true,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              },
		//	              "Resources": {
		//	                "insertionOrder": true,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              }
		//	            },
		//	            "type": "object"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "Source": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "PacketHeaderStatement": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "DestinationAddresses": {
		//	                "insertionOrder": true,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              },
		//	              "DestinationPorts": {
		//	                "insertionOrder": true,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              },
		//	              "DestinationPrefixLists": {
		//	                "insertionOrder": true,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              },
		//	              "Protocols": {
		//	                "insertionOrder": true,
		//	                "items": {
		//	                  "enum": [
		//	                    "tcp",
		//	                    "udp"
		//	                  ],
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              },
		//	              "SourceAddresses": {
		//	                "insertionOrder": true,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              },
		//	              "SourcePorts": {
		//	                "insertionOrder": true,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              },
		//	              "SourcePrefixLists": {
		//	                "insertionOrder": true,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              }
		//	            },
		//	            "type": "object"
		//	          },
		//	          "ResourceStatement": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "ResourceTypes": {
		//	                "insertionOrder": true,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              },
		//	              "Resources": {
		//	                "insertionOrder": true,
		//	                "items": {
		//	                  "type": "string"
		//	                },
		//	                "type": "array"
		//	              }
		//	            },
		//	            "type": "object"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "ThroughResources": {
		//	        "insertionOrder": true,
		//	        "items": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "ResourceStatement": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "ResourceTypes": {
		//	                  "insertionOrder": true,
		//	                  "items": {
		//	                    "type": "string"
		//	                  },
		//	                  "type": "array"
		//	                },
		//	                "Resources": {
		//	                  "insertionOrder": true,
		//	                  "items": {
		//	                    "type": "string"
		//	                  },
		//	                  "type": "array"
		//	                }
		//	              },
		//	              "type": "object"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "type": "array"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"match_paths": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Destination
					"destination": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: PacketHeaderStatement
							"packet_header_statement": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: DestinationAddresses
									"destination_addresses": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: DestinationPorts
									"destination_ports": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: DestinationPrefixLists
									"destination_prefix_lists": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: Protocols
									"protocols": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: SourceAddresses
									"source_addresses": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: SourcePorts
									"source_ports": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: SourcePrefixLists
									"source_prefix_lists": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Computed:    true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: ResourceStatement
							"resource_statement": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: ResourceTypes
									"resource_types": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: Resources
									"resources": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Computed:    true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Computed: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Source
					"source": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: PacketHeaderStatement
							"packet_header_statement": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: DestinationAddresses
									"destination_addresses": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: DestinationPorts
									"destination_ports": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: DestinationPrefixLists
									"destination_prefix_lists": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: Protocols
									"protocols": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: SourceAddresses
									"source_addresses": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: SourcePorts
									"source_ports": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: SourcePrefixLists
									"source_prefix_lists": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Computed:    true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: ResourceStatement
							"resource_statement": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: ResourceTypes
									"resource_types": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: Resources
									"resources": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Computed:    true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Computed: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: ThroughResources
					"through_resources": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: ResourceStatement
								"resource_statement": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: ResourceTypes
										"resource_types": schema.ListAttribute{ /*START ATTRIBUTE*/
											ElementType: types.StringType,
											Computed:    true,
										}, /*END ATTRIBUTE*/
										// Property: Resources
										"resources": schema.ListAttribute{ /*START ATTRIBUTE*/
											ElementType: types.StringType,
											Computed:    true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Computed: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: NetworkInsightsAccessScopeArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"network_insights_access_scope_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: NetworkInsightsAccessScopeId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"network_insights_access_scope_id": schema.StringAttribute{ /*START ATTRIBUTE*/
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
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
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
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: UpdatedDate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"updated_date": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::EC2::NetworkInsightsAccessScope",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::NetworkInsightsAccessScope").WithTerraformTypeName("awscc_ec2_network_insights_access_scope")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"created_date":                      "CreatedDate",
		"destination":                       "Destination",
		"destination_addresses":             "DestinationAddresses",
		"destination_ports":                 "DestinationPorts",
		"destination_prefix_lists":          "DestinationPrefixLists",
		"exclude_paths":                     "ExcludePaths",
		"key":                               "Key",
		"match_paths":                       "MatchPaths",
		"network_insights_access_scope_arn": "NetworkInsightsAccessScopeArn",
		"network_insights_access_scope_id":  "NetworkInsightsAccessScopeId",
		"packet_header_statement":           "PacketHeaderStatement",
		"protocols":                         "Protocols",
		"resource_statement":                "ResourceStatement",
		"resource_types":                    "ResourceTypes",
		"resources":                         "Resources",
		"source":                            "Source",
		"source_addresses":                  "SourceAddresses",
		"source_ports":                      "SourcePorts",
		"source_prefix_lists":               "SourcePrefixLists",
		"tags":                              "Tags",
		"through_resources":                 "ThroughResources",
		"updated_date":                      "UpdatedDate",
		"value":                             "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
