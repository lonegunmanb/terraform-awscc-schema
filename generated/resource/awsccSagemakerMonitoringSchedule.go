package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerMonitoringSchedule = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description": "The time at which the schedule was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint_name": {
        "computed": true,
        "description": "The name of the endpoint used to run the monitoring job.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "failure_reason": {
        "computed": true,
        "description": "Contains the reason a monitoring job failed, if it failed.",
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
      "last_modified_time": {
        "computed": true,
        "description": "A timestamp that indicates the last time the monitoring job was modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_monitoring_execution_summary": {
        "computed": true,
        "description": "Describes metadata on the last execution to run, if there was one.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "creation_time": {
              "computed": true,
              "description": "The time at which the monitoring job was created.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "endpoint_name": {
              "computed": true,
              "description": "The name of the endpoint used to run the monitoring job.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "failure_reason": {
              "computed": true,
              "description": "Contains the reason a monitoring job failed, if it failed.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "last_modified_time": {
              "computed": true,
              "description": "A timestamp that indicates the last time the monitoring job was modified.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "monitoring_execution_status": {
              "computed": true,
              "description": "The status of the monitoring job.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "monitoring_schedule_name": {
              "computed": true,
              "description": "The name of the monitoring schedule.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "processing_job_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the monitoring job.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "scheduled_time": {
              "computed": true,
              "description": "The time the monitoring job was scheduled.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "monitoring_schedule_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the monitoring schedule.",
        "description_kind": "plain",
        "type": "string"
      },
      "monitoring_schedule_config": {
        "description": "The configuration object that specifies the monitoring schedule and defines the monitoring job.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "monitoring_job_definition": {
              "computed": true,
              "description": "Defines the monitoring job.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "baseline_config": {
                    "computed": true,
                    "description": "Baseline configuration used to validate that the data conforms to the specified constraints and statistics.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "constraints_resource": {
                          "computed": true,
                          "description": "The baseline constraints resource for a monitoring job.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "s3_uri": {
                                "computed": true,
                                "description": "The Amazon S3 URI for baseline constraint file in Amazon S3 that the current monitoring job should validated against.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "statistics_resource": {
                          "computed": true,
                          "description": "The baseline statistics resource for a monitoring job.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "s3_uri": {
                                "computed": true,
                                "description": "The Amazon S3 URI for the baseline statistics file in Amazon S3 that the current monitoring job should be validated against.",
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
                  },
                  "environment": {
                    "computed": true,
                    "description": "Sets the environment variables in the Docker container",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "monitoring_app_specification": {
                    "computed": true,
                    "description": "Container image configuration object for the monitoring job.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "container_arguments": {
                          "computed": true,
                          "description": "An array of arguments for the container used to run the monitoring job.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "container_entrypoint": {
                          "computed": true,
                          "description": "Specifies the entrypoint for a container used to run the monitoring job.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "image_uri": {
                          "computed": true,
                          "description": "The container image to be run by the monitoring job.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "post_analytics_processor_source_uri": {
                          "computed": true,
                          "description": "An Amazon S3 URI to a script that is called after analysis has been performed. Applicable only for the built-in (first party) containers.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "record_preprocessor_source_uri": {
                          "computed": true,
                          "description": "An Amazon S3 URI to a script that is called per row prior to running analysis. It can base64 decode the payload and convert it into a flatted json so that the built-in container can use the converted data. Applicable only for the built-in (first party) containers",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "monitoring_inputs": {
                    "computed": true,
                    "description": "The array of inputs for the monitoring job.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "batch_transform_input": {
                          "computed": true,
                          "description": "The batch transform input for a monitoring job.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "data_captured_destination_s3_uri": {
                                "computed": true,
                                "description": "A URI that identifies the Amazon S3 storage location where Batch Transform Job captures data.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "dataset_format": {
                                "computed": true,
                                "description": "The dataset format of the data to monitor",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "csv": {
                                      "computed": true,
                                      "description": "The CSV format",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "header": {
                                            "computed": true,
                                            "description": "A boolean flag indicating if given CSV has header",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "bool"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "json": {
                                      "computed": true,
                                      "description": "The Json format",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "line": {
                                            "computed": true,
                                            "description": "A boolean flag indicating if it is JSON line format",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "bool"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "parquet": {
                                      "computed": true,
                                      "description": "A flag indicating if the dataset format is Parquet",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              },
                              "exclude_features_attribute": {
                                "computed": true,
                                "description": "Indexes or names of the features to be excluded from analysis",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "local_path": {
                                "computed": true,
                                "description": "Path to the filesystem where the endpoint data is available to the container.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "s3_data_distribution_type": {
                                "computed": true,
                                "description": "Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key. Defauts to FullyReplicated",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "s3_input_mode": {
                                "computed": true,
                                "description": "Whether the Pipe or File is used as the input mode for transfering data for the monitoring job. Pipe mode is recommended for large datasets. File mode is useful for small files that fit in memory. Defaults to File.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "endpoint_input": {
                          "computed": true,
                          "description": "The endpoint for a monitoring job.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "endpoint_name": {
                                "computed": true,
                                "description": "The name of the endpoint used to run the monitoring job.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "exclude_features_attribute": {
                                "computed": true,
                                "description": "Indexes or names of the features to be excluded from analysis",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "local_path": {
                                "computed": true,
                                "description": "Path to the filesystem where the endpoint data is available to the container.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "s3_data_distribution_type": {
                                "computed": true,
                                "description": "Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key. Defauts to FullyReplicated",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "s3_input_mode": {
                                "computed": true,
                                "description": "Whether the Pipe or File is used as the input mode for transfering data for the monitoring job. Pipe mode is recommended for large datasets. File mode is useful for small files that fit in memory. Defaults to File.",
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
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "monitoring_output_config": {
                    "computed": true,
                    "description": "The output configuration for monitoring jobs.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "kms_key_id": {
                          "computed": true,
                          "description": "The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "monitoring_outputs": {
                          "computed": true,
                          "description": "Monitoring outputs for monitoring jobs. This is where the output of the periodic monitoring jobs is uploaded.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "s3_output": {
                                "computed": true,
                                "description": "Information about where and how to store the results of a monitoring job.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "local_path": {
                                      "computed": true,
                                      "description": "The local path to the Amazon S3 storage location where Amazon SageMaker saves the results of a monitoring job. LocalPath is an absolute path for the output data.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "s3_upload_mode": {
                                      "computed": true,
                                      "description": "Whether to upload the results of the monitoring job continuously or after the job completes.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "s3_uri": {
                                      "computed": true,
                                      "description": "A URI that identifies the Amazon S3 storage location where Amazon SageMaker saves the results of a monitoring job.",
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
                            "nesting_mode": "list"
                          },
                          "optional": true
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "monitoring_resources": {
                    "computed": true,
                    "description": "Identifies the resources to deploy for a monitoring job.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "cluster_config": {
                          "computed": true,
                          "description": "Configuration for the cluster used to run model monitoring jobs.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "instance_count": {
                                "computed": true,
                                "description": "The number of ML compute instances to use in the model monitoring job. For distributed processing jobs, specify a value greater than 1. The default value is 1.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "instance_type": {
                                "computed": true,
                                "description": "The ML compute instance type for the processing job.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "volume_kms_key_id": {
                                "computed": true,
                                "description": "The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance(s) that run the model monitoring job.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "volume_size_in_gb": {
                                "computed": true,
                                "description": "The size of the ML storage volume, in gigabytes, that you want to provision. You must specify sufficient ML storage for your scenario.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
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
                  "network_config": {
                    "computed": true,
                    "description": "Networking options for a job, such as network traffic encryption between containers, whether to allow inbound and outbound network calls to and from containers, and the VPC subnets and security groups to use for VPC-enabled jobs.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "enable_inter_container_traffic_encryption": {
                          "computed": true,
                          "description": "Whether to encrypt all communications between distributed processing jobs. Choose True to encrypt communications. Encryption provides greater security for distributed processing jobs, but the processing might take longer.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "enable_network_isolation": {
                          "computed": true,
                          "description": "Whether to allow inbound and outbound network calls to and from the containers used for the processing job.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "vpc_config": {
                          "computed": true,
                          "description": "Specifies a VPC that your training jobs and hosted models have access to. Control access to and from your training and model containers by configuring the VPC.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "security_group_ids": {
                                "computed": true,
                                "description": "The VPC security group IDs, in the form sg-xxxxxxxx. Specify the security groups for the VPC that is specified in the Subnets field.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "subnets": {
                                "computed": true,
                                "description": "The ID of the subnets in the VPC to which you want to connect to your monitoring jobs.",
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
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "role_arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "stopping_condition": {
                    "computed": true,
                    "description": "Specifies a time limit for how long the monitoring job is allowed to run.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "max_runtime_in_seconds": {
                          "computed": true,
                          "description": "The maximum runtime allowed in seconds.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
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
            "monitoring_job_definition_name": {
              "computed": true,
              "description": "Name of the job definition",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "monitoring_type": {
              "computed": true,
              "description": "The type of monitoring job.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "schedule_config": {
              "computed": true,
              "description": "Configuration details about the monitoring schedule.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "data_analysis_end_time": {
                    "computed": true,
                    "description": "Data Analysis end time, e.g. PT0H",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "data_analysis_start_time": {
                    "computed": true,
                    "description": "Data Analysis start time, e.g. -PT1H",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "schedule_expression": {
                    "computed": true,
                    "description": "A cron expression or 'NOW' that describes details about the monitoring schedule.",
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
        "required": true
      },
      "monitoring_schedule_name": {
        "description": "The name of the monitoring schedule.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "monitoring_schedule_status": {
        "computed": true,
        "description": "The status of a schedule job.",
        "description_kind": "plain",
        "optional": true,
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
              "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
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
    "description": "Resource Type definition for AWS::SageMaker::MonitoringSchedule",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSagemakerMonitoringScheduleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerMonitoringSchedule), &result)
	return &result
}
