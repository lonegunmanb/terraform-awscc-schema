package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLambdaCodeSigningConfig = `{
  "block": {
    "attributes": {
      "allowed_publishers": {
        "computed": true,
        "description": "When the CodeSigningConfig is later on attached to a function, the function code will be expected to be signed by profiles from this list",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "signing_profile_version_arns": {
              "computed": true,
              "description": "List of Signing profile version Arns",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        }
      },
      "code_signing_config_arn": {
        "computed": true,
        "description": "A unique Arn for CodeSigningConfig resource",
        "description_kind": "plain",
        "type": "string"
      },
      "code_signing_config_id": {
        "computed": true,
        "description": "A unique identifier for CodeSigningConfig resource",
        "description_kind": "plain",
        "type": "string"
      },
      "code_signing_policies": {
        "computed": true,
        "description": "Policies to control how to act if a signature is invalid",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "untrusted_artifact_on_deployment": {
              "computed": true,
              "description": "Indicates how Lambda operations involve updating the code artifact will operate. Default to Warn if not provided",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "description": {
        "computed": true,
        "description": "A description of the CodeSigningConfig",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of tags to apply to CodeSigningConfig resource",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::Lambda::CodeSigningConfig",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLambdaCodeSigningConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLambdaCodeSigningConfig), &result)
	return &result
}
