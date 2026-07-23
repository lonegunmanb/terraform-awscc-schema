package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockagentcoreRuntime = `{
  "block": {
    "attributes": {
      "agent_runtime_arn": {
        "computed": true,
        "description": "The Amazon Resource Name(ARN) that uniquely identifies the Agent",
        "description_kind": "plain",
        "type": "string"
      },
      "agent_runtime_artifact": {
        "computed": true,
        "description": "The artifact of the agent",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "code_configuration": {
              "computed": true,
              "description": "Representation of a code configuration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "code": {
                    "computed": true,
                    "description": "Object represents source code from zip file",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "s3": {
                          "computed": true,
                          "description": "S3 Location Configuration",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "bucket": {
                                "computed": true,
                                "description": "S3 bucket name",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "prefix": {
                                "computed": true,
                                "description": "S3 object key prefix",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "version_id": {
                                "computed": true,
                                "description": "S3 object version ID",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "entry_point": {
                    "computed": true,
                    "description": "List of entry points",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "runtime": {
                    "computed": true,
                    "description": "Managed runtime types",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "container_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "container_uri": {
                    "computed": true,
                    "description": "The ECR URI of the container",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "agent_runtime_id": {
        "computed": true,
        "description": "Identifier for a resource",
        "description_kind": "plain",
        "type": "string"
      },
      "agent_runtime_name": {
        "computed": true,
        "description": "Name for a resource",
        "description_kind": "plain",
        "type": "string"
      },
      "agent_runtime_version": {
        "computed": true,
        "description": "Version of the Agent",
        "description_kind": "plain",
        "type": "string"
      },
      "authorizer_configuration": {
        "computed": true,
        "description": "Authorizer configuration for the agent runtime",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "custom_jwt_authorizer": {
              "computed": true,
              "description": "Configuration for custom JWT authorizer",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "allowed_audience": {
                    "computed": true,
                    "description": "List of allowed audiences",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "allowed_clients": {
                    "computed": true,
                    "description": "List of allowed clients",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "allowed_scopes": {
                    "computed": true,
                    "description": "List of allowed scopes",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "allowed_workload_configuration": {
                    "computed": true,
                    "description": "Allow-list of upstream workloads permitted to reach this resource via the workload identity chain. When set, the data plane enforces that the introspected workload chain's caller matches one of the configured hosting environments or workload identities; absent means no chain enforcement.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "hosting_environments": {
                          "computed": true,
                          "description": "List of allow-listed hosting environments",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "arn": {
                                "computed": true,
                                "description": "The ARN of the bedrock-agentcore hosting environment",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "set"
                          }
                        },
                        "workload_identities": {
                          "computed": true,
                          "description": "List of allow-listed workload identity names",
                          "description_kind": "plain",
                          "type": [
                            "set",
                            "string"
                          ]
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "custom_claims": {
                    "computed": true,
                    "description": "List of required custom claims",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "authorizing_claim_match_value": {
                          "computed": true,
                          "description": "The value or values in the custom claim to match and relationship of match",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "claim_match_operator": {
                                "computed": true,
                                "description": "The relationship between the claim field value and the value or values being matched",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "claim_match_value": {
                                "computed": true,
                                "description": "The value or values in the custom claim to match for",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "match_value_string": {
                                      "computed": true,
                                      "description": "The string value to match for",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "match_value_string_list": {
                                      "computed": true,
                                      "description": "The list of strings to check for a match",
                                      "description_kind": "plain",
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "inbound_token_claim_name": {
                          "computed": true,
                          "description": "The name of the custom claim to validate",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "inbound_token_claim_value_type": {
                          "computed": true,
                          "description": "Token claim data type",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "discovery_url": {
                    "computed": true,
                    "description": "OpenID Connect discovery URL",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "private_endpoint": {
                    "computed": true,
                    "description": "Private endpoint configuration. Exactly one of SelfManagedLatticeResource or ManagedVpcResource must be specified.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "managed_vpc_resource": {
                          "computed": true,
                          "description": "Managed VPC resource configuration",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "endpoint_ip_address_type": {
                                "computed": true,
                                "description": "The IP address type for the endpoint",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "routing_domain": {
                                "computed": true,
                                "description": "An intermediate domain to use as the resource configuration endpoint instead of the actual target domain",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "security_group_ids": {
                                "computed": true,
                                "description": "The security group IDs",
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "subnet_ids": {
                                "computed": true,
                                "description": "The subnet IDs",
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "tags": {
                                "computed": true,
                                "description": "Tags to apply to the managed VPC Lattice resource gateway",
                                "description_kind": "plain",
                                "type": [
                                  "map",
                                  "string"
                                ]
                              },
                              "vpc_identifier": {
                                "computed": true,
                                "description": "The VPC identifier",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "self_managed_lattice_resource": {
                          "computed": true,
                          "description": "Self-managed VPC Lattice resource configuration",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "resource_configuration_identifier": {
                                "computed": true,
                                "description": "The identifier of the VPC Lattice resource configuration",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "private_endpoint_overrides": {
                    "computed": true,
                    "description": "List of private endpoint overrides",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "domain": {
                          "computed": true,
                          "description": "The domain to override",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "private_endpoint": {
                          "computed": true,
                          "description": "Private endpoint configuration. Exactly one of SelfManagedLatticeResource or ManagedVpcResource must be specified.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "managed_vpc_resource": {
                                "computed": true,
                                "description": "Managed VPC resource configuration",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "endpoint_ip_address_type": {
                                      "computed": true,
                                      "description": "The IP address type for the endpoint",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "routing_domain": {
                                      "computed": true,
                                      "description": "An intermediate domain to use as the resource configuration endpoint instead of the actual target domain",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "security_group_ids": {
                                      "computed": true,
                                      "description": "The security group IDs",
                                      "description_kind": "plain",
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    },
                                    "subnet_ids": {
                                      "computed": true,
                                      "description": "The subnet IDs",
                                      "description_kind": "plain",
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    },
                                    "tags": {
                                      "computed": true,
                                      "description": "Tags to apply to the managed VPC Lattice resource gateway",
                                      "description_kind": "plain",
                                      "type": [
                                        "map",
                                        "string"
                                      ]
                                    },
                                    "vpc_identifier": {
                                      "computed": true,
                                      "description": "The VPC identifier",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "self_managed_lattice_resource": {
                                "computed": true,
                                "description": "Self-managed VPC Lattice resource configuration",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "resource_configuration_identifier": {
                                      "computed": true,
                                      "description": "The identifier of the VPC Lattice resource configuration",
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              }
                            },
                            "nesting_mode": "single"
                          }
                        }
                      },
                      "nesting_mode": "list"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "created_at": {
        "computed": true,
        "description": "Timestamp when the Agent was created",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Description of the resource",
        "description_kind": "plain",
        "type": "string"
      },
      "environment_variables": {
        "computed": true,
        "description": "Environment variables for the agent runtime",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "failure_reason": {
        "computed": true,
        "description": "The reason for failure if the agent is in a failed state.",
        "description_kind": "plain",
        "type": "string"
      },
      "filesystem_configurations": {
        "computed": true,
        "description": "Filesystem configurations for the agent runtime",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "efs_access_point": {
              "computed": true,
              "description": "Configuration for EFS access point filesystem",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "access_point_arn": {
                    "computed": true,
                    "description": "ARN of the EFS access point",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "mount_path": {
                    "computed": true,
                    "description": "Mount path for filesystem configuration",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "s3_files_access_point": {
              "computed": true,
              "description": "Configuration for S3 Files access point filesystem",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "access_point_arn": {
                    "computed": true,
                    "description": "ARN of the S3 Files access point",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "mount_path": {
                    "computed": true,
                    "description": "Mount path for filesystem configuration",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "session_storage": {
              "computed": true,
              "description": "Configuration for session storage",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "mount_path": {
                    "computed": true,
                    "description": "Mount path for filesystem configuration",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "list"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_updated_at": {
        "computed": true,
        "description": "When resource was last updated",
        "description_kind": "plain",
        "type": "string"
      },
      "lifecycle_configuration": {
        "computed": true,
        "description": "Lifecycle configuration for managing runtime sessions",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "idle_runtime_session_timeout": {
              "computed": true,
              "description": "Timeout in seconds for idle runtime sessions",
              "description_kind": "plain",
              "type": "number"
            },
            "max_lifetime": {
              "computed": true,
              "description": "Maximum lifetime in seconds for runtime sessions",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "network_configuration": {
        "computed": true,
        "description": "Network access configuration for the Agent",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "network_mode": {
              "computed": true,
              "description": "Network mode configuration type",
              "description_kind": "plain",
              "type": "string"
            },
            "network_mode_config": {
              "computed": true,
              "description": "Network mode configuration for VPC",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "security_groups": {
                    "computed": true,
                    "description": "Security groups for VPC",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "subnets": {
                    "computed": true,
                    "description": "Subnets for VPC",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "protocol_configuration": {
        "computed": true,
        "description": "Protocol configuration for the agent runtime",
        "description_kind": "plain",
        "type": "string"
      },
      "request_header_configuration": {
        "computed": true,
        "description": "Configuration for HTTP request headers",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "request_header_allowlist": {
              "computed": true,
              "description": "List of allowed HTTP headers for agent runtime requests",
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        }
      },
      "role_arn": {
        "computed": true,
        "description": "Amazon Resource Name (ARN) of an IAM role",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "Current status of the agent",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A map of tag keys and values",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "workload_identity_details": {
        "computed": true,
        "description": "Workload identity details for the agent",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "workload_identity_arn": {
              "computed": true,
              "description": "ARN of the workload identity",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::BedrockAgentCore::Runtime",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBedrockagentcoreRuntimeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockagentcoreRuntime), &result)
	return &result
}
