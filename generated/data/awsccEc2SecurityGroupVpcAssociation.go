package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2SecurityGroupVpcAssociation = `{
  "block": {
    "attributes": {
      "group_id": {
        "computed": true,
        "description": "The group ID of the specified security group.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
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
        "computed": true,
        "description": "The ID of the VPC in the security group vpc association.",
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_owner_id": {
        "computed": true,
        "description": "The owner of the VPC in the security group vpc association.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::SecurityGroupVpcAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2SecurityGroupVpcAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2SecurityGroupVpcAssociation), &result)
	return &result
}
