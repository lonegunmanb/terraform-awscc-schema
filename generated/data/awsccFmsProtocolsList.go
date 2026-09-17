package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccFmsProtocolsList = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the protocols list.",
        "description_kind": "plain",
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The time that the Firewall Manager protocols list was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_update_time": {
        "computed": true,
        "description": "The time that the Firewall Manager protocols list was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "list_id": {
        "computed": true,
        "description": "The ID of the Firewall Manager protocols list.",
        "description_kind": "plain",
        "type": "string"
      },
      "list_name": {
        "computed": true,
        "description": "The name of the Firewall Manager protocols list.",
        "description_kind": "plain",
        "type": "string"
      },
      "protocols_list": {
        "computed": true,
        "description": "An array of protocols in the Firewall Manager protocols list.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to the protocols list.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::FMS::ProtocolsList",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccFmsProtocolsListSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccFmsProtocolsList), &result)
	return &result
}
