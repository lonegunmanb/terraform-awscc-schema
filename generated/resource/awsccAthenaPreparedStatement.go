package resource

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
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "query_statement": {
        "description": "The query string for the prepared statement.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "statement_name": {
        "description": "The name of the prepared statement.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "work_group": {
        "description": "The name of the workgroup to which the prepared statement belongs.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource schema for AWS::Athena::PreparedStatement",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAthenaPreparedStatementSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAthenaPreparedStatement), &result)
	return &result
}
