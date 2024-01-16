package resource

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
        "optional": true,
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
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "orgunit": {
              "computed": true,
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
        "optional": true
      },
      "exclude_resource_tags": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
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
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "orgunit": {
              "computed": true,
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
        "optional": true
      },
      "policy_description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "policy_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "remediation_enabled": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "resource_set_ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "resource_type": {
        "computed": true,
        "description": "An AWS resource type",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_type_list": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "resources_clean_up": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "security_service_policy_data": {
        "description": "Firewall security service policy data.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "managed_service_data": {
              "computed": true,
              "description": "Firewall managed service data.",
              "description_kind": "plain",
              "optional": true,
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
                          "description": "Firewall deployment mode.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "third_party_firewall_policy": {
                    "computed": true,
                    "description": "Third party firewall policy.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "firewall_deployment_model": {
                          "description": "Firewall deployment mode.",
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
                "nesting_mode": "single"
              },
              "optional": true
            },
            "type": {
              "description": "Firewall policy type.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
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
    "description": "Creates an AWS Firewall Manager policy.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccFmsPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccFmsPolicy), &result)
	return &result
}
