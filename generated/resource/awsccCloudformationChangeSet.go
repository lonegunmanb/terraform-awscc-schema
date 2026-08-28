package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudformationChangeSet = `{
  "block": {
    "attributes": {
      "capabilities": {
        "computed": true,
        "description": "The capabilities that are allowed in the stack.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "change_set_id": {
        "computed": true,
        "description": "The ARN of the change set.",
        "description_kind": "plain",
        "type": "string"
      },
      "change_set_name": {
        "description": "The name of the change set. Must be unique among all change sets associated with the specified stack.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "change_set_type": {
        "computed": true,
        "description": "The type of change set operation.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description": "The time the change set was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "deployment_mode": {
        "computed": true,
        "description": "Determines how CloudFormation handles configuration drift during deployment.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description to help you identify this change set.",
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
      "import_existing_resources": {
        "computed": true,
        "description": "Indicates if the change set imports resources that already exist.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "include_nested_stacks": {
        "computed": true,
        "description": "Creates a change set for all nested stacks specified in the template.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "notification_ar_ns": {
        "computed": true,
        "description": "The ARNs of Amazon SNS topics that CloudFormation associates with the stack.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "on_stack_failure": {
        "computed": true,
        "description": "Determines what action will be taken if stack creation fails.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description": "The ARN of an IAM role that CloudFormation assumes when executing the change set.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "stack_id": {
        "computed": true,
        "description": "The unique ID of the stack.",
        "description_kind": "plain",
        "type": "string"
      },
      "stack_name": {
        "description": "The name or unique ID of the stack for which you are creating a change set.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Key-value pairs to associate with the change set.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "template_body": {
        "computed": true,
        "description": "A structure that contains the body of the revised template.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "template_url": {
        "computed": true,
        "description": "The URL of the file that contains the revised template.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "use_previous_template": {
        "computed": true,
        "description": "Whether to reuse the template associated with the stack to create the change set.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "description": "Resource type definition for AWS::CloudFormation::ChangeSet",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudformationChangeSetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudformationChangeSet), &result)
	return &result
}
