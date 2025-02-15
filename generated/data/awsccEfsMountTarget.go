package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEfsMountTarget = `{
  "block": {
    "attributes": {
      "file_system_id": {
        "computed": true,
        "description": "The ID of the file system for which to create the mount target.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ip_address": {
        "computed": true,
        "description": "Valid IPv4 address within the address range of the specified subnet.",
        "description_kind": "plain",
        "type": "string"
      },
      "mount_target_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "security_groups": {
        "computed": true,
        "description": "VPC security group IDs, of the form ` + "`" + `` + "`" + `sg-xxxxxxxx` + "`" + `` + "`" + `. These must be for the same VPC as the subnet specified. The maximum number of security groups depends on account quota. For more information, see [Amazon VPC Quotas](https://docs.aws.amazon.com/vpc/latest/userguide/amazon-vpc-limits.html) in the *Amazon VPC User Guide* (see the *Security Groups* table).",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "subnet_id": {
        "computed": true,
        "description": "The ID of the subnet to add the mount target in. For One Zone file systems, use the subnet that is associated with the file system's Availability Zone.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EFS::MountTarget",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEfsMountTargetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEfsMountTarget), &result)
	return &result
}
