package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDirectconnectConnection = `{
  "block": {
    "attributes": {
      "bandwidth": {
        "description": "The bandwidth of the connection.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "connection_arn": {
        "computed": true,
        "description": "The ARN of the connection.",
        "description_kind": "plain",
        "type": "string"
      },
      "connection_id": {
        "computed": true,
        "description": "The ID of the connection.",
        "description_kind": "plain",
        "type": "string"
      },
      "connection_name": {
        "description": "The name of the connection.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "connection_state": {
        "computed": true,
        "description": "The state of the connection.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "lag_id": {
        "computed": true,
        "description": "The ID or ARN of the LAG to associate the connection with.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "The location of the connection.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "provider_name": {
        "computed": true,
        "description": "The name of the service provider associated with the requested connection.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "request_mac_sec": {
        "computed": true,
        "description": "Indicates whether you want the connection to support MAC Security (MACsec).",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "tags": {
        "computed": true,
        "description": "The tags associated with the connection.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
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
    "description": "Resource Type definition for AWS::DirectConnect::Connection",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDirectconnectConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDirectconnectConnection), &result)
	return &result
}
