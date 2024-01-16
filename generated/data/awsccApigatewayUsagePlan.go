package data

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
              "type": "string"
            },
            "stage": {
              "computed": true,
              "description": "API stage name of the associated API stage in a usage plan.",
              "description_kind": "plain",
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
                    "type": "number"
                  },
                  "rate_limit": {
                    "computed": true,
                    "description": "The API target request rate limit.",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "map"
              }
            }
          },
          "nesting_mode": "list"
        }
      },
      "description": {
        "computed": true,
        "description": "The description of a usage plan.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
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
              "type": "number"
            },
            "offset": {
              "computed": true,
              "description": "The number of requests subtracted from the given limit in the initial time period.",
              "description_kind": "plain",
              "type": "number"
            },
            "period": {
              "computed": true,
              "description": "The time period in which the limit applies. Valid values are \"DAY\", \"WEEK\" or \"MONTH\".",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description": "The collection of tags. Each tag element is associated with a given resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
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
              "type": "number"
            },
            "rate_limit": {
              "computed": true,
              "description": "The API target request rate limit.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "usage_plan_name": {
        "computed": true,
        "description": "The name of a usage plan.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ApiGateway::UsagePlan",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccApigatewayUsagePlanSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayUsagePlan), &result)
	return &result
}
