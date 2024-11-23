package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudfrontMonitoringSubscription = `{
  "block": {
    "attributes": {
      "distribution_id": {
        "description": "The ID of the distribution that you are enabling metrics for.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "monitoring_subscription": {
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
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      }
    },
    "description": "A monitoring subscription. This structure contains information about whether additional CloudWatch metrics are enabled for a given CloudFront distribution.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudfrontMonitoringSubscriptionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudfrontMonitoringSubscription), &result)
	return &result
}
