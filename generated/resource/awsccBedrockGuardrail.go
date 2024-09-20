package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockGuardrail = `{
  "block": {
    "attributes": {
      "blocked_input_messaging": {
        "description": "Messaging for when violations are detected in text",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "blocked_outputs_messaging": {
        "description": "Messaging for when violations are detected in text",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "content_policy_config": {
        "computed": true,
        "description": "Content policy config for a guardrail.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "filters_config": {
              "computed": true,
              "description": "List of content filter configs in content policy.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "input_strength": {
                    "computed": true,
                    "description": "Strength for filters",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "output_strength": {
                    "computed": true,
                    "description": "Strength for filters",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "type": {
                    "computed": true,
                    "description": "Type of filter in content policy",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "contextual_grounding_policy_config": {
        "computed": true,
        "description": "Contextual grounding policy config for a guardrail.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "filters_config": {
              "computed": true,
              "description": "List of contextual grounding filter configs.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "threshold": {
                    "computed": true,
                    "description": "The threshold for this filter.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "type": {
                    "computed": true,
                    "description": "Type of contextual grounding filter",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
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
        "description": "Time Stamp",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Description of the guardrail or its version",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "failure_recommendations": {
        "computed": true,
        "description": "List of failure recommendations",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "guardrail_arn": {
        "computed": true,
        "description": "Arn representation for the guardrail",
        "description_kind": "plain",
        "type": "string"
      },
      "guardrail_id": {
        "computed": true,
        "description": "Unique id for the guardrail",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_arn": {
        "computed": true,
        "description": "The KMS key with which the guardrail was encrypted at rest",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "Name of the guardrail",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sensitive_information_policy_config": {
        "computed": true,
        "description": "Sensitive information policy config for a guardrail.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "pii_entities_config": {
              "computed": true,
              "description": "List of entities.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "action": {
                    "computed": true,
                    "description": "Options for sensitive information action.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "type": {
                    "computed": true,
                    "description": "The currently supported PII entities",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "regexes_config": {
              "computed": true,
              "description": "List of regex.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "action": {
                    "computed": true,
                    "description": "Options for sensitive information action.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "description": {
                    "computed": true,
                    "description": "The regex description.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "The regex name.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "pattern": {
                    "computed": true,
                    "description": "The regex pattern.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "status": {
        "computed": true,
        "description": "Status of the guardrail",
        "description_kind": "plain",
        "type": "string"
      },
      "status_reasons": {
        "computed": true,
        "description": "List of status reasons",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "List of Tags",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "Tag Key",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "Tag Value",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "topic_policy_config": {
        "computed": true,
        "description": "Topic policy config for a guardrail.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "topics_config": {
              "computed": true,
              "description": "List of topic configs in topic policy.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "definition": {
                    "computed": true,
                    "description": "Definition of topic in topic policy",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "examples": {
                    "computed": true,
                    "description": "List of text examples",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "name": {
                    "computed": true,
                    "description": "Name of topic in topic policy",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "type": {
                    "computed": true,
                    "description": "Type of topic in a policy",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "updated_at": {
        "computed": true,
        "description": "Time Stamp",
        "description_kind": "plain",
        "type": "string"
      },
      "version": {
        "computed": true,
        "description": "Guardrail version",
        "description_kind": "plain",
        "type": "string"
      },
      "word_policy_config": {
        "computed": true,
        "description": "Word policy config for a guardrail.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "managed_word_lists_config": {
              "computed": true,
              "description": "A config for the list of managed words.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "type": {
                    "computed": true,
                    "description": "Options for managed words.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "words_config": {
              "computed": true,
              "description": "List of custom word configs.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "text": {
                    "computed": true,
                    "description": "The custom word text.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      }
    },
    "description": "Definition of AWS::Bedrock::Guardrail Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBedrockGuardrailSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockGuardrail), &result)
	return &result
}
