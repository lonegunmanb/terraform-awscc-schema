package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockagentcoreWorkloadIdentity = `{
  "block": {
    "attributes": {
      "allowed_resource_oauth_2_return_urls": {
        "computed": true,
        "description": "The list of allowed OAuth2 return URLs for resources associated with this workload identity.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "created_time": {
        "computed": true,
        "description": "The timestamp when the workload identity was created.",
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_updated_time": {
        "computed": true,
        "description": "The timestamp when the workload identity was last updated.",
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "computed": true,
        "description": "The name of the workload identity. The name must be unique within your account.",
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
      },
      "workload_identity_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the workload identity.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::BedrockAgentCore::WorkloadIdentity",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBedrockagentcoreWorkloadIdentitySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockagentcoreWorkloadIdentity), &result)
	return &result
}
