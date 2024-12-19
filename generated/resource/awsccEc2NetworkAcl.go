package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2NetworkAcl = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "network_acl_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags for the network ACL.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value.",
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
        "description": "The ID of the VPC for the network ACL.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Specifies a network ACL for your VPC.\n To add a network ACL entry, see [AWS::EC2::NetworkAclEntry](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkaclentry.html).",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2NetworkAclSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2NetworkAcl), &result)
	return &result
}
