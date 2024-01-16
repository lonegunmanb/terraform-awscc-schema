package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEventschemasRegistryPolicy = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "registry_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "revision_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EventSchemas::RegistryPolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEventschemasRegistryPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEventschemasRegistryPolicy), &result)
	return &result
}
