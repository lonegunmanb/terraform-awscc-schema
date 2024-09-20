package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNeptuneDbCluster = `{
  "block": {
    "attributes": {
      "associated_roles": {
        "computed": true,
        "description": "Provides a list of the AWS Identity and Access Management (IAM) roles that are associated with the DB cluster. IAM roles that are associated with a DB cluster grant permission for the DB cluster to access other AWS services on your behalf.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "feature_name": {
              "computed": true,
              "description": "The name of the feature associated with the AWS Identity and Access Management (IAM) role. For the list of supported feature names, see DBEngineVersion in the Amazon Neptune API Reference.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "role_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the IAM role that is associated with the DB cluster.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "availability_zones": {
        "computed": true,
        "description": "Provides the list of EC2 Availability Zones that instances in the DB cluster can be created in.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "backup_retention_period": {
        "computed": true,
        "description": "Specifies the number of days for which automatic DB snapshots are retained.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "cluster_resource_id": {
        "computed": true,
        "description": "The resource id for the DB cluster. For example: ` + "`" + `cluster-ABCD1234EFGH5678IJKL90MNOP` + "`" + `. The cluster ID uniquely identifies the cluster and is used in things like IAM authentication policies.",
        "description_kind": "plain",
        "type": "string"
      },
      "copy_tags_to_snapshot": {
        "computed": true,
        "description": "A value that indicates whether to copy all tags from the DB cluster to snapshots of the DB cluster. The default behaviour is not to copy them.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "db_cluster_identifier": {
        "computed": true,
        "description": "The DB cluster identifier. Contains a user-supplied DB cluster identifier. This identifier is the unique key that identifies a DB cluster stored as a lowercase string.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_cluster_parameter_group_name": {
        "computed": true,
        "description": "Provides the name of the DB cluster parameter group.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_instance_parameter_group_name": {
        "computed": true,
        "description": "The name of the DB parameter group to apply to all instances of the DB cluster. Used only in case of a major EngineVersion upgrade request.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_port": {
        "computed": true,
        "description": "The port number on which the DB instances in the DB cluster accept connections. \n\nIf not specified, the default port used is ` + "`" + `8182` + "`" + `. \n\nNote: ` + "`" + `Port` + "`" + ` property will soon be deprecated from this resource. Please update existing templates to rename it with new property ` + "`" + `DBPort` + "`" + ` having same functionalities.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "db_subnet_group_name": {
        "computed": true,
        "description": "Specifies information on the subnet group associated with the DB cluster, including the name, description, and subnets in the subnet group.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "deletion_protection": {
        "computed": true,
        "description": "Indicates whether or not the DB cluster has deletion protection enabled. The database can't be deleted when deletion protection is enabled.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_cloudwatch_logs_exports": {
        "computed": true,
        "description": "Specifies a list of log types that are enabled for export to CloudWatch Logs.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "endpoint": {
        "computed": true,
        "description": "The connection endpoint for the DB cluster. For example: ` + "`" + `mystack-mydbcluster-1apw1j4phylrk.cg034hpkmmjt.us-east-2.rds.amazonaws.com` + "`" + `",
        "description_kind": "plain",
        "type": "string"
      },
      "engine_version": {
        "computed": true,
        "description": "Indicates the database engine version.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "iam_auth_enabled": {
        "computed": true,
        "description": "True if mapping of Amazon Identity and Access Management (IAM) accounts to database accounts is enabled, and otherwise false.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the AWS KMS key that is used to encrypt the database instances in the DB cluster, such as arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef. If you enable the StorageEncrypted property but don't specify this property, the default KMS key is used. If you specify this property, you must set the StorageEncrypted property to true.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "port": {
        "computed": true,
        "description": "The port number on which the DB cluster accepts connections. For example: ` + "`" + `8182` + "`" + `.",
        "description_kind": "plain",
        "type": "string"
      },
      "preferred_backup_window": {
        "computed": true,
        "description": "Specifies the daily time range during which automated backups are created if automated backups are enabled, as determined by the BackupRetentionPeriod.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "preferred_maintenance_window": {
        "computed": true,
        "description": "Specifies the weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "read_endpoint": {
        "computed": true,
        "description": "The reader endpoint for the DB cluster. For example: ` + "`" + `mystack-mydbcluster-ro-1apw1j4phylrk.cg034hpkmmjt.us-east-2.rds.amazonaws.com` + "`" + `",
        "description_kind": "plain",
        "type": "string"
      },
      "restore_to_time": {
        "computed": true,
        "description": "Creates a new DB cluster from a DB snapshot or DB cluster snapshot.\n\nIf a DB snapshot is specified, the target DB cluster is created from the source DB snapshot with a default configuration and default security group.\n\nIf a DB cluster snapshot is specified, the target DB cluster is created from the source DB cluster restore point with the same configuration as the original source DB cluster, except that the new DB cluster is created with the default security group.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "restore_type": {
        "computed": true,
        "description": "Creates a new DB cluster from a DB snapshot or DB cluster snapshot.\n\nIf a DB snapshot is specified, the target DB cluster is created from the source DB snapshot with a default configuration and default security group.\n\nIf a DB cluster snapshot is specified, the target DB cluster is created from the source DB cluster restore point with the same configuration as the original source DB cluster, except that the new DB cluster is created with the default security group.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "serverless_scaling_configuration": {
        "computed": true,
        "description": "Contains the scaling configuration used by the Neptune Serverless Instances within this DB cluster.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "max_capacity": {
              "computed": true,
              "description": "The maximum number of Neptune capacity units (NCUs) for a DB instance in an Neptune Serverless cluster. You can specify NCU values in half-step increments, such as 40, 40.5, 41, and so on. The smallest value you can use is 2.5, whereas the largest is 128.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "min_capacity": {
              "computed": true,
              "description": "The minimum number of Neptune capacity units (NCUs) for a DB instance in an Neptune Serverless cluster. You can specify NCU values in half-step increments, such as 8, 8.5, 9, and so on. The smallest value you can use is 1, whereas the largest is 128.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "snapshot_identifier": {
        "computed": true,
        "description": "Specifies the identifier for a DB cluster snapshot. Must match the identifier of an existing snapshot.\n\nAfter you restore a DB cluster using a SnapshotIdentifier, you must specify the same SnapshotIdentifier for any future updates to the DB cluster. When you specify this property for an update, the DB cluster is not restored from the snapshot again, and the data in the database is not changed.\n\nHowever, if you don't specify the SnapshotIdentifier, an empty DB cluster is created, and the original DB cluster is deleted. If you specify a property that is different from the previous snapshot restore property, the DB cluster is restored from the snapshot specified by the SnapshotIdentifier, and the original DB cluster is deleted.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_db_cluster_identifier": {
        "computed": true,
        "description": "Creates a new DB cluster from a DB snapshot or DB cluster snapshot.\n\nIf a DB snapshot is specified, the target DB cluster is created from the source DB snapshot with a default configuration and default security group.\n\nIf a DB cluster snapshot is specified, the target DB cluster is created from the source DB cluster restore point with the same configuration as the original source DB cluster, except that the new DB cluster is created with the default security group.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "storage_encrypted": {
        "computed": true,
        "description": "Indicates whether the DB cluster is encrypted.\n\nIf you specify the KmsKeyId property, then you must enable encryption and set this property to true.\n\nIf you enable the StorageEncrypted property but don't specify KmsKeyId property, then the default KMS key is used. If you specify KmsKeyId property, then that KMS Key is used to encrypt the database instances in the DB cluster.\n\nIf you specify the SourceDBClusterIdentifier property and don't specify this property or disable it. The value is inherited from the source DB cluster, and if the DB cluster is encrypted, the KmsKeyId property from the source cluster is used.\n\nIf you specify the DBSnapshotIdentifier and don't specify this property or disable it. The value is inherited from the snapshot, and the specified KmsKeyId property from the snapshot is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "tags": {
        "computed": true,
        "description": "The tags assigned to this cluster.",
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
      },
      "use_latest_restorable_time": {
        "computed": true,
        "description": "Creates a new DB cluster from a DB snapshot or DB cluster snapshot.\n\nIf a DB snapshot is specified, the target DB cluster is created from the source DB snapshot with a default configuration and default security group.\n\nIf a DB cluster snapshot is specified, the target DB cluster is created from the source DB cluster restore point with the same configuration as the original source DB cluster, except that the new DB cluster is created with the default security group.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "vpc_security_group_ids": {
        "computed": true,
        "description": "Provides a list of VPC security groups that the DB cluster belongs to.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "The AWS::Neptune::DBCluster resource creates an Amazon Neptune DB cluster.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccNeptuneDbClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNeptuneDbCluster), &result)
	return &result
}
