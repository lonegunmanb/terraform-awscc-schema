package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudfrontOriginAccessControl = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "origin_access_control_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "description": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "origin_access_control_origin_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "signing_behavior": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "signing_protocol": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::CloudFront::OriginAccessControl",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCloudfrontOriginAccessControlSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudfrontOriginAccessControl), &result)
	return &result
}
