package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccOpensearchmemorySpace = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "encryption_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kms_key_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "event_expiry_days": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "extraction_model_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "memory_expiry_days": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "memory_strategies": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "consolidation_enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "extract_every_n_messages": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "extract_every_n_tokens": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "extraction_prompt": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "extraction_seed_recent_n": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "extraction_seed_top_k": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "idle_timeout_seconds": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "description": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "space_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::OpenSearchMemory::Space",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccOpensearchmemorySpaceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOpensearchmemorySpace), &result)
	return &result
}
