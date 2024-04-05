package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSyntheticsCanary = `{
  "block": {
    "attributes": {
      "artifact_config": {
        "computed": true,
        "description": "Provide artifact configuration",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "s3_encryption": {
              "computed": true,
              "description": "Encryption configuration for uploading artifacts to S3",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "encryption_mode": {
                    "computed": true,
                    "description": "Encryption mode for encrypting artifacts when uploading to S3. Valid values: SSE_S3 and SSE_KMS.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "kms_key_arn": {
                    "computed": true,
                    "description": "KMS key Arn for encrypting artifacts when uploading to S3. You must specify KMS key Arn for SSE_KMS encryption mode only.",
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
      "artifact_s3_location": {
        "computed": true,
        "description": "Provide the s3 bucket output location for test results",
        "description_kind": "plain",
        "type": "string"
      },
      "canary_id": {
        "computed": true,
        "description": "Id of the canary",
        "description_kind": "plain",
        "type": "string"
      },
      "code": {
        "computed": true,
        "description": "Provide the canary script source",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "handler": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "s3_bucket": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "s3_key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "s3_object_version": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "script": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "source_location_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "delete_lambda_resources_on_canary_deletion": {
        "computed": true,
        "description": "Deletes associated lambda resources created by Synthetics if set to True. Default is False",
        "description_kind": "plain",
        "type": "bool"
      },
      "execution_role_arn": {
        "computed": true,
        "description": "Lambda Execution role used to run your canaries",
        "description_kind": "plain",
        "type": "string"
      },
      "failure_retention_period": {
        "computed": true,
        "description": "Retention period of failed canary runs represented in number of days",
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Name of the canary.",
        "description_kind": "plain",
        "type": "string"
      },
      "run_config": {
        "computed": true,
        "description": "Provide canary run configuration",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "active_tracing": {
              "computed": true,
              "description": "Enable active tracing if set to true",
              "description_kind": "plain",
              "type": "bool"
            },
            "environment_variables": {
              "computed": true,
              "description": "Environment variable key-value pairs.",
              "description_kind": "plain",
              "type": [
                "map",
                "string"
              ]
            },
            "memory_in_mb": {
              "computed": true,
              "description": "Provide maximum memory available for canary in MB",
              "description_kind": "plain",
              "type": "number"
            },
            "timeout_in_seconds": {
              "computed": true,
              "description": "Provide maximum canary timeout per run in seconds",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "runtime_version": {
        "computed": true,
        "description": "Runtime version of Synthetics Library",
        "description_kind": "plain",
        "type": "string"
      },
      "schedule": {
        "computed": true,
        "description": "Frequency to run your canaries",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "duration_in_seconds": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "expression": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "start_canary_after_creation": {
        "computed": true,
        "description": "Runs canary if set to True. Default is False",
        "description_kind": "plain",
        "type": "bool"
      },
      "state": {
        "computed": true,
        "description": "State of the canary",
        "description_kind": "plain",
        "type": "string"
      },
      "success_retention_period": {
        "computed": true,
        "description": "Retention period of successful canary runs represented in number of days",
        "description_kind": "plain",
        "type": "number"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "visual_reference": {
        "computed": true,
        "description": "Visual reference configuration for visual testing",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "base_canary_run_id": {
              "computed": true,
              "description": "Canary run id to be used as base reference for visual testing",
              "description_kind": "plain",
              "type": "string"
            },
            "base_screenshots": {
              "computed": true,
              "description": "List of screenshots used as base reference for visual testing",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "ignore_coordinates": {
                    "computed": true,
                    "description": "List of coordinates of rectangles to be ignored during visual testing",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "screenshot_name": {
                    "computed": true,
                    "description": "Name of the screenshot to be used as base reference for visual testing",
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
      "vpc_config": {
        "computed": true,
        "description": "Provide VPC Configuration if enabled.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "security_group_ids": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "subnet_ids": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "vpc_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::Synthetics::Canary",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSyntheticsCanarySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSyntheticsCanary), &result)
	return &result
}
