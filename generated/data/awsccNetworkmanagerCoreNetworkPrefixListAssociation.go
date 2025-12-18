package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNetworkmanagerCoreNetworkPrefixListAssociation = `{
  "block": {
    "attributes": {
      "core_network_id": {
        "computed": true,
        "description": "The ID of the core network.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "prefix_list_alias": {
        "computed": true,
        "description": "The alias of the prefix list",
        "description_kind": "plain",
        "type": "string"
      },
      "prefix_list_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the prefix list.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::NetworkManager::CoreNetworkPrefixListAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccNetworkmanagerCoreNetworkPrefixListAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworkmanagerCoreNetworkPrefixListAssociation), &result)
	return &result
}
