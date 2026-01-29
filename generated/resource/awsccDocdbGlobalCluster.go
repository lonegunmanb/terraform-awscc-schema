package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDocdbGlobalCluster = `{
  "block": {
    "attributes": {
      "deletion_protection": {
        "computed": true,
        "description": "Indicates whether the global cluster has deletion protection enabled. The global cluster can't be deleted when deletion protection is enabled.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "engine": {
        "computed": true,
        "description": "The database engine to use for this global cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "engine_version": {
        "computed": true,
        "description": "The engine version to use for this global cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "global_cluster_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the global cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "global_cluster_identifier": {
        "description": "The cluster identifier of the global cluster.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "global_cluster_resource_id": {
        "computed": true,
        "description": "The AWS Region-unique, immutable identifier for the global database cluster. This identifier is found in AWS CloudTrail log entries whenever the AWS KMS key for the DB cluster is accessed.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_db_cluster_identifier": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) to use as the primary cluster of the global cluster. You may also choose to instead specify the DBClusterIdentifier. If you provide a value for this parameter, don't specify values for the following settings because Amazon DocumentDB uses the values from the specified source DB cluster: Engine, EngineVersion, StorageEncrypted",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "storage_encrypted": {
        "computed": true,
        "description": "Indicates whether the global cluster has storage encryption enabled.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "tags": {
        "computed": true,
        "description": "The tags to be assigned to the Amazon DocumentDB resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "The AWS::DocDB::GlobalCluster resource represents an Amazon DocumentDB Global Cluster.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDocdbGlobalClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDocdbGlobalCluster), &result)
	return &result
}
