package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatasyncTask = `{
  "block": {
    "attributes": {
      "cloudwatch_log_group_arn": {
        "computed": true,
        "description": "The ARN of the Amazon CloudWatch log group that is used to monitor and log events in the task.",
        "description_kind": "plain",
        "type": "string"
      },
      "destination_location_arn": {
        "computed": true,
        "description": "The ARN of an AWS storage resource's location.",
        "description_kind": "plain",
        "type": "string"
      },
      "destination_network_interface_arns": {
        "computed": true,
        "description": "The Amazon Resource Names (ARNs) of the destination ENIs (Elastic Network Interfaces) that were created for your subnet.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "excludes": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "filter_type": {
              "computed": true,
              "description": "The type of filter rule to apply. AWS DataSync only supports the SIMPLE_PATTERN rule type.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "A single filter string that consists of the patterns to include or exclude. The patterns are delimited by \"|\".",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "includes": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "filter_type": {
              "computed": true,
              "description": "The type of filter rule to apply. AWS DataSync only supports the SIMPLE_PATTERN rule type.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "A single filter string that consists of the patterns to include or exclude. The patterns are delimited by \"|\".",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "manifest_config": {
        "computed": true,
        "description": "Configures a manifest, which is a list of files or objects that you want DataSync to transfer.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "action": {
              "computed": true,
              "description": "Specifies what DataSync uses the manifest for.",
              "description_kind": "plain",
              "type": "string"
            },
            "format": {
              "computed": true,
              "description": "Specifies the file format of your manifest.",
              "description_kind": "plain",
              "type": "string"
            },
            "source": {
              "computed": true,
              "description": "Specifies the manifest that you want DataSync to use and where it's hosted.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "s3": {
                    "computed": true,
                    "description": "Specifies the S3 bucket where you're hosting the manifest that you want AWS DataSync to use.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bucket_access_role_arn": {
                          "computed": true,
                          "description": "Specifies the AWS Identity and Access Management (IAM) role that allows DataSync to access your manifest.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "manifest_object_path": {
                          "computed": true,
                          "description": "Specifies the Amazon S3 object key of your manifest.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "manifest_object_version_id": {
                          "computed": true,
                          "description": "Specifies the object version ID of the manifest that you want DataSync to use.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "s3_bucket_arn": {
                          "computed": true,
                          "description": "Specifies the Amazon Resource Name (ARN) of the S3 bucket where you're hosting your manifest.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "name": {
        "computed": true,
        "description": "The name of a task. This value is a text reference that is used to identify the task in the console.",
        "description_kind": "plain",
        "type": "string"
      },
      "options": {
        "computed": true,
        "description": "Represents the options that are available to control the behavior of a StartTaskExecution operation.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "atime": {
              "computed": true,
              "description": "A file metadata value that shows the last time a file was accessed (that is, when the file was read or written to).",
              "description_kind": "plain",
              "type": "string"
            },
            "bytes_per_second": {
              "computed": true,
              "description": "A value that limits the bandwidth used by AWS DataSync.",
              "description_kind": "plain",
              "type": "number"
            },
            "gid": {
              "computed": true,
              "description": "The group ID (GID) of the file's owners.",
              "description_kind": "plain",
              "type": "string"
            },
            "log_level": {
              "computed": true,
              "description": "A value that determines the types of logs that DataSync publishes to a log stream in the Amazon CloudWatch log group that you provide.",
              "description_kind": "plain",
              "type": "string"
            },
            "mtime": {
              "computed": true,
              "description": "A value that indicates the last time that a file was modified (that is, a file was written to) before the PREPARING phase.",
              "description_kind": "plain",
              "type": "string"
            },
            "object_tags": {
              "computed": true,
              "description": "A value that determines whether object tags should be read from the source object store and written to the destination object store.",
              "description_kind": "plain",
              "type": "string"
            },
            "overwrite_mode": {
              "computed": true,
              "description": "A value that determines whether files at the destination should be overwritten or preserved when copying files.",
              "description_kind": "plain",
              "type": "string"
            },
            "posix_permissions": {
              "computed": true,
              "description": "A value that determines which users or groups can access a file for a specific purpose such as reading, writing, or execution of the file.",
              "description_kind": "plain",
              "type": "string"
            },
            "preserve_deleted_files": {
              "computed": true,
              "description": "A value that specifies whether files in the destination that don't exist in the source file system should be preserved.",
              "description_kind": "plain",
              "type": "string"
            },
            "preserve_devices": {
              "computed": true,
              "description": "A value that determines whether AWS DataSync should preserve the metadata of block and character devices in the source file system, and recreate the files with that device name and metadata on the destination.",
              "description_kind": "plain",
              "type": "string"
            },
            "security_descriptor_copy_flags": {
              "computed": true,
              "description": "A value that determines which components of the SMB security descriptor are copied during transfer.",
              "description_kind": "plain",
              "type": "string"
            },
            "task_queueing": {
              "computed": true,
              "description": "A value that determines whether tasks should be queued before executing the tasks.",
              "description_kind": "plain",
              "type": "string"
            },
            "transfer_mode": {
              "computed": true,
              "description": "A value that determines whether DataSync transfers only the data and metadata that differ between the source and the destination location, or whether DataSync transfers all the content from the source, without comparing to the destination location.",
              "description_kind": "plain",
              "type": "string"
            },
            "uid": {
              "computed": true,
              "description": "The user ID (UID) of the file's owner.",
              "description_kind": "plain",
              "type": "string"
            },
            "verify_mode": {
              "computed": true,
              "description": "A value that determines whether a data integrity verification should be performed at the end of a task execution after all data and metadata have been transferred.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "schedule": {
        "computed": true,
        "description": "Specifies the schedule you want your task to use for repeated executions.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "schedule_expression": {
              "computed": true,
              "description": "A cron expression that specifies when AWS DataSync initiates a scheduled transfer from a source to a destination location",
              "description_kind": "plain",
              "type": "string"
            },
            "status": {
              "computed": true,
              "description": "Specifies status of a schedule.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "source_location_arn": {
        "computed": true,
        "description": "The ARN of the source location for the task.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_network_interface_arns": {
        "computed": true,
        "description": "The Amazon Resource Names (ARNs) of the source ENIs (Elastic Network Interfaces) that were created for your subnet.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "status": {
        "computed": true,
        "description": "The status of the task that was described.",
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
              "description": "The key for an AWS resource tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for an AWS resource tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "task_arn": {
        "computed": true,
        "description": "The ARN of the task.",
        "description_kind": "plain",
        "type": "string"
      },
      "task_mode": {
        "computed": true,
        "description": "Specifies the task mode for the task.",
        "description_kind": "plain",
        "type": "string"
      },
      "task_report_config": {
        "computed": true,
        "description": "Specifies how you want to configure a task report, which provides detailed information about for your Datasync transfer.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "destination": {
              "computed": true,
              "description": "Specifies where DataSync uploads your task report.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "s3": {
                    "computed": true,
                    "description": "Specifies the Amazon S3 bucket where DataSync uploads your task report.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bucket_access_role_arn": {
                          "computed": true,
                          "description": "Specifies the Amazon Resource Name (ARN) of the IAM policy that allows Datasync to upload a task report to your S3 bucket.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "s3_bucket_arn": {
                          "computed": true,
                          "description": "Specifies the ARN of the S3 bucket where Datasync uploads your report.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "subdirectory": {
                          "computed": true,
                          "description": "Specifies a bucket prefix for your report.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            },
            "object_version_ids": {
              "computed": true,
              "description": "Specifies whether your task report includes the new version of each object transferred into an S3 bucket, this only applies if you enable versioning on your bucket.",
              "description_kind": "plain",
              "type": "string"
            },
            "output_type": {
              "computed": true,
              "description": "Specifies the type of task report that you want.",
              "description_kind": "plain",
              "type": "string"
            },
            "overrides": {
              "computed": true,
              "description": "Customizes the reporting level for aspects of your task report. For example, your report might generally only include errors, but you could specify that you want a list of successes and errors just for the files that Datasync attempted to delete in your destination location.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "deleted": {
                    "computed": true,
                    "description": "Specifies the level of reporting for the files, objects, and directories that Datasync attempted to delete in your destination location. This only applies if you configure your task to delete data in the destination that isn't in the source.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "report_level": {
                          "computed": true,
                          "description": "Specifies whether you want your task report to include only what went wrong with your transfer or a list of what succeeded and didn't.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "skipped": {
                    "computed": true,
                    "description": "Specifies the level of reporting for the files, objects, and directories that Datasync attempted to skip during your transfer.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "report_level": {
                          "computed": true,
                          "description": "Specifies whether you want your task report to include only what went wrong with your transfer or a list of what succeeded and didn't.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "transferred": {
                    "computed": true,
                    "description": "Specifies the level of reporting for the files, objects, and directories that Datasync attempted to transfer.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "report_level": {
                          "computed": true,
                          "description": "Specifies whether you want your task report to include only what went wrong with your transfer or a list of what succeeded and didn't.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "verified": {
                    "computed": true,
                    "description": "Specifies the level of reporting for the files, objects, and directories that Datasync attempted to verify at the end of your transfer. This only applies if you configure your task to verify data during and after the transfer (which Datasync does by default)",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "report_level": {
                          "computed": true,
                          "description": "Specifies whether you want your task report to include only what went wrong with your transfer or a list of what succeeded and didn't.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            },
            "report_level": {
              "computed": true,
              "description": "Specifies whether you want your task report to include only what went wrong with your transfer or a list of what succeeded and didn't.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::DataSync::Task",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDatasyncTaskSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatasyncTask), &result)
	return &result
}
