package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWorkspaceswebIpAccessSettings = `{
  "block": {
    "attributes": {
      "additional_encryption_context": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "associated_portal_arns": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "creation_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "customer_managed_key": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "ip_access_settings_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ip_rules": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "description": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ip_range": {
              "description": "A single IP address or an IP address range in CIDR notation",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Definition of AWS::WorkSpacesWeb::IpAccessSettings Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccWorkspaceswebIpAccessSettingsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWorkspaceswebIpAccessSettings), &result)
	return &result
}
