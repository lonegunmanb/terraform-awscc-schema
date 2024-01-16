package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccXraySamplingRule = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "rule_arn": {
        "computed": true,
        "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
        "description_kind": "plain",
        "type": "string"
      },
      "rule_name": {
        "computed": true,
        "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sampling_rule": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "attributes": {
              "computed": true,
              "description": "Matches attributes derived from the request.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "fixed_rate": {
              "description": "The percentage of matching requests to instrument, after the reservoir is exhausted.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "host": {
              "description": "Matches the hostname from a request URL.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "http_method": {
              "description": "Matches the HTTP method from a request URL.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "priority": {
              "description": "The priority of the sampling rule.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "reservoir_size": {
              "description": "A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "resource_arn": {
              "description": "Matches the ARN of the AWS resource on which the service runs.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "rule_arn": {
              "computed": true,
              "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "rule_name": {
              "computed": true,
              "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "service_name": {
              "description": "Matches the name that the service uses to identify itself in segments.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "service_type": {
              "description": "Matches the origin that the service uses to identify its type in segments.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "url_path": {
              "description": "Matches the path from a request URL.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "version": {
              "computed": true,
              "description": "The version of the sampling rule format (1)",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "sampling_rule_record": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "created_at": {
              "computed": true,
              "description": "When the rule was created, in Unix time seconds.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "modified_at": {
              "computed": true,
              "description": "When the rule was modified, in Unix time seconds.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sampling_rule": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "attributes": {
                    "computed": true,
                    "description": "Matches attributes derived from the request.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "fixed_rate": {
                    "description": "The percentage of matching requests to instrument, after the reservoir is exhausted.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "host": {
                    "description": "Matches the hostname from a request URL.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "http_method": {
                    "description": "Matches the HTTP method from a request URL.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "priority": {
                    "description": "The priority of the sampling rule.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "reservoir_size": {
                    "description": "A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "resource_arn": {
                    "description": "Matches the ARN of the AWS resource on which the service runs.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "rule_arn": {
                    "computed": true,
                    "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "rule_name": {
                    "computed": true,
                    "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "service_name": {
                    "description": "Matches the name that the service uses to identify itself in segments.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "service_type": {
                    "description": "Matches the origin that the service uses to identify its type in segments.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "url_path": {
                    "description": "Matches the path from a request URL.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "version": {
                    "computed": true,
                    "description": "The version of the sampling rule format (1)",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
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
      "sampling_rule_update": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "attributes": {
              "computed": true,
              "description": "Matches attributes derived from the request.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "fixed_rate": {
              "computed": true,
              "description": "The percentage of matching requests to instrument, after the reservoir is exhausted.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "host": {
              "computed": true,
              "description": "Matches the hostname from a request URL.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "http_method": {
              "computed": true,
              "description": "Matches the HTTP method from a request URL.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "priority": {
              "computed": true,
              "description": "The priority of the sampling rule.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "reservoir_size": {
              "computed": true,
              "description": "A fixed number of matching requests to instrument per second, prior to applying the fixed rate. The reservoir is not used directly by services, but applies to all services using the rule collectively.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "resource_arn": {
              "computed": true,
              "description": "Matches the ARN of the AWS resource on which the service runs.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "rule_arn": {
              "computed": true,
              "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "rule_name": {
              "computed": true,
              "description": "The ARN of the sampling rule. Specify a rule by either name or ARN, but not both.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "service_name": {
              "computed": true,
              "description": "Matches the name that the service uses to identify itself in segments.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "service_type": {
              "computed": true,
              "description": "Matches the origin that the service uses to identify its type in segments.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "url_path": {
              "computed": true,
              "description": "Matches the path from a request URL.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the tag.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "This schema provides construct and validation rules for AWS-XRay SamplingRule resource parameters.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccXraySamplingRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccXraySamplingRule), &result)
	return &result
}
