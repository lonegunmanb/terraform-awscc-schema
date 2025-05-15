package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2Host = `{
  "block": {
    "attributes": {
      "asset_id": {
        "computed": true,
        "description": "The ID of the Outpost hardware asset.",
        "description_kind": "plain",
        "type": "string"
      },
      "auto_placement": {
        "computed": true,
        "description": "Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "availability_zone": {
        "computed": true,
        "description": "The Availability Zone in which to allocate the Dedicated Host.",
        "description_kind": "plain",
        "type": "string"
      },
      "host_id": {
        "computed": true,
        "description": "ID of the host created.",
        "description_kind": "plain",
        "type": "string"
      },
      "host_maintenance": {
        "computed": true,
        "description": "Automatically allocates a new dedicated host and moves your instances on to it if a degradation is detected on your current host.",
        "description_kind": "plain",
        "type": "string"
      },
      "host_recovery": {
        "computed": true,
        "description": "Indicates whether to enable or disable host recovery for the Dedicated Host. Host recovery is disabled by default.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_family": {
        "computed": true,
        "description": "Specifies the instance family to be supported by the Dedicated Hosts. If you specify an instance family, the Dedicated Hosts support multiple instance types within that instance family.",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_type": {
        "computed": true,
        "description": "Specifies the instance type to be supported by the Dedicated Hosts. If you specify an instance type, the Dedicated Hosts support instances of the specified instance type only.",
        "description_kind": "plain",
        "type": "string"
      },
      "outpost_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the Amazon Web Services Outpost on which to allocate the Dedicated Host.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Any tags assigned to the Host.",
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
    "description": "Data Source schema for AWS::EC2::Host",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2HostSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2Host), &result)
	return &result
}
