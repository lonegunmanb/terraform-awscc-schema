package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccTransferProfile = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Specifies the unique Amazon Resource Name (ARN) for the profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "as_2_id": {
        "description": "AS2 identifier agreed with a trading partner.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "certificate_ids": {
        "computed": true,
        "description": "List of the certificate IDs associated with this profile to be used for encryption and signing of AS2 messages.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "profile_id": {
        "computed": true,
        "description": "A unique identifier for the profile",
        "description_kind": "plain",
        "type": "string"
      },
      "profile_type": {
        "description": "Enum specifying whether the profile is local or associated with a trading partner.",
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
              "description": "The name assigned to the tag that you create.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "Contains one or more values that you assigned to the key name you create.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::Transfer::Profile",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccTransferProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccTransferProfile), &result)
	return &result
}
