package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2FcaPolicyAssociation = `{
  "block": {
    "attributes": {
      "associated_at": {
        "computed": true,
        "description": "The timestamp when the association was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "association_mode": {
        "computed": true,
        "description": "The mode of the association.",
        "description_kind": "plain",
        "type": "string"
      },
      "expected_effective_at": {
        "computed": true,
        "description": "The timestamp when the association is expected to take effect.",
        "description_kind": "plain",
        "type": "string"
      },
      "fca_policy_association_id": {
        "computed": true,
        "description": "The ID of the FCA policy association.",
        "description_kind": "plain",
        "type": "string"
      },
      "fca_policy_id": {
        "description": "The ID of the FCA policy to associate.",
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
      "managed_by": {
        "computed": true,
        "description": "Who manages the association: the account that created it, or a declarative policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "network_resource_id": {
        "description": "The ID of the network resource to associate the FCA policy with.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_resource_scope_type": {
        "computed": true,
        "description": "Associate by scope instead of directly. When set, NetworkResourceId is the ID of the scope (an account ID, a VPC ID, or a subnet ID) and the association covers every eligible network resource in that scope; when omitted, NetworkResourceId is a direct network resource (for example a NAT gateway or instance) and must match the policy's NetworkResourceType.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "owner_account": {
        "computed": true,
        "description": "The AWS account ID that owns the FCA policy association.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Associates a Flex Cost Allocation (FCA) policy with a network resource such as a NAT gateway, EC2 instance, or VPC peering connection.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2FcaPolicyAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2FcaPolicyAssociation), &result)
	return &result
}
