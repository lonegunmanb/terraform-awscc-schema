package data

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
        "computed": true,
        "description": "The ARN of the Domain dataset group that the recommender is created for.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_updated_date_time": {
        "computed": true,
        "description": "The time at which the recommender was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the recommender.",
        "description_kind": "plain",
        "type": "string"
      },
      "recipe_arn": {
        "computed": true,
        "description": "The ARN of the recipe (Domain dataset group use case) that the recommender uses.",
        "description_kind": "plain",
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
              "type": "bool"
            },
            "item_exploration_config": {
              "computed": true,
              "description": "Specifies the exploration configuration hyperparameters.",
              "description_kind": "plain",
              "type": [
                "map",
                "string"
              ]
            },
            "min_recommendation_requests_per_second": {
              "computed": true,
              "description": "Specifies the requested minimum provisioned recommendation requests per second.",
              "description_kind": "plain",
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
              }
            }
          },
          "nesting_mode": "single"
        }
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
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::Personalize::Recommender",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccPersonalizeRecommenderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPersonalizeRecommender), &result)
	return &result
}
