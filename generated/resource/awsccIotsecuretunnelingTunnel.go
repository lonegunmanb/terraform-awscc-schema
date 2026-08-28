package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotsecuretunnelingTunnel = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "A short text description of the tunnel.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_config": {
        "computed": true,
        "description": "The destination configuration for the tunnel.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "services": {
              "computed": true,
              "description": "A list of service names that identify the target application.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "thing_name": {
              "computed": true,
              "description": "The name of the IoT thing to which you want to connect.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the tunnel. Valid values are OPEN and CLOSED.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A collection of tag metadata.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "timeout_config": {
        "computed": true,
        "description": "Timeout configuration for the tunnel.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "max_lifetime_timeout_minutes": {
              "computed": true,
              "description": "The maximum amount of time (in minutes) a tunnel can remain open.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "tunnel_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the tunnel.",
        "description_kind": "plain",
        "type": "string"
      },
      "tunnel_id": {
        "computed": true,
        "description": "A unique alpha-numeric tunnel ID.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "A connection between a source computer and a destination device using AWS IoT Secure Tunneling.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIotsecuretunnelingTunnelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotsecuretunnelingTunnel), &result)
	return &result
}
