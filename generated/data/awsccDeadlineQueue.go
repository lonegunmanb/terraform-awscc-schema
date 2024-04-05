package data

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
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "farm_id": {
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
      "job_attachment_settings": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "root_prefix": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "s3_bucket_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
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
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "user": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "run_as": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "windows": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "password_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "user": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "queue_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "required_file_system_location_names": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "role_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Deadline::Queue",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDeadlineQueueSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDeadlineQueue), &result)
	return &result
}
