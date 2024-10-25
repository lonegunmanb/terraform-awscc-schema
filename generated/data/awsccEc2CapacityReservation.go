package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2CapacityReservation = `{
  "block": {
    "attributes": {
      "availability_zone": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "available_instance_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "capacity_reservation_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ebs_optimized": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "end_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "end_date_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ephemeral_storage": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "instance_match_criteria": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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
      "out_post_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "placement_group_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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
      "total_instance_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "unused_reservation_billing_owner_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::CapacityReservation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2CapacityReservationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2CapacityReservation), &result)
	return &result
}
