package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerCodeRepository = `{
  "block": {
    "attributes": {
      "code_repository_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the code repository.",
        "description_kind": "plain",
        "type": "string"
      },
      "code_repository_name": {
        "computed": true,
        "description": "The name of the Git repository.",
        "description_kind": "plain",
        "type": "string"
      },
      "git_config": {
        "computed": true,
        "description": "Configuration details for the Git repository, including the URL where it is located and the ARN of the AWS Secrets Manager secret that contains the credentials used to access the repository.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "branch": {
              "computed": true,
              "description": "The default branch for the Git repository.",
              "description_kind": "plain",
              "type": "string"
            },
            "repository_url": {
              "computed": true,
              "description": "The URL where the Git repository is located.",
              "description_kind": "plain",
              "type": "string"
            },
            "secret_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the AWS Secrets Manager secret that contains the credentials used to access the git repository.",
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
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::SageMaker::CodeRepository",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSagemakerCodeRepositorySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerCodeRepository), &result)
	return &result
}
