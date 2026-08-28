package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDeadlineWorker = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the worker.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The date and time the resource was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_by": {
        "computed": true,
        "description": "The user or system that created this resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "farm_id": {
        "description": "The farm ID.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "fleet_id": {
        "description": "The fleet ID.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "host_properties": {
        "computed": true,
        "description": "The IP address and host name of the worker.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "host_name": {
              "computed": true,
              "description": "The host name.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ip_addresses": {
              "computed": true,
              "description": "The IP addresses for a host.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "ip_v4_addresses": {
                    "computed": true,
                    "description": "The IpV4 address of the network.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "ip_v6_addresses": {
                    "computed": true,
                    "description": "The IpV6 address for the network and node component.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the worker.",
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
      "worker_id": {
        "computed": true,
        "description": "The worker ID.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Definition of AWS::Deadline::Worker Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDeadlineWorkerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDeadlineWorker), &result)
	return &result
}
