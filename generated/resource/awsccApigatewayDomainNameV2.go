package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayDomainNameV2 = `{
  "block": {
    "attributes": {
      "certificate_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "domain_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "domain_name_arn": {
        "computed": true,
        "description": "The amazon resource name (ARN) of the domain name resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_name_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ip_address_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "types": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
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
      "policy": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "routing_mode": {
        "computed": true,
        "description": "The valid routing modes are [BASE_PATH_MAPPING_ONLY], [ROUTING_RULE_THEN_BASE_PATH_MAPPING] and [ROUTING_RULE_ONLY]. All other inputs are invalid.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_policy": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
    "description": "Resource Type definition for AWS::ApiGateway::DomainNameV2.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApigatewayDomainNameV2Schema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayDomainNameV2), &result)
	return &result
}
