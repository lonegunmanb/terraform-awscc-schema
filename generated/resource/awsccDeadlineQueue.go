package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDeadlineQueue = `{
  "block": {
    "attributes": {
      "allowed_storage_profile_ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "default_budget_action": {
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
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "farm_id": {
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
      "job_attachment_settings": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "root_prefix": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "s3_bucket_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "job_run_as_user": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "posix": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "group": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "user": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "run_as": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "windows": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "password_arn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "user": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "queue_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "required_file_system_location_names": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "role_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Definition of AWS::Deadline::Queue Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDeadlineQueueSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDeadlineQueue), &result)
	return &result
}
