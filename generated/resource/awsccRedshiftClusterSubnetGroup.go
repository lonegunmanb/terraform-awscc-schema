package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRedshiftClusterSubnetGroup = `{
  "block": {
    "attributes": {
      "cluster_subnet_group_name": {
        "computed": true,
        "description": "This name must be unique for all subnet groups that are created by your AWS account. If costumer do not provide it, cloudformation will generate it. Must not be \"Default\". ",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "The description of the parameter group.",
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
      "subnet_ids": {
        "description": "The list of VPC subnet IDs",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "The list of tags for the cluster parameter group.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Specifies an Amazon Redshift subnet group.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRedshiftClusterSubnetGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRedshiftClusterSubnetGroup), &result)
	return &result
}
