package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEcsDaemon = `{
  "block": {
    "attributes": {
      "capacity_provider_arns": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "cluster_arn": {
        "computed": true,
        "description_kind": "plain",
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
        "type": "string"
      },
      "daemon_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "daemon_task_definition_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "deployment_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "deployment_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "alarms": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "alarm_names": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "enable": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "bake_time_in_minutes": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "drain_percent": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "enable_ecs_managed_tags": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_execute_command": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "propagate_tags": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
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
      },
      "updated_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ECS::Daemon",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEcsDaemonSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEcsDaemon), &result)
	return &result
}
