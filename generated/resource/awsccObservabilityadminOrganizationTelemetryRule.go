package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccObservabilityadminOrganizationTelemetryRule = `{
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
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "resource_type": {
              "description": "Resource Type associated with the Organization Telemetry Rule",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "scope": {
              "computed": true,
              "description": "Selection Criteria on scope level for rule application",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "selection_criteria": {
              "computed": true,
              "description": "Selection Criteria on resource level for rule application",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "telemetry_type": {
              "description": "Telemetry Type associated with the Organization Telemetry Rule",
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
        "description": "The arn of the organization telemetry rule",
        "description_kind": "plain",
        "type": "string"
      },
      "rule_name": {
        "description": "The name of the organization telemetry rule",
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
    "description": "The AWS::ObservabilityAdmin::OrganizationTelemetryRule resource defines a CloudWatch Observability Admin Organization Telemetry Rule.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccObservabilityadminOrganizationTelemetryRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccObservabilityadminOrganizationTelemetryRule), &result)
	return &result
}
