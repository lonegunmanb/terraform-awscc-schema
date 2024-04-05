package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApplicationautoscalingScalableTarget = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "max_capacity": {
        "computed": true,
        "description": "The maximum value that you plan to scale in to. When a scaling policy is in effect, Application Auto Scaling can scale in (contract) as needed to the minimum capacity limit in response to changing demand",
        "description_kind": "plain",
        "type": "number"
      },
      "min_capacity": {
        "computed": true,
        "description": "The minimum value that you plan to scale in to. When a scaling policy is in effect, Application Auto Scaling can scale in (contract) as needed to the minimum capacity limit in response to changing demand",
        "description_kind": "plain",
        "type": "number"
      },
      "resource_id": {
        "computed": true,
        "description": "The identifier of the resource associated with the scalable target",
        "description_kind": "plain",
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description": "Specify the Amazon Resource Name (ARN) of an Identity and Access Management (IAM) role that allows Application Auto Scaling to modify the scalable target on your behalf. ",
        "description_kind": "plain",
        "type": "string"
      },
      "scalable_dimension": {
        "computed": true,
        "description": "The scalable dimension associated with the scalable target. This string consists of the service namespace, resource type, and scaling property",
        "description_kind": "plain",
        "type": "string"
      },
      "scalable_target_id": {
        "computed": true,
        "description": "This value can be returned by using the Ref function. Ref returns the Cloudformation generated ID of the resource in format - ResourceId|ScalableDimension|ServiceNamespace",
        "description_kind": "plain",
        "type": "string"
      },
      "scheduled_actions": {
        "computed": true,
        "description": "The scheduled actions for the scalable target. Duplicates aren't allowed.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "end_time": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "scalable_target_action": {
              "computed": true,
              "description": "specifies the minimum and maximum capacity",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "max_capacity": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "min_capacity": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "schedule": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "scheduled_action_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "start_time": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "timezone": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "service_namespace": {
        "computed": true,
        "description": "The namespace of the AWS service that provides the resource, or a custom-resource",
        "description_kind": "plain",
        "type": "string"
      },
      "suspended_state": {
        "computed": true,
        "description": "An embedded object that contains attributes and attribute values that are used to suspend and resume automatic scaling. Setting the value of an attribute to true suspends the specified scaling activities. Setting it to false (default) resumes the specified scaling activities.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "dynamic_scaling_in_suspended": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "dynamic_scaling_out_suspended": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "scheduled_scaling_suspended": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::ApplicationAutoScaling::ScalableTarget",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccApplicationautoscalingScalableTargetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApplicationautoscalingScalableTarget), &result)
	return &result
}
