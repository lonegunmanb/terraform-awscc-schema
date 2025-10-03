package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGameliftFleet = `{
  "block": {
    "attributes": {
      "anywhere_configuration": {
        "computed": true,
        "description": "Configuration for Anywhere fleet.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cost": {
              "computed": true,
              "description": "Cost of compute can be specified on Anywhere Fleets to prioritize placement across Queue destinations based on Cost.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "apply_capacity": {
        "computed": true,
        "description": "Determines when and how to apply fleet or location capacities. Allowed options are ON_UPDATE (default), ON_CREATE_AND_UPDATE and ON_CREATE_AND_UPDATE_WITH_AUTOSCALING. If you choose ON_CREATE_AND_UPDATE_WITH_AUTOSCALING, MinSize and MaxSize will still be applied on creation and on updates, but DesiredEC2Instances will only be applied once on fleet creation and will be ignored during updates to prevent conflicts with auto-scaling. During updates with ON_CREATE_AND_UPDATE_WITH_AUTOSCALING chosen, if current desired instance is lower than the new MinSize, it will be increased to the new MinSize; if current desired instance is larger than the new MaxSize, it will be decreased to the new MaxSize.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "build_id": {
        "computed": true,
        "description": "A unique identifier for a build to be deployed on the new fleet. If you are deploying the fleet with a custom game build, you must specify this property. The build must have been successfully uploaded to Amazon GameLift and be in a READY status. This fleet setting cannot be changed once the fleet is created.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "certificate_configuration": {
        "computed": true,
        "description": "Indicates whether to generate a TLS/SSL certificate for the new fleet. TLS certificates are used for encrypting traffic between game clients and game servers running on GameLift. If this parameter is not set, certificate generation is disabled. This fleet setting cannot be changed once the fleet is created.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "certificate_type": {
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
      "compute_type": {
        "computed": true,
        "description": "ComputeType to differentiate EC2 hardware managed by GameLift and Anywhere hardware managed by the customer.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A human-readable description of a fleet.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "desired_ec2_instances": {
        "computed": true,
        "description": "[DEPRECATED] The number of EC2 instances that you want this fleet to host. When creating a new fleet, GameLift automatically sets this value to \"1\" and initiates a single instance. Once the fleet is active, update this value to trigger GameLift to add or remove instances from the fleet.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "ec2_inbound_permissions": {
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
      "ec2_instance_type": {
        "computed": true,
        "description": "The name of an EC2 instance type that is supported in Amazon GameLift. A fleet instance type determines the computing resources of each instance in the fleet, including CPU, memory, storage, and networking capacity. Amazon GameLift supports the following EC2 instance types. See Amazon EC2 Instance Types for detailed descriptions.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "fleet_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) that is assigned to a Amazon GameLift Servers Fleet resource and uniquely identifies it. ARNs are unique across all Regions. In a GameLift Fleet ARN, the resource ID matches the FleetId value.",
        "description_kind": "plain",
        "type": "string"
      },
      "fleet_id": {
        "computed": true,
        "description": "Unique fleet ID",
        "description_kind": "plain",
        "type": "string"
      },
      "fleet_type": {
        "computed": true,
        "description": "Indicates whether to use On-Demand instances or Spot instances for this fleet. If empty, the default is ON_DEMAND. Both categories of instances use identical hardware and configurations based on the instance type selected for this fleet.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_role_arn": {
        "computed": true,
        "description": "A unique identifier for an AWS IAM role that manages access to your AWS services. With an instance role ARN set, any application that runs on an instance in this fleet can assume the role, including install scripts, server processes, and daemons (background processes). Create a role or look up a role's ARN from the IAM dashboard in the AWS Management Console.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_role_credentials_provider": {
        "computed": true,
        "description": "Credentials provider implementation that loads credentials from the Amazon EC2 Instance Metadata Service.",
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
                    "description": "Defaults to MinSize if not defined. The number of EC2 instances you want to maintain in the specified fleet location. This value must fall between the minimum and maximum size limits.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "max_size": {
                    "computed": true,
                    "description": "The maximum value that is allowed for the fleet's instance count for a location. When creating a new fleet, GameLift automatically sets this value to \"1\". Once the fleet is active, you can change this value.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "min_size": {
                    "computed": true,
                    "description": "The minimum value allowed for the fleet's instance count for a location. When creating a new fleet, GameLift automatically sets this value to \"0\". After the fleet is active, you can change this value.",
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
          "nesting_mode": "list"
        },
        "optional": true
      },
      "log_paths": {
        "computed": true,
        "description": "This parameter is no longer used. When hosting a custom game build, specify where Amazon GameLift should store log files using the Amazon GameLift server API call ProcessReady()",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "max_size": {
        "computed": true,
        "description": "[DEPRECATED] The maximum value that is allowed for the fleet's instance count. When creating a new fleet, GameLift automatically sets this value to \"1\". Once the fleet is active, you can change this value.",
        "description_kind": "plain",
        "optional": true,
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
      "min_size": {
        "computed": true,
        "description": "[DEPRECATED] The minimum value allowed for the fleet's instance count. When creating a new fleet, GameLift automatically sets this value to \"0\". After the fleet is active, you can change this value.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "description": "A descriptive label that is associated with a fleet. Fleet names do not need to be unique.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "new_game_session_protection_policy": {
        "computed": true,
        "description": "A game session protection policy to apply to all game sessions hosted on instances in this fleet. When protected, active game sessions cannot be terminated during a scale-down event. If this parameter is not set, instances in this fleet default to no protection. You can change a fleet's protection policy to affect future game sessions on the fleet. You can also set protection for individual game sessions.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "peer_vpc_aws_account_id": {
        "computed": true,
        "description": "A unique identifier for the AWS account with the VPC that you want to peer your Amazon GameLift fleet with. You can find your account ID in the AWS Management Console under account settings.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "peer_vpc_id": {
        "computed": true,
        "description": "A unique identifier for a VPC with resources to be accessed by your Amazon GameLift fleet. The VPC must be in the same Region as your fleet. To look up a VPC ID, use the VPC Dashboard in the AWS Management Console.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_creation_limit_policy": {
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
      "runtime_configuration": {
        "computed": true,
        "description": "Instructions for launching server processes on each instance in the fleet. Server processes run either a custom game build executable or a Realtime script. The runtime configuration defines the server executables or launch script file, launch parameters, and the number of processes to run concurrently on each instance. When creating a fleet, the runtime configuration must have at least one server process configuration; otherwise the request fails with an invalid request exception.\n\nThis parameter is required unless the parameters ServerLaunchPath and ServerLaunchParameters are defined. Runtime configuration has replaced these parameters, but fleets that use them will continue to work.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "game_session_activation_timeout_seconds": {
              "computed": true,
              "description": "The maximum amount of time (in seconds) that a game session can remain in status ACTIVATING. If the game session is not active before the timeout, activation is terminated and the game session status is changed to TERMINATED.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_concurrent_game_session_activations": {
              "computed": true,
              "description": "The maximum number of game sessions with status ACTIVATING to allow on an instance simultaneously. This setting limits the amount of instance resources that can be used for new game activations at any one time.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "server_processes": {
              "computed": true,
              "description": "A collection of server process configurations that describe which server processes to run on each instance in a fleet.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "concurrent_executions": {
                    "computed": true,
                    "description": "The number of server processes that use this configuration to run concurrently on an instance.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "launch_path": {
                    "computed": true,
                    "description": "The location of the server executable in a custom game build or the name of the Realtime script file that contains the Init() function. Game builds and Realtime scripts are installed on instances at the root:\n\nWindows (for custom game builds only): C:\\game. Example: \"C:\\game\\MyGame\\server.exe\"\n\nLinux: /local/game. Examples: \"/local/game/MyGame/server.exe\" or \"/local/game/MyRealtimeScript.js\"",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "parameters": {
                    "computed": true,
                    "description": "An optional list of parameters to pass to the server executable or Realtime script on launch.",
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
            "location": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
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
            "status": {
              "computed": true,
              "description": "Current status of the scaling policy. The scaling policy can be in force only when in an ACTIVE status. Scaling policies can be suspended for individual fleets. If the policy is suspended for a fleet, the policy status does not change.",
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
            },
            "update_status": {
              "computed": true,
              "description": "The current status of the fleet's scaling policies in a requested fleet location. The status PENDING_UPDATE indicates that an update was requested for the fleet but has not yet been completed for the location.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "script_id": {
        "computed": true,
        "description": "A unique identifier for a Realtime script to be deployed on a new Realtime Servers fleet. The script must have been successfully uploaded to Amazon GameLift. This fleet setting cannot be changed once the fleet is created.\n\nNote: It is not currently possible to use the !Ref command to reference a script created with a CloudFormation template for the fleet property ScriptId. Instead, use Fn::GetAtt Script.Arn or Fn::GetAtt Script.Id to retrieve either of these properties as input for ScriptId. Alternatively, enter a ScriptId string manually.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "server_launch_parameters": {
        "computed": true,
        "description": "This parameter is no longer used but is retained for backward compatibility. Instead, specify server launch parameters in the RuntimeConfiguration parameter. A request must specify either a runtime configuration or values for both ServerLaunchParameters and ServerLaunchPath.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "server_launch_path": {
        "computed": true,
        "description": "This parameter is no longer used. Instead, specify a server launch path using the RuntimeConfiguration parameter. Requests that specify a server launch path and launch parameters instead of a runtime configuration will continue to work.",
        "description_kind": "plain",
        "optional": true,
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
    "description": "The AWS::GameLift::Fleet resource creates an Amazon GameLift (GameLift) fleet to host game servers. A fleet is a set of EC2 or Anywhere instances, each of which can host multiple game sessions.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccGameliftFleetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGameliftFleet), &result)
	return &result
}
