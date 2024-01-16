package resource

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
        "description": "The Name of the AppImageConfig.",
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
                    "optional": true,
                    "type": "number"
                  },
                  "default_uid": {
                    "computed": true,
                    "description": "The default POSIX user ID (UID). If not specified, defaults to 1000.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "mount_path": {
                    "computed": true,
                    "description": "The path within the image to mount the user's EFS home directory. The directory should be empty. If not specified, defaults to /home/sagemaker-user.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "kernel_specs": {
              "description": "The specification of the Jupyter kernels in the image.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "display_name": {
                    "computed": true,
                    "description": "The display name of the kernel.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "description": "The name of the kernel.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "required": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "tags": {
        "computed": true,
        "description": "A list of tags to apply to the AppImageConfig.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::SageMaker::AppImageConfig",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSagemakerAppImageConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerAppImageConfig), &result)
	return &result
}
