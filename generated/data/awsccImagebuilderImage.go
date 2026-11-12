package data

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
        "type": "string"
      },
      "deletion_settings": {
        "computed": true,
        "description": "The deletion settings of the image, indicating whether to delete the underlying resources in addition to the image.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "execution_role": {
              "computed": true,
              "description": "The execution role to use for deleting the image, as well as underlying resources.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "distribution_configuration_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the distribution configuration.",
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
              "type": "string"
            },
            "on_update": {
              "computed": true,
              "description": "Whether to trigger the image pipeline when the pipeline is updated. False by default.",
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "image_recipe_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the image recipe that defines how images are configured, tested, and assessed.",
        "description_kind": "plain",
        "type": "string"
      },
      "image_scanning_configuration": {
        "computed": true,
        "description": "Contains settings for vulnerability scans that Amazon Inspector runs against the test instance during image creation.",
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
                    "description": "The name of the container repository that Amazon Inspector scans to identify findings for your container images. The name includes the path for the repository location. If you don’t provide this information, Image Builder creates a repository in your account named image-builder-image-scanning-repository to use for vulnerability scans for your output container images.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "image_scanning_enabled": {
              "computed": true,
              "description": "This sets whether Image Builder keeps a snapshot of the vulnerability scans that Amazon Inspector runs against the test instance when you create a new image.",
              "description_kind": "plain",
              "type": "bool"
            },
            "image_scanning_types": {
              "computed": true,
              "description": "The type of vulnerability scan that Image Builder runs for AMI output when ImageScanningEnabled is true. Specify a single value. If you omit ImageScanningTypes, Image Builder uses TEST_INSTANCE_CVE. With TEST_INSTANCE_CVE, Amazon Inspector scans the test instance during image tests, and you must activate Amazon Inspector for Amazon EC2 scanning in your account. With ON_DEMAND_AMI_CVE, Amazon Inspector starts an on-demand scan of the output AMI during image tests. You must build from an image recipe, but you do not need to activate Amazon Inspector in your account.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "scan_result_condition": {
              "computed": true,
              "description": "The condition that the output AMI's scan results must meet for the image to pass. You can set ScanResultCondition only when ImageScanningTypes is ON_DEMAND_AMI_CVE. If you omit it, Image Builder starts the scan but does not wait for it or use its results to pass or fail the image.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "fix_available": {
                    "computed": true,
                    "description": "Specifies whether only findings that have a fix available count toward MaximumSeverityLabel. When false, Image Builder counts every finding. The default is true.",
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "maximum_severity_label": {
                    "computed": true,
                    "description": "The maximum severity allowed for findings on the output AMI. A finding above this severity fails the image, unless FixAvailable excludes it.",
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
              "type": "bool"
            },
            "timeout_minutes": {
              "computed": true,
              "description": "TimeoutMinutes",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
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
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
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
    "description": "Data Source schema for AWS::ImageBuilder::Image",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccImagebuilderImageSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccImagebuilderImage), &result)
	return &result
}
