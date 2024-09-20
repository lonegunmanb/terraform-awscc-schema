package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCeAnomalySubscription = `{
  "block": {
    "attributes": {
      "account_id": {
        "computed": true,
        "description": "The accountId",
        "description_kind": "plain",
        "type": "string"
      },
      "frequency": {
        "description": "The frequency at which anomaly reports are sent over email. ",
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
      "monitor_arn_list": {
        "description": "A list of cost anomaly monitors.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "resource_tags": {
        "computed": true,
        "description": "Tags to assign to subscription.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name for the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "subscribers": {
        "description": "A list of subscriber",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "address": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "status": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "subscription_arn": {
        "computed": true,
        "description": "Subscription ARN",
        "description_kind": "plain",
        "type": "string"
      },
      "subscription_name": {
        "description": "The name of the subscription.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "threshold": {
        "computed": true,
        "description": "The dollar value that triggers a notification if the threshold is exceeded. ",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "threshold_expression": {
        "computed": true,
        "description": "An Expression object in JSON String format used to specify the anomalies that you want to generate alerts for.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "AWS Cost Anomaly Detection leverages advanced Machine Learning technologies to identify anomalous spend and root causes, so you can quickly take action. Create subscription to be notified",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCeAnomalySubscriptionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCeAnomalySubscription), &result)
	return &result
}
