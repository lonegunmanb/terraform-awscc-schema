package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2VpcBlockPublicAccessExclusion = `{
  "block": {
    "attributes": {
      "exclusion_id": {
        "computed": true,
        "description": "The ID of the exclusion",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "internet_gateway_exclusion_mode": {
        "description": "The desired Block Public Access Exclusion Mode for a specific VPC/Subnet.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "subnet_id": {
        "computed": true,
        "description": "The ID of the subnet. Required only if you don't specify VpcId",
        "description_kind": "plain",
        "optional": true,
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "vpc_id": {
        "computed": true,
        "description": "The ID of the vpc. Required only if you don't specify SubnetId.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::EC2::VPCBlockPublicAccessExclusion.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2VpcBlockPublicAccessExclusionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2VpcBlockPublicAccessExclusion), &result)
	return &result
}
