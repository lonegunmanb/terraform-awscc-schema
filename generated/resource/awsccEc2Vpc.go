package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2Vpc = `{
  "block": {
    "attributes": {
      "cidr_block": {
        "computed": true,
        "description": "The IPv4 network range for the VPC, in CIDR notation. For example, ` + "`" + `` + "`" + `10.0.0.0/16` + "`" + `` + "`" + `. We modify the specified CIDR block to its canonical form; for example, if you specify ` + "`" + `` + "`" + `100.68.0.18/18` + "`" + `` + "`" + `, we modify it to ` + "`" + `` + "`" + `100.68.0.0/18` + "`" + `` + "`" + `.\n You must specify either` + "`" + `` + "`" + `CidrBlock` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `Ipv4IpamPoolId` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cidr_block_associations": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "default_network_acl": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "default_security_group": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enable_dns_hostnames": {
        "computed": true,
        "description": "Indicates whether the instances launched in the VPC get DNS hostnames. If enabled, instances in the VPC get DNS hostnames; otherwise, they do not. Disabled by default for nondefault VPCs. For more information, see [DNS attributes in your VPC](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-dns.html#vpc-dns-support).\n You can only enable DNS hostnames if you've enabled DNS support.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_dns_support": {
        "computed": true,
        "description": "Indicates whether the DNS resolution is supported for the VPC. If enabled, queries to the Amazon provided DNS server at the 169.254.169.253 IP address, or the reserved IP address at the base of the VPC network range \"plus two\" succeed. If disabled, the Amazon provided DNS service in the VPC that resolves public DNS hostnames to IP addresses is not enabled. Enabled by default. For more information, see [DNS attributes in your VPC](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-dns.html#vpc-dns-support).",
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
      "instance_tenancy": {
        "computed": true,
        "description": "The allowed tenancy of instances launched into the VPC.\n  +  ` + "`" + `` + "`" + `default` + "`" + `` + "`" + `: An instance launched into the VPC runs on shared hardware by default, unless you explicitly specify a different tenancy during instance launch.\n  +  ` + "`" + `` + "`" + `dedicated` + "`" + `` + "`" + `: An instance launched into the VPC runs on dedicated hardware by default, unless you explicitly specify a tenancy of ` + "`" + `` + "`" + `host` + "`" + `` + "`" + ` during instance launch. You cannot specify a tenancy of ` + "`" + `` + "`" + `default` + "`" + `` + "`" + ` during instance launch.\n  \n Updating ` + "`" + `` + "`" + `InstanceTenancy` + "`" + `` + "`" + ` requires no replacement only if you are updating its value from ` + "`" + `` + "`" + `dedicated` + "`" + `` + "`" + ` to ` + "`" + `` + "`" + `default` + "`" + `` + "`" + `. Updating ` + "`" + `` + "`" + `InstanceTenancy` + "`" + `` + "`" + ` from ` + "`" + `` + "`" + `default` + "`" + `` + "`" + ` to ` + "`" + `` + "`" + `dedicated` + "`" + `` + "`" + ` requires replacement.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipv_4_ipam_pool_id": {
        "computed": true,
        "description": "The ID of an IPv4 IPAM pool you want to use for allocating this VPC's CIDR. For more information, see [What is IPAM?](https://docs.aws.amazon.com//vpc/latest/ipam/what-is-it-ipam.html) in the *Amazon VPC IPAM User Guide*.\n You must specify either` + "`" + `` + "`" + `CidrBlock` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `Ipv4IpamPoolId` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipv_4_netmask_length": {
        "computed": true,
        "description": "The netmask length of the IPv4 CIDR you want to allocate to this VPC from an Amazon VPC IP Address Manager (IPAM) pool. For more information about IPAM, see [What is IPAM?](https://docs.aws.amazon.com//vpc/latest/ipam/what-is-it-ipam.html) in the *Amazon VPC IPAM User Guide*.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "ipv_6_cidr_blocks": {
        "computed": true,
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
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Specifies a virtual private cloud (VPC).\n You can optionally request an IPv6 CIDR block for the VPC. You can request an Amazon-provided IPv6 CIDR block from Amazon's pool of IPv6 addresses, or an IPv6 CIDR block from an IPv6 address pool that you provisioned through bring your own IP addresses (BYOIP).\n For more information, see [Virtual private clouds (VPC)](https://docs.aws.amazon.com/vpc/latest/userguide/configure-your-vpc.html) in the *Amazon VPC User Guide*.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2VpcSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2Vpc), &result)
	return &result
}
