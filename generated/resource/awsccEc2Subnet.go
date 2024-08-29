package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2Subnet = `{
  "block": {
    "attributes": {
      "assign_ipv_6_address_on_creation": {
        "computed": true,
        "description": "Indicates whether a network interface created in this subnet receives an IPv6 address. The default value is ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `.\n If you specify ` + "`" + `` + "`" + `AssignIpv6AddressOnCreation` + "`" + `` + "`" + `, you must also specify an IPv6 CIDR block.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "availability_zone": {
        "computed": true,
        "description": "The Availability Zone of the subnet.\n If you update this property, you must also update the ` + "`" + `` + "`" + `CidrBlock` + "`" + `` + "`" + ` property.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "availability_zone_id": {
        "computed": true,
        "description": "The AZ ID of the subnet.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cidr_block": {
        "computed": true,
        "description": "The IPv4 CIDR block assigned to the subnet.\n If you update this property, we create a new subnet, and then delete the existing one.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_dns_64": {
        "computed": true,
        "description": "Indicates whether DNS queries made to the Amazon-provided DNS Resolver in this subnet should return synthetic IPv6 addresses for IPv4-only destinations.\n  You must first configure a NAT gateway in a public subnet (separate from the subnet containing the IPv6-only workloads). For example, the subnet containing the NAT gateway should have a ` + "`" + `` + "`" + `0.0.0.0/0` + "`" + `` + "`" + ` route pointing to the internet gateway. For more information, see [Configure DNS64 and NAT64](https://docs.aws.amazon.com/vpc/latest/userguide/nat-gateway-nat64-dns64.html#nat-gateway-nat64-dns64-walkthrough) in the *User Guide*.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_lni_at_device_index": {
        "computed": true,
        "description": "Indicates the device position for local network interfaces in this subnet. For example, ` + "`" + `` + "`" + `1` + "`" + `` + "`" + ` indicates local network interfaces in this subnet are the secondary network interface (eth1).",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "ipv_4_ipam_pool_id": {
        "computed": true,
        "description": "An IPv4 IPAM pool ID for the subnet.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipv_4_netmask_length": {
        "computed": true,
        "description": "An IPv4 netmask length for the subnet.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "ipv_6_cidr_block": {
        "computed": true,
        "description": "The IPv6 CIDR block.\n If you specify ` + "`" + `` + "`" + `AssignIpv6AddressOnCreation` + "`" + `` + "`" + `, you must also specify an IPv6 CIDR block.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipv_6_cidr_blocks": {
        "computed": true,
        "description": "The IPv6 network ranges for the subnet, in CIDR notation.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "ipv_6_ipam_pool_id": {
        "computed": true,
        "description": "An IPv6 IPAM pool ID for the subnet.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipv_6_native": {
        "computed": true,
        "description": "Indicates whether this is an IPv6 only subnet. For more information, see [Subnet basics](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html#subnet-basics) in the *User Guide*.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "ipv_6_netmask_length": {
        "computed": true,
        "description": "An IPv6 netmask length for the subnet.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "map_public_ip_on_launch": {
        "computed": true,
        "description": "Indicates whether instances launched in this subnet receive a public IPv4 address. The default value is ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `.\n  AWS charges for all public IPv4 addresses, including public IPv4 addresses associated with running instances and Elastic IP addresses. For more information, see the *Public IPv4 Address* tab on the [VPC pricing page](https://docs.aws.amazon.com/vpc/pricing/).",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "network_acl_association_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "outpost_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the Outpost.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "private_dns_name_options_on_launch": {
        "computed": true,
        "description": "The hostname type for EC2 instances launched into this subnet and how DNS A and AAAA record queries to the instances should be handled. For more information, see [Amazon EC2 instance hostname types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-naming.html) in the *User Guide*.\n Available options:\n  +  EnableResourceNameDnsAAAARecord (true | false)\n  +  EnableResourceNameDnsARecord (true | false)\n  +  HostnameType (ip-name | resource-name)",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enable_resource_name_dns_a_record": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_resource_name_dns_aaaa_record": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "hostname_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "subnet_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Any tags assigned to the subnet.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The tag key.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The tag value.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "vpc_id": {
        "description": "The ID of the VPC the subnet is in.\n If you update this property, you must also update the ` + "`" + `` + "`" + `CidrBlock` + "`" + `` + "`" + ` property.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Specifies a subnet for the specified VPC.\n For an IPv4 only subnet, specify an IPv4 CIDR block. If the VPC has an IPv6 CIDR block, you can create an IPv6 only subnet or a dual stack subnet instead. For an IPv6 only subnet, specify an IPv6 CIDR block. For a dual stack subnet, specify both an IPv4 CIDR block and an IPv6 CIDR block.\n For more information, see [Subnets for your VPC](https://docs.aws.amazon.com/vpc/latest/userguide/configure-subnets.html) in the *Amazon VPC User Guide*.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2SubnetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2Subnet), &result)
	return &result
}
