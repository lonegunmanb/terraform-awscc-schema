package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudfrontCloudfrontOriginAccessIdentity = `{
  "block": {
    "attributes": {
      "cloudfront_origin_access_identity_config": {
        "description": "The current configuration information for the identity.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "comment": {
              "description": "A comment to describe the origin access identity. The comment cannot be longer than 128 characters.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "cloudfront_origin_access_identity_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "s3_canonical_user_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "The request to create a new origin access identity (OAI). An origin access identity is a special CloudFront user that you can associate with Amazon S3 origins, so that you can secure all or just some of your Amazon S3 content. For more information, see [Restricting Access to Amazon S3 Content by Using an Origin Access Identity](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-restricting-access-to-s3.html) in the *Amazon CloudFront Developer Guide*.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudfrontCloudfrontOriginAccessIdentitySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudfrontCloudfrontOriginAccessIdentity), &result)
	return &result
}
