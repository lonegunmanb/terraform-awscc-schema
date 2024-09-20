package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53HealthCheck = `{
  "block": {
    "attributes": {
      "health_check_config": {
        "description": "A complex type that contains information about the health check.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "alarm_identifier": {
              "computed": true,
              "description": "A complex type that identifies the CloudWatch alarm that you want Amazon Route 53 health checkers to use to determine whether the specified health check is healthy.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
                    "computed": true,
                    "description": "The name of the CloudWatch alarm that you want Amazon Route 53 health checkers to use to determine whether this health check is healthy.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "region": {
                    "computed": true,
                    "description": "For the CloudWatch alarm that you want Route 53 health checkers to use to determine whether this health check is healthy, the region that the alarm was created in.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "child_health_checks": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "enable_sni": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "failure_threshold": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "fully_qualified_domain_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "health_threshold": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "insufficient_data_health_status": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "inverted": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "ip_address": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "measure_latency": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "port": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "regions": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "request_interval": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "resource_path": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "routing_control_arn": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "search_string": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "health_check_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "health_check_tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
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
      }
    },
    "description": "Resource schema for AWS::Route53::HealthCheck.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRoute53HealthCheckSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53HealthCheck), &result)
	return &result
}
