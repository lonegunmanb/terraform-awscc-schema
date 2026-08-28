package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCodecommitRepository = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the repository.",
        "description_kind": "plain",
        "type": "string"
      },
      "clone_url_http": {
        "computed": true,
        "description": "The URL to use for cloning the repository over HTTPS.",
        "description_kind": "plain",
        "type": "string"
      },
      "clone_url_ssh": {
        "computed": true,
        "description": "The URL to use for cloning the repository over SSH.",
        "description_kind": "plain",
        "type": "string"
      },
      "code": {
        "computed": true,
        "description": "Information about code to be committed to a repository after it is created in an AWS CloudFormation stack. Information about code is only used in resource creation. Updates to a stack will not reflect changes made to code properties after initial resource creation.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "branch_name": {
              "computed": true,
              "description": "Optional. Specifies a branch name to be used as the default branch when importing code into a repository on initial creation. If this property is not set, the name main will be used for the default branch for the repository. Changes to this property are ignored after initial resource creation. We recommend using this parameter to set the name to main to align with the default behavior of CodeCommit unless another name is needed.",
              "description_kind": "plain",
              "type": "string"
            },
            "s3": {
              "computed": true,
              "description": "Information about the Amazon S3 bucket that contains a ZIP file of code to be committed to the repository. Changes to this property are ignored after initial resource creation.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket": {
                    "computed": true,
                    "description": "The name of the Amazon S3 bucket that contains the ZIP file with the content that will be committed to the new repository. This can be specified using the name of the bucket in the AWS account. Changes to this property are ignored after initial resource creation.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "key": {
                    "computed": true,
                    "description": "The key to use for accessing the Amazon S3 bucket. Changes to this property are ignored after initial resource creation.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "object_version": {
                    "computed": true,
                    "description": "The object version of the ZIP file, if versioning is enabled for the Amazon S3 bucket. Changes to this property are ignored after initial resource creation.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
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
      "kms_key_id": {
        "computed": true,
        "description": "The ID of the AWS Key Management Service encryption key used to encrypt and decrypt the repository.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The repository's name.",
        "description_kind": "plain",
        "type": "string"
      },
      "repository_description": {
        "computed": true,
        "description": "A comment or description about the new repository.",
        "description_kind": "plain",
        "type": "string"
      },
      "repository_id": {
        "computed": true,
        "description": "The ID of the repository.",
        "description_kind": "plain",
        "type": "string"
      },
      "repository_name": {
        "computed": true,
        "description": "The name of the new repository to be created.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "One or more tag key-value pairs to use when tagging this repository.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag's key.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag's value.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "triggers": {
        "computed": true,
        "description": "Information about a trigger for a repository.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "branches": {
              "computed": true,
              "description": "The branches to be included in the trigger configuration. If you specify an empty array, the trigger applies to all branches.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "custom_data": {
              "computed": true,
              "description": "Any custom data associated with the trigger to be included in the information sent to the target of the trigger.",
              "description_kind": "plain",
              "type": "string"
            },
            "destination_arn": {
              "computed": true,
              "description": "The ARN of the resource that is the target for a trigger (for example, the ARN of a topic in Amazon SNS).",
              "description_kind": "plain",
              "type": "string"
            },
            "events": {
              "computed": true,
              "description": "The repository events that cause the trigger to run actions in another service, such as sending a notification through Amazon SNS.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "name": {
              "computed": true,
              "description": "The name of the trigger.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::CodeCommit::Repository",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCodecommitRepositorySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCodecommitRepository), &result)
	return &result
}
