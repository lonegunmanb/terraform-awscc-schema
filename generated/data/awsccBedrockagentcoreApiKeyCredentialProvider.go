package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockagentcoreApiKeyCredentialProvider = `{
  "block": {
    "attributes": {
      "api_key": {
        "computed": true,
        "description": "The API key to use for authentication",
        "description_kind": "plain",
        "type": "string"
      },
      "api_key_secret_arn": {
        "computed": true,
        "description": "The ARN of the API key secret in AWS Secrets Manager",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "secret_arn": {
              "computed": true,
              "description": "The ARN of the secret in AWS Secrets Manager",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "created_time": {
        "computed": true,
        "description": "The timestamp when the credential provider was created",
        "description_kind": "plain",
        "type": "string"
      },
      "credential_provider_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the API key credential provider",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_updated_time": {
        "computed": true,
        "description": "The timestamp when the credential provider was last updated",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the API key credential provider",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags to assign to the API key credential provider",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::BedrockAgentCore::ApiKeyCredentialProvider",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBedrockagentcoreApiKeyCredentialProviderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockagentcoreApiKeyCredentialProvider), &result)
	return &result
}
