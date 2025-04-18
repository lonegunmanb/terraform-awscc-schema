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
        "description": "The amount of storage in gibibytes (GiB) to allocate to each DB instance in the Multi-AZ DB cluster.\n Valid for Cluster Type: Multi-AZ DB clusters only\n This setting is required to create a Multi-AZ DB cluster.",
        "description_kind": "plain",
        "type": "number"
      },
      "associated_roles": {
        "computed": true,
        "description": "Provides a list of the AWS Identity and Access Management (IAM) roles that are associated with the DB cluster. IAM roles that are associated with a DB cluster grant permission for the DB cluster to access other Amazon Web Services on your behalf.\n Valid for: Aurora DB clusters and Multi-AZ DB clusters",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "feature_name": {
              "computed": true,
              "description": "The name of the feature associated with the AWS Identity and Access Management (IAM) role. IAM roles that are associated with a DB cluster grant permission for the DB cluster to access other AWS services on your behalf. For the list of supported feature names, see the ` + "`" + `` + "`" + `SupportedFeatureNames` + "`" + `` + "`" + ` description in [DBEngineVersion](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBEngineVersion.html) in the *Amazon RDS API Reference*.",
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
        "description": "Specifies whether minor engine upgrades are applied automatically to the DB cluster during the maintenance window. By default, minor engine upgrades are applied automatically.\n Valid for Cluster Type: Aurora DB clusters and Multi-AZ DB cluster",
        "description_kind": "plain",
        "type": "bool"
      },
      "availability_zones": {
        "computed": true,
        "description": "A list of Availability Zones (AZs) where instances in the DB cluster can be created. For information on AWS Regions and Availability Zones, see [Choosing the Regions and Availability Zones](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.RegionsAndAvailabilityZones.html) in the *Amazon Aurora User Guide*. \n Valid for: Aurora DB clusters only",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "backtrack_window": {
        "computed": true,
        "description": "The target backtrack window, in seconds. To disable backtracking, set this value to ` + "`" + `` + "`" + `0` + "`" + `` + "`" + `.\n Valid for Cluster Type: Aurora MySQL DB clusters only\n Default: ` + "`" + `` + "`" + `0` + "`" + `` + "`" + ` \n Constraints:\n  +  If specified, this value must be set to a number from 0 to 259,200 (72 hours).",
        "description_kind": "plain",
        "type": "number"
      },
      "backup_retention_period": {
        "computed": true,
        "description": "The number of days for which automated backups are retained.\n Default: 1\n Constraints:\n  +  Must be a value from 1 to 35\n  \n Valid for: Aurora DB clusters and Multi-AZ DB clusters",
        "description_kind": "plain",
        "type": "number"
      },
      "cluster_scalability_type": {
        "computed": true,
        "description": "Specifies the scalability mode of the Aurora DB cluster. When set to ` + "`" + `` + "`" + `limitless` + "`" + `` + "`" + `, the cluster operates as an Aurora Limitless Database, allowing you to create a DB shard group for horizontal scaling (sharding) capabilities. When set to ` + "`" + `` + "`" + `standard` + "`" + `` + "`" + ` (the default), the cluster uses normal DB instance creation.",
        "description_kind": "plain",
        "type": "string"
      },
      "copy_tags_to_snapshot": {
        "computed": true,
        "description": "A value that indicates whether to copy all tags from the DB cluster to snapshots of the DB cluster. The default is not to copy them.\n Valid for: Aurora DB clusters and Multi-AZ DB clusters",
        "description_kind": "plain",
        "type": "bool"
      },
      "database_insights_mode": {
        "computed": true,
        "description": "The mode of Database Insights to enable for the DB cluster.\n If you set this value to ` + "`" + `` + "`" + `advanced` + "`" + `` + "`" + `, you must also set the ` + "`" + `` + "`" + `PerformanceInsightsEnabled` + "`" + `` + "`" + ` parameter to ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` and the ` + "`" + `` + "`" + `PerformanceInsightsRetentionPeriod` + "`" + `` + "`" + ` parameter to 465.\n Valid for Cluster Type: Aurora DB clusters and Multi-AZ DB clusters",
        "description_kind": "plain",
        "type": "string"
      },
      "database_name": {
        "computed": true,
        "description": "The name of your database. If you don't provide a name, then Amazon RDS won't create a database in this DB cluster. For naming constraints, see [Naming Constraints](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_Limits.html#RDS_Limits.Constraints) in the *Amazon Aurora User Guide*. \n Valid for: Aurora DB clusters and Multi-AZ DB clusters",
        "description_kind": "plain",
        "type": "string"
      },
      "db_cluster_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "db_cluster_identifier": {
        "computed": true,
        "description": "The DB cluster identifier. This parameter is stored as a lowercase string.\n Constraints:\n  +  Must contain from 1 to 63 letters, numbers, or hyphens.\n  +  First character must be a letter.\n  +  Can't end with a hyphen or contain two consecutive hyphens.\n  \n Example: ` + "`" + `` + "`" + `my-cluster1` + "`" + `` + "`" + ` \n Valid for: Aurora DB clusters and Multi-AZ DB clusters",
        "description_kind": "plain",
        "type": "string"
      },
      "db_cluster_instance_class": {
        "computed": true,
        "description": "The compute and memory capacity of each DB instance in the Multi-AZ DB cluster, for example ` + "`" + `` + "`" + `db.m6gd.xlarge` + "`" + `` + "`" + `. Not all DB instance classes are available in all AWS-Regions, or for all database engines.\n For the full list of DB instance classes and availability for your engine, see [DB instance class](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html) in the *Amazon RDS User Guide*.\n This setting is required to create a Multi-AZ DB cluster.\n Valid for Cluster Type: Multi-AZ DB clusters only",
        "description_kind": "plain",
        "type": "string"
      },
      "db_cluster_parameter_group_name": {
        "computed": true,
        "description": "The name of the DB cluster parameter group to associate with this DB cluster.\n  If you apply a parameter group to an existing DB cluster, then its DB instances might need to reboot. This can result in an outage while the DB instances are rebooting.\n If you apply a change to parameter group associated with a stopped DB cluster, then the update stack waits until the DB cluster is started.\n  To list all of the available DB cluster parameter group names, use the following command:\n  ` + "`" + `` + "`" + `aws rds describe-db-cluster-parameter-groups --query \"DBClusterParameterGroups[].DBClusterParameterGroupName\" --output text` + "`" + `` + "`" + ` \n Valid for: Aurora DB clusters and Multi-AZ DB clusters",
        "description_kind": "plain",
        "type": "string"
      },
      "db_cluster_resource_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "db_instance_parameter_group_name": {
        "computed": true,
        "description": "The name of the DB parameter group to apply to all instances of the DB cluster.\n  When you apply a parameter group using the ` + "`" + `` + "`" + `DBInstanceParameterGroupName` + "`" + `` + "`" + ` parameter, the DB cluster isn't rebooted automatically. Also, parameter changes are applied immediately rather than during the next maintenance window.\n  Valid for Cluster Type: Aurora DB clusters only\n Default: The existing name setting\n Constraints:\n  +  The DB parameter group must be in the same DB parameter group family as this DB cluster.\n  +  The ` + "`" + `` + "`" + `DBInstanceParameterGroupName` + "`" + `` + "`" + ` parameter is valid in combination with the ` + "`" + `` + "`" + `AllowMajorVersionUpgrade` + "`" + `` + "`" + ` parameter for a major version upgrade only.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_subnet_group_name": {
        "computed": true,
        "description": "A DB subnet group that you want to associate with this DB cluster. \n If you are restoring a DB cluster to a point in time with ` + "`" + `` + "`" + `RestoreType` + "`" + `` + "`" + ` set to ` + "`" + `` + "`" + `copy-on-write` + "`" + `` + "`" + `, and don't specify a DB subnet group name, then the DB cluster is restored with a default DB subnet group.\n Valid for: Aurora DB clusters and Multi-AZ DB clusters",
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
        "description": "A value that indicates whether the DB cluster has deletion protection enabled. The database can't be deleted when deletion protection is enabled. By default, deletion protection is disabled.\n Valid for: Aurora DB clusters and Multi-AZ DB clusters",
        "description_kind": "plain",
        "type": "bool"
      },
      "domain": {
        "computed": true,
        "description": "Indicates the directory ID of the Active Directory to create the DB cluster.\n For Amazon Aurora DB clusters, Amazon RDS can use Kerberos authentication to authenticate users that connect to the DB cluster.\n For more information, see [Kerberos authentication](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/kerberos-authentication.html) in the *Amazon Aurora User Guide*.\n Valid for: Aurora DB clusters only",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_iam_role_name": {
        "computed": true,
        "description": "Specifies the name of the IAM role to use when making API calls to the Directory Service.\n Valid for: Aurora DB clusters only",
        "description_kind": "plain",
        "type": "string"
      },
      "enable_cloudwatch_logs_exports": {
        "computed": true,
        "description": "The list of log types that need to be enabled for exporting to CloudWatch Logs. The values in the list depend on the DB engine being used. For more information, see [Publishing Database Logs to Amazon CloudWatch Logs](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch) in the *Amazon Aurora User Guide*.\n  *Aurora MySQL* \n Valid values: ` + "`" + `` + "`" + `audit` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `error` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `general` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `slowquery` + "`" + `` + "`" + ` \n  *Aurora PostgreSQL* \n Valid values: ` + "`" + `` + "`" + `postgresql` + "`" + `` + "`" + ` \n Valid for: Aurora DB clusters and Multi-AZ DB clusters",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "enable_global_write_forwarding": {
        "computed": true,
        "description": "Specifies whether to enable this DB cluster to forward write operations to the primary cluster of a global cluster (Aurora global database). By default, write operations are not allowed on Aurora DB clusters that are secondary clusters in an Aurora global database.\n You can set this value only on Aurora DB clusters that are members of an Aurora global database. With this parameter enabled, a secondary cluster can forward writes to the current primary cluster, and the resulting changes are replicated back to this cluster. For the primary DB cluster of an Aurora global database, this value is used immediately if the primary is demoted by a global cluster API operation, but it does nothing until then.\n Valid for Cluster Type: Aurora DB clusters only",
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_http_endpoint": {
        "computed": true,
        "description": "Specifies whether to enable the HTTP endpoint for the DB cluster. By default, the HTTP endpoint isn't enabled.\n When enabled, the HTTP endpoint provides a connectionless web service API (RDS Data API) for running SQL queries on the DB cluster. You can also query your database from inside the RDS console with the RDS query editor.\n For more information, see [Using RDS Data API](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/data-api.html) in the *Amazon Aurora User Guide*.\n Valid for Cluster Type: Aurora DB clusters only",
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_iam_database_authentication": {
        "computed": true,
        "description": "A value that indicates whether to enable mapping of AWS Identity and Access Management (IAM) accounts to database accounts. By default, mapping is disabled.\n For more information, see [IAM Database Authentication](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.IAMDBAuth.html) in the *Amazon Aurora User Guide.* \n Valid for: Aurora DB clusters only",
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_local_write_forwarding": {
        "computed": true,
        "description": "Specifies whether read replicas can forward write operations to the writer DB instance in the DB cluster. By default, write operations aren't allowed on reader DB instances.\n Valid for: Aurora DB clusters only",
        "description_kind": "plain",
        "type": "bool"
      },
      "endpoint": {
        "computed": true,
        "description": "The ` + "`" + `` + "`" + `Endpoint` + "`" + `` + "`" + ` return value specifies the connection endpoint for the primary instance of the DB cluster.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "address": {
              "computed": true,
              "description": "Specifies the connection endpoint for the primary instance of the DB cluster.",
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
        "description": "The name of the database engine to be used for this DB cluster.\n Valid Values:\n  +   ` + "`" + `` + "`" + `aurora-mysql` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `aurora-postgresql` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `mysql` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `postgres` + "`" + `` + "`" + ` \n  \n Valid for: Aurora DB clusters and Multi-AZ DB clusters",
        "description_kind": "plain",
        "type": "string"
      },
      "engine_lifecycle_support": {
        "computed": true,
        "description": "The life cycle type for this DB cluster.\n  By default, this value is set to ` + "`" + `` + "`" + `open-source-rds-extended-support` + "`" + `` + "`" + `, which enrolls your DB cluster into Amazon RDS Extended Support. At the end of standard support, you can avoid charges for Extended Support by setting the value to ` + "`" + `` + "`" + `open-source-rds-extended-support-disabled` + "`" + `` + "`" + `. In this case, creating the DB cluster will fail if the DB major version is past its end of standard support date.\n  You can use this setting to enroll your DB cluster into Amazon RDS Extended Support. With RDS Extended Support, you can run the selected major engine version on your DB cluster past the end of standard support for that engine version. For more information, see the following sections:\n  +  Amazon Aurora - [Using Amazon RDS Extended Support](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/extended-support.html) in the *Amazon Aurora User Guide* \n  +  Amazon RDS - [Using Amazon RDS Extended Support](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/extended-support.html) in the *Amazon RDS User Guide* \n  \n Valid for Cluster Type: Aurora DB clusters and Multi-AZ DB clusters\n Valid Values: ` + "`" + `` + "`" + `open-source-rds-extended-support | open-source-rds-extended-support-disabled` + "`" + `` + "`" + ` \n Default: ` + "`" + `` + "`" + `open-source-rds-extended-support` + "`" + `` + "`" + `",
        "description_kind": "plain",
        "type": "string"
      },
      "engine_mode": {
        "computed": true,
        "description": "The DB engine mode of the DB cluster, either ` + "`" + `` + "`" + `provisioned` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `serverless` + "`" + `` + "`" + `.\n The ` + "`" + `` + "`" + `serverless` + "`" + `` + "`" + ` engine mode only applies for Aurora Serverless v1 DB clusters. Aurora Serverless v2 DB clusters use the ` + "`" + `` + "`" + `provisioned` + "`" + `` + "`" + ` engine mode.\n For information about limitations and requirements for Serverless DB clusters, see the following sections in the *Amazon Aurora User Guide*:\n  +   [Limitations of Aurora Serverless v1](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.html#aurora-serverless.limitations) \n  +   [Requirements for Aurora Serverless v2](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless-v2.requirements.html) \n  \n Valid for Cluster Type: Aurora DB clusters only",
        "description_kind": "plain",
        "type": "string"
      },
      "engine_version": {
        "computed": true,
        "description": "The version number of the database engine to use.\n To list all of the available engine versions for Aurora MySQL version 2 (5.7-compatible) and version 3 (8.0-compatible), use the following command:\n  ` + "`" + `` + "`" + `aws rds describe-db-engine-versions --engine aurora-mysql --query \"DBEngineVersions[].EngineVersion\"` + "`" + `` + "`" + ` \n You can supply either ` + "`" + `` + "`" + `5.7` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `8.0` + "`" + `` + "`" + ` to use the default engine version for Aurora MySQL version 2 or version 3, respectively.\n To list all of the available engine versions for Aurora PostgreSQL, use the following command:\n  ` + "`" + `` + "`" + `aws rds describe-db-engine-versions --engine aurora-postgresql --query \"DBEngineVersions[].EngineVersion\"` + "`" + `` + "`" + ` \n To list all of the available engine versions for RDS for MySQL, use the following command:\n  ` + "`" + `` + "`" + `aws rds describe-db-engine-versions --engine mysql --query \"DBEngineVersions[].EngineVersion\"` + "`" + `` + "`" + ` \n To list all of the available engine versions for RDS for PostgreSQL, use the following command:\n  ` + "`" + `` + "`" + `aws rds describe-db-engine-versions --engine postgres --query \"DBEngineVersions[].EngineVersion\"` + "`" + `` + "`" + ` \n  *Aurora MySQL* \n For information, see [Database engine updates for Amazon Aurora MySQL](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Updates.html) in the *Amazon Aurora User Guide*.\n  *Aurora PostgreSQL* \n For information, see [Amazon Aurora PostgreSQL releases and engine versions](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Updates.20180305.html) in the *Amazon Aurora User Guide*.\n  *MySQL* \n For information, see [Amazon RDS for MySQL](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_MySQL.html#MySQL.Concepts.VersionMgmt) in the *Amazon RDS User Guide*.\n  *PostgreSQL* \n For information, see [Amazon RDS for PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_PostgreSQL.html#PostgreSQL.Concepts) in the *Amazon RDS User Guide*.\n Valid for: Aurora DB clusters and Multi-AZ DB clusters",
        "description_kind": "plain",
        "type": "string"
      },
      "global_cluster_identifier": {
        "computed": true,
        "description": "If you are configuring an Aurora global database cluster and want your Aurora DB cluster to be a secondary member in the global database cluster, specify the global cluster ID of the global database cluster. To define the primary database cluster of the global cluster, use the [AWS::RDS::GlobalCluster](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-globalcluster.html) resource. \n  If you aren't configuring a global database cluster, don't specify this property. \n  To remove the DB cluster from a global database cluster, specify an empty value for the ` + "`" + `` + "`" + `GlobalClusterIdentifier` + "`" + `` + "`" + ` property.\n  For information about Aurora global databases, see [Working with Amazon Aurora Global Databases](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database.html) in the *Amazon Aurora User Guide*.\n Valid for: Aurora DB clusters only",
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
        "description": "The amount of Provisioned IOPS (input/output operations per second) to be initially allocated for each DB instance in the Multi-AZ DB cluster.\n For information about valid IOPS values, see [Provisioned IOPS storage](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html#USER_PIOPS) in the *Amazon RDS User Guide*.\n This setting is required to create a Multi-AZ DB cluster.\n Valid for Cluster Type: Multi-AZ DB clusters only\n Constraints:\n  +  Must be a multiple between .5 and 50 of the storage amount for the DB cluster.",
        "description_kind": "plain",
        "type": "number"
      },
      "kms_key_id": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the AWS KMS key that is used to encrypt the database instances in the DB cluster, such as ` + "`" + `` + "`" + `arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef` + "`" + `` + "`" + `. If you enable the ` + "`" + `` + "`" + `StorageEncrypted` + "`" + `` + "`" + ` property but don't specify this property, the default KMS key is used. If you specify this property, you must set the ` + "`" + `` + "`" + `StorageEncrypted` + "`" + `` + "`" + ` property to ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `.\n If you specify the ` + "`" + `` + "`" + `SnapshotIdentifier` + "`" + `` + "`" + ` property, the ` + "`" + `` + "`" + `StorageEncrypted` + "`" + `` + "`" + ` property value is inherited from the snapshot, and if the DB cluster is encrypted, the specified ` + "`" + `` + "`" + `KmsKeyId` + "`" + `` + "`" + ` property is used.\n If you create a read replica of an encrypted DB cluster in another AWS Region, make sure to set ` + "`" + `` + "`" + `KmsKeyId` + "`" + `` + "`" + ` to a KMS key identifier that is valid in the destination AWS Region. This KMS key is used to encrypt the read replica in that AWS Region.\n Valid for: Aurora DB clusters and Multi-AZ DB clusters",
        "description_kind": "plain",
        "type": "string"
      },
      "manage_master_user_password": {
        "computed": true,
        "description": "Specifies whether to manage the master user password with AWS Secrets Manager.\n For more information, see [Password management with Secrets Manager](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-secrets-manager.html) in the *Amazon RDS User Guide* and [Password management with Secrets Manager](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/rds-secrets-manager.html) in the *Amazon Aurora User Guide.* \n Valid for Cluster Type: Aurora DB clusters and Multi-AZ DB clusters\n Constraints:\n  +  Can't manage the master user password with AWS Secrets Manager if ` + "`" + `` + "`" + `MasterUserPassword` + "`" + `` + "`" + ` is specified.",
        "description_kind": "plain",
        "type": "bool"
      },
      "master_user_password": {
        "computed": true,
        "description": "The master password for the DB instance.\n  If you specify the ` + "`" + `` + "`" + `SourceDBClusterIdentifier` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `SnapshotIdentifier` + "`" + `` + "`" + `, or ` + "`" + `` + "`" + `GlobalClusterIdentifier` + "`" + `` + "`" + ` property, don't specify this property. The value is inherited from the source DB cluster, the snapshot, or the primary DB cluster for the global database cluster, respectively.\n  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
        "description_kind": "plain",
        "type": "string"
      },
      "master_user_secret": {
        "computed": true,
        "description": "The secret managed by RDS in AWS Secrets Manager for the master user password.\n  When you restore a DB cluster from a snapshot, Amazon RDS generates a new secret instead of reusing the secret specified in the ` + "`" + `` + "`" + `SecretArn` + "`" + `` + "`" + ` property. This ensures that the restored DB cluster is securely managed with a dedicated secret. To maintain consistent integration with your application, you might need to update resource configurations to reference the newly created secret.\n  For more information, see [Password management with Secrets Manager](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-secrets-manager.html) in the *Amazon RDS User Guide* and [Password management with Secrets Manager](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/rds-secrets-manager.html) in the *Amazon Aurora User Guide.*",
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
              "description": "The Amazon Resource Name (ARN) of the secret. This parameter is a return value that you can retrieve using the ` + "`" + `` + "`" + `Fn::GetAtt` + "`" + `` + "`" + ` intrinsic function. For more information, see [Return values](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#aws-resource-rds-dbcluster-return-values).",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "master_username": {
        "computed": true,
        "description": "The name of the master user for the DB cluster.\n  If you specify the ` + "`" + `` + "`" + `SourceDBClusterIdentifier` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `SnapshotIdentifier` + "`" + `` + "`" + `, or ` + "`" + `` + "`" + `GlobalClusterIdentifier` + "`" + `` + "`" + ` property, don't specify this property. The value is inherited from the source DB cluster, the snapshot, or the primary DB cluster for the global database cluster, respectively.\n  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
        "description_kind": "plain",
        "type": "string"
      },
      "monitoring_interval": {
        "computed": true,
        "description": "The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB cluster. To turn off collecting Enhanced Monitoring metrics, specify ` + "`" + `` + "`" + `0` + "`" + `` + "`" + `.\n If ` + "`" + `` + "`" + `MonitoringRoleArn` + "`" + `` + "`" + ` is specified, also set ` + "`" + `` + "`" + `MonitoringInterval` + "`" + `` + "`" + ` to a value other than ` + "`" + `` + "`" + `0` + "`" + `` + "`" + `.\n Valid for Cluster Type: Aurora DB clusters and Multi-AZ DB clusters\n Valid Values: ` + "`" + `` + "`" + `0 | 1 | 5 | 10 | 15 | 30 | 60` + "`" + `` + "`" + ` \n Default: ` + "`" + `` + "`" + `0` + "`" + `` + "`" + `",
        "description_kind": "plain",
        "type": "number"
      },
      "monitoring_role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the IAM role that permits RDS to send Enhanced Monitoring metrics to Amazon CloudWatch Logs. An example is ` + "`" + `` + "`" + `arn:aws:iam:123456789012:role/emaccess` + "`" + `` + "`" + `. For information on creating a monitoring role, see [Setting up and enabling Enhanced Monitoring](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.OS.html#USER_Monitoring.OS.Enabling) in the *Amazon RDS User Guide*.\n If ` + "`" + `` + "`" + `MonitoringInterval` + "`" + `` + "`" + ` is set to a value other than ` + "`" + `` + "`" + `0` + "`" + `` + "`" + `, supply a ` + "`" + `` + "`" + `MonitoringRoleArn` + "`" + `` + "`" + ` value.\n Valid for Cluster Type: Aurora DB clusters and Multi-AZ DB clusters",
        "description_kind": "plain",
        "type": "string"
      },
      "network_type": {
        "computed": true,
        "description": "The network type of the DB cluster.\n Valid values:\n  +   ` + "`" + `` + "`" + `IPV4` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `DUAL` + "`" + `` + "`" + ` \n  \n The network type is determined by the ` + "`" + `` + "`" + `DBSubnetGroup` + "`" + `` + "`" + ` specified for the DB cluster. A ` + "`" + `` + "`" + `DBSubnetGroup` + "`" + `` + "`" + ` can support only the IPv4 protocol or the IPv4 and IPv6 protocols (` + "`" + `` + "`" + `DUAL` + "`" + `` + "`" + `).\n For more information, see [Working with a DB instance in a VPC](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_VPC.WorkingWithRDSInstanceinaVPC.html) in the *Amazon Aurora User Guide.* \n Valid for: Aurora DB clusters only",
        "description_kind": "plain",
        "type": "string"
      },
      "performance_insights_enabled": {
        "computed": true,
        "description": "Specifies whether to turn on Performance Insights for the DB cluster.\n For more information, see [Using Amazon Performance Insights](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.html) in the *Amazon RDS User Guide*.\n Valid for Cluster Type: Aurora DB clusters and Multi-AZ DB clusters",
        "description_kind": "plain",
        "type": "bool"
      },
      "performance_insights_kms_key_id": {
        "computed": true,
        "description": "The AWS KMS key identifier for encryption of Performance Insights data.\n The AWS KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key.\n If you don't specify a value for ` + "`" + `` + "`" + `PerformanceInsightsKMSKeyId` + "`" + `` + "`" + `, then Amazon RDS uses your default KMS key. There is a default KMS key for your AWS-account. Your AWS-account has a different default KMS key for each AWS-Region.\n Valid for Cluster Type: Aurora DB clusters and Multi-AZ DB clusters",
        "description_kind": "plain",
        "type": "string"
      },
      "performance_insights_retention_period": {
        "computed": true,
        "description": "The number of days to retain Performance Insights data. When creating a DB cluster without enabling Performance Insights, you can't specify the parameter ` + "`" + `` + "`" + `PerformanceInsightsRetentionPeriod` + "`" + `` + "`" + `.\n Valid for Cluster Type: Aurora DB clusters and Multi-AZ DB clusters\n Valid Values:\n  +   ` + "`" + `` + "`" + `7` + "`" + `` + "`" + ` \n  +   *month* * 31, where *month* is a number of months from 1-23. Examples: ` + "`" + `` + "`" + `93` + "`" + `` + "`" + ` (3 months * 31), ` + "`" + `` + "`" + `341` + "`" + `` + "`" + ` (11 months * 31), ` + "`" + `` + "`" + `589` + "`" + `` + "`" + ` (19 months * 31)\n  +   ` + "`" + `` + "`" + `731` + "`" + `` + "`" + ` \n  \n Default: ` + "`" + `` + "`" + `7` + "`" + `` + "`" + ` days\n If you specify a retention period that isn't valid, such as ` + "`" + `` + "`" + `94` + "`" + `` + "`" + `, Amazon RDS issues an error.",
        "description_kind": "plain",
        "type": "number"
      },
      "port": {
        "computed": true,
        "description": "The port number on which the DB instances in the DB cluster accept connections.\n Default:\n  +  When ` + "`" + `` + "`" + `EngineMode` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `provisioned` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `3306` + "`" + `` + "`" + ` (for both Aurora MySQL and Aurora PostgreSQL)\n  +  When ` + "`" + `` + "`" + `EngineMode` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `serverless` + "`" + `` + "`" + `:\n  +   ` + "`" + `` + "`" + `3306` + "`" + `` + "`" + ` when ` + "`" + `` + "`" + `Engine` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `aurora` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `aurora-mysql` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `5432` + "`" + `` + "`" + ` when ` + "`" + `` + "`" + `Engine` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `aurora-postgresql` + "`" + `` + "`" + ` \n  \n  \n  The ` + "`" + `` + "`" + `No interruption` + "`" + `` + "`" + ` on update behavior only applies to DB clusters. If you are updating a DB instance, see [Port](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-port) for the AWS::RDS::DBInstance resource.\n  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
        "description_kind": "plain",
        "type": "number"
      },
      "preferred_backup_window": {
        "computed": true,
        "description": "The daily time range during which automated backups are created. For more information, see [Backup Window](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Managing.Backups.html#Aurora.Managing.Backups.BackupWindow) in the *Amazon Aurora User Guide.* \n Constraints:\n  +  Must be in the format ` + "`" + `` + "`" + `hh24:mi-hh24:mi` + "`" + `` + "`" + `.\n  +  Must be in Universal Coordinated Time (UTC).\n  +  Must not conflict with the preferred maintenance window.\n  +  Must be at least 30 minutes.\n  \n Valid for: Aurora DB clusters and Multi-AZ DB clusters",
        "description_kind": "plain",
        "type": "string"
      },
      "preferred_maintenance_window": {
        "computed": true,
        "description": "The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).\n Format: ` + "`" + `` + "`" + `ddd:hh24:mi-ddd:hh24:mi` + "`" + `` + "`" + ` \n The default is a 30-minute window selected at random from an 8-hour block of time for each AWS Region, occurring on a random day of the week. To see the time blocks available, see [Maintaining an Amazon Aurora DB cluster](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_UpgradeDBInstance.Maintenance.html#AdjustingTheMaintenanceWindow.Aurora) in the *Amazon Aurora User Guide.* \n Valid Days: Mon, Tue, Wed, Thu, Fri, Sat, Sun.\n Constraints: Minimum 30-minute window.\n Valid for: Aurora DB clusters and Multi-AZ DB clusters",
        "description_kind": "plain",
        "type": "string"
      },
      "publicly_accessible": {
        "computed": true,
        "description": "Specifies whether the DB cluster is publicly accessible.\n When the DB cluster is publicly accessible and you connect from outside of the DB cluster's virtual private cloud (VPC), its Domain Name System (DNS) endpoint resolves to the public IP address. When you connect from within the same VPC as the DB cluster, the endpoint resolves to the private IP address. Access to the DB cluster is ultimately controlled by the security group it uses. That public access isn't permitted if the security group assigned to the DB cluster doesn't permit it.\n When the DB cluster isn't publicly accessible, it is an internal DB cluster with a DNS name that resolves to a private IP address.\n Valid for Cluster Type: Multi-AZ DB clusters only\n Default: The default behavior varies depending on whether ` + "`" + `` + "`" + `DBSubnetGroupName` + "`" + `` + "`" + ` is specified.\n If ` + "`" + `` + "`" + `DBSubnetGroupName` + "`" + `` + "`" + ` isn't specified, and ` + "`" + `` + "`" + `PubliclyAccessible` + "`" + `` + "`" + ` isn't specified, the following applies:\n  +  If the default VPC in the target Region doesn?t have an internet gateway attached to it, the DB cluster is private.\n  +  If the default VPC in the target Region has an internet gateway attached to it, the DB cluster is public.\n  \n If ` + "`" + `` + "`" + `DBSubnetGroupName` + "`" + `` + "`" + ` is specified, and ` + "`" + `` + "`" + `PubliclyAccessible` + "`" + `` + "`" + ` isn't specified, the following applies:\n  +  If the subnets are part of a VPC that doesn?t have an internet gateway attached to it, the DB cluster is private.\n  +  If the subnets are part of a VPC that has an internet gateway attached to it, the DB cluster is public.",
        "description_kind": "plain",
        "type": "bool"
      },
      "read_endpoint": {
        "computed": true,
        "description": "The ` + "`" + `` + "`" + `ReadEndpoint` + "`" + `` + "`" + ` return value specifies the reader endpoint for the DB cluster.\n The reader endpoint for a DB cluster load-balances connections across the Aurora Replicas that are available in a DB cluster. As clients request new connections to the reader endpoint, Aurora distributes the connection requests among the Aurora Replicas in the DB cluster. This functionality can help balance your read workload across multiple Aurora Replicas in your DB cluster.\n If a failover occurs, and the Aurora Replica that you are connected to is promoted to be the primary instance, your connection is dropped. To continue sending your read workload to other Aurora Replicas in the cluster, you can then reconnect to the reader endpoint.\n For more information about Aurora endpoints, see [Amazon Aurora connection management](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Overview.Endpoints.html) in the *Amazon Aurora User Guide*.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "address": {
              "computed": true,
              "description": "The host address of the reader endpoint.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "replication_source_identifier": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the source DB instance or DB cluster if this DB cluster is created as a read replica.\n Valid for: Aurora DB clusters only",
        "description_kind": "plain",
        "type": "string"
      },
      "restore_to_time": {
        "computed": true,
        "description": "The date and time to restore the DB cluster to.\n Valid Values: Value must be a time in Universal Coordinated Time (UTC) format\n Constraints:\n  +  Must be before the latest restorable time for the DB instance\n  +  Must be specified if ` + "`" + `` + "`" + `UseLatestRestorableTime` + "`" + `` + "`" + ` parameter isn't provided\n  +  Can't be specified if the ` + "`" + `` + "`" + `UseLatestRestorableTime` + "`" + `` + "`" + ` parameter is enabled\n  +  Can't be specified if the ` + "`" + `` + "`" + `RestoreType` + "`" + `` + "`" + ` parameter is ` + "`" + `` + "`" + `copy-on-write` + "`" + `` + "`" + ` \n  \n This property must be used with ` + "`" + `` + "`" + `SourceDBClusterIdentifier` + "`" + `` + "`" + ` property. The resulting cluster will have the identifier that matches the value of the ` + "`" + `` + "`" + `DBclusterIdentifier` + "`" + `` + "`" + ` property.\n Example: ` + "`" + `` + "`" + `2015-03-07T23:45:00Z` + "`" + `` + "`" + ` \n Valid for: Aurora DB clusters and Multi-AZ DB clusters",
        "description_kind": "plain",
        "type": "string"
      },
      "restore_type": {
        "computed": true,
        "description": "The type of restore to be performed. You can specify one of the following values:\n  +   ` + "`" + `` + "`" + `full-copy` + "`" + `` + "`" + ` - The new DB cluster is restored as a full copy of the source DB cluster.\n  +   ` + "`" + `` + "`" + `copy-on-write` + "`" + `` + "`" + ` - The new DB cluster is restored as a clone of the source DB cluster.\n  \n If you don't specify a ` + "`" + `` + "`" + `RestoreType` + "`" + `` + "`" + ` value, then the new DB cluster is restored as a full copy of the source DB cluster.\n Valid for: Aurora DB clusters and Multi-AZ DB clusters",
        "description_kind": "plain",
        "type": "string"
      },
      "scaling_configuration": {
        "computed": true,
        "description": "The scaling configuration of an Aurora Serverless v1 DB cluster.\n This property is only supported for Aurora Serverless v1. For Aurora Serverless v2, Use the ` + "`" + `` + "`" + `ServerlessV2ScalingConfiguration` + "`" + `` + "`" + ` property.\n Valid for: Aurora Serverless v1 DB clusters only",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "auto_pause": {
              "computed": true,
              "description": "Indicates whether to allow or disallow automatic pause for an Aurora DB cluster in ` + "`" + `` + "`" + `serverless` + "`" + `` + "`" + ` DB engine mode. A DB cluster can be paused only when it's idle (it has no connections).\n  If a DB cluster is paused for more than seven days, the DB cluster might be backed up with a snapshot. In this case, the DB cluster is restored when there is a request to connect to it.",
              "description_kind": "plain",
              "type": "bool"
            },
            "max_capacity": {
              "computed": true,
              "description": "The maximum capacity for an Aurora DB cluster in ` + "`" + `` + "`" + `serverless` + "`" + `` + "`" + ` DB engine mode.\n For Aurora MySQL, valid capacity values are ` + "`" + `` + "`" + `1` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `2` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `4` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `8` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `16` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `32` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `64` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `128` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `256` + "`" + `` + "`" + `.\n For Aurora PostgreSQL, valid capacity values are ` + "`" + `` + "`" + `2` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `4` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `8` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `16` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `32` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `64` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `192` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `384` + "`" + `` + "`" + `.\n The maximum capacity must be greater than or equal to the minimum capacity.",
              "description_kind": "plain",
              "type": "number"
            },
            "min_capacity": {
              "computed": true,
              "description": "The minimum capacity for an Aurora DB cluster in ` + "`" + `` + "`" + `serverless` + "`" + `` + "`" + ` DB engine mode.\n For Aurora MySQL, valid capacity values are ` + "`" + `` + "`" + `1` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `2` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `4` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `8` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `16` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `32` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `64` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `128` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `256` + "`" + `` + "`" + `.\n For Aurora PostgreSQL, valid capacity values are ` + "`" + `` + "`" + `2` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `4` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `8` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `16` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `32` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `64` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `192` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `384` + "`" + `` + "`" + `.\n The minimum capacity must be less than or equal to the maximum capacity.",
              "description_kind": "plain",
              "type": "number"
            },
            "seconds_before_timeout": {
              "computed": true,
              "description": "The amount of time, in seconds, that Aurora Serverless v1 tries to find a scaling point to perform seamless scaling before enforcing the timeout action. The default is 300.\n Specify a value between 60 and 600 seconds.",
              "description_kind": "plain",
              "type": "number"
            },
            "seconds_until_auto_pause": {
              "computed": true,
              "description": "The time, in seconds, before an Aurora DB cluster in ` + "`" + `` + "`" + `serverless` + "`" + `` + "`" + ` mode is paused.\n Specify a value between 300 and 86,400 seconds.",
              "description_kind": "plain",
              "type": "number"
            },
            "timeout_action": {
              "computed": true,
              "description": "The action to take when the timeout is reached, either ` + "`" + `` + "`" + `ForceApplyCapacityChange` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `RollbackCapacityChange` + "`" + `` + "`" + `.\n  ` + "`" + `` + "`" + `ForceApplyCapacityChange` + "`" + `` + "`" + ` sets the capacity to the specified value as soon as possible.\n  ` + "`" + `` + "`" + `RollbackCapacityChange` + "`" + `` + "`" + `, the default, ignores the capacity change if a scaling point isn't found in the timeout period.\n  If you specify ` + "`" + `` + "`" + `ForceApplyCapacityChange` + "`" + `` + "`" + `, connections that prevent Aurora Serverless v1 from finding a scaling point might be dropped.\n  For more information, see [Autoscaling for Aurora Serverless v1](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.how-it-works.html#aurora-serverless.how-it-works.auto-scaling) in the *Amazon Aurora User Guide*.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "serverless_v2_scaling_configuration": {
        "computed": true,
        "description": "The scaling configuration of an Aurora Serverless V2 DB cluster. \n This property is only supported for Aurora Serverless v2. For Aurora Serverless v1, Use the ` + "`" + `` + "`" + `ScalingConfiguration` + "`" + `` + "`" + ` property.\n Valid for: Aurora Serverless v2 DB clusters only",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "max_capacity": {
              "computed": true,
              "description": "The maximum number of Aurora capacity units (ACUs) for a DB instance in an Aurora Serverless v2 cluster. You can specify ACU values in half-step increments, such as 40, 40.5, 41, and so on. The largest value that you can use is 128.\n The maximum capacity must be higher than 0.5 ACUs. For more information, see [Choosing the maximum Aurora Serverless v2 capacity setting for a cluster](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless-v2.setting-capacity.html#aurora-serverless-v2.max_capacity_considerations) in the *Amazon Aurora User Guide*.\n Aurora automatically sets certain parameters for Aurora Serverless V2 DB instances to values that depend on the maximum ACU value in the capacity range. When you update the maximum capacity value, the ` + "`" + `` + "`" + `ParameterApplyStatus` + "`" + `` + "`" + ` value for the DB instance changes to ` + "`" + `` + "`" + `pending-reboot` + "`" + `` + "`" + `. You can update the parameter values by rebooting the DB instance after changing the capacity range.",
              "description_kind": "plain",
              "type": "number"
            },
            "min_capacity": {
              "computed": true,
              "description": "The minimum number of Aurora capacity units (ACUs) for a DB instance in an Aurora Serverless v2 cluster. You can specify ACU values in half-step increments, such as 8, 8.5, 9, and so on. For Aurora versions that support the Aurora Serverless v2 auto-pause feature, the smallest value that you can use is 0. For versions that don't support Aurora Serverless v2 auto-pause, the smallest value that you can use is 0.5.",
              "description_kind": "plain",
              "type": "number"
            },
            "seconds_until_auto_pause": {
              "computed": true,
              "description": "Specifies the number of seconds an Aurora Serverless v2 DB instance must be idle before Aurora attempts to automatically pause it. \n Specify a value between 300 seconds (five minutes) and 86,400 seconds (one day). The default is 300 seconds.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "snapshot_identifier": {
        "computed": true,
        "description": "The identifier for the DB snapshot or DB cluster snapshot to restore from.\n You can use either the name or the Amazon Resource Name (ARN) to specify a DB cluster snapshot. However, you can use only the ARN to specify a DB snapshot.\n After you restore a DB cluster with a ` + "`" + `` + "`" + `SnapshotIdentifier` + "`" + `` + "`" + ` property, you must specify the same ` + "`" + `` + "`" + `SnapshotIdentifier` + "`" + `` + "`" + ` property for any future updates to the DB cluster. When you specify this property for an update, the DB cluster is not restored from the snapshot again, and the data in the database is not changed. However, if you don't specify the ` + "`" + `` + "`" + `SnapshotIdentifier` + "`" + `` + "`" + ` property, an empty DB cluster is created, and the original DB cluster is deleted. If you specify a property that is different from the previous snapshot restore property, a new DB cluster is restored from the specified ` + "`" + `` + "`" + `SnapshotIdentifier` + "`" + `` + "`" + ` property, and the original DB cluster is deleted.\n If you specify the ` + "`" + `` + "`" + `SnapshotIdentifier` + "`" + `` + "`" + ` property to restore a DB cluster (as opposed to specifying it for DB cluster updates), then don't specify the following properties:\n  +   ` + "`" + `` + "`" + `GlobalClusterIdentifier` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `MasterUsername` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `MasterUserPassword` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `ReplicationSourceIdentifier` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `RestoreType` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `SourceDBClusterIdentifier` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `SourceRegion` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `StorageEncrypted` + "`" + `` + "`" + ` (for an encrypted snapshot)\n  +   ` + "`" + `` + "`" + `UseLatestRestorableTime` + "`" + `` + "`" + ` \n  \n Constraints:\n  +  Must match the identifier of an existing Snapshot.\n  \n Valid for: Aurora DB clusters and Multi-AZ DB clusters",
        "description_kind": "plain",
        "type": "string"
      },
      "source_db_cluster_identifier": {
        "computed": true,
        "description": "When restoring a DB cluster to a point in time, the identifier of the source DB cluster from which to restore.\n Constraints:\n  +  Must match the identifier of an existing DBCluster.\n  \n Valid for: Aurora DB clusters and Multi-AZ DB clusters",
        "description_kind": "plain",
        "type": "string"
      },
      "source_region": {
        "computed": true,
        "description": "The AWS Region which contains the source DB cluster when replicating a DB cluster. For example, ` + "`" + `` + "`" + `us-east-1` + "`" + `` + "`" + `. \n Valid for: Aurora DB clusters only",
        "description_kind": "plain",
        "type": "string"
      },
      "storage_encrypted": {
        "computed": true,
        "description": "Indicates whether the DB cluster is encrypted.\n If you specify the ` + "`" + `` + "`" + `KmsKeyId` + "`" + `` + "`" + ` property, then you must enable encryption.\n If you specify the ` + "`" + `` + "`" + `SourceDBClusterIdentifier` + "`" + `` + "`" + ` property, don't specify this property. The value is inherited from the source DB cluster, and if the DB cluster is encrypted, the specified ` + "`" + `` + "`" + `KmsKeyId` + "`" + `` + "`" + ` property is used.\n If you specify the ` + "`" + `` + "`" + `SnapshotIdentifier` + "`" + `` + "`" + ` and the specified snapshot is encrypted, don't specify this property. The value is inherited from the snapshot, and the specified ` + "`" + `` + "`" + `KmsKeyId` + "`" + `` + "`" + ` property is used.\n If you specify the ` + "`" + `` + "`" + `SnapshotIdentifier` + "`" + `` + "`" + ` and the specified snapshot isn't encrypted, you can use this property to specify that the restored DB cluster is encrypted. Specify the ` + "`" + `` + "`" + `KmsKeyId` + "`" + `` + "`" + ` property for the KMS key to use for encryption. If you don't want the restored DB cluster to be encrypted, then don't set this property or set it to ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `.\n  If you specify both the ` + "`" + `` + "`" + `StorageEncrypted` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `SnapshotIdentifier` + "`" + `` + "`" + ` properties without specifying the ` + "`" + `` + "`" + `KmsKeyId` + "`" + `` + "`" + ` property, then the restored DB cluster inherits the encryption settings from the DB snapshot that provide.\n  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
        "description_kind": "plain",
        "type": "bool"
      },
      "storage_throughput": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "storage_type": {
        "computed": true,
        "description": "The storage type to associate with the DB cluster.\n For information on storage types for Aurora DB clusters, see [Storage configurations for Amazon Aurora DB clusters](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Overview.StorageReliability.html#aurora-storage-type). For information on storage types for Multi-AZ DB clusters, see [Settings for creating Multi-AZ DB clusters](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/create-multi-az-db-cluster.html#create-multi-az-db-cluster-settings).\n This setting is required to create a Multi-AZ DB cluster.\n When specified for a Multi-AZ DB cluster, a value for the ` + "`" + `` + "`" + `Iops` + "`" + `` + "`" + ` parameter is required.\n Valid for Cluster Type: Aurora DB clusters and Multi-AZ DB clusters\n Valid Values:\n  +  Aurora DB clusters - ` + "`" + `` + "`" + `aurora | aurora-iopt1` + "`" + `` + "`" + ` \n  +  Multi-AZ DB clusters - ` + "`" + `` + "`" + `io1 | io2 | gp3` + "`" + `` + "`" + ` \n  \n Default:\n  +  Aurora DB clusters - ` + "`" + `` + "`" + `aurora` + "`" + `` + "`" + ` \n  +  Multi-AZ DB clusters - ` + "`" + `` + "`" + `io1` + "`" + `` + "`" + ` \n  \n  When you create an Aurora DB cluster with the storage type set to ` + "`" + `` + "`" + `aurora-iopt1` + "`" + `` + "`" + `, the storage type is returned in the response. The storage type isn't returned when you set it to ` + "`" + `` + "`" + `aurora` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags to assign to the DB cluster.\n Valid for Cluster Type: Aurora DB clusters and Multi-AZ DB clusters",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "A key is the required name of the tag. The string value can be from 1 to 128 Unicode characters in length and can't be prefixed with ` + "`" + `` + "`" + `aws:` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `rds:` + "`" + `` + "`" + `. The string can only contain only the set of Unicode letters, digits, white-space, '_', '.', ':', '/', '=', '+', '-', '@' (Java regex: \"^([\\\\p{L}\\\\p{Z}\\\\p{N}_.:/=+\\\\-@]*)$\").",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "A value is the optional value of the tag. The string value can be from 1 to 256 Unicode characters in length and can't be prefixed with ` + "`" + `` + "`" + `aws:` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `rds:` + "`" + `` + "`" + `. The string can only contain only the set of Unicode letters, digits, white-space, '_', '.', ':', '/', '=', '+', '-', '@' (Java regex: \"^([\\\\p{L}\\\\p{Z}\\\\p{N}_.:/=+\\\\-@]*)$\").",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "use_latest_restorable_time": {
        "computed": true,
        "description": "A value that indicates whether to restore the DB cluster to the latest restorable backup time. By default, the DB cluster is not restored to the latest restorable backup time. \n Valid for: Aurora DB clusters and Multi-AZ DB clusters",
        "description_kind": "plain",
        "type": "bool"
      },
      "vpc_security_group_ids": {
        "computed": true,
        "description": "A list of EC2 VPC security groups to associate with this DB cluster.\n If you plan to update the resource, don't specify VPC security groups in a shared VPC.\n Valid for: Aurora DB clusters and Multi-AZ DB clusters",
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
