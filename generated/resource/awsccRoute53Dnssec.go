package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53Dnssec = `{
  "block": {
    "attributes": {
      "hosted_zone_id": {
        "description": "The unique string (ID) used to identify a hosted zone.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource used to control (enable/disable) DNSSEC in a specific hosted zone.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRoute53DnssecSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53Dnssec), &result)
	return &result
}
