package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudwatchAccessProfile = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the access profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "assume_status": {
        "computed": true,
        "description": "Whether the calling principal is authorized to assume the access profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The time the access profile was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the access profile.",
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
      "name": {
        "description": "The name of the access profile.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "profile_id": {
        "computed": true,
        "description": "The unique identifier of the access profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "profile_type": {
        "computed": true,
        "description": "Whether the access profile is managed by the service or by the customer.",
        "description_kind": "plain",
        "type": "string"
      },
      "space_id": {
        "description": "The unique identifier of the space the access profile belongs to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "updated_at": {
        "computed": true,
        "description": "The time the access profile was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource type definition for AWS::CloudWatch::AccessProfile",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudwatchAccessProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudwatchAccessProfile), &result)
	return &result
}
