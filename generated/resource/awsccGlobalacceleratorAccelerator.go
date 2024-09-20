package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGlobalacceleratorAccelerator = `{
  "block": {
    "attributes": {
      "accelerator_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the accelerator.",
        "description_kind": "plain",
        "type": "string"
      },
      "dns_name": {
        "computed": true,
        "description": "The Domain Name System (DNS) name that Global Accelerator creates that points to your accelerator's static IPv4 addresses.",
        "description_kind": "plain",
        "type": "string"
      },
      "dual_stack_dns_name": {
        "computed": true,
        "description": "The Domain Name System (DNS) name that Global Accelerator creates that points to your accelerator's static IPv4 and IPv6 addresses.",
        "description_kind": "plain",
        "type": "string"
      },
      "enabled": {
        "computed": true,
        "description": "Indicates whether an accelerator is enabled. The value is true or false.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "ip_address_type": {
        "computed": true,
        "description": "IP Address type.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ip_addresses": {
        "computed": true,
        "description": "The IP addresses from BYOIP Prefix pool.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "ipv_4_addresses": {
        "computed": true,
        "description": "The IPv4 addresses assigned to the accelerator.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "ipv_6_addresses": {
        "computed": true,
        "description": "The IPv6 addresses assigned if the accelerator is dualstack",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "name": {
        "description": "Name of accelerator.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "Key of the tag. Value can be 1 to 127 characters.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "Value for the tag. Value can be 1 to 255 characters.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::GlobalAccelerator::Accelerator",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccGlobalacceleratorAcceleratorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGlobalacceleratorAccelerator), &result)
	return &result
}
