package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEcsDaemon = `{
  "block": {
    "attributes": {
      "capacity_provider_arns": {
        "computed": true,
        "description": "The Amazon Resource Names (ARNs) of the capacity providers associated with the daemon.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "cluster_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the cluster that the daemon is running in.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "daemon_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "daemon_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "daemon_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "daemon_task_definition_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the daemon task definition used by this revision.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "deployment_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "deployment_configuration": {
        "computed": true,
        "description": "The deployment configuration used for this daemon deployment.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "alarms": {
              "computed": true,
              "description": "The CloudWatch alarm configuration for the daemon deployment. When alarms are triggered during a deployment, the deployment can be automatically rolled back.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "alarm_names": {
                    "computed": true,
                    "description": "The CloudWatch alarm names to monitor during a daemon deployment.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "enable": {
                    "computed": true,
                    "description": "Determines whether to use the CloudWatch alarm option in the daemon deployment process. The default value is ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "bake_time_in_minutes": {
              "computed": true,
              "description": "The amount of time (in minutes) to wait after a successful deployment step before proceeding. This allows time to monitor for issues before continuing. The default value is 0.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "drain_percent": {
              "computed": true,
              "description": "The percentage of container instances to drain simultaneously during a daemon deployment. Valid values are between 0.0 and 100.0.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "enable_ecs_managed_tags": {
        "computed": true,
        "description": "Specifies whether Amazon ECS managed tags are turned on for the daemon tasks.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_execute_command": {
        "computed": true,
        "description": "Specifies whether the execute command functionality is turned on for the daemon tasks.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "propagate_tags": {
        "computed": true,
        "description": "Specifies whether tags are propagated from the daemon to the daemon tasks.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "One part of a key-value pair that make up a tag. A ` + "`" + `` + "`" + `key` + "`" + `` + "`" + ` is a general label that acts like a category for more specific tag values.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The optional part of a key-value pair that make up a tag. A ` + "`" + `` + "`" + `value` + "`" + `` + "`" + ` acts as a descriptor within a tag category (key).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "updated_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Information about a daemon resource.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEcsDaemonSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEcsDaemon), &result)
	return &result
}
