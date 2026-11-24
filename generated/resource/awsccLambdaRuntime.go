package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLambdaRuntime = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the runtime.",
        "description_kind": "plain",
        "type": "string"
      },
      "created": {
        "computed": true,
        "description": "The timestamp when the runtime was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the runtime.",
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
        "description": "The timestamp when the runtime was last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "license_info": {
        "computed": true,
        "description": "Licensing information for the runtime.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the runtime.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "release_lifecycle": {
        "computed": true,
        "description": "The publisher's lifecycle declaration for the runtime. Omitting it removes any stored declaration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "deprecation_date": {
              "computed": true,
              "description": "The date on which the runtime becomes deprecated. A UTC calendar date in yyyy-mm-dd format.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "function_create_block_date": {
              "computed": true,
              "description": "The date on which creating new functions on this runtime is blocked. A UTC calendar date in yyyy-mm-dd format.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "function_update_block_date": {
              "computed": true,
              "description": "The date on which updating existing functions on this runtime is blocked. A UTC calendar date in yyyy-mm-dd format.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "release_state": {
              "computed": true,
              "description": "The release state of the runtime. DEPRECATED requires a DeprecationDate on or before today (UTC); PREVIEW and GA require a later DeprecationDate, or none. Reads report DEPRECATED once DeprecationDate is reached.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "state": {
        "computed": true,
        "description": "The lifecycle state of the runtime.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of tags to apply to the runtime.",
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
    "description": "Resource Type definition for AWS::Lambda::Runtime. A Lambda runtime backed by an OCI image, used by Lambda functions.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLambdaRuntimeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLambdaRuntime), &result)
	return &result
}
