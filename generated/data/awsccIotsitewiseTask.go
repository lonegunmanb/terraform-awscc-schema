package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotsitewiseTask = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "A description of the task.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The current lifecycle status of the task.",
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
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "task_arn": {
        "computed": true,
        "description": "The ARN of the task.",
        "description_kind": "plain",
        "type": "string"
      },
      "task_configuration": {
        "computed": true,
        "description": "The task execution configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "container_task_configuration": {
              "computed": true,
              "description": "Configuration for running a custom container image on managed compute.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "command": {
                    "computed": true,
                    "description": "The command to execute in the container.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "ecr_uri": {
                    "computed": true,
                    "description": "The Amazon ECR image URI for the task container.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "environment_variables": {
                    "computed": true,
                    "description": "A map of environment variable key-value pairs.",
                    "description_kind": "plain",
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "ephemeral_storage_configuration": {
                    "computed": true,
                    "description": "Configuration for ephemeral storage attached to the container task.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "storage_class": {
                          "computed": true,
                          "description": "The storage type that determines I/O performance characteristics. Family name indicates workload pattern, level number indicates performance within that family.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "storage_size_in_gi_b": {
                          "computed": true,
                          "description": "Storage volume size in GiB.",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "mounts": {
                    "computed": true,
                    "description": "Mounts attached to the container filesystem. Each mount exposes an external data source as a local directory inside the container.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "name": {
                          "computed": true,
                          "description": "A unique name for the mount within the task.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "relative_path": {
                          "computed": true,
                          "description": "The relative path under the service-owned mount root where this mount is attached inside the container.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "source": {
                          "computed": true,
                          "description": "The data source configuration for a mount.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "s3_access_point": {
                                "computed": true,
                                "description": "Configures a mount that reads from an Amazon S3 access point.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "access_point_arn": {
                                      "computed": true,
                                      "description": "The Amazon Resource Name (ARN) of the Amazon S3 access point. The mount reads objects from the bucket associated with this access point. Access is governed by the access point policy and the task execution role's IAM permissions.",
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "prefix": {
                                      "computed": true,
                                      "description": "An object key name prefix. If specified, the mount includes only objects whose keys begin with this prefix. To include all objects at the access point, omit this field.",
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
                        "storage_type": {
                          "computed": true,
                          "description": "The type of storage used for the mount inside the container.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "processing_type": {
                    "computed": true,
                    "description": "The processing type for compute resources.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "processing_unit": {
                    "computed": true,
                    "description": "The processing unit allocation that determines vCPU, memory, and GPU resources.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "task_execution_role": {
                    "computed": true,
                    "description": "The ARN of the IAM role that grants the containerized workload permissions to access AWS resources.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "timeout_seconds": {
                    "computed": true,
                    "description": "The timeout in seconds for task execution. Default: 3600 (1 hour).",
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
      "task_name": {
        "computed": true,
        "description": "The name of the task. Must be unique within the workspace.",
        "description_kind": "plain",
        "type": "string"
      },
      "workspace_name": {
        "computed": true,
        "description": "The name of the workspace.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::IoTSiteWise::Task",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIotsitewiseTaskSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotsitewiseTask), &result)
	return &result
}
