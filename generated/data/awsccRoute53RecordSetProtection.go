package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53RecordSetProtection = `{
  "block": {
    "attributes": {
      "comment": {
        "computed": true,
        "description": "A comment that describes the protection.",
        "description_kind": "plain",
        "type": "string"
      },
      "hosted_zone_id": {
        "computed": true,
        "description": "The ID of the hosted zone that contains the record set to protect.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "record_name": {
        "computed": true,
        "description": "The name of the record set to protect.",
        "description_kind": "plain",
        "type": "string"
      },
      "record_type": {
        "computed": true,
        "description": "The DNS record type of the record set to protect.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Route53::RecordSetProtection",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRoute53RecordSetProtectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53RecordSetProtection), &result)
	return &result
}
