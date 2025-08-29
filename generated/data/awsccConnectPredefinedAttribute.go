package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConnectPredefinedAttribute = `{
  "block": {
    "attributes": {
      "attribute_configuration": {
        "computed": true,
        "description": "Custom metadata associated to a Predefined attribute that controls how the attribute behaves when used by upstream services.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enable_value_validation_on_association": {
              "computed": true,
              "description": "Enables customers to enforce strict validation on the specific values that this predefined attribute can hold.",
              "description_kind": "plain",
              "type": "bool"
            },
            "is_read_only": {
              "computed": true,
              "description": "Allows the predefined attribute to show up and be managed in the Amazon Connect UI.",
              "description_kind": "plain",
              "type": "bool"
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
        "description": "The identifier of the Amazon Connect instance.",
        "description_kind": "plain",
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
        "computed": true,
        "description": "The name of the predefined attribute.",
        "description_kind": "plain",
        "type": "string"
      },
      "purposes": {
        "computed": true,
        "description": "The assigned purposes of the predefined attribute.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
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
