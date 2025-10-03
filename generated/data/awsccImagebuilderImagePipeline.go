package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccImagebuilderImagePipeline = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the image pipeline.",
        "description_kind": "plain",
        "type": "string"
      },
      "container_recipe_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the container recipe that defines how images are configured and tested.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the image pipeline.",
        "description_kind": "plain",
        "type": "string"
      },
      "distribution_configuration_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the distribution configuration associated with this image pipeline.",
        "description_kind": "plain",
        "type": "string"
      },
      "enhanced_image_metadata_enabled": {
        "computed": true,
        "description": "Collects additional information about the image being created, including the operating system (OS) version and package list.",
        "description_kind": "plain",
        "type": "bool"
      },
      "execution_role": {
        "computed": true,
        "description": "The execution role name/ARN for the image build, if provided",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "image_recipe_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the image recipe that defines how images are configured, tested, and assessed.",
        "description_kind": "plain",
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
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "repository_name": {
                    "computed": true,
                    "description": "The name of the container repository that Amazon Inspector scans to identify findings for your container images. The name includes the path for the repository location. If you don't provide this information, Image Builder creates a repository in your account named image-builder-image-scanning-repository to use for vulnerability scans for your output container images.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "image_scanning_enabled": {
              "computed": true,
              "description": "This sets whether Image Builder keeps a snapshot of the vulnerability scans that Amazon Inspector runs against the build instance when you create a new image.",
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "image_tests_configuration": {
        "computed": true,
        "description": "The image tests configuration of the image pipeline.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "image_tests_enabled": {
              "computed": true,
              "description": "Defines if tests should be executed when building this image.",
              "description_kind": "plain",
              "type": "bool"
            },
            "timeout_minutes": {
              "computed": true,
              "description": "The maximum time in minutes that tests are permitted to run.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "infrastructure_configuration_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the infrastructure configuration associated with this image pipeline.",
        "description_kind": "plain",
        "type": "string"
      },
      "logging_configuration": {
        "computed": true,
        "description": "The logging configuration settings for the image pipeline.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "image_log_group_name": {
              "computed": true,
              "description": "The name of the log group for image build logs.",
              "description_kind": "plain",
              "type": "string"
            },
            "pipeline_log_group_name": {
              "computed": true,
              "description": "The name of the log group for pipeline execution logs.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "name": {
        "computed": true,
        "description": "The name of the image pipeline.",
        "description_kind": "plain",
        "type": "string"
      },
      "schedule": {
        "computed": true,
        "description": "The schedule of the image pipeline.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "auto_disable_policy": {
              "computed": true,
              "description": "The auto-disable policy for the image pipeline.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "failure_count": {
                    "computed": true,
                    "description": "The number of consecutive failures after which the pipeline should be automatically disabled.",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "pipeline_execution_start_condition": {
              "computed": true,
              "description": "The condition configures when the pipeline should trigger a new image build.",
              "description_kind": "plain",
              "type": "string"
            },
            "schedule_expression": {
              "computed": true,
              "description": "The expression determines how often EC2 Image Builder evaluates your pipelineExecutionStartCondition.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "status": {
        "computed": true,
        "description": "The status of the image pipeline.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags of this image pipeline.",
        "description_kind": "plain",
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
              "type": "string"
            },
            "parallel_group": {
              "computed": true,
              "description": "The parallel group name",
              "description_kind": "plain",
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
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "list"
              }
            },
            "workflow_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the workflow",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::ImageBuilder::ImagePipeline",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccImagebuilderImagePipelineSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccImagebuilderImagePipeline), &result)
	return &result
}
