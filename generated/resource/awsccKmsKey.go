package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccKmsKey = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bypass_policy_lockout_safety_check": {
        "computed": true,
        "description": "Skips (\"bypasses\") the key policy lockout safety check. The default value is false.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "description": {
        "computed": true,
        "description": "A description of the AWS KMS key. Use a description that helps you to distinguish this AWS KMS key from others in the account, such as its intended use.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_key_rotation": {
        "computed": true,
        "description": "Enables automatic rotation of the key material for the specified AWS KMS key. By default, automation key rotation is not enabled.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enabled": {
        "computed": true,
        "description": "Specifies whether the AWS KMS key is enabled. Disabled AWS KMS keys cannot be used in cryptographic operations.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "key_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "key_policy": {
        "computed": true,
        "description": "The key policy that authorizes use of the AWS KMS key. The key policy must observe the following rules.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "key_spec": {
        "computed": true,
        "description": "Specifies the type of AWS KMS key to create. The default value is SYMMETRIC_DEFAULT. This property is required only for asymmetric AWS KMS keys. You can't change the KeySpec value after the AWS KMS key is created.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "key_usage": {
        "computed": true,
        "description": "Determines the cryptographic operations for which you can use the AWS KMS key. The default value is ENCRYPT_DECRYPT. This property is required only for asymmetric AWS KMS keys. You can't change the KeyUsage value after the AWS KMS key is created.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "multi_region": {
        "computed": true,
        "description": "Specifies whether the AWS KMS key should be Multi-Region. You can't change the MultiRegion value after the AWS KMS key is created.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "origin": {
        "computed": true,
        "description": "The source of the key material for the KMS key. You cannot change the origin after you create the KMS key. The default is AWS_KMS, which means that AWS KMS creates the key material.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pending_window_in_days": {
        "computed": true,
        "description": "Specifies the number of days in the waiting period before AWS KMS deletes an AWS KMS key that has been removed from a CloudFormation stack. Enter a value between 7 and 30 days. The default value is 30 days.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "The AWS::KMS::Key resource specifies an AWS KMS key in AWS Key Management Service (AWS KMS). Authorized users can use the AWS KMS key to encrypt and decrypt small amounts of data (up to 4096 bytes), but they are more commonly used to generate data keys. You can also use AWS KMS keys to encrypt data stored in AWS services that are integrated with AWS KMS or within their applications.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccKmsKeySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccKmsKey), &result)
	return &result
}
