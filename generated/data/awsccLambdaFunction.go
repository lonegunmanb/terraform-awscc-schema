package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLambdaFunction = `{
  "block": {
    "attributes": {
      "architectures": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "arn": {
        "computed": true,
        "description": "Unique identifier for function resources",
        "description_kind": "plain",
        "type": "string"
      },
      "code": {
        "computed": true,
        "description": "The code for the function.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "image_uri": {
              "computed": true,
              "description": "ImageUri.",
              "description_kind": "plain",
              "type": "string"
            },
            "s3_bucket": {
              "computed": true,
              "description": "An Amazon S3 bucket in the same AWS Region as your function. The bucket can be in a different AWS account.",
              "description_kind": "plain",
              "type": "string"
            },
            "s3_key": {
              "computed": true,
              "description": "The Amazon S3 key of the deployment package.",
              "description_kind": "plain",
              "type": "string"
            },
            "s3_object_version": {
              "computed": true,
              "description": "For versioned objects, the version of the deployment package object to use.",
              "description_kind": "plain",
              "type": "string"
            },
            "zip_file": {
              "computed": true,
              "description": "The source code of your Lambda function. If you include your function source inline with this parameter, AWS CloudFormation places it in a file named index and zips it to create a deployment package..",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "code_signing_config_arn": {
        "computed": true,
        "description": "A unique Arn for CodeSigningConfig resource",
        "description_kind": "plain",
        "type": "string"
      },
      "dead_letter_config": {
        "computed": true,
        "description": "A dead letter queue configuration that specifies the queue or topic where Lambda sends asynchronous events when they fail processing.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "target_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of an Amazon SQS queue or Amazon SNS topic.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "description": {
        "computed": true,
        "description": "A description of the function.",
        "description_kind": "plain",
        "type": "string"
      },
      "environment": {
        "computed": true,
        "description": "Environment variables that are accessible from function code during execution.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "variables": {
              "computed": true,
              "description": "Environment variable key-value pairs.",
              "description_kind": "plain",
              "type": [
                "map",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        }
      },
      "ephemeral_storage": {
        "computed": true,
        "description": "A function's ephemeral storage settings.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "size": {
              "computed": true,
              "description": "The amount of ephemeral storage that your function has access to.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "file_system_configs": {
        "computed": true,
        "description": "Connection settings for an Amazon EFS file system. To connect a function to a file system, a mount target must be available in every Availability Zone that your function connects to. If your template contains an AWS::EFS::MountTarget resource, you must also specify a DependsOn attribute to ensure that the mount target is created or updated before the function.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the Amazon EFS access point that provides access to the file system.",
              "description_kind": "plain",
              "type": "string"
            },
            "local_mount_path": {
              "computed": true,
              "description": "The path where the function can access the file system, starting with /mnt/.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "function_name": {
        "computed": true,
        "description": "The name of the Lambda function, up to 64 characters in length. If you don't specify a name, AWS CloudFormation generates one.",
        "description_kind": "plain",
        "type": "string"
      },
      "handler": {
        "computed": true,
        "description": "The name of the method within your code that Lambda calls to execute your function. The format includes the file name. It can also include namespaces and other qualifiers, depending on the runtime",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "image_config": {
        "computed": true,
        "description": "ImageConfig",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "command": {
              "computed": true,
              "description": "Command.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "entry_point": {
              "computed": true,
              "description": "EntryPoint.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "working_directory": {
              "computed": true,
              "description": "WorkingDirectory.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "kms_key_arn": {
        "computed": true,
        "description": "The ARN of the AWS Key Management Service (AWS KMS) key that's used to encrypt your function's environment variables. If it's not provided, AWS Lambda uses a default service key.",
        "description_kind": "plain",
        "type": "string"
      },
      "layers": {
        "computed": true,
        "description": "A list of function layers to add to the function's execution environment. Specify each layer by its ARN, including the version.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "logging_config": {
        "computed": true,
        "description": "The logging configuration of your function",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "application_log_level": {
              "computed": true,
              "description": "Application log granularity level, can only be used when LogFormat is set to JSON",
              "description_kind": "plain",
              "type": "string"
            },
            "log_format": {
              "computed": true,
              "description": "Log delivery format for the lambda function",
              "description_kind": "plain",
              "type": "string"
            },
            "log_group": {
              "computed": true,
              "description": "The log group name.",
              "description_kind": "plain",
              "type": "string"
            },
            "system_log_level": {
              "computed": true,
              "description": "System log granularity level, can only be used when LogFormat is set to JSON",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "memory_size": {
        "computed": true,
        "description": "The amount of memory that your function has access to. Increasing the function's memory also increases its CPU allocation. The default value is 128 MB. The value must be a multiple of 64 MB.",
        "description_kind": "plain",
        "type": "number"
      },
      "package_type": {
        "computed": true,
        "description": "PackageType.",
        "description_kind": "plain",
        "type": "string"
      },
      "reserved_concurrent_executions": {
        "computed": true,
        "description": "The number of simultaneous executions to reserve for the function.",
        "description_kind": "plain",
        "type": "number"
      },
      "role": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the function's execution role.",
        "description_kind": "plain",
        "type": "string"
      },
      "runtime": {
        "computed": true,
        "description": "The identifier of the function's runtime.",
        "description_kind": "plain",
        "type": "string"
      },
      "runtime_management_config": {
        "computed": true,
        "description": "RuntimeManagementConfig",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "runtime_version_arn": {
              "computed": true,
              "description": "Unique identifier for a runtime version arn",
              "description_kind": "plain",
              "type": "string"
            },
            "update_runtime_on": {
              "computed": true,
              "description": "Trigger for runtime update",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "snap_start": {
        "computed": true,
        "description": "The SnapStart setting of your function",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "apply_on": {
              "computed": true,
              "description": "Applying SnapStart setting on function resource type.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "snap_start_response": {
        "computed": true,
        "description": "The SnapStart response of your function",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "apply_on": {
              "computed": true,
              "description": "Applying SnapStart setting on function resource type.",
              "description_kind": "plain",
              "type": "string"
            },
            "optimization_status": {
              "computed": true,
              "description": "Indicates whether SnapStart is activated for the specified function version.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description": "A list of tags to apply to the function.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "timeout": {
        "computed": true,
        "description": "The amount of time that Lambda allows a function to run before stopping it. The default is 3 seconds. The maximum allowed value is 900 seconds.",
        "description_kind": "plain",
        "type": "number"
      },
      "tracing_config": {
        "computed": true,
        "description": "Set Mode to Active to sample and trace a subset of incoming requests with AWS X-Ray.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "mode": {
              "computed": true,
              "description": "The tracing mode.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "vpc_config": {
        "computed": true,
        "description": "For network connectivity to AWS resources in a VPC, specify a list of security groups and subnets in the VPC.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ipv_6_allowed_for_dual_stack": {
              "computed": true,
              "description": "A boolean indicating whether IPv6 protocols will be allowed for dual stack subnets",
              "description_kind": "plain",
              "type": "bool"
            },
            "security_group_ids": {
              "computed": true,
              "description": "A list of VPC security groups IDs.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "subnet_ids": {
              "computed": true,
              "description": "A list of VPC subnet IDs.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::Lambda::Function",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLambdaFunctionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLambdaFunction), &result)
	return &result
}
