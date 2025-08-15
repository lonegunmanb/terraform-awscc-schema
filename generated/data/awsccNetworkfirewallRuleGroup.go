package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNetworkfirewallRuleGroup = `{
  "block": {
    "attributes": {
      "capacity": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "description": {
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
                          "type": "string"
                        }
                      },
                      "nesting_mode": "map"
                    }
                  }
                },
                "nesting_mode": "single"
              }
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
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "nesting_mode": "map"
                    }
                  },
                  "port_sets": {
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
            "rules_source": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "rules_source_list": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "generated_rules_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "target_types": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "targets": {
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
                  "rules_string": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "stateful_rules": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "action": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "header": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "destination": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "destination_port": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "direction": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "protocol": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "source": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "source_port": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "rule_options": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "keyword": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "settings": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "nesting_mode": "list"
                          }
                        }
                      },
                      "nesting_mode": "list"
                    }
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
                        "stateless_rules": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "priority": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "number"
                              },
                              "rule_definition": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "actions": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    },
                                    "match_attributes": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "destination_ports": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "from_port": {
                                                  "computed": true,
                                                  "description_kind": "plain",
                                                  "type": "number"
                                                },
                                                "to_port": {
                                                  "computed": true,
                                                  "description_kind": "plain",
                                                  "type": "number"
                                                }
                                              },
                                              "nesting_mode": "list"
                                            }
                                          },
                                          "destinations": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "address_definition": {
                                                  "computed": true,
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "list"
                                            }
                                          },
                                          "protocols": {
                                            "computed": true,
                                            "description_kind": "plain",
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
                                                  "computed": true,
                                                  "description_kind": "plain",
                                                  "type": "number"
                                                },
                                                "to_port": {
                                                  "computed": true,
                                                  "description_kind": "plain",
                                                  "type": "number"
                                                }
                                              },
                                              "nesting_mode": "list"
                                            }
                                          },
                                          "sources": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "address_definition": {
                                                  "computed": true,
                                                  "description_kind": "plain",
                                                  "type": "string"
                                                }
                                              },
                                              "nesting_mode": "list"
                                            }
                                          },
                                          "tcp_flags": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "flags": {
                                                  "computed": true,
                                                  "description_kind": "plain",
                                                  "type": [
                                                    "list",
                                                    "string"
                                                  ]
                                                },
                                                "masks": {
                                                  "computed": true,
                                                  "description_kind": "plain",
                                                  "type": [
                                                    "list",
                                                    "string"
                                                  ]
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
            "stateful_rule_options": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "rule_order": {
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
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "summary_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "rule_options": {
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
      },
      "type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::NetworkFirewall::RuleGroup",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccNetworkfirewallRuleGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworkfirewallRuleGroup), &result)
	return &result
}
