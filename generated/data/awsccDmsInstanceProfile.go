package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDmsInstanceProfile = `{
  "block": {
    "attributes": {
      "availability_zone": {
        "computed": true,
        "description": "The property describes an availability zone of the instance profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The optional description of the instance profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_profile_arn": {
        "computed": true,
        "description": "The property describes an ARN of the instance profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_profile_creation_time": {
        "computed": true,
        "description": "The property describes a creating time of the instance profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_profile_identifier": {
        "computed": true,
        "description": "The property describes an identifier for the instance profile. It is used for describing/deleting/modifying. Can be name/arn",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_profile_name": {
        "computed": true,
        "description": "The property describes a name for the instance profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_arn": {
        "computed": true,
        "description": "The property describes kms key arn for the instance profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "network_type": {
        "computed": true,
        "description": "The property describes a network type for the instance profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "publicly_accessible": {
        "computed": true,
        "description": "The property describes the publicly accessible of the instance profile",
        "description_kind": "plain",
        "type": "bool"
      },
      "subnet_group_identifier": {
        "computed": true,
        "description": "The property describes a subnet group identifier for the instance profile.",
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
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "vpc_security_groups": {
        "computed": true,
        "description": "The property describes vps security groups for the instance profile.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::DMS::InstanceProfile",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDmsInstanceProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDmsInstanceProfile), &result)
	return &result
}
