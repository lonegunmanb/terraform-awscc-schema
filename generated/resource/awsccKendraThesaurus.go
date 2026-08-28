package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccKendraThesaurus = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the thesaurus.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description for the thesaurus.",
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
        "description": "The identifier of the index for the thesaurus.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "A name for the thesaurus.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "role_arn": {
        "description": "An IAM role that gives Amazon Kendra permissions to access the thesaurus file specified in SourceS3Path.",
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
        "description": "A list of key-value pairs that identify or categorize the thesaurus.",
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
      },
      "thesaurus_id": {
        "computed": true,
        "description": "The identifier of the thesaurus.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "A thesaurus for an Amazon Kendra index. The thesaurus contains a list of synonyms in Solr format.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccKendraThesaurusSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccKendraThesaurus), &result)
	return &result
}
