package resource

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
        "optional": true,
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
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_updated_time": {
        "computed": true,
        "description": "The timestamp when the workload identity was last updated.",
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "description": "The name of the workload identity. The name must be unique within your account.",
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
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
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
      "workload_identity_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the workload identity.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Definition of AWS::BedrockAgentCore::WorkloadIdentity Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBedrockagentcoreWorkloadIdentitySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockagentcoreWorkloadIdentity), &result)
	return &result
}
