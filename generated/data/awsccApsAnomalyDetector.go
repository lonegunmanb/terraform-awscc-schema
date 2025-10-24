package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApsAnomalyDetector = `{
  "block": {
    "attributes": {
      "alias": {
        "computed": true,
        "description": "The AnomalyDetector alias.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The AnomalyDetector ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "configuration": {
        "computed": true,
        "description": "Determines the anomaly detector's algorithm and its configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "random_cut_forest": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "ignore_near_expected_from_above": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "amount": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "ratio": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "ignore_near_expected_from_below": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "amount": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "ratio": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "query": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "sample_size": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "shingle_size": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "evaluation_interval_in_seconds": {
        "computed": true,
        "description": "The AnomalyDetector period of detection and metric generation.",
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "labels": {
        "computed": true,
        "description": "An array of key-value pairs to provide meta-data.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "missing_data_action": {
        "computed": true,
        "description": "The action to perform when running the expression returns no data.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "mark_as_anomaly": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "skip": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "workspace": {
        "computed": true,
        "description": "Required to identify a specific APS Workspace associated with this Anomaly Detector.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::APS::AnomalyDetector",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccApsAnomalyDetectorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApsAnomalyDetector), &result)
	return &result
}
