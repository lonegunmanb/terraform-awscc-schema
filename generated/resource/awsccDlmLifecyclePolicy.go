package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDlmLifecyclePolicy = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the lifecycle policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "copy_tags": {
        "computed": true,
        "description": "**[Default policies only]** Indicates whether the policy should copy tags from the source resource to the snapshot or AMI. If you do not specify a value, the default is false.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "create_interval": {
        "computed": true,
        "description": "**[Default policies only]** Specifies how often the policy should run and create snapshots or AMIs. The creation frequency can range from 1 to 7 days.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "cross_region_copy_targets": {
        "computed": true,
        "description": "**[Default policies only]** Specifies destination Regions for snapshot or AMI copies. You can specify up to 3 destination Regions. If you do not want to create cross-Region copies, omit this parameter.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "target_region": {
              "computed": true,
              "description": "The target Region, for example ` + "`" + `us-east-1` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "default_policy": {
        "computed": true,
        "description": "**[Default policies only]** Specify the type of default policy to create.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the lifecycle policy. The characters ^[0-9A-Za-z _-]+$ are supported.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "exclusions": {
        "computed": true,
        "description": "**[Default policies only]** Specifies exclusion parameters for volumes or instances for which you do not want to create snapshots or AMIs. The policy will not create snapshots or AMIs for target resources that match any of the specified exclusion parameters.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "exclude_boot_volumes": {
              "computed": true,
              "description": "**[Default policies for EBS snapshots only]** Indicates whether to exclude volumes that are attached to instances as the boot volume. If you exclude boot volumes, only volumes attached as data (non-boot) volumes will be backed up by the policy. To exclude boot volumes, specify ` + "`" + `true` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "exclude_tags": {
              "computed": true,
              "description": "**[Default policies for EBS-backed AMIs only]** Specifies whether to exclude volumes that have specific tags.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "key": {
                    "computed": true,
                    "description": "The tag key.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description": "The tag value.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "exclude_volume_types": {
              "computed": true,
              "description": "**[Default policies for EBS snapshots only]** Specifies the volume types to exclude. Volumes of the specified types will not be targeted by the policy.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "execution_role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the IAM role used to run the operations specified by the lifecycle policy.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "extend_deletion": {
        "computed": true,
        "description": "**[Default policies only]** Defines the snapshot or AMI retention behavior for the policy if the source volume or instance is deleted, or if the policy enters the error, disabled, or deleted state.",
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
      "policy_details": {
        "computed": true,
        "description": "The configuration details of the lifecycle policy.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "actions": {
              "computed": true,
              "description": "**[Event-based policies only]** The actions to be performed when the event-based policy is activated. You can specify only one action per policy.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cross_region_copy": {
                    "computed": true,
                    "description": "The rule for copying shared snapshots across Regions.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "encryption_configuration": {
                          "computed": true,
                          "description": "The encryption settings for the copied snapshot.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "cmk_arn": {
                                "computed": true,
                                "description": "The Amazon Resource Name (ARN) of the AWS KMS key to use for EBS encryption. If this parameter is not specified, the default KMS key for the account is used.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "encrypted": {
                                "computed": true,
                                "description": "To encrypt a copy of an unencrypted snapshot when encryption by default is not enabled, enable encryption using this parameter. Copies of encrypted snapshots are encrypted, even if this parameter is ` + "`" + `false` + "`" + ` or when encryption by default is not enabled.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "retain_rule": {
                          "computed": true,
                          "description": "The retention rule that indicates how long the cross-Region snapshot or AMI copies are to be retained in the destination Region.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "interval": {
                                "computed": true,
                                "description": "The amount of time to retain a cross-Region snapshot or AMI copy. The maximum is 100 years. This is equivalent to 1200 months, 5200 weeks, or 36500 days.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "interval_unit": {
                                "computed": true,
                                "description": "The unit of time for time-based retention. For example, to retain a cross-Region copy for 3 months, specify ` + "`" + `Interval=3` + "`" + ` and ` + "`" + `IntervalUnit=MONTHS` + "`" + `.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "target": {
                          "computed": true,
                          "description": "The target Region.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "name": {
                    "computed": true,
                    "description": "A descriptive name for the action.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "copy_tags": {
              "computed": true,
              "description": "**[Default policies only]** Indicates whether the policy should copy tags from the source resource to the snapshot or AMI. If you do not specify a value, the default is ` + "`" + `false` + "`" + `.\n\nDefault: ` + "`" + `false` + "`" + `",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "create_interval": {
              "computed": true,
              "description": "**[Default policies only]** Specifies how often the policy should run and create snapshots or AMIs. The creation frequency can range from 1 to 7 days. If you do not specify a value, the default is 1.\n\nDefault: 1",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "cross_region_copy_targets": {
              "computed": true,
              "description": "**[Default policies only]** Specifies destination Regions for snapshot or AMI copies. You can specify up to 3 destination Regions. If you do not want to create cross-Region copies, omit this parameter.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "target_region": {
                    "computed": true,
                    "description": "The target Region, for example ` + "`" + `us-east-1` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "event_source": {
              "computed": true,
              "description": "**[Event-based policies only]** The event that activates the event-based policy.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "parameters": {
                    "computed": true,
                    "description": "Information about the event.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "description_regex": {
                          "computed": true,
                          "description": "The snapshot description that can trigger the policy. The description pattern is specified using a regular expression. The policy runs only if a snapshot with a description that matches the specified pattern is shared with your account.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "event_type": {
                          "computed": true,
                          "description": "The type of event. Currently, only snapshot sharing events are supported.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "snapshot_owner": {
                          "computed": true,
                          "description": "The IDs of the AWS accounts that can trigger policy by sharing snapshots with your account. The policy only runs if one of the specified AWS accounts shares a snapshot with your account.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "type": {
                    "computed": true,
                    "description": "The source of the event. Currently only managed Amazon EventBridge events are supported.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "exclusions": {
              "computed": true,
              "description": "**[Default policies only]** Specifies exclusion parameters for volumes or instances for which you do not want to create snapshots or AMIs. The policy will not create snapshots or AMIs for target resources that match any of the specified exclusion parameters.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "exclude_boot_volumes": {
                    "computed": true,
                    "description": "**[Default policies for EBS snapshots only]** Indicates whether to exclude volumes that are attached to instances as the boot volume. If you exclude boot volumes, only volumes attached as data (non-boot) volumes will be backed up by the policy. To exclude boot volumes, specify ` + "`" + `true` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "exclude_tags": {
                    "computed": true,
                    "description": "**[Default policies for EBS-backed AMIs only]** Specifies whether to exclude volumes that have specific tags.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "key": {
                          "computed": true,
                          "description": "The tag key.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
                          "computed": true,
                          "description": "The tag value.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "exclude_volume_types": {
                    "computed": true,
                    "description": "**[Default policies for EBS snapshots only]** Specifies the volume types to exclude. Volumes of the specified types will not be targeted by the policy.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "extend_deletion": {
              "computed": true,
              "description": "**[Default policies only]** Defines the snapshot or AMI retention behavior for the policy if the source volume or instance is deleted, or if the policy enters the error, disabled, or deleted state.\n\nDefault: ` + "`" + `false` + "`" + `",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "parameters": {
              "computed": true,
              "description": "**[Custom snapshot and AMI policies only]** A set of optional parameters for snapshot and AMI lifecycle policies.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "exclude_boot_volume": {
                    "computed": true,
                    "description": "**[Custom snapshot policies that target instances only]** Indicates whether to exclude the root volume from multi-volume snapshot sets. The default is ` + "`" + `false` + "`" + `. If you specify ` + "`" + `true` + "`" + `, then the root volumes attached to targeted instances will be excluded from the multi-volume snapshot sets created by the policy.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "exclude_data_volume_tags": {
                    "computed": true,
                    "description": "**[Custom snapshot policies that target instances only]** The tags used to identify data (non-root) volumes to exclude from multi-volume snapshot sets. If you create a snapshot lifecycle policy that targets instances and you specify tags for this parameter, then data volumes with the specified tags that are attached to targeted instances will be excluded from the multi-volume snapshot sets created by the policy.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "key": {
                          "computed": true,
                          "description": "The tag key.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
                          "computed": true,
                          "description": "The tag value.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "no_reboot": {
                    "computed": true,
                    "description": "**[Custom AMI policies only]** Indicates whether targeted instances are rebooted when the lifecycle policy runs. ` + "`" + `true` + "`" + ` indicates that targeted instances are not rebooted when the policy runs. ` + "`" + `false` + "`" + ` indicates that target instances are rebooted when the policy runs.\n\nThe default is ` + "`" + `true` + "`" + ` (instances are not rebooted).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "policy_language": {
              "computed": true,
              "description": "The type of policy to create. Specify one of the following:\n\n- ` + "`" + `SIMPLIFIED` + "`" + ` -- To create a default policy.\n- ` + "`" + `STANDARD` + "`" + ` -- To create a custom policy.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "policy_type": {
              "computed": true,
              "description": "The type of policy. Specify ` + "`" + `EBS_SNAPSHOT_MANAGEMENT` + "`" + ` to create a lifecycle policy that manages the lifecycle of Amazon EBS snapshots. Specify ` + "`" + `IMAGE_MANAGEMENT` + "`" + ` to create a lifecycle policy that manages the lifecycle of EBS-backed AMIs. Specify ` + "`" + `EVENT_BASED_POLICY` + "`" + ` to create an event-based policy that performs specific actions when a defined event occurs in your AWS account.\n\nThe default is ` + "`" + `EBS_SNAPSHOT_MANAGEMENT` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_locations": {
              "computed": true,
              "description": "**[Custom snapshot and AMI policies only]** The location of the resources to backup. If the source resources are located in a Region, specify ` + "`" + `CLOUD` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "resource_type": {
              "computed": true,
              "description": "**[Default policies only]** Specify the type of default policy to create.\n\n- To create a default policy for EBS snapshots, that creates snapshots of all volumes in the Region that do not have recent backups, specify ` + "`" + `VOLUME` + "`" + `.\n- To create a default policy for EBS-backed AMIs, that creates EBS-backed AMIs from all instances in the Region that do not have recent backups, specify ` + "`" + `INSTANCE` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "resource_types": {
              "computed": true,
              "description": "**[Custom snapshot policies only]** The target resource type for snapshot and AMI lifecycle policies. Use ` + "`" + `VOLUME` + "`" + ` to create snapshots of individual volumes or use ` + "`" + `INSTANCE` + "`" + ` to create multi-volume snapshots from the volumes for an instance.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "retain_interval": {
              "computed": true,
              "description": "**[Default policies only]** Specifies how long the policy should retain snapshots or AMIs before deleting them. The retention period can range from 2 to 14 days, but it must be greater than the creation frequency to ensure that the policy retains at least 1 snapshot or AMI at any given time. If you do not specify a value, the default is 7.\n\nDefault: 7",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "schedules": {
              "computed": true,
              "description": "**[Custom snapshot and AMI policies only]** The schedules of policy-defined actions for snapshot and AMI lifecycle policies. A policy can have up to four schedules -- one mandatory schedule and up to three optional schedules.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "archive_rule": {
                    "computed": true,
                    "description": "**[Custom snapshot policies that target volumes only]** The snapshot archiving rule for the schedule. When you specify an archiving rule, snapshots are automatically moved from the standard tier to the archive tier once the schedule's retention threshold is met. Snapshots are then retained in the archive tier for the archive retention period that you specify.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "retain_rule": {
                          "computed": true,
                          "description": "Information about the retention period for the snapshot archiving rule.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "retention_archive_tier": {
                                "computed": true,
                                "description": "Information about retention period in the Amazon EBS Snapshots Archive.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "count": {
                                      "computed": true,
                                      "description": "The maximum number of snapshots to retain in the archive storage tier for each volume. The count must ensure that each snapshot remains in the archive tier for at least 90 days.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "interval": {
                                      "computed": true,
                                      "description": "Specifies the period of time to retain snapshots in the archive tier. After this period expires, the snapshot is permanently deleted.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "interval_unit": {
                                      "computed": true,
                                      "description": "The unit of time in which to measure the **Interval**. For example, to retain snapshots in the archive tier for 6 months, specify ` + "`" + `Interval=6` + "`" + ` and ` + "`" + `IntervalUnit=MONTHS` + "`" + `.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "copy_tags": {
                    "computed": true,
                    "description": "Copy all user-defined tags on a source volume to snapshots of the volume created by this policy.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "create_rule": {
                    "computed": true,
                    "description": "The creation rule.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "cron_expression": {
                          "computed": true,
                          "description": "The schedule, as a Cron expression. The schedule interval must be between 1 hour and 1 year.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "interval": {
                          "computed": true,
                          "description": "The interval between snapshots. The supported values are 1, 2, 3, 4, 6, 8, 12, and 24.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "interval_unit": {
                          "computed": true,
                          "description": "The interval unit.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "location": {
                          "computed": true,
                          "description": "**[Custom snapshot policies only]** Specifies the destination for snapshots created by the policy. The allowed destinations depend on the location of the targeted resources.\n\n- If the policy targets resources in a Region, then you must create snapshots in the same Region as the source resource.\n- If the policy targets resources in a Local Zone, you can create snapshots in the same Local Zone or in its parent Region.\n- If the policy targets resources on an Outpost, then you can create snapshots on the same Outpost or in its parent Region.\n\nDefault: ` + "`" + `CLOUD` + "`" + `",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "scripts": {
                          "computed": true,
                          "description": "**[Custom snapshot policies that target instances only]** Specifies pre and/or post scripts for a snapshot lifecycle policy that targets instances. This is useful for creating application-consistent snapshots, or for performing specific administrative tasks before or after Amazon Data Lifecycle Manager initiates snapshot creation.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "execute_operation_on_script_failure": {
                                "computed": true,
                                "description": "Indicates whether Amazon Data Lifecycle Manager should default to crash-consistent snapshots if the pre script fails.\n\n- To default to crash consistent snapshot if the pre script fails, specify ` + "`" + `true` + "`" + `.\n- To skip the instance for snapshot creation if the pre script fails, specify ` + "`" + `false` + "`" + `.\n\nThis parameter is supported only if you run a pre script. If you run a post script only, omit this parameter.\n\nDefault: ` + "`" + `true` + "`" + `",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "execution_handler": {
                                "computed": true,
                                "description": "The SSM document that includes the pre and/or post scripts to run.\n\nIf you are automating VSS backups, specify ` + "`" + `AWS_VSS_BACKUP` + "`" + `. In this case, Amazon Data Lifecycle Manager automatically uses the ` + "`" + `AWSEC2-CreateVssSnapshot` + "`" + ` SSM document.\n\nIf you are using a custom SSM document that you own, specify either the name or ARN of the SSM document. If you are using a custom SSM document that is shared with you, specify the ARN of the SSM document.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "execution_handler_service": {
                                "computed": true,
                                "description": "Indicates the service used to execute the pre and/or post scripts.\n\nDefault: ` + "`" + `AWS_SYSTEMS_MANAGER` + "`" + `",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "execution_timeout": {
                                "computed": true,
                                "description": "Specifies a timeout period, in seconds, after which Amazon Data Lifecycle Manager fails the script run attempt if it has not completed. If a script does not complete within its timeout period, Amazon Data Lifecycle Manager fails the attempt. The timeout period applies to the pre and post scripts individually.\n\nDefault: 10",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "maximum_retry_count": {
                                "computed": true,
                                "description": "Specifies the number of times Amazon Data Lifecycle Manager should retry scripts that fail.\n\nIf the pre script fails, Amazon Data Lifecycle Manager retries the entire snapshot creation process, including running the pre and post scripts.\n\nIf the post script fails, Amazon Data Lifecycle Manager retries the post script only; in this case, the pre script will have completed and the snapshot might have been created.\n\nIf you do not want Amazon Data Lifecycle Manager to retry failed scripts, specify ` + "`" + `0` + "`" + `.\n\nDefault: 0",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "stages": {
                                "computed": true,
                                "description": "Indicate which scripts Amazon Data Lifecycle Manager should run on target instances. Pre scripts run before Amazon Data Lifecycle Manager initiates snapshot creation. Post scripts run after Amazon Data Lifecycle Manager initiates snapshot creation.\n\n- To run a pre script only, specify ` + "`" + `PRE` + "`" + `.\n- To run a post script only, specify ` + "`" + `POST` + "`" + `.\n- To run both pre and post scripts, specify both ` + "`" + `PRE` + "`" + ` and ` + "`" + `POST` + "`" + `.\n\nDefault: PRE and POST",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "optional": true
                        },
                        "times": {
                          "computed": true,
                          "description": "The time, in UTC, to start the operation. The supported format is hh:mm.\n\nThe operation occurs within a one-hour window following the specified time. If you do not specify a time, Amazon Data Lifecycle Manager selects a time within the next 24 hours.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "cross_region_copy_rules": {
                    "computed": true,
                    "description": "Specifies a rule for copying snapshots or AMIs across Regions.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "cmk_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the AWS KMS key to use for EBS encryption. If this parameter is not specified, the default KMS key for the account is used.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "copy_tags": {
                          "computed": true,
                          "description": "Indicates whether to copy all user-defined tags from the source snapshot or AMI to the cross-Region copy.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "deprecate_rule": {
                          "computed": true,
                          "description": "**[Custom AMI policies only]** The AMI deprecation rule for cross-Region AMI copies created by the rule.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "interval": {
                                "computed": true,
                                "description": "The period after which to deprecate the cross-Region AMI copies. The period must be less than or equal to the cross-Region AMI copy retention period, and it can't be greater than 10 years. This is equivalent to 120 months, 520 weeks, or 3650 days.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "interval_unit": {
                                "computed": true,
                                "description": "The unit of time in which to measure the **Interval**. For example, to deprecate a cross-Region AMI copy after 3 months, specify ` + "`" + `Interval=3` + "`" + ` and ` + "`" + `IntervalUnit=MONTHS` + "`" + `.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "encrypted": {
                          "computed": true,
                          "description": "To encrypt a copy of an unencrypted snapshot if encryption by default is not enabled, enable encryption using this parameter. Copies of encrypted snapshots are encrypted, even if this parameter is ` + "`" + `false` + "`" + ` or if encryption by default is not enabled.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "retain_rule": {
                          "computed": true,
                          "description": "The retention rule that indicates how long the cross-Region snapshot or AMI copies are to be retained in the destination Region.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "interval": {
                                "computed": true,
                                "description": "The amount of time to retain a cross-Region snapshot or AMI copy. The maximum is 100 years. This is equivalent to 1200 months, 5200 weeks, or 36500 days.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "interval_unit": {
                                "computed": true,
                                "description": "The unit of time for time-based retention. For example, to retain a cross-Region copy for 3 months, specify ` + "`" + `Interval=3` + "`" + ` and ` + "`" + `IntervalUnit=MONTHS` + "`" + `.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "target": {
                          "computed": true,
                          "description": "**[Custom snapshot policies only]** The target Region or the Amazon Resource Name (ARN) of the target Outpost for the snapshot copies.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "target_region": {
                          "computed": true,
                          "description": "**[Custom AMI policies only]** The target Region or the Amazon Resource Name (ARN) of the target Outpost for the AMI copies.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "deprecate_rule": {
                    "computed": true,
                    "description": "**[Custom AMI policies only]** The AMI deprecation rule for the schedule.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "count": {
                          "computed": true,
                          "description": "If the schedule has a count-based retention rule, this parameter specifies the number of oldest AMIs to deprecate. The count must be less than or equal to the schedule's retention count, and it can't be greater than 1000.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "interval": {
                          "computed": true,
                          "description": "If the schedule has an age-based retention rule, this parameter specifies the period after which to deprecate AMIs created by the schedule. The period must be less than or equal to the schedule's retention period, and it can't be greater than 10 years. This is equivalent to 120 months, 520 weeks, or 3650 days.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "interval_unit": {
                          "computed": true,
                          "description": "The unit of time in which to measure the **Interval**.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "fast_restore_rule": {
                    "computed": true,
                    "description": "**[Custom snapshot policies only]** The rule for enabling fast snapshot restore.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "availability_zone_ids": {
                          "computed": true,
                          "description": "The Availability Zone IDs in which to enable fast snapshot restore.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "availability_zones": {
                          "computed": true,
                          "description": "The Availability Zones in which to enable fast snapshot restore.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "count": {
                          "computed": true,
                          "description": "The number of snapshots to be enabled with fast snapshot restore.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "interval": {
                          "computed": true,
                          "description": "The amount of time to enable fast snapshot restore. The maximum is 100 years. This is equivalent to 1200 months, 5200 weeks, or 36500 days.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "interval_unit": {
                          "computed": true,
                          "description": "The unit of time for enabling fast snapshot restore.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "name": {
                    "computed": true,
                    "description": "The name of the schedule.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "retain_rule": {
                    "computed": true,
                    "description": "The retention rule for snapshots or AMIs created by the policy.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "count": {
                          "computed": true,
                          "description": "The number of snapshots to retain for each volume, up to a maximum of 1000. For example if you want to retain a maximum of three snapshots, specify ` + "`" + `3` + "`" + `. When the fourth snapshot is created, the oldest retained snapshot is deleted, or it is moved to the archive tier if you have specified an ` + "`" + `ArchiveRule` + "`" + `.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "interval": {
                          "computed": true,
                          "description": "The amount of time to retain each snapshot. The maximum is 100 years. This is equivalent to 1200 months, 5200 weeks, or 36500 days.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "interval_unit": {
                          "computed": true,
                          "description": "The unit of time for time-based retention. For example, to retain snapshots for 3 months, specify ` + "`" + `Interval=3` + "`" + ` and ` + "`" + `IntervalUnit=MONTHS` + "`" + `. Once the snapshot has been retained for 3 months, it is deleted, or it is moved to the archive tier if you have specified an ` + "`" + `ArchiveRule` + "`" + `.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "share_rules": {
                    "computed": true,
                    "description": "**[Custom snapshot policies only]** The rule for sharing snapshots with other AWS accounts.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "target_accounts": {
                          "computed": true,
                          "description": "The IDs of the AWS accounts with which to share the snapshots.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "unshare_interval": {
                          "computed": true,
                          "description": "The period after which snapshots that are shared with other AWS accounts are automatically unshared.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "unshare_interval_unit": {
                          "computed": true,
                          "description": "The unit of time for the automatic unsharing interval.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "tags_to_add": {
                    "computed": true,
                    "description": "The tags to apply to policy-created resources. These user-defined tags are in addition to the AWS-added lifecycle tags.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "key": {
                          "computed": true,
                          "description": "The tag key.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
                          "computed": true,
                          "description": "The tag value.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "variable_tags": {
                    "computed": true,
                    "description": "**[AMI policies and snapshot policies that target instances only]** A collection of key/value pairs with values determined dynamically when the policy is executed. Keys may be any valid Amazon EC2 tag key. Values must be in one of the two following formats: ` + "`" + `$(instance-id)` + "`" + ` or ` + "`" + `$(timestamp)` + "`" + `. Variable tags are only valid for EBS Snapshot Management -- Instance policies.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "key": {
                          "computed": true,
                          "description": "The tag key.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
                          "computed": true,
                          "description": "The tag value.",
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
                "nesting_mode": "list"
              },
              "optional": true
            },
            "target_tags": {
              "computed": true,
              "description": "**[Custom snapshot and AMI policies only]** The single tag that identifies targeted resources for this policy.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "key": {
                    "computed": true,
                    "description": "The tag key.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description": "The tag value.",
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
      "policy_id": {
        "computed": true,
        "description": "The identifier of the lifecycle policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "retain_interval": {
        "computed": true,
        "description": "**[Default policies only]** Specifies how long the policy should retain snapshots or AMIs before deleting them. The retention period can range from 2 to 14 days, but it must be greater than the creation frequency to ensure that the policy retains at least 1 snapshot or AMI at any given time.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "state": {
        "computed": true,
        "description": "The activation state of the lifecycle policy.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags to apply to the lifecycle policy during creation.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value.",
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
    "description": "Resource Type definition for AWS::DLM::LifecyclePolicy",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDlmLifecyclePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDlmLifecyclePolicy), &result)
	return &result
}
