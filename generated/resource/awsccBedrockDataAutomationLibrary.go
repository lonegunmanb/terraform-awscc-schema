package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBedrockDataAutomationLibrary = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description": "Time Stamp",
        "description_kind": "plain",
        "type": "string"
      },
      "encryption_configuration": {
        "computed": true,
        "description": "KMS Encryption Configuration",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kms_encryption_context": {
              "computed": true,
              "description": "KMS Encryption Context",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "kms_key_id": {
              "computed": true,
              "description": "KMS Key Identifier",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "entity_types": {
        "computed": true,
        "description": "List of info for each entity type in the DataAutomationLibrary",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "entity_metadata": {
              "computed": true,
              "description": "JSON string representing relevant metadata for the entity type",
              "description_kind": "plain",
              "type": "string"
            },
            "entity_type": {
              "computed": true,
              "description": "Entity types supported in DataAutomationLibraries",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "library_arn": {
        "computed": true,
        "description": "ARN generated at the server side when a DataAutomationLibrary is created",
        "description_kind": "plain",
        "type": "string"
      },
      "library_description": {
        "computed": true,
        "description": "Description of the DataAutomationLibrary",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "library_name": {
        "description": "Name of the DataAutomationLibrary",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "Status of DataAutomationLibrary",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "List of tags",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "Tag key",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "Tag value",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::Bedrock::DataAutomationLibrary",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBedrockDataAutomationLibrarySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBedrockDataAutomationLibrary), &result)
	return &result
}
