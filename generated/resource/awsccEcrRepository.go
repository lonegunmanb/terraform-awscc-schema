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
        "description": "If true, deleting the repository force deletes the contents of the repository. Without a force delete, you can only delete empty repositories.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "encryption_configuration": {
        "computed": true,
        "description": "The encryption configuration for the repository. This determines how the contents of your repository are encrypted at rest.\n\nBy default, when no encryption configuration is set or the AES256 encryption type is used, Amazon ECR uses server-side encryption with Amazon S3-managed encryption keys which encrypts your data at rest using an AES-256 encryption algorithm. This does not require any action on your part.\n\nFor more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/encryption-at-rest.html",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "encryption_type": {
              "description": "The encryption type to use.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "kms_key": {
              "computed": true,
              "description": "If you use the KMS encryption type, specify the CMK to use for encryption. The alias, key ID, or full ARN of the CMK can be specified. The key must exist in the same Region as the repository. If no key is specified, the default AWS managed CMK for Amazon ECR will be used.",
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
        "description": "The image scanning configuration for the repository. This setting determines whether images are scanned for known vulnerabilities after being pushed to the repository.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "scan_on_push": {
              "computed": true,
              "description": "The setting that determines whether images are scanned after being pushed to a repository.",
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
        "description": "The image tag mutability setting for the repository.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "lifecycle_policy": {
        "computed": true,
        "description": "The LifecyclePolicy property type specifies a lifecycle policy. For information about lifecycle policy syntax, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html",
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
              "description": "The AWS account ID associated with the registry that contains the repository. If you do not specify a registry, the default registry is assumed. ",
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
        "description": "The name to use for the repository. The repository name may be specified on its own (such as nginx-web-app) or it can be prepended with a namespace to group the repository into a category (such as project-a/nginx-web-app). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the repository name. For more information, see https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "repository_policy_text": {
        "computed": true,
        "description": "The JSON repository policy text to apply to the repository. For more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/RepositoryPolicyExamples.html in the Amazon Elastic Container Registry User Guide. ",
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
              "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
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
    "description": "The AWS::ECR::Repository resource specifies an Amazon Elastic Container Registry (Amazon ECR) repository, where users can push and pull Docker images. For more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/Repositories.html",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEcrRepositorySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEcrRepository), &result)
	return &result
}
