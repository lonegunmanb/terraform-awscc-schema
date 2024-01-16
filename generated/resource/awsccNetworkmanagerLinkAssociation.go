package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNetworkmanagerLinkAssociation = `{
  "block": {
    "attributes": {
      "device_id": {
        "description": "The ID of the device",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "global_network_id": {
        "description": "The ID of the global network.",
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
      "link_id": {
        "description": "The ID of the link",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "The AWS::NetworkManager::LinkAssociation type associates a link to a device. The device and link must be in the same global network and the same site.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccNetworkmanagerLinkAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworkmanagerLinkAssociation), &result)
	return &result
}
