package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudfrontCloudfrontOriginAccessIdentity = `{
  "block": {
    "attributes": {
      "cloudfront_origin_access_identity_config": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "comment": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "s3_canonical_user_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::CloudFront::CloudFrontOriginAccessIdentity",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudfrontCloudfrontOriginAccessIdentitySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudfrontCloudfrontOriginAccessIdentity), &result)
	return &result
}
