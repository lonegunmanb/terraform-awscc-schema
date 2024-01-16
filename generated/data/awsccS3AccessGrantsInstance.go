package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccS3AccessGrantsInstance = `{
  "block": {
    "attributes": {
      "access_grants_instance_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the specified Access Grants instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "access_grants_instance_id": {
        "computed": true,
        "description": "A unique identifier for the specified access grants instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "identity_center_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the specified AWS Identity Center.",
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
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::S3::AccessGrantsInstance",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccS3AccessGrantsInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccS3AccessGrantsInstance), &result)
	return &result
}
