package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53HealthCheck = `{
  "block": {
    "attributes": {
      "health_check_config": {
        "computed": true,
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
                    "type": "string"
                  },
                  "region": {
                    "computed": true,
                    "description": "For the CloudWatch alarm that you want Route 53 health checkers to use to determine whether this health check is healthy, the region that the alarm was created in.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "child_health_checks": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "enable_sni": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "endpoint_type": {
              "computed": true,
              "description": "Whether the health check targets a public endpoint or a private endpoint inside a VPC. Defaults to PUBLIC. Adding this property to an existing health check's template replaces the health check; leave it unset for existing public health checks.",
              "description_kind": "plain",
              "type": "string"
            },
            "failure_threshold": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "fully_qualified_domain_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "health_check_source_id": {
              "computed": true,
              "description": "The ID of the health check source that Route 53 uses to reach a private endpoint. Required when EndpointType is PRIVATE.",
              "description_kind": "plain",
              "type": "string"
            },
            "health_threshold": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "insufficient_data_health_status": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "inverted": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "ip_address": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "measure_latency": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "port": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "regions": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "request_interval": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "resource_path": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "routing_control_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "search_string": {
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
          "nesting_mode": "single"
        }
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
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Route53::HealthCheck",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRoute53HealthCheckSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53HealthCheck), &result)
	return &result
}
