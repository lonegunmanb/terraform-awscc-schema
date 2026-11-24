package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLambdaRuntimeAlias = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the runtime alias.",
        "description_kind": "plain",
        "type": "string"
      },
      "compatibility": {
        "computed": true,
        "description": "Compatibility matrix for the runtime version the alias points to.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "architectures": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            },
            "integration_models": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            },
            "managed_instances": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "created": {
        "computed": true,
        "description": "The timestamp when the runtime alias was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the runtime alias.",
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
      "last_modified": {
        "computed": true,
        "description": "The timestamp when the runtime alias was last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The alias name.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "release_lifecycle": {
        "computed": true,
        "description": "The publisher's lifecycle declaration for the alias. Omitting it removes any stored declaration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "deprecation_date": {
              "computed": true,
              "description": "The date on which the alias becomes deprecated. A UTC calendar date in yyyy-mm-dd format.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "function_create_block_date": {
              "computed": true,
              "description": "The date on which creating new functions on this alias is blocked. A UTC calendar date in yyyy-mm-dd format.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "function_update_block_date": {
              "computed": true,
              "description": "The date on which updating existing functions on this alias is blocked. A UTC calendar date in yyyy-mm-dd format.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "release_state": {
              "computed": true,
              "description": "The release state of the alias. DEPRECATED requires a DeprecationDate on or before today (UTC); PREVIEW and GA require a later DeprecationDate, or none. Reads report DEPRECATED once DeprecationDate is reached.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "runtime_name": {
        "description": "The name of the parent runtime.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "runtime_version": {
        "description": "The runtime version this alias points to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "runtime_version_arn": {
        "computed": true,
        "description": "The ARN of the runtime version this alias points to.",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The lifecycle state of the runtime alias.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Lambda::RuntimeAlias. A mutable named pointer to a specific runtime version.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLambdaRuntimeAliasSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLambdaRuntimeAlias), &result)
	return &result
}
