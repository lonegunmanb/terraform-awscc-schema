package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRdsDbInstance = `{
  "block": {
    "attributes": {
      "allocated_storage": {
        "computed": true,
        "description": "The amount of storage in gibibytes (GiB) to be initially allocated for the database instance.\n  If any value is set in the ` + "`" + `` + "`" + `Iops` + "`" + `` + "`" + ` parameter, ` + "`" + `` + "`" + `AllocatedStorage` + "`" + `` + "`" + ` must be at least 100 GiB, which corresponds to the minimum Iops value of 1,000. If you increase the ` + "`" + `` + "`" + `Iops` + "`" + `` + "`" + ` value (in 1,000 IOPS increments), then you must also increase the ` + "`" + `` + "`" + `AllocatedStorage` + "`" + `` + "`" + ` value (in 100-GiB increments). \n   *Amazon Aurora* \n Not applicable. Aurora cluster volumes automatically grow as the amount of data in your database increases, though you are only charged for the space that you use in an Aurora cluster volume.\n  *Db2* \n Constraints to the amount of storage for each storage type are the following:\n  +  General Purpose (SSD) storage (gp3): Must be an integer from 20 to 64000.\n  +  Provisioned IOPS storage (io1): Must be an integer from 100 to 64000.\n  \n  *MySQL* \n Constraints to the amount of storage for each storage type are the following: \n  +  General Purpose (SSD) storage (gp2): Must be an integer from 20 to 65536.\n  +  Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.\n  +  Magnetic storage (standard): Must be an integer from 5 to 3072.\n  \n  *MariaDB* \n Constraints to the amount of storage for each storage type are the following: \n  +  General Purpose (SSD) storage (gp2): Must be an integer from 20 to 65536.\n  +  Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.\n  +  Magnetic storage (standard): Must be an integer from 5 to 3072.\n  \n  *PostgreSQL* \n Constraints to the amount of storage for each storage type are the following: \n  +  General Purpose (SSD) storage (gp2): Must be an integer from 20 to 65536.\n  +  Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.\n  +  Magnetic storage (standard): Must be an integer from 5 to 3072.\n  \n  *Oracle* \n Constraints to the amount of storage for each storage type are the following: \n  +  General Purpose (SSD) storage (gp2): Must be an integer from 20 to 65536.\n  +  Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.\n  +  Magnetic storage (standard): Must be an integer from 10 to 3072.\n  \n  *SQL Server* \n Constraints to the amount of storage for each storage type are the following: \n  +  General Purpose (SSD) storage (gp2):\n  +  Enterprise and Standard editions: Must be an integer from 20 to 16384.\n  +  Web and Express editions: Must be an integer from 20 to 16384.\n  \n  +  Provisioned IOPS storage (io1):\n  +  Enterprise and Standard editions: Must be an integer from 20 to 16384.\n  +  Web and Express editions: Must be an integer from 20 to 16384.\n  \n  +  Magnetic storage (standard):\n  +  Enterprise and Standard editions: Must be an integer from 20 to 1024.\n  +  Web and Express editions: Must be an integer from 20 to 1024.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "allow_major_version_upgrade": {
        "computed": true,
        "description": "A value that indicates whether major version upgrades are allowed. Changing this parameter doesn't result in an outage and the change is asynchronously applied as soon as possible.\n Constraints: Major version upgrades must be allowed when specifying a value for the ` + "`" + `` + "`" + `EngineVersion` + "`" + `` + "`" + ` parameter that is a different major version than the DB instance's current version.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "associated_roles": {
        "computed": true,
        "description": "The IAMlong (IAM) roles associated with the DB instance. \n  *Amazon Aurora* \n Not applicable. The associated roles are managed by the DB cluster.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "feature_name": {
              "computed": true,
              "description": "The name of the feature associated with the AWS Identity and Access Management (IAM) role. IAM roles that are associated with a DB instance grant permission for the DB instance to access other AWS services on your behalf. For the list of supported feature names, see the ` + "`" + `` + "`" + `SupportedFeatureNames` + "`" + `` + "`" + ` description in [DBEngineVersion](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBEngineVersion.html) in the *Amazon RDS API Reference*.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "role_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the IAM role that is associated with the DB instance.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "auto_minor_version_upgrade": {
        "computed": true,
        "description": "A value that indicates whether minor engine upgrades are applied automatically to the DB instance during the maintenance window. By default, minor engine upgrades are applied automatically.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "automatic_backup_replication_kms_key_id": {
        "computed": true,
        "description": "The AWS KMS key identifier for encryption of the replicated automated backups. The KMS key ID is the Amazon Resource Name (ARN) for the KMS encryption key in the destination AWS-Region, for example, ` + "`" + `` + "`" + `arn:aws:kms:us-east-1:123456789012:key/AKIAIOSFODNN7EXAMPLE` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "automatic_backup_replication_region": {
        "computed": true,
        "description": "The AWS-Region associated with the automated backup.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "availability_zone": {
        "computed": true,
        "description": "The Availability Zone (AZ) where the database will be created. For information on AWS-Regions and Availability Zones, see [Regions and Availability Zones](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.RegionsAndAvailabilityZones.html).\n For Amazon Aurora, each Aurora DB cluster hosts copies of its storage in three separate Availability Zones. Specify one of these Availability Zones. Aurora automatically chooses an appropriate Availability Zone if you don't specify one.\n Default: A random, system-chosen Availability Zone in the endpoint's AWS-Region.\n Constraints:\n  +  The ` + "`" + `` + "`" + `AvailabilityZone` + "`" + `` + "`" + ` parameter can't be specified if the DB instance is a Multi-AZ deployment.\n  +  The specified Availability Zone must be in the same AWS-Region as the current endpoint.\n  \n Example: ` + "`" + `` + "`" + `us-east-1d` + "`" + `` + "`" + `",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "backup_retention_period": {
        "computed": true,
        "description": "The number of days for which automated backups are retained. Setting this parameter to a positive number enables backups. Setting this parameter to 0 disables automated backups.\n  *Amazon Aurora* \n Not applicable. The retention period for automated backups is managed by the DB cluster.\n Default: 1\n Constraints:\n  +  Must be a value from 0 to 35\n  +  Can't be set to 0 if the DB instance is a source to read replicas",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "ca_certificate_identifier": {
        "computed": true,
        "description": "The identifier of the CA certificate for this DB instance.\n For more information, see [Using SSL/TLS to encrypt a connection to a DB instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.SSL.html) in the *Amazon RDS User Guide* and [Using SSL/TLS to encrypt a connection to a DB cluster](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.SSL.html) in the *Amazon Aurora User Guide*.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "certificate_details": {
        "computed": true,
        "description": "The details of the DB instance's server certificate.",
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
              "description": "The expiration date of the DB instance?s server certificate.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "certificate_rotation_restart": {
        "computed": true,
        "description": "Specifies whether the DB instance is restarted when you rotate your SSL/TLS certificate.\n By default, the DB instance is restarted when you rotate your SSL/TLS certificate. The certificate is not updated until the DB instance is restarted.\n  Set this parameter only if you are *not* using SSL/TLS to connect to the DB instance.\n  If you are using SSL/TLS to connect to the DB instance, follow the appropriate instructions for your DB engine to rotate your SSL/TLS certificate:\n  +  For more information about rotating your SSL/TLS certificate for RDS DB engines, see [Rotating Your SSL/TLS Certificate.](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.SSL-certificate-rotation.html) in the *Amazon RDS User Guide.* \n  +  For more information about rotating your SSL/TLS certificate for Aurora DB engines, see [Rotating Your SSL/TLS Certificate](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.SSL-certificate-rotation.html) in the *Amazon Aurora User Guide*.\n  \n This setting doesn't apply to RDS Custom DB instances.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "character_set_name": {
        "computed": true,
        "description": "For supported engines, indicates that the DB instance should be associated with the specified character set.\n  *Amazon Aurora* \n Not applicable. The character set is managed by the DB cluster. For more information, see [AWS::RDS::DBCluster](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "copy_tags_to_snapshot": {
        "computed": true,
        "description": "Specifies whether to copy tags from the DB instance to snapshots of the DB instance. By default, tags are not copied.\n This setting doesn't apply to Amazon Aurora DB instances. Copying tags to snapshots is managed by the DB cluster. Setting this value for an Aurora DB instance has no effect on the DB cluster setting.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "custom_iam_instance_profile": {
        "computed": true,
        "description": "The instance profile associated with the underlying Amazon EC2 instance of an RDS Custom DB instance.\n This setting is required for RDS Custom.\n Constraints:\n  +  The profile must exist in your account.\n  +  The profile must have an IAM role that Amazon EC2 has permissions to assume.\n  +  The instance profile name and the associated IAM role name must start with the prefix ` + "`" + `` + "`" + `AWSRDSCustom` + "`" + `` + "`" + `.\n  \n For the list of permissions required for the IAM role, see [Configure IAM and your VPC](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-setup-orcl.html#custom-setup-orcl.iam-vpc) in the *Amazon RDS User Guide*.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "database_insights_mode": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "db_cluster_identifier": {
        "computed": true,
        "description": "The identifier of the DB cluster that this DB instance will belong to.\n This setting doesn't apply to RDS Custom DB instances.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_cluster_snapshot_identifier": {
        "computed": true,
        "description": "The identifier for the Multi-AZ DB cluster snapshot to restore from.\n For more information on Multi-AZ DB clusters, see [Multi-AZ DB cluster deployments](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html) in the *Amazon RDS User Guide*.\n Constraints:\n  +  Must match the identifier of an existing Multi-AZ DB cluster snapshot.\n  +  Can't be specified when ` + "`" + `` + "`" + `DBSnapshotIdentifier` + "`" + `` + "`" + ` is specified.\n  +  Must be specified when ` + "`" + `` + "`" + `DBSnapshotIdentifier` + "`" + `` + "`" + ` isn't specified.\n  +  If you are restoring from a shared manual Multi-AZ DB cluster snapshot, the ` + "`" + `` + "`" + `DBClusterSnapshotIdentifier` + "`" + `` + "`" + ` must be the ARN of the shared snapshot.\n  +  Can't be the identifier of an Aurora DB cluster snapshot.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_instance_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "db_instance_class": {
        "computed": true,
        "description": "The compute and memory capacity of the DB instance, for example ` + "`" + `` + "`" + `db.m5.large` + "`" + `` + "`" + `. Not all DB instance classes are available in all AWS-Regions, or for all database engines. For the full list of DB instance classes, and availability for your engine, see [DB instance classes](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html) in the *Amazon RDS User Guide* or [Aurora DB instance classes](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.DBInstanceClass.html) in the *Amazon Aurora User Guide*.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_instance_identifier": {
        "computed": true,
        "description": "A name for the DB instance. If you specify a name, AWS CloudFormation converts it to lowercase. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the DB instance. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).\n For information about constraints that apply to DB instance identifiers, see [Naming constraints in Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Limits.html#RDS_Limits.Constraints) in the *Amazon RDS User Guide*.\n  If you specify a name, you can't perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_name": {
        "computed": true,
        "description": "The meaning of this parameter differs according to the database engine you use.\n  If you specify the ` + "`" + `` + "`" + `DBSnapshotIdentifier` + "`" + `` + "`" + ` property, this property only applies to RDS for Oracle.\n   *Amazon Aurora* \n Not applicable. The database name is managed by the DB cluster.\n  *Db2* \n The name of the database to create when the DB instance is created. If this parameter isn't specified, no database is created in the DB instance.\n Constraints:\n  +  Must contain 1 to 64 letters or numbers.\n  +  Must begin with a letter. Subsequent characters can be letters, underscores, or digits (0-9).\n  +  Can't be a word reserved by the specified database engine.\n  \n  *MySQL* \n The name of the database to create when the DB instance is created. If this parameter is not specified, no database is created in the DB instance.\n Constraints:\n  +  Must contain 1 to 64 letters or numbers.\n  +  Can't be a word reserved by the specified database engine\n  \n  *MariaDB* \n The name of the database to create when the DB instance is created. If this parameter is not specified, no database is created in the DB instance.\n Constraints:\n  +  Must contain 1 to 64 letters or numbers.\n  +  Can't be a word reserved by the specified database engine\n  \n  *PostgreSQL* \n The name of the database to create when the DB instance is created. If this parameter is not specified, the default ` + "`" + `` + "`" + `postgres` + "`" + `` + "`" + ` database is created in the DB instance.\n Constraints:\n  +  Must begin with a letter. Subsequent characters can be letters, underscores, or digits (0-9).\n  +  Must contain 1 to 63 characters.\n  +  Can't be a word reserved by the specified database engine\n  \n  *Oracle* \n The Oracle System ID (SID) of the created DB instance. If you specify ` + "`" + `` + "`" + `null` + "`" + `` + "`" + `, the default value ` + "`" + `` + "`" + `ORCL` + "`" + `` + "`" + ` is used. You can't specify the string NULL, or any other reserved word, for ` + "`" + `` + "`" + `DBName` + "`" + `` + "`" + `. \n Default: ` + "`" + `` + "`" + `ORCL` + "`" + `` + "`" + ` \n Constraints:\n  +  Can't be longer than 8 characters\n  \n  *SQL Server* \n Not applicable. Must be null.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_parameter_group_name": {
        "computed": true,
        "description": "The name of an existing DB parameter group or a reference to an [AWS::RDS::DBParameterGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbparametergroup.html) resource created in the template.\n To list all of the available DB parameter group names, use the following command:\n  ` + "`" + `` + "`" + `aws rds describe-db-parameter-groups --query \"DBParameterGroups[].DBParameterGroupName\" --output text` + "`" + `` + "`" + ` \n  If any of the data members of the referenced parameter group are changed during an update, the DB instance might need to be restarted, which causes some interruption. If the parameter group contains static parameters, whether they were changed or not, an update triggers a reboot.\n  If you don't specify a value for ` + "`" + `` + "`" + `DBParameterGroupName` + "`" + `` + "`" + ` property, the default DB parameter group for the specified engine and engine version is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_security_groups": {
        "computed": true,
        "description": "A list of the DB security groups to assign to the DB instance. The list can include both the name of existing DB security groups or references to AWS::RDS::DBSecurityGroup resources created in the template.\n  If you set DBSecurityGroups, you must not set VPCSecurityGroups, and vice versa. Also, note that the DBSecurityGroups property exists only for backwards compatibility with older regions and is no longer recommended for providing security information to an RDS DB instance. Instead, use VPCSecurityGroups.\n  If you specify this property, AWS CloudFormation sends only the following properties (if specified) to Amazon RDS during create operations:\n  +   ` + "`" + `` + "`" + `AllocatedStorage` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `AutoMinorVersionUpgrade` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `AvailabilityZone` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `BackupRetentionPeriod` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `CharacterSetName` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `DBInstanceClass` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `DBName` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `DBParameterGroupName` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `DBSecurityGroups` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `DBSubnetGroupName` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `Engine` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `EngineVersion` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `Iops` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `LicenseModel` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `MasterUsername` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `MasterUserPassword` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `MultiAZ` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `OptionGroupName` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `PreferredBackupWindow` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `PreferredMaintenanceWindow` + "`" + `` + "`" + ` \n  \n All other properties are ignored. Specify a virtual private cloud (VPC) security group if you want to submit other properties, such as ` + "`" + `` + "`" + `StorageType` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `StorageEncrypted` + "`" + `` + "`" + `, or ` + "`" + `` + "`" + `KmsKeyId` + "`" + `` + "`" + `. If you're already using the ` + "`" + `` + "`" + `DBSecurityGroups` + "`" + `` + "`" + ` property, you can't use these other properties by updating your DB instance to use a VPC security group. You must recreate the DB instance.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "db_snapshot_identifier": {
        "computed": true,
        "description": "The name or Amazon Resource Name (ARN) of the DB snapshot that's used to restore the DB instance. If you're restoring from a shared manual DB snapshot, you must specify the ARN of the snapshot.\n By specifying this property, you can create a DB instance from the specified DB snapshot. If the ` + "`" + `` + "`" + `DBSnapshotIdentifier` + "`" + `` + "`" + ` property is an empty string or the ` + "`" + `` + "`" + `AWS::RDS::DBInstance` + "`" + `` + "`" + ` declaration has no ` + "`" + `` + "`" + `DBSnapshotIdentifier` + "`" + `` + "`" + ` property, AWS CloudFormation creates a new database. If the property contains a value (other than an empty string), AWS CloudFormation creates a database from the specified snapshot. If a snapshot with the specified name doesn't exist, AWS CloudFormation can't create the database and it rolls back the stack.\n Some DB instance properties aren't valid when you restore from a snapshot, such as the ` + "`" + `` + "`" + `MasterUsername` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `MasterUserPassword` + "`" + `` + "`" + ` properties, and the point-in-time recovery properties ` + "`" + `` + "`" + `RestoreTime` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `UseLatestRestorableTime` + "`" + `` + "`" + `. For information about the properties that you can specify, see the [RestoreDBInstanceFromDBSnapshot](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_RestoreDBInstanceFromDBSnapshot.html) action in the *Amazon RDS API Reference*.\n After you restore a DB instance with a ` + "`" + `` + "`" + `DBSnapshotIdentifier` + "`" + `` + "`" + ` property, you must specify the same ` + "`" + `` + "`" + `DBSnapshotIdentifier` + "`" + `` + "`" + ` property for any future updates to the DB instance. When you specify this property for an update, the DB instance is not restored from the DB snapshot again, and the data in the database is not changed. However, if you don't specify the ` + "`" + `` + "`" + `DBSnapshotIdentifier` + "`" + `` + "`" + ` property, an empty DB instance is created, and the original DB instance is deleted. If you specify a property that is different from the previous snapshot restore property, a new DB instance is restored from the specified ` + "`" + `` + "`" + `DBSnapshotIdentifier` + "`" + `` + "`" + ` property, and the original DB instance is deleted.\n If you specify the ` + "`" + `` + "`" + `DBSnapshotIdentifier` + "`" + `` + "`" + ` property to restore a DB instance (as opposed to specifying it for DB instance updates), then don't specify the following properties:\n  +   ` + "`" + `` + "`" + `CharacterSetName` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `DBClusterIdentifier` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `DBName` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `KmsKeyId` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `MasterUsername` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `MasterUserPassword` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `PromotionTier` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `SourceDBInstanceIdentifier` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `SourceRegion` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `StorageEncrypted` + "`" + `` + "`" + ` (for an unencrypted snapshot)\n  +   ` + "`" + `` + "`" + `Timezone` + "`" + `` + "`" + ` \n  \n  *Amazon Aurora* \n Not applicable. Snapshot restore is managed by the DB cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_subnet_group_name": {
        "computed": true,
        "description": "A DB subnet group to associate with the DB instance. If you update this value, the new subnet group must be a subnet group in a new VPC. \n If there's no DB subnet group, then the DB instance isn't a VPC DB instance.\n For more information about using Amazon RDS in a VPC, see [Amazon VPC and Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.html) in the *Amazon RDS User Guide*. \n This setting doesn't apply to Amazon Aurora DB instances. The DB subnet group is managed by the DB cluster. If specified, the setting must match the DB cluster setting.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_system_id": {
        "computed": true,
        "description": "The Oracle system identifier (SID), which is the name of the Oracle database instance that manages your database files. In this context, the term \"Oracle database instance\" refers exclusively to the system global area (SGA) and Oracle background processes. If you don't specify a SID, the value defaults to ` + "`" + `` + "`" + `RDSCDB` + "`" + `` + "`" + `. The Oracle SID is also the name of your CDB.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dbi_resource_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dedicated_log_volume": {
        "computed": true,
        "description": "Indicates whether the DB instance has a dedicated log volume (DLV) enabled.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "delete_automated_backups": {
        "computed": true,
        "description": "A value that indicates whether to remove automated backups immediately after the DB instance is deleted. This parameter isn't case-sensitive. The default is to remove automated backups immediately after the DB instance is deleted.\n  *Amazon Aurora* \n Not applicable. When you delete a DB cluster, all automated backups for that DB cluster are deleted and can't be recovered. Manual DB cluster snapshots of the DB cluster are not deleted.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "deletion_protection": {
        "computed": true,
        "description": "Specifies whether the DB instance has deletion protection enabled. The database can't be deleted when deletion protection is enabled. By default, deletion protection isn't enabled. For more information, see [Deleting a DB Instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_DeleteInstance.html).\n This setting doesn't apply to Amazon Aurora DB instances. You can enable or disable deletion protection for the DB cluster. For more information, see ` + "`" + `` + "`" + `CreateDBCluster` + "`" + `` + "`" + `. DB instances in a DB cluster can be deleted even when deletion protection is enabled for the DB cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "domain": {
        "computed": true,
        "description": "The Active Directory directory ID to create the DB instance in. Currently, only Db2, MySQL, Microsoft SQL Server, Oracle, and PostgreSQL DB instances can be created in an Active Directory Domain.\n For more information, see [Kerberos Authentication](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/kerberos-authentication.html) in the *Amazon RDS User Guide*.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "domain_auth_secret_arn": {
        "computed": true,
        "description": "The ARN for the Secrets Manager secret with the credentials for the user joining the domain.\n Example: ` + "`" + `` + "`" + `arn:aws:secretsmanager:region:account-number:secret:myselfmanagedADtestsecret-123456` + "`" + `` + "`" + `",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "domain_dns_ips": {
        "computed": true,
        "description": "The IPv4 DNS IP addresses of your primary and secondary Active Directory domain controllers.\n Constraints:\n  +  Two IP addresses must be provided. If there isn't a secondary domain controller, use the IP address of the primary domain controller for both entries in the list.\n  \n Example: ` + "`" + `` + "`" + `123.124.125.126,234.235.236.237` + "`" + `` + "`" + `",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "domain_fqdn": {
        "computed": true,
        "description": "The fully qualified domain name (FQDN) of an Active Directory domain.\n Constraints:\n  +  Can't be longer than 64 characters.\n  \n Example: ` + "`" + `` + "`" + `mymanagedADtest.mymanagedAD.mydomain` + "`" + `` + "`" + `",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "domain_iam_role_name": {
        "computed": true,
        "description": "The name of the IAM role to use when making API calls to the Directory Service.\n This setting doesn't apply to the following DB instances:\n  +  Amazon Aurora (The domain is managed by the DB cluster.)\n  +  RDS Custom",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "domain_ou": {
        "computed": true,
        "description": "The Active Directory organizational unit for your DB instance to join.\n Constraints:\n  +  Must be in the distinguished name format.\n  +  Can't be longer than 64 characters.\n  \n Example: ` + "`" + `` + "`" + `OU=mymanagedADtestOU,DC=mymanagedADtest,DC=mymanagedAD,DC=mydomain` + "`" + `` + "`" + `",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_cloudwatch_logs_exports": {
        "computed": true,
        "description": "The list of log types that need to be enabled for exporting to CloudWatch Logs. The values in the list depend on the DB engine being used. For more information, see [Publishing Database Logs to Amazon CloudWatch Logs](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch) in the *Amazon Relational Database Service User Guide*.\n  *Amazon Aurora* \n Not applicable. CloudWatch Logs exports are managed by the DB cluster. \n  *Db2* \n Valid values: ` + "`" + `` + "`" + `diag.log` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `notify.log` + "`" + `` + "`" + ` \n  *MariaDB* \n Valid values: ` + "`" + `` + "`" + `audit` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `error` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `general` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `slowquery` + "`" + `` + "`" + ` \n  *Microsoft SQL Server* \n Valid values: ` + "`" + `` + "`" + `agent` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `error` + "`" + `` + "`" + ` \n  *MySQL* \n Valid values: ` + "`" + `` + "`" + `audit` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `error` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `general` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `slowquery` + "`" + `` + "`" + ` \n  *Oracle* \n Valid values: ` + "`" + `` + "`" + `alert` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `audit` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `listener` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `trace` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `oemagent` + "`" + `` + "`" + ` \n  *PostgreSQL* \n Valid values: ` + "`" + `` + "`" + `postgresql` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `upgrade` + "`" + `` + "`" + `",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "enable_iam_database_authentication": {
        "computed": true,
        "description": "A value that indicates whether to enable mapping of AWS Identity and Access Management (IAM) accounts to database accounts. By default, mapping is disabled.\n This property is supported for RDS for MariaDB, RDS for MySQL, and RDS for PostgreSQL. For more information, see [IAM Database Authentication for MariaDB, MySQL, and PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html) in the *Amazon RDS User Guide.* \n  *Amazon Aurora* \n Not applicable. Mapping AWS IAM accounts to database accounts is managed by the DB cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_performance_insights": {
        "computed": true,
        "description": "Specifies whether to enable Performance Insights for the DB instance. For more information, see [Using Amazon Performance Insights](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.html) in the *Amazon RDS User Guide*.\n This setting doesn't apply to RDS Custom DB instances.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "endpoint": {
        "computed": true,
        "description": "The connection endpoint for the DB instance.\n  The endpoint might not be shown for instances with the status of ` + "`" + `` + "`" + `creating` + "`" + `` + "`" + `.",
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
        },
        "optional": true
      },
      "engine": {
        "computed": true,
        "description": "The name of the database engine to use for this DB instance. Not every database engine is available in every AWS Region.\n This property is required when creating a DB instance.\n  You can convert an Oracle database from the non-CDB architecture to the container database (CDB) architecture by updating the ` + "`" + `` + "`" + `Engine` + "`" + `` + "`" + ` value in your templates from ` + "`" + `` + "`" + `oracle-ee` + "`" + `` + "`" + ` to ` + "`" + `` + "`" + `oracle-ee-cdb` + "`" + `` + "`" + ` or from ` + "`" + `` + "`" + `oracle-se2` + "`" + `` + "`" + ` to ` + "`" + `` + "`" + `oracle-se2-cdb` + "`" + `` + "`" + `. Converting to the CDB architecture requires an interruption.\n  Valid Values:\n  +   ` + "`" + `` + "`" + `aurora-mysql` + "`" + `` + "`" + ` (for Aurora MySQL DB instances)\n  +   ` + "`" + `` + "`" + `aurora-postgresql` + "`" + `` + "`" + ` (for Aurora PostgreSQL DB instances)\n  +   ` + "`" + `` + "`" + `custom-oracle-ee` + "`" + `` + "`" + ` (for RDS Custom for Oracle DB instances)\n  +   ` + "`" + `` + "`" + `custom-oracle-ee-cdb` + "`" + `` + "`" + ` (for RDS Custom for Oracle DB instances)\n  +   ` + "`" + `` + "`" + `custom-sqlserver-ee` + "`" + `` + "`" + ` (for RDS Custom for SQL Server DB instances)\n  +   ` + "`" + `` + "`" + `custom-sqlserver-se` + "`" + `` + "`" + ` (for RDS Custom for SQL Server DB instances)\n  +   ` + "`" + `` + "`" + `custom-sqlserver-web` + "`" + `` + "`" + ` (for RDS Custom for SQL Server DB instances)\n  +   ` + "`" + `` + "`" + `db2-ae` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `db2-se` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `mariadb` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `mysql` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `oracle-ee` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `oracle-ee-cdb` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `oracle-se2` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `oracle-se2-cdb` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `postgres` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `sqlserver-ee` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `sqlserver-se` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `sqlserver-ex` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `sqlserver-web` + "`" + `` + "`" + `",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "engine_lifecycle_support": {
        "computed": true,
        "description": "The life cycle type for this DB instance.\n  By default, this value is set to ` + "`" + `` + "`" + `open-source-rds-extended-support` + "`" + `` + "`" + `, which enrolls your DB instance into Amazon RDS Extended Support. At the end of standard support, you can avoid charges for Extended Support by setting the value to ` + "`" + `` + "`" + `open-source-rds-extended-support-disabled` + "`" + `` + "`" + `. In this case, creating the DB instance will fail if the DB major version is past its end of standard support date.\n  This setting applies only to RDS for MySQL and RDS for PostgreSQL. For Amazon Aurora DB instances, the life cycle type is managed by the DB cluster.\n You can use this setting to enroll your DB instance into Amazon RDS Extended Support. With RDS Extended Support, you can run the selected major engine version on your DB instance past the end of standard support for that engine version. For more information, see [Using Amazon RDS Extended Support](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/extended-support.html) in the *Amazon RDS User Guide*.\n Valid Values: ` + "`" + `` + "`" + `open-source-rds-extended-support | open-source-rds-extended-support-disabled` + "`" + `` + "`" + ` \n Default: ` + "`" + `` + "`" + `open-source-rds-extended-support` + "`" + `` + "`" + `",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "engine_version": {
        "computed": true,
        "description": "The version number of the database engine to use.\n For a list of valid engine versions, use the ` + "`" + `` + "`" + `DescribeDBEngineVersions` + "`" + `` + "`" + ` action.\n The following are the database engines and links to information about the major and minor versions that are available with Amazon RDS. Not every database engine is available for every AWS Region.\n  *Amazon Aurora* \n Not applicable. The version number of the database engine to be used by the DB instance is managed by the DB cluster.\n  *Db2* \n See [Amazon RDS for Db2](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Db2.html#Db2.Concepts.VersionMgmt) in the *Amazon RDS User Guide.* \n  *MariaDB* \n See [MariaDB on Amazon RDS Versions](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_MariaDB.html#MariaDB.Concepts.VersionMgmt) in the *Amazon RDS User Guide.* \n  *Microsoft SQL Server* \n See [Microsoft SQL Server Versions on Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SQLServer.html#SQLServer.Concepts.General.VersionSupport) in the *Amazon RDS User Guide.* \n  *MySQL* \n See [MySQL on Amazon RDS Versions](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_MySQL.html#MySQL.Concepts.VersionMgmt) in the *Amazon RDS User Guide.* \n  *Oracle* \n See [Oracle Database Engine Release Notes](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.Oracle.PatchComposition.html) in the *Amazon RDS User Guide.* \n  *PostgreSQL* \n See [Supported PostgreSQL Database Versions](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_PostgreSQL.html#PostgreSQL.Concepts.General.DBVersions) in the *Amazon RDS User Guide.*",
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
      "iops": {
        "computed": true,
        "description": "The number of I/O operations per second (IOPS) that the database provisions. The value must be equal to or greater than 1000. \n If you specify this property, you must follow the range of allowed ratios of your requested IOPS rate to the amount of storage that you allocate (IOPS to allocated storage). For example, you can provision an Oracle database instance with 1000 IOPS and 200 GiB of storage (a ratio of 5:1), or specify 2000 IOPS with 200 GiB of storage (a ratio of 10:1). For more information, see [Amazon RDS Provisioned IOPS Storage to Improve Performance](https://docs.aws.amazon.com/AmazonRDS/latest/DeveloperGuide/CHAP_Storage.html#USER_PIOPS) in the *Amazon RDS User Guide*.\n  If you specify ` + "`" + `` + "`" + `io1` + "`" + `` + "`" + ` for the ` + "`" + `` + "`" + `StorageType` + "`" + `` + "`" + ` property, then you must also specify the ` + "`" + `` + "`" + `Iops` + "`" + `` + "`" + ` property.\n  Constraints:\n  +  For RDS for Db2, MariaDB, MySQL, Oracle, and PostgreSQL - Must be a multiple between .5 and 50 of the storage amount for the DB instance.\n  +  For RDS for SQL Server - Must be a multiple between 1 and 50 of the storage amount for the DB instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "kms_key_id": {
        "computed": true,
        "description": "The ARN of the AWS KMS key that's used to encrypt the DB instance, such as ` + "`" + `` + "`" + `arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef` + "`" + `` + "`" + `. If you enable the StorageEncrypted property but don't specify this property, AWS CloudFormation uses the default KMS key. If you specify this property, you must set the StorageEncrypted property to true. \n If you specify the ` + "`" + `` + "`" + `SourceDBInstanceIdentifier` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `SourceDbiResourceId` + "`" + `` + "`" + ` property, don't specify this property. The value is inherited from the source DB instance, and if the DB instance is encrypted, the specified ` + "`" + `` + "`" + `KmsKeyId` + "`" + `` + "`" + ` property is used. However, if the source DB instance is in a different AWS Region, you must specify a KMS key ID.\n If you specify the ` + "`" + `` + "`" + `SourceDBInstanceAutomatedBackupsArn` + "`" + `` + "`" + ` property, don't specify this property. The value is inherited from the source DB instance automated backup, and if the automated backup is encrypted, the specified ` + "`" + `` + "`" + `KmsKeyId` + "`" + `` + "`" + ` property is used.\n If you create an encrypted read replica in a different AWS Region, then you must specify a KMS key for the destination AWS Region. KMS encryption keys are specific to the region that they're created in, and you can't use encryption keys from one region in another region.\n If you specify the ` + "`" + `` + "`" + `DBSnapshotIdentifier` + "`" + `` + "`" + ` property, don't specify this property. The ` + "`" + `` + "`" + `StorageEncrypted` + "`" + `` + "`" + ` property value is inherited from the snapshot. If the DB instance is encrypted, the specified ` + "`" + `` + "`" + `KmsKeyId` + "`" + `` + "`" + ` property is also inherited from the snapshot.\n If you specify ` + "`" + `` + "`" + `DBSecurityGroups` + "`" + `` + "`" + `, AWS CloudFormation ignores this property. To specify both a security group and this property, you must use a VPC security group. For more information about Amazon RDS and VPC, see [Using Amazon RDS with Amazon VPC](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.html) in the *Amazon RDS User Guide*.\n  *Amazon Aurora* \n Not applicable. The KMS key identifier is managed by the DB cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "license_model": {
        "computed": true,
        "description": "License model information for this DB instance.\n  Valid Values:\n  +  Aurora MySQL - ` + "`" + `` + "`" + `general-public-license` + "`" + `` + "`" + ` \n  +  Aurora PostgreSQL - ` + "`" + `` + "`" + `postgresql-license` + "`" + `` + "`" + ` \n  +  RDS for Db2 - ` + "`" + `` + "`" + `bring-your-own-license` + "`" + `` + "`" + `. For more information about RDS for Db2 licensing, see [](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-licensing.html) in the *Amazon RDS User Guide.* \n  +  RDS for MariaDB - ` + "`" + `` + "`" + `general-public-license` + "`" + `` + "`" + ` \n  +  RDS for Microsoft SQL Server - ` + "`" + `` + "`" + `license-included` + "`" + `` + "`" + ` \n  +  RDS for MySQL - ` + "`" + `` + "`" + `general-public-license` + "`" + `` + "`" + ` \n  +  RDS for Oracle - ` + "`" + `` + "`" + `bring-your-own-license` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `license-included` + "`" + `` + "`" + ` \n  +  RDS for PostgreSQL - ` + "`" + `` + "`" + `postgresql-license` + "`" + `` + "`" + ` \n  \n  If you've specified ` + "`" + `` + "`" + `DBSecurityGroups` + "`" + `` + "`" + ` and then you update the license model, AWS CloudFormation replaces the underlying DB instance. This will incur some interruptions to database availability.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "manage_master_user_password": {
        "computed": true,
        "description": "Specifies whether to manage the master user password with AWS Secrets Manager.\n For more information, see [Password management with Secrets Manager](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-secrets-manager.html) in the *Amazon RDS User Guide.* \n Constraints:\n  +  Can't manage the master user password with AWS Secrets Manager if ` + "`" + `` + "`" + `MasterUserPassword` + "`" + `` + "`" + ` is specified.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "master_user_password": {
        "computed": true,
        "description": "The password for the master user. The password can include any printable ASCII character except \"/\", \"\"\", or \"@\".\n  *Amazon Aurora* \n Not applicable. The password for the master user is managed by the DB cluster.\n  *RDS for Db2* \n Must contain from 8 to 255 characters.\n  *RDS for MariaDB* \n Constraints: Must contain from 8 to 41 characters.\n  *RDS for Microsoft SQL Server* \n Constraints: Must contain from 8 to 128 characters.\n  *RDS for MySQL* \n Constraints: Must contain from 8 to 41 characters.\n  *RDS for Oracle* \n Constraints: Must contain from 8 to 30 characters.\n  *RDS for PostgreSQL* \n Constraints: Must contain from 8 to 128 characters.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "master_user_secret": {
        "computed": true,
        "description": "The secret managed by RDS in AWS Secrets Manager for the master user password.\n For more information, see [Password management with Secrets Manager](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-secrets-manager.html) in the *Amazon RDS User Guide.*",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kms_key_id": {
              "computed": true,
              "description": "The AWS KMS key identifier that is used to encrypt the secret.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "secret_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the secret. This parameter is a return value that you can retrieve using the ` + "`" + `` + "`" + `Fn::GetAtt` + "`" + `` + "`" + ` intrinsic function. For more information, see [Return values](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbinstance.html#aws-resource-rds-dbinstance-return-values).",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "master_username": {
        "computed": true,
        "description": "The master user name for the DB instance.\n  If you specify the ` + "`" + `` + "`" + `SourceDBInstanceIdentifier` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `DBSnapshotIdentifier` + "`" + `` + "`" + ` property, don't specify this property. The value is inherited from the source DB instance or snapshot.\n When migrating a self-managed Db2 database, we recommend that you use the same master username as your self-managed Db2 instance name.\n   *Amazon Aurora* \n Not applicable. The name for the master user is managed by the DB cluster. \n  *RDS for Db2* \n Constraints:\n  +  Must be 1 to 16 letters or numbers.\n  +  First character must be a letter.\n  +  Can't be a reserved word for the chosen database engine.\n  \n  *RDS for MariaDB* \n Constraints:\n  +  Must be 1 to 16 letters or numbers.\n  +  Can't be a reserved word for the chosen database engine.\n  \n  *RDS for Microsoft SQL Server* \n Constraints:\n  +  Must be 1 to 128 letters or numbers.\n  +  First character must be a letter.\n  +  Can't be a reserved word for the chosen database engine.\n  \n  *RDS for MySQL* \n Constraints:\n  +  Must be 1 to 16 letters or numbers.\n  +  First character must be a letter.\n  +  Can't be a reserved word for the chosen database engine.\n  \n  *RDS for Oracle* \n Constraints:\n  +  Must be 1 to 30 letters or numbers.\n  +  First character must be a letter.\n  +  Can't be a reserved word for the chosen database engine.\n  \n  *RDS for PostgreSQL* \n Constraints:\n  +  Must be 1 to 63 letters or numbers.\n  +  First character must be a letter.\n  +  Can't be a reserved word for the chosen database engine.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "max_allocated_storage": {
        "computed": true,
        "description": "The upper limit in gibibytes (GiB) to which Amazon RDS can automatically scale the storage of the DB instance.\n For more information about this setting, including limitations that apply to it, see [Managing capacity automatically with Amazon RDS storage autoscaling](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIOPS.StorageTypes.html#USER_PIOPS.Autoscaling) in the *Amazon RDS User Guide*.\n This setting doesn't apply to the following DB instances:\n  +  Amazon Aurora (Storage is managed by the DB cluster.)\n  +  RDS Custom",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "monitoring_interval": {
        "computed": true,
        "description": "The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance. To disable collection of Enhanced Monitoring metrics, specify ` + "`" + `` + "`" + `0` + "`" + `` + "`" + `.\n If ` + "`" + `` + "`" + `MonitoringRoleArn` + "`" + `` + "`" + ` is specified, then you must set ` + "`" + `` + "`" + `MonitoringInterval` + "`" + `` + "`" + ` to a value other than ` + "`" + `` + "`" + `0` + "`" + `` + "`" + `.\n This setting doesn't apply to RDS Custom DB instances.\n Valid Values: ` + "`" + `` + "`" + `0 | 1 | 5 | 10 | 15 | 30 | 60` + "`" + `` + "`" + ` \n Default: ` + "`" + `` + "`" + `0` + "`" + `` + "`" + `",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "monitoring_role_arn": {
        "computed": true,
        "description": "The ARN for the IAM role that permits RDS to send enhanced monitoring metrics to Amazon CloudWatch Logs. For example, ` + "`" + `` + "`" + `arn:aws:iam:123456789012:role/emaccess` + "`" + `` + "`" + `. For information on creating a monitoring role, see [Setting Up and Enabling Enhanced Monitoring](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.OS.html#USER_Monitoring.OS.Enabling) in the *Amazon RDS User Guide*.\n If ` + "`" + `` + "`" + `MonitoringInterval` + "`" + `` + "`" + ` is set to a value other than ` + "`" + `` + "`" + `0` + "`" + `` + "`" + `, then you must supply a ` + "`" + `` + "`" + `MonitoringRoleArn` + "`" + `` + "`" + ` value.\n This setting doesn't apply to RDS Custom DB instances.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "multi_az": {
        "computed": true,
        "description": "Specifies whether the DB instance is a Multi-AZ deployment. You can't set the ` + "`" + `` + "`" + `AvailabilityZone` + "`" + `` + "`" + ` parameter if the DB instance is a Multi-AZ deployment.\n This setting doesn't apply to the following DB instances:\n  +  Amazon Aurora (DB instance Availability Zones (AZs) are managed by the DB cluster.)\n  +  RDS Custom",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "nchar_character_set_name": {
        "computed": true,
        "description": "The name of the NCHAR character set for the Oracle DB instance.\n This setting doesn't apply to RDS Custom DB instances.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "network_type": {
        "computed": true,
        "description": "The network type of the DB instance.\n Valid values:\n  +   ` + "`" + `` + "`" + `IPV4` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `DUAL` + "`" + `` + "`" + ` \n  \n The network type is determined by the ` + "`" + `` + "`" + `DBSubnetGroup` + "`" + `` + "`" + ` specified for the DB instance. A ` + "`" + `` + "`" + `DBSubnetGroup` + "`" + `` + "`" + ` can support only the IPv4 protocol or the IPv4 and IPv6 protocols (` + "`" + `` + "`" + `DUAL` + "`" + `` + "`" + `).\n For more information, see [Working with a DB instance in a VPC](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.WorkingWithRDSInstanceinaVPC.html) in the *Amazon RDS User Guide.*",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "option_group_name": {
        "computed": true,
        "description": "Indicates that the DB instance should be associated with the specified option group.\n Permanent options, such as the TDE option for Oracle Advanced Security TDE, can't be removed from an option group. Also, that option group can't be removed from a DB instance once it is associated with a DB instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "performance_insights_kms_key_id": {
        "computed": true,
        "description": "The AWS KMS key identifier for encryption of Performance Insights data.\n The KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key.\n If you do not specify a value for ` + "`" + `` + "`" + `PerformanceInsightsKMSKeyId` + "`" + `` + "`" + `, then Amazon RDS uses your default KMS key. There is a default KMS key for your AWS account. Your AWS account has a different default KMS key for each AWS Region.\n For information about enabling Performance Insights, see [EnablePerformanceInsights](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-enableperformanceinsights).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "performance_insights_retention_period": {
        "computed": true,
        "description": "The number of days to retain Performance Insights data.\n This setting doesn't apply to RDS Custom DB instances.\n Valid Values:\n  +   ` + "`" + `` + "`" + `7` + "`" + `` + "`" + ` \n  +   *month* * 31, where *month* is a number of months from 1-23. Examples: ` + "`" + `` + "`" + `93` + "`" + `` + "`" + ` (3 months * 31), ` + "`" + `` + "`" + `341` + "`" + `` + "`" + ` (11 months * 31), ` + "`" + `` + "`" + `589` + "`" + `` + "`" + ` (19 months * 31)\n  +   ` + "`" + `` + "`" + `731` + "`" + `` + "`" + ` \n  \n Default: ` + "`" + `` + "`" + `7` + "`" + `` + "`" + ` days\n If you specify a retention period that isn't valid, such as ` + "`" + `` + "`" + `94` + "`" + `` + "`" + `, Amazon RDS returns an error.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "port": {
        "computed": true,
        "description": "The port number on which the database accepts connections.\n This setting doesn't apply to Aurora DB instances. The port number is managed by the cluster.\n Valid Values: ` + "`" + `` + "`" + `1150-65535` + "`" + `` + "`" + ` \n Default:\n  +  RDS for Db2 - ` + "`" + `` + "`" + `50000` + "`" + `` + "`" + ` \n  +  RDS for MariaDB - ` + "`" + `` + "`" + `3306` + "`" + `` + "`" + ` \n  +  RDS for Microsoft SQL Server - ` + "`" + `` + "`" + `1433` + "`" + `` + "`" + ` \n  +  RDS for MySQL - ` + "`" + `` + "`" + `3306` + "`" + `` + "`" + ` \n  +  RDS for Oracle - ` + "`" + `` + "`" + `1521` + "`" + `` + "`" + ` \n  +  RDS for PostgreSQL - ` + "`" + `` + "`" + `5432` + "`" + `` + "`" + ` \n  \n Constraints:\n  +  For RDS for Microsoft SQL Server, the value can't be ` + "`" + `` + "`" + `1234` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `1434` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `3260` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `3343` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `3389` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `47001` + "`" + `` + "`" + `, or ` + "`" + `` + "`" + `49152-49156` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "preferred_backup_window": {
        "computed": true,
        "description": "The daily time range during which automated backups are created if automated backups are enabled, using the ` + "`" + `` + "`" + `BackupRetentionPeriod` + "`" + `` + "`" + ` parameter. For more information, see [Backup Window](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithAutomatedBackups.html#USER_WorkingWithAutomatedBackups.BackupWindow) in the *Amazon RDS User Guide.* \n Constraints:\n  +  Must be in the format ` + "`" + `` + "`" + `hh24:mi-hh24:mi` + "`" + `` + "`" + `.\n  +  Must be in Universal Coordinated Time (UTC).\n  +  Must not conflict with the preferred maintenance window.\n  +  Must be at least 30 minutes.\n  \n  *Amazon Aurora* \n Not applicable. The daily time range for creating automated backups is managed by the DB cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "preferred_maintenance_window": {
        "computed": true,
        "description": "The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).\n Format: ` + "`" + `` + "`" + `ddd:hh24:mi-ddd:hh24:mi` + "`" + `` + "`" + ` \n The default is a 30-minute window selected at random from an 8-hour block of time for each AWS Region, occurring on a random day of the week. To see the time blocks available, see [Maintaining a DB instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.Maintenance.html#AdjustingTheMaintenanceWindow) in the *Amazon RDS User Guide.* \n  This property applies when AWS CloudFormation initially creates the DB instance. If you use AWS CloudFormation to update the DB instance, those updates are applied immediately.\n  Constraints: Minimum 30-minute window.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "processor_features": {
        "computed": true,
        "description": "The number of CPU cores and the number of threads per core for the DB instance class of the DB instance.\n This setting doesn't apply to Amazon Aurora or RDS Custom DB instances.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
              "computed": true,
              "description": "The name of the processor feature. Valid names are ` + "`" + `` + "`" + `coreCount` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `threadsPerCore` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of a processor feature.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "promotion_tier": {
        "computed": true,
        "description": "The order of priority in which an Aurora Replica is promoted to the primary instance after a failure of the existing primary instance. For more information, see [Fault Tolerance for an Aurora DB Cluster](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.AuroraHighAvailability.html#Aurora.Managing.FaultTolerance) in the *Amazon Aurora User Guide*.\n This setting doesn't apply to RDS Custom DB instances.\n Default: ` + "`" + `` + "`" + `1` + "`" + `` + "`" + ` \n Valid Values: ` + "`" + `` + "`" + `0 - 15` + "`" + `` + "`" + `",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "publicly_accessible": {
        "computed": true,
        "description": "Indicates whether the DB instance is an internet-facing instance. If you specify true, AWS CloudFormation creates an instance with a publicly resolvable DNS name, which resolves to a public IP address. If you specify false, AWS CloudFormation creates an internal instance with a DNS name that resolves to a private IP address. \n The default behavior value depends on your VPC setup and the database subnet group. For more information, see the ` + "`" + `` + "`" + `PubliclyAccessible` + "`" + `` + "`" + ` parameter in the [CreateDBInstance](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html) in the *Amazon RDS API Reference*.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "replica_mode": {
        "computed": true,
        "description": "The open mode of an Oracle read replica. For more information, see [Working with Oracle Read Replicas for Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-read-replicas.html) in the *Amazon RDS User Guide*.\n This setting is only supported in RDS for Oracle.\n Default: ` + "`" + `` + "`" + `open-read-only` + "`" + `` + "`" + ` \n Valid Values: ` + "`" + `` + "`" + `open-read-only` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `mounted` + "`" + `` + "`" + `",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "restore_time": {
        "computed": true,
        "description": "The date and time to restore from. This parameter applies to point-in-time recovery. For more information, see [Restoring a DB instance to a specified time](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIT.html) in the in the *Amazon RDS User Guide*.\n Constraints:\n  +  Must be a time in Universal Coordinated Time (UTC) format.\n  +  Must be before the latest restorable time for the DB instance.\n  +  Can't be specified if the ` + "`" + `` + "`" + `UseLatestRestorableTime` + "`" + `` + "`" + ` parameter is enabled.\n  \n Example: ` + "`" + `` + "`" + `2009-09-07T23:45:00Z` + "`" + `` + "`" + `",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_db_cluster_identifier": {
        "computed": true,
        "description": "The identifier of the Multi-AZ DB cluster that will act as the source for the read replica. Each DB cluster can have up to 15 read replicas.\n Constraints:\n  +  Must be the identifier of an existing Multi-AZ DB cluster.\n  +  Can't be specified if the ` + "`" + `` + "`" + `SourceDBInstanceIdentifier` + "`" + `` + "`" + ` parameter is also specified.\n  +  The specified DB cluster must have automatic backups enabled, that is, its backup retention period must be greater than 0.\n  +  The source DB cluster must be in the same AWS-Region as the read replica. Cross-Region replication isn't supported.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_db_instance_automated_backups_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the replicated automated backups from which to restore, for example, ` + "`" + `` + "`" + `arn:aws:rds:us-east-1:123456789012:auto-backup:ab-L2IJCEXJP7XQ7HOJ4SIEXAMPLE` + "`" + `` + "`" + `.\n This setting doesn't apply to RDS Custom.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_db_instance_identifier": {
        "computed": true,
        "description": "If you want to create a read replica DB instance, specify the ID of the source DB instance. Each DB instance can have a limited number of read replicas. For more information, see [Working with Read Replicas](https://docs.aws.amazon.com/AmazonRDS/latest/DeveloperGuide/USER_ReadRepl.html) in the *Amazon RDS User Guide*.\n For information about constraints that apply to DB instance identifiers, see [Naming constraints in Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Limits.html#RDS_Limits.Constraints) in the *Amazon RDS User Guide*.\n The ` + "`" + `` + "`" + `SourceDBInstanceIdentifier` + "`" + `` + "`" + ` property determines whether a DB instance is a read replica. If you remove the ` + "`" + `` + "`" + `SourceDBInstanceIdentifier` + "`" + `` + "`" + ` property from your template and then update your stack, AWS CloudFormation promotes the read replica to a standalone DB instance.\n If you specify the ` + "`" + `` + "`" + `UseLatestRestorableTime` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `RestoreTime` + "`" + `` + "`" + ` properties in conjunction with the ` + "`" + `` + "`" + `SourceDBInstanceIdentifier` + "`" + `` + "`" + ` property, RDS restores the DB instance to the requested point in time, thereby creating a new DB instance.\n   +  If you specify a source DB instance that uses VPC security groups, we recommend that you specify the ` + "`" + `` + "`" + `VPCSecurityGroups` + "`" + `` + "`" + ` property. If you don't specify the property, the read replica inherits the value of the ` + "`" + `` + "`" + `VPCSecurityGroups` + "`" + `` + "`" + ` property from the source DB when you create the replica. However, if you update the stack, AWS CloudFormation reverts the replica's ` + "`" + `` + "`" + `VPCSecurityGroups` + "`" + `` + "`" + ` property to the default value because it's not defined in the stack's template. This change might cause unexpected issues.\n  +  Read replicas don't support deletion policies. AWS CloudFormation ignores any deletion policy that's associated with a read replica.\n  +  If you specify ` + "`" + `` + "`" + `SourceDBInstanceIdentifier` + "`" + `` + "`" + `, don't specify the ` + "`" + `` + "`" + `DBSnapshotIdentifier` + "`" + `` + "`" + ` property. You can't create a read replica from a snapshot.\n  +  Don't set the ` + "`" + `` + "`" + `BackupRetentionPeriod` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `DBName` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `MasterUsername` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `MasterUserPassword` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `PreferredBackupWindow` + "`" + `` + "`" + ` properties. The database attributes are inherited from the source DB instance, and backups are disabled for read replicas.\n  +  If the source DB instance is in a different region than the read replica, specify the source region in ` + "`" + `` + "`" + `SourceRegion` + "`" + `` + "`" + `, and specify an ARN for a valid DB instance in ` + "`" + `` + "`" + `SourceDBInstanceIdentifier` + "`" + `` + "`" + `. For more information, see [Constructing a Amazon RDS Amazon Resource Name (ARN)](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html#USER_Tagging.ARN) in the *Amazon RDS User Guide*.\n  +  For DB instances in Amazon Aurora clusters, don't specify this property. Amazon RDS automatically assigns writer and reader DB instances.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_dbi_resource_id": {
        "computed": true,
        "description": "The resource ID of the source DB instance from which to restore.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_region": {
        "computed": true,
        "description": "The ID of the region that contains the source DB instance for the read replica.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "storage_encrypted": {
        "computed": true,
        "description": "A value that indicates whether the DB instance is encrypted. By default, it isn't encrypted.\n If you specify the ` + "`" + `` + "`" + `KmsKeyId` + "`" + `` + "`" + ` property, then you must enable encryption.\n If you specify the ` + "`" + `` + "`" + `SourceDBInstanceIdentifier` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `SourceDbiResourceId` + "`" + `` + "`" + ` property, don't specify this property. The value is inherited from the source DB instance, and if the DB instance is encrypted, the specified ` + "`" + `` + "`" + `KmsKeyId` + "`" + `` + "`" + ` property is used.\n If you specify the ` + "`" + `` + "`" + `SourceDBInstanceAutomatedBackupsArn` + "`" + `` + "`" + ` property, don't specify this property. The value is inherited from the source DB instance automated backup. \n If you specify ` + "`" + `` + "`" + `DBSnapshotIdentifier` + "`" + `` + "`" + ` property, don't specify this property. The value is inherited from the snapshot.\n  *Amazon Aurora* \n Not applicable. The encryption for DB instances is managed by the DB cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "storage_throughput": {
        "computed": true,
        "description": "Specifies the storage throughput value, in mebibyte per second (MiBps), for the DB instance. This setting applies only to the ` + "`" + `` + "`" + `gp3` + "`" + `` + "`" + ` storage type. \n This setting doesn't apply to RDS Custom or Amazon Aurora.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "storage_type": {
        "computed": true,
        "description": "The storage type to associate with the DB instance.\n If you specify ` + "`" + `` + "`" + `io1` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `io2` + "`" + `` + "`" + `, or ` + "`" + `` + "`" + `gp3` + "`" + `` + "`" + `, you must also include a value for the ` + "`" + `` + "`" + `Iops` + "`" + `` + "`" + ` parameter.\n This setting doesn't apply to Amazon Aurora DB instances. Storage is managed by the DB cluster.\n Valid Values: ` + "`" + `` + "`" + `gp2 | gp3 | io1 | io2 | standard` + "`" + `` + "`" + ` \n Default: ` + "`" + `` + "`" + `io1` + "`" + `` + "`" + `, if the ` + "`" + `` + "`" + `Iops` + "`" + `` + "`" + ` parameter is specified. Otherwise, ` + "`" + `` + "`" + `gp2` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
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
              "description": "A key is the required name of the tag. The string value can be from 1 to 128 Unicode characters in length and can't be prefixed with ` + "`" + `` + "`" + `aws:` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `rds:` + "`" + `` + "`" + `. The string can only contain only the set of Unicode letters, digits, white-space, '_', '.', ':', '/', '=', '+', '-', '@' (Java regex: \"^([\\\\p{L}\\\\p{Z}\\\\p{N}_.:/=+\\\\-@]*)$\").",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "A value is the optional value of the tag. The string value can be from 1 to 256 Unicode characters in length and can't be prefixed with ` + "`" + `` + "`" + `aws:` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `rds:` + "`" + `` + "`" + `. The string can only contain only the set of Unicode letters, digits, white-space, '_', '.', ':', '/', '=', '+', '-', '@' (Java regex: \"^([\\\\p{L}\\\\p{Z}\\\\p{N}_.:/=+\\\\-@]*)$\").",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "tde_credential_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tde_credential_password": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "timezone": {
        "computed": true,
        "description": "The time zone of the DB instance. The time zone parameter is currently supported only by [RDS for Db2](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-time-zone) and [RDS for SQL Server](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SQLServer.html#SQLServer.Concepts.General.TimeZone).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "use_default_processor_features": {
        "computed": true,
        "description": "Specifies whether the DB instance class of the DB instance uses its default processor features.\n This setting doesn't apply to RDS Custom DB instances.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "use_latest_restorable_time": {
        "computed": true,
        "description": "Specifies whether the DB instance is restored from the latest backup time. By default, the DB instance isn't restored from the latest backup time. This parameter applies to point-in-time recovery. For more information, see [Restoring a DB instance to a specified time](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIT.html) in the in the *Amazon RDS User Guide*.\n Constraints:\n  +  Can't be specified if the ` + "`" + `` + "`" + `RestoreTime` + "`" + `` + "`" + ` parameter is provided.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "vpc_security_groups": {
        "computed": true,
        "description": "A list of the VPC security group IDs to assign to the DB instance. The list can include both the physical IDs of existing VPC security groups and references to [AWS::EC2::SecurityGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group.html) resources created in the template.\n If you plan to update the resource, don't specify VPC security groups in a shared VPC.\n  If you set ` + "`" + `` + "`" + `VPCSecurityGroups` + "`" + `` + "`" + `, you must not set [DBSecurityGroups](https://docs.aws.amazon.com//AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-dbsecuritygroups), and vice versa.\n  You can migrate a DB instance in your stack from an RDS DB security group to a VPC security group, but keep the following in mind:\n  +  You can't revert to using an RDS security group after you establish a VPC security group membership.\n  +  When you migrate your DB instance to VPC security groups, if your stack update rolls back because the DB instance update fails or because an update fails in another AWS CloudFormation resource, the rollback fails because it can't revert to an RDS security group.\n  +  To use the properties that are available when you use a VPC security group, you must recreate the DB instance. If you don't, AWS CloudFormation submits only the property values that are listed in the [DBSecurityGroups](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-dbsecuritygroups) property.\n  \n  To avoid this situation, migrate your DB instance to using VPC security groups only when that is the only change in your stack template. \n  *Amazon Aurora* \n Not applicable. The associated list of EC2 VPC security groups is managed by the DB cluster. If specified, the setting must match the DB cluster setting.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "The ` + "`" + `` + "`" + `AWS::RDS::DBInstance` + "`" + `` + "`" + ` resource creates an Amazon DB instance. The new DB instance can be an RDS DB instance, or it can be a DB instance in an Aurora DB cluster.\n For more information about creating an RDS DB instance, see [Creating an Amazon RDS DB instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_CreateDBInstance.html) in the *Amazon RDS User Guide*.\n For more information about creating a DB instance in an Aurora DB cluster, see [Creating an Amazon Aurora DB cluster](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.CreateInstance.html) in the *Amazon Aurora User Guide*.\n If you import an existing DB instance, and the template configuration doesn't match the actual configuration of the DB instance, AWS CloudFormation applies the changes in the template during the import operation.\n  If a DB instance is deleted or replaced during an update, AWS CloudFormation deletes all automated snapshots. However, it retains manual DB snapshots. During an update that requires replacement, you can apply a stack policy to prevent DB instances from being replaced. For more information, see [Prevent Updates to Stack Resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/protect-stack-resources.html).\n   *Updating DB instances* \n When properties labeled \"*Update requires:* [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)\" are updated, AWS CloudFormation first creates a replacement DB instance, then changes references from other dependent resources to point to the replacement DB instance, and finally deletes the old DB instance.\n  We highly recommend that you take a snapshot of the database before updating the stack. If you don't, you lose the data when AWS CloudFormation replaces your DB instance. To preserve your data, perform the following procedure:\n  1.  Deactivate any applications that are using the DB instance so that there's no activity on the DB instance.\n  1.  Create a snapshot of the DB instance. For more information, see [Creating a DB Snapshot](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_CreateSnapshot.html).\n  1.  If you want to restore your instance using a DB snapshot, modify the updated template with your DB instance changes and add the ` + "`" + `` + "`" + `DBSnapshotIdentifier` + "`" + `` + "`" + ` property with the ID of the DB snapshot that you want to use.\n After you restore a DB instance with a ` + "`" + `` + "`" + `DBSnapshotIdentifier` + "`" + `` + "`" + ` property, you can delete the ` + "`" + `` + "`" + `DBSnapshotIdentifier` + "`" + `` + "`" + ` property. When you specify this property for an update, the DB instance is not restored from the DB snapshot again, and the data in the database is not changed. However, if you don't specify the ` + "`" + `` + "`" + `DBSnapshotIdentifier` + "`" + `` + "`" + ` property, an empty DB instance is created, and the original DB instance is deleted. If you specify a property that is different from the previous snapshot restore property, a new DB instance is restored from the specified ` + "`" + `` + "`" + `DBSnapshotIdentifier` + "`" + `` + "`" + ` property, and the original DB instance is deleted.\n  1.  Update the stack.\n  \n  For more information about updating other properties of this resource, see ` + "`" + `` + "`" + `ModifyDBInstance` + "`" + `` + "`" + `. For more information about updating stacks, see [CloudFormation Stacks Updates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks.html).\n  *Deleting DB instances* \n For DB instances that are part of an Aurora DB cluster, you can set a deletion policy for your DB instance to control how AWS CloudFormation handles the DB instance when the stack is deleted. For Amazon RDS DB instances, you can choose to *retain* the DB instance, to *delete* the DB instance, or to *create a snapshot* of the DB instance. The default AWS CloudFormation behavior depends on the ` + "`" + `` + "`" + `DBClusterIdentifier` + "`" + `` + "`" + ` property:\n  1.  For ` + "`" + `` + "`" + `AWS::RDS::DBInstance` + "`" + `` + "`" + ` resources that don't specify the ` + "`" + `` + "`" + `DBClusterIdentifier` + "`" + `` + "`" + ` property, AWS CloudFormation saves a snapshot of the DB instance.\n  1.   For ` + "`" + `` + "`" + `AWS::RDS::DBInstance` + "`" + `` + "`" + ` resources that do specify the ` + "`" + `` + "`" + `DBClusterIdentifier` + "`" + `` + "`" + ` property, AWS CloudFormation deletes the DB instance.\n  \n  For more information, see [DeletionPolicy Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html).",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRdsDbInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRdsDbInstance), &result)
	return &result
}
