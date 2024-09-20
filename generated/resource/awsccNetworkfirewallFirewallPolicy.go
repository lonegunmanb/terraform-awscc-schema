package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNetworkfirewallFirewallPolicy = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "firewall_policy": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "policy_variables": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "rule_variables": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "definition": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "nesting_mode": "map"
                    },
                    "optional": true
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "stateful_default_actions": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "stateful_engine_options": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "rule_order": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "stream_exception_policy": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "stateful_rule_group_references": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "override": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "action": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "priority": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "resource_arn": {
                    "computed": true,
                    "description": "A resource ARN.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "stateless_custom_actions": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "action_definition": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "publish_metric_action": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "dimensions": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
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
                  "action_name": {
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
            "stateless_default_actions": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "stateless_fragment_default_actions": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "stateless_rule_group_references": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "priority": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "resource_arn": {
                    "computed": true,
                    "description": "A resource ARN.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "tls_inspection_configuration_arn": {
              "computed": true,
              "description": "A resource ARN.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "firewall_policy_arn": {
        "computed": true,
        "description": "A resource ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "firewall_policy_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "firewall_policy_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
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
    "description": "Resource type definition for AWS::NetworkFirewall::FirewallPolicy",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccNetworkfirewallFirewallPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworkfirewallFirewallPolicy), &result)
	return &result
}
