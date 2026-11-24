package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNetworkmanagerContextStore = `{
  "block": {
    "attributes": {
      "context_store_arn": {
        "computed": true,
        "description": "The ARN of the context store.",
        "description_kind": "plain",
        "type": "string"
      },
      "context_store_id": {
        "computed": true,
        "description": "The ID of the context store.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "Creation timestamp.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the context store.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "global_network_id": {
        "computed": true,
        "description": "The ID of the global network.",
        "description_kind": "plain",
        "optional": true,
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
        "description": "The state of the context store (PENDING / AVAILABLE / UPDATING / DELETING / DELETED).",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Creates a Context Store paired to a Global Network for sharing network context with firewall vendors.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccNetworkmanagerContextStoreSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworkmanagerContextStore), &result)
	return &result
}
