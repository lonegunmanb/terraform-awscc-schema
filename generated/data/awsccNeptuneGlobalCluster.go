package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNeptuneGlobalCluster = `{
  "block": {
    "attributes": {
      "deletion_protection": {
        "computed": true,
        "description": "Whether deletion protection is enabled.",
        "description_kind": "plain",
        "type": "bool"
      },
      "engine": {
        "computed": true,
        "description": "The name of the database engine.",
        "description_kind": "plain",
        "type": "string"
      },
      "engine_version": {
        "computed": true,
        "description": "The version number of the database engine.",
        "description_kind": "plain",
        "type": "string"
      },
      "global_cluster_identifier": {
        "computed": true,
        "description": "The cluster identifier of the global database cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_db_cluster_identifier": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of an existing Neptune DB cluster to use as the primary cluster of the new global database.",
        "description_kind": "plain",
        "type": "string"
      },
      "storage_encrypted": {
        "computed": true,
        "description": "Whether the global database cluster is storage encrypted.",
        "description_kind": "plain",
        "type": "bool"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::Neptune::GlobalCluster",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccNeptuneGlobalClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNeptuneGlobalCluster), &result)
	return &result
}
