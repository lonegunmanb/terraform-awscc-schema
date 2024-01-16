package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccArczonalshiftZonalAutoshiftConfiguration = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
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
              "type": [
                "list",
                "string"
              ]
            },
            "blocked_windows": {
              "computed": true,
              "description_kind": "plain",
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
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "outcome_alarms": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "alarm_identifier": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "resource_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "zonal_autoshift_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ARCZonalShift::ZonalAutoshiftConfiguration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccArczonalshiftZonalAutoshiftConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccArczonalshiftZonalAutoshiftConfiguration), &result)
	return &result
}
