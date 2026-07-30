package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccServicediscoveryService = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the service.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description for the service.",
        "description_kind": "plain",
        "type": "string"
      },
      "dns_config": {
        "computed": true,
        "description": "DNS-related configurations for the service.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "dns_records": {
              "computed": true,
              "description": "A list of DNS records associated with the service.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "ttl": {
                    "computed": true,
                    "description": "The time-to-live (TTL) for the DNS record.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "type": {
                    "computed": true,
                    "description": "The DNS record type (e.g., A, AAAA, SRV).",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "namespace_id": {
              "computed": true,
              "description": "The ID of the namespace for the DNS configuration.",
              "description_kind": "plain",
              "type": "string"
            },
            "routing_policy": {
              "computed": true,
              "description": "The routing policy to use for DNS queries.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "health_check_config": {
        "computed": true,
        "description": "Settings for health checks. Used when routing is DNS-based.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "failure_threshold": {
              "computed": true,
              "description": "The number of consecutive health check failures that must occur before declaring the service unhealthy.",
              "description_kind": "plain",
              "type": "number"
            },
            "resource_path": {
              "computed": true,
              "description": "The path to ping on the service for health checks.",
              "description_kind": "plain",
              "type": "string"
            },
            "type": {
              "computed": true,
              "description": "The type of health check (e.g., HTTP, HTTPS, TCP).",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "health_check_custom_config": {
        "computed": true,
        "description": "Settings for custom health checks.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "failure_threshold": {
              "computed": true,
              "description": "The number of consecutive health check failures required before the service is considered unhealthy.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the service.",
        "description_kind": "plain",
        "type": "string"
      },
      "namespace_id": {
        "computed": true,
        "description": "The ID of the namespace in which the service is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "service_attributes": {
        "computed": true,
        "description": "A string map that contains attributes and values for the service. You can specify a maximum of 30 key-value pairs.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "service_id": {
        "computed": true,
        "description": "The unique identifier for the service.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to associate with the service.",
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
              "description": "The value of the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "type": {
        "computed": true,
        "description": "The type of service. Supported values are HTTP or DNS.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ServiceDiscovery::Service",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccServicediscoveryServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccServicediscoveryService), &result)
	return &result
}
