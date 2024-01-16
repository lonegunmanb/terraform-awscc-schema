package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLambdaCodeSigningConfig = `{
  "block": {
    "attributes": {
      "allowed_publishers": {
        "description": "When the CodeSigningConfig is later on attached to a function, the function code will be expected to be signed by profiles from this list",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "signing_profile_version_arns": {
              "description": "List of Signing profile version Arns",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        },
        "required": true
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
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "description": {
        "computed": true,
        "description": "A description of the CodeSigningConfig",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Lambda::CodeSigningConfig.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLambdaCodeSigningConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLambdaCodeSigningConfig), &result)
	return &result
}
