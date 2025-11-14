package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMediaconnectRouterNetworkInterface = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "associated_input_count": {
        "computed": true,
        "description": "The number of router inputs associated with the network interface.",
        "description_kind": "plain",
        "type": "number"
      },
      "associated_output_count": {
        "computed": true,
        "description": "The number of router outputs associated with the network interface.",
        "description_kind": "plain",
        "type": "number"
      },
      "configuration": {
        "description": "The configuration settings for a router network interface.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "public": {
              "computed": true,
              "description": "The configuration settings for a public router network interface, including the list of allowed CIDR blocks.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "allow_rules": {
                    "computed": true,
                    "description": "The list of allowed CIDR blocks for the public router network interface.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "cidr": {
                          "computed": true,
                          "description": "The CIDR block that is allowed to access the public router network interface.",
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
                "nesting_mode": "single"
              },
              "optional": true
            },
            "vpc": {
              "computed": true,
              "description": "The configuration settings for a router network interface within a VPC, including the security group IDs and subnet ID.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "security_group_ids": {
                    "computed": true,
                    "description": "The IDs of the security groups to associate with the router network interface within the VPC.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "subnet_id": {
                    "computed": true,
                    "description": "The ID of the subnet within the VPC to associate the router network interface with.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp when the router network interface was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the router network interface.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_interface_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "region_name": {
        "computed": true,
        "description": "The AWS Region for the router network interface. Defaults to the current region if not specified.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "router_network_interface_id": {
        "computed": true,
        "description": "The unique identifier of the router network interface.",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Key-value pairs that can be used to tag and organize this router network interface.",
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
      "updated_at": {
        "computed": true,
        "description": "The timestamp when the router network interface was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Represents a router network interface in AWS Elemental MediaConnect that is used to define a network boundary for router resources",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMediaconnectRouterNetworkInterfaceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMediaconnectRouterNetworkInterface), &result)
	return &result
}
