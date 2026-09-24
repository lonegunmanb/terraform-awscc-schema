package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccElasticbeanstalkApplicationVersion = `{
  "block": {
    "attributes": {
      "application_name": {
        "computed": true,
        "description": "The name of the Elastic Beanstalk application that is associated with this application version. ",
        "description_kind": "plain",
        "type": "string"
      },
      "application_version_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "build_configuration": {
        "computed": true,
        "description": "Settings for an AWS CodeBuild build that packages and builds an application version from source code.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "artifact_name": {
              "computed": true,
              "description": "The name of the build artifact.",
              "description_kind": "plain",
              "type": "string"
            },
            "code_build_service_role": {
              "computed": true,
              "description": "The ARN of the IAM role that AWS CodeBuild assumes to build the application version.",
              "description_kind": "plain",
              "type": "string"
            },
            "compute_type": {
              "computed": true,
              "description": "The compute type for the CodeBuild build environment.",
              "description_kind": "plain",
              "type": "string"
            },
            "image": {
              "computed": true,
              "description": "The CodeBuild image used for the build environment.",
              "description_kind": "plain",
              "type": "string"
            },
            "timeout_in_minutes": {
              "computed": true,
              "description": "The timeout for the CodeBuild build, in minutes.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "description": {
        "computed": true,
        "description": "A description of this application version.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "image_configuration": {
        "computed": true,
        "description": "Configuration for image-based application versions.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "build": {
              "computed": true,
              "description": "Configuration for building a container image from source code.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "architecture": {
                    "computed": true,
                    "description": "The target architecture for the built container image.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "buildpack": {
                    "computed": true,
                    "description": "The buildpack to use for building the image.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "code_build_service_role": {
                    "computed": true,
                    "description": "The ARN of the IAM role that AWS CodeBuild assumes to build the application version.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "compute_type": {
                    "computed": true,
                    "description": "The compute type for the CodeBuild build environment.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "dockerfile_location": {
                    "computed": true,
                    "description": "The path to the Dockerfile, relative to the source root.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "timeout_in_minutes": {
                    "computed": true,
                    "description": "The timeout for the CodeBuild build, in minutes.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "type": {
                    "computed": true,
                    "description": "The type of image build: docker or buildpack.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "source": {
              "computed": true,
              "description": "The container image source for this version, as an ECR image URI.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "uri": {
                    "computed": true,
                    "description": "The URI of the container image, e.g. an ECR image URI.",
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
      "process": {
        "computed": true,
        "description": "Pre-process and validate the environment manifest (` + "`" + `env.yaml` + "`" + `) and configuration files in the source bundle. Leave unset for the service default.",
        "description_kind": "plain",
        "type": "bool"
      },
      "source_bundle": {
        "computed": true,
        "description": "The Amazon S3 bucket and key that identify the location of the source bundle for this version. ",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "s3_bucket": {
              "computed": true,
              "description": "The Amazon S3 bucket where the data is located.",
              "description_kind": "plain",
              "type": "string"
            },
            "s3_key": {
              "computed": true,
              "description": "The Amazon S3 key where the data is located.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::ElasticBeanstalk::ApplicationVersion",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccElasticbeanstalkApplicationVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccElasticbeanstalkApplicationVersion), &result)
	return &result
}
