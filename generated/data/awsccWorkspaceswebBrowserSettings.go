package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWorkspaceswebBrowserSettings = `{
  "block": {
    "attributes": {
      "additional_encryption_context": {
        "computed": true,
        "description_kind": "plain",
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
      "browser_policy": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "browser_settings_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "customer_managed_key": {
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
      },
      "web_content_filtering_policy": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "allowed_urls": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "blocked_categories": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "blocked_urls": {
              "computed": true,
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
    "description": "Data Source schema for AWS::WorkSpacesWeb::BrowserSettings",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccWorkspaceswebBrowserSettingsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWorkspaceswebBrowserSettings), &result)
	return &result
}
