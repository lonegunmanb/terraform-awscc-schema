package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEventschemasRegistryPolicy = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "policy": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "registry_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "revision_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::EventSchemas::RegistryPolicy",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEventschemasRegistryPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEventschemasRegistryPolicy), &result)
	return &result
}
