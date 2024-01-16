package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccS3AccessGrantsLocation = `{
  "block": {
    "attributes": {
      "access_grants_location_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the specified Access Grants location.",
        "description_kind": "plain",
        "type": "string"
      },
      "access_grants_location_id": {
        "computed": true,
        "description": "The unique identifier for the specified Access Grants location.",
        "description_kind": "plain",
        "type": "string"
      },
      "iam_role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the access grant location's associated IAM role.",
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
      "location_scope": {
        "computed": true,
        "description": "Descriptor for where the location actually points",
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
    "description": "The AWS::S3::AccessGrantsLocation resource is an Amazon S3 resource type hosted in an access grants instance which can be the target of S3 access grants.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccS3AccessGrantsLocationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccS3AccessGrantsLocation), &result)
	return &result
}
