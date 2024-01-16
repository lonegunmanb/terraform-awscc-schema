package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccArczonalshiftZonalAutoshiftConfiguration = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "practice_run_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "blocked_dates": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "blocked_windows": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "blocking_alarms": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "alarm_identifier": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "outcome_alarms": {
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "alarm_identifier": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "required": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "resource_identifier": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "zonal_autoshift_status": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Definition of AWS::ARCZonalShift::ZonalAutoshiftConfiguration Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccArczonalshiftZonalAutoshiftConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccArczonalshiftZonalAutoshiftConfiguration), &result)
	return &result
}
