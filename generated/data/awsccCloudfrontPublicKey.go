package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudfrontPublicKey = `{
  "block": {
    "attributes": {
      "created_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "public_key_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "caller_reference": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "comment": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "encoded_key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::CloudFront::PublicKey",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCloudfrontPublicKeySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudfrontPublicKey), &result)
	return &result
}
