package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSesMailManagerRelay = `{
  "block": {
    "attributes": {
      "authentication": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "no_authentication": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "secret_arn": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "relay_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "relay_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "relay_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "server_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "server_port": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
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
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Definition of AWS::SES::MailManagerRelay Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSesMailManagerRelaySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSesMailManagerRelay), &result)
	return &result
}
