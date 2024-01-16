package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRedshiftScheduledAction = `{
  "block": {
    "attributes": {
      "enable": {
        "computed": true,
        "description": "If true, the schedule is enabled. If false, the scheduled action does not trigger.",
        "description_kind": "plain",
        "type": "bool"
      },
      "end_time": {
        "computed": true,
        "description": "The end time in UTC of the scheduled action. After this time, the scheduled action does not trigger.",
        "description_kind": "plain",
        "type": "string"
      },
      "iam_role": {
        "computed": true,
        "description": "The IAM role to assume to run the target action.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "next_invocations": {
        "computed": true,
        "description": "List of times when the scheduled action will run.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "schedule": {
        "computed": true,
        "description": "The schedule in ` + "`" + `at( )` + "`" + ` or ` + "`" + `cron( )` + "`" + ` format.",
        "description_kind": "plain",
        "type": "string"
      },
      "scheduled_action_description": {
        "computed": true,
        "description": "The description of the scheduled action.",
        "description_kind": "plain",
        "type": "string"
      },
      "scheduled_action_name": {
        "computed": true,
        "description": "The name of the scheduled action. The name must be unique within an account.",
        "description_kind": "plain",
        "type": "string"
      },
      "start_time": {
        "computed": true,
        "description": "The start time in UTC of the scheduled action. Before this time, the scheduled action does not trigger.",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The state of the scheduled action.",
        "description_kind": "plain",
        "type": "string"
      },
      "target_action": {
        "computed": true,
        "description": "A JSON format string of the Amazon Redshift API operation with input parameters.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "pause_cluster": {
              "computed": true,
              "description": "Describes a pause cluster operation. For example, a scheduled action to run the ` + "`" + `PauseCluster` + "`" + ` API operation.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cluster_identifier": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "resize_cluster": {
              "computed": true,
              "description": "Describes a resize cluster operation. For example, a scheduled action to run the ` + "`" + `ResizeCluster` + "`" + ` API operation.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "classic": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "cluster_identifier": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "cluster_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "node_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "number_of_nodes": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "resume_cluster": {
              "computed": true,
              "description": "Describes a resume cluster operation. For example, a scheduled action to run the ` + "`" + `ResumeCluster` + "`" + ` API operation.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cluster_identifier": {
                    "computed": true,
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
    "description": "Data Source schema for AWS::Redshift::ScheduledAction",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRedshiftScheduledActionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRedshiftScheduledAction), &result)
	return &result
}
