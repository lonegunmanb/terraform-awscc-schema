package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockagentcoreGatewayRateLimit = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Optional human-readable description for this limit.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dimension_keys": {
        "description": "Ordered list of dimension names defining the scope of this limit.\nUnique per gateway — no two limits can share the same dimensionKeys.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "entries": {
        "description": "Rule entries mapping dimension values to rate configurations.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "connections": {
              "computed": true,
              "description": "Connection rate limits (per second only). Limited to 1 entry for now. — P2",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "period": {
                    "computed": true,
                    "description": "Time period for rate limiting",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "rate": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "dimensions": {
              "description": "Map of dimension name to dimension value for a rule entry",
              "description_kind": "plain",
              "required": true,
              "type": [
                "map",
                "string"
              ]
            },
            "requests": {
              "computed": true,
              "description": "Request rate limits (RPS or RPM). Limited to 1 entry for now.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "period": {
                    "computed": true,
                    "description": "Time period for rate limiting",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "rate": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "tokens": {
              "computed": true,
              "description": "Token rate limits (TPM). Limited to 1 entry for now. — P1",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "period": {
                    "computed": true,
                    "description": "Time period for rate limiting",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "rate": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "gateway_identifier": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "rate_limit_id": {
        "computed": true,
        "description": "Limit identifier. Optional on Create (system-generates if not provided by customer).\nAlways present in responses.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "Status of a gateway limit",
        "description_kind": "plain",
        "type": "string"
      },
      "updated_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Definition of AWS::BedrockAgentCore::GatewayRateLimit Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBedrockagentcoreGatewayRateLimitSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockagentcoreGatewayRateLimit), &result)
	return &result
}
