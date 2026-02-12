package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLogsScheduledQuery = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "s3_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "destination_identifier": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "role_arn": {
                    "computed": true,
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
        "optional": true
      },
      "execution_role_arn": {
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
      "last_execution_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "last_triggered_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "last_updated_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "log_group_identifiers": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "query_language": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "query_string": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "schedule_end_time": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "schedule_expression": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "schedule_start_time": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "scheduled_query_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "start_time_offset": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "state": {
        "computed": true,
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
      "timezone": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Creates a new Scheduled Query that allows you to define a Logs Insights query that will run on a schedule and configure actions to take with the query results.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLogsScheduledQuerySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLogsScheduledQuery), &result)
	return &result
}
