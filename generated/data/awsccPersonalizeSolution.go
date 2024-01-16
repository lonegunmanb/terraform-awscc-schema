package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccPersonalizeSolution = `{
  "block": {
    "attributes": {
      "dataset_group_arn": {
        "computed": true,
        "description": "The ARN of the dataset group that provides the training data.",
        "description_kind": "plain",
        "type": "string"
      },
      "event_type": {
        "computed": true,
        "description": "When your have multiple event types (using an EVENT_TYPE schema field), this parameter specifies which event type (for example, 'click' or 'like') is used for training the model. If you do not provide an eventType, Amazon Personalize will use all interactions for training with equal weight regardless of type.",
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
        "description": "The name for the solution",
        "description_kind": "plain",
        "type": "string"
      },
      "perform_auto_ml": {
        "computed": true,
        "description": "Whether to perform automated machine learning (AutoML). The default is false. For this case, you must specify recipeArn.",
        "description_kind": "plain",
        "type": "bool"
      },
      "perform_hpo": {
        "computed": true,
        "description": "Whether to perform hyperparameter optimization (HPO) on the specified or selected recipe. The default is false. When performing AutoML, this parameter is always true and you should not set it to false.",
        "description_kind": "plain",
        "type": "bool"
      },
      "recipe_arn": {
        "computed": true,
        "description": "The ARN of the recipe to use for model training. Only specified when performAutoML is false.",
        "description_kind": "plain",
        "type": "string"
      },
      "solution_arn": {
        "computed": true,
        "description": "The ARN of the solution",
        "description_kind": "plain",
        "type": "string"
      },
      "solution_config": {
        "computed": true,
        "description": "The configuration to use with the solution. When performAutoML is set to true, Amazon Personalize only evaluates the autoMLConfig section of the solution configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "algorithm_hyper_parameters": {
              "computed": true,
              "description": "Lists the hyperparameter names and ranges.",
              "description_kind": "plain",
              "type": [
                "map",
                "string"
              ]
            },
            "auto_ml_config": {
              "computed": true,
              "description": "The AutoMLConfig object containing a list of recipes to search when AutoML is performed.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "metric_name": {
                    "computed": true,
                    "description": "The metric to optimize.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "recipe_list": {
                    "computed": true,
                    "description": "The list of candidate recipes.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              }
            },
            "event_value_threshold": {
              "computed": true,
              "description": "Only events with a value greater than or equal to this threshold are used for training a model.",
              "description_kind": "plain",
              "type": "string"
            },
            "feature_transformation_parameters": {
              "computed": true,
              "description": "Lists the feature transformation parameters.",
              "description_kind": "plain",
              "type": [
                "map",
                "string"
              ]
            },
            "hpo_config": {
              "computed": true,
              "description": "Describes the properties for hyperparameter optimization (HPO)",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "algorithm_hyper_parameter_ranges": {
                    "computed": true,
                    "description": "The hyperparameters and their allowable ranges",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "categorical_hyper_parameter_ranges": {
                          "computed": true,
                          "description": "The categorical hyperparameters and their ranges.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "name": {
                                "computed": true,
                                "description": "The name of the hyperparameter.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "values": {
                                "computed": true,
                                "description": "A list of the categories for the hyperparameter.",
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
                        "continuous_hyper_parameter_ranges": {
                          "computed": true,
                          "description": "The continuous hyperparameters and their ranges.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "max_value": {
                                "computed": true,
                                "description": "The maximum allowable value for the hyperparameter.",
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "min_value": {
                                "computed": true,
                                "description": "The minimum allowable value for the hyperparameter.",
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "name": {
                                "computed": true,
                                "description": "The name of the hyperparameter.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        },
                        "integer_hyper_parameter_ranges": {
                          "computed": true,
                          "description": "The integer hyperparameters and their ranges.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "max_value": {
                                "computed": true,
                                "description": "The maximum allowable value for the hyperparameter.",
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "min_value": {
                                "computed": true,
                                "description": "The minimum allowable value for the hyperparameter.",
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "name": {
                                "computed": true,
                                "description": "The name of the hyperparameter.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "hpo_objective": {
                    "computed": true,
                    "description": "The metric to optimize during HPO.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "metric_name": {
                          "computed": true,
                          "description": "The name of the metric",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "metric_regex": {
                          "computed": true,
                          "description": "A regular expression for finding the metric in the training job logs.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "type": {
                          "computed": true,
                          "description": "The type of the metric. Valid values are Maximize and Minimize.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "hpo_resource_config": {
                    "computed": true,
                    "description": "Describes the resource configuration for hyperparameter optimization (HPO).",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "max_number_of_training_jobs": {
                          "computed": true,
                          "description": "The maximum number of training jobs when you create a solution version. The maximum value for maxNumberOfTrainingJobs is 40.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "max_parallel_training_jobs": {
                          "computed": true,
                          "description": "The maximum number of parallel training jobs when you create a solution version. The maximum value for maxParallelTrainingJobs is 10.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::Personalize::Solution",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccPersonalizeSolutionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPersonalizeSolution), &result)
	return &result
}
