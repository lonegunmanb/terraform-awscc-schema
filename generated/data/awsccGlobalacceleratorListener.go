package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGlobalacceleratorListener = `{
  "block": {
    "attributes": {
      "accelerator_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the accelerator.",
        "description_kind": "plain",
        "type": "string"
      },
      "client_affinity": {
        "computed": true,
        "description": "Client affinity lets you direct all requests from a user to the same endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "listener_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the listener.",
        "description_kind": "plain",
        "type": "string"
      },
      "port_ranges": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "from_port": {
              "computed": true,
              "description": "A network port number",
              "description_kind": "plain",
              "type": "number"
            },
            "to_port": {
              "computed": true,
              "description": "A network port number",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "list"
        }
      },
      "protocol": {
        "computed": true,
        "description": "The protocol for the listener.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::GlobalAccelerator::Listener",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccGlobalacceleratorListenerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGlobalacceleratorListener), &result)
	return &result
}
