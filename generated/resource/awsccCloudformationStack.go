package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudformationStack = `{
  "block": {
    "attributes": {
      "capabilities": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "change_set_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "creation_time": {
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
      "disable_rollback": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_termination_protection": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_update_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "notification_ar_ns": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "outputs": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "description": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "export_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "output_key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "output_value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "parameters": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "parent_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "root_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "stack_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "stack_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "stack_policy_body": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "stack_policy_url": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "stack_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "stack_status_reason": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
      "template_body": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "template_url": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "timeout_in_minutes": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "description": "The AWS::CloudFormation::Stack resource nests a stack as a resource in a top-level template.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudformationStackSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudformationStack), &result)
	return &result
}
