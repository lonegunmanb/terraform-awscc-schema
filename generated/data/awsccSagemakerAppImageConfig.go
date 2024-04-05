package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerAppImageConfig = `{
  "block": {
    "attributes": {
      "app_image_config_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the AppImageConfig.",
        "description_kind": "plain",
        "type": "string"
      },
      "app_image_config_name": {
        "computed": true,
        "description": "The Name of the AppImageConfig.",
        "description_kind": "plain",
        "type": "string"
      },
      "code_editor_app_image_config": {
        "computed": true,
        "description": "The CodeEditorAppImageConfig.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "container_config": {
              "computed": true,
              "description": "The container configuration for a SageMaker image.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "container_arguments": {
                    "computed": true,
                    "description": "A list of arguments to apply to the container.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "container_entrypoint": {
                    "computed": true,
                    "description": "The custom entry point to use on container.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "container_environment_variables": {
                    "computed": true,
                    "description": "A list of variables to apply to the custom container.",
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
      "jupyter_lab_app_image_config": {
        "computed": true,
        "description": "The JupyterLabAppImageConfig.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "container_config": {
              "computed": true,
              "description": "The container configuration for a SageMaker image.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "container_arguments": {
                    "computed": true,
                    "description": "A list of arguments to apply to the container.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "container_entrypoint": {
                    "computed": true,
                    "description": "The custom entry point to use on container.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "container_environment_variables": {
                    "computed": true,
                    "description": "A list of variables to apply to the custom container.",
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
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "kernel_gateway_image_config": {
        "computed": true,
        "description": "The KernelGatewayImageConfig.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "file_system_config": {
              "computed": true,
              "description": "The Amazon Elastic File System (EFS) storage configuration for a SageMaker image.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "default_gid": {
                    "computed": true,
                    "description": "The default POSIX group ID (GID). If not specified, defaults to 100.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "default_uid": {
                    "computed": true,
                    "description": "The default POSIX user ID (UID). If not specified, defaults to 1000.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "mount_path": {
                    "computed": true,
                    "description": "The path within the image to mount the user's EFS home directory. The directory should be empty. If not specified, defaults to /home/sagemaker-user.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "kernel_specs": {
              "computed": true,
              "description": "The specification of the Jupyter kernels in the image.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "display_name": {
                    "computed": true,
                    "description": "The display name of the kernel.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "The name of the kernel.",
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
      "tags": {
        "computed": true,
        "description": "A list of tags to apply to the AppImageConfig.",
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
    "description": "Data Source schema for AWS::SageMaker::AppImageConfig",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSagemakerAppImageConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerAppImageConfig), &result)
	return &result
}
