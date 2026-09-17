package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccFmsApplicationsList = `{
  "block": {
    "attributes": {
      "apps_list": {
        "computed": true,
        "description": "An array of applications in the Firewall Manager applications list.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "app_name": {
              "computed": true,
              "description": "The application's name.",
              "description_kind": "plain",
              "type": "string"
            },
            "port": {
              "computed": true,
              "description": "The application's port number, for example 80.",
              "description_kind": "plain",
              "type": "number"
            },
            "protocol": {
              "computed": true,
              "description": "The IP protocol name or number. The name can be one of tcp, udp, or icmp. For information on possible numbers, see Protocol Numbers (https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml).",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the applications list.",
        "description_kind": "plain",
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The time that the Firewall Manager applications list was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_update_time": {
        "computed": true,
        "description": "The time that the Firewall Manager applications list was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "list_id": {
        "computed": true,
        "description": "The ID of the Firewall Manager applications list.",
        "description_kind": "plain",
        "type": "string"
      },
      "list_name": {
        "computed": true,
        "description": "The name of the Firewall Manager applications list.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to the applications list.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::FMS::ApplicationsList",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccFmsApplicationsListSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccFmsApplicationsList), &result)
	return &result
}
