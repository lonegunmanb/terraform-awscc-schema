package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAmplifyBranch = `{
  "block": {
    "attributes": {
      "app_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "backend": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "stack_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "basic_auth_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enable_basic_auth": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "password": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "username": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "branch_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "build_spec": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "compute_role_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enable_auto_build": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_performance_mode": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_pull_request_preview": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_skew_protection": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "environment_variables": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "framework": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "pull_request_environment_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "stage": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::Amplify::Branch",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAmplifyBranchSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAmplifyBranch), &result)
	return &result
}
