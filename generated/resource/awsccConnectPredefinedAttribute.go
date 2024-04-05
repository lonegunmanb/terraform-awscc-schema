package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConnectPredefinedAttribute = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_arn": {
        "description": "The identifier of the Amazon Connect instance.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_modified_region": {
        "computed": true,
        "description": "Last modified region.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description": "Last modified time.",
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "description": "The name of the predefined attribute.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "values": {
        "description": "The values of a predefined attribute.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "string_list": {
              "computed": true,
              "description": "Predefined attribute values of type string list.",
              "description_kind": "plain",
              "optional": true,
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
    "description": "Resource Type definition for AWS::Connect::PredefinedAttribute",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccConnectPredefinedAttributeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectPredefinedAttribute), &result)
	return &result
}
