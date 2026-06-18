package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockagentcoreConfigurationBundle = `{
  "block": {
    "attributes": {
      "branch_name": {
        "computed": true,
        "description": "The branch name for version tracking.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "bundle_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the configuration bundle.",
        "description_kind": "plain",
        "type": "string"
      },
      "bundle_id": {
        "computed": true,
        "description": "The unique identifier of the configuration bundle.",
        "description_kind": "plain",
        "type": "string"
      },
      "bundle_name": {
        "description": "The name for the configuration bundle. Names must be unique within your account.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "commit_message": {
        "computed": true,
        "description": "A commit message describing the version of the configuration bundle.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "components": {
        "description": "A map of component identifiers to their configurations.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "configuration": {
              "computed": true,
              "description": "The configuration values as a flexible JSON document.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "map"
        },
        "required": true
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp when the configuration bundle was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_by": {
        "computed": true,
        "description": "The source that created a configuration bundle version.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the source, if applicable.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "The name of the source (for example, user, optimization-job, or system).",
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
        "description": "The description for the configuration bundle.",
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
      "kms_key_arn": {
        "computed": true,
        "description": "The ARN of the KMS key used to encrypt component configurations.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "lineage_metadata": {
        "computed": true,
        "description": "The version lineage metadata that tracks parent versions and creation source.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "branch_name": {
              "computed": true,
              "description": "The branch name for this version.",
              "description_kind": "plain",
              "type": "string"
            },
            "commit_message": {
              "computed": true,
              "description": "A commit message describing the changes in this version.",
              "description_kind": "plain",
              "type": "string"
            },
            "created_by": {
              "computed": true,
              "description": "The source that created a configuration bundle version.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of the source, if applicable.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "The name of the source (for example, user, optimization-job, or system).",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "parent_version_ids": {
              "computed": true,
              "description": "A list of parent version identifiers.",
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
      "tags": {
        "computed": true,
        "description": "Tags to assign to the configuration bundle.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp when the configuration bundle was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "version_id": {
        "computed": true,
        "description": "The version identifier of the configuration bundle.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Definition of AWS::BedrockAgentCore::ConfigurationBundle Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBedrockagentcoreConfigurationBundleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockagentcoreConfigurationBundle), &result)
	return &result
}
