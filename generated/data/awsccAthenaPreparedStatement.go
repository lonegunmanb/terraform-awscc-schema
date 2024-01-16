package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAthenaPreparedStatement = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "The description of the prepared statement.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "query_statement": {
        "computed": true,
        "description": "The query string for the prepared statement.",
        "description_kind": "plain",
        "type": "string"
      },
      "statement_name": {
        "computed": true,
        "description": "The name of the prepared statement.",
        "description_kind": "plain",
        "type": "string"
      },
      "work_group": {
        "computed": true,
        "description": "The name of the workgroup to which the prepared statement belongs.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Athena::PreparedStatement",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAthenaPreparedStatementSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAthenaPreparedStatement), &result)
	return &result
}
