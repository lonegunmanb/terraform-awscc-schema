package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2SecurityGroupIngress = `{
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
        "description": "Updates the description of an ingress (inbound) security group rule. You can replace an existing description, or add a description to a rule that did not have one previously",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "from_port": {
        "computed": true,
        "description": "The start of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 type number. A value of -1 indicates all ICMP/ICMPv6 types. If you specify all ICMP/ICMPv6 types, you must specify all codes.\n\nUse this for ICMP and any protocol that uses ports.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "group_id": {
        "computed": true,
        "description": "The ID of the security group. You must specify either the security group ID or the security group name in the request. For security groups in a nondefault VPC, you must specify the security group ID.\n\nYou must specify the GroupName property or the GroupId property. For security groups that are in a VPC, you must use the GroupId property.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "group_name": {
        "computed": true,
        "description": "The name of the security group.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "ip_protocol": {
        "description": "The IP protocol name (tcp, udp, icmp, icmpv6) or number (see Protocol Numbers).\n\n[VPC only] Use -1 to specify all protocols. When authorizing security group rules, specifying -1 or a protocol number other than tcp, udp, icmp, or icmpv6 allows traffic on all ports, regardless of any port range you specify. For tcp, udp, and icmp, you must specify a port range. For icmpv6, the port range is optional; if you omit the port range, traffic for all types and codes is allowed.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "security_group_ingress_id": {
        "computed": true,
        "description": "The Security Group Rule Id",
        "description_kind": "plain",
        "type": "string"
      },
      "source_prefix_list_id": {
        "computed": true,
        "description": "[EC2-VPC only] The ID of a prefix list.\n\n",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_security_group_id": {
        "computed": true,
        "description": "The ID of the security group. You must specify either the security group ID or the security group name. For security groups in a nondefault VPC, you must specify the security group ID.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_security_group_name": {
        "computed": true,
        "description": "[EC2-Classic, default VPC] The name of the source security group.\n\nYou must specify the GroupName property or the GroupId property. For security groups that are in a VPC, you must use the GroupId property.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_security_group_owner_id": {
        "computed": true,
        "description": "[nondefault VPC] The AWS account ID that owns the source security group. You can't specify this property with an IP address range.\n\nIf you specify SourceSecurityGroupName or SourceSecurityGroupId and that security group is owned by a different account than the account creating the stack, you must specify the SourceSecurityGroupOwnerId; otherwise, this property is optional.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "to_port": {
        "computed": true,
        "description": "The end of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 code. A value of -1 indicates all ICMP/ICMPv6 codes for the specified ICMP type. If you specify all ICMP/ICMPv6 types, you must specify all codes.\n\nUse this for ICMP and any protocol that uses ports.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "description": "Resource Type definition for AWS::EC2::SecurityGroupIngress",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2SecurityGroupIngressSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2SecurityGroupIngress), &result)
	return &result
}
