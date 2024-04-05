package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockAgent = `{
  "block": {
    "attributes": {
      "action_groups": {
        "computed": true,
        "description": "List of ActionGroups",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "action_group_executor": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "lambda": {
                    "description": "ARN of a Lambda.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "action_group_name": {
              "description": "Name of the action group",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "action_group_state": {
              "computed": true,
              "description": "State of the action group",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "api_schema": {
              "computed": true,
              "description": "Contains information about the API Schema for the Action Group",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "payload": {
                    "computed": true,
                    "description": "String OpenAPI Payload",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "s3": {
                    "computed": true,
                    "description": "The identifier for the S3 resource.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "s3_bucket_name": {
                          "computed": true,
                          "description": "A bucket in S3.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "s3_object_key": {
                          "computed": true,
                          "description": "A object key in S3.",
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
            "description": {
              "computed": true,
              "description": "Description of action group",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "parent_action_group_signature": {
              "computed": true,
              "description": "Action Group Signature for a BuiltIn Action",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "skip_resource_in_use_check_on_delete": {
              "computed": true,
              "description": "Specifies whether to allow deleting action group while it is in use.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "agent_arn": {
        "computed": true,
        "description": "Arn representation of the Agent.",
        "description_kind": "plain",
        "type": "string"
      },
      "agent_id": {
        "computed": true,
        "description": "Identifier for a resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "agent_name": {
        "description": "Name for a resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "agent_resource_role_arn": {
        "computed": true,
        "description": "ARN of a IAM role.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "agent_status": {
        "computed": true,
        "description": "Schema Type for Action APIs.",
        "description_kind": "plain",
        "type": "string"
      },
      "agent_version": {
        "computed": true,
        "description": "Draft Agent Version.",
        "description_kind": "plain",
        "type": "string"
      },
      "auto_prepare": {
        "computed": true,
        "description": "Specifies whether to automatically prepare after creating or updating the agent.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "created_at": {
        "computed": true,
        "description": "Time Stamp.",
        "description_kind": "plain",
        "type": "string"
      },
      "customer_encryption_key_arn": {
        "computed": true,
        "description": "A KMS key ARN",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Description of the Resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "failure_reasons": {
        "computed": true,
        "description": "Failure Reasons for Error.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "foundation_model": {
        "computed": true,
        "description": "ARN or name of a Bedrock model.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "idle_session_ttl_in_seconds": {
        "computed": true,
        "description": "Max Session Time.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "instruction": {
        "computed": true,
        "description": "Instruction for the agent.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "knowledge_bases": {
        "computed": true,
        "description": "List of Agent Knowledge Bases",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "description": {
              "description": "Description of the Resource.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "knowledge_base_id": {
              "description": "Identifier for a resource.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "knowledge_base_state": {
              "computed": true,
              "description": "State of the knowledge base; whether it is enabled or disabled",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "prepared_at": {
        "computed": true,
        "description": "Time Stamp.",
        "description_kind": "plain",
        "type": "string"
      },
      "prompt_override_configuration": {
        "computed": true,
        "description": "Configuration for prompt override.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "override_lambda": {
              "computed": true,
              "description": "ARN of a Lambda.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "prompt_configurations": {
              "description": "List of BasePromptConfiguration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "base_prompt_template": {
                    "computed": true,
                    "description": "Base Prompt Template.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "inference_configuration": {
                    "computed": true,
                    "description": "Configuration for inference in prompt configuration",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "maximum_length": {
                          "computed": true,
                          "description": "Maximum length of output",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "stop_sequences": {
                          "computed": true,
                          "description": "List of stop sequences",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "temperature": {
                          "computed": true,
                          "description": "Controls randomness, higher values increase diversity",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "top_k": {
                          "computed": true,
                          "description": "Sample from the k most likely next tokens",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "top_p": {
                          "computed": true,
                          "description": "Cumulative probability cutoff for token selection",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "parser_mode": {
                    "computed": true,
                    "description": "Creation Mode for Prompt Configuration.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "prompt_creation_mode": {
                    "computed": true,
                    "description": "Creation Mode for Prompt Configuration.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "prompt_state": {
                    "computed": true,
                    "description": "Prompt State.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "prompt_type": {
                    "computed": true,
                    "description": "Prompt Type.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "required": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "recommended_actions": {
        "computed": true,
        "description": "The recommended actions users can take to resolve an error in failureReasons.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "skip_resource_in_use_check_on_delete": {
        "computed": true,
        "description": "Specifies whether to allow deleting agent while it is in use.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "updated_at": {
        "computed": true,
        "description": "Time Stamp.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Definition of AWS::Bedrock::Agent Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBedrockAgentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockAgent), &result)
	return &result
}
