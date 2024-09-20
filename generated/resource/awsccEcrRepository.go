package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEcrRepository = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "empty_on_delete": {
        "computed": true,
        "description": "If true, deleting the repository force deletes the contents of the repository. If false, the repository must be empty before attempting to delete it.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "encryption_configuration": {
        "computed": true,
        "description": "The encryption configuration for the repository. This determines how the contents of your repository are encrypted at rest.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "encryption_type": {
              "computed": true,
              "description": "The encryption type to use.\n If you use the ` + "`" + `` + "`" + `KMS` + "`" + `` + "`" + ` encryption type, the contents of the repository will be encrypted using server-side encryption with KMSlong key stored in KMS. When you use KMS to encrypt your data, you can either use the default AWS managed KMS key for Amazon ECR, or specify your own KMS key, which you already created. For more information, see [Protecting data using server-side encryption with an key stored in (SSE-KMS)](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingKMSEncryption.html) in the *Amazon Simple Storage Service Console Developer Guide*.\n If you use the ` + "`" + `` + "`" + `AES256` + "`" + `` + "`" + ` encryption type, Amazon ECR uses server-side encryption with Amazon S3-managed encryption keys which encrypts the images in the repository using an AES-256 encryption algorithm. For more information, see [Protecting data using server-side encryption with Amazon S3-managed encryption keys (SSE-S3)](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingServerSideEncryption.html) in the *Amazon Simple Storage Service Console Developer Guide*.",
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
      "image_scanning_configuration": {
        "computed": true,
        "description": "The image scanning configuration for the repository. This determines whether images are scanned for known vulnerabilities after being pushed to the repository.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "scan_on_push": {
              "computed": true,
              "description": "The setting that determines whether images are scanned after being pushed to a repository. If set to ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `, images will be scanned after being pushed. If this parameter is not specified, it will default to ` + "`" + `` + "`" + `false` + "`" + `` + "`" + ` and images will not be scanned unless a scan is manually started.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "image_tag_mutability": {
        "computed": true,
        "description": "The tag mutability setting for the repository. If this parameter is omitted, the default setting of ` + "`" + `` + "`" + `MUTABLE` + "`" + `` + "`" + ` will be used which will allow image tags to be overwritten. If ` + "`" + `` + "`" + `IMMUTABLE` + "`" + `` + "`" + ` is specified, all image tags within the repository will be immutable which will prevent them from being overwritten.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "lifecycle_policy": {
        "computed": true,
        "description": "Creates or updates a lifecycle policy. For information about lifecycle policy syntax, see [Lifecycle policy template](https://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html).",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "lifecycle_policy_text": {
              "computed": true,
              "description": "The JSON repository policy text to apply to the repository.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "registry_id": {
              "computed": true,
              "description": "The AWS account ID associated with the registry that contains the repository. If you do? not specify a registry, the default registry is assumed.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "repository_name": {
        "computed": true,
        "description": "The name to use for the repository. The repository name may be specified on its own (such as ` + "`" + `` + "`" + `nginx-web-app` + "`" + `` + "`" + `) or it can be prepended with a namespace to group the repository into a category (such as ` + "`" + `` + "`" + `project-a/nginx-web-app` + "`" + `` + "`" + `). If you don't specify a name, CFNlong generates a unique physical ID and uses that ID for the repository name. For more information, see [Name type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).\n The repository name must start with a letter and can only contain lowercase letters, numbers, hyphens, underscores, and forward slashes.\n  If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "repository_policy_text": {
        "computed": true,
        "description": "The JSON repository policy text to apply to the repository. For more information, see [Amazon ECR repository policies](https://docs.aws.amazon.com/AmazonECR/latest/userguide/repository-policy-examples.html) in the *Amazon Elastic Container Registry User Guide*.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "repository_uri": {
        "computed": true,
        "description_kind": "plain",
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
      }
    },
    "description": "The ` + "`" + `` + "`" + `AWS::ECR::Repository` + "`" + `` + "`" + ` resource specifies an Amazon Elastic Container Registry (Amazon ECR) repository, where users can push and pull Docker images, Open Container Initiative (OCI) images, and OCI compatible artifacts. For more information, see [Amazon ECR private repositories](https://docs.aws.amazon.com/AmazonECR/latest/userguide/Repositories.html) in the *Amazon ECR User Guide*.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEcrRepositorySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEcrRepository), &result)
	return &result
}
