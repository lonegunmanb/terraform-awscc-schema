package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccChimeAppInstanceUser = `{
  "block": {
    "attributes": {
      "app_instance_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "app_instance_user_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "app_instance_user_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "expiration_settings": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "expiration_criterion": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "expiration_days": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
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
      "metadata": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
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
    "description": "Data Source schema for AWS::Chime::AppInstanceUser",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccChimeAppInstanceUserSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccChimeAppInstanceUser), &result)
	return &result
}
