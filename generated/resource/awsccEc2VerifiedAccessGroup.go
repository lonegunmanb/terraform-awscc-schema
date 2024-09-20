package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2VerifiedAccessGroup = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description": "Time this Verified Access Group was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description for the AWS Verified Access group.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_updated_time": {
        "computed": true,
        "description": "Time this Verified Access Group was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "owner": {
        "computed": true,
        "description": "The AWS account number that owns the group.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_document": {
        "computed": true,
        "description": "The AWS Verified Access policy document.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "policy_enabled": {
        "computed": true,
        "description": "The status of the Verified Access policy.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "sse_specification": {
        "computed": true,
        "description": "The configuration options for customer provided KMS encryption.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "customer_managed_key_enabled": {
              "computed": true,
              "description": "Whether to encrypt the policy with the provided key or disable encryption",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "kms_key_arn": {
              "computed": true,
              "description": "KMS Key Arn used to encrypt the group policy",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
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
      },
      "verified_access_group_arn": {
        "computed": true,
        "description": "The ARN of the Verified Access group.",
        "description_kind": "plain",
        "type": "string"
      },
      "verified_access_group_id": {
        "computed": true,
        "description": "The ID of the AWS Verified Access group.",
        "description_kind": "plain",
        "type": "string"
      },
      "verified_access_instance_id": {
        "description": "The ID of the AWS Verified Access instance.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "The AWS::EC2::VerifiedAccessGroup resource creates an AWS EC2 Verified Access Group.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2VerifiedAccessGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2VerifiedAccessGroup), &result)
	return &result
}
