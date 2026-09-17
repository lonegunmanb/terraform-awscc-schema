package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccKendraFeaturedResultsSet = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the featured results set.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description for the set of featured results.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "featured_documents": {
        "computed": true,
        "description": "A list of document IDs for the documents you want to feature at the top of the search results page.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "id": {
              "computed": true,
              "description": "The identifier of the document to feature in the search results.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "featured_results_set_id": {
        "computed": true,
        "description": "The identifier of the set of featured results.",
        "description_kind": "plain",
        "type": "string"
      },
      "featured_results_set_name": {
        "description": "A name for the set of featured results.",
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
      "index_id": {
        "description": "The identifier of the index that you want to use for featuring results.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "query_texts": {
        "computed": true,
        "description": "A list of queries for featuring results.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "status": {
        "computed": true,
        "description": "The current status of the set of featured results. When the value is ACTIVE, featured results are ready for use.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of key-value pairs that identify or categorize the featured results set.",
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
    "description": "Resource Type definition for AWS::Kendra::FeaturedResultsSet. A set of featured results that are displayed at the top of your search results. Featured results are placed above all other results for certain queries. If there's an exact match of a query, then one or more specific documents are featured in the search results.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccKendraFeaturedResultsSetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccKendraFeaturedResultsSet), &result)
	return &result
}
