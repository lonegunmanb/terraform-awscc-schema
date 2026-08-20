package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWellarchitectedWorkload = `{
  "block": {
    "attributes": {
      "account_ids": {
        "computed": true,
        "description": "The list of Amazon Web Services account IDs associated with the workload.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "architectural_design": {
        "computed": true,
        "description": "The URL of the architectural design for the workload.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "aws_regions": {
        "computed": true,
        "description": "The list of Amazon Web Services Regions associated with the workload.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "description": {
        "description": "The description for the workload.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "discovery_config": {
        "computed": true,
        "description": "Discovery configuration associated to the workload.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "trusted_advisor_integration_status": {
              "computed": true,
              "description": "Discovery integration status in respect to Trusted Advisor for the workload.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "workload_resource_definition": {
              "computed": true,
              "description": "The mode to use for identifying resources associated with the workload.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "environment": {
        "description": "The environment for the workload.",
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
      "improvement_status": {
        "computed": true,
        "description": "The improvement status for a workload.",
        "description_kind": "plain",
        "type": "string"
      },
      "industry": {
        "computed": true,
        "description": "The industry for the workload.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "industry_type": {
        "computed": true,
        "description": "The industry type for the workload.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "lenses": {
        "description": "The list of lenses associated with the workload.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "non_aws_regions": {
        "computed": true,
        "description": "The list of non-Amazon Web Services Regions associated with the workload.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "notes": {
        "computed": true,
        "description": "The notes associated with the workload.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "review_owner": {
        "computed": true,
        "description": "The review owner of the workload.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags associated with the workload.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "workload_arn": {
        "computed": true,
        "description": "The ARN for the workload.",
        "description_kind": "plain",
        "type": "string"
      },
      "workload_id": {
        "computed": true,
        "description": "The ID assigned to the workload.",
        "description_kind": "plain",
        "type": "string"
      },
      "workload_name": {
        "description": "The name of the workload.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Definition of AWS::WellArchitected::Workload Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccWellarchitectedWorkloadSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWellarchitectedWorkload), &result)
	return &result
}
