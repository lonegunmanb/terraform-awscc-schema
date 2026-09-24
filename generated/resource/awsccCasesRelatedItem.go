package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCasesRelatedItem = `{
  "block": {
    "attributes": {
      "case_id": {
        "description": "A unique identifier of the case.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "content": {
        "description": "The content of a related item to be created.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "comment": {
              "computed": true,
              "description": "Represents a comment.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "body": {
                    "computed": true,
                    "description": "Text in the body of a comment.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "content_type": {
                    "computed": true,
                    "description": "Type of the text in the comment.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "domain_id": {
        "description": "The unique identifier of the Cases domain.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "related_item_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the related item.",
        "description_kind": "plain",
        "type": "string"
      },
      "related_item_id": {
        "computed": true,
        "description": "The unique identifier of the related item.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of tags on the related item.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "type": {
        "description": "The type of a related item.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Cases::RelatedItem. Creates a related item (comments) and associates it with a case.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCasesRelatedItemSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCasesRelatedItem), &result)
	return &result
}
