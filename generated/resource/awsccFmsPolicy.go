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
        "description": "Uniquely identifies the resource.",
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
      "policy_id": {
        "computed": true,
        "description_kind": "plain",
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
                  "network_acl_common_policy": {
                    "computed": true,
                    "description": "Network ACL common policy.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "network_acl_entry_set": {
                          "description": "Network ACL entry set.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "first_entries": {
                                "computed": true,
                                "description": "NetworkAcl entry list.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "cidr_block": {
                                      "computed": true,
                                      "description": "CIDR block.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "egress": {
                                      "description": "Whether the entry is an egress entry.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "bool"
                                    },
                                    "icmp_type_code": {
                                      "computed": true,
                                      "description": "ICMP type and code.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "code": {
                                            "description": "Code.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "number"
                                          },
                                          "type": {
                                            "description": "Type.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "number"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "ipv_6_cidr_block": {
                                      "computed": true,
                                      "description": "IPv6 CIDR block.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "port_range": {
                                      "computed": true,
                                      "description": "Port range.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "from": {
                                            "description": "From Port.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "number"
                                          },
                                          "to": {
                                            "description": "To Port.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "number"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "protocol": {
                                      "description": "Protocol.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "rule_action": {
                                      "description": "Rule Action.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "list"
                                },
                                "optional": true
                              },
                              "force_remediate_for_first_entries": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "bool"
                              },
                              "force_remediate_for_last_entries": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "bool"
                              },
                              "last_entries": {
                                "computed": true,
                                "description": "NetworkAcl entry list.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "cidr_block": {
                                      "computed": true,
                                      "description": "CIDR block.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "egress": {
                                      "description": "Whether the entry is an egress entry.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "bool"
                                    },
                                    "icmp_type_code": {
                                      "computed": true,
                                      "description": "ICMP type and code.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "code": {
                                            "description": "Code.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "number"
                                          },
                                          "type": {
                                            "description": "Type.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "number"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "ipv_6_cidr_block": {
                                      "computed": true,
                                      "description": "IPv6 CIDR block.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "port_range": {
                                      "computed": true,
                                      "description": "Port range.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "from": {
                                            "description": "From Port.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "number"
                                          },
                                          "to": {
                                            "description": "To Port.",
                                            "description_kind": "plain",
                                            "required": true,
                                            "type": "number"
                                          }
                                        },
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    },
                                    "protocol": {
                                      "description": "Protocol.",
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "rule_action": {
                                      "description": "Rule Action.",
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
                            "nesting_mode": "single"
                          },
                          "required": true
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
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
