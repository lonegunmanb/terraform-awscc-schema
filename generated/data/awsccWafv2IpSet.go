package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWafv2IpSet = `{
  "block": {
    "attributes": {
      "addresses": {
        "computed": true,
        "description": "List of IPAddresses.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "arn": {
        "computed": true,
        "description": "ARN of the WAF entity.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Description of the entity.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ip_address_version": {
        "computed": true,
        "description": "Type of addresses in the IPSet, use IPV4 for IPV4 IP addresses, IPV6 for IPV6 address.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Name of the IPSet.",
        "description_kind": "plain",
        "type": "string"
      },
      "scope": {
        "computed": true,
        "description": "Use CLOUDFRONT for CloudFront IPSet, use REGIONAL for Application Load Balancer and API Gateway.",
        "description_kind": "plain",
        "type": "string"
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
    "description": "Data Source schema for AWS::WAFv2::IPSet",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccWafv2IpSetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWafv2IpSet), &result)
	return &result
}
