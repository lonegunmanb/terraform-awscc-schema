package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRdsDbInstance = `{
  "block": {
    "attributes": {
      "allocated_storage": {
        "computed": true,
        "description": "The amount of storage (in gigabytes) to be initially allocated for the database instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "allow_major_version_upgrade": {
        "computed": true,
        "description": "A value that indicates whether major version upgrades are allowed. Changing this parameter doesn't result in an outage and the change is asynchronously applied as soon as possible.",
        "description_kind": "plain",
        "type": "bool"
      },
      "associated_roles": {
        "computed": true,
        "description": "The AWS Identity and Access Management (IAM) roles associated with the DB instance.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "feature_name": {
              "computed": true,
              "description": "The name of the feature associated with the AWS Identity and Access Management (IAM) role. IAM roles that are associated with a DB instance grant permission for the DB instance to access other AWS services on your behalf.",
              "description_kind": "plain",
              "type": "string"
            },
            "role_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the IAM role that is associated with the DB instance.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "auto_minor_version_upgrade": {
        "computed": true,
        "description": "A value that indicates whether minor engine upgrades are applied automatically to the DB instance during the maintenance window. By default, minor engine upgrades are applied automatically.",
        "description_kind": "plain",
        "type": "bool"
      },
      "automatic_backup_replication_region": {
        "computed": true,
        "description": "Enables replication of automated backups to a different Amazon Web Services Region.",
        "description_kind": "plain",
        "type": "string"
      },
      "availability_zone": {
        "computed": true,
        "description": "The Availability Zone (AZ) where the database will be created. For information on AWS Regions and Availability Zones.",
        "description_kind": "plain",
        "type": "string"
      },
      "backup_retention_period": {
        "computed": true,
        "description": "The number of days for which automated backups are retained. Setting this parameter to a positive number enables backups. Setting this parameter to 0 disables automated backups.",
        "description_kind": "plain",
        "type": "number"
      },
      "ca_certificate_identifier": {
        "computed": true,
        "description": "The identifier of the CA certificate for this DB instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_details": {
        "computed": true,
        "description": "Returns the details of the DB instance's server certificate.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ca_identifier": {
              "computed": true,
              "description": "The CA identifier of the CA certificate used for the DB instance's server certificate.",
              "description_kind": "plain",
              "type": "string"
            },
            "valid_till": {
              "computed": true,
              "description": "The expiration date of the DB instance’s server certificate.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "certificate_rotation_restart": {
        "computed": true,
        "description": "A value that indicates whether the DB instance is restarted when you rotate your SSL/TLS certificate.\nBy default, the DB instance is restarted when you rotate your SSL/TLS certificate. The certificate is not updated until the DB instance is restarted.\nIf you are using SSL/TLS to connect to the DB instance, follow the appropriate instructions for your DB engine to rotate your SSL/TLS certificate\nThis setting doesn't apply to RDS Custom.",
        "description_kind": "plain",
        "type": "bool"
      },
      "character_set_name": {
        "computed": true,
        "description": "For supported engines, indicates that the DB instance should be associated with the specified character set.",
        "description_kind": "plain",
        "type": "string"
      },
      "copy_tags_to_snapshot": {
        "computed": true,
        "description": "A value that indicates whether to copy tags from the DB instance to snapshots of the DB instance. By default, tags are not copied.",
        "description_kind": "plain",
        "type": "bool"
      },
      "custom_iam_instance_profile": {
        "computed": true,
        "description": "The instance profile associated with the underlying Amazon EC2 instance of an RDS Custom DB instance. The instance profile must meet the following requirements:\n * The profile must exist in your account.\n * The profile must have an IAM role that Amazon EC2 has permissions to assume.\n * The instance profile name and the associated IAM role name must start with the prefix AWSRDSCustom .\nFor the list of permissions required for the IAM role, see Configure IAM and your VPC in the Amazon RDS User Guide .\n\nThis setting is required for RDS Custom.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_cluster_identifier": {
        "computed": true,
        "description": "The identifier of the DB cluster that the instance will belong to.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_cluster_snapshot_identifier": {
        "computed": true,
        "description": "The identifier for the RDS for MySQL Multi-AZ DB cluster snapshot to restore from. For more information on Multi-AZ DB clusters, see Multi-AZ deployments with two readable standby DB instances in the Amazon RDS User Guide .\n\nConstraints:\n * Must match the identifier of an existing Multi-AZ DB cluster snapshot.\n * Can't be specified when DBSnapshotIdentifier is specified.\n * Must be specified when DBSnapshotIdentifier isn't specified.\n * If you are restoring from a shared manual Multi-AZ DB cluster snapshot, the DBClusterSnapshotIdentifier must be the ARN of the shared snapshot.\n * Can't be the identifier of an Aurora DB cluster snapshot.\n * Can't be the identifier of an RDS for PostgreSQL Multi-AZ DB cluster snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_instance_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the DB instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_instance_class": {
        "computed": true,
        "description": "The compute and memory capacity of the DB instance, for example, db.m4.large. Not all DB instance classes are available in all AWS Regions, or for all database engines.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_instance_identifier": {
        "computed": true,
        "description": "A name for the DB instance. If you specify a name, AWS CloudFormation converts it to lowercase. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the DB instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_name": {
        "computed": true,
        "description": "The meaning of this parameter differs according to the database engine you use.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_parameter_group_name": {
        "computed": true,
        "description": "The name of an existing DB parameter group or a reference to an AWS::RDS::DBParameterGroup resource created in the template.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_security_groups": {
        "computed": true,
        "description": "A list of the DB security groups to assign to the DB instance. The list can include both the name of existing DB security groups or references to AWS::RDS::DBSecurityGroup resources created in the template.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "db_snapshot_identifier": {
        "computed": true,
        "description": "The name or Amazon Resource Name (ARN) of the DB snapshot that's used to restore the DB instance. If you're restoring from a shared manual DB snapshot, you must specify the ARN of the snapshot.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_subnet_group_name": {
        "computed": true,
        "description": "A DB subnet group to associate with the DB instance. If you update this value, the new subnet group must be a subnet group in a new VPC.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_system_id": {
        "computed": true,
        "description": "The Oracle system ID (Oracle SID) for a container database (CDB). The Oracle SID is also the name of the CDB. This setting is valid for RDS Custom only.",
        "description_kind": "plain",
        "type": "string"
      },
      "dbi_resource_id": {
        "computed": true,
        "description": "The AWS Region-unique, immutable identifier for the DB instance. This identifier is found in AWS CloudTrail log entries whenever the AWS KMS key for the DB instance is accessed.",
        "description_kind": "plain",
        "type": "string"
      },
      "dedicated_log_volume": {
        "computed": true,
        "description": "Indicates whether the DB instance has a dedicated log volume (DLV) enabled.",
        "description_kind": "plain",
        "type": "bool"
      },
      "delete_automated_backups": {
        "computed": true,
        "description": "A value that indicates whether to remove automated backups immediately after the DB instance is deleted. This parameter isn't case-sensitive. The default is to remove automated backups immediately after the DB instance is deleted.",
        "description_kind": "plain",
        "type": "bool"
      },
      "deletion_protection": {
        "computed": true,
        "description": "A value that indicates whether the DB instance has deletion protection enabled. The database can't be deleted when deletion protection is enabled. By default, deletion protection is disabled.",
        "description_kind": "plain",
        "type": "bool"
      },
      "domain": {
        "computed": true,
        "description": "The Active Directory directory ID to create the DB instance in. Currently, only MySQL, Microsoft SQL Server, Oracle, and PostgreSQL DB instances can be created in an Active Directory Domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_auth_secret_arn": {
        "computed": true,
        "description": "The ARN for the Secrets Manager secret with the credentials for the user joining the domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_dns_ips": {
        "computed": true,
        "description": "The IPv4 DNS IP addresses of your primary and secondary Active Directory domain controllers.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "domain_fqdn": {
        "computed": true,
        "description": "The fully qualified domain name (FQDN) of an Active Directory domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_iam_role_name": {
        "computed": true,
        "description": "Specify the name of the IAM role to be used when making API calls to the Directory Service.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_ou": {
        "computed": true,
        "description": "The Active Directory organizational unit for your DB instance to join.",
        "description_kind": "plain",
        "type": "string"
      },
      "enable_cloudwatch_logs_exports": {
        "computed": true,
        "description": "The list of log types that need to be enabled for exporting to CloudWatch Logs. The values in the list depend on the DB engine being used.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "enable_iam_database_authentication": {
        "computed": true,
        "description": "A value that indicates whether to enable mapping of AWS Identity and Access Management (IAM) accounts to database accounts. By default, mapping is disabled.",
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_performance_insights": {
        "computed": true,
        "description": "A value that indicates whether to enable Performance Insights for the DB instance.",
        "description_kind": "plain",
        "type": "bool"
      },
      "endpoint": {
        "computed": true,
        "description": "Specifies the connection endpoint.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "address": {
              "computed": true,
              "description": "Specifies the DNS address of the DB instance.",
              "description_kind": "plain",
              "type": "string"
            },
            "hosted_zone_id": {
              "computed": true,
              "description": "Specifies the ID that Amazon Route 53 assigns when you create a hosted zone.",
              "description_kind": "plain",
              "type": "string"
            },
            "port": {
              "computed": true,
              "description": "Specifies the port that the database engine is listening on.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "engine": {
        "computed": true,
        "description": "The name of the database engine that you want to use for this DB instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "engine_version": {
        "computed": true,
        "description": "The version number of the database engine to use.",
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
        "description": "The number of I/O operations per second (IOPS) that the database provisions.",
        "description_kind": "plain",
        "type": "number"
      },
      "kms_key_id": {
        "computed": true,
        "description": "The ARN of the AWS Key Management Service (AWS KMS) master key that's used to encrypt the DB instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "license_model": {
        "computed": true,
        "description": "License model information for this DB instance.",
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
        "description": "The password for the master user.",
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
        "description": "The master user name for the DB instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "max_allocated_storage": {
        "computed": true,
        "description": "The upper limit to which Amazon RDS can automatically scale the storage of the DB instance.",
        "description_kind": "plain",
        "type": "number"
      },
      "monitoring_interval": {
        "computed": true,
        "description": "The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance. To disable collecting Enhanced Monitoring metrics, specify 0. The default is 0.",
        "description_kind": "plain",
        "type": "number"
      },
      "monitoring_role_arn": {
        "computed": true,
        "description": "The ARN for the IAM role that permits RDS to send enhanced monitoring metrics to Amazon CloudWatch Logs.",
        "description_kind": "plain",
        "type": "string"
      },
      "multi_az": {
        "computed": true,
        "description": "Specifies whether the database instance is a multiple Availability Zone deployment.",
        "description_kind": "plain",
        "type": "bool"
      },
      "nchar_character_set_name": {
        "computed": true,
        "description": "The name of the NCHAR character set for the Oracle DB instance. This parameter doesn't apply to RDS Custom.",
        "description_kind": "plain",
        "type": "string"
      },
      "network_type": {
        "computed": true,
        "description": "The network type of the DB cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "option_group_name": {
        "computed": true,
        "description": "Indicates that the DB instance should be associated with the specified option group.",
        "description_kind": "plain",
        "type": "string"
      },
      "performance_insights_kms_key_id": {
        "computed": true,
        "description": "The AWS KMS key identifier for encryption of Performance Insights data. The KMS key ID is the Amazon Resource Name (ARN), KMS key identifier, or the KMS key alias for the KMS encryption key.",
        "description_kind": "plain",
        "type": "string"
      },
      "performance_insights_retention_period": {
        "computed": true,
        "description": "The amount of time, in days, to retain Performance Insights data. Valid values are 7 or 731 (2 years).",
        "description_kind": "plain",
        "type": "number"
      },
      "port": {
        "computed": true,
        "description": "The port number on which the database accepts connections.",
        "description_kind": "plain",
        "type": "string"
      },
      "preferred_backup_window": {
        "computed": true,
        "description": "The daily time range during which automated backups are created if automated backups are enabled, using the BackupRetentionPeriod parameter.",
        "description_kind": "plain",
        "type": "string"
      },
      "preferred_maintenance_window": {
        "computed": true,
        "description": "he weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).",
        "description_kind": "plain",
        "type": "string"
      },
      "processor_features": {
        "computed": true,
        "description": "The number of CPU cores and the number of threads per core for the DB instance class of the DB instance.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
              "computed": true,
              "description": "The name of the processor feature. Valid names are coreCount and threadsPerCore.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of a processor feature name.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "promotion_tier": {
        "computed": true,
        "description": "A value that specifies the order in which an Aurora Replica is promoted to the primary instance after a failure of the existing primary instance.",
        "description_kind": "plain",
        "type": "number"
      },
      "publicly_accessible": {
        "computed": true,
        "description": "Indicates whether the DB instance is an internet-facing instance. If you specify true, AWS CloudFormation creates an instance with a publicly resolvable DNS name, which resolves to a public IP address. If you specify false, AWS CloudFormation creates an internal instance with a DNS name that resolves to a private IP address.",
        "description_kind": "plain",
        "type": "bool"
      },
      "replica_mode": {
        "computed": true,
        "description": "The open mode of an Oracle read replica. The default is open-read-only.",
        "description_kind": "plain",
        "type": "string"
      },
      "restore_time": {
        "computed": true,
        "description": "The date and time to restore from.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_db_cluster_identifier": {
        "computed": true,
        "description": "The identifier of the Multi-AZ DB cluster that will act as the source for the read replica. Each DB cluster can have up to 15 read replicas.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_db_instance_automated_backups_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the replicated automated backups from which to restore.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_db_instance_identifier": {
        "computed": true,
        "description": "If you want to create a Read Replica DB instance, specify the ID of the source DB instance. Each DB instance can have a limited number of Read Replicas.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_dbi_resource_id": {
        "computed": true,
        "description": "The resource ID of the source DB instance from which to restore.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_region": {
        "computed": true,
        "description": "The ID of the region that contains the source DB instance for the Read Replica.",
        "description_kind": "plain",
        "type": "string"
      },
      "storage_encrypted": {
        "computed": true,
        "description": "A value that indicates whether the DB instance is encrypted. By default, it isn't encrypted.",
        "description_kind": "plain",
        "type": "bool"
      },
      "storage_throughput": {
        "computed": true,
        "description": "Specifies the storage throughput for the DB instance.",
        "description_kind": "plain",
        "type": "number"
      },
      "storage_type": {
        "computed": true,
        "description": "Specifies the storage type to be associated with the DB instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags to assign to the DB instance.",
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
          "nesting_mode": "list"
        }
      },
      "tde_credential_arn": {
        "computed": true,
        "description": "The ARN from the key store with which to associate the instance for TDE encryption.",
        "description_kind": "plain",
        "type": "string"
      },
      "tde_credential_password": {
        "computed": true,
        "description": "The password for the given ARN from the key store in order to access the device.",
        "description_kind": "plain",
        "type": "string"
      },
      "timezone": {
        "computed": true,
        "description": "The time zone of the DB instance. The time zone parameter is currently supported only by Microsoft SQL Server.",
        "description_kind": "plain",
        "type": "string"
      },
      "use_default_processor_features": {
        "computed": true,
        "description": "A value that indicates whether the DB instance class of the DB instance uses its default processor features.",
        "description_kind": "plain",
        "type": "bool"
      },
      "use_latest_restorable_time": {
        "computed": true,
        "description": "A value that indicates whether the DB instance is restored from the latest backup time. By default, the DB instance isn't restored from the latest backup time.",
        "description_kind": "plain",
        "type": "bool"
      },
      "vpc_security_groups": {
        "computed": true,
        "description": "A list of the VPC security group IDs to assign to the DB instance. The list can include both the physical IDs of existing VPC security groups and references to AWS::EC2::SecurityGroup resources created in the template.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::RDS::DBInstance",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRdsDbInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRdsDbInstance), &result)
	return &result
}
