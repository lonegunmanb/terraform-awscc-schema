package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNetworkmanagerConnection = `{
  "block": {
    "attributes": {
      "connected_device_id": {
        "computed": true,
        "description": "The ID of the second device in the connection.",
        "description_kind": "plain",
        "type": "string"
      },
      "connected_link_id": {
        "computed": true,
        "description": "The ID of the link for the second device.",
        "description_kind": "plain",
        "type": "string"
      },
      "connection_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the connection.",
        "description_kind": "plain",
        "type": "string"
      },
      "connection_id": {
        "computed": true,
        "description": "The ID of the connection.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The date and time that the connection was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the connection.",
        "description_kind": "plain",
        "type": "string"
      },
      "device_id": {
        "computed": true,
        "description": "The ID of the first device in the connection.",
        "description_kind": "plain",
        "type": "string"
      },
      "global_network_id": {
        "computed": true,
        "description": "The ID of the global network that the connection belongs to.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "link_id": {
        "computed": true,
        "description": "The ID of the link for the first device.",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The state of the connection.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags for the connection.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key. Printable ASCII with at least one non-space character and no leading or trailing space, at most 128 characters, and cannot be prefixed with aws:.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value. Printable ASCII with at least one non-space character and no leading or trailing space, at most 256 characters, and cannot be prefixed with aws:.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::NetworkManager::Connection",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccNetworkmanagerConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworkmanagerConnection), &result)
	return &result
}
