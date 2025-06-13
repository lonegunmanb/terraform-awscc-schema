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
        "description": "Valid IPv4 address within the address range of the specified subnet.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ip_address_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipv_6_address": {
        "computed": true,
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
        "description": "VPC security group IDs, of the form ` + "`" + `` + "`" + `sg-xxxxxxxx` + "`" + `` + "`" + `. These must be for the same VPC as the subnet specified. The maximum number of security groups depends on account quota. For more information, see [Amazon VPC Quotas](https://docs.aws.amazon.com/vpc/latest/userguide/amazon-vpc-limits.html) in the *Amazon VPC User Guide* (see the *Security Groups* table).",
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "subnet_id": {
        "description": "The ID of the subnet to add the mount target in. For One Zone file systems, use the subnet that is associated with the file system's Availability Zone.",
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
