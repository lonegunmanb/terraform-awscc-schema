package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIvsStreamKey = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Stream Key ARN is automatically generated on creation and assigned as the unique identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "channel_arn": {
        "description": "Channel ARN for the stream.",
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
      "tags": {
        "computed": true,
        "description": "A list of key-value pairs that contain metadata for the asset model.",
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
      },
      "value": {
        "computed": true,
        "description": "Stream-key value.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::IVS::StreamKey",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIvsStreamKeySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIvsStreamKey), &result)
	return &result
}
