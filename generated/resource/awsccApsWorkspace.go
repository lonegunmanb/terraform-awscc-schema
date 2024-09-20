package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApsWorkspace = `{
  "block": {
    "attributes": {
      "alert_manager_definition": {
        "computed": true,
        "description": "The AMP Workspace alert manager definition data",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "alias": {
        "computed": true,
        "description": "AMP Workspace alias.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "Workspace arn.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_arn": {
        "computed": true,
        "description": "KMS Key ARN used to encrypt and decrypt AMP workspace data.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "logging_configuration": {
        "computed": true,
        "description": "Logging configuration",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "log_group_arn": {
              "computed": true,
              "description": "CloudWatch log group ARN",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "prometheus_endpoint": {
        "computed": true,
        "description": "AMP Workspace prometheus endpoint",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "workspace_id": {
        "computed": true,
        "description": "Required to identify a specific APS Workspace.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::APS::Workspace",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApsWorkspaceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApsWorkspace), &result)
	return &result
}
