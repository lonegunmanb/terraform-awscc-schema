package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudformationGuardHook = `{
  "block": {
    "attributes": {
      "alias": {
        "computed": true,
        "description": "The typename alias for the hook.",
        "description_kind": "plain",
        "type": "string"
      },
      "execution_role": {
        "computed": true,
        "description": "The execution role ARN assumed by hooks to read Guard rules from S3 and write Guard outputs to S3.",
        "description_kind": "plain",
        "type": "string"
      },
      "failure_mode": {
        "computed": true,
        "description": "Attribute to specify CloudFormation behavior on hook failure.",
        "description_kind": "plain",
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
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "log_bucket": {
        "computed": true,
        "description": "S3 Bucket where the guard validate report will be uploaded to",
        "description_kind": "plain",
        "type": "string"
      },
      "options": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "input_params": {
              "computed": true,
              "description": "S3 Source Location for the Guard files.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "uri": {
                    "computed": true,
                    "description": "S3 uri of Guard files.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "version_id": {
                    "computed": true,
                    "description": "S3 object version",
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
      "rule_location": {
        "computed": true,
        "description": "S3 Source Location for the Guard files.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "uri": {
              "computed": true,
              "description": "S3 uri of Guard files.",
              "description_kind": "plain",
              "type": "string"
            },
            "version_id": {
              "computed": true,
              "description": "S3 object version",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
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
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "include": {
                    "computed": true,
                    "description": "List of stack names that the hook is going to target",
                    "description_kind": "plain",
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              }
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
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "include": {
                    "computed": true,
                    "description": "List of stack roles that the hook is going to target",
                    "description_kind": "plain",
                    "type": [
                      "set",
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
              "type": [
                "set",
                "string"
              ]
            },
            "invocation_points": {
              "computed": true,
              "description": "List of invocation points that the hook is going to target",
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            },
            "target_names": {
              "computed": true,
              "description": "List of type names that the hook is going to target",
              "description_kind": "plain",
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
                    "type": "string"
                  },
                  "invocation_point": {
                    "computed": true,
                    "description": "Invocation points are the point in provisioning workflow where hooks will be executed.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "target_name": {
                    "computed": true,
                    "description": "Type name of hook target. Hook targets are the destination where hooks will be invoked against.",
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
      },
      "target_operations": {
        "computed": true,
        "description": "Which operations should this Hook run against? Resource changes, stacks or change sets.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::CloudFormation::GuardHook",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCloudformationGuardHookSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudformationGuardHook), &result)
	return &result
}
