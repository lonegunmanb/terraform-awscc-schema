package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccImagebuilderImage = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the image.",
        "description_kind": "plain",
        "type": "string"
      },
      "container_recipe_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the container recipe that defines how images are configured and tested.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "distribution_configuration_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the distribution configuration.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enhanced_image_metadata_enabled": {
        "computed": true,
        "description": "Collects additional information about the image being created, including the operating system (OS) version and package list.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "execution_role": {
        "computed": true,
        "description": "The execution role name/ARN for the image build, if provided",
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
      "image_id": {
        "computed": true,
        "description": "The AMI ID of the EC2 AMI in current region.",
        "description_kind": "plain",
        "type": "string"
      },
      "image_pipeline_execution_settings": {
        "computed": true,
        "description": "The image pipeline execution settings of the image.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "deployment_id": {
              "computed": true,
              "description": "The deployment ID of the pipeline, used to trigger new image pipeline executions.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "on_update": {
              "computed": true,
              "description": "Whether to trigger the image pipeline when the pipeline is updated. False by default.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "image_recipe_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the image recipe that defines how images are configured, tested, and assessed.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "image_scanning_configuration": {
        "computed": true,
        "description": "Contains settings for vulnerability scans.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ecr_configuration": {
              "computed": true,
              "description": "Contains ECR settings for vulnerability scans.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "container_tags": {
                    "computed": true,
                    "description": "Tags for Image Builder to apply the output container image that is scanned. Tags can help you identify and manage your scanned images.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "repository_name": {
                    "computed": true,
                    "description": "The name of the container repository that Amazon Inspector scans to identify findings for your container images. The name includes the path for the repository location. If you don’t provide this information, Image Builder creates a repository in your account named image-builder-image-scanning-repository to use for vulnerability scans for your output container images.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "image_scanning_enabled": {
              "computed": true,
              "description": "This sets whether Image Builder keeps a snapshot of the vulnerability scans that Amazon Inspector runs against the build instance when you create a new image.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "image_tests_configuration": {
        "computed": true,
        "description": "The image tests configuration used when creating this image.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "image_tests_enabled": {
              "computed": true,
              "description": "ImageTestsEnabled",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "timeout_minutes": {
              "computed": true,
              "description": "TimeoutMinutes",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "image_uri": {
        "computed": true,
        "description": "URI for containers created in current Region with default ECR image tag",
        "description_kind": "plain",
        "type": "string"
      },
      "infrastructure_configuration_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the infrastructure configuration.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "latest_version": {
        "computed": true,
        "description": "The latest version references of the image.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "arn": {
              "computed": true,
              "description": "The latest version ARN of the created image.",
              "description_kind": "plain",
              "type": "string"
            },
            "major": {
              "computed": true,
              "description": "The latest version ARN of the created image, with the same major version.",
              "description_kind": "plain",
              "type": "string"
            },
            "minor": {
              "computed": true,
              "description": "The latest version ARN of the created image, with the same minor version.",
              "description_kind": "plain",
              "type": "string"
            },
            "patch": {
              "computed": true,
              "description": "The latest version ARN of the created image, with the same patch version.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "logging_configuration": {
        "computed": true,
        "description": "The logging configuration settings for the image.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "log_group_name": {
              "computed": true,
              "description": "The name of the log group for image build logs.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "name": {
        "computed": true,
        "description": "The name of the image.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags associated with the image.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "workflows": {
        "computed": true,
        "description": "Workflows to define the image build process",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "on_failure": {
              "computed": true,
              "description": "Define execution decision in case of workflow failure",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "parallel_group": {
              "computed": true,
              "description": "The parallel group name",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "parameters": {
              "computed": true,
              "description": "The parameters associated with the workflow",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
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
            "workflow_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the workflow",
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
    "description": "Resource schema for AWS::ImageBuilder::Image",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccImagebuilderImageSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccImagebuilderImage), &result)
	return &result
}
