package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConnectDataTableAttribute = `{
  "block": {
    "attributes": {
      "attribute_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "data_table_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
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
      "instance_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified_region": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "lock_version": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "attribute": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "data_table": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "validation": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enum": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "strict": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "values": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              }
            },
            "exclusive_maximum": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "exclusive_minimum": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "max_length": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "max_values": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "maximum": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "min_length": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "min_values": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "minimum": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "multiple_of": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "value_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Connect::DataTableAttribute",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccConnectDataTableAttributeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectDataTableAttribute), &result)
	return &result
}
