package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGameliftContainerFleet = `{
  "block": {
    "attributes": {
      "billing_type": {
        "computed": true,
        "description": "Indicates whether to use On-Demand instances or Spot instances for this fleet. If empty, the default is ON_DEMAND. Both categories of instances use identical hardware and configurations based on the instance type selected for this fleet.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description": "A time stamp indicating when this data object was created. Format is a number expressed in Unix time as milliseconds (for example \"1469498468.057\").",
        "description_kind": "plain",
        "type": "string"
      },
      "deployment_configuration": {
        "computed": true,
        "description": "Provides details about how to drain old tasks and replace them with new updated tasks.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "impairment_strategy": {
              "computed": true,
              "description": "The strategy to apply in case of impairment; defaults to MAINTAIN.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "minimum_healthy_percentage": {
              "computed": true,
              "description": "The minimum percentage of healthy required; defaults to 75.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "protection_strategy": {
              "computed": true,
              "description": "The protection strategy for deployment on the container fleet; defaults to WITH_PROTECTION.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "deployment_details": {
        "computed": true,
        "description": "Provides information about the last deployment ID and its status.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "latest_deployment_id": {
              "computed": true,
              "description": "The ID of the last deployment on the container fleet. This field will be empty if the container fleet does not have a ContainerGroupDefinition attached.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "description": {
        "computed": true,
        "description": "A human-readable description of a fleet.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "fleet_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) that is assigned to a Amazon GameLift container fleet resource and uniquely identifies it across all AWS Regions.",
        "description_kind": "plain",
        "type": "string"
      },
      "fleet_id": {
        "computed": true,
        "description": "Unique fleet ID",
        "description_kind": "plain",
        "type": "string"
      },
      "fleet_role_arn": {
        "description": "A unique identifier for an AWS IAM role that manages access to your AWS services. Create a role or look up a role's ARN from the IAM dashboard in the AWS Management Console.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "game_server_container_group_definition_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the game server container group definition. This field will be empty if GameServerContainerGroupDefinitionName is not specified.",
        "description_kind": "plain",
        "type": "string"
      },
      "game_server_container_group_definition_name": {
        "computed": true,
        "description": "The name of the container group definition that will be created per game server. You must specify GAME_SERVER container group. You have the option to also specify one PER_INSTANCE container group.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "game_server_container_groups_per_instance": {
        "computed": true,
        "description": "The number of desired game server container groups per instance, a number between 1-5000.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "game_session_creation_limit_policy": {
        "computed": true,
        "description": "A policy that limits the number of game sessions an individual player can create over a span of time for this fleet.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "new_game_sessions_per_creator": {
              "computed": true,
              "description": "The maximum number of game sessions that an individual can create during the policy period.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "policy_period_in_minutes": {
              "computed": true,
              "description": "The time span used in evaluating the resource creation limit policy.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_connection_port_range": {
        "computed": true,
        "description": "Defines the range of ports on the instance that allow inbound traffic to connect with containers in a fleet.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "from_port": {
              "computed": true,
              "description": "A starting value for a range of allowed port numbers.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "to_port": {
              "computed": true,
              "description": "An ending value for a range of allowed port numbers. Port numbers are end-inclusive. This value must be higher than FromPort.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "instance_inbound_permissions": {
        "computed": true,
        "description": "A range of IP addresses and port settings that allow inbound traffic to connect to server processes on an Amazon GameLift server.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "from_port": {
              "computed": true,
              "description": "A starting value for a range of allowed port numbers.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "ip_range": {
              "computed": true,
              "description": "A range of allowed IP addresses. This value must be expressed in CIDR notation. Example: \"000.000.000.000/[subnet mask]\" or optionally the shortened version \"0.0.0.0/[subnet mask]\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "protocol": {
              "computed": true,
              "description": "The network communication protocol used by the fleet.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "to_port": {
              "computed": true,
              "description": "An ending value for a range of allowed port numbers. Port numbers are end-inclusive. This value must be higher than FromPort.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "instance_type": {
        "computed": true,
        "description": "The name of an EC2 instance type that is supported in Amazon GameLift. A fleet instance type determines the computing resources of each instance in the fleet, including CPU, memory, storage, and networking capacity. Amazon GameLift supports the following EC2 instance types. See Amazon EC2 Instance Types for detailed descriptions.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "locations": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "location": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "location_capacity": {
              "computed": true,
              "description": "Current resource capacity settings in a specified fleet or location. The location value might refer to a fleet's remote location or its home Region.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "desired_ec2_instances": {
                    "computed": true,
                    "description": "Defaults to MinSize if not defined. The number of EC2 instances you want to maintain in the specified fleet location. This value must fall between the minimum and maximum size limits. If any auto-scaling policy is defined for the container fleet, the desired instance will only be applied once during fleet creation and will be ignored in updates to avoid conflicts with auto-scaling. During updates with any auto-scaling policy defined, if current desired instance is lower than the new MinSize, it will be increased to the new MinSize; if current desired instance is larger than the new MaxSize, it will be decreased to the new MaxSize.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "max_size": {
                    "computed": true,
                    "description": "The maximum value that is allowed for the fleet's instance count for a location.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "min_size": {
                    "computed": true,
                    "description": "The minimum value allowed for the fleet's instance count for a location.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "stopped_actions": {
              "computed": true,
              "description": "A list of fleet actions that have been suspended in the fleet location.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "log_configuration": {
        "computed": true,
        "description": "A policy the location and provider of logs from the fleet.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "log_destination": {
              "computed": true,
              "description": "Configures the service that provides logs.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "log_group_arn": {
              "computed": true,
              "description": "If log destination is CLOUDWATCH, logs are sent to the specified log group in Amazon CloudWatch.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "s3_bucket_name": {
              "computed": true,
              "description": "The name of the S3 bucket to pull logs from if S3 is the LogDestination",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "maximum_game_server_container_groups_per_instance": {
        "computed": true,
        "description": "The maximum number of game server container groups per instance, a number between 1-5000.",
        "description_kind": "plain",
        "type": "number"
      },
      "metric_groups": {
        "computed": true,
        "description": "The name of an Amazon CloudWatch metric group. A metric group aggregates the metrics for all fleets in the group. Specify a string containing the metric group name. You can use an existing name or use a new name to create a new metric group. Currently, this parameter can have only one string.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "new_game_session_protection_policy": {
        "computed": true,
        "description": "A game session protection policy to apply to all game sessions hosted on instances in this fleet. When protected, active game sessions cannot be terminated during a scale-down event. If this parameter is not set, instances in this fleet default to no protection. You can change a fleet's protection policy to affect future game sessions on the fleet. You can also set protection for individual game sessions.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "per_instance_container_group_definition_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the per instance container group definition. This field will be empty if PerInstanceContainerGroupDefinitionName is not specified.",
        "description_kind": "plain",
        "type": "string"
      },
      "per_instance_container_group_definition_name": {
        "computed": true,
        "description": "The name of the container group definition that will be created per instance. This field is optional if you specify GameServerContainerGroupDefinitionName.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "scaling_policies": {
        "computed": true,
        "description": "A list of rules that control how a fleet is scaled.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "comparison_operator": {
              "computed": true,
              "description": "Comparison operator to use when measuring a metric against the threshold value.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "evaluation_periods": {
              "computed": true,
              "description": "Length of time (in minutes) the metric must be at or beyond the threshold before a scaling event is triggered.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "metric_name": {
              "computed": true,
              "description": "Name of the Amazon GameLift-defined metric that is used to trigger a scaling adjustment.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "A descriptive label that is associated with a fleet's scaling policy. Policy names do not need to be unique.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "policy_type": {
              "computed": true,
              "description": "The type of scaling policy to create. For a target-based policy, set the parameter MetricName to 'PercentAvailableGameSessions' and specify a TargetConfiguration. For a rule-based policy set the following parameters: MetricName, ComparisonOperator, Threshold, EvaluationPeriods, ScalingAdjustmentType, and ScalingAdjustment.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "scaling_adjustment": {
              "computed": true,
              "description": "Amount of adjustment to make, based on the scaling adjustment type.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "scaling_adjustment_type": {
              "computed": true,
              "description": "The type of adjustment to make to a fleet's instance count.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "target_configuration": {
              "computed": true,
              "description": "An object that contains settings for a target-based scaling policy.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "target_value": {
                    "computed": true,
                    "description": "Desired value to use with a target-based scaling policy. The value must be relevant for whatever metric the scaling policy is using. For example, in a policy using the metric PercentAvailableGameSessions, the target value should be the preferred size of the fleet's buffer (the percent of capacity that should be idle and ready for new game sessions).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "threshold": {
              "computed": true,
              "description": "Metric value used to trigger a scaling event.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "status": {
        "computed": true,
        "description": "The current status of the container fleet.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "The AWS::GameLift::ContainerFleet resource creates an Amazon GameLift (GameLift) container fleet to host game servers.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccGameliftContainerFleetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGameliftContainerFleet), &result)
	return &result
}
