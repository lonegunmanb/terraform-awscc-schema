package data

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
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  }
                },
                "nesting_mode": "single"
              }
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
                          "type": "string"
                        },
                        "value": {
                          "computed": true,
                          "description": "The number of copies for the inference component",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "maximum_execution_timeout_in_seconds": {
                    "computed": true,
                    "description_kind": "plain",
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
                          "type": "string"
                        },
                        "value": {
                          "computed": true,
                          "description": "The number of copies for the inference component",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "wait_interval_in_seconds": {
                    "computed": true,
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
      "endpoint_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the endpoint the inference component is associated with",
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint_name": {
        "computed": true,
        "description": "The name of the endpoint the inference component is associated with",
        "description_kind": "plain",
        "type": "string"
      },
      "failure_reason": {
        "computed": true,
        "description": "The failure reason if the inference component is in a failed state",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
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
            }
          },
          "nesting_mode": "single"
        }
      },
      "specification": {
        "computed": true,
        "description": "The specification for the inference component",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "base_inference_component_name": {
              "computed": true,
              "description": "The name of the base inference component",
              "description_kind": "plain",
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
                    "type": "number"
                  },
                  "min_memory_required_in_mb": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "number_of_accelerator_devices_required": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "number_of_cpu_cores_required": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "container": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "artifact_url": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
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
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "image": {
                    "computed": true,
                    "description": "The image to use for the container that will be materialized for the inference component",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "model_name": {
              "computed": true,
              "description": "The name of the model to use with the inference component",
              "description_kind": "plain",
              "type": "string"
            },
            "startup_parameters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "container_startup_health_check_timeout_in_seconds": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "model_data_download_timeout_in_seconds": {
                    "computed": true,
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
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "variant_name": {
        "computed": true,
        "description": "The name of the endpoint variant the inference component is associated with",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SageMaker::InferenceComponent",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSagemakerInferenceComponentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerInferenceComponent), &result)
	return &result
}
