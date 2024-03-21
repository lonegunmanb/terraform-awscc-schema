package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGameliftGameServerGroup = `{
  "block": {
    "attributes": {
      "auto_scaling_group_arn": {
        "computed": true,
        "description": "A generated unique ID for the EC2 Auto Scaling group that is associated with this game server group.",
        "description_kind": "plain",
        "type": "string"
      },
      "auto_scaling_policy": {
        "computed": true,
        "description": "Configuration settings to define a scaling policy for the Auto Scaling group that is optimized for game hosting. Updating this game server group property will not take effect for the created EC2 Auto Scaling group, please update the EC2 Auto Scaling group directly after creating the resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "estimated_instance_warmup": {
              "computed": true,
              "description": "Length of time, in seconds, it takes for a new instance to start new game server processes and register with GameLift FleetIQ.",
              "description_kind": "plain",
              "type": "number"
            },
            "target_tracking_configuration": {
              "computed": true,
              "description": "Settings for a target-based scaling policy applied to Auto Scaling group.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "target_value": {
                    "computed": true,
                    "description": "Desired value to use with a game server group target-based scaling policy.",
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
      "balancing_strategy": {
        "computed": true,
        "description": "The fallback balancing method to use for the game server group when Spot Instances in a Region become unavailable or are not viable for game hosting.",
        "description_kind": "plain",
        "type": "string"
      },
      "delete_option": {
        "computed": true,
        "description": "The type of delete to perform.",
        "description_kind": "plain",
        "type": "string"
      },
      "game_server_group_arn": {
        "computed": true,
        "description": "A generated unique ID for the game server group.",
        "description_kind": "plain",
        "type": "string"
      },
      "game_server_group_name": {
        "computed": true,
        "description": "An identifier for the new game server group.",
        "description_kind": "plain",
        "type": "string"
      },
      "game_server_protection_policy": {
        "computed": true,
        "description": "A flag that indicates whether instances in the game server group are protected from early termination.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_definitions": {
        "computed": true,
        "description": "A set of EC2 instance types to use when creating instances in the group.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "instance_type": {
              "computed": true,
              "description": "An EC2 instance type designation.",
              "description_kind": "plain",
              "type": "string"
            },
            "weighted_capacity": {
              "computed": true,
              "description": "Instance weighting that indicates how much this instance type contributes to the total capacity of a game server group.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "launch_template": {
        "computed": true,
        "description": "The EC2 launch template that contains configuration settings and game server code to be deployed to all instances in the game server group. Updating this game server group property will not take effect for the created EC2 Auto Scaling group, please update the EC2 Auto Scaling group directly after creating the resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "launch_template_id": {
              "computed": true,
              "description": "A unique identifier for an existing EC2 launch template.",
              "description_kind": "plain",
              "type": "string"
            },
            "launch_template_name": {
              "computed": true,
              "description": "A readable identifier for an existing EC2 launch template.",
              "description_kind": "plain",
              "type": "string"
            },
            "version": {
              "computed": true,
              "description": "The version of the EC2 launch template to use.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "max_size": {
        "computed": true,
        "description": "The maximum number of instances allowed in the EC2 Auto Scaling group. Updating this game server group property will not take effect for the created EC2 Auto Scaling group, please update the EC2 Auto Scaling group directly after creating the resource.",
        "description_kind": "plain",
        "type": "number"
      },
      "min_size": {
        "computed": true,
        "description": "The minimum number of instances allowed in the EC2 Auto Scaling group. Updating this game server group property will not take effect for the created EC2 Auto Scaling group, please update the EC2 Auto Scaling group directly after creating the resource.",
        "description_kind": "plain",
        "type": "number"
      },
      "role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for an IAM role that allows Amazon GameLift to access your EC2 Auto Scaling groups.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of labels to assign to the new game server group resource. Updating game server group tags with CloudFormation will not take effect. Please update this property using AWS GameLift APIs instead.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key for a developer-defined key:value pair for tagging an AWS resource.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for a developer-defined key:value pair for tagging an AWS resource.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "vpc_subnets": {
        "computed": true,
        "description": "A list of virtual private cloud (VPC) subnets to use with instances in the game server group. Updating this game server group property will not take effect for the created EC2 Auto Scaling group, please update the EC2 Auto Scaling group directly after creating the resource.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::GameLift::GameServerGroup",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccGameliftGameServerGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGameliftGameServerGroup), &result)
	return &result
}
