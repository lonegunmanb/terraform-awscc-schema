package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAthenaNamedQuery = `{
  "block": {
    "attributes": {
      "database": {
        "computed": true,
        "description": "The database to which the query belongs.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The query description.",
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
        "description": "The query name.",
        "description_kind": "plain",
        "type": "string"
      },
      "named_query_id": {
        "computed": true,
        "description": "The unique ID of the query.",
        "description_kind": "plain",
        "type": "string"
      },
      "query_string": {
        "computed": true,
        "description": "The contents of the query with all query statements.",
        "description_kind": "plain",
        "type": "string"
      },
      "work_group": {
        "computed": true,
        "description": "The name of the workgroup that contains the named query.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Athena::NamedQuery",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAthenaNamedQuerySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAthenaNamedQuery), &result)
	return &result
}
