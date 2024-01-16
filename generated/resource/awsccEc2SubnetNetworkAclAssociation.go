package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2SubnetNetworkAclAssociation = `{
  "block": {
    "attributes": {
      "association_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "network_acl_id": {
        "description": "The ID of the network ACL",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "subnet_id": {
        "description": "The ID of the subnet",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::EC2::SubnetNetworkAclAssociation",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2SubnetNetworkAclAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2SubnetNetworkAclAssociation), &result)
	return &result
}
