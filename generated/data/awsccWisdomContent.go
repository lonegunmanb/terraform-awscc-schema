package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWisdomContent = `{
  "block": {
    "attributes": {
      "content_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the content.",
        "description_kind": "plain",
        "type": "string"
      },
      "content_id": {
        "computed": true,
        "description": "The identifier of the content.",
        "description_kind": "plain",
        "type": "string"
      },
      "content_type": {
        "computed": true,
        "description": "The media type of the content.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "knowledge_base_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the knowledge base.",
        "description_kind": "plain",
        "type": "string"
      },
      "knowledge_base_id": {
        "computed": true,
        "description": "The identifier of the knowledge base.",
        "description_kind": "plain",
        "type": "string"
      },
      "link_out_uri": {
        "computed": true,
        "description": "The URI of the content.",
        "description_kind": "plain",
        "type": "string"
      },
      "metadata": {
        "computed": true,
        "description": "A key/value map to store attributes without affecting tagging or recommendations.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "computed": true,
        "description": "The name of the content.",
        "description_kind": "plain",
        "type": "string"
      },
      "override_link_out_uri": {
        "computed": true,
        "description": "The URI you want to use for the article.",
        "description_kind": "plain",
        "type": "string"
      },
      "revision_id": {
        "computed": true,
        "description": "The identifier of the content revision.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the content.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags used to organize, track, or control access for this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "title": {
        "computed": true,
        "description": "The title of the content.",
        "description_kind": "plain",
        "type": "string"
      },
      "upload_id": {
        "computed": true,
        "description": "A pointer to the uploaded asset. This value is returned by StartContentUpload.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Wisdom::Content",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccWisdomContentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWisdomContent), &result)
	return &result
}
