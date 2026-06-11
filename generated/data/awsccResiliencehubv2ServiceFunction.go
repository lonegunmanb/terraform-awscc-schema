package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccResiliencehubv2ServiceFunction = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "The timestamp when the service function was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "criticality": {
        "computed": true,
        "description": "The criticality of the service function.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the service function.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the service function.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_count": {
        "computed": true,
        "description": "The number of resources associated with this function.",
        "description_kind": "plain",
        "type": "number"
      },
      "service_arn": {
        "computed": true,
        "description": "The ARN of the parent service.",
        "description_kind": "plain",
        "type": "string"
      },
      "service_function_id": {
        "computed": true,
        "description": "The server-generated service function ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "source": {
        "computed": true,
        "description": "The source of the service function.",
        "description_kind": "plain",
        "type": "string"
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp when the service function was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ResilienceHubV2::ServiceFunction",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccResiliencehubv2ServiceFunctionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccResiliencehubv2ServiceFunction), &result)
	return &result
}
