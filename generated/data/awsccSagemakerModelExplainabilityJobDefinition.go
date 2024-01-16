package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerModelExplainabilityJobDefinition = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description": "The time at which the job definition was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint_name": {
        "computed": true,
        "description": "The name of the endpoint used to run the monitoring job.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "job_definition_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of job definition.",
        "description_kind": "plain",
        "type": "string"
      },
      "job_definition_name": {
        "computed": true,
        "description": "The name of the job definition.",
        "description_kind": "plain",
        "type": "string"
      },
      "job_resources": {
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
                    "description": "The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance(s) that run the model monitoring job.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "volume_size_in_gb": {
                    "computed": true,
                    "description": "The size of the ML storage volume, in gigabytes, that you want to provision. You must specify sufficient ML storage for your scenario.",
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
      "model_explainability_app_specification": {
        "computed": true,
        "description": "Container image configuration object for the monitoring job.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "config_uri": {
              "computed": true,
              "description": "The S3 URI to an analysis configuration file",
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
            "image_uri": {
              "computed": true,
              "description": "The container image to be run by the monitoring job.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "model_explainability_baseline_config": {
        "computed": true,
        "description": "Baseline configuration used to validate that the data conforms to the specified constraints and statistics.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "baselining_job_name": {
              "computed": true,
              "description": "The name of a processing job",
              "description_kind": "plain",
              "type": "string"
            },
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
      "model_explainability_job_input": {
        "computed": true,
        "description": "The inputs for a monitoring job.",
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
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          }
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
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "parquet": {
                          "computed": true,
                          "description": "A flag indicating if the dataset format is Parquet",
                          "description_kind": "plain",
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "features_attribute": {
                    "computed": true,
                    "description": "JSONpath to locate features in JSONlines dataset",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "inference_attribute": {
                    "computed": true,
                    "description": "Index or JSONpath to locate predicted label(s)",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "local_path": {
                    "computed": true,
                    "description": "Path to the filesystem where the endpoint data is available to the container.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "probability_attribute": {
                    "computed": true,
                    "description": "Index or JSONpath to locate probabilities",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "s3_data_distribution_type": {
                    "computed": true,
                    "description": "Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key. Defauts to FullyReplicated",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "s3_input_mode": {
                    "computed": true,
                    "description": "Whether the Pipe or File is used as the input mode for transfering data for the monitoring job. Pipe mode is recommended for large datasets. File mode is useful for small files that fit in memory. Defaults to File.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
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
                    "type": "string"
                  },
                  "features_attribute": {
                    "computed": true,
                    "description": "JSONpath to locate features in JSONlines dataset",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "inference_attribute": {
                    "computed": true,
                    "description": "Index or JSONpath to locate predicted label(s)",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "local_path": {
                    "computed": true,
                    "description": "Path to the filesystem where the endpoint data is available to the container.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "probability_attribute": {
                    "computed": true,
                    "description": "Index or JSONpath to locate probabilities",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "s3_data_distribution_type": {
                    "computed": true,
                    "description": "Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key. Defauts to FullyReplicated",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "s3_input_mode": {
                    "computed": true,
                    "description": "Whether the Pipe or File is used as the input mode for transfering data for the monitoring job. Pipe mode is recommended for large datasets. File mode is useful for small files that fit in memory. Defaults to File.",
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
      "model_explainability_job_output_config": {
        "computed": true,
        "description": "The output configuration for monitoring jobs.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kms_key_id": {
              "computed": true,
              "description": "The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.",
              "description_kind": "plain",
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
                          "type": "string"
                        },
                        "s3_upload_mode": {
                          "computed": true,
                          "description": "Whether to upload the results of the monitoring job continuously or after the job completes.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "s3_uri": {
                          "computed": true,
                          "description": "A URI that identifies the Amazon S3 storage location where Amazon SageMaker saves the results of a monitoring job.",
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
              "description": "Specifies a VPC that your training jobs and hosted models have access to. Control access to and from your training and model containers by configuring the VPC.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "security_group_ids": {
                    "computed": true,
                    "description": "The VPC security group IDs, in the form sg-xxxxxxxx. Specify the security groups for the VPC that is specified in the Subnets field.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "subnets": {
                    "computed": true,
                    "description": "The ID of the subnets in the VPC to which you want to connect to your monitoring jobs.",
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
      "role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.",
        "description_kind": "plain",
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
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
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
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::SageMaker::ModelExplainabilityJobDefinition",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSagemakerModelExplainabilityJobDefinitionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerModelExplainabilityJobDefinition), &result)
	return &result
}
