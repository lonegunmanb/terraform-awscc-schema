package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudfrontMonitoringSubscription = `{
  "block": {
    "attributes": {
      "distribution_id": {
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
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "realtime_metrics_subscription_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "realtime_metrics_subscription_status": {
                    "description_kind": "plain",
                    "required": true,
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
    "description": "Resource Type definition for AWS::CloudFront::MonitoringSubscription",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudfrontMonitoringSubscriptionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudfrontMonitoringSubscription), &result)
	return &result
}
