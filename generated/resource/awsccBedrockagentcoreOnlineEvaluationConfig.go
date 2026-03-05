package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockagentcoreOnlineEvaluationConfig = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "The timestamp when the online evaluation configuration was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_source_config": {
        "description": "The data source configuration that specifies CloudWatch log groups and service names to monitor.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cloudwatch_logs": {
              "description": "The configuration for reading agent traces from CloudWatch logs.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "log_group_names": {
                    "description": "The list of CloudWatch log group names to monitor for agent traces.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "service_names": {
                    "description": "The list of service names to filter traces within the specified log groups.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "description": {
        "computed": true,
        "description": "The description of the online evaluation configuration.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "evaluation_execution_role_arn": {
        "description": "The Amazon Resource Name (ARN) of the IAM role that grants permissions for evaluation.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "evaluators": {
        "description": "The list of evaluators to apply during online evaluation.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "evaluator_id": {
              "description": "The unique identifier of the evaluator.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "execution_status": {
        "computed": true,
        "description": "The execution status indicating whether the online evaluation is currently running.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "online_evaluation_config_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the online evaluation configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "online_evaluation_config_id": {
        "computed": true,
        "description": "The unique identifier of the online evaluation configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "online_evaluation_config_name": {
        "description": "The name of the online evaluation configuration. Must be unique within your account.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "output_config": {
        "computed": true,
        "description": "The configuration that specifies where evaluation results should be written.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cloudwatch_config": {
              "computed": true,
              "description": "The CloudWatch configuration for writing evaluation results.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "log_group_name": {
                    "computed": true,
                    "description": "The CloudWatch log group name for evaluation results.",
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
      "rule": {
        "description": "The evaluation rule that defines sampling configuration, filters, and session detection settings.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "filters": {
              "computed": true,
              "description": "The list of filters that determine which agent traces should be included in the evaluation.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "key": {
                    "computed": true,
                    "description": "The key or field name to filter on within the agent trace data.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "operator": {
                    "computed": true,
                    "description": "The comparison operator to use for filtering.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description": "The value used in filter comparisons.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "boolean_value": {
                          "computed": true,
                          "description": "The boolean value for true/false filtering conditions.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "double_value": {
                          "computed": true,
                          "description": "The numeric value for numerical filtering.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "string_value": {
                          "computed": true,
                          "description": "The string value for text-based filtering.",
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
                "nesting_mode": "list"
              },
              "optional": true
            },
            "sampling_config": {
              "description": "The configuration that controls what percentage of agent traces are sampled for evaluation.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "sampling_percentage": {
                    "description": "The percentage of agent traces to sample for evaluation.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            },
            "session_config": {
              "computed": true,
              "description": "The configuration that defines how agent sessions are detected.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "session_timeout_minutes": {
                    "computed": true,
                    "description": "The number of minutes of inactivity after which an agent session is considered complete.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
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
      "status": {
        "computed": true,
        "description": "The status of the online evaluation configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of tags to assign to the online evaluation configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp when the online evaluation configuration was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::BedrockAgentCore::OnlineEvaluationConfig - Creates an online evaluation configuration for continuous monitoring of agent performance.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBedrockagentcoreOnlineEvaluationConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockagentcoreOnlineEvaluationConfig), &result)
	return &result
}
