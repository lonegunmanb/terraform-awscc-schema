package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2SecurityGroupEgress = `{
  "block": {
    "attributes": {
      "cidr_ip": {
        "computed": true,
        "description": "The IPv4 ranges",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cidr_ipv_6": {
        "computed": true,
        "description": "[VPC only] The IPv6 ranges",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Resource Type definition for an egress (outbound) security group rule.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_prefix_list_id": {
        "computed": true,
        "description": "[EC2-VPC only] The ID of a prefix list.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_security_group_id": {
        "computed": true,
        "description": "You must specify a destination security group (DestinationPrefixListId or DestinationSecurityGroupId) or a CIDR range (CidrIp or CidrIpv6).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "from_port": {
        "computed": true,
        "description": "The start of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 type number. A value of -1 indicates all ICMP/ICMPv6 types. If you specify all ICMP/ICMPv6 types, you must specify all codes.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "group_id": {
        "description": "The ID of the security group. You must specify either the security group ID or the security group name in the request. For security groups in a nondefault VPC, you must specify the security group ID.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "The Security Group Rule Id",
        "description_kind": "plain",
        "type": "string"
      },
      "ip_protocol": {
        "description": "[VPC only] Use -1 to specify all protocols. When authorizing security group rules, specifying -1 or a protocol number other than tcp, udp, icmp, or icmpv6 allows traffic on all ports, regardless of any port range you specify. For tcp, udp, and icmp, you must specify a port range. For icmpv6, the port range is optional; if you omit the port range, traffic for all types and codes is allowed.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "to_port": {
        "computed": true,
        "description": "The end of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 code. A value of -1 indicates all ICMP/ICMPv6 codes. If you specify all ICMP/ICMPv6 types, you must specify all codes.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "description": "Resource Type definition for AWS::EC2::SecurityGroupEgress",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2SecurityGroupEgressSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2SecurityGroupEgress), &result)
	return &result
}
