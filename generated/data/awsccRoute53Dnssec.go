package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53Dnssec = `{
  "block": {
    "attributes": {
      "hosted_zone_id": {
        "computed": true,
        "description": "The unique string (ID) used to identify a hosted zone.",
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
    "description": "Data Source schema for AWS::Route53::DNSSEC",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRoute53DnssecSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53Dnssec), &result)
	return &result
}
