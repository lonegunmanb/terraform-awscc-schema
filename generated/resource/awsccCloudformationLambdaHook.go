package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudformationLambdaHook = `{
  "block": {
    "attributes": {
      "alias": {
        "description": "The typename alias for the hook.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "execution_role": {
        "description": "The execution role ARN assumed by Hooks to invoke Lambda.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "failure_mode": {
        "description": "Attribute to specify CloudFormation behavior on hook failure.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "hook_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the activated hook",
        "description_kind": "plain",
        "type": "string"
      },
      "hook_status": {
        "computed": true,
        "description": "Attribute to specify which stacks this hook applies to or should get invoked for",
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
      "lambda_function": {
        "description": "Amazon Resource Name (ARN), Partial ARN, name, version, or alias of the Lambda function to invoke with this hook.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "stack_filters": {
        "computed": true,
        "description": "Filters to allow hooks to target specific stack attributes",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "filtering_criteria": {
              "computed": true,
              "description": "Attribute to specify the filtering behavior. ANY will make the Hook pass if one filter matches. ALL will make the Hook pass if all filters match",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "stack_names": {
              "computed": true,
              "description": "List of stack names as filters",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "exclude": {
                    "computed": true,
                    "description": "List of stack names that the hook is going to be excluded from",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "include": {
                    "computed": true,
                    "description": "List of stack names that the hook is going to target",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "stack_roles": {
              "computed": true,
              "description": "List of stack roles that are performing the stack operations.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "exclude": {
                    "computed": true,
                    "description": "List of stack roles that the hook is going to be excluded from",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "include": {
                    "computed": true,
                    "description": "List of stack roles that the hook is going to target",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
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
      "target_filters": {
        "computed": true,
        "description": "Attribute to specify which targets should invoke the hook",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "actions": {
              "computed": true,
              "description": "List of actions that the hook is going to target",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "invocation_points": {
              "computed": true,
              "description": "List of invocation points that the hook is going to target",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "target_names": {
              "computed": true,
              "description": "List of type names that the hook is going to target",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "targets": {
              "computed": true,
              "description": "List of hook targets",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "action": {
                    "computed": true,
                    "description": "Target actions are the type of operation hooks will be executed at.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "invocation_point": {
                    "computed": true,
                    "description": "Invocation points are the point in provisioning workflow where hooks will be executed.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "target_name": {
                    "computed": true,
                    "description": "Type name of hook target. Hook targets are the destination where hooks will be invoked against.",
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
      "target_operations": {
        "description": "Which operations should this Hook run against? Resource changes, stacks or change sets.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "This is a CloudFormation resource for the first-party AWS::Hooks::LambdaHook.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudformationLambdaHookSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudformationLambdaHook), &result)
	return &result
}
