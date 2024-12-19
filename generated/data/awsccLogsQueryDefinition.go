package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLogsQueryDefinition = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "log_group_names": {
        "computed": true,
        "description": "Optionally define specific log groups as part of your query definition",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "name": {
        "computed": true,
        "description": "A name for the saved query definition",
        "description_kind": "plain",
        "type": "string"
      },
      "query_definition_id": {
        "computed": true,
        "description": "Unique identifier of a query definition",
        "description_kind": "plain",
        "type": "string"
      },
      "query_language": {
        "computed": true,
        "description": "Query language of the query string. Possible values are CWLI, SQL, PPL, with CWLI being the default.",
        "description_kind": "plain",
        "type": "string"
      },
      "query_string": {
        "computed": true,
        "description": "The query string to use for this definition",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Logs::QueryDefinition",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLogsQueryDefinitionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLogsQueryDefinition), &result)
	return &result
}
