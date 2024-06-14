package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEcsCluster = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "capacity_providers": {
        "computed": true,
        "description": "The short name of one or more capacity providers to associate with the cluster. A capacity provider must be associated with a cluster before it can be included as part of the default capacity provider strategy of the cluster or used in a capacity provider strategy when calling the [CreateService](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_CreateService.html) or [RunTask](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_RunTask.html) actions.\n If specifying a capacity provider that uses an Auto Scaling group, the capacity provider must be created but not associated with another cluster. New Auto Scaling group capacity providers can be created with the [CreateCapacityProvider](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_CreateCapacityProvider.html) API operation.\n To use a FARGATElong capacity provider, specify either the ` + "`" + `` + "`" + `FARGATE` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `FARGATE_SPOT` + "`" + `` + "`" + ` capacity providers. The FARGATElong capacity providers are available to all accounts and only need to be associated with a cluster to be used.\n The [PutCapacityProvider](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PutCapacityProvider.html) API operation is used to update the list of available capacity providers for a cluster after the cluster is created.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "cluster_name": {
        "computed": true,
        "description": "A user-generated string that you use to identify your cluster. If you don't specify a name, CFNlong generates a unique physical ID for the name.",
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_settings": {
        "computed": true,
        "description": "The settings to use when creating a cluster. This parameter is used to turn on CloudWatch Container Insights for a cluster.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
              "computed": true,
              "description": "The name of the cluster setting. The value is ` + "`" + `` + "`" + `containerInsights` + "`" + `` + "`" + ` .",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value to set for the cluster setting. The supported values are ` + "`" + `` + "`" + `enabled` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `disabled` + "`" + `` + "`" + `. \n If you set ` + "`" + `` + "`" + `name` + "`" + `` + "`" + ` to ` + "`" + `` + "`" + `containerInsights` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `value` + "`" + `` + "`" + ` to ` + "`" + `` + "`" + `enabled` + "`" + `` + "`" + `, CloudWatch Container Insights will be on for the cluster, otherwise it will be off unless the ` + "`" + `` + "`" + `containerInsights` + "`" + `` + "`" + ` account setting is turned on. If a cluster value is specified, it will override the ` + "`" + `` + "`" + `containerInsights` + "`" + `` + "`" + ` value set with [PutAccountSetting](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PutAccountSetting.html) or [PutAccountSettingDefault](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PutAccountSettingDefault.html).",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "configuration": {
        "computed": true,
        "description": "The execute command configuration for the cluster.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "execute_command_configuration": {
              "computed": true,
              "description": "The details of the execute command configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "kms_key_id": {
                    "computed": true,
                    "description": "Specify an KMSlong key ID to encrypt the data between the local client and the container.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "log_configuration": {
                    "computed": true,
                    "description": "The log configuration for the results of the execute command actions. The logs can be sent to CloudWatch Logs or an Amazon S3 bucket. When ` + "`" + `` + "`" + `logging=OVERRIDE` + "`" + `` + "`" + ` is specified, a ` + "`" + `` + "`" + `logConfiguration` + "`" + `` + "`" + ` must be provided.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "cloudwatch_encryption_enabled": {
                          "computed": true,
                          "description": "Determines whether to use encryption on the CloudWatch logs. If not specified, encryption will be off.",
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "cloudwatch_log_group_name": {
                          "computed": true,
                          "description": "The name of the CloudWatch log group to send logs to.\n  The CloudWatch log group must already be created.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "s3_bucket_name": {
                          "computed": true,
                          "description": "The name of the S3 bucket to send logs to.\n  The S3 bucket must already be created.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "s3_encryption_enabled": {
                          "computed": true,
                          "description": "Determines whether to use encryption on the S3 logs. If not specified, encryption is not used.",
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "s3_key_prefix": {
                          "computed": true,
                          "description": "An optional folder in the S3 bucket to place logs in.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "logging": {
                    "computed": true,
                    "description": "The log setting to use for redirecting logs for your execute command results. The following log settings are available.\n  +   ` + "`" + `` + "`" + `NONE` + "`" + `` + "`" + `: The execute command session is not logged.\n  +   ` + "`" + `` + "`" + `DEFAULT` + "`" + `` + "`" + `: The ` + "`" + `` + "`" + `awslogs` + "`" + `` + "`" + ` configuration in the task definition is used. If no logging parameter is specified, it defaults to this value. If no ` + "`" + `` + "`" + `awslogs` + "`" + `` + "`" + ` log driver is configured in the task definition, the output won't be logged.\n  +   ` + "`" + `` + "`" + `OVERRIDE` + "`" + `` + "`" + `: Specify the logging details as a part of ` + "`" + `` + "`" + `logConfiguration` + "`" + `` + "`" + `. If the ` + "`" + `` + "`" + `OVERRIDE` + "`" + `` + "`" + ` logging option is specified, the ` + "`" + `` + "`" + `logConfiguration` + "`" + `` + "`" + ` is required.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "managed_storage_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "fargate_ephemeral_storage_kms_key_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "kms_key_id": {
                    "computed": true,
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
      "default_capacity_provider_strategy": {
        "computed": true,
        "description": "The default capacity provider strategy for the cluster. When services or tasks are run in the cluster with no launch type or capacity provider strategy specified, the default capacity provider strategy is used.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "base": {
              "computed": true,
              "description": "The *base* value designates how many tasks, at a minimum, to run on the specified capacity provider. Only one capacity provider in a capacity provider strategy can have a *base* defined. If no value is specified, the default value of ` + "`" + `` + "`" + `0` + "`" + `` + "`" + ` is used.",
              "description_kind": "plain",
              "type": "number"
            },
            "capacity_provider": {
              "computed": true,
              "description": "The short name of the capacity provider.",
              "description_kind": "plain",
              "type": "string"
            },
            "weight": {
              "computed": true,
              "description": "The *weight* value designates the relative percentage of the total number of tasks launched that should use the specified capacity provider. The ` + "`" + `` + "`" + `weight` + "`" + `` + "`" + ` value is taken into consideration after the ` + "`" + `` + "`" + `base` + "`" + `` + "`" + ` value, if defined, is satisfied.\n If no ` + "`" + `` + "`" + `weight` + "`" + `` + "`" + ` value is specified, the default value of ` + "`" + `` + "`" + `0` + "`" + `` + "`" + ` is used. When multiple capacity providers are specified within a capacity provider strategy, at least one of the capacity providers must have a weight value greater than zero and any capacity providers with a weight of ` + "`" + `` + "`" + `0` + "`" + `` + "`" + ` can't be used to place tasks. If you specify multiple capacity providers in a strategy that all have a weight of ` + "`" + `` + "`" + `0` + "`" + `` + "`" + `, any ` + "`" + `` + "`" + `RunTask` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `CreateService` + "`" + `` + "`" + ` actions using the capacity provider strategy will fail.\n An example scenario for using weights is defining a strategy that contains two capacity providers and both have a weight of ` + "`" + `` + "`" + `1` + "`" + `` + "`" + `, then when the ` + "`" + `` + "`" + `base` + "`" + `` + "`" + ` is satisfied, the tasks will be split evenly across the two capacity providers. Using that same logic, if you specify a weight of ` + "`" + `` + "`" + `1` + "`" + `` + "`" + ` for *capacityProviderA* and a weight of ` + "`" + `` + "`" + `4` + "`" + `` + "`" + ` for *capacityProviderB*, then for every one task that's run using *capacityProviderA*, four tasks would use *capacityProviderB*.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "list"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "service_connect_defaults": {
        "computed": true,
        "description": "Use this parameter to set a default Service Connect namespace. After you set a default Service Connect namespace, any new services with Service Connect turned on that are created in the cluster are added as client services in the namespace. This setting only applies to new services that set the ` + "`" + `` + "`" + `enabled` + "`" + `` + "`" + ` parameter to ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` in the ` + "`" + `` + "`" + `ServiceConnectConfiguration` + "`" + `` + "`" + `. You can set the namespace of each service individually in the ` + "`" + `` + "`" + `ServiceConnectConfiguration` + "`" + `` + "`" + ` to override this default parameter.\n Tasks that run in a namespace can use short names to connect to services in the namespace. Tasks can connect to services across all of the clusters in the namespace. Tasks connect through a managed proxy container that collects logs and metrics for increased visibility. Only the tasks that Amazon ECS services create are supported with Service Connect. For more information, see [Service Connect](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-connect.html) in the *Amazon Elastic Container Service Developer Guide*.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "namespace": {
              "computed": true,
              "description": "The namespace name or full Amazon Resource Name (ARN) of the CMAPlong namespace that's used when you create a service and don't specify a Service Connect configuration. The namespace name can include up to 1024 characters. The name is case-sensitive. The name can't include hyphens (-), tilde (~), greater than (\u003e), less than (\u003c), or slash (/).\n If you enter an existing namespace name or ARN, then that namespace will be used. Any namespace type is supported. The namespace must be in this account and this AWS Region.\n If you enter a new name, a CMAPlong namespace will be created. Amazon ECS creates a CMAP namespace with the \"API calls\" method of instance discovery only. This instance discovery method is the \"HTTP\" namespace type in the CLIlong. Other types of instance discovery aren't used by Service Connect.\n If you update the cluster with an empty string ` + "`" + `` + "`" + `\"\"` + "`" + `` + "`" + ` for the namespace name, the cluster configuration for Service Connect is removed. Note that the namespace will remain in CMAP and must be deleted separately.\n For more information about CMAPlong, see [Working with Services](https://docs.aws.amazon.com/cloud-map/latest/dg/working-with-services.html) in the *Developer Guide*.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description": "The metadata that you apply to the cluster to help you categorize and organize them. Each tag consists of a key and an optional value. You define both.\n The following basic restrictions apply to tags:\n  +  Maximum number of tags per resource - 50\n  +  For each resource, each tag key must be unique, and each tag key can have only one value.\n  +  Maximum key length - 128 Unicode characters in UTF-8\n  +  Maximum value length - 256 Unicode characters in UTF-8\n  +  If your tagging schema is used across multiple services and resources, remember that other services may have restrictions on allowed characters. Generally allowed characters are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.\n  +  Tag keys and values are case-sensitive.\n  +  Do not use ` + "`" + `` + "`" + `aws:` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `AWS:` + "`" + `` + "`" + `, or any upper or lowercase combination of such as a prefix for either keys or values as it is reserved for AWS use. You cannot edit or delete tag keys or values with this prefix. Tags with this prefix do not count against your tags per resource limit.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "One part of a key-value pair that make up a tag. A ` + "`" + `` + "`" + `key` + "`" + `` + "`" + ` is a general label that acts like a category for more specific tag values.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The optional part of a key-value pair that make up a tag. A ` + "`" + `` + "`" + `value` + "`" + `` + "`" + ` acts as a descriptor within a tag category (key).",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::ECS::Cluster",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEcsClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEcsCluster), &result)
	return &result
}
