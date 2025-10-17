package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2NatGateway = `{
  "block": {
    "attributes": {
      "allocation_id": {
        "computed": true,
        "description": "[Public NAT gateway only] The allocation ID of the Elastic IP address that's associated with the NAT gateway. This property is required for a public NAT gateway and cannot be specified with a private NAT gateway.",
        "description_kind": "plain",
        "type": "string"
      },
      "connectivity_type": {
        "computed": true,
        "description": "Indicates whether the NAT gateway supports public or private connectivity. The default is public connectivity.",
        "description_kind": "plain",
        "type": "string"
      },
      "eni_id": {
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
      "max_drain_duration_seconds": {
        "computed": true,
        "description": "The maximum amount of time to wait (in seconds) before forcibly releasing the IP addresses if connections are still in progress. Default value is 350 seconds.",
        "description_kind": "plain",
        "type": "number"
      },
      "nat_gateway_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "private_ip_address": {
        "computed": true,
        "description": "The private IPv4 address to assign to the NAT gateway. If you don't provide an address, a private IPv4 address will be automatically assigned.",
        "description_kind": "plain",
        "type": "string"
      },
      "secondary_allocation_ids": {
        "computed": true,
        "description": "Secondary EIP allocation IDs. For more information, see [Create a NAT gateway](https://docs.aws.amazon.com/vpc/latest/userguide/nat-gateway-working-with.html) in the *Amazon VPC User Guide*.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "secondary_private_ip_address_count": {
        "computed": true,
        "description": "[Private NAT gateway only] The number of secondary private IPv4 addresses you want to assign to the NAT gateway. For more information about secondary addresses, see [Create a NAT gateway](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html#nat-gateway-creating) in the *Amazon Virtual Private Cloud User Guide*.\n ` + "`" + `` + "`" + `SecondaryPrivateIpAddressCount` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `SecondaryPrivateIpAddresses` + "`" + `` + "`" + ` cannot be set at the same time.",
        "description_kind": "plain",
        "type": "number"
      },
      "secondary_private_ip_addresses": {
        "computed": true,
        "description": "Secondary private IPv4 addresses. For more information about secondary addresses, see [Create a NAT gateway](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html#nat-gateway-creating) in the *Amazon Virtual Private Cloud User Guide*.\n ` + "`" + `` + "`" + `SecondaryPrivateIpAddressCount` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `SecondaryPrivateIpAddresses` + "`" + `` + "`" + ` cannot be set at the same time.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "subnet_id": {
        "computed": true,
        "description": "The ID of the subnet in which the NAT gateway is located.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags for the NAT gateway.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "vpc_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::NatGateway",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2NatGatewaySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2NatGateway), &result)
	return &result
}
