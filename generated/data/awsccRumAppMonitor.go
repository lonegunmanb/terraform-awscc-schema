package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRumAppMonitor = `{
  "block": {
    "attributes": {
      "app_monitor_configuration": {
        "computed": true,
        "description": "AppMonitor configuration",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "allow_cookies": {
              "computed": true,
              "description": "If you set this to true, the RUM web client sets two cookies, a session cookie and a user cookie. The cookies allow the RUM web client to collect data relating to the number of users an application has and the behavior of the application across a sequence of events. Cookies are stored in the top-level domain of the current page.",
              "description_kind": "plain",
              "type": "bool"
            },
            "enable_x_ray": {
              "computed": true,
              "description": "If you set this to true, RUM enables xray tracing for the user sessions that RUM samples. RUM adds an xray trace header to allowed HTTP requests. It also records an xray segment for allowed HTTP requests. You can see traces and segments from these user sessions in the xray console and the CW ServiceLens console.",
              "description_kind": "plain",
              "type": "bool"
            },
            "excluded_pages": {
              "computed": true,
              "description": "A list of URLs in your website or application to exclude from RUM data collection. You can't include both ExcludedPages and IncludedPages in the same operation.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "favorite_pages": {
              "computed": true,
              "description": "A list of pages in the RUM console that are to be displayed with a favorite icon.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "guest_role_arn": {
              "computed": true,
              "description": "The ARN of the guest IAM role that is attached to the identity pool that is used to authorize the sending of data to RUM.",
              "description_kind": "plain",
              "type": "string"
            },
            "identity_pool_id": {
              "computed": true,
              "description": "The ID of the identity pool that is used to authorize the sending of data to RUM.",
              "description_kind": "plain",
              "type": "string"
            },
            "included_pages": {
              "computed": true,
              "description": "If this app monitor is to collect data from only certain pages in your application, this structure lists those pages. You can't include both ExcludedPages and IncludedPages in the same operation.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "metric_destinations": {
              "computed": true,
              "description": "An array of structures which define the destinations and the metrics that you want to send.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "destination": {
                    "computed": true,
                    "description": "Defines the destination to send the metrics to. Valid values are CloudWatch and Evidently. If you specify Evidently, you must also specify the ARN of the Evidently experiment that is to be the destination and an IAM role that has permission to write to the experiment.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "destination_arn": {
                    "computed": true,
                    "description": "Use this parameter only if Destination is Evidently. This parameter specifies the ARN of the Evidently experiment that will receive the extended metrics.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "iam_role_arn": {
                    "computed": true,
                    "description": "This parameter is required if Destination is Evidently. If Destination is CloudWatch, do not use this parameter.\n\nThis parameter specifies the ARN of an IAM role that RUM will assume to write to the Evidently experiment that you are sending metrics to. This role must have permission to write to that experiment.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "metric_definitions": {
                    "computed": true,
                    "description": "An array of structures which define the metrics that you want to send.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "dimension_keys": {
                          "computed": true,
                          "description": "Use this field only if you are sending the metric to CloudWatch.\n\nThis field is a map of field paths to dimension names. It defines the dimensions to associate with this metric in CloudWatch. For extended metrics, valid values for the entries in this field are the following:\n\n\"metadata.pageId\": \"PageId\"\n\n\"metadata.browserName\": \"BrowserName\"\n\n\"metadata.deviceType\": \"DeviceType\"\n\n\"metadata.osName\": \"OSName\"\n\n\"metadata.countryCode\": \"CountryCode\"\n\n\"event_details.fileType\": \"FileType\"\n\nAll dimensions listed in this field must also be included in EventPattern.",
                          "description_kind": "plain",
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "event_pattern": {
                          "computed": true,
                          "description": "The pattern that defines the metric, specified as a JSON object. RUM checks events that happen in a user's session against the pattern, and events that match the pattern are sent to the metric destination.\n\nWhen you define extended metrics, the metric definition is not valid if EventPattern is omitted.\n\nExample event patterns:\n\n'{ \"event_type\": [\"com.amazon.rum.js_error_event\"], \"metadata\": { \"browserName\": [ \"Chrome\", \"Safari\" ], } }'\n\n'{ \"event_type\": [\"com.amazon.rum.performance_navigation_event\"], \"metadata\": { \"browserName\": [ \"Chrome\", \"Firefox\" ] }, \"event_details\": { \"duration\": [{ \"numeric\": [ \"\u003c\", 2000 ] }] } }'\n\n'{ \"event_type\": [\"com.amazon.rum.performance_navigation_event\"], \"metadata\": { \"browserName\": [ \"Chrome\", \"Safari\" ], \"countryCode\": [ \"US\" ] }, \"event_details\": { \"duration\": [{ \"numeric\": [ \"\u003e=\", 2000, \"\u003c\", 8000 ] }] } }'\n\nIf the metrics destination' is CloudWatch and the event also matches a value in DimensionKeys, then the metric is published with the specified dimensions.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "name": {
                          "computed": true,
                          "description": "The name for the metric that is defined in this structure. For extended metrics, valid values are the following:\n\nPerformanceNavigationDuration\n\nPerformanceResourceDuration\n\nNavigationSatisfiedTransaction\n\nNavigationToleratedTransaction\n\nNavigationFrustratedTransaction\n\nWebVitalsCumulativeLayoutShift\n\nWebVitalsFirstInputDelay\n\nWebVitalsLargestContentfulPaint\n\nJsErrorCount\n\nHttpErrorCount\n\nSessionCount",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "namespace": {
                          "computed": true,
                          "description": "The namespace used by CloudWatch Metrics for the metric that is defined in this structure",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "unit_label": {
                          "computed": true,
                          "description": "The CloudWatch metric unit to use for this metric. If you omit this field, the metric is recorded with no unit.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "value_key": {
                          "computed": true,
                          "description": "The field within the event object that the metric value is sourced from.\n\nIf you omit this field, a hardcoded value of 1 is pushed as the metric value. This is useful if you just want to count the number of events that the filter catches.\n\nIf this metric is sent to Evidently, this field will be passed to Evidently raw and Evidently will handle data extraction from the event.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "set"
                    }
                  }
                },
                "nesting_mode": "set"
              }
            },
            "session_sample_rate": {
              "computed": true,
              "description": "Specifies the percentage of user sessions to use for RUM data collection. Choosing a higher percentage gives you more data but also incurs more costs. The number you specify is the percentage of user sessions that will be used. If you omit this parameter, the default of 10 is used.",
              "description_kind": "plain",
              "type": "number"
            },
            "telemetries": {
              "computed": true,
              "description": "An array that lists the types of telemetry data that this app monitor is to collect.",
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
      "app_monitor_id": {
        "computed": true,
        "description": "The unique ID of the new app monitor.",
        "description_kind": "plain",
        "type": "string"
      },
      "custom_events": {
        "computed": true,
        "description": "AppMonitor custom events configuration",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "status": {
              "computed": true,
              "description": "Indicates whether AppMonitor accepts custom events.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "cw_log_enabled": {
        "computed": true,
        "description": "Data collected by RUM is kept by RUM for 30 days and then deleted. This parameter specifies whether RUM sends a copy of this telemetry data to CWLlong in your account. This enables you to keep the telemetry data for more than 30 days, but it does incur CWLlong charges. If you omit this parameter, the default is false",
        "description_kind": "plain",
        "type": "bool"
      },
      "deobfuscation_configuration": {
        "computed": true,
        "description": "A structure that contains the configuration for how an app monitor can deobfuscate stack traces.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "java_script_source_maps": {
              "computed": true,
              "description": "A structure that contains the configuration for how an app monitor can unminify JavaScript error stack traces using source maps.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "s3_uri": {
                    "computed": true,
                    "description": "The S3Uri of the bucket or folder that stores the source map files. It is required if status is ENABLED.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "status": {
                    "computed": true,
                    "description": "Specifies whether JavaScript error stack traces should be unminified for this app monitor. The default is for JavaScript error stack trace unminification to be DISABLED",
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
      "domain": {
        "computed": true,
        "description": "The top-level internet domain name for which your application has administrative authority. The CreateAppMonitor requires either the domain or the domain list.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_list": {
        "computed": true,
        "description": "The top-level internet domain names for which your application has administrative authority. The CreateAppMonitor requires either the domain or the domain list.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "A name for the app monitor",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_policy": {
        "computed": true,
        "description": "A structure that defines resource policy attached to your app monitor.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "policy_document": {
              "computed": true,
              "description": "The JSON to use as the resource policy. The document can be up to 4 KB in size. ",
              "description_kind": "plain",
              "type": "string"
            },
            "policy_revision_id": {
              "computed": true,
              "description": "A string value that you can use to conditionally update your policy. You can provide the revision ID of your existing policy to make mutating requests against that policy. \n\n When you assign a policy revision ID, then later requests about that policy will be rejected with an InvalidPolicyRevisionIdException error if they don't provide the correct current revision ID.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description": "Assigns one or more tags (key-value pairs) to the app monitor. Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a user permission to access or change only resources with certain tag values. Tags don't have any semantic meaning to AWS and are interpreted strictly as strings of characters.You can associate as many as 50 tags with an app monitor.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::RUM::AppMonitor",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRumAppMonitorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRumAppMonitor), &result)
	return &result
}
