package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockagentcoreHarnessEndpoint = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp when the endpoint was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint_name": {
        "computed": true,
        "description": "The name of the endpoint. Must start with a letter and contain only alphanumeric characters and underscores.",
        "description_kind": "plain",
        "type": "string"
      },
      "harness_id": {
        "computed": true,
        "description": "The ID of the harness that the endpoint belongs to.",
        "description_kind": "plain",
        "type": "string"
      },
      "harness_name": {
        "computed": true,
        "description": "The name of the harness that the endpoint belongs to.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "live_version": {
        "computed": true,
        "description": "The harness version that the endpoint is currently serving.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags to apply to the harness endpoint resource.",
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
      "target_version": {
        "computed": true,
        "description": "The harness version that the endpoint points to and serves invocations from.",
        "description_kind": "plain",
        "type": "string"
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp when the endpoint was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::BedrockAgentCore::HarnessEndpoint",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBedrockagentcoreHarnessEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockagentcoreHarnessEndpoint), &result)
	return &result
}
