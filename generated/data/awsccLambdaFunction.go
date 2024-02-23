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
        "description": "The instruction set architecture that the function supports. Enter a string array with one of the valid values (arm64 or x86_64). The default value is ` + "`" + `` + "`" + `x86_64` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "arn": {
        "computed": true,
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
              "description": "URI of a [container image](https://docs.aws.amazon.com/lambda/latest/dg/lambda-images.html) in the Amazon ECR registry.",
              "description_kind": "plain",
              "type": "string"
            },
            "s3_bucket": {
              "computed": true,
              "description": "An Amazon S3 bucket in the same AWS-Region as your function. The bucket can be in a different AWS-account.",
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
              "description": "(Node.js and Python) The source code of your Lambda function. If you include your function source inline with this parameter, CFN places it in a file named ` + "`" + `` + "`" + `index` + "`" + `` + "`" + ` and zips it to create a [deployment package](https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-package.html). This zip file cannot exceed 4MB. For the ` + "`" + `` + "`" + `Handler` + "`" + `` + "`" + ` property, the first part of the handler identifier must be ` + "`" + `` + "`" + `index` + "`" + `` + "`" + `. For example, ` + "`" + `` + "`" + `index.handler` + "`" + `` + "`" + `.\n  For JSON, you must escape quotes and special characters such as newline (` + "`" + `` + "`" + `\\n` + "`" + `` + "`" + `) with a backslash.\n If you specify a function that interacts with an AWS CloudFormation custom resource, you don't have to write your own functions to send responses to the custom resource that invoked the function. AWS CloudFormation provides a response module ([cfn-response](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-lambda-function-code-cfnresponsemodule.html)) that simplifies sending responses. See [Using Lambda with CloudFormation](https://docs",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "code_signing_config_arn": {
        "computed": true,
        "description": "To enable code signing for this function, specify the ARN of a code-signing configuration. A code-signing configuration includes a set of signing profiles, which define the trusted publishers for this function.",
        "description_kind": "plain",
        "type": "string"
      },
      "dead_letter_config": {
        "computed": true,
        "description": "A dead-letter queue configuration that specifies the queue or topic where Lambda sends asynchronous events when they fail processing. For more information, see [Dead-letter queues](https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#invocation-dlq).",
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
              "description": "Environment variable key-value pairs. For more information, see [Using Lambda environment variables](https://docs.aws.amazon.com/lambda/latest/dg/configuration-envvars.html).",
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
        "description": "The size of the function's ` + "`" + `` + "`" + `/tmp` + "`" + `` + "`" + ` directory in MB. The default value is 512, but it can be any whole number between 512 and 10,240 MB.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "size": {
              "computed": true,
              "description": "The size of the function's ` + "`" + `` + "`" + `/tmp` + "`" + `` + "`" + ` directory.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "file_system_configs": {
        "computed": true,
        "description": "Connection settings for an Amazon EFS file system. To connect a function to a file system, a mount target must be available in every Availability Zone that your function connects to. If your template contains an [AWS::EFS::MountTarget](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html) resource, you must also specify a ` + "`" + `` + "`" + `DependsOn` + "`" + `` + "`" + ` attribute to ensure that the mount target is created or updated before the function.\n For more information about using the ` + "`" + `` + "`" + `DependsOn` + "`" + `` + "`" + ` attribute, see [DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html).",
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
              "description": "The path where the function can access the file system, starting with ` + "`" + `` + "`" + `/mnt/` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "function_name": {
        "computed": true,
        "description": "The name of the Lambda function, up to 64 characters in length. If you don't specify a name, CFN generates one.\n If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.",
        "description_kind": "plain",
        "type": "string"
      },
      "handler": {
        "computed": true,
        "description": "The name of the method within your code that Lambda calls to run your function. Handler is required if the deployment package is a .zip file archive. The format includes the file name. It can also include namespaces and other qualifiers, depending on the runtime. For more information, see [Lambda programming model](https://docs.aws.amazon.com/lambda/latest/dg/foundation-progmodel.html).",
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
        "description": "Configuration values that override the container image Dockerfile settings. For more information, see [Container image settings](https://docs.aws.amazon.com/lambda/latest/dg/images-create.html#images-parms).",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "command": {
              "computed": true,
              "description": "Specifies parameters that you want to pass in with ENTRYPOINT. You can specify a maximum of 1,500 parameters in the list.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "entry_point": {
              "computed": true,
              "description": "Specifies the entry point to their application, which is typically the location of the runtime executable. You can specify a maximum of 1,500 string entries in the list.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "working_directory": {
              "computed": true,
              "description": "Specifies the working directory. The length of the directory string cannot exceed 1,000 characters.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "kms_key_arn": {
        "computed": true,
        "description": "The ARN of the KMSlong (KMS) customer managed key that's used to encrypt your function's [environment variables](https://docs.aws.amazon.com/lambda/latest/dg/configuration-envvars.html#configuration-envvars-encryption). When [Lambda SnapStart](https://docs.aws.amazon.com/lambda/latest/dg/snapstart-security.html) is activated, Lambda also uses this key is to encrypt your function's snapshot. If you deploy your function using a container image, Lambda also uses this key to encrypt your function when it's deployed. Note that this is not the same key that's used to protect your container image in the Amazon Elastic Container Registry (Amazon ECR). If you don't provide a customer managed key, Lambda uses a default service key.",
        "description_kind": "plain",
        "type": "string"
      },
      "layers": {
        "computed": true,
        "description": "A list of [function layers](https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html) to add to the function's execution environment. Specify each layer by its ARN, including the version.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "logging_config": {
        "computed": true,
        "description": "The function's Amazon CloudWatch Logs configuration settings.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "application_log_level": {
              "computed": true,
              "description": "Set this property to filter the application logs for your function that Lambda sends to CloudWatch. Lambda only sends application logs at the selected level of detail and lower, where ` + "`" + `` + "`" + `TRACE` + "`" + `` + "`" + ` is the highest level and ` + "`" + `` + "`" + `FATAL` + "`" + `` + "`" + ` is the lowest.",
              "description_kind": "plain",
              "type": "string"
            },
            "log_format": {
              "computed": true,
              "description": "The format in which Lambda sends your function's application and system logs to CloudWatch. Select between plain text and structured JSON.",
              "description_kind": "plain",
              "type": "string"
            },
            "log_group": {
              "computed": true,
              "description": "The name of the Amazon CloudWatch log group the function sends logs to. By default, Lambda functions send logs to a default log group named ` + "`" + `` + "`" + `/aws/lambda/\u003cfunction name\u003e` + "`" + `` + "`" + `. To use a different log group, enter an existing log group or enter a new log group name.",
              "description_kind": "plain",
              "type": "string"
            },
            "system_log_level": {
              "computed": true,
              "description": "Set this property to filter the system logs for your function that Lambda sends to CloudWatch. Lambda only sends system logs at the selected level of detail and lower, where ` + "`" + `` + "`" + `DEBUG` + "`" + `` + "`" + ` is the highest level and ` + "`" + `` + "`" + `WARN` + "`" + `` + "`" + ` is the lowest.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "memory_size": {
        "computed": true,
        "description": "The amount of [memory available to the function](https://docs.aws.amazon.com/lambda/latest/dg/configuration-function-common.html#configuration-memory-console) at runtime. Increasing the function memory also increases its CPU allocation. The default value is 128 MB. The value can be any multiple of 1 MB. Note that new AWS accounts have reduced concurrency and memory quotas. AWS raises these quotas automatically based on your usage. You can also request a quota increase.",
        "description_kind": "plain",
        "type": "number"
      },
      "package_type": {
        "computed": true,
        "description": "The type of deployment package. Set to ` + "`" + `` + "`" + `Image` + "`" + `` + "`" + ` for container image and set ` + "`" + `` + "`" + `Zip` + "`" + `` + "`" + ` for .zip file archive.",
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
        "description": "The identifier of the function's [runtime](https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html). Runtime is required if the deployment package is a .zip file archive.\n The following list includes deprecated runtimes. For more information, see [Runtime deprecation policy](https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html#runtime-support-policy).",
        "description_kind": "plain",
        "type": "string"
      },
      "runtime_management_config": {
        "computed": true,
        "description": "Sets the runtime management configuration for a function's version. For more information, see [Runtime updates](https://docs.aws.amazon.com/lambda/latest/dg/runtimes-update.html).",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "runtime_version_arn": {
              "computed": true,
              "description": "The ARN of the runtime version you want the function to use.\n  This is only required if you're using the *Manual* runtime update mode.",
              "description_kind": "plain",
              "type": "string"
            },
            "update_runtime_on": {
              "computed": true,
              "description": "Specify the runtime update mode.\n  + *Auto (default)* - Automatically update to the most recent and secure runtime version using a [Two-phase runtime version rollout](https://docs.aws.amazon.com/lambda/latest/dg/runtimes-update.html#runtime-management-two-phase). This is the best choice for most customers to ensure they always benefit from runtime updates.\n + *FunctionUpdate* - LAM updates the runtime of you function to the most recent and secure runtime version when you update your function. This approach synchronizes runtime updates with function deployments, giving you control over when runtime updates are applied and allowing you to detect and mitigate rare runtime update incompatibilities early. When using this setting, you need to regularly update your functions to keep their runtime up-to-date.\n + *Manual* - You specify a runtime version in your function configuration. The function will use this runtime version indefinitely. In the rare case where a new runtime version is incomp",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "snap_start": {
        "computed": true,
        "description": "The function's [SnapStart](https://docs.aws.amazon.com/lambda/latest/dg/snapstart.html) setting.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "apply_on": {
              "computed": true,
              "description": "Set ` + "`" + `` + "`" + `ApplyOn` + "`" + `` + "`" + ` to ` + "`" + `` + "`" + `PublishedVersions` + "`" + `` + "`" + ` to create a snapshot of the initialized execution environment when you publish a function version.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "snap_start_response": {
        "computed": true,
        "description": "The function's [SnapStart](https://docs.aws.amazon.com/lambda/latest/dg/snapstart.html) setting.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "apply_on": {
              "computed": true,
              "description": "When set to ` + "`" + `` + "`" + `PublishedVersions` + "`" + `` + "`" + `, Lambda creates a snapshot of the execution environment when you publish a function version.",
              "description_kind": "plain",
              "type": "string"
            },
            "optimization_status": {
              "computed": true,
              "description": "When you provide a [qualified Amazon Resource Name (ARN)](https://docs.aws.amazon.com/lambda/latest/dg/configuration-versions.html#versioning-versions-using), this response element indicates whether SnapStart is activated for the specified function version.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description": "A list of [tags](https://docs.aws.amazon.com/lambda/latest/dg/tagging.html) to apply to the function.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "timeout": {
        "computed": true,
        "description": "The amount of time (in seconds) that Lambda allows a function to run before stopping it. The default is 3 seconds. The maximum allowed value is 900 seconds. For more information, see [Lambda execution environment](https://docs.aws.amazon.com/lambda/latest/dg/runtimes-context.html).",
        "description_kind": "plain",
        "type": "number"
      },
      "tracing_config": {
        "computed": true,
        "description": "Set ` + "`" + `` + "`" + `Mode` + "`" + `` + "`" + ` to ` + "`" + `` + "`" + `Active` + "`" + `` + "`" + ` to sample and trace a subset of incoming requests with [X-Ray](https://docs.aws.amazon.com/lambda/latest/dg/services-xray.html).",
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
        "description": "For network connectivity to AWS resources in a VPC, specify a list of security groups and subnets in the VPC. When you connect a function to a VPC, it can access resources and the internet only through that VPC. For more information, see [Configuring a Lambda function to access resources in a VPC](https://docs.aws.amazon.com/lambda/latest/dg/configuration-vpc.html).",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ipv_6_allowed_for_dual_stack": {
              "computed": true,
              "description": "Allows outbound IPv6 traffic on VPC functions that are connected to dual-stack subnets.",
              "description_kind": "plain",
              "type": "bool"
            },
            "security_group_ids": {
              "computed": true,
              "description": "A list of VPC security group IDs.",
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
