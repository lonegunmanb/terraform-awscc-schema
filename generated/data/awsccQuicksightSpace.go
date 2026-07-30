package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccQuicksightSpace = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the space.",
        "description_kind": "plain",
        "type": "string"
      },
      "aws_account_id": {
        "computed": true,
        "description": "The ID of the Amazon Web Services account where the space is being created.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The date and time the space was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_by": {
        "computed": true,
        "description": "The user name of the principal who created the space.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the space.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The display name of the space.",
        "description_kind": "plain",
        "type": "string"
      },
      "permissions": {
        "computed": true,
        "description": "A list of permissions granted on the space.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "actions": {
              "computed": true,
              "description": "The list of actions granted to the principal.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "principal": {
              "computed": true,
              "description": "The ARN of the principal (user or group) receiving the permission.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "resources": {
        "computed": true,
        "description": "A list of QuickSight resources attached to the space.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "resource_arn": {
              "computed": true,
              "description": "The ARN of the QuickSight resource.",
              "description_kind": "plain",
              "type": "string"
            },
            "resource_type": {
              "computed": true,
              "description": "The type of QuickSight resource.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "space_id": {
        "computed": true,
        "description": "The unique identifier for the space.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of key-value pairs to associate with the space resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "updated_at": {
        "computed": true,
        "description": "The date and time the space was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::QuickSight::Space",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccQuicksightSpaceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccQuicksightSpace), &result)
	return &result
}
