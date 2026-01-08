package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEcrRepositoryCreationTemplate = `{
  "block": {
    "attributes": {
      "applied_for": {
        "description": "A list of enumerable Strings representing the repository creation scenarios that this template will apply towards. The supported scenarios are PULL_THROUGH_CACHE, REPLICATION, and CREATE_ON_PUSH",
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "custom_role_arn": {
        "computed": true,
        "description": "The ARN of the role to be assumed by Amazon ECR. Amazon ECR will assume your supplied role when the customRoleArn is specified. When this field isn't specified, Amazon ECR will use the service-linked role for the repository creation template.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description associated with the repository creation template.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "encryption_configuration": {
        "computed": true,
        "description": "The encryption configuration associated with the repository creation template.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "encryption_type": {
              "computed": true,
              "description": "The encryption type to use.\n If you use the ` + "`" + `` + "`" + `KMS` + "`" + `` + "`" + ` encryption type, the contents of the repository will be encrypted using server-side encryption with KMSlong key stored in KMS. When you use KMS to encrypt your data, you can either use the default AWS managed KMS key for Amazon ECR, or specify your own KMS key, which you already created.\n If you use the ` + "`" + `` + "`" + `KMS_DSSE` + "`" + `` + "`" + ` encryption type, the contents of the repository will be encrypted with two layers of encryption using server-side encryption with the KMS Management Service key stored in KMS. Similar to the ` + "`" + `` + "`" + `KMS` + "`" + `` + "`" + ` encryption type, you can either use the default AWS managed KMS key for Amazon ECR, or specify your own KMS key, which you've already created. \n If you use the ` + "`" + `` + "`" + `AES256` + "`" + `` + "`" + ` encryption type, Amazon ECR uses server-side encryption with Amazon S3-managed encryption keys which encrypts the images in the repository using an AES256 encryption algorithm.\n For more information, see [Amazon ECR encryption at rest](https://docs.aws.amazon.com/AmazonECR/latest/userguide/encryption-at-rest.html) in the *Amazon Elastic Container Registry User Guide*.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "kms_key": {
              "computed": true,
              "description": "If you use the ` + "`" + `` + "`" + `KMS` + "`" + `` + "`" + ` encryption type, specify the KMS key to use for encryption. The alias, key ID, or full ARN of the KMS key can be specified. The key must exist in the same Region as the repository. If no key is specified, the default AWS managed KMS key for Amazon ECR will be used.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "image_tag_mutability": {
        "computed": true,
        "description": "The tag mutability setting for the repository. If this parameter is omitted, the default setting of ` + "`" + `` + "`" + `MUTABLE` + "`" + `` + "`" + ` will be used which will allow image tags to be overwritten. If ` + "`" + `` + "`" + `IMMUTABLE` + "`" + `` + "`" + ` is specified, all image tags within the repository will be immutable which will prevent them from being overwritten.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "image_tag_mutability_exclusion_filters": {
        "computed": true,
        "description": "A list of filters that specify which image tags are excluded from the repository creation template's image tag mutability setting.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "image_tag_mutability_exclusion_filter_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "image_tag_mutability_exclusion_filter_value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "lifecycle_policy": {
        "computed": true,
        "description": "The lifecycle policy to use for repositories created using the template.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "prefix": {
        "description": "The repository namespace prefix associated with the repository creation template.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "repository_policy": {
        "computed": true,
        "description": "The repository policy to apply to repositories created using the template. A repository policy is a permissions policy associated with a repository to control access permissions.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_tags": {
        "computed": true,
        "description": "The metadata to apply to the repository to help you categorize and organize. Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "One part of a key-value pair that make up a tag. A ` + "`" + `` + "`" + `key` + "`" + `` + "`" + ` is a general label that acts like a category for more specific tag values.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "A ` + "`" + `` + "`" + `value` + "`" + `` + "`" + ` acts as a descriptor within a tag category (key).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "updated_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "The details of the repository creation template associated with the request.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEcrRepositoryCreationTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEcrRepositoryCreationTemplate), &result)
	return &result
}
