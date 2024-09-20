package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotSecurityProfile = `{
  "block": {
    "attributes": {
      "additional_metrics_to_retain_v2": {
        "computed": true,
        "description": "A list of metrics whose data is retained (stored). By default, data is retained for any metric used in the profile's behaviors, but it is also retained for any metric specified here.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "export_metric": {
              "computed": true,
              "description": "Flag to enable/disable metrics export for metric to be retained.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "metric": {
              "computed": true,
              "description": "What is measured by the behavior.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "metric_dimension": {
              "computed": true,
              "description": "The dimension of a metric.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "dimension_name": {
                    "computed": true,
                    "description": "A unique identifier for the dimension.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "operator": {
                    "computed": true,
                    "description": "Defines how the dimensionValues of a dimension are interpreted.",
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
      "alert_targets": {
        "computed": true,
        "description": "Specifies the destinations to which alerts are sent.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "alert_target_arn": {
              "computed": true,
              "description": "The ARN of the notification target to which alerts are sent.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "role_arn": {
              "computed": true,
              "description": "The ARN of the role that grants permission to send alerts to the notification target.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "map"
        },
        "optional": true
      },
      "behaviors": {
        "computed": true,
        "description": "Specifies the behaviors that, when violated by a device (thing), cause an alert.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "criteria": {
              "computed": true,
              "description": "The criteria by which the behavior is determined to be normal.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "comparison_operator": {
                    "computed": true,
                    "description": "The operator that relates the thing measured (metric) to the criteria (containing a value or statisticalThreshold).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "consecutive_datapoints_to_alarm": {
                    "computed": true,
                    "description": "If a device is in violation of the behavior for the specified number of consecutive datapoints, an alarm occurs. If not specified, the default is 1.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "consecutive_datapoints_to_clear": {
                    "computed": true,
                    "description": "If an alarm has occurred and the offending device is no longer in violation of the behavior for the specified number of consecutive datapoints, the alarm is cleared. If not specified, the default is 1.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "duration_seconds": {
                    "computed": true,
                    "description": "Use this to specify the time duration over which the behavior is evaluated.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "ml_detection_config": {
                    "computed": true,
                    "description": "The configuration of an ML Detect Security Profile.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "confidence_level": {
                          "computed": true,
                          "description": "The sensitivity of anomalous behavior evaluation. Can be Low, Medium, or High.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "statistical_threshold": {
                    "computed": true,
                    "description": "A statistical ranking (percentile) which indicates a threshold value by which a behavior is determined to be in compliance or in violation of the behavior.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "statistic": {
                          "computed": true,
                          "description": "The percentile which resolves to a threshold value by which compliance with a behavior is determined",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "value": {
                    "computed": true,
                    "description": "The value to be compared with the metric.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "cidrs": {
                          "computed": true,
                          "description": "If the ComparisonOperator calls for a set of CIDRs, use this to specify that set to be compared with the metric.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "count": {
                          "computed": true,
                          "description": "If the ComparisonOperator calls for a numeric value, use this to specify that (integer) numeric value to be compared with the metric.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "number": {
                          "computed": true,
                          "description": "The numeral value of a metric.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "numbers": {
                          "computed": true,
                          "description": "The numeral values of a metric.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "number"
                          ]
                        },
                        "ports": {
                          "computed": true,
                          "description": "If the ComparisonOperator calls for a set of ports, use this to specify that set to be compared with the metric.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "number"
                          ]
                        },
                        "strings": {
                          "computed": true,
                          "description": "The string values of a metric.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
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
            "export_metric": {
              "computed": true,
              "description": "Flag to enable/disable metrics export for metric to be retained.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "metric": {
              "computed": true,
              "description": "What is measured by the behavior.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "metric_dimension": {
              "computed": true,
              "description": "The dimension of a metric.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "dimension_name": {
                    "computed": true,
                    "description": "A unique identifier for the dimension.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "operator": {
                    "computed": true,
                    "description": "Defines how the dimensionValues of a dimension are interpreted.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "name": {
              "computed": true,
              "description": "The name for the behavior.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "suppress_alerts": {
              "computed": true,
              "description": "Manage Detect alarm SNS notifications by setting behavior notification to on or suppressed. Detect will continue to performing device behavior evaluations. However, suppressed alarms wouldn't be forwarded for SNS notification.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "metrics_export_config": {
        "computed": true,
        "description": "A structure containing the mqtt topic for metrics export.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "mqtt_topic": {
              "computed": true,
              "description": "The topic for metrics export.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "role_arn": {
              "computed": true,
              "description": "The ARN of the role that grants permission to publish to mqtt topic.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "security_profile_arn": {
        "computed": true,
        "description": "The ARN (Amazon resource name) of the created security profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "security_profile_description": {
        "computed": true,
        "description": "A description of the security profile.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_profile_name": {
        "computed": true,
        "description": "A unique identifier for the security profile.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Metadata that can be used to manage the security profile.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag's key.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag's value.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "target_arns": {
        "computed": true,
        "description": "A set of target ARNs that the security profile is attached to.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      }
    },
    "description": "A security profile defines a set of expected behaviors for devices in your account.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIotSecurityProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotSecurityProfile), &result)
	return &result
}
