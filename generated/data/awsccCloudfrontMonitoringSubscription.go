package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudfrontMonitoringSubscription = `{
  "block": {
    "attributes": {
      "distribution_id": {
        "computed": true,
        "description": "The ID of the distribution that you are enabling metrics for.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "monitoring_subscription": {
        "computed": true,
        "description": "A subscription configuration for additional CloudWatch metrics.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "realtime_metrics_subscription_config": {
              "computed": true,
              "description": "A subscription configuration for additional CloudWatch metrics.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "realtime_metrics_subscription_status": {
                    "computed": true,
                    "description": "A flag that indicates whether additional CloudWatch metrics are enabled for a given CloudFront distribution.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::CloudFront::MonitoringSubscription",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCloudfrontMonitoringSubscriptionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudfrontMonitoringSubscription), &result)
	return &result
}
