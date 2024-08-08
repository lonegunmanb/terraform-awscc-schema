package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRedshiftCluster = `{
  "block": {
    "attributes": {
      "allow_version_upgrade": {
        "computed": true,
        "description": "Major version upgrades can be applied during the maintenance window to the Amazon Redshift engine that is running on the cluster. Default value is True",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "aqua_configuration_status": {
        "computed": true,
        "description": "The value represents how the cluster is configured to use AQUA (Advanced Query Accelerator) after the cluster is restored. Possible values include the following.\n\nenabled - Use AQUA if it is available for the current Region and Amazon Redshift node type.\ndisabled - Don't use AQUA.\nauto - Amazon Redshift determines whether to use AQUA.\n",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "automated_snapshot_retention_period": {
        "computed": true,
        "description": "The number of days that automated snapshots are retained. If the value is 0, automated snapshots are disabled. Default value is 1",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "availability_zone": {
        "computed": true,
        "description": "The EC2 Availability Zone (AZ) in which you want Amazon Redshift to provision the cluster. Default: A random, system-chosen Availability Zone in the region that is specified by the endpoint",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "availability_zone_relocation": {
        "computed": true,
        "description": "The option to enable relocation for an Amazon Redshift cluster between Availability Zones after the cluster modification is complete.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "availability_zone_relocation_status": {
        "computed": true,
        "description": "The availability zone relocation status of the cluster",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "classic": {
        "computed": true,
        "description": "A boolean value indicating whether the resize operation is using the classic resize process. If you don't provide this parameter or set the value to false , the resize type is elastic.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "cluster_identifier": {
        "computed": true,
        "description": "A unique identifier for the cluster. You use this identifier to refer to the cluster for any subsequent cluster operations such as deleting or modifying. All alphabetical characters must be lower case, no hypens at the end, no two consecutive hyphens. Cluster name should be unique for all clusters within an AWS account",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cluster_namespace_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the cluster namespace.",
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_parameter_group_name": {
        "computed": true,
        "description": "The name of the parameter group to be associated with this cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cluster_security_groups": {
        "computed": true,
        "description": "A list of security groups to be associated with this cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "cluster_subnet_group_name": {
        "computed": true,
        "description": "The name of a cluster subnet group to be associated with this cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cluster_type": {
        "description": "The type of the cluster. When cluster type is specified as single-node, the NumberOfNodes parameter is not required and if multi-node, the NumberOfNodes parameter is required",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cluster_version": {
        "computed": true,
        "description": "The version of the Amazon Redshift engine software that you want to deploy on the cluster.The version selected runs on all the nodes in the cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "db_name": {
        "description": "The name of the first database to be created when the cluster is created. To create additional databases after the cluster is created, connect to the cluster with a SQL client and use SQL commands to create a database.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "defer_maintenance": {
        "computed": true,
        "description": "A boolean indicating whether to enable the deferred maintenance window.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "defer_maintenance_duration": {
        "computed": true,
        "description": "An integer indicating the duration of the maintenance window in days. If you specify a duration, you can't specify an end time. The duration must be 45 days or less.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "defer_maintenance_end_time": {
        "computed": true,
        "description": "A timestamp indicating end time for the deferred maintenance window. If you specify an end time, you can't specify a duration.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "defer_maintenance_identifier": {
        "computed": true,
        "description": "A unique identifier for the deferred maintenance window.",
        "description_kind": "plain",
        "type": "string"
      },
      "defer_maintenance_start_time": {
        "computed": true,
        "description": "A timestamp indicating the start time for the deferred maintenance window.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_region": {
        "computed": true,
        "description": "The destination AWS Region that you want to copy snapshots to. Constraints: Must be the name of a valid AWS Region. For more information, see Regions and Endpoints in the Amazon Web Services [https://docs.aws.amazon.com/general/latest/gr/rande.html#redshift_region] General Reference",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "elastic_ip": {
        "computed": true,
        "description": "The Elastic IP (EIP) address for the cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "encrypted": {
        "computed": true,
        "description": "If true, the data in the cluster is encrypted at rest.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "endpoint": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "address": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "port": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "enhanced_vpc_routing": {
        "computed": true,
        "description": "An option that specifies whether to create the cluster with enhanced VPC routing enabled. To create a cluster that uses enhanced VPC routing, the cluster must be in a VPC. For more information, see Enhanced VPC Routing in the Amazon Redshift Cluster Management Guide.\n\nIf this option is true , enhanced VPC routing is enabled.\n\nDefault: false",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "hsm_client_certificate_identifier": {
        "computed": true,
        "description": "Specifies the name of the HSM client certificate the Amazon Redshift cluster uses to retrieve the data encryption keys stored in an HSM",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "hsm_configuration_identifier": {
        "computed": true,
        "description": "Specifies the name of the HSM configuration that contains the information the Amazon Redshift cluster can use to retrieve and store keys in an HSM.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "iam_roles": {
        "computed": true,
        "description": "A list of AWS Identity and Access Management (IAM) roles that can be used by the cluster to access other AWS services. You must supply the IAM roles in their Amazon Resource Name (ARN) format. You can supply up to 50 IAM roles in a single request",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description": "The AWS Key Management Service (KMS) key ID of the encryption key that you want to use to encrypt data in the cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "logging_properties": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "bucket_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "log_destination_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "log_exports": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "s3_key_prefix": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "maintenance_track_name": {
        "computed": true,
        "description": "The name for the maintenance track that you want to assign for the cluster. This name change is asynchronous. The new track name stays in the PendingModifiedValues for the cluster until the next maintenance window. When the maintenance track changes, the cluster is switched to the latest cluster release available for the maintenance track. At this point, the maintenance track name is applied.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "manage_master_password": {
        "computed": true,
        "description": "A boolean indicating if the redshift cluster's admin user credentials is managed by Redshift or not. You can't use MasterUserPassword if ManageMasterPassword is true. If ManageMasterPassword is false or not set, Amazon Redshift uses MasterUserPassword for the admin user account's password.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "manual_snapshot_retention_period": {
        "computed": true,
        "description": "The number of days to retain newly copied snapshots in the destination AWS Region after they are copied from the source AWS Region. If the value is -1, the manual snapshot is retained indefinitely.\n\nThe value must be either -1 or an integer between 1 and 3,653.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "master_password_secret_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the cluster's admin user credentials secret.",
        "description_kind": "plain",
        "type": "string"
      },
      "master_password_secret_kms_key_id": {
        "computed": true,
        "description": "The ID of the Key Management Service (KMS) key used to encrypt and store the cluster's admin user credentials secret.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "master_user_password": {
        "computed": true,
        "description": "The password associated with the master user account for the cluster that is being created. You can't use MasterUserPassword if ManageMasterPassword is true. Password must be between 8 and 64 characters in length, should have at least one uppercase letter.Must contain at least one lowercase letter.Must contain one number.Can be any printable ASCII character.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "master_username": {
        "description": "The user name associated with the master user account for the cluster that is being created. The user name can't be PUBLIC and first character must be a letter.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "multi_az": {
        "computed": true,
        "description": "A boolean indicating if the redshift cluster is multi-az or not. If you don't provide this parameter or set the value to false, the redshift cluster will be single-az.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "namespace_resource_policy": {
        "computed": true,
        "description": "The namespace resource policy document that will be attached to a Redshift cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "node_type": {
        "description": "The node type to be provisioned for the cluster.Valid Values: ds2.xlarge | ds2.8xlarge | dc1.large | dc1.8xlarge | dc2.large | dc2.8xlarge | ra3.4xlarge | ra3.16xlarge",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "number_of_nodes": {
        "computed": true,
        "description": "The number of compute nodes in the cluster. This parameter is required when the ClusterType parameter is specified as multi-node.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "owner_account": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "port": {
        "computed": true,
        "description": "The port number on which the cluster accepts incoming connections. The cluster is accessible only via the JDBC and ODBC connection strings",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "preferred_maintenance_window": {
        "computed": true,
        "description": "The weekly time range (in UTC) during which automated cluster maintenance can occur.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "publicly_accessible": {
        "computed": true,
        "description": "If true, the cluster can be accessed from a public network.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "resource_action": {
        "computed": true,
        "description": "The Redshift operation to be performed. Resource Action supports pause-cluster, resume-cluster, failover-primary-compute APIs",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "revision_target": {
        "computed": true,
        "description": "The identifier of the database revision. You can retrieve this value from the response to the DescribeClusterDbRevisions request.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rotate_encryption_key": {
        "computed": true,
        "description": "A boolean indicating if we want to rotate Encryption Keys.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "snapshot_cluster_identifier": {
        "computed": true,
        "description": "The name of the cluster the source snapshot was created from. This parameter is required if your IAM user has a policy containing a snapshot resource element that specifies anything other than * for the cluster name.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "snapshot_copy_grant_name": {
        "computed": true,
        "description": "The name of the snapshot copy grant to use when snapshots of an AWS KMS-encrypted cluster are copied to the destination region.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "snapshot_copy_manual": {
        "computed": true,
        "description": "Indicates whether to apply the snapshot retention period to newly copied manual snapshots instead of automated snapshots.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "snapshot_copy_retention_period": {
        "computed": true,
        "description": "The number of days to retain automated snapshots in the destination region after they are copied from the source region. \n\n Default is 7. \n\n Constraints: Must be at least 1 and no more than 35.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "snapshot_identifier": {
        "computed": true,
        "description": "The name of the snapshot from which to create the new cluster. This parameter isn't case sensitive.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The list of tags for the cluster parameter group.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "vpc_security_group_ids": {
        "computed": true,
        "description": "A list of Virtual Private Cloud (VPC) security groups to be associated with the cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "An example resource schema demonstrating some basic constructs and validation rules.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRedshiftClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRedshiftCluster), &result)
	return &result
}
