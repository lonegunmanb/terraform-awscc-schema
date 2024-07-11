package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudformationStackSet = `{
  "block": {
    "attributes": {
      "administration_role_arn": {
        "computed": true,
        "description": "The Amazon Resource Number (ARN) of the IAM role to use to create this stack set. Specify an IAM role only if you are using customized administrator roles to control which users or groups can manage specific stack sets within the same administrator account.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "auto_deployment": {
        "computed": true,
        "description": "Describes whether StackSets automatically deploys to AWS Organizations accounts that are added to the target organization or organizational unit (OU). Specify only if PermissionModel is SERVICE_MANAGED.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enabled": {
              "computed": true,
              "description": "If set to true, StackSets automatically deploys additional stack instances to AWS Organizations accounts that are added to a target organization or organizational unit (OU) in the specified Regions. If an account is removed from a target organization or OU, StackSets deletes stack instances from the account in the specified Regions.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "retain_stacks_on_account_removal": {
              "computed": true,
              "description": "If set to true, stack resources are retained when an account is removed from a target organization or OU. If set to false, stack resources are deleted. Specify only if Enabled is set to True.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "call_as": {
        "computed": true,
        "description": "Specifies the AWS account that you are acting from. By default, SELF is specified. For self-managed permissions, specify SELF; for service-managed permissions, if you are signed in to the organization's management account, specify SELF. If you are signed in to a delegated administrator account, specify DELEGATED_ADMIN.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "capabilities": {
        "computed": true,
        "description": "In some cases, you must explicitly acknowledge that your stack set template contains certain capabilities in order for AWS CloudFormation to create the stack set and related stack instances.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "description": {
        "computed": true,
        "description": "A description of the stack set. You can use the description to identify the stack set's purpose or other important information.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "execution_role_name": {
        "computed": true,
        "description": "The name of the IAM execution role to use to create the stack set. If you do not specify an execution role, AWS CloudFormation uses the AWSCloudFormationStackSetExecutionRole role for the stack set operation.",
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
      "managed_execution": {
        "computed": true,
        "description": "Describes whether StackSets performs non-conflicting operations concurrently and queues conflicting operations.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "active": {
              "computed": true,
              "description": "When true, StackSets performs non-conflicting operations concurrently and queues conflicting operations. After conflicting operations finish, StackSets starts queued operations in request order.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "operation_preferences": {
        "computed": true,
        "description": "The user-specified preferences for how AWS CloudFormation performs a stack set operation.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "concurrency_mode": {
              "computed": true,
              "description": "Specifies how the concurrency level behaves during the operation execution.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "failure_tolerance_count": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "failure_tolerance_percentage": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_concurrent_count": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_concurrent_percentage": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "region_concurrency_type": {
              "computed": true,
              "description": "The concurrency type of deploying StackSets operations in regions, could be in parallel or one region at a time",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "region_order": {
              "computed": true,
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
      },
      "parameters": {
        "computed": true,
        "description": "The input parameters for the stack set template.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "parameter_key": {
              "description": "The key associated with the parameter. If you don't specify a key and value for a particular parameter, AWS CloudFormation uses the default value that is specified in your template.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "parameter_value": {
              "description": "The input value associated with the parameter.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "permission_model": {
        "description": "Describes how the IAM roles required for stack set operations are created. By default, SELF-MANAGED is specified.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "stack_instances_group": {
        "computed": true,
        "description": "A group of stack instances with parameters in some specific accounts and regions.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "deployment_targets": {
              "description": " The AWS OrganizationalUnitIds or Accounts for which to create stack instances in the specified Regions.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "account_filter_type": {
                    "computed": true,
                    "description": "The filter type you want to apply on organizational units and accounts.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "accounts": {
                    "computed": true,
                    "description": "AWS accounts that you want to create stack instances in the specified Region(s) for.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "accounts_url": {
                    "computed": true,
                    "description": "Returns the value of the AccountsUrl property.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "organizational_unit_ids": {
                    "computed": true,
                    "description": "The organization root ID or organizational unit (OU) IDs to which StackSets deploys.",
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
              "required": true
            },
            "parameter_overrides": {
              "computed": true,
              "description": "A list of stack set parameters whose values you want to override in the selected stack instances.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "parameter_key": {
                    "description": "The key associated with the parameter. If you don't specify a key and value for a particular parameter, AWS CloudFormation uses the default value that is specified in your template.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "parameter_value": {
                    "description": "The input value associated with the parameter.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "set"
              },
              "optional": true
            },
            "regions": {
              "description": "The names of one or more Regions where you want to create stack instances using the specified AWS account(s).",
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "stack_set_id": {
        "computed": true,
        "description": "The ID of the stack set that you're creating.",
        "description_kind": "plain",
        "type": "string"
      },
      "stack_set_name": {
        "description": "The name to associate with the stack set. The name must be unique in the Region where you create your stack set.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The key-value pairs to associate with this stack set and the stacks created from it. AWS CloudFormation also propagates these tags to supported resources that are created in the stacks. A maximum number of 50 tags can be specified.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "A string used to identify this tag. You can specify a maximum of 127 characters for a tag key.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "A string containing the value for this tag. You can specify a maximum of 256 characters for a tag value.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "template_body": {
        "computed": true,
        "description": "The structure that contains the template body, with a minimum length of 1 byte and a maximum length of 51,200 bytes.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "template_url": {
        "computed": true,
        "description": "Location of file containing the template body. The URL must point to a template (max size: 460,800 bytes) that is located in an Amazon S3 bucket.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "StackSet as a resource provides one-click experience for provisioning a StackSet and StackInstances",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudformationStackSetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudformationStackSet), &result)
	return &result
}
