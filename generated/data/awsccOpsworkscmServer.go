package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccOpsworkscmServer = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "associate_public_ip_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "backup_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "backup_retention_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "custom_certificate": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "custom_domain": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "custom_private_key": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "disable_automated_backup": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "engine": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "engine_attributes": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
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
      "engine_model": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "engine_version": {
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
      "instance_profile_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "instance_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "key_pair": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "preferred_backup_window": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "preferred_maintenance_window": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "security_group_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "server_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "server_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_role_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
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
    "description": "Data Source schema for AWS::OpsWorksCM::Server",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccOpsworkscmServerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOpsworkscmServer), &result)
	return &result
}
