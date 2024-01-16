package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAthenaCapacityReservation = `{
  "block": {
    "attributes": {
      "allocated_dpus": {
        "computed": true,
        "description": "The number of DPUs Athena has provisioned and allocated for the reservation",
        "description_kind": "plain",
        "type": "number"
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the specified capacity reservation",
        "description_kind": "plain",
        "type": "string"
      },
      "capacity_assignment_configuration": {
        "computed": true,
        "description": "Assignment configuration to assign workgroups to a reservation",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "capacity_assignments": {
              "description": "List of capacity assignments",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "workgroup_names": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
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
      "creation_time": {
        "computed": true,
        "description": "The date and time the reservation was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_successful_allocation_time": {
        "computed": true,
        "description": "The timestamp when the last successful allocated was made",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The reservation name.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the reservation.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "target_dpus": {
        "description": "The number of DPUs to request to be allocated to the reservation.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      }
    },
    "description": "Resource schema for AWS::Athena::CapacityReservation",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAthenaCapacityReservationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAthenaCapacityReservation), &result)
	return &result
}
