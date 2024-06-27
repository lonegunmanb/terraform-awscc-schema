package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRdsGlobalCluster = `{
  "block": {
    "attributes": {
      "deletion_protection": {
        "computed": true,
        "description": "The deletion protection setting for the new global database. The global database can't be deleted when deletion protection is enabled.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "engine": {
        "computed": true,
        "description": "The name of the database engine to be used for this DB cluster. Valid Values: aurora (for MySQL 5.6-compatible Aurora), aurora-mysql (for MySQL 5.7-compatible Aurora).\nIf you specify the SourceDBClusterIdentifier property, don't specify this property. The value is inherited from the cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "engine_lifecycle_support": {
        "computed": true,
        "description": "The life cycle type of the global cluster. You can use this setting to enroll your global cluster into Amazon RDS Extended Support.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "engine_version": {
        "computed": true,
        "description": "The version number of the database engine to use. If you specify the SourceDBClusterIdentifier property, don't specify this property. The value is inherited from the cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "global_cluster_identifier": {
        "computed": true,
        "description": "The cluster identifier of the new global database cluster. This parameter is stored as a lowercase string.",
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
      "source_db_cluster_identifier": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) to use as the primary cluster of the global database. This parameter is optional. This parameter is stored as a lowercase string.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "storage_encrypted": {
        "computed": true,
        "description": " The storage encryption setting for the new global database cluster.\nIf you specify the SourceDBClusterIdentifier property, don't specify this property. The value is inherited from the cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "description": "Resource Type definition for AWS::RDS::GlobalCluster",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRdsGlobalClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRdsGlobalCluster), &result)
	return &result
}
