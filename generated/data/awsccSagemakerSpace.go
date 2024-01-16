package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerSpace = `{
  "block": {
    "attributes": {
      "domain_id": {
        "computed": true,
        "description": "The ID of the associated Domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "space_arn": {
        "computed": true,
        "description": "The space Amazon Resource Name (ARN).",
        "description_kind": "plain",
        "type": "string"
      },
      "space_name": {
        "computed": true,
        "description": "A name for the Space.",
        "description_kind": "plain",
        "type": "string"
      },
      "space_settings": {
        "computed": true,
        "description": "A collection of settings.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "jupyter_server_app_settings": {
              "computed": true,
              "description": "The Jupyter server's app settings.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "default_resource_spec": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "instance_type": {
                          "computed": true,
                          "description": "The instance type that the image version runs on.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "sage_maker_image_arn": {
                          "computed": true,
                          "description": "The ARN of the SageMaker image that the image version belongs to.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "sage_maker_image_version_arn": {
                          "computed": true,
                          "description": "The ARN of the image version created on the instance.",
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
            "kernel_gateway_app_settings": {
              "computed": true,
              "description": "The kernel gateway app settings.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "custom_images": {
                    "computed": true,
                    "description": "A list of custom SageMaker images that are configured to run as a KernelGateway app.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "app_image_config_name": {
                          "computed": true,
                          "description": "The Name of the AppImageConfig.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "image_name": {
                          "computed": true,
                          "description": "The name of the CustomImage. Must be unique to your account.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "image_version_number": {
                          "computed": true,
                          "description": "The version number of the CustomImage.",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "default_resource_spec": {
                    "computed": true,
                    "description": "The default instance type and the Amazon Resource Name (ARN) of the default SageMaker image used by the KernelGateway app.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "instance_type": {
                          "computed": true,
                          "description": "The instance type that the image version runs on.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "sage_maker_image_arn": {
                          "computed": true,
                          "description": "The ARN of the SageMaker image that the image version belongs to.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "sage_maker_image_version_arn": {
                          "computed": true,
                          "description": "The ARN of the image version created on the instance.",
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
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description": "A list of tags to apply to the space.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::SageMaker::Space",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSagemakerSpaceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerSpace), &result)
	return &result
}
