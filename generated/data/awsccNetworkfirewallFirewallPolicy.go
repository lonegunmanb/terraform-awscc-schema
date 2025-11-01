package data

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
        "type": "string"
      },
      "firewall_policy": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enable_tls_session_holding": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
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
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "nesting_mode": "map"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            },
            "stateful_default_actions": {
              "computed": true,
              "description_kind": "plain",
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
                  "flow_timeouts": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "tcp_idle_timeout_seconds": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "rule_order": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "stream_exception_policy": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "stateful_rule_group_references": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "deep_threat_inspection": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "override": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "action": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "priority": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "resource_arn": {
                    "computed": true,
                    "description": "A resource ARN.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
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
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "list"
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
                  "action_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "stateless_default_actions": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "stateless_fragment_default_actions": {
              "computed": true,
              "description_kind": "plain",
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
                    "type": "number"
                  },
                  "resource_arn": {
                    "computed": true,
                    "description": "A resource ARN.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "tls_inspection_configuration_arn": {
              "computed": true,
              "description": "A resource ARN.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
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
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
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
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::NetworkFirewall::FirewallPolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccNetworkfirewallFirewallPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworkfirewallFirewallPolicy), &result)
	return &result
}
