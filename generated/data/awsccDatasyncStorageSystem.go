package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatasyncStorageSystem = `{
  "block": {
    "attributes": {
      "agent_arns": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "cloudwatch_log_group_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "connectivity_status": {
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
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secrets_manager_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "server_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "server_hostname": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "server_port": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "server_credentials": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "password": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "username": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "storage_system_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "storage_system_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "system_type": {
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
      }
    },
    "description": "Data Source schema for AWS::DataSync::StorageSystem",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDatasyncStorageSystemSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatasyncStorageSystem), &result)
	return &result
}
