package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAutoscalingLifecycleHook = `{
  "block": {
    "attributes": {
      "auto_scaling_group_name": {
        "description": "The name of the Auto Scaling group for the lifecycle hook.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "default_result": {
        "computed": true,
        "description": "The action the Auto Scaling group takes when the lifecycle hook timeout elapses or if an unexpected failure occurs. The valid values are CONTINUE and ABANDON (default).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "heartbeat_timeout": {
        "computed": true,
        "description": "The maximum time, in seconds, that can elapse before the lifecycle hook times out. The range is from 30 to 7200 seconds. The default value is 3600 seconds (1 hour). If the lifecycle hook times out, Amazon EC2 Auto Scaling performs the action that you specified in the DefaultResult property.",
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
      "lifecycle_hook_name": {
        "computed": true,
        "description": "The name of the lifecycle hook.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "lifecycle_transition": {
        "description": "The instance state to which you want to attach the lifecycle hook.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "notification_metadata": {
        "computed": true,
        "description": "Additional information that is included any time Amazon EC2 Auto Scaling sends a message to the notification target.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "notification_target_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the notification target that Amazon EC2 Auto Scaling uses to notify you when an instance is in the transition state for the lifecycle hook. You can specify an Amazon SQS queue or an Amazon SNS topic. The notification message includes the following information: lifecycle action token, user account ID, Auto Scaling group name, lifecycle hook name, instance ID, lifecycle transition, and notification metadata.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description": "The ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target, for example, an Amazon SNS topic or an Amazon SQS queue.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::AutoScaling::LifecycleHook",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAutoscalingLifecycleHookSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAutoscalingLifecycleHook), &result)
	return &result
}
