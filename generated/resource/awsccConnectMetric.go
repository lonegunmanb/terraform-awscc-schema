package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConnectMetric = `{
  "block": {
    "attributes": {
      "category": {
        "computed": true,
        "description": "The category of the custom metric",
        "description_kind": "plain",
        "type": "string"
      },
      "created_time": {
        "computed": true,
        "description": "The timestamp where the metric was created",
        "description_kind": "plain",
        "type": "number"
      },
      "created_user": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "aws_identity_arn": {
              "computed": true,
              "description": "STS or IAM ARN representing the identity of API Caller. SDK users cannot populate this and this value is calculated automatically if ConnectUserArn is not provided.",
              "description_kind": "plain",
              "type": "string"
            },
            "connect_user_arn": {
              "computed": true,
              "description": "An agent ARN representing a connect user.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "creation_method": {
        "computed": true,
        "description": "Whether the metric was built with the guided Service Level (SL) experience, or with the free-form metric builder",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the custom metric",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "effective_time": {
        "computed": true,
        "description": "Earliest time that can be queried for this metric",
        "description_kind": "plain",
        "type": "number"
      },
      "filters": {
        "computed": true,
        "description": "List of filter types that may be used with this metric",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "groupings": {
        "computed": true,
        "description": "List of groupings that may be used with this metric",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_arn": {
        "description": "The identifier of the Amazon Connect instance.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_modified_region": {
        "computed": true,
        "description": "The AWS region where the metric was last modified",
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description": "The timestamp where the metric was last modified",
        "description_kind": "plain",
        "type": "number"
      },
      "last_modified_user": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "aws_identity_arn": {
              "computed": true,
              "description": "STS or IAM ARN representing the identity of API Caller. SDK users cannot populate this and this value is calculated automatically if ConnectUserArn is not provided.",
              "description_kind": "plain",
              "type": "string"
            },
            "connect_user_arn": {
              "computed": true,
              "description": "An agent ARN representing a connect user.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "metric_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the custom metric.",
        "description_kind": "plain",
        "type": "string"
      },
      "metric_calculation": {
        "computed": true,
        "description": "The calculation configuration for the metric",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "calculation": {
              "computed": true,
              "description": "The calculation formula",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "calculation_components": {
              "computed": true,
              "description": "The calculation components for the metric",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "alias": {
                    "computed": true,
                    "description": "Metric calculation component alias for use within a calculation",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "metric_filters": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "boolean_condition": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "comparison": {
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
                        "metric_filter_key": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "negate": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "number_condition": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "comparison": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "values": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "number"
                                ]
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "string_condition": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "comparison": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "values": {
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
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "metric_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "metric_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
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
      "name": {
        "computed": true,
        "description": "The name of the custom metric",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "positive_trend_indicator": {
        "computed": true,
        "description": "Indicates how to classify a positive trend in metric data on the UI",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "primary_event_source": {
        "computed": true,
        "description": "Main provider of the document/row-level data for the metric; should match Data Lake table names",
        "description_kind": "plain",
        "type": "string"
      },
      "primary_event_source_effective_timestamp_type": {
        "computed": true,
        "description": "Identifies the timestamp used to place the metrics on a time-series; should match public attribute name",
        "description_kind": "plain",
        "type": "string"
      },
      "refresh_rate": {
        "computed": true,
        "description": "Recommended time to wait between each refresh of data for the metric",
        "description_kind": "plain",
        "type": "number"
      },
      "status": {
        "computed": true,
        "description": "The status of the custom metric",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "supported_stats": {
        "computed": true,
        "description": "List of stat aggregations available for the metric",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "supports_custom_calculation": {
        "computed": true,
        "description": "The metric may be used to compose other (custom) metrics",
        "description_kind": "plain",
        "type": "bool"
      },
      "supports_preaggregate_calculation": {
        "computed": true,
        "description": "The metric may be used to compose other (custom) metrics, meaning it can be used inside of aggregate stat functions",
        "description_kind": "plain",
        "type": "bool"
      },
      "tags": {
        "computed": true,
        "description": "One or more tags.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is maximum of 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "type": {
        "computed": true,
        "description": "Whether the metric is provided out-of-the-box or created by each customer",
        "description_kind": "plain",
        "type": "string"
      },
      "unit": {
        "computed": true,
        "description": "Display unit for the metric data",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Connect::Metric, a custom metric configured for an Amazon Connect instance",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccConnectMetricSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectMetric), &result)
	return &result
}
