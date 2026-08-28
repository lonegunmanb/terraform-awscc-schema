package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccKendraQuerySuggestionsBlockList = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the query suggestions block list.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description for the block list.",
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
      "index_id": {
        "description": "The identifier of the index for the block list.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the block list.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "query_suggestions_block_list_id": {
        "computed": true,
        "description": "The identifier of the block list.",
        "description_kind": "plain",
        "type": "string"
      },
      "role_arn": {
        "description": "The Amazon Resource Name (ARN) of an IAM role with permission to access the S3 bucket that contains the block list text file.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_s3_path": {
        "description": "Information required to find a specific file in an Amazon S3 bucket.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "bucket": {
              "description": "The name of the S3 bucket that contains the file.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "key": {
              "description": "The name of the file.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "tags": {
        "computed": true,
        "description": "A list of key-value pairs that identify or categorize the block list.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key for the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value associated with the tag.",
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
    "description": "A block list used for query suggestions for an Amazon Kendra index. A block list contains words or phrases that should not appear as query suggestions.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccKendraQuerySuggestionsBlockListSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccKendraQuerySuggestionsBlockList), &result)
	return &result
}
