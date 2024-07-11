package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSesMailManagerAddonInstance = `{
  "block": {
    "attributes": {
      "addon_instance_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "addon_instance_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "addon_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "addon_subscription_id": {
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
    "description": "Data Source schema for AWS::SES::MailManagerAddonInstance",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSesMailManagerAddonInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSesMailManagerAddonInstance), &result)
	return &result
}
