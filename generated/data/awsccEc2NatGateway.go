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
      "auto_provision_zones": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "auto_scaling_ips": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "availability_mode": {
        "computed": true,
        "description": "Indicates whether this is a zonal (single-AZ) or regional (multi-AZ) NAT gateway.\n A zonal NAT gateway is a NAT Gateway that provides redundancy and scalability within a single availability zone. A regional NAT gateway is a single NAT Gateway that works across multiple availability zones (AZs) in your VPC, providing redundancy, scalability and availability across all the AZs in a Region.\n For more information, see [Regional NAT gateways for automatic multi-AZ expansion](https://docs.aws.amazon.com/vpc/latest/userguide/nat-gateways-regional.html) in the *Amazon VPC User Guide*.",
        "description_kind": "plain",
        "type": "string"
      },
      "availability_zone_addresses": {
        "computed": true,
        "description": "For regional NAT gateways only: Specifies which Availability Zones you want the NAT gateway to support and the Elastic IP addresses (EIPs) to use in each AZ. The regional NAT gateway uses these EIPs to handle outbound NAT traffic from their respective AZs. If not specified, the NAT gateway will automatically expand to new AZs and associate EIPs upon detection of an elastic network interface. If you specify this parameter, auto-expansion is disabled and you must manually manage AZ coverage.\n A regional NAT gateway is a single NAT Gateway that works across multiple availability zones (AZs) in your VPC, providing redundancy, scalability and availability across all the AZs in a Region.\n For more information, see [Regional NAT gateways for automatic multi-AZ expansion](https://docs.aws.amazon.com/vpc/latest/userguide/nat-gateways-regional.html) in the *Amazon VPC User Guide*.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "allocation_ids": {
              "computed": true,
              "description": "The allocation IDs of the Elastic IP addresses (EIPs) to be used for handling outbound NAT traffic in this specific Availability Zone.",
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            },
            "availability_zone": {
              "computed": true,
              "description": "For regional NAT gateways only: The Availability Zone where this specific NAT gateway configuration will be active. Each AZ in a regional NAT gateway has its own configuration to handle outbound NAT traffic from that AZ. \n A regional NAT gateway is a single NAT Gateway that works across multiple availability zones (AZs) in your VPC, providing redundancy, scalability and availability across all the AZs in a Region.",
              "description_kind": "plain",
              "type": "string"
            },
            "availability_zone_id": {
              "computed": true,
              "description": "For regional NAT gateways only: The ID of the Availability Zone where this specific NAT gateway configuration will be active. Each AZ in a regional NAT gateway has its own configuration to handle outbound NAT traffic from that AZ. Use this instead of AvailabilityZone for consistent identification of AZs across AWS Regions. \n A regional NAT gateway is a single NAT Gateway that works across multiple availability zones (AZs) in your VPC, providing redundancy, scalability and availability across all the AZs in a Region.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
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
      "route_table_id": {
        "computed": true,
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
        "description": "The ID of the VPC in which the NAT gateway is located.",
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
