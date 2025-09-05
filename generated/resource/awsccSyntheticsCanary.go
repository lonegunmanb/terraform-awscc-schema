package resource

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
                    "optional": true,
                    "type": "string"
                  },
                  "kms_key_arn": {
                    "computed": true,
                    "description": "KMS key Arn for encrypting artifacts when uploading to S3. You must specify KMS key Arn for SSE_KMS encryption mode only.",
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
      "artifact_s3_location": {
        "description": "Provide the s3 bucket output location for test results",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "browser_configs": {
        "computed": true,
        "description": "List of browser configurations for the canary",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "browser_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "canary_id": {
        "computed": true,
        "description": "Id of the canary",
        "description_kind": "plain",
        "type": "string"
      },
      "code": {
        "description": "Provide the canary script source",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "dependencies": {
              "computed": true,
              "description": "List of Lambda layers to attach to the canary",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "reference": {
                    "computed": true,
                    "description": "ARN of the Lambda layer",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "type": {
                    "computed": true,
                    "description": "Type of dependency",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "handler": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "s3_bucket": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "s3_key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "s3_object_version": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "script": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_location_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "delete_lambda_resources_on_canary_deletion": {
        "computed": true,
        "description": "Deletes associated lambda resources created by Synthetics if set to True. Default is False",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "dry_run_and_update": {
        "computed": true,
        "description": "Setting to control if UpdateCanary will perform a DryRun and validate it is PASSING before performing the Update. Default is FALSE.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "execution_role_arn": {
        "description": "Lambda Execution role used to run your canaries",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "failure_retention_period": {
        "computed": true,
        "description": "Retention period of failed canary runs represented in number of days",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "Name of the canary.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "provisioned_resource_cleanup": {
        "computed": true,
        "description": "Setting to control if provisioned resources created by Synthetics are deleted alongside the canary. Default is AUTOMATIC.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resources_to_replicate_tags": {
        "computed": true,
        "description": "List of resources which canary tags should be replicated to.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
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
              "optional": true,
              "type": "bool"
            },
            "environment_variables": {
              "computed": true,
              "description": "Environment variable key-value pairs.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "ephemeral_storage": {
              "computed": true,
              "description": "Provide ephemeralStorage available for canary in MB",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "memory_in_mb": {
              "computed": true,
              "description": "Provide maximum memory available for canary in MB",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "timeout_in_seconds": {
              "computed": true,
              "description": "Provide maximum canary timeout per run in seconds",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "runtime_version": {
        "description": "Runtime version of Synthetics Library",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "schedule": {
        "description": "Frequency to run your canaries",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "duration_in_seconds": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "expression": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "retry_config": {
              "computed": true,
              "description": "Provide canary auto retry configuration",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "max_retries": {
                    "computed": true,
                    "description": "maximum times the canary will be retried upon the scheduled run failure",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "start_canary_after_creation": {
        "computed": true,
        "description": "Runs canary if set to True. Default is False",
        "description_kind": "plain",
        "optional": true,
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
        "optional": true,
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
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
              "optional": true,
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
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "screenshot_name": {
                    "computed": true,
                    "description": "Name of the screenshot to be used as base reference for visual testing",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "browser_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "visual_references": {
        "computed": true,
        "description": "List of visual references for the canary",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "base_canary_run_id": {
              "computed": true,
              "description": "Canary run id to be used as base reference for visual testing",
              "description_kind": "plain",
              "optional": true,
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
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "screenshot_name": {
                    "computed": true,
                    "description": "Name of the screenshot to be used as base reference for visual testing",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "browser_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "vpc_config": {
        "computed": true,
        "description": "Provide VPC Configuration if enabled.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ipv_6_allowed_for_dual_stack": {
              "computed": true,
              "description": "Allow outbound IPv6 traffic on VPC canaries that are connected to dual-stack subnets if set to true",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "security_group_ids": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "subnet_ids": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "vpc_id": {
              "computed": true,
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
    "description": "Resource Type definition for AWS::Synthetics::Canary",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSyntheticsCanarySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSyntheticsCanary), &result)
	return &result
}
