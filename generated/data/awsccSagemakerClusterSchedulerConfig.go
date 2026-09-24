package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerClusterSchedulerConfig = `{
  "block": {
    "attributes": {
      "cluster_arn": {
        "computed": true,
        "description": "ARN of the cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_scheduler_config_arn": {
        "computed": true,
        "description": "ARN of the cluster policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_scheduler_config_id": {
        "computed": true,
        "description": "ID of the cluster policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_scheduler_config_version": {
        "computed": true,
        "description": "Version of the cluster policy.",
        "description_kind": "plain",
        "type": "number"
      },
      "creation_time": {
        "computed": true,
        "description": "Creation time of the cluster policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Description of the cluster policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Name for the cluster policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "scheduler_config": {
        "computed": true,
        "description": "Cluster policy configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "fair_share": {
              "computed": true,
              "description": "When enabled, entities borrow idle compute based on assigned FairShareWeight.",
              "description_kind": "plain",
              "type": "string"
            },
            "idle_resource_sharing": {
              "computed": true,
              "description": "Configuration for sharing idle compute resources across entities.",
              "description_kind": "plain",
              "type": "string"
            },
            "priority_classes": {
              "computed": true,
              "description": "List of priority class configurations.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
                    "computed": true,
                    "description": "Name of the priority class.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "weight": {
                    "computed": true,
                    "description": "Weight of the priority class. Range 0-100, default 0.",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "status": {
        "computed": true,
        "description": "Status of the cluster policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags of the cluster policy.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::SageMaker::ClusterSchedulerConfig",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSagemakerClusterSchedulerConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerClusterSchedulerConfig), &result)
	return &result
}
