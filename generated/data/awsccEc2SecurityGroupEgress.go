package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2SecurityGroupEgress = `{
  "block": {
    "attributes": {
      "cidr_ip": {
        "computed": true,
        "description": "The IPv4 address range, in CIDR format.\n You must specify a destination security group (` + "`" + `` + "`" + `DestinationPrefixListId` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `DestinationSecurityGroupId` + "`" + `` + "`" + `) or a CIDR range (` + "`" + `` + "`" + `CidrIp` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `CidrIpv6` + "`" + `` + "`" + `).\n For examples of rules that you can add to security groups for specific access scenarios, see [Security group rules for different use cases](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/security-group-rules-reference.html) in the *User Guide*.",
        "description_kind": "plain",
        "type": "string"
      },
      "cidr_ipv_6": {
        "computed": true,
        "description": "The IPv6 address range, in CIDR format.\n You must specify a destination security group (` + "`" + `` + "`" + `DestinationPrefixListId` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `DestinationSecurityGroupId` + "`" + `` + "`" + `) or a CIDR range (` + "`" + `` + "`" + `CidrIp` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `CidrIpv6` + "`" + `` + "`" + `).\n For examples of rules that you can add to security groups for specific access scenarios, see [Security group rules for different use cases](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/security-group-rules-reference.html) in the *User Guide*.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of an egress (outbound) security group rule.\n Constraints: Up to 255 characters in length. Allowed characters are a-z, A-Z, 0-9, spaces, and ._-:/()#,@[]+=;{}!$*",
        "description_kind": "plain",
        "type": "string"
      },
      "destination_prefix_list_id": {
        "computed": true,
        "description": "The prefix list IDs for an AWS service. This is the AWS service that you want to access through a VPC endpoint from instances associated with the security group.\n You must specify a destination security group (` + "`" + `` + "`" + `DestinationPrefixListId` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `DestinationSecurityGroupId` + "`" + `` + "`" + `) or a CIDR range (` + "`" + `` + "`" + `CidrIp` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `CidrIpv6` + "`" + `` + "`" + `).",
        "description_kind": "plain",
        "type": "string"
      },
      "destination_security_group_id": {
        "computed": true,
        "description": "The ID of the security group.\n You must specify a destination security group (` + "`" + `` + "`" + `DestinationPrefixListId` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `DestinationSecurityGroupId` + "`" + `` + "`" + `) or a CIDR range (` + "`" + `` + "`" + `CidrIp` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `CidrIpv6` + "`" + `` + "`" + `).",
        "description_kind": "plain",
        "type": "string"
      },
      "from_port": {
        "computed": true,
        "description": "If the protocol is TCP or UDP, this is the start of the port range. If the protocol is ICMP or ICMPv6, this is the ICMP type or -1 (all ICMP types).",
        "description_kind": "plain",
        "type": "number"
      },
      "group_id": {
        "computed": true,
        "description": "The ID of the security group. You must specify either the security group ID or the security group name in the request. For security groups in a nondefault VPC, you must specify the security group ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ip_protocol": {
        "computed": true,
        "description": "The IP protocol name (` + "`" + `` + "`" + `tcp` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `udp` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `icmp` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `icmpv6` + "`" + `` + "`" + `) or number (see [Protocol Numbers](https://docs.aws.amazon.com/http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml)).\n Use ` + "`" + `` + "`" + `-1` + "`" + `` + "`" + ` to specify all protocols. When authorizing security group rules, specifying ` + "`" + `` + "`" + `-1` + "`" + `` + "`" + ` or a protocol number other than ` + "`" + `` + "`" + `tcp` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `udp` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `icmp` + "`" + `` + "`" + `, or ` + "`" + `` + "`" + `icmpv6` + "`" + `` + "`" + ` allows traffic on all ports, regardless of any port range you specify. For ` + "`" + `` + "`" + `tcp` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `udp` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `icmp` + "`" + `` + "`" + `, you must specify a port range. For ` + "`" + `` + "`" + `icmpv6` + "`" + `` + "`" + `, the port range is optional; if you omit the port range, traffic for all types and codes is allowed.",
        "description_kind": "plain",
        "type": "string"
      },
      "security_group_egress_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "to_port": {
        "computed": true,
        "description": "If the protocol is TCP or UDP, this is the end of the port range. If the protocol is ICMP or ICMPv6, this is the ICMP code or -1 (all ICMP codes). If the start port is -1 (all ICMP types), then the end port must be -1 (all ICMP codes).",
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description": "Data Source schema for AWS::EC2::SecurityGroupEgress",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2SecurityGroupEgressSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2SecurityGroupEgress), &result)
	return &result
}
