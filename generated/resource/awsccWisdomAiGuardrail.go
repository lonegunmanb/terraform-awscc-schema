package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWisdomAiGuardrail = `{
  "block": {
    "attributes": {
      "ai_guardrail_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ai_guardrail_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "assistant_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "assistant_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
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
                    "description": "Type of text to text filter in content policy",
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
      "description": {
        "computed": true,
        "description": "Description of the guardrail or its version",
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
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
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
    "description": "Definition of AWS::Wisdom::AIGuardrail Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccWisdomAiGuardrailSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWisdomAiGuardrail), &result)
	return &result
}
