package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudfrontOriginAccessControl = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "origin_access_control_config": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "description": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "origin_access_control_origin_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "signing_behavior": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "signing_protocol": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      }
    },
    "description": "Resource Type definition for AWS::CloudFront::OriginAccessControl",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudfrontOriginAccessControlSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudfrontOriginAccessControl), &result)
	return &result
}
