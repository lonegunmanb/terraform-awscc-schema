package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccKmsReplicaKey = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the AWS KMS key. Use a description that helps you to distinguish this AWS KMS key from others in the account, such as its intended use.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
        "description": "The key policy that authorizes use of the AWS KMS key. The key policy must observe the following rules.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "pending_window_in_days": {
        "computed": true,
        "description": "Specifies the number of days in the waiting period before AWS KMS deletes an AWS KMS key that has been removed from a CloudFormation stack. Enter a value between 7 and 30 days. The default value is 30 days.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "primary_key_arn": {
        "description": "Identifies the primary AWS KMS key to create a replica of. Specify the Amazon Resource Name (ARN) of the AWS KMS key. You cannot specify an alias or key ID. For help finding the ARN, see Finding the Key ID and ARN in the AWS Key Management Service Developer Guide.",
        "description_kind": "plain",
        "required": true,
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
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "The AWS::KMS::ReplicaKey resource specifies a multi-region replica AWS KMS key in AWS Key Management Service (AWS KMS).",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccKmsReplicaKeySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccKmsReplicaKey), &result)
	return &result
}
