package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccChimeAppInstance = `{
  "block": {
    "attributes": {
      "app_instance_arn": {
        "computed": true,
        "description": "The Amazon Resource Number (ARN) of the AppInstance.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_timestamp": {
        "computed": true,
        "description": "The time at which an AppInstance was created. In epoch milliseconds.",
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_updated_timestamp": {
        "computed": true,
        "description": "The time an AppInstance was last updated. In epoch milliseconds.",
        "description_kind": "plain",
        "type": "number"
      },
      "metadata": {
        "computed": true,
        "description": "The metadata of the AppInstance. Limited to a 1KB string in UTF-8.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the AppInstance.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags assigned to the AppInstance.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key in a tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value in a tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::Chime::AppInstance",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccChimeAppInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccChimeAppInstance), &result)
	return &result
}
