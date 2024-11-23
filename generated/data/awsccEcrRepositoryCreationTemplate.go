package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEcrRepositoryCreationTemplate = `{
  "block": {
    "attributes": {
      "applied_for": {
        "computed": true,
        "description": "A list of enumerable Strings representing the repository creation scenarios that the template will apply towards.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "created_at": {
        "computed": true,
        "description": "Create timestamp of the template.",
        "description_kind": "plain",
        "type": "string"
      },
      "custom_role_arn": {
        "computed": true,
        "description": "The ARN of the role to be assumed by ECR. This role must be in the same account as the registry that you are configuring.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the template.",
        "description_kind": "plain",
        "type": "string"
      },
      "encryption_configuration": {
        "computed": true,
        "description": "The encryption configuration for the repository. This determines how the contents of your repository are encrypted at rest. By default, when no encryption configuration is set or the AES256 encryption type is used, Amazon ECR uses server-side encryption with Amazon S3-managed encryption keys which encrypts your data at rest using an AES-256 encryption algorithm. This does not require any action on your part.\n\nFor more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/encryption-at-rest.html",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "encryption_type": {
              "computed": true,
              "description": "The encryption type to use.",
              "description_kind": "plain",
              "type": "string"
            },
            "kms_key": {
              "computed": true,
              "description": "If you use the KMS or KMS_DSSE encryption type, specify the CMK to use for encryption. The alias, key ID, or full ARN of the CMK can be specified. The key must exist in the same Region as the repository. If no key is specified, the default AWS managed CMK for Amazon ECR will be used.",
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
      "image_tag_mutability": {
        "computed": true,
        "description": "The image tag mutability setting for the repository.",
        "description_kind": "plain",
        "type": "string"
      },
      "lifecycle_policy": {
        "computed": true,
        "description": "The JSON lifecycle policy text to apply to the repository. For information about lifecycle policy syntax, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html",
        "description_kind": "plain",
        "type": "string"
      },
      "prefix": {
        "computed": true,
        "description": "The prefix use to match the repository name and apply the template.",
        "description_kind": "plain",
        "type": "string"
      },
      "repository_policy": {
        "computed": true,
        "description": "The JSON repository policy text to apply to the repository. For more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/RepositoryPolicyExamples.html",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "updated_at": {
        "computed": true,
        "description": "Update timestamp of the template.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ECR::RepositoryCreationTemplate",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEcrRepositoryCreationTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEcrRepositoryCreationTemplate), &result)
	return &result
}
