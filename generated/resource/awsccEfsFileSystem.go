package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEfsFileSystem = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "availability_zone_name": {
        "computed": true,
        "description": "For One Zone file systems, specify the AWS Availability Zone in which to create the file system. Use the format ` + "`" + `` + "`" + `us-east-1a` + "`" + `` + "`" + ` to specify the Availability Zone. For more information about One Zone file systems, see [EFS file system types](https://docs.aws.amazon.com/efs/latest/ug/availability-durability.html#file-system-type) in the *Amazon EFS User Guide*.\n  One Zone file systems are not available in all Availability Zones in AWS-Regions where Amazon EFS is available.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "backup_policy": {
        "computed": true,
        "description": "Use the ` + "`" + `` + "`" + `BackupPolicy` + "`" + `` + "`" + ` to turn automatic backups on or off for the file system.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "status": {
              "computed": true,
              "description": "Set the backup policy status for the file system.\n  +  *ENABLED* - Turns automatic backups on for the file system. \n  +  *DISABLED* - Turns automatic backups off for the file system.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "bypass_policy_lockout_safety_check": {
        "computed": true,
        "description": "(Optional) A boolean that specifies whether or not to bypass the ` + "`" + `` + "`" + `FileSystemPolicy` + "`" + `` + "`" + ` lockout safety check. The lockout safety check determines whether the policy in the request will lock out, or prevent, the IAM principal that is making the request from making future ` + "`" + `` + "`" + `PutFileSystemPolicy` + "`" + `` + "`" + ` requests on this file system. Set ` + "`" + `` + "`" + `BypassPolicyLockoutSafetyCheck` + "`" + `` + "`" + ` to ` + "`" + `` + "`" + `True` + "`" + `` + "`" + ` only when you intend to prevent the IAM principal that is making the request from making subsequent ` + "`" + `` + "`" + `PutFileSystemPolicy` + "`" + `` + "`" + ` requests on this file system. The default value is ` + "`" + `` + "`" + `False` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "encrypted": {
        "computed": true,
        "description": "A Boolean value that, if true, creates an encrypted file system. When creating an encrypted file system, you have the option of specifying a KmsKeyId for an existing kms-key-long. If you don't specify a kms-key, then the default kms-key for EFS, ` + "`" + `` + "`" + `/aws/elasticfilesystem` + "`" + `` + "`" + `, is used to protect the encrypted file system.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "file_system_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "file_system_policy": {
        "computed": true,
        "description": "The ` + "`" + `` + "`" + `FileSystemPolicy` + "`" + `` + "`" + ` for the EFS file system. A file system policy is an IAM resource policy used to control NFS access to an EFS file system. For more information, see [Using to control NFS access to Amazon EFS](https://docs.aws.amazon.com/efs/latest/ug/iam-access-control-nfs-efs.html) in the *Amazon EFS User Guide*.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "file_system_protection": {
        "computed": true,
        "description": "Describes the protection on the file system.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "replication_overwrite_protection": {
              "computed": true,
              "description": "The status of the file system's replication overwrite protection.\n  +  ` + "`" + `` + "`" + `ENABLED` + "`" + `` + "`" + ` ? The file system cannot be used as the destination file system in a replication configuration. The file system is writeable. Replication overwrite protection is ` + "`" + `` + "`" + `ENABLED` + "`" + `` + "`" + ` by default. \n  +  ` + "`" + `` + "`" + `DISABLED` + "`" + `` + "`" + ` ? The file system can be used as the destination file system in a replication configuration. The file system is read-only and can only be modified by EFS replication.\n  +  ` + "`" + `` + "`" + `REPLICATING` + "`" + `` + "`" + ` ? The file system is being used as the destination file system in a replication configuration. The file system is read-only and is modified only by EFS replication.\n  \n If the replication configuration is deleted, the file system's replication overwrite protection is re-enabled, the file system becomes writeable.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "file_system_tags": {
        "computed": true,
        "description": "Use to create one or more tags associated with the file system. Each tag is a user-defined key-value pair. Name your file system on creation by including a ` + "`" + `` + "`" + `\"Key\":\"Name\",\"Value\":\"{value}\"` + "`" + `` + "`" + ` key-value pair. Each key must be unique. For more information, see [Tagging resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the *General Reference Guide*.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key (String). The key can't start with ` + "`" + `` + "`" + `aws:` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the tag key.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description": "The ID of the kms-key-long to be used to protect the encrypted file system. This parameter is only required if you want to use a nondefault kms-key. If this parameter is not specified, the default kms-key for EFS is used. This ID can be in one of the following formats:\n  +  Key ID - A unique identifier of the key, for example ` + "`" + `` + "`" + `1234abcd-12ab-34cd-56ef-1234567890ab` + "`" + `` + "`" + `.\n  +  ARN - An Amazon Resource Name (ARN) for the key, for example ` + "`" + `` + "`" + `arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab` + "`" + `` + "`" + `.\n  +  Key alias - A previously created display name for a key, for example ` + "`" + `` + "`" + `alias/projectKey1` + "`" + `` + "`" + `.\n  +  Key alias ARN - An ARN for a key alias, for example ` + "`" + `` + "`" + `arn:aws:kms:us-west-2:444455556666:alias/projectKey1` + "`" + `` + "`" + `.\n  \n If ` + "`" + `` + "`" + `KmsKeyId` + "`" + `` + "`" + ` is specified, the ` + "`" + `` + "`" + `Encrypted` + "`" + `` + "`" + ` parameter must be set to true.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "lifecycle_policies": {
        "computed": true,
        "description": "An array of ` + "`" + `` + "`" + `LifecyclePolicy` + "`" + `` + "`" + ` objects that define the file system's ` + "`" + `` + "`" + `LifecycleConfiguration` + "`" + `` + "`" + ` object. A ` + "`" + `` + "`" + `LifecycleConfiguration` + "`" + `` + "`" + ` object informs Lifecycle management of the following:\n  +  When to move files in the file system from primary storage to IA storage.\n  +  When to move files in the file system from primary storage or IA storage to Archive storage.\n  +  When to move files that are in IA or Archive storage to primary storage.\n  \n  EFS requires that each ` + "`" + `` + "`" + `LifecyclePolicy` + "`" + `` + "`" + ` object have only a single transition. This means that in a request body, ` + "`" + `` + "`" + `LifecyclePolicies` + "`" + `` + "`" + ` needs to be structured as an array of ` + "`" + `` + "`" + `LifecyclePolicy` + "`" + `` + "`" + ` objects, one object for each transition, ` + "`" + `` + "`" + `TransitionToIA` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `TransitionToArchive` + "`" + `` + "`" + `` + "`" + `` + "`" + `TransitionToPrimaryStorageClass` + "`" + `` + "`" + `. See the example requests in the following section for more information.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "transition_to_archive": {
              "computed": true,
              "description": "The number of days after files were last accessed in primary storage (the Standard storage class) at which to move them to Archive storage. Metadata operations such as listing the contents of a directory don't count as file access events.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "transition_to_ia": {
              "computed": true,
              "description": "The number of days after files were last accessed in primary storage (the Standard storage class) at which to move them to Infrequent Access (IA) storage. Metadata operations such as listing the contents of a directory don't count as file access events.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "transition_to_primary_storage_class": {
              "computed": true,
              "description": "Whether to move files back to primary (Standard) storage after they are accessed in IA or Archive storage. Metadata operations such as listing the contents of a directory don't count as file access events.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "performance_mode": {
        "computed": true,
        "description": "The performance mode of the file system. We recommend ` + "`" + `` + "`" + `generalPurpose` + "`" + `` + "`" + ` performance mode for all file systems. File systems using the ` + "`" + `` + "`" + `maxIO` + "`" + `` + "`" + ` performance mode can scale to higher levels of aggregate throughput and operations per second with a tradeoff of slightly higher latencies for most file operations. The performance mode can't be changed after the file system has been created. The ` + "`" + `` + "`" + `maxIO` + "`" + `` + "`" + ` mode is not supported on One Zone file systems.\n  Due to the higher per-operation latencies with Max I/O, we recommend using General Purpose performance mode for all file systems.\n  Default is ` + "`" + `` + "`" + `generalPurpose` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "provisioned_throughput_in_mibps": {
        "computed": true,
        "description": "The throughput, measured in mebibytes per second (MiBps), that you want to provision for a file system that you're creating. Required if ` + "`" + `` + "`" + `ThroughputMode` + "`" + `` + "`" + ` is set to ` + "`" + `` + "`" + `provisioned` + "`" + `` + "`" + `. Valid values are 1-3414 MiBps, with the upper limit depending on Region. To increase this limit, contact SUP. For more information, see [Amazon EFS quotas that you can increase](https://docs.aws.amazon.com/efs/latest/ug/limits.html#soft-limits) in the *Amazon EFS User Guide*.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "replication_configuration": {
        "computed": true,
        "description": "Describes the replication configuration for a specific file system.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "destinations": {
              "computed": true,
              "description": "An array of destination objects. Only one destination object is supported.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "availability_zone_name": {
                    "computed": true,
                    "description": "For One Zone file systems, the replication configuration must specify the Availability Zone in which the destination file system is located. \n Use the format ` + "`" + `` + "`" + `us-east-1a` + "`" + `` + "`" + ` to specify the Availability Zone. For more information about One Zone file systems, see [EFS file system types](https://docs.aws.amazon.com/efs/latest/ug/storage-classes.html) in the *Amazon EFS User Guide*.\n  One Zone file system type is not available in all Availability Zones in AWS-Regions where Amazon EFS is available.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "file_system_id": {
                    "computed": true,
                    "description": "The ID of the destination Amazon EFS file system.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "kms_key_id": {
                    "computed": true,
                    "description": "The ID of an kms-key-long used to protect the encrypted file system.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "region": {
                    "computed": true,
                    "description": "The AWS-Region in which the destination file system is located.\n  For One Zone file systems, the replication configuration must specify the AWS-Region in which the destination file system is located.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "role_arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of the current source file system in the replication configuration.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "status": {
                    "computed": true,
                    "description": "Describes the status of the replication configuration. For more information about replication status, see [Viewing replication details](https://docs.aws.amazon.com//efs/latest/ug/awsbackup.html#restoring-backup-efsmonitoring-replication-status.html) in the *Amazon EFS User Guide*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "status_message": {
                    "computed": true,
                    "description": "Message that provides details about the ` + "`" + `` + "`" + `PAUSED` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `ERRROR` + "`" + `` + "`" + ` state of the replication destination configuration. For more information about replication status messages, see [Viewing replication details](https://docs.aws.amazon.com//efs/latest/ug/awsbackup.html#restoring-backup-efsmonitoring-replication-status.html) in the *Amazon EFS User Guide*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "throughput_mode": {
        "computed": true,
        "description": "Specifies the throughput mode for the file system. The mode can be ` + "`" + `` + "`" + `bursting` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `provisioned` + "`" + `` + "`" + `, or ` + "`" + `` + "`" + `elastic` + "`" + `` + "`" + `. If you set ` + "`" + `` + "`" + `ThroughputMode` + "`" + `` + "`" + ` to ` + "`" + `` + "`" + `provisioned` + "`" + `` + "`" + `, you must also set a value for ` + "`" + `` + "`" + `ProvisionedThroughputInMibps` + "`" + `` + "`" + `. After you create the file system, you can decrease your file system's Provisioned throughput or change between the throughput modes, with certain time restrictions. For more information, see [Specifying throughput with provisioned mode](https://docs.aws.amazon.com/efs/latest/ug/performance.html#provisioned-throughput) in the *Amazon EFS User Guide*. \n Default is ` + "`" + `` + "`" + `bursting` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "The ` + "`" + `` + "`" + `AWS::EFS::FileSystem` + "`" + `` + "`" + ` resource creates a new, empty file system in EFSlong (EFS). You must create a mount target ([AWS::EFS::MountTarget](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html)) to mount your EFS file system on an EC2 or other AWS cloud compute resource.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEfsFileSystemSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEfsFileSystem), &result)
	return &result
}
