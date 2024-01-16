package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudfrontCloudfrontOriginAccessIdentity = `{
  "block": {
    "attributes": {
      "cloudfront_origin_access_identity_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "comment": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "s3_canonical_user_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::CloudFront::CloudFrontOriginAccessIdentity",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCloudfrontCloudfrontOriginAccessIdentitySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudfrontCloudfrontOriginAccessIdentity), &result)
	return &result
}
