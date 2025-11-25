package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccImagebuilderImageRecipe = `{
  "block": {
    "attributes": {
      "additional_instance_configuration": {
        "computed": true,
        "description": "Specify additional settings and launch scripts for your build instances.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "systems_manager_agent": {
              "computed": true,
              "description": "Contains settings for the SSM agent on your build instance.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "uninstall_after_build": {
                    "computed": true,
                    "description": "Controls whether the SSM agent is removed from your final build image, prior to creating the new AMI. If this is set to true, then the agent is removed from the final image. If it's set to false, then the agent is left in, so that it is included in the new AMI. The default value is false.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "user_data_override": {
              "computed": true,
              "description": "Use this property to provide commands or a command script to run when you launch your build instance.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "ami_tags": {
        "computed": true,
        "description": "The tags to apply to the AMI created by this image recipe.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the image recipe.",
        "description_kind": "plain",
        "type": "string"
      },
      "block_device_mappings": {
        "computed": true,
        "description": "The block device mappings to apply when creating images from this recipe.",
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
      "components": {
        "computed": true,
        "description": "The components of the image recipe.",
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
      "description": {
        "computed": true,
        "description": "The description of the image recipe.",
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
      "latest_version": {
        "computed": true,
        "description": "The latest version references of the image recipe.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "arn": {
              "computed": true,
              "description": "The latest version ARN of the created image recipe.",
              "description_kind": "plain",
              "type": "string"
            },
            "major": {
              "computed": true,
              "description": "The latest version ARN of the created image recipe, with the same major version.",
              "description_kind": "plain",
              "type": "string"
            },
            "minor": {
              "computed": true,
              "description": "The latest version ARN of the created image recipe, with the same minor version.",
              "description_kind": "plain",
              "type": "string"
            },
            "patch": {
              "computed": true,
              "description": "The latest version ARN of the created image recipe, with the same patch version.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "name": {
        "description": "The name of the image recipe.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parent_image": {
        "description": "The parent image of the image recipe.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags of the image recipe.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "version": {
        "description": "The version of the image recipe.",
        "description_kind": "plain",
        "required": true,
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
    "description": "Resource schema for AWS::ImageBuilder::ImageRecipe",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccImagebuilderImageRecipeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccImagebuilderImageRecipe), &result)
	return &result
}
