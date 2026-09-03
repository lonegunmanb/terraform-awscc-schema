package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerEndpointConfig = `{
  "block": {
    "attributes": {
      "async_inference_config": {
        "computed": true,
        "description": "Specifies configuration for how an endpoint performs asynchronous inference.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "client_config": {
              "computed": true,
              "description": "Configures the behavior of the client used by SageMaker to interact with the model container during asynchronous inference.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "max_concurrent_invocations_per_instance": {
                    "computed": true,
                    "description": "The maximum number of concurrent requests sent by the SageMaker client to the model container. If no value is provided, SageMaker will choose an optimal value for you.",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "output_config": {
              "computed": true,
              "description": "Specifies the configuration for asynchronous inference invocation outputs.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "kms_key_id": {
                    "computed": true,
                    "description": "The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt the asynchronous inference output in Amazon S3.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "notification_config": {
                    "computed": true,
                    "description": "Specifies the configuration for notifications of inference results for asynchronous inference.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "error_topic": {
                          "computed": true,
                          "description": "Amazon SNS topic to post a notification to when an inference fails. If no topic is provided, no notification is sent on failure.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "include_inference_response_in": {
                          "computed": true,
                          "description": "The Amazon SNS topics where you want the inference response to be included.",
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "success_topic": {
                          "computed": true,
                          "description": "Amazon SNS topic to post a notification to when an inference completes successfully. If no topic is provided, no notification is sent on success.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "s3_failure_path": {
                    "computed": true,
                    "description": "The Amazon S3 location to upload failure inference responses to.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "s3_output_path": {
                    "computed": true,
                    "description": "The Amazon S3 location to upload inference responses to.",
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
      "data_capture_config": {
        "computed": true,
        "description": "Specifies how to capture endpoint data for model monitor. The data capture configuration applies to all production variants hosted at the endpoint.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "capture_content_type_header": {
              "computed": true,
              "description": "A list of the JSON and CSV content type that the endpoint captures.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "csv_content_types": {
                    "computed": true,
                    "description": "A list of the CSV content types of the data that the endpoint captures. For the endpoint to capture the data, you must also specify the content type when you invoke the endpoint.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "json_content_types": {
                    "computed": true,
                    "description": "A list of the JSON content types of the data that the endpoint captures. For the endpoint to capture the data, you must also specify the content type when you invoke the endpoint.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              }
            },
            "capture_options": {
              "computed": true,
              "description": "Specifies whether the endpoint captures input data to your model, output data from your model, or both.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "capture_mode": {
                    "computed": true,
                    "description": "Specifies whether the endpoint captures input data or output data.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "destination_s3_uri": {
              "computed": true,
              "description": "The S3 bucket where model monitor stores captured data.",
              "description_kind": "plain",
              "type": "string"
            },
            "enable_capture": {
              "computed": true,
              "description": "Set to True to enable data capture.",
              "description_kind": "plain",
              "type": "bool"
            },
            "initial_sampling_percentage": {
              "computed": true,
              "description": "The percentage of data to capture.",
              "description_kind": "plain",
              "type": "number"
            },
            "kms_key_id": {
              "computed": true,
              "description": "The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt the captured data at rest using Amazon S3 server-side encryption.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "enable_network_isolation": {
        "computed": true,
        "description": "Sets whether all model containers deployed to the endpoint are isolated. If they are, no inbound or outbound network calls can be made to or from the model containers.",
        "description_kind": "plain",
        "type": "bool"
      },
      "endpoint_config_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the endpoint configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint_config_name": {
        "computed": true,
        "description": "The name of the endpoint configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "execution_role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker AI can assume to perform actions on your behalf.",
        "description_kind": "plain",
        "type": "string"
      },
      "explainer_config": {
        "computed": true,
        "description": "A parameter to activate explainers.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "clarify_explainer_config": {
              "computed": true,
              "description": "A member of ExplainerConfig that contains configuration parameters for the SageMaker Clarify explainer.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enable_explanations": {
                    "computed": true,
                    "description": "A JMESPath boolean expression used to filter which records to explain. Explanations are activated by default.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "inference_config": {
                    "computed": true,
                    "description": "The inference configuration parameter for the model container.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "content_template": {
                          "computed": true,
                          "description": "A template string used to format a JSON record into an acceptable model container input.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "feature_headers": {
                          "computed": true,
                          "description": "The names of the features. If provided, these are included in the endpoint response payload to help readability of the InvokeEndpoint output.",
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "feature_types": {
                          "computed": true,
                          "description": "A list of data types of the features (optional). Applicable only to NLP explainability. If provided, FeatureTypes must have at least one 'text' string (for example, ['text']). If FeatureTypes is not provided, the explainer infers the feature types based on the baseline data.",
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "features_attribute": {
                          "computed": true,
                          "description": "Provides the JMESPath expression to extract the features from a model container input in JSON Lines format.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "label_attribute": {
                          "computed": true,
                          "description": "A JMESPath expression used to locate the list of label headers in the model container output.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "label_headers": {
                          "computed": true,
                          "description": "For multiclass classification problems, the label headers are the names of the classes. Otherwise, the label header is the name of the predicted label.",
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "label_index": {
                          "computed": true,
                          "description": "A zero-based index used to extract a label header or list of label headers from model container output in CSV format.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "max_payload_in_mb": {
                          "computed": true,
                          "description": "The maximum payload size (MB) allowed of a request from the explainer to the model container. Defaults to 6 MB.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "max_record_count": {
                          "computed": true,
                          "description": "The maximum number of records in a request that the model container can process when querying the model container for the predictions of a synthetic dataset. A record is a unit of input data that inference can be made on, for example, a single line in CSV data.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "probability_attribute": {
                          "computed": true,
                          "description": "A JMESPath expression used to extract the probability (or score) from the model container output if the model container is in JSON Lines format.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "probability_index": {
                          "computed": true,
                          "description": "A zero-based index used to extract a probability value (score) or list from model container output in CSV format. If this value is not provided, the entire model container output will be treated as a probability value (score) or list.",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "shap_config": {
                    "computed": true,
                    "description": "The configuration for SHAP analysis.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "number_of_samples": {
                          "computed": true,
                          "description": "The number of samples to be used for analysis by the Kernal SHAP algorithm.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "seed": {
                          "computed": true,
                          "description": "The starting value used to initialize the random number generator in the explainer. Provide a value for this parameter to obtain a deterministic SHAP result.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "shap_baseline_config": {
                          "computed": true,
                          "description": "The configuration for the SHAP baseline of the Kernal SHAP algorithm.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "mime_type": {
                                "computed": true,
                                "description": "The MIME type of the baseline data. Choose from 'text/csv' or 'application/jsonlines'. Defaults to 'text/csv'.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "shap_baseline": {
                                "computed": true,
                                "description": "The inline SHAP baseline data in string format. ShapBaseline can have one or multiple records to be used as the baseline dataset. The format of the SHAP baseline file should be the same format as the training dataset.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "shap_baseline_uri": {
                                "computed": true,
                                "description": "The uniform resource identifier (URI) of the S3 bucket where the SHAP baseline file is stored. The format of the SHAP baseline file should be the same format as the format of the training dataset.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "text_config": {
                          "computed": true,
                          "description": "A parameter that indicates if text features are treated as text and explanations are provided for individual units of text. Required for natural language processing (NLP) explainability only.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "granularity": {
                                "computed": true,
                                "description": "The unit of granularity for the analysis of text features. For example, if the unit is 'token', then each token (like a word in English) of the text is treated as a feature. SHAP values are computed for each unit/feature.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "language": {
                                "computed": true,
                                "description": "Specifies the language of the text features in ISO 639-1 or ISO 639-3 code of a supported language.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "use_logit": {
                          "computed": true,
                          "description": "A Boolean toggle to indicate if you want to use the logit function (true) or log-odds units (false) for model predictions. Defaults to false.",
                          "description_kind": "plain",
                          "type": "bool"
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
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of an AWS Key Management Service key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "metrics_config": {
        "computed": true,
        "description": "Specifies the metrics that the endpoint publishes to Amazon CloudWatch, the frequency of publication, and whether to enable enhanced or detailed observability metrics.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enable_detailed_observability": {
              "computed": true,
              "description": "Specifies whether to enable detailed observability for the endpoint. When set to true, the endpoint publishes container-level inference metrics, per-GPU metrics, per-instance host metrics, and inference component placement metrics.",
              "description_kind": "plain",
              "type": "bool"
            },
            "enable_enhanced_metrics": {
              "computed": true,
              "description": "Specifies whether to enable enhanced metrics for the endpoint. Enhanced metrics provide utilization and invocation data at instance and container granularity.",
              "description_kind": "plain",
              "type": "bool"
            },
            "metric_publish_frequency_in_seconds": {
              "computed": true,
              "description": "The interval, in seconds, at which the endpoint publishes metrics to Amazon CloudWatch. Valid values are 10, 30, 60, 120, 180, 240, and 300. The default is 60.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "production_variants": {
        "computed": true,
        "description": "A list of ProductionVariant objects, one for each model that you want to host at this endpoint.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "capacity_reservation_config": {
              "computed": true,
              "description": "Settings for the capacity reservation for the compute instances that SageMaker AI reserves for an endpoint.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "capacity_reservation_preference": {
                    "computed": true,
                    "description": "Options that you can choose for the capacity reservation.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "ml_reservation_arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) that uniquely identifies the ML capacity reservation that SageMaker AI applies when it deploys the endpoint.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "container_startup_health_check_timeout_in_seconds": {
              "computed": true,
              "description": "The timeout value, in seconds, for your inference container to pass health check by SageMaker Hosting.",
              "description_kind": "plain",
              "type": "number"
            },
            "core_dump_config": {
              "computed": true,
              "description": "Specifies configuration for a core dump from the model container when the process crashes.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "destination_s3_uri": {
                    "computed": true,
                    "description": "The Amazon S3 bucket to send the core dump to.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "kms_key_id": {
                    "computed": true,
                    "description": "The AWS Key Management Service (AWS KMS) key that SageMaker uses to encrypt the core dump data at rest using Amazon S3 server-side encryption. If you use a KMS key ID or an alias of your KMS key, the SageMaker execution role must include permissions to call kms:Encrypt.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "enable_ssm_access": {
              "computed": true,
              "description": "You can use this parameter to turn on native AWS Systems Manager (SSM) access for a production variant behind an endpoint. By default, SSM access is disabled for all production variants behind an endpoint.",
              "description_kind": "plain",
              "type": "bool"
            },
            "inference_ami_version": {
              "computed": true,
              "description": "Specifies an option from a collection of preconfigured Amazon Machine Image (AMI) images. Each image is configured by AWS with a set of software and driver versions. AWS optimizes these configurations for different machine learning workloads. By selecting an AMI version, you can ensure that your inference environment is compatible with specific software requirements, such as CUDA driver versions, Linux kernel versions, or AWS Neuron driver versions",
              "description_kind": "plain",
              "type": "string"
            },
            "initial_instance_count": {
              "computed": true,
              "description": "Number of instances to launch initially.",
              "description_kind": "plain",
              "type": "number"
            },
            "initial_variant_weight": {
              "computed": true,
              "description": "Determines initial traffic distribution among all of the models that you specify in the endpoint configuration.",
              "description_kind": "plain",
              "type": "number"
            },
            "instance_pools": {
              "computed": true,
              "description": "A list of instance pools for the production variant. Each instance pool specifies an instance type and its priority for provisioning. Use instance pools to configure heterogeneous endpoints that deploy models across multiple instance types.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "instance_type": {
                    "computed": true,
                    "description": "The ML compute instance type for the instance pool.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "model_name_override": {
                    "computed": true,
                    "description": "The name of a SageMaker model to use for this instance pool instead of the model specified for the production variant. Use this to deploy a different model optimized for the instance type in this pool.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "priority": {
                    "computed": true,
                    "description": "The priority for the instance pool. SageMaker attempts to provision instances in order of priority, starting with the lowest value. If instances for a higher-priority pool are unavailable, SageMaker attempts to provision from the next pool. Valid values: 1 to 5, where 1 is the highest priority.",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "instance_type": {
              "computed": true,
              "description": "The ML compute instance type.",
              "description_kind": "plain",
              "type": "string"
            },
            "managed_instance_scaling": {
              "computed": true,
              "description": "Settings that control the range in the number of instances that the endpoint provisions as it scales up or down to accommodate traffic.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "max_instance_count": {
                    "computed": true,
                    "description": "The maximum number of instances that the endpoint can provision when it scales up to accommodate an increase in traffic.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "min_instance_count": {
                    "computed": true,
                    "description": "The minimum number of instances that the endpoint must retain when it scales down to accommodate a decrease in traffic.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "scale_in_policy": {
                    "computed": true,
                    "description": "Configures the scale-in behavior for managed instance scaling.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "cooldown_in_minutes": {
                          "computed": true,
                          "description": "The cooldown period, in minutes, after the last endpoint operation before the endpoint evaluates consolidation scale-in opportunities. Valid values are 5 to 1440. The default is 20.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "maximum_step_size": {
                          "computed": true,
                          "description": "The maximum number of instances that the endpoint can terminate at a time during a consolidation scale-in operation. Valid values are 1 to 100. The default is 1.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "strategy": {
                          "computed": true,
                          "description": "The strategy for scaling in instances. IDLE_RELEASE releases instances that have no hosted inference component copies. CONSOLIDATION consolidates inference component copies onto fewer instances to release more instances.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "status": {
                    "computed": true,
                    "description": "Indicates whether managed instance scaling is enabled.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "model_data_download_timeout_in_seconds": {
              "computed": true,
              "description": "The timeout value, in seconds, to download and extract the model that you want to host from Amazon S3 to the individual inference instance associated with this production variant.",
              "description_kind": "plain",
              "type": "number"
            },
            "model_name": {
              "computed": true,
              "description": "The name of the model that you want to host. This is the name that you specified when creating the model.",
              "description_kind": "plain",
              "type": "string"
            },
            "routing_config": {
              "computed": true,
              "description": "Settings that control how the endpoint routes incoming traffic to the instances that the endpoint hosts.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "prefix_aware_routing_config": {
                    "computed": true,
                    "description": "The configuration for prefix-aware routing. Specify this property only when you set RoutingStrategy to PREFIX_AWARE.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "concurrency_threshold": {
                          "computed": true,
                          "description": "The maximum number of in-flight requests on the target instance before the endpoint routes to another instance. Required when RoutingStrategy is PREFIX_AWARE. Valid values are 1 to 1024.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "prefix_length": {
                          "computed": true,
                          "description": "The maximum length of the prefix used for routing decisions. Required when RoutingStrategy is PREFIX_AWARE. Valid values are 1024 to 65536.",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "routing_strategy": {
                    "computed": true,
                    "description": "Sets how the endpoint routes incoming traffic.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "serverless_config": {
              "computed": true,
              "description": "The serverless configuration for an endpoint. Specifies a serverless endpoint configuration instead of an instance-based endpoint configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "max_concurrency": {
                    "computed": true,
                    "description": "The maximum number of concurrent invocations your serverless endpoint can process.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "memory_size_in_mb": {
                    "computed": true,
                    "description": "The memory size of your serverless endpoint. Valid values are in 1 GB increments: 1024 MB, 2048 MB, 3072 MB, 4096 MB, 5120 MB, or 6144 MB.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "provisioned_concurrency": {
                    "computed": true,
                    "description": "The amount of provisioned concurrency to allocate for the serverless endpoint. Should be less than or equal to MaxConcurrency.",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "variant_instance_provision_timeout_in_seconds": {
              "computed": true,
              "description": "The timeout value, in seconds, for provisioning instances for the production variant. When SageMaker encounters an insufficient capacity error while provisioning instances, it retries with the next instance pool (if configured) or waits until the timeout expires. This timeout applies only to capacity provisioning and does not include the time for model download or container startup.",
              "description_kind": "plain",
              "type": "number"
            },
            "variant_name": {
              "computed": true,
              "description": "The name of the production variant.",
              "description_kind": "plain",
              "type": "string"
            },
            "volume_size_in_gb": {
              "computed": true,
              "description": "The size, in GB, of the ML storage volume attached to individual inference instance associated with the production variant. Currently only Amazon EBS gp2 storage volumes are supported.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "list"
        }
      },
      "shadow_production_variants": {
        "computed": true,
        "description": "Array of ProductionVariant objects. There is one for each model that you want to host at this endpoint in shadow mode with production traffic replicated from the model specified on ProductionVariants. If you use this field, you can only specify one variant for ProductionVariants and one variant for ShadowProductionVariants.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "capacity_reservation_config": {
              "computed": true,
              "description": "Settings for the capacity reservation for the compute instances that SageMaker AI reserves for an endpoint.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "capacity_reservation_preference": {
                    "computed": true,
                    "description": "Options that you can choose for the capacity reservation.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "ml_reservation_arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) that uniquely identifies the ML capacity reservation that SageMaker AI applies when it deploys the endpoint.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "container_startup_health_check_timeout_in_seconds": {
              "computed": true,
              "description": "The timeout value, in seconds, for your inference container to pass health check by SageMaker Hosting.",
              "description_kind": "plain",
              "type": "number"
            },
            "core_dump_config": {
              "computed": true,
              "description": "Specifies configuration for a core dump from the model container when the process crashes.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "destination_s3_uri": {
                    "computed": true,
                    "description": "The Amazon S3 bucket to send the core dump to.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "kms_key_id": {
                    "computed": true,
                    "description": "The AWS Key Management Service (AWS KMS) key that SageMaker uses to encrypt the core dump data at rest using Amazon S3 server-side encryption. If you use a KMS key ID or an alias of your KMS key, the SageMaker execution role must include permissions to call kms:Encrypt.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "enable_ssm_access": {
              "computed": true,
              "description": "You can use this parameter to turn on native AWS Systems Manager (SSM) access for a production variant behind an endpoint. By default, SSM access is disabled for all production variants behind an endpoint.",
              "description_kind": "plain",
              "type": "bool"
            },
            "inference_ami_version": {
              "computed": true,
              "description": "Specifies an option from a collection of preconfigured Amazon Machine Image (AMI) images. Each image is configured by AWS with a set of software and driver versions. AWS optimizes these configurations for different machine learning workloads. By selecting an AMI version, you can ensure that your inference environment is compatible with specific software requirements, such as CUDA driver versions, Linux kernel versions, or AWS Neuron driver versions",
              "description_kind": "plain",
              "type": "string"
            },
            "initial_instance_count": {
              "computed": true,
              "description": "Number of instances to launch initially.",
              "description_kind": "plain",
              "type": "number"
            },
            "initial_variant_weight": {
              "computed": true,
              "description": "Determines initial traffic distribution among all of the models that you specify in the endpoint configuration.",
              "description_kind": "plain",
              "type": "number"
            },
            "instance_pools": {
              "computed": true,
              "description": "A list of instance pools for the production variant. Each instance pool specifies an instance type and its priority for provisioning. Use instance pools to configure heterogeneous endpoints that deploy models across multiple instance types.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "instance_type": {
                    "computed": true,
                    "description": "The ML compute instance type for the instance pool.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "model_name_override": {
                    "computed": true,
                    "description": "The name of a SageMaker model to use for this instance pool instead of the model specified for the production variant. Use this to deploy a different model optimized for the instance type in this pool.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "priority": {
                    "computed": true,
                    "description": "The priority for the instance pool. SageMaker attempts to provision instances in order of priority, starting with the lowest value. If instances for a higher-priority pool are unavailable, SageMaker attempts to provision from the next pool. Valid values: 1 to 5, where 1 is the highest priority.",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "instance_type": {
              "computed": true,
              "description": "The ML compute instance type.",
              "description_kind": "plain",
              "type": "string"
            },
            "managed_instance_scaling": {
              "computed": true,
              "description": "Settings that control the range in the number of instances that the endpoint provisions as it scales up or down to accommodate traffic.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "max_instance_count": {
                    "computed": true,
                    "description": "The maximum number of instances that the endpoint can provision when it scales up to accommodate an increase in traffic.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "min_instance_count": {
                    "computed": true,
                    "description": "The minimum number of instances that the endpoint must retain when it scales down to accommodate a decrease in traffic.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "scale_in_policy": {
                    "computed": true,
                    "description": "Configures the scale-in behavior for managed instance scaling.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "cooldown_in_minutes": {
                          "computed": true,
                          "description": "The cooldown period, in minutes, after the last endpoint operation before the endpoint evaluates consolidation scale-in opportunities. Valid values are 5 to 1440. The default is 20.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "maximum_step_size": {
                          "computed": true,
                          "description": "The maximum number of instances that the endpoint can terminate at a time during a consolidation scale-in operation. Valid values are 1 to 100. The default is 1.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "strategy": {
                          "computed": true,
                          "description": "The strategy for scaling in instances. IDLE_RELEASE releases instances that have no hosted inference component copies. CONSOLIDATION consolidates inference component copies onto fewer instances to release more instances.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "status": {
                    "computed": true,
                    "description": "Indicates whether managed instance scaling is enabled.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "model_data_download_timeout_in_seconds": {
              "computed": true,
              "description": "The timeout value, in seconds, to download and extract the model that you want to host from Amazon S3 to the individual inference instance associated with this production variant.",
              "description_kind": "plain",
              "type": "number"
            },
            "model_name": {
              "computed": true,
              "description": "The name of the model that you want to host. This is the name that you specified when creating the model.",
              "description_kind": "plain",
              "type": "string"
            },
            "routing_config": {
              "computed": true,
              "description": "Settings that control how the endpoint routes incoming traffic to the instances that the endpoint hosts.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "prefix_aware_routing_config": {
                    "computed": true,
                    "description": "The configuration for prefix-aware routing. Specify this property only when you set RoutingStrategy to PREFIX_AWARE.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "concurrency_threshold": {
                          "computed": true,
                          "description": "The maximum number of in-flight requests on the target instance before the endpoint routes to another instance. Required when RoutingStrategy is PREFIX_AWARE. Valid values are 1 to 1024.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "prefix_length": {
                          "computed": true,
                          "description": "The maximum length of the prefix used for routing decisions. Required when RoutingStrategy is PREFIX_AWARE. Valid values are 1024 to 65536.",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "routing_strategy": {
                    "computed": true,
                    "description": "Sets how the endpoint routes incoming traffic.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "serverless_config": {
              "computed": true,
              "description": "The serverless configuration for an endpoint. Specifies a serverless endpoint configuration instead of an instance-based endpoint configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "max_concurrency": {
                    "computed": true,
                    "description": "The maximum number of concurrent invocations your serverless endpoint can process.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "memory_size_in_mb": {
                    "computed": true,
                    "description": "The memory size of your serverless endpoint. Valid values are in 1 GB increments: 1024 MB, 2048 MB, 3072 MB, 4096 MB, 5120 MB, or 6144 MB.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "provisioned_concurrency": {
                    "computed": true,
                    "description": "The amount of provisioned concurrency to allocate for the serverless endpoint. Should be less than or equal to MaxConcurrency.",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "variant_instance_provision_timeout_in_seconds": {
              "computed": true,
              "description": "The timeout value, in seconds, for provisioning instances for the production variant. When SageMaker encounters an insufficient capacity error while provisioning instances, it retries with the next instance pool (if configured) or waits until the timeout expires. This timeout applies only to capacity provisioning and does not include the time for model download or container startup.",
              "description_kind": "plain",
              "type": "number"
            },
            "variant_name": {
              "computed": true,
              "description": "The name of the production variant.",
              "description_kind": "plain",
              "type": "string"
            },
            "volume_size_in_gb": {
              "computed": true,
              "description": "The size, in GB, of the ML storage volume attached to individual inference instance associated with the production variant. Currently only Amazon EBS gp2 storage volumes are supported.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "list"
        }
      },
      "tags": {
        "computed": true,
        "description": "A list of key-value pairs to apply to this resource.",
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
      "vpc_config": {
        "computed": true,
        "description": "Specifies an Amazon Virtual Private Cloud (VPC) that your SageMaker jobs, hosted models, and compute resources have access to. You can control access to and from your resources by configuring a VPC.",
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
              "description": "The ID of the subnets in the VPC to which you want to connect your training job or model.",
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
    "description": "Data Source schema for AWS::SageMaker::EndpointConfig",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSagemakerEndpointConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerEndpointConfig), &result)
	return &result
}
