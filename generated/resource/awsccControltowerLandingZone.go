package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccControltowerLandingZone = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "drift_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "landing_zone_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "latest_available_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "manifest": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "remediation_types": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "version": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Definition of AWS::ControlTower::LandingZone Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccControltowerLandingZoneSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccControltowerLandingZone), &result)
	return &result
}
