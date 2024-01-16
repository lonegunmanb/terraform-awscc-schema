package resource

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
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "identity_center_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the specified AWS Identity Center.",
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
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "The AWS::S3::AccessGrantsInstance resource is an Amazon S3 resource type that hosts Access Grants and their associated locations",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccS3AccessGrantsInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccS3AccessGrantsInstance), &result)
	return &result
}
