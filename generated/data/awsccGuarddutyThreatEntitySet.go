package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGuarddutyThreatEntitySet = `{
  "block": {
    "attributes": {
      "activate": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "detector_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "error_details": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "expected_bucket_owner": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "format": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "threat_entity_set_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "updated_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::GuardDuty::ThreatEntitySet",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccGuarddutyThreatEntitySetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGuarddutyThreatEntitySet), &result)
	return &result
}
