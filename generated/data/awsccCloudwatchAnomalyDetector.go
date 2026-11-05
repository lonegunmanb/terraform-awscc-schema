package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudwatchAnomalyDetector = `{
  "block": {
    "attributes": {
      "anomaly_detector_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "excluded_time_ranges": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "end_time": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "start_time": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "metric_time_zone": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "dimensions": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
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
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "metric_characteristics": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "periodic_spikes": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "metric_math_anomaly_detector": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "metric_data_queries": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "account_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "expression": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "label": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "metric_stat": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "metric": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "dimensions": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "name": {
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
                              "metric_name": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "namespace": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "period": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "stat": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "unit": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "period": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "return_data": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "metric_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "namespace": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "single_metric_anomaly_detector": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "account_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "dimensions": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
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
            "metric_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "namespace": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "stat": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "stat": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::CloudWatch::AnomalyDetector",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCloudwatchAnomalyDetectorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudwatchAnomalyDetector), &result)
	return &result
}
