package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccPersonalizeRecommender = `{
  "block": {
    "attributes": {
      "creation_date_time": {
        "computed": true,
        "description": "The time at which the recommender was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "dataset_group_arn": {
        "description": "The ARN of the Domain dataset group that the recommender is created for.",
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
      "last_updated_date_time": {
        "computed": true,
        "description": "The time at which the recommender was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the recommender.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "recipe_arn": {
        "description": "The ARN of the recipe (Domain dataset group use case) that the recommender uses.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "recommender_arn": {
        "computed": true,
        "description": "The ARN of the recommender.",
        "description_kind": "plain",
        "type": "string"
      },
      "recommender_config": {
        "computed": true,
        "description": "The configuration details of the recommender.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enable_metadata_with_recommendations": {
              "computed": true,
              "description": "Whether metadata with recommendations is enabled for the recommender.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "item_exploration_config": {
              "computed": true,
              "description": "Specifies the exploration configuration hyperparameters.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "min_recommendation_requests_per_second": {
              "computed": true,
              "description": "Specifies the requested minimum provisioned recommendation requests per second.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "training_data_config": {
              "computed": true,
              "description": "The training data configuration to use when creating a domain recommender.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "excluded_dataset_columns": {
                    "computed": true,
                    "description": "Specifies the columns to exclude from training. Each key is a dataset type, and each value is a list of columns.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      [
                        "list",
                        "string"
                      ]
                    ]
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "status": {
        "computed": true,
        "description": "The status of the recommender.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags to associate with the recommender.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "A recommender for a Domain dataset group that generates recommendations using a domain use case recipe.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccPersonalizeRecommenderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPersonalizeRecommender), &result)
	return &result
}
