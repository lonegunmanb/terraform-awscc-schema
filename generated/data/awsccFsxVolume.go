package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccFsxVolume = `{
  "block": {
    "attributes": {
      "backup_id": {
        "computed": true,
        "description": "Specifies the ID of the volume backup to use to create a new volume.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the volume.",
        "description_kind": "plain",
        "type": "string"
      },
      "ontap_configuration": {
        "computed": true,
        "description": "The configuration of an Amazon FSx for NetApp ONTAP volume.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "aggregate_configuration": {
              "computed": true,
              "description": "Used to specify the configuration options for an FSx for ONTAP volume's storage aggregate or aggregates.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "aggregates": {
                    "computed": true,
                    "description": "The list of aggregates that this volume resides on. Aggregates are storage pools which make up your primary storage tier.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "constituents_per_aggregate": {
                    "computed": true,
                    "description": "Used to explicitly set the number of constituents within the FlexGroup per storage aggregate. This field is optional when creating a FlexGroup volume. If unspecified, the default value will be 8. This field cannot be provided when creating a FlexVol volume.",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "copy_tags_to_backups": {
              "computed": true,
              "description": "A boolean flag indicating whether tags for the volume should be copied to backups.",
              "description_kind": "plain",
              "type": "string"
            },
            "junction_path": {
              "computed": true,
              "description": "Specifies the location in the SVM's namespace where the volume is mounted. This parameter is required. The JunctionPath must have a leading forward slash, such as /vol3.",
              "description_kind": "plain",
              "type": "string"
            },
            "ontap_volume_type": {
              "computed": true,
              "description": "Specifies the type of volume you are creating. Valid values are the following: RW or DP",
              "description_kind": "plain",
              "type": "string"
            },
            "security_style": {
              "computed": true,
              "description": "Specifies the security style for the volume. If a volume's security style is not specified, it is automatically set to the root volume's security style.",
              "description_kind": "plain",
              "type": "string"
            },
            "size_in_bytes": {
              "computed": true,
              "description": "Specifies the configured size of the volume, in bytes.",
              "description_kind": "plain",
              "type": "string"
            },
            "size_in_megabytes": {
              "computed": true,
              "description": "Use SizeInBytes instead. Specifies the size of the volume, in megabytes (MB), that you are creating",
              "description_kind": "plain",
              "type": "string"
            },
            "snaplock_configuration": {
              "computed": true,
              "description": "The SnapLock configuration object for an FSx for ONTAP SnapLock volume.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "audit_log_volume": {
                    "computed": true,
                    "description": "Enables or disables the audit log volume for an FSx for ONTAP SnapLock volume",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "autocommit_period": {
                    "computed": true,
                    "description": "The configuration object for setting the autocommit period of files in an FSx for ONTAP SnapLock volume.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "type": {
                          "computed": true,
                          "description": "Defines the type of time for the autocommit period of a file in an FSx for ONTAP SnapLock volume. Setting this value to NONE disables autocommit. The default value is NONE.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "value": {
                          "computed": true,
                          "description": "Defines the amount of time for the autocommit period of a file in an FSx for ONTAP SnapLock volume.",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "privileged_delete": {
                    "computed": true,
                    "description": "Enables, disables, or permanently disables privileged delete on an FSx for ONTAP SnapLock Enterprise volume.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "retention_period": {
                    "computed": true,
                    "description": "Specifies the retention period of an FSx for ONTAP SnapLock volume.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "default_retention": {
                          "computed": true,
                          "description": "The retention period assigned to a write once, read many (WORM) file by default if an explicit retention period is not set for an FSx for ONTAP SnapLock volume.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "type": {
                                "computed": true,
                                "description": "Defines the type of time for the retention period of an FSx for ONTAP SnapLock volume. Set it to one of the valid types. If you set it to INFINITE, the files are retained forever. If you set it to UNSPECIFIED, the files are retained until you set an explicit retention period.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "value": {
                                "computed": true,
                                "description": "Defines the amount of time for the retention period of an FSx for ONTAP SnapLock volume. You can't set a value for INFINITE or UNSPECIFIED.",
                                "description_kind": "plain",
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "maximum_retention": {
                          "computed": true,
                          "description": "The longest retention period that can be assigned to a WORM file on an FSx for ONTAP SnapLock volume.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "type": {
                                "computed": true,
                                "description": "Defines the type of time for the retention period of an FSx for ONTAP SnapLock volume. Set it to one of the valid types. If you set it to INFINITE, the files are retained forever. If you set it to UNSPECIFIED, the files are retained until you set an explicit retention period.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "value": {
                                "computed": true,
                                "description": "Defines the amount of time for the retention period of an FSx for ONTAP SnapLock volume. You can't set a value for INFINITE or UNSPECIFIED.",
                                "description_kind": "plain",
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "minimum_retention": {
                          "computed": true,
                          "description": "The shortest retention period that can be assigned to a WORM file on an FSx for ONTAP SnapLock volume.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "type": {
                                "computed": true,
                                "description": "Defines the type of time for the retention period of an FSx for ONTAP SnapLock volume. Set it to one of the valid types. If you set it to INFINITE, the files are retained forever. If you set it to UNSPECIFIED, the files are retained until you set an explicit retention period.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "value": {
                                "computed": true,
                                "description": "Defines the amount of time for the retention period of an FSx for ONTAP SnapLock volume. You can't set a value for INFINITE or UNSPECIFIED.",
                                "description_kind": "plain",
                                "type": "number"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "snaplock_type": {
                    "computed": true,
                    "description": "Specifies the retention mode of an FSx for ONTAP SnapLock volume. After it is set, it can't be changed.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "volume_append_mode_enabled": {
                    "computed": true,
                    "description": "Enables or disables volume-append mode on an FSx for ONTAP SnapLock volume.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "snapshot_policy": {
              "computed": true,
              "description": "Specifies the snapshot policy for the volume. There are three built-in snapshot policies: default, default-1weekly, none.",
              "description_kind": "plain",
              "type": "string"
            },
            "storage_efficiency_enabled": {
              "computed": true,
              "description": "Set to true to enable deduplication, compression, and compaction storage efficiency features on the volume, or set to false to disable them.",
              "description_kind": "plain",
              "type": "string"
            },
            "storage_virtual_machine_id": {
              "computed": true,
              "description": "Specifies the ONTAP SVM in which to create the volume.",
              "description_kind": "plain",
              "type": "string"
            },
            "tiering_policy": {
              "computed": true,
              "description": "Describes the data tiering policy for an ONTAP volume.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cooling_period": {
                    "computed": true,
                    "description": "Specifies the number of days that user data in a volume must remain inactive before it is considered \"cold\" and moved to the capacity pool.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "name": {
                    "computed": true,
                    "description": "Specifies the tiering policy used to transition data. Default value is SNAPSHOT_ONLY.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "volume_style": {
              "computed": true,
              "description": "Use to specify the style of an ONTAP volume.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "open_zfs_configuration": {
        "computed": true,
        "description": "The configuration of an Amazon FSx for OpenZFS volume.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "copy_tags_to_snapshots": {
              "computed": true,
              "description": "A Boolean value indicating whether tags for the volume should be copied to snapshots. This value defaults to false. If this value is set to true, and you do not specify any tags, all tags for the original volume are copied over to snapshots. If this value is set to true, and you do specify one or more tags, only the specified tags for the original volume are copied over to snapshots. If you specify one or more tags when creating a new snapshot, no tags are copied over from the original volume, regardless of this value.",
              "description_kind": "plain",
              "type": "bool"
            },
            "data_compression_type": {
              "computed": true,
              "description": "Specifies the method used to compress the data on the volume",
              "description_kind": "plain",
              "type": "string"
            },
            "nfs_exports": {
              "computed": true,
              "description": "The configuration object for mounting a Network File System (NFS) file system.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "client_configurations": {
                    "computed": true,
                    "description": "The configuration object for mounting a Network File System (NFS) file system.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "clients": {
                          "computed": true,
                          "description": "A value that specifies who can mount the file system. You can provide a wildcard character (*), an IP address (0.0.0.0), or a CIDR address (192.0.2.0/24). By default, Amazon FSx uses the wildcard character when specifying the client.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "options": {
                          "computed": true,
                          "description": "The configuration object for mounting a Network File System (NFS) file system.",
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "nesting_mode": "list"
                    }
                  }
                },
                "nesting_mode": "list"
              }
            },
            "options": {
              "computed": true,
              "description": "The configuration object for mounting a Network File System (NFS) file system.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "origin_snapshot": {
              "computed": true,
              "description": "The configuration of an Amazon FSx for OpenZFS volume.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "copy_strategy": {
                    "computed": true,
                    "description": "The configuration object for mounting a Network File System (NFS) file system.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "snapshot_arn": {
                    "computed": true,
                    "description": "Specifies the snapshot to use when creating an OpenZFS volume from a snapshot.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "parent_volume_id": {
              "computed": true,
              "description": "The ID of the volume to use as the parent volume of the volume that you are creating.",
              "description_kind": "plain",
              "type": "string"
            },
            "read_only": {
              "computed": true,
              "description": "A Boolean value indicating whether the volume is read-only.",
              "description_kind": "plain",
              "type": "bool"
            },
            "record_size_ki_b": {
              "computed": true,
              "description": "Specifies the suggested block size for a volume in a ZFS dataset, in kibibytes (KiB).",
              "description_kind": "plain",
              "type": "number"
            },
            "storage_capacity_quota_gi_b": {
              "computed": true,
              "description": "Sets the maximum storage size in gibibytes (GiB) for the volume. You can specify a quota that is larger than the storage on the parent volume. A volume quota limits the amount of storage that the volume can consume to the configured amount, but does not guarantee the space will be available on the parent volume. To guarantee quota space, you must also set StorageCapacityReservationGiB. To not specify a storage capacity quota, set this to -1.",
              "description_kind": "plain",
              "type": "number"
            },
            "storage_capacity_reservation_gi_b": {
              "computed": true,
              "description": "Specifies the amount of storage in gibibytes (GiB) to reserve from the parent volume. Setting StorageCapacityReservationGiB guarantees that the specified amount of storage space on the parent volume will always be available for the volume. You can't reserve more storage than the parent volume has. To not specify a storage capacity reservation, set this to 0 or -1. For more information, see Volume properties in the Amazon FSx for OpenZFS User Guide.",
              "description_kind": "plain",
              "type": "number"
            },
            "user_and_group_quotas": {
              "computed": true,
              "description": "Configures how much storage users and groups can use on the volume.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "id": {
                    "computed": true,
                    "description": "The ID of the user or group that the quota applies to.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "storage_capacity_quota_gi_b": {
                    "computed": true,
                    "description": "The user or group's storage quota, in gibibytes (GiB).",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "type": {
                    "computed": true,
                    "description": "Specifies whether the quota applies to a user or group.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "resource_arn": {
        "computed": true,
        "description": "Returns the volume's Amazon Resource Name (ARN).",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "One or more tags.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "A value that specifies the TagKey, the name of the tag. Tag keys must be unique for the resource to which they are attached.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "A value that specifies the TagValue, the value assigned to the corresponding tag key. Tag values can be null and don't have to be unique in a tag set. For example, you can have a key-value pair in a tag set of finances : April and also of payroll : April.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "uuid": {
        "computed": true,
        "description": "Returns the volume's ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "volume_id": {
        "computed": true,
        "description": "Returns the volume's universally unique identifier (UUID).",
        "description_kind": "plain",
        "type": "string"
      },
      "volume_type": {
        "computed": true,
        "description": "The type of the volume.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::FSx::Volume",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccFsxVolumeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccFsxVolume), &result)
	return &result
}
