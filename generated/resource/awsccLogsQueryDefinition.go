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
      "parameters": {
        "computed": true,
        "description": "Use this parameter to include specific query parameters as part of your query definition. Query parameters are supported only for Logs Insights QL queries. Query parameters allow you to use placeholder variables in your query string that are substituted with values at execution time. Use the {{parameterName}} syntax in your query string to reference a parameter.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "default_value": {
              "computed": true,
              "description": "The default value to use for this query parameter if no value is supplied at execution time.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "description": {
              "computed": true,
              "description": "A description of the query parameter that explains its purpose or expected values.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "The name of the query parameter. A query parameter name must start with a letter or underscore, and contain only letters, digits, and underscores.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
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
        "optional": true,
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
