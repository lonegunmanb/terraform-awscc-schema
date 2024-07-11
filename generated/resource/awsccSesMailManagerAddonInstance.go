package resource

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
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Definition of AWS::SES::MailManagerAddonInstance Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSesMailManagerAddonInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSesMailManagerAddonInstance), &result)
	return &result
}
