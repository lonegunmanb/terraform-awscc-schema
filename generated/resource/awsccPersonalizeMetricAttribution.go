package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccPersonalizeMetricAttribution = `{
  "block": {
    "attributes": {
      "dataset_group_arn": {
        "description": "The ARN of the destination dataset group.",
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
      "metric_attribution_arn": {
        "computed": true,
        "description": "The ARN of the metric attribution.",
        "description_kind": "plain",
        "type": "string"
      },
      "metrics": {
        "description": "A list of metric attributes for the metric attribution.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "event_type": {
              "description": "The metric's event type.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "expression": {
              "description": "The attribute's expression.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "metric_name": {
              "description": "The metric's name.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "metrics_output_config": {
        "description": "The output configuration details for the metric attribution.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "role_arn": {
              "description": "The ARN of the IAM role for the metric attribution.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "s3_data_destination": {
              "computed": true,
              "description": "The configuration details of an Amazon S3 output bucket.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "kms_key_arn": {
                    "computed": true,
                    "description": "The ARN of the KMS key.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "path": {
                    "computed": true,
                    "description": "The file path of the Amazon S3 bucket.",
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
      "name": {
        "description": "The name of the metric attribution.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the metric attribution.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Creates a metric attribution for reporting on recommendation impact in Amazon Personalize.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccPersonalizeMetricAttributionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPersonalizeMetricAttribution), &result)
	return &result
}
