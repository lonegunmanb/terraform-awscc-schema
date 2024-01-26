package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRdsDbCluster = `{
  "block": {
    "attributes": {
      "allocated_storage": {
        "computed": true,
        "description": "The amount of storage in gibibytes (GiB) to allocate to each DB instance in the Multi-AZ DB cluster.",
        "description_kind": "plain",
        "type": "number"
      },
      "associated_roles": {
        "computed": true,
        "description": "Provides a list of the AWS Identity and Access Management (IAM) roles that are associated with the DB cluster. IAM roles that are associated with a DB cluster grant permission for the DB cluster to access other AWS services on your behalf.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "feature_name": {
              "computed": true,
              "description": "The name of the feature associated with the AWS Identity and Access Management (IAM) role. For the list of supported feature names, see DBEngineVersion in the Amazon RDS API Reference.",
              "description_kind": "plain",
              "type": "string"
            },
            "role_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the IAM role that is associated with the DB cluster.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "auto_minor_version_upgrade": {
        "computed": true,
        "description": "A value that indicates whether minor engine upgrades are applied automatically to the DB cluster during the maintenance window. By default, minor engine upgrades are applied automatically.",
        "description_kind": "plain",
        "type": "bool"
      },
      "availability_zones": {
        "computed": true,
        "description": "A list of Availability Zones (AZs) where instances in the DB cluster can be created. For information on AWS Regions and Availability Zones, see Choosing the Regions and Availability Zones in the Amazon Aurora User Guide.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "backtrack_window": {
        "computed": true,
        "description": "The target backtrack window, in seconds. To disable backtracking, set this value to 0.",
        "description_kind": "plain",
        "type": "number"
      },
      "backup_retention_period": {
        "computed": true,
        "description": "The number of days for which automated backups are retained.",
        "description_kind": "plain",
        "type": "number"
      },
      "copy_tags_to_snapshot": {
        "computed": true,
        "description": "A value that indicates whether to copy all tags from the DB cluster to snapshots of the DB cluster. The default is not to copy them.",
        "description_kind": "plain",
        "type": "bool"
      },
      "database_name": {
        "computed": true,
        "description": "The name of your database. If you don't provide a name, then Amazon RDS won't create a database in this DB cluster. For naming constraints, see Naming Constraints in the Amazon RDS User Guide.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_cluster_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the DB cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_cluster_identifier": {
        "computed": true,
        "description": "The DB cluster identifier. This parameter is stored as a lowercase string.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_cluster_instance_class": {
        "computed": true,
        "description": "The compute and memory capacity of each DB instance in the Multi-AZ DB cluster, for example db.m6g.xlarge.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_cluster_parameter_group_name": {
        "computed": true,
        "description": "The name of the DB cluster parameter group to associate with this DB cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_cluster_resource_id": {
        "computed": true,
        "description": "The AWS Region-unique, immutable identifier for the DB cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_instance_parameter_group_name": {
        "computed": true,
        "description": "The name of the DB parameter group to apply to all instances of the DB cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_subnet_group_name": {
        "computed": true,
        "description": "A DB subnet group that you want to associate with this DB cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_system_id": {
        "computed": true,
        "description": "Reserved for future use.",
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_protection": {
        "computed": true,
        "description": "A value that indicates whether the DB cluster has deletion protection enabled. The database can't be deleted when deletion protection is enabled. By default, deletion protection is disabled.",
        "description_kind": "plain",
        "type": "bool"
      },
      "domain": {
        "computed": true,
        "description": "The Active Directory directory ID to create the DB cluster in.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_iam_role_name": {
        "computed": true,
        "description": "Specify the name of the IAM role to be used when making API calls to the Directory Service.",
        "description_kind": "plain",
        "type": "string"
      },
      "enable_cloudwatch_logs_exports": {
        "computed": true,
        "description": "The list of log types that need to be enabled for exporting to CloudWatch Logs. The values in the list depend on the DB engine being used. For more information, see Publishing Database Logs to Amazon CloudWatch Logs in the Amazon Aurora User Guide.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "enable_global_write_forwarding": {
        "computed": true,
        "description": "Specifies whether to enable this DB cluster to forward write operations to the primary cluster of a global cluster (Aurora global database). By default, write operations are not allowed on Aurora DB clusters that are secondary clusters in an Aurora global database.",
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_http_endpoint": {
        "computed": true,
        "description": "A value that indicates whether to enable the HTTP endpoint for DB cluster. By default, the HTTP endpoint is disabled.",
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_iam_database_authentication": {
        "computed": true,
        "description": "A value that indicates whether to enable mapping of AWS Identity and Access Management (IAM) accounts to database accounts. By default, mapping is disabled.",
        "description_kind": "plain",
        "type": "bool"
      },
      "endpoint": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "address": {
              "computed": true,
              "description": "The connection endpoint for the DB cluster.",
              "description_kind": "plain",
              "type": "string"
            },
            "port": {
              "computed": true,
              "description": "The port number that will accept connections on this DB cluster.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "engine": {
        "computed": true,
        "description": "The name of the database engine to be used for this DB cluster. Valid Values: aurora (for MySQL 5.6-compatible Aurora), aurora-mysql (for MySQL 5.7-compatible Aurora), and aurora-postgresql",
        "description_kind": "plain",
        "type": "string"
      },
      "engine_mode": {
        "computed": true,
        "description": "The DB engine mode of the DB cluster, either provisioned, serverless, parallelquery, global, or multimaster.",
        "description_kind": "plain",
        "type": "string"
      },
      "engine_version": {
        "computed": true,
        "description": "The version number of the database engine to use.",
        "description_kind": "plain",
        "type": "string"
      },
      "global_cluster_identifier": {
        "computed": true,
        "description": "If you are configuring an Aurora global database cluster and want your Aurora DB cluster to be a secondary member in the global database cluster, specify the global cluster ID of the global database cluster. To define the primary database cluster of the global cluster, use the AWS::RDS::GlobalCluster resource.\n\nIf you aren't configuring a global database cluster, don't specify this property.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "iops": {
        "computed": true,
        "description": "The amount of Provisioned IOPS (input/output operations per second) to be initially allocated for each DB instance in the Multi-AZ DB cluster.",
        "description_kind": "plain",
        "type": "number"
      },
      "kms_key_id": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the AWS Key Management Service master key that is used to encrypt the database instances in the DB cluster, such as arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef. If you enable the StorageEncrypted property but don't specify this property, the default master key is used. If you specify this property, you must set the StorageEncrypted property to true.",
        "description_kind": "plain",
        "type": "string"
      },
      "manage_master_user_password": {
        "computed": true,
        "description": "A value that indicates whether to manage the master user password with AWS Secrets Manager.",
        "description_kind": "plain",
        "type": "bool"
      },
      "master_user_password": {
        "computed": true,
        "description": "The master password for the DB instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "master_user_secret": {
        "computed": true,
        "description": "Contains the secret managed by RDS in AWS Secrets Manager for the master user password.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kms_key_id": {
              "computed": true,
              "description": "The AWS KMS key identifier that is used to encrypt the secret.",
              "description_kind": "plain",
              "type": "string"
            },
            "secret_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the secret.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "master_username": {
        "computed": true,
        "description": "The name of the master user for the DB cluster. You must specify MasterUsername, unless you specify SnapshotIdentifier. In that case, don't specify MasterUsername.",
        "description_kind": "plain",
        "type": "string"
      },
      "monitoring_interval": {
        "computed": true,
        "description": "The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB cluster. To turn off collecting Enhanced Monitoring metrics, specify 0. The default is 0.",
        "description_kind": "plain",
        "type": "number"
      },
      "monitoring_role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the IAM role that permits RDS to send Enhanced Monitoring metrics to Amazon CloudWatch Logs.",
        "description_kind": "plain",
        "type": "string"
      },
      "network_type": {
        "computed": true,
        "description": "The network type of the DB cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "performance_insights_enabled": {
        "computed": true,
        "description": "A value that indicates whether to turn on Performance Insights for the DB cluster.",
        "description_kind": "plain",
        "type": "bool"
      },
      "performance_insights_kms_key_id": {
        "computed": true,
        "description": "The Amazon Web Services KMS key identifier for encryption of Performance Insights data.",
        "description_kind": "plain",
        "type": "string"
      },
      "performance_insights_retention_period": {
        "computed": true,
        "description": "The amount of time, in days, to retain Performance Insights data.",
        "description_kind": "plain",
        "type": "number"
      },
      "port": {
        "computed": true,
        "description": "The port number on which the instances in the DB cluster accept connections. Default: 3306 if engine is set as aurora or 5432 if set to aurora-postgresql.",
        "description_kind": "plain",
        "type": "number"
      },
      "preferred_backup_window": {
        "computed": true,
        "description": "The daily time range during which automated backups are created if automated backups are enabled using the BackupRetentionPeriod parameter. The default is a 30-minute window selected at random from an 8-hour block of time for each AWS Region. To see the time blocks available, see Adjusting the Preferred DB Cluster Maintenance Window in the Amazon Aurora User Guide.",
        "description_kind": "plain",
        "type": "string"
      },
      "preferred_maintenance_window": {
        "computed": true,
        "description": "The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC). The default is a 30-minute window selected at random from an 8-hour block of time for each AWS Region, occurring on a random day of the week. To see the time blocks available, see Adjusting the Preferred DB Cluster Maintenance Window in the Amazon Aurora User Guide.",
        "description_kind": "plain",
        "type": "string"
      },
      "publicly_accessible": {
        "computed": true,
        "description": "A value that indicates whether the DB cluster is publicly accessible.",
        "description_kind": "plain",
        "type": "bool"
      },
      "read_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "address": {
              "computed": true,
              "description": "The reader endpoint for the DB cluster.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "replication_source_identifier": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the source DB instance or DB cluster if this DB cluster is created as a Read Replica.",
        "description_kind": "plain",
        "type": "string"
      },
      "restore_to_time": {
        "computed": true,
        "description": "The date and time to restore the DB cluster to. Value must be a time in Universal Coordinated Time (UTC) format. An example: 2015-03-07T23:45:00Z",
        "description_kind": "plain",
        "type": "string"
      },
      "restore_type": {
        "computed": true,
        "description": "The type of restore to be performed. You can specify one of the following values:\nfull-copy - The new DB cluster is restored as a full copy of the source DB cluster.\ncopy-on-write - The new DB cluster is restored as a clone of the source DB cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "scaling_configuration": {
        "computed": true,
        "description": "The ScalingConfiguration property type specifies the scaling configuration of an Aurora Serverless DB cluster.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "auto_pause": {
              "computed": true,
              "description": "A value that indicates whether to allow or disallow automatic pause for an Aurora DB cluster in serverless DB engine mode. A DB cluster can be paused only when it's idle (it has no connections).",
              "description_kind": "plain",
              "type": "bool"
            },
            "max_capacity": {
              "computed": true,
              "description": "The maximum capacity for an Aurora DB cluster in serverless DB engine mode.\nFor Aurora MySQL, valid capacity values are 1, 2, 4, 8, 16, 32, 64, 128, and 256.\nFor Aurora PostgreSQL, valid capacity values are 2, 4, 8, 16, 32, 64, 192, and 384.\nThe maximum capacity must be greater than or equal to the minimum capacity.",
              "description_kind": "plain",
              "type": "number"
            },
            "min_capacity": {
              "computed": true,
              "description": "The minimum capacity for an Aurora DB cluster in serverless DB engine mode.\nFor Aurora MySQL, valid capacity values are 1, 2, 4, 8, 16, 32, 64, 128, and 256.\nFor Aurora PostgreSQL, valid capacity values are 2, 4, 8, 16, 32, 64, 192, and 384.\nThe minimum capacity must be less than or equal to the maximum capacity.",
              "description_kind": "plain",
              "type": "number"
            },
            "seconds_before_timeout": {
              "computed": true,
              "description": "The amount of time, in seconds, that Aurora Serverless v1 tries to find a scaling point to perform seamless scaling before enforcing the timeout action.\nThe default is 300.",
              "description_kind": "plain",
              "type": "number"
            },
            "seconds_until_auto_pause": {
              "computed": true,
              "description": "The time, in seconds, before an Aurora DB cluster in serverless mode is paused.",
              "description_kind": "plain",
              "type": "number"
            },
            "timeout_action": {
              "computed": true,
              "description": "The action to take when the timeout is reached, either ForceApplyCapacityChange or RollbackCapacityChange.\nForceApplyCapacityChange sets the capacity to the specified value as soon as possible.\nRollbackCapacityChange, the default, ignores the capacity change if a scaling point isn't found in the timeout period.\n\nFor more information, see Autoscaling for Aurora Serverless v1 in the Amazon Aurora User Guide.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "serverless_v2_scaling_configuration": {
        "computed": true,
        "description": "Contains the scaling configuration of an Aurora Serverless v2 DB cluster.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "max_capacity": {
              "computed": true,
              "description": "The maximum number of Aurora capacity units (ACUs) for a DB instance in an Aurora Serverless v2 cluster. You can specify ACU values in half-step increments, such as 40, 40.5, 41, and so on. The largest value that you can use is 128.",
              "description_kind": "plain",
              "type": "number"
            },
            "min_capacity": {
              "computed": true,
              "description": "The minimum number of Aurora capacity units (ACUs) for a DB instance in an Aurora Serverless v2 cluster. You can specify ACU values in half-step increments, such as 8, 8.5, 9, and so on. The smallest value that you can use is 0.5.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "snapshot_identifier": {
        "computed": true,
        "description": "The identifier for the DB snapshot or DB cluster snapshot to restore from.\nYou can use either the name or the Amazon Resource Name (ARN) to specify a DB cluster snapshot. However, you can use only the ARN to specify a DB snapshot.\nAfter you restore a DB cluster with a SnapshotIdentifier property, you must specify the same SnapshotIdentifier property for any future updates to the DB cluster. When you specify this property for an update, the DB cluster is not restored from the snapshot again, and the data in the database is not changed. However, if you don't specify the SnapshotIdentifier property, an empty DB cluster is created, and the original DB cluster is deleted. If you specify a property that is different from the previous snapshot restore property, the DB cluster is restored from the specified SnapshotIdentifier property, and the original DB cluster is deleted.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_db_cluster_identifier": {
        "computed": true,
        "description": "The identifier of the source DB cluster from which to restore.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_region": {
        "computed": true,
        "description": "The AWS Region which contains the source DB cluster when replicating a DB cluster. For example, us-east-1.",
        "description_kind": "plain",
        "type": "string"
      },
      "storage_encrypted": {
        "computed": true,
        "description": "Indicates whether the DB instance is encrypted.\nIf you specify the DBClusterIdentifier, SnapshotIdentifier, or SourceDBInstanceIdentifier property, don't specify this property. The value is inherited from the cluster, snapshot, or source DB instance.",
        "description_kind": "plain",
        "type": "bool"
      },
      "storage_type": {
        "computed": true,
        "description": "Specifies the storage type to be associated with the DB cluster.",
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
      },
      "use_latest_restorable_time": {
        "computed": true,
        "description": "A value that indicates whether to restore the DB cluster to the latest restorable backup time. By default, the DB cluster is not restored to the latest restorable backup time.",
        "description_kind": "plain",
        "type": "bool"
      },
      "vpc_security_group_ids": {
        "computed": true,
        "description": "A list of EC2 VPC security groups to associate with this DB cluster.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::RDS::DBCluster",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRdsDbClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRdsDbCluster), &result)
	return &result
}
