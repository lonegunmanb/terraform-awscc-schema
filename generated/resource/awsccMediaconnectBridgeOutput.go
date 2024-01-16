package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMediaconnectBridgeOutput = `{
  "block": {
    "attributes": {
      "bridge_arn": {
        "description": "The Amazon Resource Number (ARN) of the bridge.",
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
      "name": {
        "description": "The network output name.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_output": {
        "description": "The output of the bridge.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ip_address": {
              "description": "The network output IP Address.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "network_name": {
              "description": "The network output's gateway network name.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "port": {
              "description": "The network output port.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "protocol": {
              "description": "The network output protocol.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "ttl": {
              "description": "The network output TTL.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      }
    },
    "description": "Resource schema for AWS::MediaConnect::BridgeOutput",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMediaconnectBridgeOutputSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMediaconnectBridgeOutput), &result)
	return &result
}
