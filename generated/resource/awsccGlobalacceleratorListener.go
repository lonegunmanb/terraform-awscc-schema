package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGlobalacceleratorListener = `{
  "block": {
    "attributes": {
      "accelerator_arn": {
        "description": "The Amazon Resource Name (ARN) of the accelerator.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "client_affinity": {
        "computed": true,
        "description": "Client affinity lets you direct all requests from a user to the same endpoint.",
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
      "listener_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the listener.",
        "description_kind": "plain",
        "type": "string"
      },
      "port_ranges": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "from_port": {
              "description": "A network port number",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "to_port": {
              "description": "A network port number",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "protocol": {
        "computed": true,
        "description": "The protocol for the listener.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::GlobalAccelerator::Listener",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccGlobalacceleratorListenerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGlobalacceleratorListener), &result)
	return &result
}
