package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatasyncStorageSystem = `{
  "block": {
    "attributes": {
      "agent_arns": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "cloudwatch_log_group_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "connectivity_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "secrets_manager_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "server_configuration": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "server_hostname": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "server_port": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "server_credentials": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "password": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "username": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
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
        "description_kind": "plain",
        "required": true,
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
      }
    },
    "description": "Resource Type definition for AWS::DataSync::StorageSystem",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDatasyncStorageSystemSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatasyncStorageSystem), &result)
	return &result
}
