package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNetworkmanagerContextStoreAssociation = `{
  "block": {
    "attributes": {
      "context_store_id": {
        "description": "The ID of the context store.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "Creation timestamp (ISO-8601).",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_arn": {
        "description": "The ARN of the Transit Gateway to associate.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The state of the association (ASSOCIATING, ASSOCIATED, DISASSOCIATING, FAILED).",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Associates a Transit Gateway with a Context Store for sharing network context.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccNetworkmanagerContextStoreAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworkmanagerContextStoreAssociation), &result)
	return &result
}
