package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApplicationsignalsServiceLevelObjective = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN of this SLO.",
        "description_kind": "plain",
        "type": "string"
      },
      "burn_rate_configurations": {
        "computed": true,
        "description": "Each object in this array defines the length of the look-back window used to calculate one burn rate metric for this SLO. The burn rate measures how fast the service is consuming the error budget, relative to the attainment goal of the SLO.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "look_back_window_minutes": {
              "computed": true,
              "description": "The number of minutes to use as the look-back window.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "created_time": {
        "computed": true,
        "description": "Epoch time in seconds of the time that this SLO was created",
        "description_kind": "plain",
        "type": "number"
      },
      "description": {
        "computed": true,
        "description": "An optional description for this SLO. Default is 'No description'",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "evaluation_type": {
        "computed": true,
        "description": "Displays whether this is a period-based SLO or a request-based SLO.",
        "description_kind": "plain",
        "type": "string"
      },
      "exclusion_windows": {
        "computed": true,
        "description": "Each object in this array defines a time exclusion window for this SLO. The time exclusion window is used to exclude breaching data points from affecting attainment rate, error budget, and burn rate metrics.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "reason": {
              "computed": true,
              "description": "An optional reason for scheduling this time exclusion window. Default is 'No reason'.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "recurrence_rule": {
              "computed": true,
              "description": "This object defines how often to repeat a time exclusion window.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "expression": {
                    "computed": true,
                    "description": "A cron or rate expression denoting how often to repeat this exclusion window.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "start_time": {
              "computed": true,
              "description": "The time you want the exclusion window to start at. Note that time exclusion windows can only be scheduled in the future, not the past.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "window": {
              "computed": true,
              "description": "This object defines the length of time an exclusion window should span.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "duration": {
                    "computed": true,
                    "description": "Specifies the duration of each interval. For example, if ` + "`" + `Duration` + "`" + ` is 1 and ` + "`" + `DurationUnit` + "`" + ` is ` + "`" + `MONTH` + "`" + `, each interval is one month, aligned with the calendar.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "duration_unit": {
                    "computed": true,
                    "description": "Specifies the interval unit.",
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
          "nesting_mode": "set"
        },
        "optional": true
      },
      "goal": {
        "computed": true,
        "description": "A structure that contains the attributes that determine the goal of the SLO. This includes the time period for evaluation and the attainment threshold.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "attainment_goal": {
              "computed": true,
              "description": "The threshold that determines if the goal is being met. An attainment goal is the ratio of good periods that meet the threshold requirements to the total periods within the interval. For example, an attainment goal of 99.9% means that within your interval, you are targeting 99.9% of the periods to be in healthy state.\nIf you omit this parameter, 99 is used to represent 99% as the attainment goal.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "interval": {
              "computed": true,
              "description": "The time period used to evaluate the SLO. It can be either a calendar interval or rolling interval.\nIf you omit this parameter, a rolling interval of 7 days is used.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "calendar_interval": {
                    "computed": true,
                    "description": "If the interval for this service level objective is a calendar interval, this structure contains the interval specifications.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "duration": {
                          "computed": true,
                          "description": "Specifies the duration of each interval. For example, if ` + "`" + `Duration` + "`" + ` is 1 and ` + "`" + `DurationUnit` + "`" + ` is ` + "`" + `MONTH` + "`" + `, each interval is one month, aligned with the calendar.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "duration_unit": {
                          "computed": true,
                          "description": "Specifies the interval unit.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "start_time": {
                          "computed": true,
                          "description": "Epoch time in seconds you want the first interval to start. Be sure to choose a time that configures the intervals the way that you want. For example, if you want weekly intervals starting on Mondays at 6 a.m., be sure to specify a start time that is a Monday at 6 a.m.\nAs soon as one calendar interval ends, another automatically begins.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "rolling_interval": {
                    "computed": true,
                    "description": "If the interval is a calendar interval, this structure contains the interval specifications.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "duration": {
                          "computed": true,
                          "description": "Specifies the duration of each interval. For example, if ` + "`" + `Duration` + "`" + ` is 1 and ` + "`" + `DurationUnit` + "`" + ` is ` + "`" + `MONTH` + "`" + `, each interval is one month, aligned with the calendar.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "duration_unit": {
                          "computed": true,
                          "description": "Specifies the interval unit.",
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
            "warning_threshold": {
              "computed": true,
              "description": "The percentage of remaining budget over total budget that you want to get warnings for. If you omit this parameter, the default of 50.0 is used.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_updated_time": {
        "computed": true,
        "description": "Epoch time in seconds of the time that this SLO was most recently updated",
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "description": "The name of this SLO.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "request_based_sli": {
        "computed": true,
        "description": "This structure contains information about the performance metric that a request-based SLO monitors.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "comparison_operator": {
              "computed": true,
              "description": "The arithmetic operation used when comparing the specified metric to the threshold.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "metric_threshold": {
              "computed": true,
              "description": "The value that the SLI metric is compared to.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "request_based_sli_metric": {
              "computed": true,
              "description": "This structure contains the information about the metric that is used for a request-based SLO.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "dependency_config": {
                    "computed": true,
                    "description": "Configuration for identifying a dependency and its operation",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "dependency_key_attributes": {
                          "computed": true,
                          "description": "If this SLO is related to a metric collected by Application Signals, you must use this field to specify which dependency the SLO metric is related to.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "dependency_operation_name": {
                          "computed": true,
                          "description": "When the SLO monitors a specific operation of the dependency, this field specifies the name of that operation in the dependency.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "key_attributes": {
                    "computed": true,
                    "description": "This is a string-to-string map that contains information about the type of object that this SLO is related to.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "metric_type": {
                    "computed": true,
                    "description": "If the SLO monitors either the LATENCY or AVAILABILITY metric that Application Signals collects, this field displays which of those metrics is used.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "monitored_request_count_metric": {
                    "computed": true,
                    "description": "This structure defines the metric that is used as the \"good request\" or \"bad request\" value for a request-based SLO. This value observed for the metric defined in ` + "`" + `TotalRequestCountMetric` + "`" + ` is divided by the number found for ` + "`" + `MonitoredRequestCountMetric` + "`" + ` to determine the percentage of successful requests that this SLO tracks.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bad_count_metric": {
                          "computed": true,
                          "description": "If you want to count \"bad requests\" to determine the percentage of successful requests for this request-based SLO, specify the metric to use as \"bad requests\" in this structure.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "account_id": {
                                "computed": true,
                                "description": "The ID of the account where the metrics are located, if this is a cross-account alarm.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "expression": {
                                "computed": true,
                                "description": "The math expression to be performed on the returned data.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "id": {
                                "computed": true,
                                "description": "A short name used to tie this object to the results in the response.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "metric_stat": {
                                "computed": true,
                                "description": "A metric to be used directly for the SLO, or to be used in the math expression that will be used for the SLO. Within one MetricDataQuery, you must specify either Expression or MetricStat but not both.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "metric": {
                                      "computed": true,
                                      "description": "This structure defines the metric used for a service level indicator, including the metric name, namespace, and dimensions.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "dimensions": {
                                            "computed": true,
                                            "description": "An array of one or more dimensions to use to define the metric that you want to use.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "name": {
                                                  "computed": true,
                                                  "description": "The name of the dimension. Dimension names must contain only ASCII characters, must include at least one non-whitespace character, and cannot start with a colon (:). ASCII control characters are not supported as part of dimension names.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "value": {
                                                  "computed": true,
                                                  "description": "The value of the dimension. Dimension values must contain only ASCII characters and must include at least one non-whitespace character. ASCII control characters are not supported as part of dimension values",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "list"
                                            },
                                            "optional": true
                                          },
                                          "metric_name": {
                                            "computed": true,
                                            "description": "The name of the metric to use.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "namespace": {
                                            "computed": true,
                                            "description": "The namespace of the metric.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "period": {
                                      "computed": true,
                                      "description": "The granularity, in seconds, to be used for the metric.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "stat": {
                                      "computed": true,
                                      "description": "The statistic to use for comparison to the threshold. It can be any CloudWatch statistic or extended statistic.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "unit": {
                                      "computed": true,
                                      "description": "If you omit Unit then all data that was collected with any unit is returned, along with the corresponding units that were specified when the data was reported to CloudWatch. If you specify a unit, the operation returns only data that was collected with that unit specified. If you specify a unit that does not match the data collected, the results of the operation are null. CloudWatch does not perform unit conversions.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              },
                              "return_data": {
                                "computed": true,
                                "description": "This option indicates whether to return the timestamps and raw data values of this metric.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "optional": true
                        },
                        "good_count_metric": {
                          "computed": true,
                          "description": "If you want to count \"good requests\" to determine the percentage of successful requests for this request-based SLO, specify the metric to use as \"good requests\" in this structure.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "account_id": {
                                "computed": true,
                                "description": "The ID of the account where the metrics are located, if this is a cross-account alarm.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "expression": {
                                "computed": true,
                                "description": "The math expression to be performed on the returned data.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "id": {
                                "computed": true,
                                "description": "A short name used to tie this object to the results in the response.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "metric_stat": {
                                "computed": true,
                                "description": "A metric to be used directly for the SLO, or to be used in the math expression that will be used for the SLO. Within one MetricDataQuery, you must specify either Expression or MetricStat but not both.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "metric": {
                                      "computed": true,
                                      "description": "This structure defines the metric used for a service level indicator, including the metric name, namespace, and dimensions.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "dimensions": {
                                            "computed": true,
                                            "description": "An array of one or more dimensions to use to define the metric that you want to use.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "name": {
                                                  "computed": true,
                                                  "description": "The name of the dimension. Dimension names must contain only ASCII characters, must include at least one non-whitespace character, and cannot start with a colon (:). ASCII control characters are not supported as part of dimension names.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "value": {
                                                  "computed": true,
                                                  "description": "The value of the dimension. Dimension values must contain only ASCII characters and must include at least one non-whitespace character. ASCII control characters are not supported as part of dimension values",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "list"
                                            },
                                            "optional": true
                                          },
                                          "metric_name": {
                                            "computed": true,
                                            "description": "The name of the metric to use.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "namespace": {
                                            "computed": true,
                                            "description": "The namespace of the metric.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "period": {
                                      "computed": true,
                                      "description": "The granularity, in seconds, to be used for the metric.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "stat": {
                                      "computed": true,
                                      "description": "The statistic to use for comparison to the threshold. It can be any CloudWatch statistic or extended statistic.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "unit": {
                                      "computed": true,
                                      "description": "If you omit Unit then all data that was collected with any unit is returned, along with the corresponding units that were specified when the data was reported to CloudWatch. If you specify a unit, the operation returns only data that was collected with that unit specified. If you specify a unit that does not match the data collected, the results of the operation are null. CloudWatch does not perform unit conversions.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              },
                              "return_data": {
                                "computed": true,
                                "description": "This option indicates whether to return the timestamps and raw data values of this metric.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
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
                  },
                  "operation_name": {
                    "computed": true,
                    "description": "If the SLO monitors a specific operation of the service, this field displays that operation name.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "total_request_count_metric": {
                    "computed": true,
                    "description": "This structure defines the metric that is used as the \"total requests\" number for a request-based SLO. The number observed for this metric is divided by the number of \"good requests\" or \"bad requests\" that is observed for the metric defined in ` + "`" + `MonitoredRequestCountMetric` + "`" + `.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "account_id": {
                          "computed": true,
                          "description": "The ID of the account where the metrics are located, if this is a cross-account alarm.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "expression": {
                          "computed": true,
                          "description": "The math expression to be performed on the returned data.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "id": {
                          "computed": true,
                          "description": "A short name used to tie this object to the results in the response.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "metric_stat": {
                          "computed": true,
                          "description": "A metric to be used directly for the SLO, or to be used in the math expression that will be used for the SLO. Within one MetricDataQuery, you must specify either Expression or MetricStat but not both.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "metric": {
                                "computed": true,
                                "description": "This structure defines the metric used for a service level indicator, including the metric name, namespace, and dimensions.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "dimensions": {
                                      "computed": true,
                                      "description": "An array of one or more dimensions to use to define the metric that you want to use.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "name": {
                                            "computed": true,
                                            "description": "The name of the dimension. Dimension names must contain only ASCII characters, must include at least one non-whitespace character, and cannot start with a colon (:). ASCII control characters are not supported as part of dimension names.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "value": {
                                            "computed": true,
                                            "description": "The value of the dimension. Dimension values must contain only ASCII characters and must include at least one non-whitespace character. ASCII control characters are not supported as part of dimension values",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "list"
                                      },
                                      "optional": true
                                    },
                                    "metric_name": {
                                      "computed": true,
                                      "description": "The name of the metric to use.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "namespace": {
                                      "computed": true,
                                      "description": "The namespace of the metric.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              },
                              "period": {
                                "computed": true,
                                "description": "The granularity, in seconds, to be used for the metric.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "stat": {
                                "computed": true,
                                "description": "The statistic to use for comparison to the threshold. It can be any CloudWatch statistic or extended statistic.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "unit": {
                                "computed": true,
                                "description": "If you omit Unit then all data that was collected with any unit is returned, along with the corresponding units that were specified when the data was reported to CloudWatch. If you specify a unit, the operation returns only data that was collected with that unit specified. If you specify a unit that does not match the data collected, the results of the operation are null. CloudWatch does not perform unit conversions.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "return_data": {
                          "computed": true,
                          "description": "This option indicates whether to return the timestamps and raw data values of this metric.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
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
      "sli": {
        "computed": true,
        "description": "This structure contains information about the performance metric that an SLO monitors.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "comparison_operator": {
              "computed": true,
              "description": "The arithmetic operation used when comparing the specified metric to the threshold.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "metric_threshold": {
              "computed": true,
              "description": "The value that the SLI metric is compared to.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "sli_metric": {
              "computed": true,
              "description": "A structure that contains information about the metric that the SLO monitors.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "dependency_config": {
                    "computed": true,
                    "description": "Configuration for identifying a dependency and its operation",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "dependency_key_attributes": {
                          "computed": true,
                          "description": "If this SLO is related to a metric collected by Application Signals, you must use this field to specify which dependency the SLO metric is related to.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "dependency_operation_name": {
                          "computed": true,
                          "description": "When the SLO monitors a specific operation of the dependency, this field specifies the name of that operation in the dependency.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "key_attributes": {
                    "computed": true,
                    "description": "This is a string-to-string map that contains information about the type of object that this SLO is related to.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "metric_data_queries": {
                    "computed": true,
                    "description": "If this SLO monitors a CloudWatch metric or the result of a CloudWatch metric math expression, this structure includes the information about that metric or expression.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "account_id": {
                          "computed": true,
                          "description": "The ID of the account where the metrics are located, if this is a cross-account alarm.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "expression": {
                          "computed": true,
                          "description": "The math expression to be performed on the returned data.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "id": {
                          "computed": true,
                          "description": "A short name used to tie this object to the results in the response.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "metric_stat": {
                          "computed": true,
                          "description": "A metric to be used directly for the SLO, or to be used in the math expression that will be used for the SLO. Within one MetricDataQuery, you must specify either Expression or MetricStat but not both.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "metric": {
                                "computed": true,
                                "description": "This structure defines the metric used for a service level indicator, including the metric name, namespace, and dimensions.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "dimensions": {
                                      "computed": true,
                                      "description": "An array of one or more dimensions to use to define the metric that you want to use.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "name": {
                                            "computed": true,
                                            "description": "The name of the dimension. Dimension names must contain only ASCII characters, must include at least one non-whitespace character, and cannot start with a colon (:). ASCII control characters are not supported as part of dimension names.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "value": {
                                            "computed": true,
                                            "description": "The value of the dimension. Dimension values must contain only ASCII characters and must include at least one non-whitespace character. ASCII control characters are not supported as part of dimension values",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "list"
                                      },
                                      "optional": true
                                    },
                                    "metric_name": {
                                      "computed": true,
                                      "description": "The name of the metric to use.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "namespace": {
                                      "computed": true,
                                      "description": "The namespace of the metric.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              },
                              "period": {
                                "computed": true,
                                "description": "The granularity, in seconds, to be used for the metric.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "stat": {
                                "computed": true,
                                "description": "The statistic to use for comparison to the threshold. It can be any CloudWatch statistic or extended statistic.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "unit": {
                                "computed": true,
                                "description": "If you omit Unit then all data that was collected with any unit is returned, along with the corresponding units that were specified when the data was reported to CloudWatch. If you specify a unit, the operation returns only data that was collected with that unit specified. If you specify a unit that does not match the data collected, the results of the operation are null. CloudWatch does not perform unit conversions.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "return_data": {
                          "computed": true,
                          "description": "This option indicates whether to return the timestamps and raw data values of this metric.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "metric_type": {
                    "computed": true,
                    "description": "If the SLO monitors either the LATENCY or AVAILABILITY metric that Application Signals collects, this field displays which of those metrics is used.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "operation_name": {
                    "computed": true,
                    "description": "If the SLO monitors a specific operation of the service, this field displays that operation name.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "period_seconds": {
                    "computed": true,
                    "description": "The number of seconds to use as the period for SLO evaluation. Your application's performance is compared to the SLI during each period. For each period, the application is determined to have either achieved or not achieved the necessary performance.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "statistic": {
                    "computed": true,
                    "description": "The statistic to use for comparison to the threshold. It can be any CloudWatch statistic or extended statistic",
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
      "tags": {
        "computed": true,
        "description": "The list of tag keys and values associated with the resource you specified",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "A string that you can use to assign a value. The combination of tag keys and values can help you organize and categorize your resources.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the specified tag key.",
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
    "description": "Resource Type definition for AWS::ApplicationSignals::ServiceLevelObjective",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApplicationsignalsServiceLevelObjectiveSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApplicationsignalsServiceLevelObjective), &result)
	return &result
}
