package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSesReceiptFilter = `{
  "block": {
    "attributes": {
      "filter": {
        "description": "A structure that describes the IP address filter to create, which consists of a name, an IP address range, and whether to allow or block mail from it.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ip_filter": {
              "description": "A structure that provides the IP addresses to block or allow, and whether to block or allow incoming mail from them.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cidr": {
                    "description": "A single IP address or a range of IP addresses to block or allow, specified in CIDR notation.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "policy": {
                    "description": "Indicates whether to block or allow incoming mail from the specified IP addresses.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            },
            "name": {
              "computed": true,
              "description": "The name of the IP address filter.",
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
      "receipt_filter_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::SES::ReceiptFilter",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSesReceiptFilterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSesReceiptFilter), &result)
	return &result
}
