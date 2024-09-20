package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLookoutmetricsAlert = `{
  "block": {
    "attributes": {
      "action": {
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
                    "optional": true,
                    "type": "string"
                  },
                  "role_arn": {
                    "computed": true,
                    "description": "ARN of an IAM role that LookoutMetrics should assume to access the Lambda function.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
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
                    "optional": true,
                    "type": "string"
                  },
                  "sns_topic_arn": {
                    "computed": true,
                    "description": "ARN of an SNS topic to send alert notifications to.",
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
      },
      "alert_description": {
        "computed": true,
        "description": "A description for the alert.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "alert_name": {
        "computed": true,
        "description": "The name of the alert. If not provided, a name is generated automatically.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "alert_sensitivity_threshold": {
        "description": "A number between 0 and 100 (inclusive) that tunes the sensitivity of the alert.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "anomaly_detector_arn": {
        "description": "The Amazon resource name (ARN) of the Anomaly Detector to alert.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "ARN assigned to the alert.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::LookoutMetrics::Alert",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLookoutmetricsAlertSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLookoutmetricsAlert), &result)
	return &result
}
