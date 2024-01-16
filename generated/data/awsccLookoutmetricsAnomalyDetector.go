package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLookoutmetricsAnomalyDetector = `{
  "block": {
    "attributes": {
      "anomaly_detector_config": {
        "computed": true,
        "description": "Configuration options for the AnomalyDetector",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "anomaly_detector_frequency": {
              "computed": true,
              "description": "Frequency of anomaly detection",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "anomaly_detector_description": {
        "computed": true,
        "description": "A description for the AnomalyDetector.",
        "description_kind": "plain",
        "type": "string"
      },
      "anomaly_detector_name": {
        "computed": true,
        "description": "Name for the Amazon Lookout for Metrics Anomaly Detector",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kms_key_arn": {
        "computed": true,
        "description": "KMS key used to encrypt the AnomalyDetector data",
        "description_kind": "plain",
        "type": "string"
      },
      "metric_set_list": {
        "computed": true,
        "description": "List of metric sets for anomaly detection",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "dimension_list": {
              "computed": true,
              "description": "Dimensions for this MetricSet.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "metric_list": {
              "computed": true,
              "description": "Metrics captured by this MetricSet.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "aggregation_function": {
                    "computed": true,
                    "description": "Operator used to aggregate metric values",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "metric_name": {
                    "computed": true,
                    "description": "Name of a column in the data.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "namespace": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "metric_set_description": {
              "computed": true,
              "description": "A description for the MetricSet.",
              "description_kind": "plain",
              "type": "string"
            },
            "metric_set_frequency": {
              "computed": true,
              "description": "A frequency period to aggregate the data",
              "description_kind": "plain",
              "type": "string"
            },
            "metric_set_name": {
              "computed": true,
              "description": "The name of the MetricSet.",
              "description_kind": "plain",
              "type": "string"
            },
            "metric_source": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "app_flow_config": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "flow_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "cloudwatch_config": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "rds_source_config": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "database_host": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "database_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "database_port": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "db_instance_identifier": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "secret_manager_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "table_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "vpc_configuration": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "security_group_id_list": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "subnet_id_list": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
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
                  "redshift_source_config": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "cluster_identifier": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "database_host": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "database_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "database_port": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "secret_manager_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "table_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "vpc_configuration": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "security_group_id_list": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "subnet_id_list": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
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
                  "s3_source_config": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "file_format_descriptor": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "csv_format_descriptor": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "charset": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "contains_header": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "bool"
                                    },
                                    "delimiter": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "file_compression": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "header_list": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    },
                                    "quote_symbol": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "json_format_descriptor": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "charset": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": "string"
                                    },
                                    "file_compression": {
                                      "computed": true,
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
                        },
                        "historical_data_path_list": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "templated_path_list": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
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
            "offset": {
              "computed": true,
              "description": "Offset, in seconds, between the frequency interval and the time at which the metrics are available.",
              "description_kind": "plain",
              "type": "number"
            },
            "timestamp_column": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "column_format": {
                    "computed": true,
                    "description": "A timestamp format for the timestamps in the dataset",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "column_name": {
                    "computed": true,
                    "description": "Name of a column in the data.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "timezone": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::LookoutMetrics::AnomalyDetector",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLookoutmetricsAnomalyDetectorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLookoutmetricsAnomalyDetector), &result)
	return &result
}
