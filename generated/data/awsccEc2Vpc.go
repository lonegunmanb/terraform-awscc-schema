package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2Vpc = `{
  "block": {
    "attributes": {
      "cidr_block": {
        "computed": true,
        "description": "The primary IPv4 CIDR block for the VPC.",
        "description_kind": "plain",
        "type": "string"
      },
      "cidr_block_associations": {
        "computed": true,
        "description": "A list of IPv4 CIDR block association IDs for the VPC.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "default_network_acl": {
        "computed": true,
        "description": "The default network ACL ID that is associated with the VPC.",
        "description_kind": "plain",
        "type": "string"
      },
      "default_security_group": {
        "computed": true,
        "description": "The default security group ID that is associated with the VPC.",
        "description_kind": "plain",
        "type": "string"
      },
      "enable_dns_hostnames": {
        "computed": true,
        "description": "Indicates whether the instances launched in the VPC get DNS hostnames. If enabled, instances in the VPC get DNS hostnames; otherwise, they do not. Disabled by default for nondefault VPCs.",
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_dns_support": {
        "computed": true,
        "description": "Indicates whether the DNS resolution is supported for the VPC. If enabled, queries to the Amazon provided DNS server at the 169.254.169.253 IP address, or the reserved IP address at the base of the VPC network range \"plus two\" succeed. If disabled, the Amazon provided DNS service in the VPC that resolves public DNS hostnames to IP addresses is not enabled. Enabled by default.",
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_tenancy": {
        "computed": true,
        "description": "The allowed tenancy of instances launched into the VPC.\n\n\"default\": An instance launched into the VPC runs on shared hardware by default, unless you explicitly specify a different tenancy during instance launch.\n\n\"dedicated\": An instance launched into the VPC is a Dedicated Instance by default, unless you explicitly specify a tenancy of host during instance launch. You cannot specify a tenancy of default during instance launch.\n\nUpdating InstanceTenancy requires no replacement only if you are updating its value from \"dedicated\" to \"default\". Updating InstanceTenancy from \"default\" to \"dedicated\" requires replacement.",
        "description_kind": "plain",
        "type": "string"
      },
      "ipv_4_ipam_pool_id": {
        "computed": true,
        "description": "The ID of an IPv4 IPAM pool you want to use for allocating this VPC's CIDR",
        "description_kind": "plain",
        "type": "string"
      },
      "ipv_4_netmask_length": {
        "computed": true,
        "description": "The netmask length of the IPv4 CIDR you want to allocate to this VPC from an Amazon VPC IP Address Manager (IPAM) pool",
        "description_kind": "plain",
        "type": "number"
      },
      "ipv_6_cidr_blocks": {
        "computed": true,
        "description": "A list of IPv6 CIDR blocks that are associated with the VPC.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "The tags for the VPC.",
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
      },
      "vpc_id": {
        "computed": true,
        "description": "The Id for the model.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::VPC",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2VpcSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2Vpc), &result)
	return &result
}
