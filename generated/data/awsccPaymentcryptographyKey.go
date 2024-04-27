package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccPaymentcryptographyKey = `{
  "block": {
    "attributes": {
      "enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "exportable": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "key_attributes": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key_algorithm": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "key_class": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "key_modes_of_use": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "decrypt": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "derive_key": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "encrypt": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "generate": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "no_restrictions": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "sign": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "unwrap": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "verify": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "wrap": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "key_usage": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "key_check_value_algorithm": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "key_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "key_origin": {
        "computed": true,
        "description": "Defines the source of a key",
        "description_kind": "plain",
        "type": "string"
      },
      "key_state": {
        "computed": true,
        "description": "Defines the state of a key",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::PaymentCryptography::Key",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccPaymentcryptographyKeySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPaymentcryptographyKey), &result)
	return &result
}
