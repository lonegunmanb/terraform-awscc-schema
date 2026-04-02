package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCustomerprofilesRecommender = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "The timestamp of when the recommender was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the recommender.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "domain_name": {
        "description": "The name of the domain for which the recommender will be created",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "failure_reason": {
        "computed": true,
        "description": "The reason for recommender failure.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_updated_at": {
        "computed": true,
        "description": "The timestamp of when the recommender was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "latest_recommender_update": {
        "computed": true,
        "description": "Information about the latest recommender update",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "creation_date_time": {
              "computed": true,
              "description": "The timestamp of when the update was created",
              "description_kind": "plain",
              "type": "string"
            },
            "failure_reason": {
              "computed": true,
              "description": "The reason for update failure",
              "description_kind": "plain",
              "type": "string"
            },
            "last_updated_date_time": {
              "computed": true,
              "description": "The timestamp of when the update was last modified",
              "description_kind": "plain",
              "type": "string"
            },
            "recommender_config": {
              "computed": true,
              "description": "Configuration for the recommender",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "events_config": {
                    "computed": true,
                    "description": "Configuration for events used in the recommender",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "event_parameters_list": {
                          "computed": true,
                          "description": "List of event parameters with their value thresholds",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "event_type": {
                                "computed": true,
                                "description": "The type of event",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "event_value_threshold": {
                                "computed": true,
                                "description": "The threshold of the event type. Only events with a value greater or equal to this threshold will be considered for solution creation.",
                                "description_kind": "plain",
                                "type": "number"
                              }
                            },
                            "nesting_mode": "list"
                          }
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
              "description": "The status of the recommender",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "recommender_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the recommender.",
        "description_kind": "plain",
        "type": "string"
      },
      "recommender_config": {
        "computed": true,
        "description": "Configuration for the recommender",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "events_config": {
              "computed": true,
              "description": "Configuration for events used in the recommender",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "event_parameters_list": {
                    "computed": true,
                    "description": "List of event parameters with their value thresholds",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "event_type": {
                          "computed": true,
                          "description": "The type of event",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "event_value_threshold": {
                          "computed": true,
                          "description": "The threshold of the event type. Only events with a value greater or equal to this threshold will be considered for solution creation.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
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
      "recommender_name": {
        "description": "The name of the recommender",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "recommender_recipe_name": {
        "description": "The name of the recommender recipe.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the recommender",
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "training_metrics": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "metrics": {
              "computed": true,
              "description": "Training metrics by type",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "coverage": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "freshness": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "hit": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "popularity": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "recall": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "similarity": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "time": {
              "computed": true,
              "description": "Timestamp of the training metrics",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Resource Type definition for AWS::CustomerProfiles::Recommender",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCustomerprofilesRecommenderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCustomerprofilesRecommender), &result)
	return &result
}
