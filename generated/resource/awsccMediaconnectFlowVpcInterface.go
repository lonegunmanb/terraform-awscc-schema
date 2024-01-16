package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMediaconnectFlowVpcInterface = `{
  "block": {
    "attributes": {
      "flow_arn": {
        "description": "The Amazon Resource Name (ARN), a unique identifier for any AWS resource, of the flow.",
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
      "name": {
        "description": "Immutable and has to be a unique against other VpcInterfaces in this Flow.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_interface_ids": {
        "computed": true,
        "description": "IDs of the network interfaces created in customer's account by MediaConnect.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "role_arn": {
        "description": "Role Arn MediaConnect can assumes to create ENIs in customer's account.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "security_group_ids": {
        "description": "Security Group IDs to be used on ENI.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "subnet_id": {
        "description": "Subnet must be in the AZ of the Flow",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource schema for AWS::MediaConnect::FlowVpcInterface",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMediaconnectFlowVpcInterfaceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMediaconnectFlowVpcInterface), &result)
	return &result
}
