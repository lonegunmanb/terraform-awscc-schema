package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSesReceiptFilter = `{
  "block": {
    "attributes": {
      "filter": {
        "computed": true,
        "description": "A structure that describes the IP address filter to create, which consists of a name, an IP address range, and whether to allow or block mail from it.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ip_filter": {
              "computed": true,
              "description": "A structure that provides the IP addresses to block or allow, and whether to block or allow incoming mail from them.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cidr": {
                    "computed": true,
                    "description": "A single IP address or a range of IP addresses to block or allow, specified in CIDR notation.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "policy": {
                    "computed": true,
                    "description": "Indicates whether to block or allow incoming mail from the specified IP addresses.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "name": {
              "computed": true,
              "description": "The name of the IP address filter.",
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
      "receipt_filter_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SES::ReceiptFilter",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSesReceiptFilterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSesReceiptFilter), &result)
	return &result
}
