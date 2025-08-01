package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerProcessingJob = `{
  "block": {
    "attributes": {
      "app_specification": {
        "computed": true,
        "description": "Configures the processing job to run a specified Docker container image.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "container_arguments": {
              "computed": true,
              "description": "The arguments for a container used to run a processing job.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "container_entrypoint": {
              "computed": true,
              "description": "The entrypoint for a container used to run a processing job.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "image_uri": {
              "computed": true,
              "description": "The container image to be run by the processing job.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "auto_ml_job_arn": {
        "computed": true,
        "description": "The ARN of an AutoML job associated with this processing job.",
        "description_kind": "plain",
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description": "The time at which the processing job was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "environment": {
        "computed": true,
        "description": "Sets the environment variables in the Docker container",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "exit_message": {
        "computed": true,
        "description": "An optional string, up to one KB in size, that contains metadata from the processing container when the processing job exits.",
        "description_kind": "plain",
        "type": "string"
      },
      "experiment_config": {
        "computed": true,
        "description": "Associates a SageMaker job as a trial component with an experiment and trial.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "experiment_name": {
              "computed": true,
              "description": "The name of an existing experiment to associate with the trial component.",
              "description_kind": "plain",
              "type": "string"
            },
            "run_name": {
              "computed": true,
              "description": "The name of the experiment run to associate with the trial component.",
              "description_kind": "plain",
              "type": "string"
            },
            "trial_component_display_name": {
              "computed": true,
              "description": "The display name for the trial component. If this key isn't specified, the display name is the trial component name.",
              "description_kind": "plain",
              "type": "string"
            },
            "trial_name": {
              "computed": true,
              "description": "The name of an existing trial to associate the trial component with. If not specified, a new trial is created.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "failure_reason": {
        "computed": true,
        "description": "A string, up to one KB in size, that contains the reason a processing job failed, if it failed.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description": "The time at which the processing job was last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "monitoring_schedule_arn": {
        "computed": true,
        "description": "The ARN of a monitoring schedule for an endpoint associated with this processing job.",
        "description_kind": "plain",
        "type": "string"
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
              "type": "bool"
            },
            "enable_network_isolation": {
              "computed": true,
              "description": "Whether to allow inbound and outbound network calls to and from the containers used for the processing job.",
              "description_kind": "plain",
              "type": "bool"
            },
            "vpc_config": {
              "computed": true,
              "description": "Specifies an Amazon Virtual Private Cloud (VPC) that your SageMaker jobs, hosted models, and compute resources have access to. You can control access to and from your resources by configuring a VPC. For more information, see https://docs.aws.amazon.com/sagemaker/latest/dg/infrastructure-give-access.html",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "security_group_ids": {
                    "computed": true,
                    "description": "The VPC security group IDs, in the form 'sg-xxxxxxxx'. Specify the security groups for the VPC that is specified in the 'Subnets' field.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "subnets": {
                    "computed": true,
                    "description": "The ID of the subnets in the VPC to which you want to connect your training job or model. For information about the availability of specific instance types, see https://docs.aws.amazon.com/sagemaker/latest/dg/regions-quotas.html",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "processing_end_time": {
        "computed": true,
        "description": "The time at which the processing job completed.",
        "description_kind": "plain",
        "type": "string"
      },
      "processing_inputs": {
        "computed": true,
        "description": "An array of inputs configuring the data to download into the processing container.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "app_managed": {
              "computed": true,
              "description": "When True, input operations such as data download are managed natively by the processing job application. When False (default), input operations are managed by Amazon SageMaker.",
              "description_kind": "plain",
              "type": "bool"
            },
            "dataset_definition": {
              "computed": true,
              "description": "Configuration for Dataset Definition inputs. The Dataset Definition input must specify exactly one of either ` + "`" + `AthenaDatasetDefinition` + "`" + ` or ` + "`" + `RedshiftDatasetDefinition` + "`" + ` types.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "athena_dataset_definition": {
                    "computed": true,
                    "description": "Configuration for Athena Dataset Definition input.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "catalog": {
                          "computed": true,
                          "description": "The name of the data catalog used in Athena query execution.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "database": {
                          "computed": true,
                          "description": "The name of the database used in the Athena query execution.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "kms_key_id": {
                          "computed": true,
                          "description": "The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt data generated from an Athena query execution.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "output_compression": {
                          "computed": true,
                          "description": "The compression used for Athena query results.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "output_format": {
                          "computed": true,
                          "description": "The data storage format for Athena query results.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "output_s3_uri": {
                          "computed": true,
                          "description": "The location in Amazon S3 where Athena query results are stored.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "query_string": {
                          "computed": true,
                          "description": "The SQL query statements, to be executed.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "work_group": {
                          "computed": true,
                          "description": "The name of the workgroup in which the Athena query is being started.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "data_distribution_type": {
                    "computed": true,
                    "description": "Whether the generated dataset is FullyReplicated or ShardedByS3Key (default).",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "input_mode": {
                    "computed": true,
                    "description": "Whether to use File or Pipe input mode. In File (default) mode, Amazon SageMaker copies the data from the input source onto the local Amazon Elastic Block Store (Amazon EBS) volumes before starting your training algorithm. This is the most commonly used input mode. In Pipe mode, Amazon SageMaker streams input data from the source directly to your algorithm without using the EBS volume.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "local_path": {
                    "computed": true,
                    "description": "The local path where you want Amazon SageMaker to download the Dataset Definition inputs to run a processing job. LocalPath is an absolute path to the input data. This is a required parameter when AppManaged is False (default).",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "redshift_dataset_definition": {
                    "computed": true,
                    "description": "Configuration for Redshift Dataset Definition input.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "cluster_id": {
                          "computed": true,
                          "description": "The Redshift cluster Identifier.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "cluster_role_arn": {
                          "computed": true,
                          "description": "The IAM role attached to your Redshift cluster that Amazon SageMaker uses to generate datasets.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "database": {
                          "computed": true,
                          "description": "The name of the Redshift database used in Redshift query execution.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "db_user": {
                          "computed": true,
                          "description": "The database user name used in Redshift query execution.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "kms_key_id": {
                          "computed": true,
                          "description": "The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt data from a Redshift execution.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "output_compression": {
                          "computed": true,
                          "description": "The compression used for Redshift query results.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "output_format": {
                          "computed": true,
                          "description": "The data storage format for Redshift query results.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "output_s3_uri": {
                          "computed": true,
                          "description": "The location in Amazon S3 where the Redshift query results are stored.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "query_string": {
                          "computed": true,
                          "description": "The SQL query statements to be executed.",
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
            "input_name": {
              "computed": true,
              "description": "The name for the processing job input.",
              "description_kind": "plain",
              "type": "string"
            },
            "s3_input": {
              "computed": true,
              "description": "Configuration for downloading input data from Amazon S3 into the processing container.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "local_path": {
                    "computed": true,
                    "description": "The local path in your container where you want Amazon SageMaker to write input data to. ` + "`" + `LocalPath` + "`" + ` is an absolute path to the input data and must begin with ` + "`" + `/opt/ml/processing/` + "`" + `. LocalPath is a required parameter when ` + "`" + `AppManaged` + "`" + ` is ` + "`" + `False` + "`" + ` (default).",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "s3_compression_type": {
                    "computed": true,
                    "description": "Whether to GZIP-decompress the data in Amazon S3 as it is streamed into the processing container. ` + "`" + `Gzip` + "`" + ` can only be used when ` + "`" + `Pipe` + "`" + ` mode is specified as the ` + "`" + `S3InputMode` + "`" + `. In ` + "`" + `Pipe` + "`" + ` mode, Amazon SageMaker streams input data from the source directly to your container without using the EBS volume.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "s3_data_distribution_type": {
                    "computed": true,
                    "description": "Whether to distribute the data from Amazon S3 to all processing instances with ` + "`" + `FullyReplicated` + "`" + `, or whether the data from Amazon S3 is shared by Amazon S3 key, downloading one shard of data to each processing instance.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "s3_data_type": {
                    "computed": true,
                    "description": "Whether you use an S3Prefix or a ManifestFile for the data type. If you choose S3Prefix, S3Uri identifies a key name prefix. Amazon SageMaker uses all objects with the specified key name prefix for the processing job. If you choose ManifestFile, S3Uri identifies an object that is a manifest file containing a list of object keys that you want Amazon SageMaker to use for the processing job.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "s3_input_mode": {
                    "computed": true,
                    "description": "Whether to use File or Pipe input mode. In File mode, Amazon SageMaker copies the data from the input source onto the local ML storage volume before starting your processing container. This is the most commonly used input mode. In Pipe mode, Amazon SageMaker streams input data from the source directly to your processing container into named pipes without using the ML storage volume.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "s3_uri": {
                    "computed": true,
                    "description": "The URI of the Amazon S3 prefix Amazon SageMaker downloads data required to run a processing job.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "list"
        }
      },
      "processing_job_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the processing job.",
        "description_kind": "plain",
        "type": "string"
      },
      "processing_job_name": {
        "computed": true,
        "description": "The name of the processing job. The name must be unique within an AWS Region in the AWS account.",
        "description_kind": "plain",
        "type": "string"
      },
      "processing_job_status": {
        "computed": true,
        "description": "Provides the status of a processing job.",
        "description_kind": "plain",
        "type": "string"
      },
      "processing_output_config": {
        "computed": true,
        "description": "Configuration for uploading output from the processing container.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kms_key_id": {
              "computed": true,
              "description": "The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt the processing job output. KmsKeyId can be an ID of a KMS key, ARN of a KMS key, or alias of a KMS key. The KmsKeyId is applied to all outputs.",
              "description_kind": "plain",
              "type": "string"
            },
            "outputs": {
              "computed": true,
              "description": "An array of outputs configuring the data to upload from the processing container.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "app_managed": {
                    "computed": true,
                    "description": "When True, output operations such as data upload are managed natively by the processing job application. When False (default), output operations are managed by Amazon SageMaker.",
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "feature_store_output": {
                    "computed": true,
                    "description": "Configuration for processing job outputs in Amazon SageMaker Feature Store.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "feature_group_name": {
                          "computed": true,
                          "description": "The name of the Amazon SageMaker FeatureGroup to use as the destination for processing job output. Note that your processing script is responsible for putting records into your Feature Store.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "output_name": {
                    "computed": true,
                    "description": "The name for the processing job output.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "s3_output": {
                    "computed": true,
                    "description": "Configuration for uploading output data to Amazon S3 from the processing container.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "local_path": {
                          "computed": true,
                          "description": "The local path of a directory where you want Amazon SageMaker to upload its contents to Amazon S3. LocalPath is an absolute path to a directory containing output files. This directory will be created by the platform and exist when your container's entrypoint is invoked.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "s3_upload_mode": {
                          "computed": true,
                          "description": "Whether to upload the results of the processing job continuously or after the job completes.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "s3_uri": {
                          "computed": true,
                          "description": "A URI that identifies the Amazon S3 bucket where you want Amazon SageMaker to save the results of a processing job.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "processing_resources": {
        "computed": true,
        "description": "Identifies the resources, ML compute instances, and ML storage volumes to deploy for a processing job. In distributed training, you specify more than one instance.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cluster_config": {
              "computed": true,
              "description": "Configuration for the cluster used to run a processing job.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "instance_count": {
                    "computed": true,
                    "description": "The number of ML compute instances to use in the processing job. For distributed processing jobs, specify a value greater than 1. The default value is 1.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "instance_type": {
                    "computed": true,
                    "description": "The ML compute instance type for the processing job.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "volume_kms_key_id": {
                    "computed": true,
                    "description": "The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance(s) that run the processing job.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "volume_size_in_gb": {
                    "computed": true,
                    "description": "The size of the ML storage volume in gigabytes that you want to provision. You must specify sufficient ML storage for your scenario.",
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
      "processing_start_time": {
        "computed": true,
        "description": "The time at which the processing job started.",
        "description_kind": "plain",
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.",
        "description_kind": "plain",
        "type": "string"
      },
      "stopping_condition": {
        "computed": true,
        "description": "Configures conditions under which the processing job should be stopped, such as how long the processing job has been running. After the condition is met, the processing job is stopped.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "max_runtime_in_seconds": {
              "computed": true,
              "description": "Specifies the maximum runtime in seconds.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description": "(Optional) An array of key-value pairs. For more information, see Using Cost Allocation Tags(https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-whatURL) in the AWS Billing and Cost Management User Guide.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key. Tag keys must be unique per resource.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "training_job_arn": {
        "computed": true,
        "description": "The ARN of a training job associated with this processing job",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SageMaker::ProcessingJob",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSagemakerProcessingJobSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerProcessingJob), &result)
	return &result
}
