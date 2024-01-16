package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWafv2LoggingConfiguration = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "log_destination_configs": {
        "description": "The Amazon Resource Names (ARNs) of the logging destinations that you want to associate with the web ACL.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "logging_filter": {
        "computed": true,
        "description": "Filtering that specifies which web requests are kept in the logs and which are dropped. You can filter on the rule action and on the web request labels that were applied by matching rules during web ACL evaluation.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "default_behavior": {
              "description": "Default handling for logs that don't match any of the specified filtering conditions.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "filters": {
              "description": "The filters that you want to apply to the logs.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "behavior": {
                    "description": "How to handle logs that satisfy the filter's conditions and requirement. ",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "conditions": {
                    "description": "Match conditions for the filter.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "action_condition": {
                          "computed": true,
                          "description": "A single action condition.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "action": {
                                "description": "Logic to apply to the filtering conditions. You can specify that, in order to satisfy the filter, a log must match all conditions or must match at least one condition.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "label_name_condition": {
                          "computed": true,
                          "description": "A single label name condition.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "label_name": {
                                "description": "The label name that a log record must contain in order to meet the condition. This must be a fully qualified label name. Fully qualified labels have a prefix, optional namespaces, and label name. The prefix identifies the rule group or web ACL context of the rule that added the label. ",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "required": true
                  },
                  "requirement": {
                    "description": "Logic to apply to the filtering conditions. You can specify that, in order to satisfy the filter, a log must match all conditions or must match at least one condition.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "required": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "managed_by_firewall_manager": {
        "computed": true,
        "description": "Indicates whether the logging configuration was created by AWS Firewall Manager, as part of an AWS WAF policy configuration. If true, only Firewall Manager can modify or delete the configuration.",
        "description_kind": "plain",
        "type": "bool"
      },
      "redacted_fields": {
        "computed": true,
        "description": "The parts of the request that you want to keep out of the logs. For example, if you redact the HEADER field, the HEADER field in the firehose will be xxx.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "json_body": {
              "computed": true,
              "description": "Inspect the request body as JSON. The request body immediately follows the request headers. This is the part of a request that contains any additional data that you want to send to your web server as the HTTP request body, such as data from a form. ",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "invalid_fallback_behavior": {
                    "computed": true,
                    "description": "What AWS WAF should do if it fails to completely parse the JSON body.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "match_pattern": {
                    "description": "The patterns to look for in the JSON body. AWS WAF inspects the results of these pattern matches against the rule inspection criteria. ",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "all": {
                          "computed": true,
                          "description": "Match all of the elements. See also MatchScope in JsonBody. You must specify either this setting or the IncludedPaths setting, but not both.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "included_paths": {
                          "computed": true,
                          "description": "Match only the specified include paths. See also MatchScope in JsonBody.",
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
                    "required": true
                  },
                  "match_scope": {
                    "description": "The parts of the JSON to match against using the MatchPattern. If you specify All, AWS WAF matches against keys and values. ",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "method": {
              "computed": true,
              "description": "Inspect the HTTP method. The method indicates the type of operation that the request is asking the origin to perform. ",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "query_string": {
              "computed": true,
              "description": "Inspect the query string. This is the part of a URL that appears after a ? character, if any. ",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "single_header": {
              "computed": true,
              "description": "Inspect a single header. Provide the name of the header to inspect, for example, User-Agent or Referer. This setting isn't case sensitive.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
                    "description": "The name of the query header to inspect.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "uri_path": {
              "computed": true,
              "description": "Inspect the request URI path. This is the part of a web request that identifies a resource, for example, /images/daily-ad.jpg. ",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "resource_arn": {
        "description": "The Amazon Resource Name (ARN) of the web ACL that you want to associate with LogDestinationConfigs.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "A WAFv2 Logging Configuration Resource Provider",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccWafv2LoggingConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWafv2LoggingConfiguration), &result)
	return &result
}
