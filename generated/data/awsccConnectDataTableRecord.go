package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConnectDataTableRecord = `{
  "block": {
    "attributes": {
      "data_table_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "data_table_record": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "primary_values": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "attribute_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "attribute_value": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "values": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "attribute_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "attribute_value": {
                    "computed": true,
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
      "record_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Connect::DataTableRecord",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccConnectDataTableRecordSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectDataTableRecord), &result)
	return &result
}
