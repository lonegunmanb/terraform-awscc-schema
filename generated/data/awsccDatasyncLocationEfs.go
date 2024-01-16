package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatasyncLocationEfs = `{
  "block": {
    "attributes": {
      "access_point_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the Amazon EFS Access point that DataSync uses when accessing the EFS file system.",
        "description_kind": "plain",
        "type": "string"
      },
      "ec_2_config": {
        "computed": true,
        "description": "The subnet and security group that DataSync uses to access target EFS file system.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "security_group_arns": {
              "computed": true,
              "description": "The Amazon Resource Names (ARNs) of the security groups that are configured for the Amazon EC2 resource.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "subnet_arn": {
              "computed": true,
              "description": "The ARN of the subnet that DataSync uses to access the target EFS file system.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "efs_filesystem_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the Amazon EFS file system.",
        "description_kind": "plain",
        "type": "string"
      },
      "file_system_access_role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the AWS IAM role that the DataSync will assume when mounting the EFS file system.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "in_transit_encryption": {
        "computed": true,
        "description": "Protocol that is used for encrypting the traffic exchanged between the DataSync Agent and the EFS file system.",
        "description_kind": "plain",
        "type": "string"
      },
      "location_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the Amazon EFS file system location that is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "location_uri": {
        "computed": true,
        "description": "The URL of the EFS location that was described.",
        "description_kind": "plain",
        "type": "string"
      },
      "subdirectory": {
        "computed": true,
        "description": "A subdirectory in the location's path. This subdirectory in the EFS file system is used to read data from the EFS source location or write data to the EFS destination.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key for an AWS resource tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for an AWS resource tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::DataSync::LocationEFS",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDatasyncLocationEfsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatasyncLocationEfs), &result)
	return &result
}
