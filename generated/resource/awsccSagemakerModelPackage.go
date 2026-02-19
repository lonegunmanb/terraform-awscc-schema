package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerModelPackage = `{
  "block": {
    "attributes": {
      "additional_inference_specifications": {
        "computed": true,
        "description": "An array of additional Inference Specification objects.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "containers": {
              "computed": true,
              "description": "The Amazon ECR registry path of the Docker image that contains the inference code.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "container_hostname": {
                    "computed": true,
                    "description": "The DNS host name for the Docker container.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
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
                  "framework": {
                    "computed": true,
                    "description": "The machine learning framework of the model package container image.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "framework_version": {
                    "computed": true,
                    "description": "The framework version of the Model Package Container Image.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "image": {
                    "computed": true,
                    "description": "The Amazon EC2 Container Registry (Amazon ECR) path where inference code is stored.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "image_digest": {
                    "computed": true,
                    "description": "An MD5 hash of the training algorithm that identifies the Docker image used for training.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "model_data_source": {
                    "computed": true,
                    "description": "Specifies the location of ML model data to deploy during endpoint creation.",
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
                                "optional": true,
                                "type": "string"
                              },
                              "model_access_config": {
                                "computed": true,
                                "description": "Specifies the access configuration file for the ML model.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "accept_eula": {
                                      "computed": true,
                                      "description": "Specifies agreement to the model end-user license agreement (EULA).",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              },
                              "s3_data_type": {
                                "computed": true,
                                "description": "Specifies the type of ML model data to deploy.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "s3_uri": {
                                "computed": true,
                                "description": "Specifies the S3 path of ML model data to deploy.",
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
                  "model_data_url": {
                    "computed": true,
                    "description": "A structure with Model Input details.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "model_input": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "data_input_config": {
                          "computed": true,
                          "description": "The input configuration object for the model.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "nearest_model_name": {
                    "computed": true,
                    "description": "The name of a pre-trained machine learning benchmarked by Amazon SageMaker Inference Recommender model that matches your model.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "description": {
              "computed": true,
              "description": "A description of the additional Inference specification.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "A unique name to identify the additional inference specification. The name must be unique within the list of your additional inference specifications for a particular model package.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "supported_content_types": {
              "computed": true,
              "description": "The supported MIME types for the input data.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "supported_realtime_inference_instance_types": {
              "computed": true,
              "description": "A list of the instance types that are used to generate inferences in real-time",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "supported_response_mime_types": {
              "computed": true,
              "description": "The supported MIME types for the output data.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "supported_transform_instance_types": {
              "computed": true,
              "description": "A list of the instance types on which a transformation job can be run or on which an endpoint can be deployed.",
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
      "additional_inference_specifications_to_add": {
        "computed": true,
        "description": "An array of additional Inference Specification objects.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "containers": {
              "computed": true,
              "description": "The Amazon ECR registry path of the Docker image that contains the inference code.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "container_hostname": {
                    "computed": true,
                    "description": "The DNS host name for the Docker container.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
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
                  "framework": {
                    "computed": true,
                    "description": "The machine learning framework of the model package container image.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "framework_version": {
                    "computed": true,
                    "description": "The framework version of the Model Package Container Image.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "image": {
                    "computed": true,
                    "description": "The Amazon EC2 Container Registry (Amazon ECR) path where inference code is stored.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "image_digest": {
                    "computed": true,
                    "description": "An MD5 hash of the training algorithm that identifies the Docker image used for training.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "model_data_source": {
                    "computed": true,
                    "description": "Specifies the location of ML model data to deploy during endpoint creation.",
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
                                "optional": true,
                                "type": "string"
                              },
                              "model_access_config": {
                                "computed": true,
                                "description": "Specifies the access configuration file for the ML model.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "accept_eula": {
                                      "computed": true,
                                      "description": "Specifies agreement to the model end-user license agreement (EULA).",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              },
                              "s3_data_type": {
                                "computed": true,
                                "description": "Specifies the type of ML model data to deploy.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "s3_uri": {
                                "computed": true,
                                "description": "Specifies the S3 path of ML model data to deploy.",
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
                  "model_data_url": {
                    "computed": true,
                    "description": "A structure with Model Input details.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "model_input": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "data_input_config": {
                          "computed": true,
                          "description": "The input configuration object for the model.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "nearest_model_name": {
                    "computed": true,
                    "description": "The name of a pre-trained machine learning benchmarked by Amazon SageMaker Inference Recommender model that matches your model.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "description": {
              "computed": true,
              "description": "A description of the additional Inference specification.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "A unique name to identify the additional inference specification. The name must be unique within the list of your additional inference specifications for a particular model package.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "supported_content_types": {
              "computed": true,
              "description": "The supported MIME types for the input data.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "supported_realtime_inference_instance_types": {
              "computed": true,
              "description": "A list of the instance types that are used to generate inferences in real-time",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "supported_response_mime_types": {
              "computed": true,
              "description": "The supported MIME types for the output data.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "supported_transform_instance_types": {
              "computed": true,
              "description": "A list of the instance types on which a transformation job can be run or on which an endpoint can be deployed.",
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
      "approval_description": {
        "computed": true,
        "description": "A description provided for the model approval.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "certify_for_marketplace": {
        "computed": true,
        "description": "Whether to certify the model package for listing on AWS Marketplace.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "client_token": {
        "computed": true,
        "description": "A unique token that guarantees that the call to this API is idempotent.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description": "The time at which the model package was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "customer_metadata_properties": {
        "computed": true,
        "description": "The metadata properties associated with the model package versions.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "domain": {
        "computed": true,
        "description": "The machine learning domain of the model package you specified.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "drift_check_baselines": {
        "computed": true,
        "description": "Represents the drift check baselines that can be used when the model monitor is set using the model package.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "bias": {
              "computed": true,
              "description": "Represents the drift check bias baselines that can be used when the model monitor is set using the model package.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "config_file": {
                    "computed": true,
                    "description": "Represents a File Source Object.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "content_digest": {
                          "computed": true,
                          "description": "The digest of the file source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "content_type": {
                          "computed": true,
                          "description": "The type of content stored in the file source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "s3_uri": {
                          "computed": true,
                          "description": "The Amazon S3 URI for the file source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "post_training_constraints": {
                    "computed": true,
                    "description": "Represents a Metric Source Object.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "content_digest": {
                          "computed": true,
                          "description": "The digest of the metric source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "content_type": {
                          "computed": true,
                          "description": "The type of content stored in the metric source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "s3_uri": {
                          "computed": true,
                          "description": "The Amazon S3 URI for the metric source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "pre_training_constraints": {
                    "computed": true,
                    "description": "Represents a Metric Source Object.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "content_digest": {
                          "computed": true,
                          "description": "The digest of the metric source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "content_type": {
                          "computed": true,
                          "description": "The type of content stored in the metric source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "s3_uri": {
                          "computed": true,
                          "description": "The Amazon S3 URI for the metric source.",
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
            "explainability": {
              "computed": true,
              "description": "Contains explainability metrics for a model.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "config_file": {
                    "computed": true,
                    "description": "Represents a File Source Object.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "content_digest": {
                          "computed": true,
                          "description": "The digest of the file source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "content_type": {
                          "computed": true,
                          "description": "The type of content stored in the file source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "s3_uri": {
                          "computed": true,
                          "description": "The Amazon S3 URI for the file source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "constraints": {
                    "computed": true,
                    "description": "Represents a Metric Source Object.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "content_digest": {
                          "computed": true,
                          "description": "The digest of the metric source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "content_type": {
                          "computed": true,
                          "description": "The type of content stored in the metric source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "s3_uri": {
                          "computed": true,
                          "description": "The Amazon S3 URI for the metric source.",
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
            "model_data_quality": {
              "computed": true,
              "description": "Represents the drift check data quality baselines that can be used when the model monitor is set using the model package.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "constraints": {
                    "computed": true,
                    "description": "Represents a Metric Source Object.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "content_digest": {
                          "computed": true,
                          "description": "The digest of the metric source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "content_type": {
                          "computed": true,
                          "description": "The type of content stored in the metric source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "s3_uri": {
                          "computed": true,
                          "description": "The Amazon S3 URI for the metric source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "statistics": {
                    "computed": true,
                    "description": "Represents a Metric Source Object.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "content_digest": {
                          "computed": true,
                          "description": "The digest of the metric source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "content_type": {
                          "computed": true,
                          "description": "The type of content stored in the metric source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "s3_uri": {
                          "computed": true,
                          "description": "The Amazon S3 URI for the metric source.",
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
            "model_quality": {
              "computed": true,
              "description": "Represents the drift check model quality baselines that can be used when the model monitor is set using the model package.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "constraints": {
                    "computed": true,
                    "description": "Represents a Metric Source Object.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "content_digest": {
                          "computed": true,
                          "description": "The digest of the metric source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "content_type": {
                          "computed": true,
                          "description": "The type of content stored in the metric source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "s3_uri": {
                          "computed": true,
                          "description": "The Amazon S3 URI for the metric source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "statistics": {
                    "computed": true,
                    "description": "Represents a Metric Source Object.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "content_digest": {
                          "computed": true,
                          "description": "The digest of the metric source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "content_type": {
                          "computed": true,
                          "description": "The type of content stored in the metric source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "s3_uri": {
                          "computed": true,
                          "description": "The Amazon S3 URI for the metric source.",
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
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "inference_specification": {
        "computed": true,
        "description": "Details about inference jobs that can be run with models based on this model package.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "containers": {
              "computed": true,
              "description": "The Amazon ECR registry path of the Docker image that contains the inference code.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "container_hostname": {
                    "computed": true,
                    "description": "The DNS host name for the Docker container.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
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
                  "framework": {
                    "computed": true,
                    "description": "The machine learning framework of the model package container image.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "framework_version": {
                    "computed": true,
                    "description": "The framework version of the Model Package Container Image.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "image": {
                    "computed": true,
                    "description": "The Amazon EC2 Container Registry (Amazon ECR) path where inference code is stored.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "image_digest": {
                    "computed": true,
                    "description": "An MD5 hash of the training algorithm that identifies the Docker image used for training.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "model_data_source": {
                    "computed": true,
                    "description": "Specifies the location of ML model data to deploy during endpoint creation.",
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
                                "optional": true,
                                "type": "string"
                              },
                              "model_access_config": {
                                "computed": true,
                                "description": "Specifies the access configuration file for the ML model.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "accept_eula": {
                                      "computed": true,
                                      "description": "Specifies agreement to the model end-user license agreement (EULA).",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              },
                              "s3_data_type": {
                                "computed": true,
                                "description": "Specifies the type of ML model data to deploy.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "s3_uri": {
                                "computed": true,
                                "description": "Specifies the S3 path of ML model data to deploy.",
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
                  "model_data_url": {
                    "computed": true,
                    "description": "A structure with Model Input details.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "model_input": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "data_input_config": {
                          "computed": true,
                          "description": "The input configuration object for the model.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "nearest_model_name": {
                    "computed": true,
                    "description": "The name of a pre-trained machine learning benchmarked by Amazon SageMaker Inference Recommender model that matches your model.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "supported_content_types": {
              "computed": true,
              "description": "The supported MIME types for the input data.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "supported_realtime_inference_instance_types": {
              "computed": true,
              "description": "A list of the instance types that are used to generate inferences in real-time",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "supported_response_mime_types": {
              "computed": true,
              "description": "The supported MIME types for the output data.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "supported_transform_instance_types": {
              "computed": true,
              "description": "A list of the instance types on which a transformation job can be run or on which an endpoint can be deployed.",
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
      "last_modified_time": {
        "computed": true,
        "description": "The time at which the model package was last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "metadata_properties": {
        "computed": true,
        "description": "Metadata properties of the tracking entity, trial, or trial component.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "commit_id": {
              "computed": true,
              "description": "The commit ID.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "generated_by": {
              "computed": true,
              "description": "The entity this entity was generated by.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "project_id": {
              "computed": true,
              "description": "The project ID metadata.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "repository": {
              "computed": true,
              "description": "The repository metadata.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "model_approval_status": {
        "computed": true,
        "description": "The approval status of the model package.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "model_card": {
        "computed": true,
        "description": "The model card associated with the model package.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "model_card_content": {
              "computed": true,
              "description": "The content of the model card.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "model_card_status": {
              "computed": true,
              "description": "The approval status of the model card within your organization.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "model_metrics": {
        "computed": true,
        "description": "A structure that contains model metrics reports.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "bias": {
              "computed": true,
              "description": "Contains bias metrics for a model.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "post_training_report": {
                    "computed": true,
                    "description": "Represents a Metric Source Object.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "content_digest": {
                          "computed": true,
                          "description": "The digest of the metric source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "content_type": {
                          "computed": true,
                          "description": "The type of content stored in the metric source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "s3_uri": {
                          "computed": true,
                          "description": "The Amazon S3 URI for the metric source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "pre_training_report": {
                    "computed": true,
                    "description": "Represents a Metric Source Object.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "content_digest": {
                          "computed": true,
                          "description": "The digest of the metric source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "content_type": {
                          "computed": true,
                          "description": "The type of content stored in the metric source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "s3_uri": {
                          "computed": true,
                          "description": "The Amazon S3 URI for the metric source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "report": {
                    "computed": true,
                    "description": "Represents a Metric Source Object.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "content_digest": {
                          "computed": true,
                          "description": "The digest of the metric source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "content_type": {
                          "computed": true,
                          "description": "The type of content stored in the metric source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "s3_uri": {
                          "computed": true,
                          "description": "The Amazon S3 URI for the metric source.",
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
            "explainability": {
              "computed": true,
              "description": "Contains explainability metrics for a model.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "report": {
                    "computed": true,
                    "description": "Represents a Metric Source Object.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "content_digest": {
                          "computed": true,
                          "description": "The digest of the metric source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "content_type": {
                          "computed": true,
                          "description": "The type of content stored in the metric source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "s3_uri": {
                          "computed": true,
                          "description": "The Amazon S3 URI for the metric source.",
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
            "model_data_quality": {
              "computed": true,
              "description": "Metrics that measure the quality of the input data for a model.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "constraints": {
                    "computed": true,
                    "description": "Represents a Metric Source Object.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "content_digest": {
                          "computed": true,
                          "description": "The digest of the metric source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "content_type": {
                          "computed": true,
                          "description": "The type of content stored in the metric source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "s3_uri": {
                          "computed": true,
                          "description": "The Amazon S3 URI for the metric source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "statistics": {
                    "computed": true,
                    "description": "Represents a Metric Source Object.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "content_digest": {
                          "computed": true,
                          "description": "The digest of the metric source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "content_type": {
                          "computed": true,
                          "description": "The type of content stored in the metric source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "s3_uri": {
                          "computed": true,
                          "description": "The Amazon S3 URI for the metric source.",
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
            "model_quality": {
              "computed": true,
              "description": "Metrics that measure the quality of a model.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "constraints": {
                    "computed": true,
                    "description": "Represents a Metric Source Object.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "content_digest": {
                          "computed": true,
                          "description": "The digest of the metric source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "content_type": {
                          "computed": true,
                          "description": "The type of content stored in the metric source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "s3_uri": {
                          "computed": true,
                          "description": "The Amazon S3 URI for the metric source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "statistics": {
                    "computed": true,
                    "description": "Represents a Metric Source Object.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "content_digest": {
                          "computed": true,
                          "description": "The digest of the metric source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "content_type": {
                          "computed": true,
                          "description": "The type of content stored in the metric source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "s3_uri": {
                          "computed": true,
                          "description": "The Amazon S3 URI for the metric source.",
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
      "model_package_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the model package group.",
        "description_kind": "plain",
        "type": "string"
      },
      "model_package_description": {
        "computed": true,
        "description": "The description of the model package.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "model_package_group_name": {
        "computed": true,
        "description": "The name of the model package group.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "model_package_name": {
        "computed": true,
        "description": "The name or arn of the model package.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "model_package_status": {
        "computed": true,
        "description": "The current status of the model package.",
        "description_kind": "plain",
        "type": "string"
      },
      "model_package_status_details": {
        "computed": true,
        "description": "Details about the current status of the model package.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "validation_statuses": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "failure_reason": {
                    "computed": true,
                    "description": "If the overall status is Failed, the reason for the failure.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "The name of the model package for which the overall status is being reported.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "status": {
                    "computed": true,
                    "description": "The current status.",
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
      "model_package_version": {
        "computed": true,
        "description": "The version of the model package.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "sample_payload_url": {
        "computed": true,
        "description": "The Amazon Simple Storage Service (Amazon S3) path where the sample payload are stored pointing to single gzip compressed tar archive.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_config": {
        "computed": true,
        "description": "An optional AWS Key Management Service key to encrypt, decrypt, and re-encrypt model package information for regulated workloads with highly sensitive data.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kms_key_id": {
              "computed": true,
              "description": "The AWS KMS Key ID (KMSKeyId) used for encryption of model package information.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "skip_model_validation": {
        "computed": true,
        "description": "Indicates if you want to skip model validation.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_algorithm_specification": {
        "computed": true,
        "description": "Details about the algorithm that was used to create the model package.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "source_algorithms": {
              "computed": true,
              "description": "A list of algorithms that were used to create a model package.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "algorithm_name": {
                    "computed": true,
                    "description": "The name of an algorithm that was used to create the model package. The algorithm must be either an algorithm resource in your Amazon SageMaker account or an algorithm in AWS Marketplace that you are subscribed to.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "model_data_url": {
                    "computed": true,
                    "description": "The Amazon S3 path where the model artifacts, which result from model training, are stored. This path must point to a single gzip compressed tar archive (.tar.gz suffix).",
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
      "source_uri": {
        "computed": true,
        "description": "The URI of the source for the model package.",
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
              "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "task": {
        "computed": true,
        "description": "The machine learning task your model package accomplishes.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "validation_specification": {
        "computed": true,
        "description": "Specifies configurations for one or more transform jobs that Amazon SageMaker runs to test the model package.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "validation_profiles": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "profile_name": {
                    "computed": true,
                    "description": "The name of the profile for the model package.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "transform_job_definition": {
                    "computed": true,
                    "description": "Defines the input needed to run a transform job using the inference specification specified in the algorithm.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "batch_strategy": {
                          "computed": true,
                          "description": "A string that determines the number of records included in a single mini-batch.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
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
                        "max_concurrent_transforms": {
                          "computed": true,
                          "description": "The maximum number of parallel requests that can be sent to each instance in a transform job. The default value is 1.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "max_payload_in_mb": {
                          "computed": true,
                          "description": "The maximum payload size allowed, in MB. A payload is the data portion of a record (without metadata).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "transform_input": {
                          "computed": true,
                          "description": "Describes the input source of a transform job and the way the transform job consumes it.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "compression_type": {
                                "computed": true,
                                "description": "If your transform data is compressed, specify the compression type. Amazon SageMaker automatically decompresses the data for the transform job accordingly. The default value is None.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "content_type": {
                                "computed": true,
                                "description": "The multipurpose internet mail extension (MIME) type of the data. Amazon SageMaker uses the MIME type with each http call to transfer data to the transform job.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "data_source": {
                                "computed": true,
                                "description": "Describes the input source of a transform job and the way the transform job consumes it.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "s3_data_source": {
                                      "computed": true,
                                      "description": "Describes the S3 data source.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "s3_data_type": {
                                            "computed": true,
                                            "description": "The S3 Data Source Type",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "s3_uri": {
                                            "computed": true,
                                            "description": "Depending on the value specified for the S3DataType, identifies either a key name prefix or a manifest.",
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
                              "split_type": {
                                "computed": true,
                                "description": "The method to use to split the transform job's data files into smaller batches. ",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "transform_output": {
                          "computed": true,
                          "description": "Describes the results of a transform job.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "accept": {
                                "computed": true,
                                "description": "The MIME type used to specify the output data. Amazon SageMaker uses the MIME type with each http call to transfer data from the transform job.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "assemble_with": {
                                "computed": true,
                                "description": "Defines how to assemble the results of the transform job as a single S3 object.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "kms_key_id": {
                                "computed": true,
                                "description": "The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "s3_output_path": {
                                "computed": true,
                                "description": "The Amazon S3 path where you want Amazon SageMaker to store the results of the transform job.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "transform_resources": {
                          "computed": true,
                          "description": "Describes the resources, including ML instance types and ML instance count, to use for transform job.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "instance_count": {
                                "computed": true,
                                "description": "The number of ML compute instances to use in the transform job. For distributed transform jobs, specify a value greater than 1. The default value is 1.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "instance_type": {
                                "computed": true,
                                "description": "The ML compute instance type for the transform job.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "volume_kms_key_id": {
                                "computed": true,
                                "description": "The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt model data on the storage volume attached to the ML compute instance(s) that run the batch transform job.",
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
                "nesting_mode": "list"
              },
              "optional": true
            },
            "validation_role": {
              "computed": true,
              "description": "The IAM roles to be used for the validation of the model package.",
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
    "description": "Resource Type definition for AWS::SageMaker::ModelPackage",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSagemakerModelPackageSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerModelPackage), &result)
	return &result
}
