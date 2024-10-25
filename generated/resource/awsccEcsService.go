package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEcsService = `{
  "block": {
    "attributes": {
      "capacity_provider_strategy": {
        "computed": true,
        "description": "The capacity provider strategy to use for the service.\n If a ` + "`" + `` + "`" + `capacityProviderStrategy` + "`" + `` + "`" + ` is specified, the ` + "`" + `` + "`" + `launchType` + "`" + `` + "`" + ` parameter must be omitted. If no ` + "`" + `` + "`" + `capacityProviderStrategy` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `launchType` + "`" + `` + "`" + ` is specified, the ` + "`" + `` + "`" + `defaultCapacityProviderStrategy` + "`" + `` + "`" + ` for the cluster is used.\n A capacity provider strategy may contain a maximum of 6 capacity providers.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "base": {
              "computed": true,
              "description": "The *base* value designates how many tasks, at a minimum, to run on the specified capacity provider. Only one capacity provider in a capacity provider strategy can have a *base* defined. If no value is specified, the default value of ` + "`" + `` + "`" + `0` + "`" + `` + "`" + ` is used.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "capacity_provider": {
              "computed": true,
              "description": "The short name of the capacity provider.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "weight": {
              "computed": true,
              "description": "The *weight* value designates the relative percentage of the total number of tasks launched that should use the specified capacity provider. The ` + "`" + `` + "`" + `weight` + "`" + `` + "`" + ` value is taken into consideration after the ` + "`" + `` + "`" + `base` + "`" + `` + "`" + ` value, if defined, is satisfied.\n If no ` + "`" + `` + "`" + `weight` + "`" + `` + "`" + ` value is specified, the default value of ` + "`" + `` + "`" + `0` + "`" + `` + "`" + ` is used. When multiple capacity providers are specified within a capacity provider strategy, at least one of the capacity providers must have a weight value greater than zero and any capacity providers with a weight of ` + "`" + `` + "`" + `0` + "`" + `` + "`" + ` can't be used to place tasks. If you specify multiple capacity providers in a strategy that all have a weight of ` + "`" + `` + "`" + `0` + "`" + `` + "`" + `, any ` + "`" + `` + "`" + `RunTask` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `CreateService` + "`" + `` + "`" + ` actions using the capacity provider strategy will fail.\n An example scenario for using weights is defining a strategy that contains two capacity providers and both have a weight of ` + "`" + `` + "`" + `1` + "`" + `` + "`" + `, then when the ` + "`" + `` + "`" + `base` + "`" + `` + "`" + ` is satisfied, the tasks will be split evenly across the two capacity providers. Using that same logic, if you specify a weight of ` + "`" + `` + "`" + `1` + "`" + `` + "`" + ` for *capacityProviderA* and a weight of ` + "`" + `` + "`" + `4` + "`" + `` + "`" + ` for *capacityProviderB*, then for every one task that's run using *capacityProviderA*, four tasks would use *capacityProviderB*.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "cluster": {
        "computed": true,
        "description": "The short name or full Amazon Resource Name (ARN) of the cluster that you run your service on. If you do not specify a cluster, the default cluster is assumed.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "deployment_configuration": {
        "computed": true,
        "description": "Optional deployment parameters that control how many tasks run during the deployment and the ordering of stopping and starting tasks.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "alarms": {
              "computed": true,
              "description": "Information about the CloudWatch alarms.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "alarm_names": {
                    "computed": true,
                    "description": "One or more CloudWatch alarm names. Use a \",\" to separate the alarms.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "enable": {
                    "computed": true,
                    "description": "Determines whether to use the CloudWatch alarm option in the service deployment process.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "rollback": {
                    "computed": true,
                    "description": "Determines whether to configure Amazon ECS to roll back the service if a service deployment fails. If rollback is used, when a service deployment fails, the service is rolled back to the last deployment that completed successfully.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "deployment_circuit_breaker": {
              "computed": true,
              "description": "The deployment circuit breaker can only be used for services using the rolling update (` + "`" + `` + "`" + `ECS` + "`" + `` + "`" + `) deployment type.\n  The *deployment circuit breaker* determines whether a service deployment will fail if the service can't reach a steady state. If you use the deployment circuit breaker, a service deployment will transition to a failed state and stop launching new tasks. If you use the rollback option, when a service deployment fails, the service is rolled back to the last deployment that completed successfully. For more information, see [Rolling update](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-type-ecs.html) in the *Amazon Elastic Container Service Developer Guide*",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enable": {
                    "computed": true,
                    "description": "Determines whether to use the deployment circuit breaker logic for the service.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "rollback": {
                    "computed": true,
                    "description": "Determines whether to configure Amazon ECS to roll back the service if a service deployment fails. If rollback is on, when a service deployment fails, the service is rolled back to the last deployment that completed successfully.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "maximum_percent": {
              "computed": true,
              "description": "If a service is using the rolling update (` + "`" + `` + "`" + `ECS` + "`" + `` + "`" + `) deployment type, the ` + "`" + `` + "`" + `maximumPercent` + "`" + `` + "`" + ` parameter represents an upper limit on the number of your service's tasks that are allowed in the ` + "`" + `` + "`" + `RUNNING` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `PENDING` + "`" + `` + "`" + ` state during a deployment, as a percentage of the ` + "`" + `` + "`" + `desiredCount` + "`" + `` + "`" + ` (rounded down to the nearest integer). This parameter enables you to define the deployment batch size. For example, if your service is using the ` + "`" + `` + "`" + `REPLICA` + "`" + `` + "`" + ` service scheduler and has a ` + "`" + `` + "`" + `desiredCount` + "`" + `` + "`" + ` of four tasks and a ` + "`" + `` + "`" + `maximumPercent` + "`" + `` + "`" + ` value of 200%, the scheduler may start four new tasks before stopping the four older tasks (provided that the cluster resources required to do this are available). The default ` + "`" + `` + "`" + `maximumPercent` + "`" + `` + "`" + ` value for a service using the ` + "`" + `` + "`" + `REPLICA` + "`" + `` + "`" + ` service scheduler is 200%.\n If a service is using either the blue/green (` + "`" + `` + "`" + `CODE_DEPLOY` + "`" + `` + "`" + `) or ` + "`" + `` + "`" + `EXTERNAL` + "`" + `` + "`" + ` deployment types, and tasks in the service use the EC2 launch type, the *maximum percent* value is set to the default value. The *maximum percent* value is used to define the upper limit on the number of the tasks in the service that remain in the ` + "`" + `` + "`" + `RUNNING` + "`" + `` + "`" + ` state while the container instances are in the ` + "`" + `` + "`" + `DRAINING` + "`" + `` + "`" + ` state.\n  You can't specify a custom ` + "`" + `` + "`" + `maximumPercent` + "`" + `` + "`" + ` value for a service that uses either the blue/green (` + "`" + `` + "`" + `CODE_DEPLOY` + "`" + `` + "`" + `) or ` + "`" + `` + "`" + `EXTERNAL` + "`" + `` + "`" + ` deployment types and has tasks that use the EC2 launch type.\n  If the tasks in the service use the Fargate launch type, the maximum percent value is not used, although it is returned when describing your service.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "minimum_healthy_percent": {
              "computed": true,
              "description": "If a service is using the rolling update (` + "`" + `` + "`" + `ECS` + "`" + `` + "`" + `) deployment type, the ` + "`" + `` + "`" + `minimumHealthyPercent` + "`" + `` + "`" + ` represents a lower limit on the number of your service's tasks that must remain in the ` + "`" + `` + "`" + `RUNNING` + "`" + `` + "`" + ` state during a deployment, as a percentage of the ` + "`" + `` + "`" + `desiredCount` + "`" + `` + "`" + ` (rounded up to the nearest integer). This parameter enables you to deploy without using additional cluster capacity. For example, if your service has a ` + "`" + `` + "`" + `desiredCount` + "`" + `` + "`" + ` of four tasks and a ` + "`" + `` + "`" + `minimumHealthyPercent` + "`" + `` + "`" + ` of 50%, the service scheduler may stop two existing tasks to free up cluster capacity before starting two new tasks. \n For services that *do not* use a load balancer, the following should be noted:\n  +  A service is considered healthy if all essential containers within the tasks in the service pass their health checks.\n  +  If a task has no essential containers with a health check defined, the service scheduler will wait for 40 seconds after a task reaches a ` + "`" + `` + "`" + `RUNNING` + "`" + `` + "`" + ` state before the task is counted towards the minimum healthy percent total.\n  +  If a task has one or more essential containers with a health check defined, the service scheduler will wait for the task to reach a healthy status before counting it towards the minimum healthy percent total. A task is considered healthy when all essential containers within the task have passed their health checks. The amount of time the service scheduler can wait for is determined by the container health check settings. \n  \n For services that *do* use a load balancer, the following should be noted:\n  +  If a task has no essential containers with a health check defined, the service scheduler will wait for the load balancer target group health check to return a healthy status before counting the task towards the minimum healthy percent total.\n  +  If a task has an essential container with a health check defined, the service scheduler will wait for both the task to reach a healthy status and the load balancer target group health check to return a healthy status before counting the task towards the minimum healthy percent total.\n  \n The default value for a replica service for ` + "`" + `` + "`" + `minimumHealthyPercent` + "`" + `` + "`" + ` is 100%. The default ` + "`" + `` + "`" + `minimumHealthyPercent` + "`" + `` + "`" + ` value for a service using the ` + "`" + `` + "`" + `DAEMON` + "`" + `` + "`" + ` service schedule is 0% for the CLI, the AWS SDKs, and the APIs and 50% for the AWS Management Console.\n The minimum number of healthy tasks during a deployment is the ` + "`" + `` + "`" + `desiredCount` + "`" + `` + "`" + ` multiplied by the ` + "`" + `` + "`" + `minimumHealthyPercent` + "`" + `` + "`" + `/100, rounded up to the nearest integer value.\n If a service is using either the blue/green (` + "`" + `` + "`" + `CODE_DEPLOY` + "`" + `` + "`" + `) or ` + "`" + `` + "`" + `EXTERNAL` + "`" + `` + "`" + ` deployment types and is running tasks that use the EC2 launch type, the *minimum healthy percent* value is set to the default value. The *minimum healthy percent* value is used to define the lower limit on the number of the tasks in the service that remain in the ` + "`" + `` + "`" + `RUNNING` + "`" + `` + "`" + ` state while the container instances are in the ` + "`" + `` + "`" + `DRAINING` + "`" + `` + "`" + ` state.\n  You can't specify a custom ` + "`" + `` + "`" + `minimumHealthyPercent` + "`" + `` + "`" + ` value for a service that uses either the blue/green (` + "`" + `` + "`" + `CODE_DEPLOY` + "`" + `` + "`" + `) or ` + "`" + `` + "`" + `EXTERNAL` + "`" + `` + "`" + ` deployment types and has tasks that use the EC2 launch type.\n  If a service is using either the blue/green (` + "`" + `` + "`" + `CODE_DEPLOY` + "`" + `` + "`" + `) or ` + "`" + `` + "`" + `EXTERNAL` + "`" + `` + "`" + ` deployment types and is running tasks that use the Fargate launch type, the minimum healthy percent value is not used, although it is returned when describing your service.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "deployment_controller": {
        "computed": true,
        "description": "The deployment controller to use for the service. If no deployment controller is specified, the default value of ` + "`" + `` + "`" + `ECS` + "`" + `` + "`" + ` is used.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "type": {
              "computed": true,
              "description": "The deployment controller type to use. There are three deployment controller types available:\n  + ECS The rolling update (ECS) deployment type involves replacing the current running version of the container with the latest version. The number of containers Amazon ECS adds or removes from the service during a rolling update is controlled by adjusting the minimum and maximum number of healthy tasks allowed during a service deployment, as specified in the DeploymentConfiguration. + CODE_DEPLOY The blue/green (CODE_DEPLOY) deployment type uses the blue/green deployment model powered by , which allows you to verify a new deployment of a service before sending production traffic to it. + EXTERNAL The external (EXTERNAL) deployment type enables you to use any third-party deployment controller for full control over the deployment process for an Amazon ECS service.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "desired_count": {
        "computed": true,
        "description": "The number of instantiations of the specified task definition to place and keep running in your service.\n For new services, if a desired count is not specified, a default value of ` + "`" + `` + "`" + `1` + "`" + `` + "`" + ` is used. When using the ` + "`" + `` + "`" + `DAEMON` + "`" + `` + "`" + ` scheduling strategy, the desired count is not required.\n For existing services, if a desired count is not specified, it is omitted from the operation.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "enable_ecs_managed_tags": {
        "computed": true,
        "description": "Specifies whether to turn on Amazon ECS managed tags for the tasks within the service. For more information, see [Tagging your Amazon ECS resources](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html) in the *Amazon Elastic Container Service Developer Guide*.\n When you use Amazon ECS managed tags, you need to set the ` + "`" + `` + "`" + `propagateTags` + "`" + `` + "`" + ` request parameter.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enable_execute_command": {
        "computed": true,
        "description": "Determines whether the execute command functionality is turned on for the service. If ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `, the execute command functionality is turned on for all containers in tasks as part of the service.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "health_check_grace_period_seconds": {
        "computed": true,
        "description": "The period of time, in seconds, that the Amazon ECS service scheduler ignores unhealthy Elastic Load Balancing target health checks after a task has first started. This is only used when your service is configured to use a load balancer. If your service has a load balancer defined and you don't specify a health check grace period value, the default value of ` + "`" + `` + "`" + `0` + "`" + `` + "`" + ` is used.\n If you do not use an Elastic Load Balancing, we recommend that you use the ` + "`" + `` + "`" + `startPeriod` + "`" + `` + "`" + ` in the task definition health check parameters. For more information, see [Health check](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_HealthCheck.html).\n If your service's tasks take a while to start and respond to Elastic Load Balancing health checks, you can specify a health check grace period of up to 2,147,483,647 seconds (about 69 years). During that time, the Amazon ECS service scheduler ignores health check status. This grace period can prevent the service scheduler from marking tasks as unhealthy and stopping them before they have time to come up.",
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
      "launch_type": {
        "computed": true,
        "description": "The launch type on which to run your service. For more information, see [Amazon ECS Launch Types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html) in the *Amazon Elastic Container Service Developer Guide*.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "load_balancers": {
        "computed": true,
        "description": "A list of load balancer objects to associate with the service. If you specify the ` + "`" + `` + "`" + `Role` + "`" + `` + "`" + ` property, ` + "`" + `` + "`" + `LoadBalancers` + "`" + `` + "`" + ` must be specified as well. For information about the number of load balancers that you can specify per service, see [Service Load Balancing](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-load-balancing.html) in the *Amazon Elastic Container Service Developer Guide*.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "container_name": {
              "computed": true,
              "description": "The name of the container (as it appears in a container definition) to associate with the load balancer.\n You need to specify the container name when configuring the target group for an Amazon ECS load balancer.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "container_port": {
              "computed": true,
              "description": "The port on the container to associate with the load balancer. This port must correspond to a ` + "`" + `` + "`" + `containerPort` + "`" + `` + "`" + ` in the task definition the tasks in the service are using. For tasks that use the EC2 launch type, the container instance they're launched on must allow ingress traffic on the ` + "`" + `` + "`" + `hostPort` + "`" + `` + "`" + ` of the port mapping.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "load_balancer_name": {
              "computed": true,
              "description": "The name of the load balancer to associate with the Amazon ECS service or task set.\n If you are using an Application Load Balancer or a Network Load Balancer the load balancer name parameter should be omitted.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "target_group_arn": {
              "computed": true,
              "description": "The full Amazon Resource Name (ARN) of the Elastic Load Balancing target group or groups associated with a service or task set.\n A target group ARN is only specified when using an Application Load Balancer or Network Load Balancer. \n For services using the ` + "`" + `` + "`" + `ECS` + "`" + `` + "`" + ` deployment controller, you can specify one or multiple target groups. For more information, see [Registering multiple target groups with a service](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/register-multiple-targetgroups.html) in the *Amazon Elastic Container Service Developer Guide*.\n For services using the ` + "`" + `` + "`" + `CODE_DEPLOY` + "`" + `` + "`" + ` deployment controller, you're required to define two target groups for the load balancer. For more information, see [Blue/green deployment with CodeDeploy](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-type-bluegreen.html) in the *Amazon Elastic Container Service Developer Guide*.\n  If your service's task definition uses the ` + "`" + `` + "`" + `awsvpc` + "`" + `` + "`" + ` network mode, you must choose ` + "`" + `` + "`" + `ip` + "`" + `` + "`" + ` as the target type, not ` + "`" + `` + "`" + `instance` + "`" + `` + "`" + `. Do this when creating your target groups because tasks that use the ` + "`" + `` + "`" + `awsvpc` + "`" + `` + "`" + ` network mode are associated with an elastic network interface, not an Amazon EC2 instance. This network mode is required for the Fargate launch type.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "network_configuration": {
        "computed": true,
        "description": "The network configuration for the service. This parameter is required for task definitions that use the ` + "`" + `` + "`" + `awsvpc` + "`" + `` + "`" + ` network mode to receive their own elastic network interface, and it is not supported for other network modes. For more information, see [Task Networking](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-networking.html) in the *Amazon Elastic Container Service Developer Guide*.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "awsvpc_configuration": {
              "computed": true,
              "description": "The VPC subnets and security groups that are associated with a task.\n  All specified subnets and security groups must be from the same VPC.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "assign_public_ip": {
                    "computed": true,
                    "description": "Whether the task's elastic network interface receives a public IP address. The default value is ` + "`" + `` + "`" + `DISABLED` + "`" + `` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "security_groups": {
                    "computed": true,
                    "description": "The IDs of the security groups associated with the task or service. If you don't specify a security group, the default security group for the VPC is used. There's a limit of 5 security groups that can be specified per ` + "`" + `` + "`" + `awsvpcConfiguration` + "`" + `` + "`" + `.\n  All specified security groups must be from the same VPC.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "subnets": {
                    "computed": true,
                    "description": "The IDs of the subnets associated with the task or service. There's a limit of 16 subnets that can be specified per ` + "`" + `` + "`" + `awsvpcConfiguration` + "`" + `` + "`" + `.\n  All specified subnets must be from the same VPC.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
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
      "placement_constraints": {
        "computed": true,
        "description": "An array of placement constraint objects to use for tasks in your service. You can specify a maximum of 10 constraints for each task. This limit includes constraints in the task definition and those specified at runtime.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "expression": {
              "computed": true,
              "description": "A cluster query language expression to apply to the constraint. The expression can have a maximum length of 2000 characters. You can't specify an expression if the constraint type is ` + "`" + `` + "`" + `distinctInstance` + "`" + `` + "`" + `. For more information, see [Cluster query language](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-query-language.html) in the *Amazon Elastic Container Service Developer Guide*.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "computed": true,
              "description": "The type of constraint. Use ` + "`" + `` + "`" + `distinctInstance` + "`" + `` + "`" + ` to ensure that each task in a particular group is running on a different container instance. Use ` + "`" + `` + "`" + `memberOf` + "`" + `` + "`" + ` to restrict the selection to a group of valid candidates.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "placement_strategies": {
        "computed": true,
        "description": "The placement strategy objects to use for tasks in your service. You can specify a maximum of 5 strategy rules for each service.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "field": {
              "computed": true,
              "description": "The field to apply the placement strategy against. For the ` + "`" + `` + "`" + `spread` + "`" + `` + "`" + ` placement strategy, valid values are ` + "`" + `` + "`" + `instanceId` + "`" + `` + "`" + ` (or ` + "`" + `` + "`" + `host` + "`" + `` + "`" + `, which has the same effect), or any platform or custom attribute that's applied to a container instance, such as ` + "`" + `` + "`" + `attribute:ecs.availability-zone` + "`" + `` + "`" + `. For the ` + "`" + `` + "`" + `binpack` + "`" + `` + "`" + ` placement strategy, valid values are ` + "`" + `` + "`" + `cpu` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `memory` + "`" + `` + "`" + `. For the ` + "`" + `` + "`" + `random` + "`" + `` + "`" + ` placement strategy, this field is not used.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "computed": true,
              "description": "The type of placement strategy. The ` + "`" + `` + "`" + `random` + "`" + `` + "`" + ` placement strategy randomly places tasks on available candidates. The ` + "`" + `` + "`" + `spread` + "`" + `` + "`" + ` placement strategy spreads placement across available candidates evenly based on the ` + "`" + `` + "`" + `field` + "`" + `` + "`" + ` parameter. The ` + "`" + `` + "`" + `binpack` + "`" + `` + "`" + ` strategy places tasks on available candidates that have the least available amount of the resource that's specified with the ` + "`" + `` + "`" + `field` + "`" + `` + "`" + ` parameter. For example, if you binpack on memory, a task is placed on the instance with the least amount of remaining memory but still enough to run the task.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "platform_version": {
        "computed": true,
        "description": "The platform version that your tasks in the service are running on. A platform version is specified only for tasks using the Fargate launch type. If one isn't specified, the ` + "`" + `` + "`" + `LATEST` + "`" + `` + "`" + ` platform version is used. For more information, see [platform versions](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html) in the *Amazon Elastic Container Service Developer Guide*.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "propagate_tags": {
        "computed": true,
        "description": "Specifies whether to propagate the tags from the task definition to the task. If no value is specified, the tags aren't propagated. Tags can only be propagated to the task during task creation. To add tags to a task after task creation, use the [TagResource](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_TagResource.html) API action.\n You must set this to a value other than ` + "`" + `` + "`" + `NONE` + "`" + `` + "`" + ` when you use Cost Explorer. For more information, see [Amazon ECS usage reports](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/usage-reports.html) in the *Amazon Elastic Container Service Developer Guide*.\n The default is ` + "`" + `` + "`" + `NONE` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "role": {
        "computed": true,
        "description": "The name or full Amazon Resource Name (ARN) of the IAM role that allows Amazon ECS to make calls to your load balancer on your behalf. This parameter is only permitted if you are using a load balancer with your service and your task definition doesn't use the ` + "`" + `` + "`" + `awsvpc` + "`" + `` + "`" + ` network mode. If you specify the ` + "`" + `` + "`" + `role` + "`" + `` + "`" + ` parameter, you must also specify a load balancer object with the ` + "`" + `` + "`" + `loadBalancers` + "`" + `` + "`" + ` parameter.\n  If your account has already created the Amazon ECS service-linked role, that role is used for your service unless you specify a role here. The service-linked role is required if your task definition uses the ` + "`" + `` + "`" + `awsvpc` + "`" + `` + "`" + ` network mode or if the service is configured to use service discovery, an external deployment controller, multiple target groups, or Elastic Inference accelerators in which case you don't specify a role here. For more information, see [Using service-linked roles for Amazon ECS](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using-service-linked-roles.html) in the *Amazon Elastic Container Service Developer Guide*.\n  If your specified role has a path other than ` + "`" + `` + "`" + `/` + "`" + `` + "`" + `, then you must either specify the full role ARN (this is recommended) or prefix the role name with the path. For example, if a role with the name ` + "`" + `` + "`" + `bar` + "`" + `` + "`" + ` has a path of ` + "`" + `` + "`" + `/foo/` + "`" + `` + "`" + ` then you would specify ` + "`" + `` + "`" + `/foo/bar` + "`" + `` + "`" + ` as the role name. For more information, see [Friendly names and paths](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-friendly-names) in the *IAM User Guide*.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scheduling_strategy": {
        "computed": true,
        "description": "The scheduling strategy to use for the service. For more information, see [Services](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs_services.html).\n There are two service scheduler strategies available:\n  +   ` + "`" + `` + "`" + `REPLICA` + "`" + `` + "`" + `-The replica scheduling strategy places and maintains the desired number of tasks across your cluster. By default, the service scheduler spreads tasks across Availability Zones. You can use task placement strategies and constraints to customize task placement decisions. This scheduler strategy is required if the service uses the ` + "`" + `` + "`" + `CODE_DEPLOY` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `EXTERNAL` + "`" + `` + "`" + ` deployment controller types.\n  +   ` + "`" + `` + "`" + `DAEMON` + "`" + `` + "`" + `-The daemon scheduling strategy deploys exactly one task on each active container instance that meets all of the task placement constraints that you specify in your cluster. The service scheduler also evaluates the task placement constraints for running tasks and will stop tasks that don't meet the placement constraints. When you're using this strategy, you don't need to specify a desired number of tasks, a task placement strategy, or use Service Auto Scaling policies.\n  Tasks using the Fargate launch type or the ` + "`" + `` + "`" + `CODE_DEPLOY` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `EXTERNAL` + "`" + `` + "`" + ` deployment controller types don't support the ` + "`" + `` + "`" + `DAEMON` + "`" + `` + "`" + ` scheduling strategy.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_connect_configuration": {
        "computed": true,
        "description": "The configuration for this service to discover and connect to services, and be discovered by, and connected from, other services within a namespace.\n Tasks that run in a namespace can use short names to connect to services in the namespace. Tasks can connect to services across all of the clusters in the namespace. Tasks connect through a managed proxy container that collects logs and metrics for increased visibility. Only the tasks that Amazon ECS services create are supported with Service Connect. For more information, see [Service Connect](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-connect.html) in the *Amazon Elastic Container Service Developer Guide*.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enabled": {
              "computed": true,
              "description": "Specifies whether to use Service Connect with this service.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "log_configuration": {
              "computed": true,
              "description": "The log configuration for the container. This parameter maps to ` + "`" + `` + "`" + `LogConfig` + "`" + `` + "`" + ` in the docker container create command and the ` + "`" + `` + "`" + `--log-driver` + "`" + `` + "`" + ` option to docker run.\n By default, containers use the same logging driver that the Docker daemon uses. However, the container might use a different logging driver than the Docker daemon by specifying a log driver configuration in the container definition.\n Understand the following when specifying a log configuration for your containers.\n  +  Amazon ECS currently supports a subset of the logging drivers available to the Docker daemon. Additional log drivers may be available in future releases of the Amazon ECS container agent.\n For tasks on FARGATElong, the supported log drivers are ` + "`" + `` + "`" + `awslogs` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `splunk` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `awsfirelens` + "`" + `` + "`" + `.\n For tasks hosted on Amazon EC2 instances, the supported log drivers are ` + "`" + `` + "`" + `awslogs` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `fluentd` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `gelf` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `json-file` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `journald` + "`" + `` + "`" + `,` + "`" + `` + "`" + `syslog` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `splunk` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `awsfirelens` + "`" + `` + "`" + `.\n  +  This parameter requires version 1.18 of the Docker Remote API or greater on your container instance.\n  +  For tasks that are hosted on Amazon EC2 instances, the Amazon ECS container agent must register the available logging drivers with the ` + "`" + `` + "`" + `ECS_AVAILABLE_LOGGING_DRIVERS` + "`" + `` + "`" + ` environment variable before containers placed on that instance can use these log configuration options. For more information, see [Amazon ECS container agent configuration](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-config.html) in the *Amazon Elastic Container Service Developer Guide*.\n  +  For tasks that are on FARGATElong, because you don't have access to the underlying infrastructure your tasks are hosted on, any additional software needed must be installed outside of the task. For example, the Fluentd output aggregators or a remote host running Logstash to send Gelf logs to.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "log_driver": {
                    "computed": true,
                    "description": "The log driver to use for the container.\n For tasks on FARGATElong, the supported log drivers are ` + "`" + `` + "`" + `awslogs` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `splunk` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `awsfirelens` + "`" + `` + "`" + `.\n For tasks hosted on Amazon EC2 instances, the supported log drivers are ` + "`" + `` + "`" + `awslogs` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `fluentd` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `gelf` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `json-file` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `journald` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `syslog` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `splunk` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `awsfirelens` + "`" + `` + "`" + `.\n For more information about using the ` + "`" + `` + "`" + `awslogs` + "`" + `` + "`" + ` log driver, see [Send Amazon ECS logs to CloudWatch](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_awslogs.html) in the *Amazon Elastic Container Service Developer Guide*.\n For more information about using the ` + "`" + `` + "`" + `awsfirelens` + "`" + `` + "`" + ` log driver, see [Send Amazon ECS logs to an service or Partner](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_firelens.html).\n  If you have a custom driver that isn't listed, you can fork the Amazon ECS container agent project that's [available on GitHub](https://docs.aws.amazon.com/https://github.com/aws/amazon-ecs-agent) and customize it to work with that driver. We encourage you to submit pull requests for changes that you would like to have included. However, we don't currently provide support for running modified copies of this software.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "options": {
                    "computed": true,
                    "description": "The configuration options to send to the log driver.\n The options you can specify depend on the log driver. Some of the options you can specify when you use the ` + "`" + `` + "`" + `awslogs` + "`" + `` + "`" + ` log driver to route logs to Amazon CloudWatch include the following:\n  + awslogs-create-group Required: No Specify whether you want the log group to be created automatically. If this option isn't specified, it defaults to false. Your IAM policy must include the logs:CreateLogGroup permission before you attempt to use awslogs-create-group. + awslogs-region Required: Yes Specify the Region that the awslogs log driver is to send your Docker logs to. You can choose to send all of your logs from clusters in different Regions to a single region in CloudWatch Logs. This is so that they're all visible in one location. Otherwise, you can separate them by Region for more granularity. Make sure that the specified log group exists in the Region that you specify with this option. + awslogs-group Required: Yes Make sure to specify a log group that the awslogs log driver sends its log streams to. + awslogs-stream-prefix Required: Yes, when using the Fargate launch type.Optional for the EC2 launch type, required for the Fargate launch type. Use the awslogs-stream-prefix option to associate a log stream with the specified prefix, the container name, and the ID of the Amazon ECS task that the container belongs to. If you specify a prefix with this option, then the log stream takes the format prefix-name/container-name/ecs-task-id. If you don't specify a prefix with this option, then the log stream is named after the container ID that's assigned by the Docker daemon on the container instance. Because it's difficult to trace logs back to the container that sent them with just the Docker container ID (which is only available on the container instance), we recommend that you specify a prefix with this option. For Amazon ECS services, you can use the service name as the prefix. Doing so, you can trace log streams to the service that the container belongs to, the name of the container that sent them, and the ID of the task that the container belongs to. You must specify a stream-prefix for your logs to have your logs appear in the Log pane when using the Amazon ECS console. + awslogs-datetime-format Required: No This option defines a multiline start pattern in Python strftime format. A log message consists of a line that matches the pattern and any following lines that don’t match the pattern. The matched line is the delimiter between log messages. One example of a use case for using this format is for parsing output such as a stack dump, which might otherwise be logged in multiple entries. The correct pattern allows it to be captured in a single entry. For more information, see awslogs-datetime-format. You cannot configure both the awslogs-datetime-format and awslogs-multiline-pattern options. Multiline logging performs regular expression parsing and matching of all log messages. This might have a negative impact on logging performance. + awslogs-multiline-pattern Required: No This option defines a multiline start pattern that uses a regular expression. A log message consists of a line that matches the pattern and any following lines that don’t match the pattern. The matched line is the delimiter between log messages. For more information, see awslogs-multiline-pattern. This option is ignored if awslogs-datetime-format is also configured. You cannot configure both the awslogs-datetime-format and awslogs-multiline-pattern options. Multiline logging performs regular expression parsing and matching of all log messages. This might have a negative impact on logging performance. + mode Required: No Valid values: non-blocking | blocking This option defines the delivery mode of log messages from the container to CloudWatch Logs. The delivery mode you choose affects application availability when the flow of logs from container to CloudWatch is interrupted. If you use the blocking mode and the flow of logs to CloudWatch is interrupted, calls from container code to write to the stdout and stderr streams will block. The logging thread of the application will block as a result. This may cause the application to become unresponsive and lead to container healthcheck failure. If you use the non-blocking mode, the container's logs are instead stored in an in-memory intermediate buffer configured with the max-buffer-size option. This prevents the application from becoming unresponsive when logs cannot be sent to CloudWatch. We recommend using this mode if you want to ensure service availability and are okay with some log loss. For more information, see Preventing log loss with non-blocking mode in the awslogs container log driver. + max-buffer-size Required: No Default value: 1m When non-blocking mode is used, the max-buffer-size log option controls the size of the buffer that's used for intermediate message storage. Make sure to specify an adequate buffer size based on your application. When the buffer fills up, further logs cannot be stored. Logs that cannot be stored are lost. \n To route logs using the ` + "`" + `` + "`" + `splunk` + "`" + `` + "`" + ` log router, you need to specify a ` + "`" + `` + "`" + `splunk-token` + "`" + `` + "`" + ` and a ` + "`" + `` + "`" + `splunk-url` + "`" + `` + "`" + `.\n When you use the ` + "`" + `` + "`" + `awsfirelens` + "`" + `` + "`" + ` log router to route logs to an AWS Service or AWS Partner Network destination for log storage and analytics, you can set the ` + "`" + `` + "`" + `log-driver-buffer-limit` + "`" + `` + "`" + ` option to limit the number of events that are buffered in memory, before being sent to the log router container. It can help to resolve potential log loss issue because high throughput might result in memory running out for the buffer inside of Docker.\n Other options you can specify when using ` + "`" + `` + "`" + `awsfirelens` + "`" + `` + "`" + ` to route logs depend on the destination. When you export logs to Amazon Data Firehose, you can specify the AWS Region with ` + "`" + `` + "`" + `region` + "`" + `` + "`" + ` and a name for the log stream with ` + "`" + `` + "`" + `delivery_stream` + "`" + `` + "`" + `.\n When you export logs to Amazon Kinesis Data Streams, you can specify an AWS Region with ` + "`" + `` + "`" + `region` + "`" + `` + "`" + ` and a data stream name with ` + "`" + `` + "`" + `stream` + "`" + `` + "`" + `.\n  When you export logs to Amazon OpenSearch Service, you can specify options like ` + "`" + `` + "`" + `Name` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `Host` + "`" + `` + "`" + ` (OpenSearch Service endpoint without protocol), ` + "`" + `` + "`" + `Port` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `Index` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `Type` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `Aws_auth` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `Aws_region` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `Suppress_Type_Name` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `tls` + "`" + `` + "`" + `.\n When you export logs to Amazon S3, you can specify the bucket using the ` + "`" + `` + "`" + `bucket` + "`" + `` + "`" + ` option. You can also specify ` + "`" + `` + "`" + `region` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `total_file_size` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `upload_timeout` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `use_put_object` + "`" + `` + "`" + ` as options.\n This parameter requires version 1.19 of the Docker Remote API or greater on your container instance. To check the Docker Remote API version on your container instance, log in to your container instance and run the following command: ` + "`" + `` + "`" + `sudo docker version --format '{{.Server.APIVersion}}'` + "`" + `` + "`" + `",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "secret_options": {
                    "computed": true,
                    "description": "The secrets to pass to the log configuration. For more information, see [Specifying sensitive data](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/specifying-sensitive-data.html) in the *Amazon Elastic Container Service Developer Guide*.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "name": {
                          "computed": true,
                          "description": "The name of the secret.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value_from": {
                          "computed": true,
                          "description": "The secret to expose to the container. The supported values are either the full ARN of the ASMlong secret or the full ARN of the parameter in the SSM Parameter Store.\n For information about the require IAMlong permissions, see [Required IAM permissions for Amazon ECS secrets](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/specifying-sensitive-data-secrets.html#secrets-iam) (for Secrets Manager) or [Required IAM permissions for Amazon ECS secrets](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/specifying-sensitive-data-parameters.html) (for Systems Manager Parameter store) in the *Amazon Elastic Container Service Developer Guide*.\n  If the SSM Parameter Store parameter exists in the same Region as the task you're launching, then you can use either the full ARN or name of the parameter. If the parameter exists in a different Region, then the full ARN must be specified.",
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
                "nesting_mode": "single"
              },
              "optional": true
            },
            "namespace": {
              "computed": true,
              "description": "The namespace name or full Amazon Resource Name (ARN) of the CMAPlong namespace for use with Service Connect. The namespace must be in the same AWS Region as the Amazon ECS service and cluster. The type of namespace doesn't affect Service Connect. For more information about CMAPlong, see [Working with Services](https://docs.aws.amazon.com/cloud-map/latest/dg/working-with-services.html) in the *Developer Guide*.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "services": {
              "computed": true,
              "description": "The list of Service Connect service objects. These are names and aliases (also known as endpoints) that are used by other Amazon ECS services to connect to this service. \n This field is not required for a \"client\" Amazon ECS service that's a member of a namespace only to connect to other services within the namespace. An example of this would be a frontend application that accepts incoming requests from either a load balancer that's attached to the service or by other means.\n An object selects a port from the task definition, assigns a name for the CMAPlong service, and a list of aliases (endpoints) and ports for client applications to refer to this service.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "client_aliases": {
                    "computed": true,
                    "description": "The list of client aliases for this Service Connect service. You use these to assign names that can be used by client applications. The maximum number of client aliases that you can have in this list is 1.\n Each alias (\"endpoint\") is a fully-qualified name and port number that other Amazon ECS tasks (\"clients\") can use to connect to this service.\n Each name and port mapping must be unique within the namespace.\n For each ` + "`" + `` + "`" + `ServiceConnectService` + "`" + `` + "`" + `, you must provide at least one ` + "`" + `` + "`" + `clientAlias` + "`" + `` + "`" + ` with one ` + "`" + `` + "`" + `port` + "`" + `` + "`" + `.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "dns_name": {
                          "computed": true,
                          "description": "The ` + "`" + `` + "`" + `dnsName` + "`" + `` + "`" + ` is the name that you use in the applications of client tasks to connect to this service. The name must be a valid DNS name but doesn't need to be fully-qualified. The name can include up to 127 characters. The name can include lowercase letters, numbers, underscores (_), hyphens (-), and periods (.). The name can't start with a hyphen.\n If this parameter isn't specified, the default value of ` + "`" + `` + "`" + `discoveryName.namespace` + "`" + `` + "`" + ` is used. If the ` + "`" + `` + "`" + `discoveryName` + "`" + `` + "`" + ` isn't specified, the port mapping name from the task definition is used in ` + "`" + `` + "`" + `portName.namespace` + "`" + `` + "`" + `.\n To avoid changing your applications in client Amazon ECS services, set this to the same name that the client application uses by default. For example, a few common names are ` + "`" + `` + "`" + `database` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `db` + "`" + `` + "`" + `, or the lowercase name of a database, such as ` + "`" + `` + "`" + `mysql` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `redis` + "`" + `` + "`" + `. For more information, see [Service Connect](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-connect.html) in the *Amazon Elastic Container Service Developer Guide*.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "port": {
                          "computed": true,
                          "description": "The listening port number for the Service Connect proxy. This port is available inside of all of the tasks within the same namespace.\n To avoid changing your applications in client Amazon ECS services, set this to the same port that the client application uses by default. For more information, see [Service Connect](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-connect.html) in the *Amazon Elastic Container Service Developer Guide*.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "discovery_name": {
                    "computed": true,
                    "description": "The ` + "`" + `` + "`" + `discoveryName` + "`" + `` + "`" + ` is the name of the new CMAP service that Amazon ECS creates for this Amazon ECS service. This must be unique within the CMAP namespace. The name can contain up to 64 characters. The name can include lowercase letters, numbers, underscores (_), and hyphens (-). The name can't start with a hyphen.\n If the ` + "`" + `` + "`" + `discoveryName` + "`" + `` + "`" + ` isn't specified, the port mapping name from the task definition is used in ` + "`" + `` + "`" + `portName.namespace` + "`" + `` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "ingress_port_override": {
                    "computed": true,
                    "description": "The port number for the Service Connect proxy to listen on.\n Use the value of this field to bypass the proxy for traffic on the port number specified in the named ` + "`" + `` + "`" + `portMapping` + "`" + `` + "`" + ` in the task definition of this application, and then use it in your VPC security groups to allow traffic into the proxy for this Amazon ECS service.\n In ` + "`" + `` + "`" + `awsvpc` + "`" + `` + "`" + ` mode and Fargate, the default value is the container port number. The container port number is in the ` + "`" + `` + "`" + `portMapping` + "`" + `` + "`" + ` in the task definition. In bridge mode, the default value is the ephemeral port of the Service Connect proxy.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "port_name": {
                    "computed": true,
                    "description": "The ` + "`" + `` + "`" + `portName` + "`" + `` + "`" + ` must match the name of one of the ` + "`" + `` + "`" + `portMappings` + "`" + `` + "`" + ` from all the containers in the task definition of this Amazon ECS service.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "timeout": {
                    "computed": true,
                    "description": "A reference to an object that represents the configured timeouts for Service Connect.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "idle_timeout_seconds": {
                          "computed": true,
                          "description": "The amount of time in seconds a connection will stay active while idle. A value of ` + "`" + `` + "`" + `0` + "`" + `` + "`" + ` can be set to disable ` + "`" + `` + "`" + `idleTimeout` + "`" + `` + "`" + `.\n The ` + "`" + `` + "`" + `idleTimeout` + "`" + `` + "`" + ` default for ` + "`" + `` + "`" + `HTTP` + "`" + `` + "`" + `/` + "`" + `` + "`" + `HTTP2` + "`" + `` + "`" + `/` + "`" + `` + "`" + `GRPC` + "`" + `` + "`" + ` is 5 minutes.\n The ` + "`" + `` + "`" + `idleTimeout` + "`" + `` + "`" + ` default for ` + "`" + `` + "`" + `TCP` + "`" + `` + "`" + ` is 1 hour.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "per_request_timeout_seconds": {
                          "computed": true,
                          "description": "The amount of time waiting for the upstream to respond with a complete response per request. A value of ` + "`" + `` + "`" + `0` + "`" + `` + "`" + ` can be set to disable ` + "`" + `` + "`" + `perRequestTimeout` + "`" + `` + "`" + `. ` + "`" + `` + "`" + `perRequestTimeout` + "`" + `` + "`" + ` can only be set if Service Connect ` + "`" + `` + "`" + `appProtocol` + "`" + `` + "`" + ` isn't ` + "`" + `` + "`" + `TCP` + "`" + `` + "`" + `. Only ` + "`" + `` + "`" + `idleTimeout` + "`" + `` + "`" + ` is allowed for ` + "`" + `` + "`" + `TCP` + "`" + `` + "`" + ` ` + "`" + `` + "`" + `appProtocol` + "`" + `` + "`" + `.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "tls": {
                    "computed": true,
                    "description": "A reference to an object that represents a Transport Layer Security (TLS) configuration.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "issuer_certificate_authority": {
                          "computed": true,
                          "description": "The signer certificate authority.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "aws_pca_authority_arn": {
                                "computed": true,
                                "description": "The ARN of the AWS Private Certificate Authority certificate.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "kms_key": {
                          "computed": true,
                          "description": "The AWS Key Management Service key.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "role_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the IAM role that's associated with the Service Connect TLS.",
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
                "nesting_mode": "list"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "service_name": {
        "computed": true,
        "description": "The name of your service. Up to 255 letters (uppercase and lowercase), numbers, underscores, and hyphens are allowed. Service names must be unique within a cluster, but you can have similarly named services in multiple clusters within a Region or across multiple Regions.\n  The stack update fails if you change any properties that require replacement and the ` + "`" + `` + "`" + `ServiceName` + "`" + `` + "`" + ` is configured. This is because AWS CloudFormation creates the replacement service first, but each ` + "`" + `` + "`" + `ServiceName` + "`" + `` + "`" + ` must be unique in the cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_registries": {
        "computed": true,
        "description": "The details of the service discovery registry to associate with this service. For more information, see [Service discovery](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-discovery.html).\n  Each service may be associated with one service registry. Multiple service registries for each service isn't supported.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "container_name": {
              "computed": true,
              "description": "The container name value to be used for your service discovery service. It's already specified in the task definition. If the task definition that your service task specifies uses the ` + "`" + `` + "`" + `bridge` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `host` + "`" + `` + "`" + ` network mode, you must specify a ` + "`" + `` + "`" + `containerName` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `containerPort` + "`" + `` + "`" + ` combination from the task definition. If the task definition that your service task specifies uses the ` + "`" + `` + "`" + `awsvpc` + "`" + `` + "`" + ` network mode and a type SRV DNS record is used, you must specify either a ` + "`" + `` + "`" + `containerName` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `containerPort` + "`" + `` + "`" + ` combination or a ` + "`" + `` + "`" + `port` + "`" + `` + "`" + ` value. However, you can't specify both.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "container_port": {
              "computed": true,
              "description": "The port value to be used for your service discovery service. It's already specified in the task definition. If the task definition your service task specifies uses the ` + "`" + `` + "`" + `bridge` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `host` + "`" + `` + "`" + ` network mode, you must specify a ` + "`" + `` + "`" + `containerName` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `containerPort` + "`" + `` + "`" + ` combination from the task definition. If the task definition your service task specifies uses the ` + "`" + `` + "`" + `awsvpc` + "`" + `` + "`" + ` network mode and a type SRV DNS record is used, you must specify either a ` + "`" + `` + "`" + `containerName` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `containerPort` + "`" + `` + "`" + ` combination or a ` + "`" + `` + "`" + `port` + "`" + `` + "`" + ` value. However, you can't specify both.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "port": {
              "computed": true,
              "description": "The port value used if your service discovery service specified an SRV record. This field might be used if both the ` + "`" + `` + "`" + `awsvpc` + "`" + `` + "`" + ` network mode and SRV records are used.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "registry_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the service registry. The currently supported service registry is CMAP. For more information, see [CreateService](https://docs.aws.amazon.com/cloud-map/latest/api/API_CreateService.html).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "tags": {
        "computed": true,
        "description": "The metadata that you apply to the service to help you categorize and organize them. Each tag consists of a key and an optional value, both of which you define. When a service is deleted, the tags are deleted as well.\n The following basic restrictions apply to tags:\n  +  Maximum number of tags per resource - 50\n  +  For each resource, each tag key must be unique, and each tag key can have only one value.\n  +  Maximum key length - 128 Unicode characters in UTF-8\n  +  Maximum value length - 256 Unicode characters in UTF-8\n  +  If your tagging schema is used across multiple services and resources, remember that other services may have restrictions on allowed characters. Generally allowed characters are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . _ : / @.\n  +  Tag keys and values are case-sensitive.\n  +  Do not use ` + "`" + `` + "`" + `aws:` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `AWS:` + "`" + `` + "`" + `, or any upper or lowercase combination of such as a prefix for either keys or values as it is reserved for AWS use. You cannot edit or delete tag keys or values with this prefix. Tags with this prefix do not count against your tags per resource limit.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "One part of a key-value pair that make up a tag. A ` + "`" + `` + "`" + `key` + "`" + `` + "`" + ` is a general label that acts like a category for more specific tag values.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The optional part of a key-value pair that make up a tag. A ` + "`" + `` + "`" + `value` + "`" + `` + "`" + ` acts as a descriptor within a tag category (key).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "task_definition": {
        "computed": true,
        "description": "The ` + "`" + `` + "`" + `family` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `revision` + "`" + `` + "`" + ` (` + "`" + `` + "`" + `family:revision` + "`" + `` + "`" + `) or full ARN of the task definition to run in your service. If a ` + "`" + `` + "`" + `revision` + "`" + `` + "`" + ` isn't specified, the latest ` + "`" + `` + "`" + `ACTIVE` + "`" + `` + "`" + ` revision is used.\n A task definition must be specified if the service uses either the ` + "`" + `` + "`" + `ECS` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `CODE_DEPLOY` + "`" + `` + "`" + ` deployment controllers.\n For more information about deployment types, see [Amazon ECS deployment types](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "volume_configurations": {
        "computed": true,
        "description": "The configuration for a volume specified in the task definition as a volume that is configured at launch time. Currently, the only supported volume type is an Amazon EBS volume.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "managed_ebs_volume": {
              "computed": true,
              "description": "The configuration for the Amazon EBS volume that Amazon ECS creates and manages on your behalf. These settings are used to create each Amazon EBS volume, with one volume created for each task in the service. The Amazon EBS volumes are visible in your account in the Amazon EC2 console once they are created.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "encrypted": {
                    "computed": true,
                    "description": "Indicates whether the volume should be encrypted. If no value is specified, encryption is turned on by default. This parameter maps 1:1 with the ` + "`" + `` + "`" + `Encrypted` + "`" + `` + "`" + ` parameter of the [CreateVolume API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateVolume.html) in the *Amazon EC2 API Reference*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "filesystem_type": {
                    "computed": true,
                    "description": "The Linux filesystem type for the volume. For volumes created from a snapshot, you must specify the same filesystem type that the volume was using when the snapshot was created. If there is a filesystem type mismatch, the task will fail to start.\n The available filesystem types are\u2028 ` + "`" + `` + "`" + `ext3` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `ext4` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `xfs` + "`" + `` + "`" + `. If no value is specified, the ` + "`" + `` + "`" + `xfs` + "`" + `` + "`" + ` filesystem type is used by default.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "iops": {
                    "computed": true,
                    "description": "The number of I/O operations per second (IOPS). For ` + "`" + `` + "`" + `gp3` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `io1` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `io2` + "`" + `` + "`" + ` volumes, this represents the number of IOPS that are provisioned for the volume. For ` + "`" + `` + "`" + `gp2` + "`" + `` + "`" + ` volumes, this represents the baseline performance of the volume and the rate at which the volume accumulates I/O credits for bursting.\n The following are the supported values for each volume type.\n  +   ` + "`" + `` + "`" + `gp3` + "`" + `` + "`" + `: 3,000 - 16,000 IOPS\n  +   ` + "`" + `` + "`" + `io1` + "`" + `` + "`" + `: 100 - 64,000 IOPS\n  +   ` + "`" + `` + "`" + `io2` + "`" + `` + "`" + `: 100 - 256,000 IOPS\n  \n This parameter is required for ` + "`" + `` + "`" + `io1` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `io2` + "`" + `` + "`" + ` volume types. The default for ` + "`" + `` + "`" + `gp3` + "`" + `` + "`" + ` volumes is ` + "`" + `` + "`" + `3,000 IOPS` + "`" + `` + "`" + `. This parameter is not supported for ` + "`" + `` + "`" + `st1` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `sc1` + "`" + `` + "`" + `, or ` + "`" + `` + "`" + `standard` + "`" + `` + "`" + ` volume types.\n This parameter maps 1:1 with the ` + "`" + `` + "`" + `Iops` + "`" + `` + "`" + ` parameter of the [CreateVolume API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateVolume.html) in the *Amazon EC2 API Reference*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "kms_key_id": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) identifier of the AWS Key Management Service key to use for Amazon EBS encryption. When encryption is turned on and no AWS Key Management Service key is specified, the default AWS managed key for Amazon EBS volumes is used. This parameter maps 1:1 with the ` + "`" + `` + "`" + `KmsKeyId` + "`" + `` + "`" + ` parameter of the [CreateVolume API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateVolume.html) in the *Amazon EC2 API Reference*.\n   AWS authenticates the AWS Key Management Service key asynchronously. Therefore, if you specify an ID, alias, or ARN that is invalid, the action can appear to complete, but eventually fails.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "role_arn": {
                    "computed": true,
                    "description": "The ARN of the IAM role to associate with this volume. This is the Amazon ECS infrastructure IAM role that is used to manage your AWS infrastructure. We recommend using the Amazon ECS-managed ` + "`" + `` + "`" + `AmazonECSInfrastructureRolePolicyForVolumes` + "`" + `` + "`" + ` IAM policy with this role. For more information, see [Amazon ECS infrastructure IAM role](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/infrastructure_IAM_role.html) in the *Amazon ECS Developer Guide*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "size_in_gi_b": {
                    "computed": true,
                    "description": "The size of the volume in GiB. You must specify either a volume size or a snapshot ID. If you specify a snapshot ID, the snapshot size is used for the volume size by default. You can optionally specify a volume size greater than or equal to the snapshot size. This parameter maps 1:1 with the ` + "`" + `` + "`" + `Size` + "`" + `` + "`" + ` parameter of the [CreateVolume API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateVolume.html) in the *Amazon EC2 API Reference*.\n The following are the supported volume size values for each volume type.\n  +   ` + "`" + `` + "`" + `gp2` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `gp3` + "`" + `` + "`" + `: 1-16,384\n  +   ` + "`" + `` + "`" + `io1` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `io2` + "`" + `` + "`" + `: 4-16,384\n  +   ` + "`" + `` + "`" + `st1` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `sc1` + "`" + `` + "`" + `: 125-16,384\n  +   ` + "`" + `` + "`" + `standard` + "`" + `` + "`" + `: 1-1,024",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "snapshot_id": {
                    "computed": true,
                    "description": "The snapshot that Amazon ECS uses to create the volume. You must specify either a snapshot ID or a volume size. This parameter maps 1:1 with the ` + "`" + `` + "`" + `SnapshotId` + "`" + `` + "`" + ` parameter of the [CreateVolume API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateVolume.html) in the *Amazon EC2 API Reference*.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "tag_specifications": {
                    "computed": true,
                    "description": "The tags to apply to the volume. Amazon ECS applies service-managed tags by default. This parameter maps 1:1 with the ` + "`" + `` + "`" + `TagSpecifications.N` + "`" + `` + "`" + ` parameter of the [CreateVolume API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateVolume.html) in the *Amazon EC2 API Reference*.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "propagate_tags": {
                          "computed": true,
                          "description": "Determines whether to propagate the tags from the task definition to \u2028the Amazon EBS volume. Tags can only propagate to a ` + "`" + `` + "`" + `SERVICE` + "`" + `` + "`" + ` specified in \u2028` + "`" + `` + "`" + `ServiceVolumeConfiguration` + "`" + `` + "`" + `. If no value is specified, the tags aren't \u2028propagated.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "resource_type": {
                          "computed": true,
                          "description": "The type of volume resource.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "tags": {
                          "computed": true,
                          "description": "The tags applied to this Amazon EBS volume. ` + "`" + `` + "`" + `AmazonECSCreated` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `AmazonECSManaged` + "`" + `` + "`" + ` are reserved tags that can't be used.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "key": {
                                "computed": true,
                                "description": "One part of a key-value pair that make up a tag. A ` + "`" + `` + "`" + `key` + "`" + `` + "`" + ` is a general label that acts like a category for more specific tag values.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "value": {
                                "computed": true,
                                "description": "The optional part of a key-value pair that make up a tag. A ` + "`" + `` + "`" + `value` + "`" + `` + "`" + ` acts as a descriptor within a tag category (key).",
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
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "throughput": {
                    "computed": true,
                    "description": "The throughput to provision for a volume, in MiB/s, with a maximum of 1,000 MiB/s. This parameter maps 1:1 with the ` + "`" + `` + "`" + `Throughput` + "`" + `` + "`" + ` parameter of the [CreateVolume API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateVolume.html) in the *Amazon EC2 API Reference*.\n  This parameter is only supported for the ` + "`" + `` + "`" + `gp3` + "`" + `` + "`" + ` volume type.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "volume_type": {
                    "computed": true,
                    "description": "The volume type. This parameter maps 1:1 with the ` + "`" + `` + "`" + `VolumeType` + "`" + `` + "`" + ` parameter of the [CreateVolume API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateVolume.html) in the *Amazon EC2 API Reference*. For more information, see [Amazon EBS volume types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-volume-types.html) in the *Amazon EC2 User Guide*.\n The following are the supported volume types.\n  +  General Purpose SSD: ` + "`" + `` + "`" + `gp2` + "`" + `` + "`" + `|` + "`" + `` + "`" + `gp3` + "`" + `` + "`" + ` \n  +  Provisioned IOPS SSD: ` + "`" + `` + "`" + `io1` + "`" + `` + "`" + `|` + "`" + `` + "`" + `io2` + "`" + `` + "`" + ` \n  +  Throughput Optimized HDD: ` + "`" + `` + "`" + `st1` + "`" + `` + "`" + ` \n  +  Cold HDD: ` + "`" + `` + "`" + `sc1` + "`" + `` + "`" + ` \n  +  Magnetic: ` + "`" + `` + "`" + `standard` + "`" + `` + "`" + ` \n  The magnetic volume type is not supported on Fargate.",
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
              "description": "The name of the volume. This value must match the volume name from the ` + "`" + `` + "`" + `Volume` + "`" + `` + "`" + ` object in the task definition.",
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
    "description": "The ` + "`" + `` + "`" + `AWS::ECS::Service` + "`" + `` + "`" + ` resource creates an Amazon Elastic Container Service (Amazon ECS) service that runs and maintains the requested number of tasks and associated load balancers.\n  The stack update fails if you change any properties that require replacement and at least one ECS Service Connect ` + "`" + `` + "`" + `ServiceConnectConfiguration` + "`" + `` + "`" + ` property the is configured. This is because AWS CloudFormation creates the replacement service first, but each ` + "`" + `` + "`" + `ServiceConnectService` + "`" + `` + "`" + ` must have a name that is unique in the namespace.\n   Starting April 15, 2023, AWS; will not onboard new customers to Amazon Elastic Inference (EI), and will help current customers migrate their workloads to options that offer better price and performance. After April 15, 2023, new customers will not be able to launch instances with Amazon EI accelerators in Amazon SageMaker, ECS, or EC2. However, customers who have used Amazon EI at least once during the past 30-day period are considered current customers and will be able to continue using the service.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEcsServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEcsService), &result)
	return &result
}
