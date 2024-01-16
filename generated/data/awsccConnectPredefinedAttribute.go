package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConnectPredefinedAttribute = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_arn": {
        "computed": true,
        "description": "The identifier of the Amazon Connect instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the predefined attribute.",
        "description_kind": "plain",
        "type": "string"
      },
      "values": {
        "computed": true,
        "description": "The values of a predefined attribute.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "string_list": {
              "computed": true,
              "description": "Predefined attribute values of type string list.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::Connect::PredefinedAttribute",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccConnectPredefinedAttributeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectPredefinedAttribute), &result)
	return &result
}
