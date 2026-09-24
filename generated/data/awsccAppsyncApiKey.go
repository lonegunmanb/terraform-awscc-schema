package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppsyncApiKey = `{
  "block": {
    "attributes": {
      "api_id": {
        "computed": true,
        "description": "Unique AWS AppSync GraphQL API ID for this API key.",
        "description_kind": "plain",
        "type": "string"
      },
      "api_key": {
        "computed": true,
        "description": "The API key.",
        "description_kind": "plain",
        "type": "string"
      },
      "api_key_id": {
        "computed": true,
        "description": "The API key ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the API key.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Unique description of your API key.",
        "description_kind": "plain",
        "type": "string"
      },
      "expires": {
        "computed": true,
        "description": "The time after which the API key expires. The date is represented as seconds since the epoch, rounded down to the nearest hour.",
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::AppSync::ApiKey",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAppsyncApiKeySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppsyncApiKey), &result)
	return &result
}
