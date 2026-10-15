package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2ClientVpnAuthorizationRule = `{
  "block": {
    "attributes": {
      "access_group_id": {
        "computed": true,
        "description": "The ID of the group to grant access to, for example, the Active Directory group or identity provider (IdP) group. Required if ` + "`" + `AuthorizeAllGroups` + "`" + ` is ` + "`" + `false` + "`" + ` or not specified.",
        "description_kind": "plain",
        "type": "string"
      },
      "authorize_all_groups": {
        "computed": true,
        "description": "Indicates whether to grant access to all clients. Specify ` + "`" + `true` + "`" + ` to grant all clients who successfully establish a VPN connection access to the network. Must be set to ` + "`" + `true` + "`" + ` if ` + "`" + `AccessGroupId` + "`" + ` is not specified.",
        "description_kind": "plain",
        "type": "bool"
      },
      "client_vpn_authorization_rule_id": {
        "computed": true,
        "description": "The ID of the Client VPN Authorization Rule.",
        "description_kind": "plain",
        "type": "string"
      },
      "client_vpn_endpoint_id": {
        "computed": true,
        "description": "The ID of the Client VPN endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A brief description of the authorization rule.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "target_network_cidr": {
        "computed": true,
        "description": "The IPv4 address range, in CIDR notation, of the network for which access is being authorized.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::ClientVpnAuthorizationRule",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2ClientVpnAuthorizationRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2ClientVpnAuthorizationRule), &result)
	return &result
}
