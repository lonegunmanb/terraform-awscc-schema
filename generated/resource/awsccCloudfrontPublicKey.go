package resource

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
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "public_key_config": {
        "description": "Configuration information about a public key that you can use with [signed URLs and signed cookies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html), or with [field-level encryption](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/field-level-encryption.html).",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "caller_reference": {
              "description": "A string included in the request to help make sure that the request can't be replayed.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "comment": {
              "computed": true,
              "description": "A comment to describe the public key. The comment cannot be longer than 128 characters.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "encoded_key": {
              "description": "The public key that you can use with [signed URLs and signed cookies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html), or with [field-level encryption](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/field-level-encryption.html).",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "name": {
              "description": "A name to help identify the public key.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "public_key_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "A public key that you can use with [signed URLs and signed cookies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html), or with [field-level encryption](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/field-level-encryption.html).\n CloudFront supports signed URLs and signed cookies with RSA 2048 or ECDSA 256 key signatures. Field-level encryption is only compatible with RSA 2048 key signatures.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudfrontPublicKeySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudfrontPublicKey), &result)
	return &result
}
