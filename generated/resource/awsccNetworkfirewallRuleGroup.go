package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNetworkfirewallRuleGroup = `{
  "block": {
    "attributes": {
      "capacity": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "rule_group": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "reference_sets": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "ip_set_references": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "reference_arn": {
                          "computed": true,
                          "description": "A resource ARN.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
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
            "rule_variables": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "ip_sets": {
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
                  },
                  "port_sets": {
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
            "rules_source": {
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "rules_source_list": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "generated_rules_type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "target_types": {
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "targets": {
                          "description_kind": "plain",
                          "required": true,
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
                  "rules_string": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "stateful_rules": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "action": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "header": {
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "destination": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "destination_port": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "direction": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "protocol": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "source": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "source_port": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "required": true
                        },
                        "rule_options": {
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "keyword": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "settings": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "required": true
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "stateless_rules_and_custom_actions": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "custom_actions": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "action_definition": {
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "publish_metric_action": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "dimensions": {
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "value": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "list"
                                            },
                                            "required": true
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
                              "action_name": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "optional": true
                        },
                        "stateless_rules": {
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "priority": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
                              },
                              "rule_definition": {
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "actions": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    },
                                    "match_attributes": {
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "destination_ports": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "from_port": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "number"
                                                },
                                                "to_port": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "number"
                                                }
                                              },
                                              "nesting_mode": "list"
                                            },
                                            "optional": true
                                          },
                                          "destinations": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "address_definition": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "list"
                                            },
                                            "optional": true
                                          },
                                          "protocols": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": [
                                              "list",
                                              "number"
                                            ]
                                          },
                                          "source_ports": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "from_port": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "number"
                                                },
                                                "to_port": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "number"
                                                }
                                              },
                                              "nesting_mode": "list"
                                            },
                                            "optional": true
                                          },
                                          "sources": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "address_definition": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "list"
                                            },
                                            "optional": true
                                          },
                                          "tcp_flags": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "flags": {
                                                  "description_kind": "plain",
                                                  "required": true,
                                                  "type": [
                                                    "list",
                                                    "string"
                                                  ]
                                                },
                                                "masks": {
                                                  "computed": true,
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": [
                                                    "list",
                                                    "string"
                                                  ]
                                                }
                                              },
                                              "nesting_mode": "list"
                                            },
                                            "optional": true
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
                            "nesting_mode": "list"
                          },
                          "required": true
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
            "stateful_rule_options": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "rule_order": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
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
      "rule_group_arn": {
        "computed": true,
        "description": "A resource ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "rule_group_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "rule_group_name": {
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
          "nesting_mode": "set"
        },
        "optional": true
      },
      "type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource type definition for AWS::NetworkFirewall::RuleGroup",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccNetworkfirewallRuleGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworkfirewallRuleGroup), &result)
	return &result
}
