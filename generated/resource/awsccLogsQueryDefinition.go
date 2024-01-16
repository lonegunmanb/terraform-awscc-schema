package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLogsQueryDefinition = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "log_group_names": {
        "computed": true,
        "description": "Optionally define specific log groups as part of your query definition",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "name": {
        "description": "A name for the saved query definition",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "query_definition_id": {
        "computed": true,
        "description": "Unique identifier of a query definition",
        "description_kind": "plain",
        "type": "string"
      },
      "query_string": {
        "description": "The query string to use for this definition",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "The resource schema for AWSLogs QueryDefinition",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLogsQueryDefinitionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLogsQueryDefinition), &result)
	return &result
}
