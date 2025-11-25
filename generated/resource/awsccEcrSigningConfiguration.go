package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEcrSigningConfiguration = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "registry_id": {
        "computed": true,
        "description": "12-digit AWS account ID of the ECR registry.",
        "description_kind": "plain",
        "type": "string"
      },
      "rules": {
        "description": "Array of signing rules that define which repositories should be signed and with which signing profiles.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "repository_filters": {
              "computed": true,
              "description": "Optional array of repository filters. If omitted, the rule matches all repositories. If provided, must contain at least one filter. Empty arrays are not allowed.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "filter": {
                    "computed": true,
                    "description": "Repository name pattern (supports '*' wildcard).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "filter_type": {
                    "computed": true,
                    "description": "Type of repository filter",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "signing_profile_arn": {
              "description": "AWS Signer signing profile ARN to use for matched repositories.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      }
    },
    "description": "The AWS::ECR::SigningConfiguration resource creates or updates the signing configuration for an Amazon ECR registry.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEcrSigningConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEcrSigningConfiguration), &result)
	return &result
}
