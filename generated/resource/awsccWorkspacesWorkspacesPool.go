package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWorkspacesWorkspacesPool = `{
  "block": {
    "attributes": {
      "application_settings": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "settings_group": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "status": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "bundle_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "capacity": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "desired_user_sessions": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "directory_id": {
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
      "pool_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "pool_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "pool_name": {
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
      },
      "timeout_settings": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "disconnect_timeout_in_seconds": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "idle_disconnect_timeout_in_seconds": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_user_duration_in_seconds": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::WorkSpaces::WorkspacesPool",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccWorkspacesWorkspacesPoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWorkspacesWorkspacesPool), &result)
	return &result
}
