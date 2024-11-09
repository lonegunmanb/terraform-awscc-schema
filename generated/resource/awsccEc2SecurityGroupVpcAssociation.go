package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2SecurityGroupVpcAssociation = `{
  "block": {
    "attributes": {
      "group_id": {
        "description": "The group ID of the specified security group.",
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
      "state": {
        "computed": true,
        "description": "The state of the security group vpc association.",
        "description_kind": "plain",
        "type": "string"
      },
      "state_reason": {
        "computed": true,
        "description": "The reason for the state of the security group vpc association.",
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_id": {
        "description": "The ID of the VPC in the security group vpc association.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vpc_owner_id": {
        "computed": true,
        "description": "The owner of the VPC in the security group vpc association.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource type definition for the AWS::EC2::SecurityGroupVpcAssociation resource",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2SecurityGroupVpcAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2SecurityGroupVpcAssociation), &result)
	return &result
}
