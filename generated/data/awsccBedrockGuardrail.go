package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockGuardrail = `{
  "block": {
    "attributes": {
      "blocked_input_messaging": {
        "computed": true,
        "description": "Messaging for when violations are detected in text",
        "description_kind": "plain",
        "type": "string"
      },
      "blocked_outputs_messaging": {
        "computed": true,
        "description": "Messaging for when violations are detected in text",
        "description_kind": "plain",
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
                  "input_action": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "input_enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "input_modalities": {
                    "computed": true,
                    "description": "List of modalities",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "input_strength": {
                    "computed": true,
                    "description": "Strength for filters",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "output_action": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "output_enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "output_modalities": {
                    "computed": true,
                    "description": "List of modalities",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "output_strength": {
                    "computed": true,
                    "description": "Strength for filters",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "type": {
                    "computed": true,
                    "description": "Type of filter in content policy",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
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
                  "action": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "threshold": {
                    "computed": true,
                    "description": "The threshold for this filter.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "type": {
                    "computed": true,
                    "description": "Type of contextual grounding filter",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "created_at": {
        "computed": true,
        "description": "Time Stamp",
        "description_kind": "plain",
        "type": "string"
      },
      "cross_region_config": {
        "computed": true,
        "description": "The system-defined guardrail profile that you?re using with your guardrail",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "guardrail_profile_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the guardrail profile",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "description": {
        "computed": true,
        "description": "Description of the guardrail or its version",
        "description_kind": "plain",
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
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kms_key_arn": {
        "computed": true,
        "description": "The KMS key with which the guardrail was encrypted at rest",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Name of the guardrail",
        "description_kind": "plain",
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
                    "type": "string"
                  },
                  "input_action": {
                    "computed": true,
                    "description": "Options for sensitive information action.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "input_enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "output_action": {
                    "computed": true,
                    "description": "Options for sensitive information action.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "output_enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "type": {
                    "computed": true,
                    "description": "The currently supported PII entities",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
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
                    "type": "string"
                  },
                  "description": {
                    "computed": true,
                    "description": "The regex description.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "input_action": {
                    "computed": true,
                    "description": "Options for sensitive information action.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "input_enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "name": {
                    "computed": true,
                    "description": "The regex name.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "output_action": {
                    "computed": true,
                    "description": "Options for sensitive information action.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "output_enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "pattern": {
                    "computed": true,
                    "description": "The regex pattern.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
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
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "Tag Value",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
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
                    "type": "string"
                  },
                  "examples": {
                    "computed": true,
                    "description": "List of text examples",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "input_action": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "input_enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "name": {
                    "computed": true,
                    "description": "Name of topic in topic policy",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "output_action": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "output_enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "type": {
                    "computed": true,
                    "description": "Type of topic in a policy",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
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
                  "input_action": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "input_enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "output_action": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "output_enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "type": {
                    "computed": true,
                    "description": "Options for managed words.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "words_config": {
              "computed": true,
              "description": "List of custom word configs.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "input_action": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "input_enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "output_action": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "output_enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "text": {
                    "computed": true,
                    "description": "The custom word text.",
                    "description_kind": "plain",
                    "type": "string"
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
    "description": "Data Source schema for AWS::Bedrock::Guardrail",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBedrockGuardrailSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockGuardrail), &result)
	return &result
}
