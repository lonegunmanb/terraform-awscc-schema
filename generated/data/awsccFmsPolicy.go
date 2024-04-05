package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccFmsPolicy = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "A resource ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "delete_all_policy_resources": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "exclude_map": {
        "computed": true,
        "description": "An FMS includeMap or excludeMap.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "account": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "orgunit": {
              "computed": true,
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
      "exclude_resource_tags": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "include_map": {
        "computed": true,
        "description": "An FMS includeMap or excludeMap.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "account": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "orgunit": {
              "computed": true,
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
      "policy_description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "policy_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "policy_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "remediation_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "resource_set_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "resource_tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "resource_type": {
        "computed": true,
        "description": "An AWS resource type",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_type_list": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "resources_clean_up": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "security_service_policy_data": {
        "computed": true,
        "description": "Firewall security service policy data.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "managed_service_data": {
              "computed": true,
              "description": "Firewall managed service data.",
              "description_kind": "plain",
              "type": "string"
            },
            "policy_option": {
              "computed": true,
              "description": "Firewall policy option.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "network_firewall_policy": {
                    "computed": true,
                    "description": "Network firewall policy.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "firewall_deployment_model": {
                          "computed": true,
                          "description": "Firewall deployment mode.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "third_party_firewall_policy": {
                    "computed": true,
                    "description": "Third party firewall policy.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "firewall_deployment_model": {
                          "computed": true,
                          "description": "Firewall deployment mode.",
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
            "type": {
              "computed": true,
              "description": "Firewall policy type.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::FMS::Policy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccFmsPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccFmsPolicy), &result)
	return &result
}
