package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEfsMountTarget = `{
  "block": {
    "attributes": {
      "file_system_id": {
        "description": "The ID of the file system for which to create the mount target.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "ip_address": {
        "computed": true,
        "description": "If the ` + "`" + `` + "`" + `IpAddressType` + "`" + `` + "`" + ` for the mount target is IPv4 ( ` + "`" + `` + "`" + `IPV4_ONLY` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `DUAL_STACK` + "`" + `` + "`" + `), then specify the IPv4 address to use. If you do not specify an ` + "`" + `` + "`" + `IpAddress` + "`" + `` + "`" + `, then Amazon EFS selects an unused IP address from the subnet specified for ` + "`" + `` + "`" + `SubnetId` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ip_address_type": {
        "computed": true,
        "description": "The IP address type for the mount target. The possible values are ` + "`" + `` + "`" + `IPV4_ONLY` + "`" + `` + "`" + ` (only IPv4 addresses), ` + "`" + `` + "`" + `IPV6_ONLY` + "`" + `` + "`" + ` (only IPv6 addresses), and ` + "`" + `` + "`" + `DUAL_STACK` + "`" + `` + "`" + ` (dual-stack, both IPv4 and IPv6 addresses). If you don?t specify an ` + "`" + `` + "`" + `IpAddressType` + "`" + `` + "`" + `, then ` + "`" + `` + "`" + `IPV4_ONLY` + "`" + `` + "`" + ` is used.\n  The ` + "`" + `` + "`" + `IPAddressType` + "`" + `` + "`" + ` must match the IP type of the subnet. Additionally, the ` + "`" + `` + "`" + `IPAddressType` + "`" + `` + "`" + ` parameter overrides the value set as the default IP address for the subnet in the VPC. For example, if the ` + "`" + `` + "`" + `IPAddressType` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `IPV4_ONLY` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `AssignIpv6AddressOnCreation` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `, then IPv4 is used for the mount target. For more information, see [Modify the IP addressing attributes of your subnet](https://docs.aws.amazon.com/vpc/latest/userguide/subnet-public-ip.html).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipv_6_address": {
        "computed": true,
        "description": "If the ` + "`" + `` + "`" + `IPAddressType` + "`" + `` + "`" + ` for the mount target is IPv6 (` + "`" + `` + "`" + `IPV6_ONLY` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `DUAL_STACK` + "`" + `` + "`" + `), then specify the IPv6 address to use. If you do not specify an ` + "`" + `` + "`" + `Ipv6Address` + "`" + `` + "`" + `, then Amazon EFS selects an unused IP address from the subnet specified for ` + "`" + `` + "`" + `SubnetId` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "mount_target_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "security_groups": {
        "description": "VPC security group IDs, of the form ` + "`" + `` + "`" + `sg-xxxxxxxx` + "`" + `` + "`" + `. These must be for the same VPC as the subnet specified. The maximum number of security groups depends on account quota. For more information, see [Amazon VPC Quotas](https://docs.aws.amazon.com/vpc/latest/userguide/amazon-vpc-limits.html) in the *Amazon VPC User Guide* (see the *Security Groups* table). If you don't specify a security group, then Amazon EFS uses the default security group for the subnet's VPC.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "subnet_id": {
        "description": "The ID of the subnet to add the mount target in. For One Zone file systems, use the subnet that is associated with the file system's Availability Zone. The subnet type must be the same type as the ` + "`" + `` + "`" + `IpAddressType` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "The ` + "`" + `` + "`" + `AWS::EFS::MountTarget` + "`" + `` + "`" + ` resource is an Amazon EFS resource that creates a mount target for an EFS file system. You can then mount the file system on Amazon EC2 instances or other resources by using the mount target.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEfsMountTargetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEfsMountTarget), &result)
	return &result
}
