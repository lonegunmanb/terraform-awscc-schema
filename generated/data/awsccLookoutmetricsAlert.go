package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLookoutmetricsAlert = `{
  "block": {
    "attributes": {
      "action": {
        "computed": true,
        "description": "The action to be taken by the alert when an anomaly is detected.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "lambda_configuration": {
              "computed": true,
              "description": "Configuration options for a Lambda alert action.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "lambda_arn": {
                    "computed": true,
                    "description": "ARN of a Lambda to send alert notifications to.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "role_arn": {
                    "computed": true,
                    "description": "ARN of an IAM role that LookoutMetrics should assume to access the Lambda function.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "sns_configuration": {
              "computed": true,
              "description": "Configuration options for an SNS alert action.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "role_arn": {
                    "computed": true,
                    "description": "ARN of an IAM role that LookoutMetrics should assume to access the SNS topic.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "sns_topic_arn": {
                    "computed": true,
                    "description": "ARN of an SNS topic to send alert notifications to.",
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
      },
      "alert_description": {
        "computed": true,
        "description": "A description for the alert.",
        "description_kind": "plain",
        "type": "string"
      },
      "alert_name": {
        "computed": true,
        "description": "The name of the alert. If not provided, a name is generated automatically.",
        "description_kind": "plain",
        "type": "string"
      },
      "alert_sensitivity_threshold": {
        "computed": true,
        "description": "A number between 0 and 100 (inclusive) that tunes the sensitivity of the alert.",
        "description_kind": "plain",
        "type": "number"
      },
      "anomaly_detector_arn": {
        "computed": true,
        "description": "The Amazon resource name (ARN) of the Anomaly Detector to alert.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "ARN assigned to the alert.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::LookoutMetrics::Alert",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLookoutmetricsAlertSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLookoutmetricsAlert), &result)
	return &result
}
