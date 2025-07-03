package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRedshiftserverlessSnapshot = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "namespace_name": {
        "computed": true,
        "description": "The namespace the snapshot is associated with.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "owner_account": {
        "computed": true,
        "description": "The owner account of the snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "retention_period": {
        "computed": true,
        "description": "The retention period of the snapshot.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "snapshot": {
        "computed": true,
        "description": "Definition for snapshot resource",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "admin_username": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "kms_key_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "namespace_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "namespace_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "owner_account": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "retention_period": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "snapshot_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "snapshot_create_time": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "snapshot_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "status": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "snapshot_name": {
        "description": "The name of the snapshot.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
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
    "description": "Resource Type definition for AWS::RedshiftServerless::Snapshot Resource Type.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRedshiftserverlessSnapshotSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRedshiftserverlessSnapshot), &result)
	return &result
}
