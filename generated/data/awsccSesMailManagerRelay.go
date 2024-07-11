package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSesMailManagerRelay = `{
  "block": {
    "attributes": {
      "authentication": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "no_authentication": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "secret_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
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
        "type": "string"
      },
      "server_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "server_port": {
        "computed": true,
        "description_kind": "plain",
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
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::SES::MailManagerRelay",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSesMailManagerRelaySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSesMailManagerRelay), &result)
	return &result
}
