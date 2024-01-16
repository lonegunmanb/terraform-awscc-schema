package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWafv2IpSet = `{
  "block": {
    "attributes": {
      "addresses": {
        "description": "List of IPAddresses.",
        "description_kind": "plain",
        "required": true,
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
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Id of the IPSet",
        "description_kind": "plain",
        "type": "string"
      },
      "ip_address_version": {
        "description": "Type of addresses in the IPSet, use IPV4 for IPV4 IP addresses, IPV6 for IPV6 address.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Name of the IPSet.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scope": {
        "description": "Use CLOUDFRONT for CloudFront IPSet, use REGIONAL for Application Load Balancer and API Gateway.",
        "description_kind": "plain",
        "required": true,
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
    "description": "Contains a list of IP addresses. This can be either IPV4 or IPV6. The list will be mutually",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccWafv2IpSetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWafv2IpSet), &result)
	return &result
}
