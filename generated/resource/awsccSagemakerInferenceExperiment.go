package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerInferenceExperiment = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the inference experiment.",
        "description_kind": "plain",
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description": "The timestamp at which you created the inference experiment.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_storage_config": {
        "computed": true,
        "description": "The Amazon S3 location and configuration for storing inference request and response data.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "content_type": {
              "computed": true,
              "description": "Configuration specifying how to treat different headers. If no headers are specified SageMaker will by default base64 encode when capturing the data.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "csv_content_types": {
                    "computed": true,
                    "description": "The list of all content type headers that SageMaker will treat as CSV and capture accordingly.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "json_content_types": {
                    "computed": true,
                    "description": "The list of all content type headers that SageMaker will treat as JSON and capture accordingly.",
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
            "destination": {
              "description": "The Amazon S3 bucket where the inference request and response data is stored.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "kms_key": {
              "computed": true,
              "description": "The AWS Key Management Service key that Amazon SageMaker uses to encrypt captured data at rest using Amazon S3 server-side encryption.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "description": {
        "computed": true,
        "description": "The description of the inference experiment.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "desired_state": {
        "computed": true,
        "description": "The desired state of the experiment after starting or stopping operation.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "endpoint_metadata": {
        "computed": true,
        "description": "The metadata of the endpoint on which the inference experiment ran.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "endpoint_config_name": {
              "computed": true,
              "description": "The name of the endpoint configuration.",
              "description_kind": "plain",
              "type": "string"
            },
            "endpoint_name": {
              "computed": true,
              "description": "The name of the endpoint used to run the inference experiment.",
              "description_kind": "plain",
              "type": "string"
            },
            "endpoint_status": {
              "computed": true,
              "description": "The status of the endpoint. For possible values of the status of an endpoint.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "endpoint_name": {
        "description": "The name of the endpoint used to run the inference experiment.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key": {
        "computed": true,
        "description": "The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description": "The timestamp at which you last modified the inference experiment.",
        "description_kind": "plain",
        "type": "string"
      },
      "model_variants": {
        "description": "An array of ModelVariantConfig objects. Each ModelVariantConfig object in the array describes the infrastructure configuration for the corresponding variant.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "infrastructure_config": {
              "description": "The configuration for the infrastructure that the model will be deployed to.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "infrastructure_type": {
                    "description": "The type of the inference experiment that you want to run.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "real_time_inference_config": {
                    "description": "The infrastructure configuration for deploying the model to a real-time inference endpoint.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "instance_count": {
                          "description": "The number of instances of the type specified by InstanceType.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "instance_type": {
                          "description": "The instance type the model is deployed to.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            },
            "model_name": {
              "description": "The name of the Amazon SageMaker Model entity.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "variant_name": {
              "description": "The name of the variant.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "name": {
        "description": "The name for the inference experiment.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "role_arn": {
        "description": "The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to access model artifacts and container images, and manage Amazon SageMaker Inference endpoints for model deployment.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "schedule": {
        "computed": true,
        "description": "The duration for which you want the inference experiment to run.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "end_time": {
              "computed": true,
              "description": "The timestamp at which the inference experiment ended or will end.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "start_time": {
              "computed": true,
              "description": "The timestamp at which the inference experiment started or will start.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "shadow_mode_config": {
        "computed": true,
        "description": "The configuration of ShadowMode inference experiment type. Use this field to specify a production variant which takes all the inference requests, and a shadow variant to which Amazon SageMaker replicates a percentage of the inference requests. For the shadow variant also specify the percentage of requests that Amazon SageMaker replicates.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "shadow_model_variants": {
              "description": "List of shadow variant configurations.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "sampling_percentage": {
                    "description": "The percentage of inference requests that Amazon SageMaker replicates from the production variant to the shadow variant.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "shadow_model_variant_name": {
                    "description": "The name of the shadow variant.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "required": true
            },
            "source_model_variant_name": {
              "description": "The name of the production variant, which takes all the inference requests.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "status": {
        "computed": true,
        "description": "The status of the inference experiment.",
        "description_kind": "plain",
        "type": "string"
      },
      "status_reason": {
        "computed": true,
        "description": "The error message or client-specified reason from the StopInferenceExperiment API, that explains the status of the inference experiment.",
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
              "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "type": {
        "description": "The type of the inference experiment that you want to run.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::SageMaker::InferenceExperiment",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSagemakerInferenceExperimentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerInferenceExperiment), &result)
	return &result
}
