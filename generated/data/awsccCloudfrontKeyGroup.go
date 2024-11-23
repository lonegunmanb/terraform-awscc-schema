package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudfrontKeyGroup = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "key_group_config": {
        "computed": true,
        "description": "The key group configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "comment": {
              "computed": true,
              "description": "A comment to describe the key group. The comment cannot be longer than 128 characters.",
              "description_kind": "plain",
              "type": "string"
            },
            "items": {
              "computed": true,
              "description": "A list of the identifiers of the public keys in the key group.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "name": {
              "computed": true,
              "description": "A name to identify the key group.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "key_group_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::CloudFront::KeyGroup",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCloudfrontKeyGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudfrontKeyGroup), &result)
	return &result
}
