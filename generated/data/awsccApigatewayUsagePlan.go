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
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "api_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "stage": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "throttle": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "burst_limit": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "rate_limit": {
                    "computed": true,
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
        "description": "` + "`" + `` + "`" + `QuotaSettings` + "`" + `` + "`" + ` is a property of the [AWS::ApiGateway::UsagePlan](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplan.html) resource that specifies a target for the maximum number of requests users can make to your REST APIs.\n In some cases clients can exceed the targets that you set. Don?t rely on usage plans to control costs. Consider using [](https://docs.aws.amazon.com/cost-management/latest/userguide/budgets-managing-costs.html) to monitor costs and [](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) to manage API requests.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "limit": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "offset": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "period": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "A string you can use to assign a value. The combination of tag keys and values can help you organize and categorize your resources.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the specified tag key.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "throttle": {
        "computed": true,
        "description": "` + "`" + `` + "`" + `ThrottleSettings` + "`" + `` + "`" + ` is a property of the [AWS::ApiGateway::UsagePlan](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplan.html) resource that specifies the overall request rate (average requests per second) and burst capacity when users call your REST APIs.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "burst_limit": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "rate_limit": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "usage_plan_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "usage_plan_name": {
        "computed": true,
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
