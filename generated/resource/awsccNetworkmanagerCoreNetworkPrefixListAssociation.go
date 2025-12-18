package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNetworkmanagerCoreNetworkPrefixListAssociation = `{
  "block": {
    "attributes": {
      "core_network_id": {
        "description": "The ID of the core network.",
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
      "prefix_list_alias": {
        "description": "The alias of the prefix list",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "prefix_list_arn": {
        "description": "The Amazon Resource Name (ARN) of the prefix list.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::NetworkManager::CoreNetworkPrefixListAssociation which associates a prefix list with a core network.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccNetworkmanagerCoreNetworkPrefixListAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworkmanagerCoreNetworkPrefixListAssociation), &result)
	return &result
}
