package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCustomerprofilesRecommenderFilter = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "The timestamp of when the recommender filter was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the recommender filter.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_name": {
        "computed": true,
        "description": "The unique name of the domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "recommender_filter_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the recommender filter.",
        "description_kind": "plain",
        "type": "string"
      },
      "recommender_filter_expression": {
        "computed": true,
        "description": "The filter expression that defines which items to include or exclude from recommendations.",
        "description_kind": "plain",
        "type": "string"
      },
      "recommender_filter_name": {
        "computed": true,
        "description": "The name of the recommender filter.",
        "description_kind": "plain",
        "type": "string"
      },
      "recommender_schema_name": {
        "computed": true,
        "description": "The name of the recommender schema to use for this recommender filter. If not specified, the default schema is used.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the recommender filter.",
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
      }
    },
    "description": "Data Source schema for AWS::CustomerProfiles::RecommenderFilter",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCustomerprofilesRecommenderFilterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCustomerprofilesRecommenderFilter), &result)
	return &result
}
