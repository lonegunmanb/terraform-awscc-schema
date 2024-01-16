package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEventsEndpoint = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "endpoint_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "event_buses": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "event_bus_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "replication_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "state": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "role_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "routing_config": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "failover_config": {
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "primary": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "health_check": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  },
                  "secondary": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "route": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "state_reason": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Events::Endpoint.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEventsEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEventsEndpoint), &result)
	return &result
}
