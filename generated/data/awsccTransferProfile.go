package data

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
        "computed": true,
        "description": "AS2 identifier agreed with a trading partner.",
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_ids": {
        "computed": true,
        "description": "List of the certificate IDs associated with this profile to be used for encryption and signing of AS2 messages.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "profile_id": {
        "computed": true,
        "description": "A unique identifier for the profile",
        "description_kind": "plain",
        "type": "string"
      },
      "profile_type": {
        "computed": true,
        "description": "Enum specifying whether the profile is local or associated with a trading partner.",
        "description_kind": "plain",
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
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "Contains one or more values that you assigned to the key name you create.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::Transfer::Profile",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccTransferProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccTransferProfile), &result)
	return &result
}
