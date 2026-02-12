package data

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
                    "type": "string"
                  },
                  "role_arn": {
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
      },
      "execution_role_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
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
        "type": [
          "list",
          "string"
        ]
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "query_language": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "query_string": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "schedule_end_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "schedule_expression": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "schedule_start_time": {
        "computed": true,
        "description_kind": "plain",
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
        "type": "number"
      },
      "state": {
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
      "timezone": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Logs::ScheduledQuery",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLogsScheduledQuerySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLogsScheduledQuery), &result)
	return &result
}
