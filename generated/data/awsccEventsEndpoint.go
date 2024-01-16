package data

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
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "event_bus_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
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
        "description_kind": "plain",
        "type": "string"
      },
      "replication_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "state": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "role_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "routing_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "failover_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "primary": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "health_check": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "secondary": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "route": {
                          "computed": true,
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
            }
          },
          "nesting_mode": "single"
        }
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
    "description": "Data Source schema for AWS::Events::Endpoint",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEventsEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEventsEndpoint), &result)
	return &result
}
