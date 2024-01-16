package data

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
        "computed": true,
        "description": "The frequency at which anomaly reports are sent over email. ",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "monitor_arn_list": {
        "computed": true,
        "description": "A list of cost anomaly monitors.",
        "description_kind": "plain",
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
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "subscribers": {
        "computed": true,
        "description": "A list of subscriber",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "address": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "status": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "subscription_arn": {
        "computed": true,
        "description": "Subscription ARN",
        "description_kind": "plain",
        "type": "string"
      },
      "subscription_name": {
        "computed": true,
        "description": "The name of the subscription.",
        "description_kind": "plain",
        "type": "string"
      },
      "threshold": {
        "computed": true,
        "description": "The dollar value that triggers a notification if the threshold is exceeded. ",
        "description_kind": "plain",
        "type": "number"
      },
      "threshold_expression": {
        "computed": true,
        "description": "An Expression object in JSON String format used to specify the anomalies that you want to generate alerts for.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::CE::AnomalySubscription",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCeAnomalySubscriptionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCeAnomalySubscription), &result)
	return &result
}
