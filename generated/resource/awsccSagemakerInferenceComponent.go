package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerInferenceComponent = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "deployment_config": {
        "computed": true,
        "description": "The deployment config for the inference component",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "auto_rollback_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "alarms": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "alarm_name": {
                          "computed": true,
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
            "rolling_update_policy": {
              "computed": true,
              "description": "The rolling update policy for the inference component",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "maximum_batch_size": {
                    "computed": true,
                    "description": "Capacity size configuration for the inference component",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "type": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
                          "computed": true,
                          "description": "The number of copies for the inference component",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "maximum_execution_timeout_in_seconds": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "rollback_maximum_batch_size": {
                    "computed": true,
                    "description": "Capacity size configuration for the inference component",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "type": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
                          "computed": true,
                          "description": "The number of copies for the inference component",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "wait_interval_in_seconds": {
                    "computed": true,
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
      "endpoint_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the endpoint the inference component is associated with",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "endpoint_name": {
        "description": "The name of the endpoint the inference component is associated with",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "failure_reason": {
        "computed": true,
        "description": "The failure reason if the inference component is in a failed state",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "inference_component_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the inference component",
        "description_kind": "plain",
        "type": "string"
      },
      "inference_component_name": {
        "computed": true,
        "description": "The name of the inference component",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "inference_component_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "runtime_config": {
        "computed": true,
        "description": "The runtime config for the inference component",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "copy_count": {
              "computed": true,
              "description": "The number of copies for the inference component",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "current_copy_count": {
              "computed": true,
              "description": "The number of copies for the inference component",
              "description_kind": "plain",
              "type": "number"
            },
            "desired_copy_count": {
              "computed": true,
              "description": "The number of copies for the inference component",
              "description_kind": "plain",
              "type": "number"
            },
            "placement_status": {
              "computed": true,
              "description": "The placement status of the inference component across instance types",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "current_copy_count": {
                    "computed": true,
                    "description": "The number of copies for the inference component",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "instance_type": {
                    "computed": true,
                    "description": "An ML compute instance type",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "specification": {
        "computed": true,
        "description": "The specification for the inference component, for an endpoint with a single instance type. Specify exactly one of Specification or Specifications. InstanceType is not accepted here; use Specifications for per instance type configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "base_inference_component_name": {
              "computed": true,
              "description": "The name of the base inference component",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "compute_resource_requirements": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "max_memory_required_in_mb": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "min_memory_required_in_mb": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "number_of_accelerator_devices_required": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "number_of_cpu_cores_required": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "container": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "artifact_url": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "container_metrics_config": {
                    "computed": true,
                    "description": "The configuration for container metrics scraping",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "metrics_endpoints": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "metric_publish_frequency_in_seconds": {
                                "computed": true,
                                "description": "The interval, in seconds, at which container metrics scraped from the endpoint are published to Amazon CloudWatch. Valid values per the SageMaker API Reference are 10, 30, 60, 120, 180, 240 and 300; the service validates the value.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "metrics_endpoint_path": {
                                "computed": true,
                                "description": "The path to the Prometheus formatted metrics endpoint exposed by the container",
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
                  "deployed_image": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "resolution_time": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "resolved_image": {
                          "computed": true,
                          "description": "The image to use for the container that will be materialized for the inference component",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "specified_image": {
                          "computed": true,
                          "description": "The image to use for the container that will be materialized for the inference component",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "environment": {
                    "computed": true,
                    "description": "Environment variables to specify on the container",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "image": {
                    "computed": true,
                    "description": "The image to use for the container that will be materialized for the inference component",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "current_data_cache_config": {
              "computed": true,
              "description": "The data caching configuration actually in effect, including a value the service chose rather than the template: SageMaker enables caching automatically on instance types with more than 232 GiB of local NVMe storage, whether or not DataCacheConfig was set. Returned by Describe and not settable; set DataCacheConfig instead.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enable_caching": {
                    "computed": true,
                    "description": "Whether the endpoint caches the model artifacts and container image on each instance it provisions for the inference component",
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "data_cache_config": {
              "computed": true,
              "description": "Settings that affect how the inference component caches data",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enable_caching": {
                    "computed": true,
                    "description": "Whether the endpoint caches the model artifacts and container image on each instance it provisions for the inference component",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "model_name": {
              "computed": true,
              "description": "The name of the model to use with the inference component",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "scheduling_config": {
              "computed": true,
              "description": "The scheduling configuration that determines how inference component copies are placed across available instances",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "availability_zone_balance": {
                    "computed": true,
                    "description": "Configuration for balancing inference component copies across Availability Zones",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "enforcement_mode": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "max_imbalance": {
                          "computed": true,
                          "description": "The maximum allowed difference in the number of inference component copies between any two Availability Zones",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "placement_strategy": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "startup_parameters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "container_startup_health_check_timeout_in_seconds": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "model_data_download_timeout_in_seconds": {
                    "computed": true,
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
      "specifications": {
        "computed": true,
        "description": "A list of specification objects for the inference component, one per instance type. The service requires at least two entries; use the singular Specification for a single instance type.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "compute_resource_requirements": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "max_memory_required_in_mb": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "min_memory_required_in_mb": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "number_of_accelerator_devices_required": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "number_of_cpu_cores_required": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "container": {
              "computed": true,
              "description": "Container specification for one Specifications entry. Distinct from InferenceComponentContainerSpecification: DescribeInferenceComponent returns no per-entry DeployedImage (VERIFIED in us-west-2), so DeployedImage is intentionally omitted here and this definition can never be aggregated into a plural READ response. The singular InferenceComponentContainerSpecification keeps DeployedImage - the service DOES return it there.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "artifact_url": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "container_metrics_config": {
                    "computed": true,
                    "description": "The configuration for container metrics scraping",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "metrics_endpoints": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "metric_publish_frequency_in_seconds": {
                                "computed": true,
                                "description": "The interval, in seconds, at which container metrics scraped from the endpoint are published to Amazon CloudWatch. Valid values per the SageMaker API Reference are 10, 30, 60, 120, 180, 240 and 300; the service validates the value.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "metrics_endpoint_path": {
                                "computed": true,
                                "description": "The path to the Prometheus formatted metrics endpoint exposed by the container",
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
                  "environment": {
                    "computed": true,
                    "description": "Environment variables to specify on the container",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "image": {
                    "computed": true,
                    "description": "The image to use for the container that will be materialized for the inference component",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "current_data_cache_config": {
              "computed": true,
              "description": "The data caching configuration actually in effect for this instance type, including a value the service chose rather than the template: SageMaker enables caching automatically on instance types with more than 232 GiB of local NVMe storage, whether or not DataCacheConfig was set. Returned by Describe and not settable; set DataCacheConfig instead.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enable_caching": {
                    "computed": true,
                    "description": "Whether the endpoint caches the model artifacts and container image on each instance it provisions for the inference component",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "data_cache_config": {
              "computed": true,
              "description": "Settings that affect how the inference component caches data",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enable_caching": {
                    "computed": true,
                    "description": "Whether the endpoint caches the model artifacts and container image on each instance it provisions for the inference component",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "instance_type": {
              "computed": true,
              "description": "An ML compute instance type",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "model_name": {
              "computed": true,
              "description": "The name of the model to use with the inference component",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "scheduling_config": {
              "computed": true,
              "description": "The scheduling configuration that determines how inference component copies are placed across available instances",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "availability_zone_balance": {
                    "computed": true,
                    "description": "Configuration for balancing inference component copies across Availability Zones",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "enforcement_mode": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "max_imbalance": {
                          "computed": true,
                          "description": "The maximum allowed difference in the number of inference component copies between any two Availability Zones",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "placement_strategy": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "startup_parameters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "container_startup_health_check_timeout_in_seconds": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "model_data_download_timeout_in_seconds": {
                    "computed": true,
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
          "nesting_mode": "list"
        },
        "optional": true
      },
      "tags": {
        "computed": true,
        "description": "An array of tags to apply to the resource",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "variant_name": {
        "computed": true,
        "description": "The name of the endpoint variant the inference component is associated with",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::SageMaker::InferenceComponent",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSagemakerInferenceComponentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerInferenceComponent), &result)
	return &result
}
