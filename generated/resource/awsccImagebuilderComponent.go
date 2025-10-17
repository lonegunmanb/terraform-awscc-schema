package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccImagebuilderComponent = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the component.",
        "description_kind": "plain",
        "type": "string"
      },
      "change_description": {
        "computed": true,
        "description": "The change description of the component.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "data": {
        "computed": true,
        "description": "The data of the component.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the component.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "encrypted": {
        "computed": true,
        "description": "The encryption status of the component.",
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description": "The KMS key identifier used to encrypt the component.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "latest_version": {
        "computed": true,
        "description": "The latest version references of the component.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "arn": {
              "computed": true,
              "description": "The latest version ARN of the created component.",
              "description_kind": "plain",
              "type": "string"
            },
            "major": {
              "computed": true,
              "description": "The latest version ARN of the created component, with the same major version.",
              "description_kind": "plain",
              "type": "string"
            },
            "minor": {
              "computed": true,
              "description": "The latest version ARN of the created component, with the same minor version.",
              "description_kind": "plain",
              "type": "string"
            },
            "patch": {
              "computed": true,
              "description": "The latest version ARN of the created component, with the same patch version.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "name": {
        "description": "The name of the component.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "platform": {
        "description": "The platform of the component.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "supported_os_versions": {
        "computed": true,
        "description": "The operating system (OS) version supported by the component.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "The tags associated with the component.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "type": {
        "computed": true,
        "description": "The type of the component denotes whether the component is used to build the image or only to test it. ",
        "description_kind": "plain",
        "type": "string"
      },
      "uri": {
        "computed": true,
        "description": "The uri of the component.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "version": {
        "description": "The version of the component.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource schema for AWS::ImageBuilder::Component",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccImagebuilderComponentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccImagebuilderComponent), &result)
	return &result
}
