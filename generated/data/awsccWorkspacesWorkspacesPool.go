package data

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
              "type": "string"
            },
            "status": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "bundle_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "capacity": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "desired_user_sessions": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "directory_id": {
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
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "running_mode": {
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
      },
      "timeout_settings": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "disconnect_timeout_in_seconds": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "idle_disconnect_timeout_in_seconds": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "max_user_duration_in_seconds": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::WorkSpaces::WorkspacesPool",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccWorkspacesWorkspacesPoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWorkspacesWorkspacesPool), &result)
	return &result
}
