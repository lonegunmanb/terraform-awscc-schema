package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayUsagePlan = `{
  "block": {
    "attributes": {
      "api_stages": {
        "computed": true,
        "description": "The associated API stages of a usage plan.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "api_id": {
              "computed": true,
              "description": "API Id of the associated API stage in a usage plan.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "stage": {
              "computed": true,
              "description": "API stage name of the associated API stage in a usage plan.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "throttle": {
              "computed": true,
              "description": "Map containing method level throttling information for API stage in a usage plan.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "burst_limit": {
                    "computed": true,
                    "description": "The API target request burst rate limit. This allows more requests through for a period of time than the target rate limit.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "rate_limit": {
                    "computed": true,
                    "description": "The API target request rate limit.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "map"
              },
              "optional": true
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "description": {
        "computed": true,
        "description": "The description of a usage plan.",
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
      "quota": {
        "computed": true,
        "description": "The target maximum number of permitted requests per a given unit time interval.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "limit": {
              "computed": true,
              "description": "The target maximum number of requests that can be made in a given time period.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "offset": {
              "computed": true,
              "description": "The number of requests subtracted from the given limit in the initial time period.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "period": {
              "computed": true,
              "description": "The time period in which the limit applies. Valid values are \"DAY\", \"WEEK\" or \"MONTH\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "tags": {
        "computed": true,
        "description": "The collection of tags. Each tag element is associated with a given resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "throttle": {
        "computed": true,
        "description": "A map containing method level throttling information for API stage in a usage plan.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "burst_limit": {
              "computed": true,
              "description": "The API target request burst rate limit. This allows more requests through for a period of time than the target rate limit.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "rate_limit": {
              "computed": true,
              "description": "The API target request rate limit.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "usage_plan_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "usage_plan_name": {
        "computed": true,
        "description": "The name of a usage plan.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "The ` + "`" + `` + "`" + `AWS::ApiGateway::UsagePlan` + "`" + `` + "`" + ` resource creates a usage plan for deployed APIs. A usage plan sets a target for the throttling and quota limits on individual client API keys. For more information, see [Creating and Using API Usage Plans in Amazon API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-api-usage-plans.html) in the *API Gateway Developer Guide*.\n In some cases clients can exceed the targets that you set. Don?t rely on usage plans to control costs. Consider using [](https://docs.aws.amazon.com/cost-management/latest/userguide/budgets-managing-costs.html) to monitor costs and [](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) to manage API requests.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApigatewayUsagePlanSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayUsagePlan), &result)
	return &result
}
