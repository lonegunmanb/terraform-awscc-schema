package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccObservabilityadminTelemetryRule = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "rule": {
        "description": "The telemetry rule",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "destination_configuration": {
              "computed": true,
              "description": "The destination configuration for telemetry data",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cloudtrail_parameters": {
                    "computed": true,
                    "description": "Telemetry parameters for Cloudtrail",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "advanced_event_selectors": {
                          "computed": true,
                          "description": "Create fine-grained selectors for AWS CloudTrail management and data.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "field_selectors": {
                                "computed": true,
                                "description": "Contains all selector statements in an advanced event selector.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "ends_with": {
                                      "computed": true,
                                      "description": "An operator that includes events that match the last few characters of the event record field specified as the value of Field.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "set",
                                        "string"
                                      ]
                                    },
                                    "equals": {
                                      "computed": true,
                                      "description": "An operator that includes events that match the exact value of the event record field specified as the value of Field.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "set",
                                        "string"
                                      ]
                                    },
                                    "field": {
                                      "computed": true,
                                      "description": "A field in a CloudTrail event record on which to filter events to be logged",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "not_ends_with": {
                                      "computed": true,
                                      "description": "An operator that excludes events that match the last few characters of the event record field specified as the value of Field.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "set",
                                        "string"
                                      ]
                                    },
                                    "not_equals": {
                                      "computed": true,
                                      "description": "An operator that excludes events that match the exact value of the event record field specified as the value of Field.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "set",
                                        "string"
                                      ]
                                    },
                                    "not_starts_with": {
                                      "computed": true,
                                      "description": "An operator that excludes events that match the first few characters of the event record field specified as the value of Field.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "set",
                                        "string"
                                      ]
                                    },
                                    "starts_with": {
                                      "computed": true,
                                      "description": "An operator that includes events that match the first few characters of the event record field specified as the value of Field.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": [
                                        "set",
                                        "string"
                                      ]
                                    }
                                  },
                                  "nesting_mode": "set"
                                },
                                "optional": true
                              },
                              "name": {
                                "computed": true,
                                "description": "An optional descriptive name for the advanced event selector",
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
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "destination_pattern": {
                    "computed": true,
                    "description": "Pattern for telemetry data destination",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "destination_type": {
                    "computed": true,
                    "description": "Type of telemetry destination",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "elb_load_balancer_logging_parameters": {
                    "computed": true,
                    "description": "Telemetry parameters for ELB/NLB Load Balancer Logs",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "field_delimiter": {
                          "computed": true,
                          "description": "A delimiter to delineate log fields",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "output_format": {
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
                  "log_delivery_parameters": {
                    "computed": true,
                    "description": "Parameters for BedrockAgentCore log delivery",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "log_types": {
                          "computed": true,
                          "description": "Types of logs to deliver for BedrockAgentCore resources",
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
                  },
                  "retention_in_days": {
                    "computed": true,
                    "description": "Number of days to retain the telemetry data in the specified destination",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "vpc_flow_log_parameters": {
                    "computed": true,
                    "description": "Telemetry parameters for VPC Flow logs",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "log_format": {
                          "computed": true,
                          "description": "The fields to include in the flow log record. If you omit this parameter, the flow log is created using the default format.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "max_aggregation_interval": {
                          "computed": true,
                          "description": "The maximum interval of time, in seconds, during which a flow of packets is captured and aggregated into a flow log record. Default is 600s.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "traffic_type": {
                          "computed": true,
                          "description": "The type of traffic captured for the flow log. Default is ALL",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "waf_logging_parameters": {
                    "computed": true,
                    "description": "Telemetry parameters for WAF v2 Web ACL",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "log_type": {
                          "computed": true,
                          "description": "The type of logs to generate for WAF.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "logging_filter": {
                          "computed": true,
                          "description": "Default handling for logs that don't match any of the specified filtering conditions.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "default_behavior": {
                                "computed": true,
                                "description": "The behavior required of the filter.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "filters": {
                                "computed": true,
                                "description": "A list of filters to be applied.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "behavior": {
                                      "computed": true,
                                      "description": "The behavior required of the filter.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "conditions": {
                                      "computed": true,
                                      "description": "A list of conditions for a filter.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "action_condition": {
                                            "computed": true,
                                            "description": "The condition of the action desired in the filter.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "action": {
                                                  "computed": true,
                                                  "description": "The enumerated action to take.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "single"
                                            },
                                            "optional": true
                                          },
                                          "label_name_condition": {
                                            "computed": true,
                                            "description": "The label name of the condition.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "label_name": {
                                                  "computed": true,
                                                  "description": "The label name of the condition.",
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
                                    "requirement": {
                                      "computed": true,
                                      "description": "The requirement portion of the filter.",
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
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "redacted_fields": {
                          "computed": true,
                          "description": "Fields not to be included in the logs.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "method": {
                                "computed": true,
                                "description": "The method with which to match this rule.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "query_string": {
                                "computed": true,
                                "description": "The query string to find the resource to match this field to.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "single_header": {
                                "computed": true,
                                "description": "Header for the field to match.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "name": {
                                      "computed": true,
                                      "description": "The name of the header",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              },
                              "uri_path": {
                                "computed": true,
                                "description": "This is the URI path to match this rule to.",
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
                      "nesting_mode": "single"
                    },
                    "optional": true
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "resource_type": {
              "description": "Resource Type associated with the Telemetry Rule",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "selection_criteria": {
              "computed": true,
              "description": "Selection Criteria on resource level for rule application",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "telemetry_source_types": {
              "computed": true,
              "description": "The telemetry source types for a telemetry rule.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "telemetry_type": {
              "description": "Telemetry Type associated with the Telemetry Rule",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "rule_arn": {
        "computed": true,
        "description": "The arn of the telemetry rule",
        "description_kind": "plain",
        "type": "string"
      },
      "rule_name": {
        "description": "The name of the telemetry rule",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
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
    "description": "The AWS::ObservabilityAdmin::TelemetryRule resource defines a CloudWatch Observability Admin Telemetry Rule.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccObservabilityadminTelemetryRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccObservabilityadminTelemetryRule), &result)
	return &result
}
