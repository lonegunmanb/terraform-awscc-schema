package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppsyncType = `{
  "block": {
    "attributes": {
      "api_id": {
        "computed": true,
        "description": "The API ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the type.",
        "description_kind": "plain",
        "type": "string"
      },
      "definition": {
        "computed": true,
        "description": "The type definition, in GraphQL Schema Definition Language (SDL) format.",
        "description_kind": "plain",
        "type": "string"
      },
      "format": {
        "computed": true,
        "description": "The type format: SDL or JSON.",
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
        "description": "The type name.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::AppSync::Type",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAppsyncTypeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppsyncType), &result)
	return &result
}
