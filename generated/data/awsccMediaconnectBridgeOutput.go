package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMediaconnectBridgeOutput = `{
  "block": {
    "attributes": {
      "bridge_arn": {
        "computed": true,
        "description": "The Amazon Resource Number (ARN) of the bridge.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The network output name.",
        "description_kind": "plain",
        "type": "string"
      },
      "network_output": {
        "computed": true,
        "description": "The output of the bridge.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ip_address": {
              "computed": true,
              "description": "The network output IP Address.",
              "description_kind": "plain",
              "type": "string"
            },
            "network_name": {
              "computed": true,
              "description": "The network output's gateway network name.",
              "description_kind": "plain",
              "type": "string"
            },
            "port": {
              "computed": true,
              "description": "The network output port.",
              "description_kind": "plain",
              "type": "number"
            },
            "protocol": {
              "computed": true,
              "description": "The network output protocol.",
              "description_kind": "plain",
              "type": "string"
            },
            "ttl": {
              "computed": true,
              "description": "The network output TTL.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::MediaConnect::BridgeOutput",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccMediaconnectBridgeOutputSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMediaconnectBridgeOutput), &result)
	return &result
}
