package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLookoutmetricsAnomalyDetector = `{
  "block": {
    "attributes": {
      "anomaly_detector_config": {
        "description": "Configuration options for the AnomalyDetector",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "anomaly_detector_frequency": {
              "description": "Frequency of anomaly detection",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "anomaly_detector_description": {
        "computed": true,
        "description": "A description for the AnomalyDetector.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "anomaly_detector_name": {
        "computed": true,
        "description": "Name for the Amazon Lookout for Metrics Anomaly Detector",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_arn": {
        "computed": true,
        "description": "KMS key used to encrypt the AnomalyDetector data",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "metric_set_list": {
        "description": "List of metric sets for anomaly detection",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "dimension_list": {
              "computed": true,
              "description": "Dimensions for this MetricSet.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "metric_list": {
              "description": "Metrics captured by this MetricSet.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "aggregation_function": {
                    "description": "Operator used to aggregate metric values",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "metric_name": {
                    "description": "Name of a column in the data.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "namespace": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "required": true
            },
            "metric_set_description": {
              "computed": true,
              "description": "A description for the MetricSet.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "metric_set_frequency": {
              "computed": true,
              "description": "A frequency period to aggregate the data",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "metric_set_name": {
              "description": "The name of the MetricSet.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "metric_source": {
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
                          "optional": true,
                          "type": "string"
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "cloudwatch_config": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "rds_source_config": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "database_host": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "database_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "database_port": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "db_instance_identifier": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "secret_manager_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "table_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
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
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "subnet_id_list": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
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
                  "redshift_source_config": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "cluster_identifier": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "database_host": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "database_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "database_port": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "secret_manager_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "table_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
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
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "subnet_id_list": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
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
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "contains_header": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "delimiter": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "file_compression": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "header_list": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    },
                                    "quote_symbol": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              },
                              "json_format_descriptor": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "charset": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "file_compression": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
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
                        "historical_data_path_list": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "role_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "templated_path_list": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
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
              "required": true
            },
            "offset": {
              "computed": true,
              "description": "Offset, in seconds, between the frequency interval and the time at which the metrics are available.",
              "description_kind": "plain",
              "optional": true,
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
                    "optional": true,
                    "type": "string"
                  },
                  "column_name": {
                    "computed": true,
                    "description": "Name of a column in the data.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "timezone": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      }
    },
    "description": "An Amazon Lookout for Metrics Detector",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLookoutmetricsAnomalyDetectorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLookoutmetricsAnomalyDetector), &result)
	return &result
}
