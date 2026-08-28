package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSdbDomain = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "Information about the SimpleDB domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_id": {
        "computed": true,
        "description": "The name of the domain to create.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SDB::Domain",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSdbDomainSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSdbDomain), &result)
	return &result
}
