package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLambdaWebFunctionRevision = `{
  "block": {
    "attributes": {
      "build_config": {
        "description": "The build configuration for the revision.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "code_config": {
              "description": "The code configuration for the revision.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "s3_object": {
                    "description": "The Amazon S3 location of the deployment artifact.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bucket": {
                          "description": "The S3 bucket name.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "key": {
                          "description": "The S3 object key.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "version_id": {
                          "computed": true,
                          "description": "The S3 object version ID.",
                          "description_kind": "plain",
                          "optional": true,
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
            "runtime_config": {
              "description": "The runtime configuration for the revision.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "runtime": {
                    "description": "The runtime identifier.",
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
      "created_at": {
        "computed": true,
        "description": "The timestamp when the revision was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the revision.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "function_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the function.",
        "description_kind": "plain",
        "type": "string"
      },
      "function_name": {
        "description": "The name of the web function this revision belongs to. The length constraint applies only to the full ARN. If you specify only the function name, it is limited to 64 characters in length.",
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
      "kms_key_arn": {
        "computed": true,
        "description": "The ARN of the KMS key used to encrypt the revision.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "revision_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the revision.",
        "description_kind": "plain",
        "type": "string"
      },
      "revision_id": {
        "computed": true,
        "description": "The unique identifier of the revision.",
        "description_kind": "plain",
        "type": "string"
      },
      "service_config": {
        "description": "The service configuration for the revision.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "environment_variables": {
              "computed": true,
              "description": "Environment variables for the function.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "execution_role_arn": {
              "description": "The ARN of the execution role.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "max_concurrency_per_environment": {
              "computed": true,
              "description": "The maximum concurrency per environment.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "telemetry_config": {
              "computed": true,
              "description": "The telemetry configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "logging_config": {
                    "computed": true,
                    "description": "The logging configuration for the web function.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "application_log_level": {
                          "computed": true,
                          "description": "The application log level.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "log_group": {
                          "computed": true,
                          "description": "The CloudWatch log group name.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "system_log_level": {
                          "computed": true,
                          "description": "The system log level.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "timeout_seconds": {
              "computed": true,
              "description": "The function timeout in seconds.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "state": {
        "computed": true,
        "description": "The current state of the revision.",
        "description_kind": "plain",
        "type": "string"
      },
      "state_reason": {
        "computed": true,
        "description": "The reason for the revision's current state.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Lambda::WebFunctionRevision. An immutable revision of a Lambda web function, containing the function code and configuration.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLambdaWebFunctionRevisionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLambdaWebFunctionRevision), &result)
	return &result
}
