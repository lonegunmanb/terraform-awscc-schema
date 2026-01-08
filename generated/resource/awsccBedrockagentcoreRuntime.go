package resource

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
                                "optional": true,
                                "type": "string"
                              },
                              "prefix": {
                                "computed": true,
                                "description": "S3 object key prefix",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "version_id": {
                                "computed": true,
                                "description": "S3 object version ID",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "entry_point": {
                    "computed": true,
                    "description": "List of entry points",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "runtime": {
                    "computed": true,
                    "description": "Managed runtime types",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
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
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "agent_runtime_id": {
        "computed": true,
        "description": "Identifier for a resource",
        "description_kind": "plain",
        "type": "string"
      },
      "agent_runtime_name": {
        "description": "Name for a resource",
        "description_kind": "plain",
        "required": true,
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
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "allowed_clients": {
                    "computed": true,
                    "description": "List of allowed clients",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "discovery_url": {
                    "computed": true,
                    "description": "OpenID Connect discovery URL",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
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
        "optional": true,
        "type": "string"
      },
      "environment_variables": {
        "computed": true,
        "description": "Environment variables for the agent runtime",
        "description_kind": "plain",
        "optional": true,
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
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
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
              "optional": true,
              "type": "number"
            },
            "max_lifetime": {
              "computed": true,
              "description": "Maximum lifetime in seconds for runtime sessions",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "network_configuration": {
        "description": "Network access configuration for the Agent",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "network_mode": {
              "description": "Network mode configuration type",
              "description_kind": "plain",
              "required": true,
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
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "subnets": {
                    "computed": true,
                    "description": "Subnets for VPC",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "protocol_configuration": {
        "computed": true,
        "description": "Protocol configuration for the agent runtime",
        "description_kind": "plain",
        "optional": true,
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
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "role_arn": {
        "description": "Amazon Resource Name (ARN) of an IAM role",
        "description_kind": "plain",
        "required": true,
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
        "optional": true,
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
    "description": "Resource Type definition for AWS::BedrockAgentCore::Runtime",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBedrockagentcoreRuntimeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockagentcoreRuntime), &result)
	return &result
}
