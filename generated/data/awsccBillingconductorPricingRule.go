package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBillingconductorPricingRule = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Pricing rule ARN",
        "description_kind": "plain",
        "type": "string"
      },
      "associated_pricing_plan_count": {
        "computed": true,
        "description": "The number of pricing plans associated with pricing rule",
        "description_kind": "plain",
        "type": "number"
      },
      "billing_entity": {
        "computed": true,
        "description": "The seller of services provided by AWS, their affiliates, or third-party providers selling services via AWS Marketplaces. Supported billing entities are AWS, AWS Marketplace, and AISPL.",
        "description_kind": "plain",
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description": "Creation timestamp in UNIX epoch time format",
        "description_kind": "plain",
        "type": "number"
      },
      "description": {
        "computed": true,
        "description": "Pricing rule description",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description": "Latest modified timestamp in UNIX epoch time format",
        "description_kind": "plain",
        "type": "number"
      },
      "modifier_percentage": {
        "computed": true,
        "description": "Pricing rule modifier percentage",
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "computed": true,
        "description": "Pricing rule name",
        "description_kind": "plain",
        "type": "string"
      },
      "operation": {
        "computed": true,
        "description": "The Operation which a SKU pricing rule is modifying",
        "description_kind": "plain",
        "type": "string"
      },
      "scope": {
        "computed": true,
        "description": "A term used to categorize the granularity of a Pricing Rule.",
        "description_kind": "plain",
        "type": "string"
      },
      "service": {
        "computed": true,
        "description": "The service which a pricing rule is applied on",
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
          "nesting_mode": "set"
        }
      },
      "tiering": {
        "computed": true,
        "description": "The set of tiering configurations for the pricing rule.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "free_tier": {
              "computed": true,
              "description": "The possible customizable free tier configurations.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "activated": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "type": {
        "computed": true,
        "description": "One of MARKUP, DISCOUNT or TIERING that describes the behaviour of the pricing rule.",
        "description_kind": "plain",
        "type": "string"
      },
      "usage_type": {
        "computed": true,
        "description": "The UsageType which a SKU pricing rule is modifying",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::BillingConductor::PricingRule",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBillingconductorPricingRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBillingconductorPricingRule), &result)
	return &result
}
