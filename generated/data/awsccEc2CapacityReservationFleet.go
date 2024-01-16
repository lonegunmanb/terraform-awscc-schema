package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2CapacityReservationFleet = `{
  "block": {
    "attributes": {
      "allocation_strategy": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "capacity_reservation_fleet_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "end_date": {
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
      "instance_match_criteria": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "instance_type_specifications": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "availability_zone": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "availability_zone_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "ebs_optimized": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "instance_platform": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "instance_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "priority": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "weight": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "set"
        }
      },
      "no_remove_end_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "remove_end_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "tag_specifications": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "resource_type": {
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
            }
          },
          "nesting_mode": "list"
        }
      },
      "tenancy": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "total_target_capacity": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description": "Data Source schema for AWS::EC2::CapacityReservationFleet",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2CapacityReservationFleetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2CapacityReservationFleet), &result)
	return &result
}
