package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccChimeVoiceConnector = `{
  "block": {
    "attributes": {
      "aws_region": {
        "computed": true,
        "description": "The AWS Region in which the Voice Connector is created.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "created_timestamp": {
        "computed": true,
        "description": "The Voice Connector creation timestamp, in ISO 8601 format.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the Voice Connector.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_type": {
        "computed": true,
        "description": "The type of network for the Voice Connector.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "outbound_host_name": {
        "computed": true,
        "description": "The outbound host name for the Voice Connector.",
        "description_kind": "plain",
        "type": "string"
      },
      "require_encryption": {
        "description": "Enables or disables encryption for the Voice Connector.",
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "tags": {
        "computed": true,
        "description": "The tags assigned to the Voice Connector.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "updated_timestamp": {
        "computed": true,
        "description": "The Voice Connector updated timestamp, in ISO 8601 format.",
        "description_kind": "plain",
        "type": "string"
      },
      "voice_connector_arn": {
        "computed": true,
        "description": "The ARN of the Voice Connector.",
        "description_kind": "plain",
        "type": "string"
      },
      "voice_connector_id": {
        "computed": true,
        "description": "The Voice Connector ID.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "An Amazon Chime SDK Voice Connector configuration, including outbound host name and encryption settings.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccChimeVoiceConnectorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccChimeVoiceConnector), &result)
	return &result
}
