package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccImagebuilderContainerRecipe = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the container recipe.",
        "description_kind": "plain",
        "type": "string"
      },
      "components": {
        "computed": true,
        "description": "Components for build and test that are included in the container recipe.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "component_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the component.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "parameters": {
              "computed": true,
              "description": "A group of parameter settings that are used to configure the component for a specific recipe.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
                    "computed": true,
                    "description": "The name of the component parameter to set.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description": "Sets the value for the named component parameter.",
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
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "container_type": {
        "computed": true,
        "description": "Specifies the type of container, such as Docker.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the container recipe.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dockerfile_template_data": {
        "computed": true,
        "description": "Dockerfiles are text documents that are used to build Docker containers, and ensure that they contain all of the elements required by the application running inside. The template data consists of contextual variables where Image Builder places build information or scripts, based on your container image recipe.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dockerfile_template_uri": {
        "computed": true,
        "description": "The S3 URI for the Dockerfile that will be used to build your container image.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "image_os_version_override": {
        "computed": true,
        "description": "Specifies the operating system version for the source image.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_configuration": {
        "computed": true,
        "description": "A group of options that can be used to configure an instance for building and testing container images.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "block_device_mappings": {
              "computed": true,
              "description": "Defines the block devices to attach for building an instance from this Image Builder AMI.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "device_name": {
                    "computed": true,
                    "description": "The device to which these mappings apply.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "ebs": {
                    "computed": true,
                    "description": "Use to manage Amazon EBS-specific configuration for this mapping.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "delete_on_termination": {
                          "computed": true,
                          "description": "Use to configure delete on termination of the associated device.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "encrypted": {
                          "computed": true,
                          "description": "Use to configure device encryption.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "iops": {
                          "computed": true,
                          "description": "Use to configure device IOPS.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "kms_key_id": {
                          "computed": true,
                          "description": "Use to configure the KMS key to use when encrypting the device.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "snapshot_id": {
                          "computed": true,
                          "description": "The snapshot that defines the device contents.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "throughput": {
                          "computed": true,
                          "description": "For GP3 volumes only - The throughput in MiB/s that the volume supports.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "volume_size": {
                          "computed": true,
                          "description": "Use to override the device's volume size.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "volume_type": {
                          "computed": true,
                          "description": "Use to override the device's volume type.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "no_device": {
                    "computed": true,
                    "description": "Use to remove a mapping from the parent image.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "virtual_name": {
                    "computed": true,
                    "description": "Use to manage instance ephemeral devices.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "image": {
              "computed": true,
              "description": "The AMI ID to use as the base image for a container build and test instance. If not specified, Image Builder will use the appropriate ECS-optimized AMI as a base image.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "kms_key_id": {
        "computed": true,
        "description": "Identifies which KMS key is used to encrypt the container image.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the container recipe.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "parent_image": {
        "computed": true,
        "description": "The source image for the container recipe.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "platform_override": {
        "computed": true,
        "description": "Specifies the operating system platform when you use a custom source image.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags that are attached to the container recipe.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "target_repository": {
        "computed": true,
        "description": "The destination repository for the container image.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "repository_name": {
              "computed": true,
              "description": "The name of the container repository where the output container image is stored. This name is prefixed by the repository location.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "service": {
              "computed": true,
              "description": "Specifies the service in which this image was registered.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "version": {
        "computed": true,
        "description": "The semantic version of the container recipe (\u003cmajor\u003e.\u003cminor\u003e.\u003cpatch\u003e).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "working_directory": {
        "computed": true,
        "description": "The working directory to be used during build and test workflows.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource schema for AWS::ImageBuilder::ContainerRecipe",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccImagebuilderContainerRecipeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccImagebuilderContainerRecipe), &result)
	return &result
}
