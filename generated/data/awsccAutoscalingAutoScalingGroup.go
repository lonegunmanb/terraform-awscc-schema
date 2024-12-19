package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAutoscalingAutoScalingGroup = `{
  "block": {
    "attributes": {
      "auto_scaling_group_name": {
        "computed": true,
        "description": "The name of the Auto Scaling group. This name must be unique per Region per account.\n The name can contain any ASCII character 33 to 126 including most punctuation characters, digits, and upper and lowercased letters.\n  You cannot use a colon (:) in the name.",
        "description_kind": "plain",
        "type": "string"
      },
      "availability_zone_distribution": {
        "computed": true,
        "description": "The instance capacity distribution across Availability Zones.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "capacity_distribution_strategy": {
              "computed": true,
              "description": "If launches fail in an Availability Zone, the following strategies are available. The default is ` + "`" + `` + "`" + `balanced-best-effort` + "`" + `` + "`" + `. \n  +   ` + "`" + `` + "`" + `balanced-only` + "`" + `` + "`" + ` - If launches fail in an Availability Zone, Auto Scaling will continue to attempt to launch in the unhealthy zone to preserve a balanced distribution.\n  +   ` + "`" + `` + "`" + `balanced-best-effort` + "`" + `` + "`" + ` - If launches fail in an Availability Zone, Auto Scaling will attempt to launch in another healthy Availability Zone instead.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "availability_zone_impairment_policy": {
        "computed": true,
        "description": "The Availability Zone impairment policy.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "impaired_zone_health_check_behavior": {
              "computed": true,
              "description": "Specifies the health check behavior for the impaired Availability Zone in an active zonal shift. If you select ` + "`" + `` + "`" + `Replace unhealthy` + "`" + `` + "`" + `, instances that appear unhealthy will be replaced in all Availability Zones. If you select ` + "`" + `` + "`" + `Ignore unhealthy` + "`" + `` + "`" + `, instances will not be replaced in the Availability Zone with the active zonal shift. For more information, see [Auto Scaling group zonal shift](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-zonal-shift.html) in the *Amazon EC2 Auto Scaling User Guide*.",
              "description_kind": "plain",
              "type": "string"
            },
            "zonal_shift_enabled": {
              "computed": true,
              "description": "If ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `, enable zonal shift for your Auto Scaling group.",
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "availability_zones": {
        "computed": true,
        "description": "A list of Availability Zones where instances in the Auto Scaling group can be created. Used for launching into the default VPC subnet in each Availability Zone when not using the ` + "`" + `` + "`" + `VPCZoneIdentifier` + "`" + `` + "`" + ` property, or for attaching a network interface when an existing network interface ID is specified in a launch template.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "capacity_rebalance": {
        "computed": true,
        "description": "Indicates whether Capacity Rebalancing is enabled. Otherwise, Capacity Rebalancing is disabled. When you turn on Capacity Rebalancing, Amazon EC2 Auto Scaling attempts to launch a Spot Instance whenever Amazon EC2 notifies that a Spot Instance is at an elevated risk of interruption. After launching a new instance, it then terminates an old instance. For more information, see [Use Capacity Rebalancing to handle Amazon EC2 Spot Interruptions](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-capacity-rebalancing.html) in the in the *Amazon EC2 Auto Scaling User Guide*.",
        "description_kind": "plain",
        "type": "bool"
      },
      "capacity_reservation_specification": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "capacity_reservation_preference": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "capacity_reservation_target": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "capacity_reservation_ids": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "capacity_reservation_resource_group_arns": {
                    "computed": true,
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
          "nesting_mode": "single"
        }
      },
      "context": {
        "computed": true,
        "description": "Reserved.",
        "description_kind": "plain",
        "type": "string"
      },
      "cooldown": {
        "computed": true,
        "description": "*Only needed if you use simple scaling policies.* \n The amount of time, in seconds, between one scaling activity ending and another one starting due to simple scaling policies. For more information, see [Scaling cooldowns for Amazon EC2 Auto Scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-scaling-cooldowns.html) in the *Amazon EC2 Auto Scaling User Guide*.\n Default: ` + "`" + `` + "`" + `300` + "`" + `` + "`" + ` seconds",
        "description_kind": "plain",
        "type": "string"
      },
      "default_instance_warmup": {
        "computed": true,
        "description": "The amount of time, in seconds, until a new instance is considered to have finished initializing and resource consumption to become stable after it enters the ` + "`" + `` + "`" + `InService` + "`" + `` + "`" + ` state. \n During an instance refresh, Amazon EC2 Auto Scaling waits for the warm-up period after it replaces an instance before it moves on to replacing the next instance. Amazon EC2 Auto Scaling also waits for the warm-up period before aggregating the metrics for new instances with existing instances in the Amazon CloudWatch metrics that are used for scaling, resulting in more reliable usage data. For more information, see [Set the default instance warmup for an Auto Scaling group](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-default-instance-warmup.html) in the *Amazon EC2 Auto Scaling User Guide*.\n  To manage various warm-up settings at the group level, we recommend that you set the default instance warmup, *even if it is set to 0 seconds*. To remove a value that you previously set, include the property but specify ` + "`" + `` + "`" + `-1` + "`" + `` + "`" + ` for the value. However, we strongly recommend keeping the default instance warmup enabled by specifying a value of ` + "`" + `` + "`" + `0` + "`" + `` + "`" + ` or other nominal value.\n  Default: None",
        "description_kind": "plain",
        "type": "number"
      },
      "desired_capacity": {
        "computed": true,
        "description": "The desired capacity is the initial capacity of the Auto Scaling group at the time of its creation and the capacity it attempts to maintain. It can scale beyond this capacity if you configure automatic scaling.\n The number must be greater than or equal to the minimum size of the group and less than or equal to the maximum size of the group. If you do not specify a desired capacity when creating the stack, the default is the minimum size of the group.\n CloudFormation marks the Auto Scaling group as successful (by setting its status to CREATE_COMPLETE) when the desired capacity is reached. However, if a maximum Spot price is set in the launch template or launch configuration that you specified, then desired capacity is not used as a criteria for success. Whether your request is fulfilled depends on Spot Instance capacity and your maximum price.",
        "description_kind": "plain",
        "type": "string"
      },
      "desired_capacity_type": {
        "computed": true,
        "description": "The unit of measurement for the value specified for desired capacity. Amazon EC2 Auto Scaling supports ` + "`" + `` + "`" + `DesiredCapacityType` + "`" + `` + "`" + ` for attribute-based instance type selection only. For more information, see [Create a mixed instances group using attribute-based instance type selection](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-mixed-instances-group-attribute-based-instance-type-selection.html) in the *Amazon EC2 Auto Scaling User Guide*.\n By default, Amazon EC2 Auto Scaling specifies ` + "`" + `` + "`" + `units` + "`" + `` + "`" + `, which translates into number of instances.\n Valid values: ` + "`" + `` + "`" + `units` + "`" + `` + "`" + ` | ` + "`" + `` + "`" + `vcpu` + "`" + `` + "`" + ` | ` + "`" + `` + "`" + `memory-mib` + "`" + `` + "`" + `",
        "description_kind": "plain",
        "type": "string"
      },
      "health_check_grace_period": {
        "computed": true,
        "description": "The amount of time, in seconds, that Amazon EC2 Auto Scaling waits before checking the health status of an EC2 instance that has come into service and marking it unhealthy due to a failed health check. This is useful if your instances do not immediately pass their health checks after they enter the ` + "`" + `` + "`" + `InService` + "`" + `` + "`" + ` state. For more information, see [Set the health check grace period for an Auto Scaling group](https://docs.aws.amazon.com/autoscaling/ec2/userguide/health-check-grace-period.html) in the *Amazon EC2 Auto Scaling User Guide*.\n Default: ` + "`" + `` + "`" + `0` + "`" + `` + "`" + ` seconds",
        "description_kind": "plain",
        "type": "number"
      },
      "health_check_type": {
        "computed": true,
        "description": "A comma-separated value string of one or more health check types.\n The valid values are ` + "`" + `` + "`" + `EC2` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `EBS` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `ELB` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `VPC_LATTICE` + "`" + `` + "`" + `. ` + "`" + `` + "`" + `EC2` + "`" + `` + "`" + ` is the default health check and cannot be disabled. For more information, see [Health checks for instances in an Auto Scaling group](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-health-checks.html) in the *Amazon EC2 Auto Scaling User Guide*.\n Only specify ` + "`" + `` + "`" + `EC2` + "`" + `` + "`" + ` if you must clear a value that was previously set.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_id": {
        "computed": true,
        "description": "The ID of the instance used to base the launch configuration on. For more information, see [Create an Auto Scaling group using an EC2 instance](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-asg-from-instance.html) in the *Amazon EC2 Auto Scaling User Guide*.\n If you specify ` + "`" + `` + "`" + `LaunchTemplate` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `MixedInstancesPolicy` + "`" + `` + "`" + `, or ` + "`" + `` + "`" + `LaunchConfigurationName` + "`" + `` + "`" + `, don't specify ` + "`" + `` + "`" + `InstanceId` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_maintenance_policy": {
        "computed": true,
        "description": "An instance maintenance policy. For more information, see [Set instance maintenance policy](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-instance-maintenance-policy.html) in the *Amazon EC2 Auto Scaling User Guide*.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "max_healthy_percentage": {
              "computed": true,
              "description": "Specifies the upper threshold as a percentage of the desired capacity of the Auto Scaling group. It represents the maximum percentage of the group that can be in service and healthy, or pending, to support your workload when replacing instances. Value range is 100 to 200. To clear a previously set value, specify a value of ` + "`" + `` + "`" + `-1` + "`" + `` + "`" + `.\n Both ` + "`" + `` + "`" + `MinHealthyPercentage` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `MaxHealthyPercentage` + "`" + `` + "`" + ` must be specified, and the difference between them cannot be greater than 100. A large range increases the number of instances that can be replaced at the same time.",
              "description_kind": "plain",
              "type": "number"
            },
            "min_healthy_percentage": {
              "computed": true,
              "description": "Specifies the lower threshold as a percentage of the desired capacity of the Auto Scaling group. It represents the minimum percentage of the group to keep in service, healthy, and ready to use to support your workload when replacing instances. Value range is 0 to 100. To clear a previously set value, specify a value of ` + "`" + `` + "`" + `-1` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "launch_configuration_name": {
        "computed": true,
        "description": "The name of the launch configuration to use to launch instances.\n Required only if you don't specify ` + "`" + `` + "`" + `LaunchTemplate` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `MixedInstancesPolicy` + "`" + `` + "`" + `, or ` + "`" + `` + "`" + `InstanceId` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "type": "string"
      },
      "launch_template": {
        "computed": true,
        "description": "Information used to specify the launch template and version to use to launch instances. You can alternatively associate a launch template to the Auto Scaling group by specifying a ` + "`" + `` + "`" + `MixedInstancesPolicy` + "`" + `` + "`" + `. For more information about creating launch templates, see [Create a launch template for an Auto Scaling group](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-launch-template.html) in the *Amazon EC2 Auto Scaling User Guide*.\n If you omit this property, you must specify ` + "`" + `` + "`" + `MixedInstancesPolicy` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `LaunchConfigurationName` + "`" + `` + "`" + `, or ` + "`" + `` + "`" + `InstanceId` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "launch_template_id": {
              "computed": true,
              "description": "The ID of the launch template.\n You must specify the ` + "`" + `` + "`" + `LaunchTemplateID` + "`" + `` + "`" + ` or the ` + "`" + `` + "`" + `LaunchTemplateName` + "`" + `` + "`" + `, but not both.",
              "description_kind": "plain",
              "type": "string"
            },
            "launch_template_name": {
              "computed": true,
              "description": "The name of the launch template.\n You must specify the ` + "`" + `` + "`" + `LaunchTemplateName` + "`" + `` + "`" + ` or the ` + "`" + `` + "`" + `LaunchTemplateID` + "`" + `` + "`" + `, but not both.",
              "description_kind": "plain",
              "type": "string"
            },
            "version": {
              "computed": true,
              "description": "The version number of the launch template.\n Specifying ` + "`" + `` + "`" + `$Latest` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `$Default` + "`" + `` + "`" + ` for the template version number is not supported. However, you can specify ` + "`" + `` + "`" + `LatestVersionNumber` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `DefaultVersionNumber` + "`" + `` + "`" + ` using the ` + "`" + `` + "`" + `Fn::GetAtt` + "`" + `` + "`" + ` intrinsic function. For more information, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).\n  For an example of using the ` + "`" + `` + "`" + `Fn::GetAtt` + "`" + `` + "`" + ` function, see the [Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html#aws-resource-autoscaling-autoscalinggroup--examples) section of the ` + "`" + `` + "`" + `AWS::AutoScaling::AutoScalingGroup` + "`" + `` + "`" + ` resource.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "lifecycle_hook_specification_list": {
        "computed": true,
        "description": "One or more lifecycle hooks to add to the Auto Scaling group before instances are launched.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "default_result": {
              "computed": true,
              "description": "The action the Auto Scaling group takes when the lifecycle hook timeout elapses or if an unexpected failure occurs. The default value is ` + "`" + `` + "`" + `ABANDON` + "`" + `` + "`" + `.\n Valid values: ` + "`" + `` + "`" + `CONTINUE` + "`" + `` + "`" + ` | ` + "`" + `` + "`" + `ABANDON` + "`" + `` + "`" + `",
              "description_kind": "plain",
              "type": "string"
            },
            "heartbeat_timeout": {
              "computed": true,
              "description": "The maximum time, in seconds, that can elapse before the lifecycle hook times out. The range is from ` + "`" + `` + "`" + `30` + "`" + `` + "`" + ` to ` + "`" + `` + "`" + `7200` + "`" + `` + "`" + ` seconds. The default value is ` + "`" + `` + "`" + `3600` + "`" + `` + "`" + ` seconds (1 hour).",
              "description_kind": "plain",
              "type": "number"
            },
            "lifecycle_hook_name": {
              "computed": true,
              "description": "The name of the lifecycle hook.",
              "description_kind": "plain",
              "type": "string"
            },
            "lifecycle_transition": {
              "computed": true,
              "description": "The lifecycle transition. For Auto Scaling groups, there are two major lifecycle transitions.\n  +  To create a lifecycle hook for scale-out events, specify ` + "`" + `` + "`" + `autoscaling:EC2_INSTANCE_LAUNCHING` + "`" + `` + "`" + `.\n  +  To create a lifecycle hook for scale-in events, specify ` + "`" + `` + "`" + `autoscaling:EC2_INSTANCE_TERMINATING` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "type": "string"
            },
            "notification_metadata": {
              "computed": true,
              "description": "Additional information that you want to include any time Amazon EC2 Auto Scaling sends a message to the notification target.",
              "description_kind": "plain",
              "type": "string"
            },
            "notification_target_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the notification target that Amazon EC2 Auto Scaling sends notifications to when an instance is in a wait state for the lifecycle hook. You can specify an Amazon SNS topic or an Amazon SQS queue.",
              "description_kind": "plain",
              "type": "string"
            },
            "role_arn": {
              "computed": true,
              "description": "The ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target. For information about creating this role, see [Prepare to add a lifecycle hook to your Auto Scaling group](https://docs.aws.amazon.com/autoscaling/ec2/userguide/prepare-for-lifecycle-notifications.html) in the *Amazon EC2 Auto Scaling User Guide*.\n Valid only if the notification target is an Amazon SNS topic or an Amazon SQS queue.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "load_balancer_names": {
        "computed": true,
        "description": "A list of Classic Load Balancers associated with this Auto Scaling group. For Application Load Balancers, Network Load Balancers, and Gateway Load Balancers, specify the ` + "`" + `` + "`" + `TargetGroupARNs` + "`" + `` + "`" + ` property instead.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "max_instance_lifetime": {
        "computed": true,
        "description": "The maximum amount of time, in seconds, that an instance can be in service. The default is null. If specified, the value must be either 0 or a number equal to or greater than 86,400 seconds (1 day). For more information, see [Replace Auto Scaling instances based on maximum instance lifetime](https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-max-instance-lifetime.html) in the *Amazon EC2 Auto Scaling User Guide*.",
        "description_kind": "plain",
        "type": "number"
      },
      "max_size": {
        "computed": true,
        "description": "The maximum size of the group.\n  With a mixed instances policy that uses instance weighting, Amazon EC2 Auto Scaling may need to go above ` + "`" + `` + "`" + `MaxSize` + "`" + `` + "`" + ` to meet your capacity requirements. In this event, Amazon EC2 Auto Scaling will never go above ` + "`" + `` + "`" + `MaxSize` + "`" + `` + "`" + ` by more than your largest instance weight (weights that define how many units each instance contributes to the desired capacity of the group).",
        "description_kind": "plain",
        "type": "string"
      },
      "metrics_collection": {
        "computed": true,
        "description": "Enables the monitoring of group metrics of an Auto Scaling group. By default, these metrics are disabled.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "granularity": {
              "computed": true,
              "description": "The frequency at which Amazon EC2 Auto Scaling sends aggregated data to CloudWatch. The only valid value is ` + "`" + `` + "`" + `1Minute` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "type": "string"
            },
            "metrics": {
              "computed": true,
              "description": "Identifies the metrics to enable.\n You can specify one or more of the following metrics:\n  +   ` + "`" + `` + "`" + `GroupMinSize` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `GroupMaxSize` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `GroupDesiredCapacity` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `GroupInServiceInstances` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `GroupPendingInstances` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `GroupStandbyInstances` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `GroupTerminatingInstances` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `GroupTotalInstances` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `GroupInServiceCapacity` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `GroupPendingCapacity` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `GroupStandbyCapacity` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `GroupTerminatingCapacity` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `GroupTotalCapacity` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `WarmPoolDesiredCapacity` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `WarmPoolWarmedCapacity` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `WarmPoolPendingCapacity` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `WarmPoolTerminatingCapacity` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `WarmPoolTotalCapacity` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `GroupAndWarmPoolDesiredCapacity` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `GroupAndWarmPoolTotalCapacity` + "`" + `` + "`" + ` \n  \n If you specify ` + "`" + `` + "`" + `Granularity` + "`" + `` + "`" + ` and don't specify any metrics, all metrics are enabled.\n For more information, see [Amazon CloudWatch metrics for Amazon EC2 Auto Scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-metrics.html) in the *Amazon EC2 Auto Scaling User Guide*.",
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
      "min_size": {
        "computed": true,
        "description": "The minimum size of the group.",
        "description_kind": "plain",
        "type": "string"
      },
      "mixed_instances_policy": {
        "computed": true,
        "description": "An embedded object that specifies a mixed instances policy.\n The policy includes properties that not only define the distribution of On-Demand Instances and Spot Instances, the maximum price to pay for Spot Instances (optional), and how the Auto Scaling group allocates instance types to fulfill On-Demand and Spot capacities, but also the properties that specify the instance configuration information—the launch template and instance types. The policy can also include a weight for each instance type and different launch templates for individual instance types.\n For more information, see [Auto Scaling groups with multiple instance types and purchase options](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-mixed-instances-groups.html) in the *Amazon EC2 Auto Scaling User Guide*.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "instances_distribution": {
              "computed": true,
              "description": "The instances distribution.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "on_demand_allocation_strategy": {
                    "computed": true,
                    "description": "The allocation strategy to apply to your On-Demand Instances when they are launched. Possible instance types are determined by the launch template overrides that you specify.\n The following lists the valid values:\n  + lowest-price Uses price to determine which instance types are the highest priority, launching the lowest priced instance types within an Availability Zone first. This is the default value for Auto Scaling groups that specify InstanceRequirements. + prioritized You set the order of instance types for the launch template overrides from highest to lowest priority (from first to last in the list). Amazon EC2 Auto Scaling launches your highest priority instance types first. If all your On-Demand capacity cannot be fulfilled using your highest priority instance type, then Amazon EC2 Auto Scaling launches the remaining capacity using the second priority instance type, and so on. This is the default value for Auto Scaling groups that don't specify InstanceRequirements and cannot be used for groups that do.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "on_demand_base_capacity": {
                    "computed": true,
                    "description": "The minimum amount of the Auto Scaling group's capacity that must be fulfilled by On-Demand Instances. This base portion is launched first as your group scales.\n This number has the same unit of measurement as the group's desired capacity. If you change the default unit of measurement (number of instances) by specifying weighted capacity values in your launch template overrides list, or by changing the default desired capacity type setting of the group, you must specify this number using the same unit of measurement.\n Default: 0\n  An update to this setting means a gradual replacement of instances to adjust the current On-Demand Instance levels. When replacing instances, Amazon EC2 Auto Scaling launches new instances before terminating the previous ones.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "on_demand_percentage_above_base_capacity": {
                    "computed": true,
                    "description": "Controls the percentages of On-Demand Instances and Spot Instances for your additional capacity beyond ` + "`" + `` + "`" + `OnDemandBaseCapacity` + "`" + `` + "`" + `. Expressed as a number (for example, 20 specifies 20% On-Demand Instances, 80% Spot Instances). If set to 100, only On-Demand Instances are used.\n Default: 100\n  An update to this setting means a gradual replacement of instances to adjust the current On-Demand and Spot Instance levels for your additional capacity higher than the base capacity. When replacing instances, Amazon EC2 Auto Scaling launches new instances before terminating the previous ones.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "spot_allocation_strategy": {
                    "computed": true,
                    "description": "The allocation strategy to apply to your Spot Instances when they are launched. Possible instance types are determined by the launch template overrides that you specify.\n The following lists the valid values:\n  + capacity-optimized Requests Spot Instances using pools that are optimally chosen based on the available Spot capacity. This strategy has the lowest risk of interruption. To give certain instance types a higher chance of launching first, use capacity-optimized-prioritized. + capacity-optimized-prioritized You set the order of instance types for the launch template overrides from highest to lowest priority (from first to last in the list). Amazon EC2 Auto Scaling honors the instance type priorities on a best effort basis but optimizes for capacity first. Note that if the On-Demand allocation strategy is set to prioritized, the same priority is applied when fulfilling On-Demand capacity. This is not a valid value for Auto Scaling groups that specify InstanceRequirements. + lowest-price Requests Spot Instances using the lowest priced pools within an Availability Zone, across the number of Spot pools that you specify for the SpotInstancePools property. To ensure that your desired capacity is met, you might receive Spot Instances from several pools. This is the default value, but it might lead to high interruption rates because this strategy only considers instance price and not available capacity. + price-capacity-optimized (recommended) The price and capacity optimized allocation strategy looks at both price and capacity to select the Spot Instance pools that are the least likely to be interrupted and have the lowest possible price.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "spot_instance_pools": {
                    "computed": true,
                    "description": "The number of Spot Instance pools across which to allocate your Spot Instances. The Spot pools are determined from the different instance types in the overrides. Valid only when the ` + "`" + `` + "`" + `SpotAllocationStrategy` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `lowest-price` + "`" + `` + "`" + `. Value must be in the range of 1–20.\n Default: 2",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "spot_max_price": {
                    "computed": true,
                    "description": "The maximum price per unit hour that you are willing to pay for a Spot Instance. If your maximum price is lower than the Spot price for the instance types that you selected, your Spot Instances are not launched. We do not recommend specifying a maximum price because it can lead to increased interruptions. When Spot Instances launch, you pay the current Spot price. To remove a maximum price that you previously set, include the property but specify an empty string (\"\") for the value.\n  If you specify a maximum price, your instances will be interrupted more frequently than if you do not specify one.\n  Valid Range: Minimum value of 0.001",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "launch_template": {
              "computed": true,
              "description": "One or more launch templates and the instance types (overrides) that are used to launch EC2 instances to fulfill On-Demand and Spot capacities.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "launch_template_specification": {
                    "computed": true,
                    "description": "The launch template.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "launch_template_id": {
                          "computed": true,
                          "description": "The ID of the launch template.\n You must specify the ` + "`" + `` + "`" + `LaunchTemplateID` + "`" + `` + "`" + ` or the ` + "`" + `` + "`" + `LaunchTemplateName` + "`" + `` + "`" + `, but not both.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "launch_template_name": {
                          "computed": true,
                          "description": "The name of the launch template.\n You must specify the ` + "`" + `` + "`" + `LaunchTemplateName` + "`" + `` + "`" + ` or the ` + "`" + `` + "`" + `LaunchTemplateID` + "`" + `` + "`" + `, but not both.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "version": {
                          "computed": true,
                          "description": "The version number of the launch template.\n Specifying ` + "`" + `` + "`" + `$Latest` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `$Default` + "`" + `` + "`" + ` for the template version number is not supported. However, you can specify ` + "`" + `` + "`" + `LatestVersionNumber` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `DefaultVersionNumber` + "`" + `` + "`" + ` using the ` + "`" + `` + "`" + `Fn::GetAtt` + "`" + `` + "`" + ` intrinsic function. For more information, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).\n  For an example of using the ` + "`" + `` + "`" + `Fn::GetAtt` + "`" + `` + "`" + ` function, see the [Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html#aws-resource-autoscaling-autoscalinggroup--examples) section of the ` + "`" + `` + "`" + `AWS::AutoScaling::AutoScalingGroup` + "`" + `` + "`" + ` resource.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "overrides": {
                    "computed": true,
                    "description": "Any properties that you specify override the same properties in the launch template.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "instance_requirements": {
                          "computed": true,
                          "description": "The instance requirements. Amazon EC2 Auto Scaling uses your specified requirements to identify instance types. Then, it uses your On-Demand and Spot allocation strategies to launch instances from these instance types.\n You can specify up to four separate sets of instance requirements per Auto Scaling group. This is useful for provisioning instances from different Amazon Machine Images (AMIs) in the same Auto Scaling group. To do this, create the AMIs and create a new launch template for each AMI. Then, create a compatible set of instance requirements for each launch template. \n  If you specify ` + "`" + `` + "`" + `InstanceRequirements` + "`" + `` + "`" + `, you can't specify ` + "`" + `` + "`" + `InstanceType` + "`" + `` + "`" + `.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "accelerator_count": {
                                "computed": true,
                                "description": "The minimum and maximum number of accelerators (GPUs, FPGAs, or AWS Inferentia chips) for an instance type.\n To exclude accelerator-enabled instance types, set ` + "`" + `` + "`" + `Max` + "`" + `` + "`" + ` to ` + "`" + `` + "`" + `0` + "`" + `` + "`" + `.\n Default: No minimum or maximum limits",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "max": {
                                      "computed": true,
                                      "description": "The maximum value.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "min": {
                                      "computed": true,
                                      "description": "The minimum value.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "accelerator_manufacturers": {
                                "computed": true,
                                "description": "Indicates whether instance types must have accelerators by specific manufacturers.\n  +  For instance types with NVIDIA devices, specify ` + "`" + `` + "`" + `nvidia` + "`" + `` + "`" + `.\n  +  For instance types with AMD devices, specify ` + "`" + `` + "`" + `amd` + "`" + `` + "`" + `.\n  +  For instance types with AWS devices, specify ` + "`" + `` + "`" + `amazon-web-services` + "`" + `` + "`" + `.\n  +  For instance types with Xilinx devices, specify ` + "`" + `` + "`" + `xilinx` + "`" + `` + "`" + `.\n  \n Default: Any manufacturer",
                                "description_kind": "plain",
                                "type": [
                                  "set",
                                  "string"
                                ]
                              },
                              "accelerator_names": {
                                "computed": true,
                                "description": "Lists the accelerators that must be on an instance type.\n  +  For instance types with NVIDIA A100 GPUs, specify ` + "`" + `` + "`" + `a100` + "`" + `` + "`" + `.\n  +  For instance types with NVIDIA V100 GPUs, specify ` + "`" + `` + "`" + `v100` + "`" + `` + "`" + `.\n  +  For instance types with NVIDIA K80 GPUs, specify ` + "`" + `` + "`" + `k80` + "`" + `` + "`" + `.\n  +  For instance types with NVIDIA T4 GPUs, specify ` + "`" + `` + "`" + `t4` + "`" + `` + "`" + `.\n  +  For instance types with NVIDIA M60 GPUs, specify ` + "`" + `` + "`" + `m60` + "`" + `` + "`" + `.\n  +  For instance types with AMD Radeon Pro V520 GPUs, specify ` + "`" + `` + "`" + `radeon-pro-v520` + "`" + `` + "`" + `.\n  +  For instance types with Xilinx VU9P FPGAs, specify ` + "`" + `` + "`" + `vu9p` + "`" + `` + "`" + `.\n  \n Default: Any accelerator",
                                "description_kind": "plain",
                                "type": [
                                  "set",
                                  "string"
                                ]
                              },
                              "accelerator_total_memory_mi_b": {
                                "computed": true,
                                "description": "The minimum and maximum total memory size for the accelerators on an instance type, in MiB.\n Default: No minimum or maximum limits",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "max": {
                                      "computed": true,
                                      "description": "The memory maximum in MiB.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "min": {
                                      "computed": true,
                                      "description": "The memory minimum in MiB.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "accelerator_types": {
                                "computed": true,
                                "description": "Lists the accelerator types that must be on an instance type.\n  +  For instance types with GPU accelerators, specify ` + "`" + `` + "`" + `gpu` + "`" + `` + "`" + `.\n  +  For instance types with FPGA accelerators, specify ` + "`" + `` + "`" + `fpga` + "`" + `` + "`" + `.\n  +  For instance types with inference accelerators, specify ` + "`" + `` + "`" + `inference` + "`" + `` + "`" + `.\n  \n Default: Any accelerator type",
                                "description_kind": "plain",
                                "type": [
                                  "set",
                                  "string"
                                ]
                              },
                              "allowed_instance_types": {
                                "computed": true,
                                "description": "The instance types to apply your specified attributes against. All other instance types are ignored, even if they match your specified attributes.\n You can use strings with one or more wild cards, represented by an asterisk (` + "`" + `` + "`" + `*` + "`" + `` + "`" + `), to allow an instance type, size, or generation. The following are examples: ` + "`" + `` + "`" + `m5.8xlarge` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `c5*.*` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `m5a.*` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `r*` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `*3*` + "`" + `` + "`" + `.\n For example, if you specify ` + "`" + `` + "`" + `c5*` + "`" + `` + "`" + `, Amazon EC2 Auto Scaling will allow the entire C5 instance family, which includes all C5a and C5n instance types. If you specify ` + "`" + `` + "`" + `m5a.*` + "`" + `` + "`" + `, Amazon EC2 Auto Scaling will allow all the M5a instance types, but not the M5n instance types.\n  If you specify ` + "`" + `` + "`" + `AllowedInstanceTypes` + "`" + `` + "`" + `, you can't specify ` + "`" + `` + "`" + `ExcludedInstanceTypes` + "`" + `` + "`" + `.\n  Default: All instance types",
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "bare_metal": {
                                "computed": true,
                                "description": "Indicates whether bare metal instance types are included, excluded, or required.\n Default: ` + "`" + `` + "`" + `excluded` + "`" + `` + "`" + `",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "baseline_ebs_bandwidth_mbps": {
                                "computed": true,
                                "description": "The minimum and maximum baseline bandwidth performance for an instance type, in Mbps. For more information, see [Amazon EBS–optimized instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-optimized.html) in the *Amazon EC2 User Guide for Linux Instances*.\n Default: No minimum or maximum limits",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "max": {
                                      "computed": true,
                                      "description": "The maximum value in Mbps.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "min": {
                                      "computed": true,
                                      "description": "The minimum value in Mbps.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "baseline_performance_factors": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "cpu": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "references": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "instance_family": {
                                                  "computed": true,
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
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "burstable_performance": {
                                "computed": true,
                                "description": "Indicates whether burstable performance instance types are included, excluded, or required. For more information, see [Burstable performance instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances.html) in the *Amazon EC2 User Guide for Linux Instances*.\n Default: ` + "`" + `` + "`" + `excluded` + "`" + `` + "`" + `",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "cpu_manufacturers": {
                                "computed": true,
                                "description": "Lists which specific CPU manufacturers to include.\n  +  For instance types with Intel CPUs, specify ` + "`" + `` + "`" + `intel` + "`" + `` + "`" + `.\n  +  For instance types with AMD CPUs, specify ` + "`" + `` + "`" + `amd` + "`" + `` + "`" + `.\n  +  For instance types with AWS CPUs, specify ` + "`" + `` + "`" + `amazon-web-services` + "`" + `` + "`" + `.\n  \n  Don't confuse the CPU hardware manufacturer with the CPU hardware architecture. Instances will be launched with a compatible CPU architecture based on the Amazon Machine Image (AMI) that you specify in your launch template. \n  Default: Any manufacturer",
                                "description_kind": "plain",
                                "type": [
                                  "set",
                                  "string"
                                ]
                              },
                              "excluded_instance_types": {
                                "computed": true,
                                "description": "The instance types to exclude. You can use strings with one or more wild cards, represented by an asterisk (` + "`" + `` + "`" + `*` + "`" + `` + "`" + `), to exclude an instance family, type, size, or generation. The following are examples: ` + "`" + `` + "`" + `m5.8xlarge` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `c5*.*` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `m5a.*` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `r*` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `*3*` + "`" + `` + "`" + `. \n For example, if you specify ` + "`" + `` + "`" + `c5*` + "`" + `` + "`" + `, you are excluding the entire C5 instance family, which includes all C5a and C5n instance types. If you specify ` + "`" + `` + "`" + `m5a.*` + "`" + `` + "`" + `, Amazon EC2 Auto Scaling will exclude all the M5a instance types, but not the M5n instance types.\n  If you specify ` + "`" + `` + "`" + `ExcludedInstanceTypes` + "`" + `` + "`" + `, you can't specify ` + "`" + `` + "`" + `AllowedInstanceTypes` + "`" + `` + "`" + `.\n  Default: No excluded instance types",
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "instance_generations": {
                                "computed": true,
                                "description": "Indicates whether current or previous generation instance types are included.\n  +  For current generation instance types, specify ` + "`" + `` + "`" + `current` + "`" + `` + "`" + `. The current generation includes EC2 instance types currently recommended for use. This typically includes the latest two to three generations in each instance family. For more information, see [Instance types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html) in the *Amazon EC2 User Guide for Linux Instances*.\n  +  For previous generation instance types, specify ` + "`" + `` + "`" + `previous` + "`" + `` + "`" + `.\n  \n Default: Any current or previous generation",
                                "description_kind": "plain",
                                "type": [
                                  "set",
                                  "string"
                                ]
                              },
                              "local_storage": {
                                "computed": true,
                                "description": "Indicates whether instance types with instance store volumes are included, excluded, or required. For more information, see [Amazon EC2 instance store](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html) in the *Amazon EC2 User Guide for Linux Instances*.\n Default: ` + "`" + `` + "`" + `included` + "`" + `` + "`" + `",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "local_storage_types": {
                                "computed": true,
                                "description": "Indicates the type of local storage that is required.\n  +  For instance types with hard disk drive (HDD) storage, specify ` + "`" + `` + "`" + `hdd` + "`" + `` + "`" + `.\n  +  For instance types with solid state drive (SSD) storage, specify ` + "`" + `` + "`" + `ssd` + "`" + `` + "`" + `.\n  \n Default: Any local storage type",
                                "description_kind": "plain",
                                "type": [
                                  "set",
                                  "string"
                                ]
                              },
                              "max_spot_price_as_percentage_of_optimal_on_demand_price": {
                                "computed": true,
                                "description": "[Price protection] The price protection threshold for Spot Instances, as a percentage of an identified On-Demand price. The identified On-Demand price is the price of the lowest priced current generation C, M, or R instance type with your specified attributes. If no current generation C, M, or R instance type matches your attributes, then the identified price is from either the lowest priced current generation instance types or, failing that, the lowest priced previous generation instance types that match your attributes. When Amazon EC2 Auto Scaling selects instance types with your attributes, we will exclude instance types whose price exceeds your specified threshold.\n The parameter accepts an integer, which Amazon EC2 Auto Scaling interprets as a percentage.\n If you set ` + "`" + `` + "`" + `DesiredCapacityType` + "`" + `` + "`" + ` to ` + "`" + `` + "`" + `vcpu` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `memory-mib` + "`" + `` + "`" + `, the price protection threshold is based on the per-vCPU or per-memory price instead of the per instance price. \n  Only one of ` + "`" + `` + "`" + `SpotMaxPricePercentageOverLowestPrice` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `MaxSpotPriceAsPercentageOfOptimalOnDemandPrice` + "`" + `` + "`" + ` can be specified. If you don't specify either, Amazon EC2 Auto Scaling will automatically apply optimal price protection to consistently select from a wide range of instance types. To indicate no price protection threshold for Spot Instances, meaning you want to consider all instance types that match your attributes, include one of these parameters and specify a high value, such as ` + "`" + `` + "`" + `999999` + "`" + `` + "`" + `.",
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "memory_gi_b_per_v_cpu": {
                                "computed": true,
                                "description": "The minimum and maximum amount of memory per vCPU for an instance type, in GiB.\n Default: No minimum or maximum limits",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "max": {
                                      "computed": true,
                                      "description": "The memory maximum in GiB.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "min": {
                                      "computed": true,
                                      "description": "The memory minimum in GiB.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "memory_mi_b": {
                                "computed": true,
                                "description": "The minimum and maximum instance memory size for an instance type, in MiB.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "max": {
                                      "computed": true,
                                      "description": "The memory maximum in MiB.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "min": {
                                      "computed": true,
                                      "description": "The memory minimum in MiB.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "network_bandwidth_gbps": {
                                "computed": true,
                                "description": "The minimum and maximum amount of network bandwidth, in gigabits per second (Gbps).\n Default: No minimum or maximum limits",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "max": {
                                      "computed": true,
                                      "description": "The maximum amount of network bandwidth, in gigabits per second (Gbps).",
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "min": {
                                      "computed": true,
                                      "description": "The minimum amount of network bandwidth, in gigabits per second (Gbps).",
                                      "description_kind": "plain",
                                      "type": "number"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "network_interface_count": {
                                "computed": true,
                                "description": "The minimum and maximum number of network interfaces for an instance type.\n Default: No minimum or maximum limits",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "max": {
                                      "computed": true,
                                      "description": "The maximum number of network interfaces.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "min": {
                                      "computed": true,
                                      "description": "The minimum number of network interfaces.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "on_demand_max_price_percentage_over_lowest_price": {
                                "computed": true,
                                "description": "[Price protection] The price protection threshold for On-Demand Instances, as a percentage higher than an identified On-Demand price. The identified On-Demand price is the price of the lowest priced current generation C, M, or R instance type with your specified attributes. If no current generation C, M, or R instance type matches your attributes, then the identified price is from either the lowest priced current generation instance types or, failing that, the lowest priced previous generation instance types that match your attributes. When Amazon EC2 Auto Scaling selects instance types with your attributes, we will exclude instance types whose price exceeds your specified threshold. \n The parameter accepts an integer, which Amazon EC2 Auto Scaling interprets as a percentage.\n To turn off price protection, specify a high value, such as ` + "`" + `` + "`" + `999999` + "`" + `` + "`" + `. \n If you set ` + "`" + `` + "`" + `DesiredCapacityType` + "`" + `` + "`" + ` to ` + "`" + `` + "`" + `vcpu` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `memory-mib` + "`" + `` + "`" + `, the price protection threshold is applied based on the per-vCPU or per-memory price instead of the per instance price. \n Default: ` + "`" + `` + "`" + `20` + "`" + `` + "`" + `",
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "require_hibernate_support": {
                                "computed": true,
                                "description": "Indicates whether instance types must provide On-Demand Instance hibernation support.\n Default: ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `",
                                "description_kind": "plain",
                                "type": "bool"
                              },
                              "spot_max_price_percentage_over_lowest_price": {
                                "computed": true,
                                "description": "[Price protection] The price protection threshold for Spot Instances, as a percentage higher than an identified Spot price. The identified Spot price is the price of the lowest priced current generation C, M, or R instance type with your specified attributes. If no current generation C, M, or R instance type matches your attributes, then the identified price is from either the lowest priced current generation instance types or, failing that, the lowest priced previous generation instance types that match your attributes. When Amazon EC2 Auto Scaling selects instance types with your attributes, we will exclude instance types whose price exceeds your specified threshold.\n The parameter accepts an integer, which Amazon EC2 Auto Scaling interprets as a percentage. \n If you set ` + "`" + `` + "`" + `DesiredCapacityType` + "`" + `` + "`" + ` to ` + "`" + `` + "`" + `vcpu` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `memory-mib` + "`" + `` + "`" + `, the price protection threshold is based on the per-vCPU or per-memory price instead of the per instance price. \n  Only one of ` + "`" + `` + "`" + `SpotMaxPricePercentageOverLowestPrice` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `MaxSpotPriceAsPercentageOfOptimalOnDemandPrice` + "`" + `` + "`" + ` can be specified. If you don't specify either, Amazon EC2 Auto Scaling will automatically apply optimal price protection to consistently select from a wide range of instance types. To indicate no price protection threshold for Spot Instances, meaning you want to consider all instance types that match your attributes, include one of these parameters and specify a high value, such as ` + "`" + `` + "`" + `999999` + "`" + `` + "`" + `.",
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "total_local_storage_gb": {
                                "computed": true,
                                "description": "The minimum and maximum total local storage size for an instance type, in GB.\n Default: No minimum or maximum limits",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "max": {
                                      "computed": true,
                                      "description": "The storage maximum in GB.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "min": {
                                      "computed": true,
                                      "description": "The storage minimum in GB.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "v_cpu_count": {
                                "computed": true,
                                "description": "The minimum and maximum number of vCPUs for an instance type.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "max": {
                                      "computed": true,
                                      "description": "The maximum number of vCPUs.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "min": {
                                      "computed": true,
                                      "description": "The minimum number of vCPUs.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "instance_type": {
                          "computed": true,
                          "description": "The instance type, such as ` + "`" + `` + "`" + `m3.xlarge` + "`" + `` + "`" + `. You must specify an instance type that is supported in your requested Region and Availability Zones. For more information, see [Instance types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html) in the *Amazon EC2 User Guide for Linux Instances*.\n You can specify up to 40 instance types per Auto Scaling group.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "launch_template_specification": {
                          "computed": true,
                          "description": "Provides a launch template for the specified instance type or set of instance requirements. For example, some instance types might require a launch template with a different AMI. If not provided, Amazon EC2 Auto Scaling uses the launch template that's specified in the ` + "`" + `` + "`" + `LaunchTemplate` + "`" + `` + "`" + ` definition. For more information, see [Specifying a different launch template for an instance type](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-mixed-instances-groups-launch-template-overrides.html) in the *Amazon EC2 Auto Scaling User Guide*. \n You can specify up to 20 launch templates per Auto Scaling group. The launch templates specified in the overrides and in the ` + "`" + `` + "`" + `LaunchTemplate` + "`" + `` + "`" + ` definition count towards this limit.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "launch_template_id": {
                                "computed": true,
                                "description": "The ID of the launch template.\n You must specify the ` + "`" + `` + "`" + `LaunchTemplateID` + "`" + `` + "`" + ` or the ` + "`" + `` + "`" + `LaunchTemplateName` + "`" + `` + "`" + `, but not both.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "launch_template_name": {
                                "computed": true,
                                "description": "The name of the launch template.\n You must specify the ` + "`" + `` + "`" + `LaunchTemplateName` + "`" + `` + "`" + ` or the ` + "`" + `` + "`" + `LaunchTemplateID` + "`" + `` + "`" + `, but not both.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "version": {
                                "computed": true,
                                "description": "The version number of the launch template.\n Specifying ` + "`" + `` + "`" + `$Latest` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `$Default` + "`" + `` + "`" + ` for the template version number is not supported. However, you can specify ` + "`" + `` + "`" + `LatestVersionNumber` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `DefaultVersionNumber` + "`" + `` + "`" + ` using the ` + "`" + `` + "`" + `Fn::GetAtt` + "`" + `` + "`" + ` intrinsic function. For more information, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).\n  For an example of using the ` + "`" + `` + "`" + `Fn::GetAtt` + "`" + `` + "`" + ` function, see the [Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html#aws-resource-autoscaling-autoscalinggroup--examples) section of the ` + "`" + `` + "`" + `AWS::AutoScaling::AutoScalingGroup` + "`" + `` + "`" + ` resource.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "weighted_capacity": {
                          "computed": true,
                          "description": "If you provide a list of instance types to use, you can specify the number of capacity units provided by each instance type in terms of virtual CPUs, memory, storage, throughput, or other relative performance characteristic. When a Spot or On-Demand Instance is launched, the capacity units count toward the desired capacity. Amazon EC2 Auto Scaling launches instances until the desired capacity is totally fulfilled, even if this results in an overage. For example, if there are two units remaining to fulfill capacity, and Amazon EC2 Auto Scaling can only launch an instance with a ` + "`" + `` + "`" + `WeightedCapacity` + "`" + `` + "`" + ` of five units, the instance is launched, and the desired capacity is exceeded by three units. For more information, see [Configure instance weighting for Amazon EC2 Auto Scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-mixed-instances-groups-instance-weighting.html) in the *Amazon EC2 Auto Scaling User Guide*. Value must be in the range of 1-999. \n If you specify a value for ` + "`" + `` + "`" + `WeightedCapacity` + "`" + `` + "`" + ` for one instance type, you must specify a value for ` + "`" + `` + "`" + `WeightedCapacity` + "`" + `` + "`" + ` for all of them.\n  Every Auto Scaling group has three size parameters (` + "`" + `` + "`" + `DesiredCapacity` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `MaxSize` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `MinSize` + "`" + `` + "`" + `). Usually, you set these sizes based on a specific number of instances. However, if you configure a mixed instances policy that defines weights for the instance types, you must specify these sizes with the same units that you use for weighting instances.",
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
            }
          },
          "nesting_mode": "single"
        }
      },
      "new_instances_protected_from_scale_in": {
        "computed": true,
        "description": "Indicates whether newly launched instances are protected from termination by Amazon EC2 Auto Scaling when scaling in. For more information about preventing instances from terminating on scale in, see [Use instance scale-in protection](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-instance-protection.html) in the *Amazon EC2 Auto Scaling User Guide*.",
        "description_kind": "plain",
        "type": "bool"
      },
      "notification_configuration": {
        "computed": true,
        "description": "A structure that specifies an Amazon SNS notification configuration for the ` + "`" + `` + "`" + `NotificationConfigurations` + "`" + `` + "`" + ` property of the [AWS::AutoScaling::AutoScalingGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html) resource.\n For an example template snippet, see [Configure Amazon EC2 Auto Scaling resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/quickref-ec2-auto-scaling.html).\n For more information, see [Get Amazon SNS notifications when your Auto Scaling group scales](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ASGettingNotifications.html) in the *Amazon EC2 Auto Scaling User Guide*.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "notification_types": {
              "computed": true,
              "description": "A list of event types that send a notification. Event types can include any of the following types. \n  *Allowed values*:\n  +   ` + "`" + `` + "`" + `autoscaling:EC2_INSTANCE_LAUNCH` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `autoscaling:EC2_INSTANCE_LAUNCH_ERROR` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `autoscaling:EC2_INSTANCE_TERMINATE` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `autoscaling:EC2_INSTANCE_TERMINATE_ERROR` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `autoscaling:TEST_NOTIFICATION` + "`" + `` + "`" + `",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "topic_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the Amazon SNS topic.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "notification_configurations": {
        "computed": true,
        "description": "Configures an Auto Scaling group to send notifications when specified events take place.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "notification_types": {
              "computed": true,
              "description": "A list of event types that send a notification. Event types can include any of the following types. \n  *Allowed values*:\n  +   ` + "`" + `` + "`" + `autoscaling:EC2_INSTANCE_LAUNCH` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `autoscaling:EC2_INSTANCE_LAUNCH_ERROR` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `autoscaling:EC2_INSTANCE_TERMINATE` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `autoscaling:EC2_INSTANCE_TERMINATE_ERROR` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `autoscaling:TEST_NOTIFICATION` + "`" + `` + "`" + `",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "topic_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the Amazon SNS topic.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "placement_group": {
        "computed": true,
        "description": "The name of the placement group into which to launch your instances. For more information, see [Placement groups](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html) in the *Amazon EC2 User Guide for Linux Instances*.\n  A *cluster* placement group is a logical grouping of instances within a single Availability Zone. You cannot specify multiple Availability Zones and a cluster placement group.",
        "description_kind": "plain",
        "type": "string"
      },
      "service_linked_role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the service-linked role that the Auto Scaling group uses to call other AWS service on your behalf. By default, Amazon EC2 Auto Scaling uses a service-linked role named ` + "`" + `` + "`" + `AWSServiceRoleForAutoScaling` + "`" + `` + "`" + `, which it creates if it does not exist. For more information, see [Service-linked roles](https://docs.aws.amazon.com/autoscaling/ec2/userguide/autoscaling-service-linked-role.html) in the *Amazon EC2 Auto Scaling User Guide*.",
        "description_kind": "plain",
        "type": "string"
      },
      "skip_zonal_shift_validation": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "tags": {
        "computed": true,
        "description": "One or more tags. You can tag your Auto Scaling group and propagate the tags to the Amazon EC2 instances it launches. Tags are not propagated to Amazon EBS volumes. To add tags to Amazon EBS volumes, specify the tags in a launch template but use caution. If the launch template specifies an instance tag with a key that is also specified for the Auto Scaling group, Amazon EC2 Auto Scaling overrides the value of that instance tag with the value specified by the Auto Scaling group. For more information, see [Tag Auto Scaling groups and instances](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-tagging.html) in the *Amazon EC2 Auto Scaling User Guide*.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key.",
              "description_kind": "plain",
              "type": "string"
            },
            "propagate_at_launch": {
              "computed": true,
              "description": "Set to ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` if you want CloudFormation to copy the tag to EC2 instances that are launched as part of the Auto Scaling group. Set to ` + "`" + `` + "`" + `false` + "`" + `` + "`" + ` if you want the tag attached only to the Auto Scaling group and not copied to any instances launched as part of the Auto Scaling group.",
              "description_kind": "plain",
              "type": "bool"
            },
            "value": {
              "computed": true,
              "description": "The tag value.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "target_group_ar_ns": {
        "computed": true,
        "description": "The Amazon Resource Names (ARN) of the Elastic Load Balancing target groups to associate with the Auto Scaling group. Instances are registered as targets with the target groups. The target groups receive incoming traffic and route requests to one or more registered targets. For more information, see [Use Elastic Load Balancing to distribute traffic across the instances in your Auto Scaling group](https://docs.aws.amazon.com/autoscaling/ec2/userguide/autoscaling-load-balancer.html) in the *Amazon EC2 Auto Scaling User Guide*.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "termination_policies": {
        "computed": true,
        "description": "A policy or a list of policies that are used to select the instance to terminate. These policies are executed in the order that you list them. For more information, see [Configure termination policies for Amazon EC2 Auto Scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-termination-policies.html) in the *Amazon EC2 Auto Scaling User Guide*.\n Valid values: ` + "`" + `` + "`" + `Default` + "`" + `` + "`" + ` | ` + "`" + `` + "`" + `AllocationStrategy` + "`" + `` + "`" + ` | ` + "`" + `` + "`" + `ClosestToNextInstanceHour` + "`" + `` + "`" + ` | ` + "`" + `` + "`" + `NewestInstance` + "`" + `` + "`" + ` | ` + "`" + `` + "`" + `OldestInstance` + "`" + `` + "`" + ` | ` + "`" + `` + "`" + `OldestLaunchConfiguration` + "`" + `` + "`" + ` | ` + "`" + `` + "`" + `OldestLaunchTemplate` + "`" + `` + "`" + ` | ` + "`" + `` + "`" + `arn:aws:lambda:region:account-id:function:my-function:my-alias` + "`" + `` + "`" + `",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "traffic_sources": {
        "computed": true,
        "description": "The traffic sources associated with this Auto Scaling group.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "identifier": {
              "computed": true,
              "description": "Identifies the traffic source.\n For Application Load Balancers, Gateway Load Balancers, Network Load Balancers, and VPC Lattice, this will be the Amazon Resource Name (ARN) for a target group in this account and Region. For Classic Load Balancers, this will be the name of the Classic Load Balancer in this account and Region.\n For example: \n  +  Application Load Balancer ARN: ` + "`" + `` + "`" + `arn:aws:elasticloadbalancing:us-west-2:123456789012:targetgroup/my-targets/1234567890123456` + "`" + `` + "`" + ` \n  +  Classic Load Balancer name: ` + "`" + `` + "`" + `my-classic-load-balancer` + "`" + `` + "`" + ` \n  +  VPC Lattice ARN: ` + "`" + `` + "`" + `arn:aws:vpc-lattice:us-west-2:123456789012:targetgroup/tg-1234567890123456` + "`" + `` + "`" + ` \n  \n To get the ARN of a target group for a Application Load Balancer, Gateway Load Balancer, or Network Load Balancer, or the name of a Classic Load Balancer, use the Elastic Load Balancing [DescribeTargetGroups](https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_DescribeTargetGroups.html) and [DescribeLoadBalancers](https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_DescribeLoadBalancers.html) API operations.\n To get the ARN of a target group for VPC Lattice, use the VPC Lattice [GetTargetGroup](https://docs.aws.amazon.com/vpc-lattice/latest/APIReference/API_GetTargetGroup.html) API operation.",
              "description_kind": "plain",
              "type": "string"
            },
            "type": {
              "computed": true,
              "description": "Provides additional context for the value of ` + "`" + `` + "`" + `Identifier` + "`" + `` + "`" + `.\n The following lists the valid values:\n  +   ` + "`" + `` + "`" + `elb` + "`" + `` + "`" + ` if ` + "`" + `` + "`" + `Identifier` + "`" + `` + "`" + ` is the name of a Classic Load Balancer.\n  +   ` + "`" + `` + "`" + `elbv2` + "`" + `` + "`" + ` if ` + "`" + `` + "`" + `Identifier` + "`" + `` + "`" + ` is the ARN of an Application Load Balancer, Gateway Load Balancer, or Network Load Balancer target group.\n  +   ` + "`" + `` + "`" + `vpc-lattice` + "`" + `` + "`" + ` if ` + "`" + `` + "`" + `Identifier` + "`" + `` + "`" + ` is the ARN of a VPC Lattice target group.\n  \n Required if the identifier is the name of a Classic Load Balancer.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "vpc_zone_identifier": {
        "computed": true,
        "description": "A list of subnet IDs for a virtual private cloud (VPC) where instances in the Auto Scaling group can be created.\n If this resource specifies public subnets and is also in a VPC that is defined in the same stack template, you must use the [DependsOn attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) to declare a dependency on the [VPC-gateway attachment](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc-gateway-attachment.html).\n  When you update ` + "`" + `` + "`" + `VPCZoneIdentifier` + "`" + `` + "`" + `, this retains the same Auto Scaling group and replaces old instances with new ones, according to the specified subnets. You can optionally specify how CloudFormation handles these updates by using an [UpdatePolicy attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html).\n  Required to launch instances into a nondefault VPC. If you specify ` + "`" + `` + "`" + `VPCZoneIdentifier` + "`" + `` + "`" + ` with ` + "`" + `` + "`" + `AvailabilityZones` + "`" + `` + "`" + `, the subnets that you specify for this property must reside in those Availability Zones.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::AutoScaling::AutoScalingGroup",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAutoscalingAutoScalingGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAutoscalingAutoScalingGroup), &result)
	return &result
}
