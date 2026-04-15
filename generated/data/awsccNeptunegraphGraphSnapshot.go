package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNeptunegraphGraphSnapshot = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN of the graph snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "graph_identifier": {
        "computed": true,
        "description": "The unique identifier of the Neptune Analytics graph to create the snapshot from.",
        "description_kind": "plain",
        "type": "string"
      },
      "graph_snapshot_id": {
        "computed": true,
        "description": "The unique identifier of the graph snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kms_key_identifier": {
        "computed": true,
        "description": "The ID of the KMS key used to encrypt and decrypt the snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "snapshot_create_time": {
        "computed": true,
        "description": "The time when the snapshot was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "snapshot_name": {
        "computed": true,
        "description": "The snapshot name.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The current status of the graph snapshot.",
        "description_kind": "plain",
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
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::NeptuneGraph::GraphSnapshot",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccNeptunegraphGraphSnapshotSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNeptunegraphGraphSnapshot), &result)
	return &result
}
