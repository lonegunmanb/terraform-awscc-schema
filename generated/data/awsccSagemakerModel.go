package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerModel = `{
  "block": {
    "attributes": {
      "containers": {
        "computed": true,
        "description": "Specifies the containers in the inference pipeline.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "container_hostname": {
              "computed": true,
              "description": "This parameter is ignored for models that contain only a PrimaryContainer.\n\nWhen a ContainerDefinition is part of an inference pipeline, the value of the parameter uniquely identifies the container for the purposes of logging and metrics. For information, see [Use Logs and Metrics to Monitor an Inference Pipeline](https://docs.aws.amazon.com/sagemaker/latest/dg/inference-pipeline-logs-metrics.html). If you don't specify a value for this parameter for a ContainerDefinition that is part of an inference pipeline, a unique name is automatically assigned based on the position of the ContainerDefinition in the pipeline. If you specify a value for the ContainerHostName for any ContainerDefinition that is part of an inference pipeline, you must specify a value for the ContainerHostName parameter of every ContainerDefinition in that pipeline.",
              "description_kind": "plain",
              "type": "string"
            },
            "environment": {
              "computed": true,
              "description": "The environment variables to set in the Docker container. Don't include any sensitive data in your environment variables.\n\nThe maximum length of each key and value in the Environment map is 1024 bytes. The maximum length of all keys and values in the map, combined, is 32 KB. If you pass multiple containers to a CreateModel request, then the maximum length of all of their maps, combined, is also 32 KB.",
              "description_kind": "plain",
              "type": "string"
            },
            "image": {
              "computed": true,
              "description": "The path where inference code is stored. This can be either in Amazon EC2 Container Registry or in a Docker registry that is accessible from the same VPC that you configure for your endpoint. If you are using your own custom algorithm instead of an algorithm provided by SageMaker, the inference code must meet SageMaker requirements. SageMaker supports both registry/repository[:tag] and registry/repository[@digest] image path formats. For more information, see [Using Your Own Algorithms with Amazon SageMaker](https://docs.aws.amazon.com/sagemaker/latest/dg/your-algorithms.html).",
              "description_kind": "plain",
              "type": "string"
            },
            "image_config": {
              "computed": true,
              "description": "Specifies whether the model container is in Amazon ECR or a private Docker registry accessible from your Amazon Virtual Private Cloud (VPC).",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "repository_access_mode": {
                    "computed": true,
                    "description": "Set this to one of the following values: Platform - The model image is hosted in Amazon ECR. Vpc - The model image is hosted in a private Docker registry in your VPC.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "repository_auth_config": {
                    "computed": true,
                    "description": "Specifies an authentication configuration for the private docker registry where your model image is hosted. Specify a value for this property only if you specified ` + "`" + `Vpc` + "`" + ` as the value for the ` + "`" + `RepositoryAccessMode` + "`" + ` field of the ` + "`" + `ImageConfig` + "`" + ` object that you passed to a call to ` + "`" + `CreateModel` + "`" + ` and the private Docker registry where the model image is hosted requires authentication.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "repository_credentials_provider_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of an AWS Lambda function that provides credentials to authenticate to the private Docker registry where your model image is hosted. For information about how to create an AWS Lambda function, see [Create a Lambda function with the console](https://docs.aws.amazon.com/lambda/latest/dg/getting-started-create-function.html) in the AWS Lambda Developer Guide",
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
            "inference_specification_name": {
              "computed": true,
              "description": "The inference specification name in the model package version.",
              "description_kind": "plain",
              "type": "string"
            },
            "mode": {
              "computed": true,
              "description": "Whether the container hosts a single model or multiple models.",
              "description_kind": "plain",
              "type": "string"
            },
            "model_data_source": {
              "computed": true,
              "description": "Specifies the location of ML model data to deploy. If specified, you must specify one and only one of the available data sources.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "s3_data_source": {
                    "computed": true,
                    "description": "Specifies the S3 location of ML model data to deploy.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "compression_type": {
                          "computed": true,
                          "description": "Specifies how the ML model data is prepared.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "hub_access_config": {
                          "computed": true,
                          "description": "Configuration information specifying which hub contents have accessible deployment options.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "hub_content_arn": {
                                "computed": true,
                                "description": "The ARN of the hub content for which deployment access is allowed.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "model_access_config": {
                          "computed": true,
                          "description": "The access configuration file to control access to the ML model. You can explicitly accept the model end-user license agreement (EULA) within the ` + "`" + `ModelAccessConfig` + "`" + `.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "accept_eula": {
                                "computed": true,
                                "description": "Specifies agreement to the model end-user license agreement (EULA). The ` + "`" + `AcceptEula` + "`" + ` value must be explicitly defined as ` + "`" + `True` + "`" + ` in order to accept the EULA that this model requires. You are responsible for reviewing and complying with any applicable license terms and making sure they are acceptable for your use case before downloading or using a model.",
                                "description_kind": "plain",
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "s3_data_type": {
                          "computed": true,
                          "description": "Specifies the type of ML model data to deploy.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "s3_uri": {
                          "computed": true,
                          "description": "Specifies the S3 path of ML model data to deploy.",
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
            "model_data_url": {
              "computed": true,
              "description": "The S3 path where the model artifacts, which result from model training, are stored. This path must point to a single gzip compressed tar archive (.tar.gz suffix). The S3 path is required for SageMaker built-in algorithms, but not if you use your own algorithms. For more information on built-in algorithms, see [Common Parameters](https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-algo-docker-registry-paths.html).\n\nIf you provide a value for this parameter, SageMaker uses AWS Security Token Service to download model artifacts from the S3 path you provide. AWS STS is activated in your AWS account by default. If you previously deactivated AWS STS for a region, you need to reactivate AWS STS for that region. For more information, see [Activating and Deactivating AWS STS in an AWS Region](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_enable-regions.html) in the AWS Identity and Access Management User Guide",
              "description_kind": "plain",
              "type": "string"
            },
            "model_package_name": {
              "computed": true,
              "description": "The name or Amazon Resource Name (ARN) of the model package to use to create the model.",
              "description_kind": "plain",
              "type": "string"
            },
            "multi_model_config": {
              "computed": true,
              "description": "Specifies additional configuration for multi-model endpoints.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "model_cache_setting": {
                    "computed": true,
                    "description": "Whether to cache models for a multi-model endpoint. By default, multi-model endpoints cache models so that a model does not have to be loaded into memory each time it is invoked. Some use cases do not benefit from model caching. For example, if an endpoint hosts a large number of models that are each invoked infrequently, the endpoint might perform better if you disable model caching. To disable model caching, set the value of this parameter to ` + "`" + `Disabled` + "`" + `.",
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
      "enable_network_isolation": {
        "computed": true,
        "description": "Isolates the model container. No inbound or outbound network calls can be made to or from the model container.",
        "description_kind": "plain",
        "type": "bool"
      },
      "execution_role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the IAM role that you specified for the model.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "inference_execution_config": {
        "computed": true,
        "description": "Specifies details about how containers in a multi-container endpoint are run.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "mode": {
              "computed": true,
              "description": "How containers in a multi-container are run.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "model_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the model.",
        "description_kind": "plain",
        "type": "string"
      },
      "model_name": {
        "computed": true,
        "description": "The name of the new model.",
        "description_kind": "plain",
        "type": "string"
      },
      "primary_container": {
        "computed": true,
        "description": "Describes the container, as part of model definition.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "container_hostname": {
              "computed": true,
              "description": "This parameter is ignored for models that contain only a PrimaryContainer.\n\nWhen a ContainerDefinition is part of an inference pipeline, the value of the parameter uniquely identifies the container for the purposes of logging and metrics. For information, see [Use Logs and Metrics to Monitor an Inference Pipeline](https://docs.aws.amazon.com/sagemaker/latest/dg/inference-pipeline-logs-metrics.html). If you don't specify a value for this parameter for a ContainerDefinition that is part of an inference pipeline, a unique name is automatically assigned based on the position of the ContainerDefinition in the pipeline. If you specify a value for the ContainerHostName for any ContainerDefinition that is part of an inference pipeline, you must specify a value for the ContainerHostName parameter of every ContainerDefinition in that pipeline.",
              "description_kind": "plain",
              "type": "string"
            },
            "environment": {
              "computed": true,
              "description": "The environment variables to set in the Docker container. Don't include any sensitive data in your environment variables.\n\nThe maximum length of each key and value in the Environment map is 1024 bytes. The maximum length of all keys and values in the map, combined, is 32 KB. If you pass multiple containers to a CreateModel request, then the maximum length of all of their maps, combined, is also 32 KB.",
              "description_kind": "plain",
              "type": "string"
            },
            "image": {
              "computed": true,
              "description": "The path where inference code is stored. This can be either in Amazon EC2 Container Registry or in a Docker registry that is accessible from the same VPC that you configure for your endpoint. If you are using your own custom algorithm instead of an algorithm provided by SageMaker, the inference code must meet SageMaker requirements. SageMaker supports both registry/repository[:tag] and registry/repository[@digest] image path formats. For more information, see [Using Your Own Algorithms with Amazon SageMaker](https://docs.aws.amazon.com/sagemaker/latest/dg/your-algorithms.html).",
              "description_kind": "plain",
              "type": "string"
            },
            "image_config": {
              "computed": true,
              "description": "Specifies whether the model container is in Amazon ECR or a private Docker registry accessible from your Amazon Virtual Private Cloud (VPC).",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "repository_access_mode": {
                    "computed": true,
                    "description": "Set this to one of the following values: Platform - The model image is hosted in Amazon ECR. Vpc - The model image is hosted in a private Docker registry in your VPC.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "repository_auth_config": {
                    "computed": true,
                    "description": "Specifies an authentication configuration for the private docker registry where your model image is hosted. Specify a value for this property only if you specified ` + "`" + `Vpc` + "`" + ` as the value for the ` + "`" + `RepositoryAccessMode` + "`" + ` field of the ` + "`" + `ImageConfig` + "`" + ` object that you passed to a call to ` + "`" + `CreateModel` + "`" + ` and the private Docker registry where the model image is hosted requires authentication.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "repository_credentials_provider_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of an AWS Lambda function that provides credentials to authenticate to the private Docker registry where your model image is hosted. For information about how to create an AWS Lambda function, see [Create a Lambda function with the console](https://docs.aws.amazon.com/lambda/latest/dg/getting-started-create-function.html) in the AWS Lambda Developer Guide",
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
            "inference_specification_name": {
              "computed": true,
              "description": "The inference specification name in the model package version.",
              "description_kind": "plain",
              "type": "string"
            },
            "mode": {
              "computed": true,
              "description": "Whether the container hosts a single model or multiple models.",
              "description_kind": "plain",
              "type": "string"
            },
            "model_data_source": {
              "computed": true,
              "description": "Specifies the location of ML model data to deploy. If specified, you must specify one and only one of the available data sources.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "s3_data_source": {
                    "computed": true,
                    "description": "Specifies the S3 location of ML model data to deploy.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "compression_type": {
                          "computed": true,
                          "description": "Specifies how the ML model data is prepared.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "hub_access_config": {
                          "computed": true,
                          "description": "Configuration information specifying which hub contents have accessible deployment options.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "hub_content_arn": {
                                "computed": true,
                                "description": "The ARN of the hub content for which deployment access is allowed.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "model_access_config": {
                          "computed": true,
                          "description": "The access configuration file to control access to the ML model. You can explicitly accept the model end-user license agreement (EULA) within the ` + "`" + `ModelAccessConfig` + "`" + `.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "accept_eula": {
                                "computed": true,
                                "description": "Specifies agreement to the model end-user license agreement (EULA). The ` + "`" + `AcceptEula` + "`" + ` value must be explicitly defined as ` + "`" + `True` + "`" + ` in order to accept the EULA that this model requires. You are responsible for reviewing and complying with any applicable license terms and making sure they are acceptable for your use case before downloading or using a model.",
                                "description_kind": "plain",
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "s3_data_type": {
                          "computed": true,
                          "description": "Specifies the type of ML model data to deploy.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "s3_uri": {
                          "computed": true,
                          "description": "Specifies the S3 path of ML model data to deploy.",
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
            "model_data_url": {
              "computed": true,
              "description": "The S3 path where the model artifacts, which result from model training, are stored. This path must point to a single gzip compressed tar archive (.tar.gz suffix). The S3 path is required for SageMaker built-in algorithms, but not if you use your own algorithms. For more information on built-in algorithms, see [Common Parameters](https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-algo-docker-registry-paths.html).\n\nIf you provide a value for this parameter, SageMaker uses AWS Security Token Service to download model artifacts from the S3 path you provide. AWS STS is activated in your AWS account by default. If you previously deactivated AWS STS for a region, you need to reactivate AWS STS for that region. For more information, see [Activating and Deactivating AWS STS in an AWS Region](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_enable-regions.html) in the AWS Identity and Access Management User Guide",
              "description_kind": "plain",
              "type": "string"
            },
            "model_package_name": {
              "computed": true,
              "description": "The name or Amazon Resource Name (ARN) of the model package to use to create the model.",
              "description_kind": "plain",
              "type": "string"
            },
            "multi_model_config": {
              "computed": true,
              "description": "Specifies additional configuration for multi-model endpoints.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "model_cache_setting": {
                    "computed": true,
                    "description": "Whether to cache models for a multi-model endpoint. By default, multi-model endpoints cache models so that a model does not have to be loaded into memory each time it is invoked. Some use cases do not benefit from model caching. For example, if an endpoint hosts a large number of models that are each invoked infrequently, the endpoint might perform better if you disable model caching. To disable model caching, set the value of this parameter to ` + "`" + `Disabled` + "`" + `.",
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
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs. You can use tags to categorize your AWS resources in different ways, for example, by purpose, owner, or environment. For more information, see [Tagging AWS Resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html).",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag value.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag key. Tag keys must be unique per resource.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "vpc_config": {
        "computed": true,
        "description": "Specifies an Amazon Virtual Private Cloud (VPC) that your SageMaker jobs, hosted models, and compute resources have access to. You can control access to and from your resources by configuring a VPC. For more information, see [Give SageMaker Access to Resources in your Amazon VPC](https://docs.aws.amazon.com/sagemaker/latest/dg/infrastructure-give-access.html).",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "security_group_ids": {
              "computed": true,
              "description": "The VPC security group IDs, in the form ` + "`" + `sg-xxxxxxxx` + "`" + `. Specify the security groups for the VPC that is specified in the ` + "`" + `Subnets` + "`" + ` field.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "subnets": {
              "computed": true,
              "description": "The ID of the subnets in the VPC to which you want to connect your training job or model. For information about the availability of specific instance types, see [Supported Instance Types and Availability Zones](https://docs.aws.amazon.com/sagemaker/latest/dg/instance-types-az.html).",
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
    "description": "Data Source schema for AWS::SageMaker::Model",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSagemakerModelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerModel), &result)
	return &result
}
