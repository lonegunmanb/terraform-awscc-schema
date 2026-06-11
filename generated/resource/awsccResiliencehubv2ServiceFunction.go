package resource

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
        "description": "The criticality of the service function.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the service function.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the service function.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_count": {
        "computed": true,
        "description": "The number of resources associated with this function.",
        "description_kind": "plain",
        "type": "number"
      },
      "service_arn": {
        "description": "The ARN of the parent service.",
        "description_kind": "plain",
        "required": true,
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
    "description": "Creates a service function within a Resilience Hub service.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccResiliencehubv2ServiceFunctionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccResiliencehubv2ServiceFunction), &result)
	return &result
}
