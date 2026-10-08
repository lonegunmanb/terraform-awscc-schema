package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSmsvoiceNotifyConfiguration = `{
  "block": {
    "attributes": {
      "created_timestamp": {
        "computed": true,
        "description": "The time when the notify configuration was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "default_template_id": {
        "computed": true,
        "description": "The default template identifier to associate with the notify configuration. If specified, this template is used when sending messages without an explicit template identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_protection_enabled": {
        "computed": true,
        "description": "By default this is set to false. When set to true the notify configuration can't be deleted.",
        "description_kind": "plain",
        "type": "bool"
      },
      "display_name": {
        "computed": true,
        "description": "The display name to associate with the notify configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "enabled_channels": {
        "computed": true,
        "description": "An array of channels to enable for the notify configuration. Supported values include SMS and VOICE.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "enabled_countries": {
        "computed": true,
        "description": "An array of two-character ISO country codes, in ISO 3166-1 alpha-2 format, that are enabled for the notify configuration.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "notify_configuration_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the notify configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "notify_configuration_id": {
        "computed": true,
        "description": "The unique identifier for the notify configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "pool_id": {
        "computed": true,
        "description": "The identifier of the pool to associate with the notify configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The current status of the notify configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of tags (key and value pairs) associated with the notify configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key identifier, or name, of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The string value associated with the key of the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "tier": {
        "computed": true,
        "description": "The tier of the notify configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "tier_upgrade_status": {
        "computed": true,
        "description": "The tier upgrade status of the notify configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "use_case": {
        "computed": true,
        "description": "The use case for the notify configuration.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SMSVOICE::NotifyConfiguration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSmsvoiceNotifyConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSmsvoiceNotifyConfiguration), &result)
	return &result
}
