package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppsyncType = `{
  "block": {
    "attributes": {
      "api_id": {
        "description": "The API ID.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the type.",
        "description_kind": "plain",
        "type": "string"
      },
      "definition": {
        "description": "The type definition, in GraphQL Schema Definition Language (SDL) format.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "format": {
        "description": "The type format: SDL or JSON.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The type name.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Represents a GraphQL type in an AWS AppSync GraphQL API.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAppsyncTypeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppsyncType), &result)
	return &result
}
