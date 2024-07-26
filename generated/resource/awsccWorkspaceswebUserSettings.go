package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWorkspaceswebUserSettings = `{
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
      "cookie_synchronization_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "allowlist": {
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "domain": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "path": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "required": true
            },
            "blocklist": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "domain": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "path": {
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
          "nesting_mode": "single"
        },
        "optional": true
      },
      "copy_allowed": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "customer_managed_key": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "deep_link_allowed": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disconnect_timeout_in_minutes": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "download_allowed": {
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
      "idle_disconnect_timeout_in_minutes": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "paste_allowed": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "print_allowed": {
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
      },
      "upload_allowed": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "user_settings_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Definition of AWS::WorkSpacesWeb::UserSettings Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccWorkspaceswebUserSettingsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWorkspaceswebUserSettings), &result)
	return &result
}
