package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSsmDocument = `{
  "block": {
    "attributes": {
      "attachments": {
        "computed": true,
        "description": "A list of key and value pairs that describe attachments to a version of a document.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key of a key-value pair that identifies the location of an attachment to a document.",
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "The name of the document attachment file.",
              "description_kind": "plain",
              "type": "string"
            },
            "values": {
              "computed": true,
              "description": "The value of a key-value pair that identifies the location of an attachment to a document. The format for Value depends on the type of key you specify.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "list"
        }
      },
      "content": {
        "computed": true,
        "description": "The content for the Systems Manager document in JSON, YAML or String format.",
        "description_kind": "plain",
        "type": "string"
      },
      "document_format": {
        "computed": true,
        "description": "Specify the document format for the request. The document format can be either JSON or YAML. JSON is the default format.",
        "description_kind": "plain",
        "type": "string"
      },
      "document_type": {
        "computed": true,
        "description": "The type of document to create.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "A name for the Systems Manager document.",
        "description_kind": "plain",
        "type": "string"
      },
      "requires": {
        "computed": true,
        "description": "A list of SSM documents required by a document. For example, an ApplicationConfiguration document requires an ApplicationConfigurationSchema document.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
              "computed": true,
              "description": "The name of the required SSM document. The name can be an Amazon Resource Name (ARN).",
              "description_kind": "plain",
              "type": "string"
            },
            "version": {
              "computed": true,
              "description": "The document version required by the current document.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "tags": {
        "computed": true,
        "description": "Optional metadata that you assign to a resource. Tags enable you to categorize a resource in different ways, such as by purpose, owner, or environment.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The name of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "target_type": {
        "computed": true,
        "description": "Specify a target type to define the kinds of resources the document can run on.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_method": {
        "computed": true,
        "description": "Update method - when set to 'Replace', the update will replace the existing document; when set to 'NewVersion', the update will create a new version.",
        "description_kind": "plain",
        "type": "string"
      },
      "version_name": {
        "computed": true,
        "description": "An optional field specifying the version of the artifact you are creating with the document. This value is unique across all versions of a document, and cannot be changed.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SSM::Document",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSsmDocumentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSsmDocument), &result)
	return &result
}
