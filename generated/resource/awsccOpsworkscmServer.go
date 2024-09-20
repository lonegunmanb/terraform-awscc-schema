package resource

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
        "optional": true,
        "type": "bool"
      },
      "backup_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "backup_retention_count": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "custom_certificate": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "custom_domain": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "custom_private_key": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "disable_automated_backup": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
        "optional": true,
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
      "engine_model": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "engine_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_profile_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "key_pair": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "preferred_backup_window": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "preferred_maintenance_window": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_group_ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "server_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_role_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "subnet_ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
    "description": "Resource Type definition for AWS::OpsWorksCM::Server",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccOpsworkscmServerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOpsworkscmServer), &result)
	return &result
}
