package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccVpclatticeListener = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "default_action": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "fixed_response": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "status_code": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "forward": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "target_groups": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "target_group_identifier": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "weight": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "nesting_mode": "list"
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
        "required": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "listener_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "port": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "protocol": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "service_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "service_identifier": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
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
    "description": "Creates a listener for a service. Before you start using your Amazon VPC Lattice service, you must add one or more listeners. A listener is a process that checks for connection requests to your services.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccVpclatticeListenerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccVpclatticeListener), &result)
	return &result
}
