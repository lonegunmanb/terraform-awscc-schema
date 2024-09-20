package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2NetworkInsightsAnalysis = `{
  "block": {
    "attributes": {
      "alternate_path_hints": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "component_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "component_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "explanations": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "acl": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "acl_rule": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cidr": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "egress": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "port_range": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "from": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "to": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "protocol": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "rule_action": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "rule_number": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "address": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "addresses": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "attached_to": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "availability_zones": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "cidrs": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "classic_load_balancer_listener": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "instance_port": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "load_balancer_port": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "component": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "customer_gateway": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "destination": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "destination_vpc": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "direction": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "elastic_load_balancer_listener": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "explanation_code": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "ingress_route_table": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "internet_gateway": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "load_balancer_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "load_balancer_listener_port": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "load_balancer_target": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "address": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "availability_zone": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "instance": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "id": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "port": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "load_balancer_target_group": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "load_balancer_target_groups": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "load_balancer_target_port": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "missing_component": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "nat_gateway": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "network_interface": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "packet_field": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "port": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "port_ranges": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "from": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "to": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "prefix_list": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "protocols": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "route_table": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "route_table_route": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "destination_cidr": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "destination_prefix_list_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "egress_only_internet_gateway_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "gateway_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "instance_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "nat_gateway_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "network_interface_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "origin": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "transit_gateway_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "vpc_peering_connection_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "security_group": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "security_group_rule": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cidr": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "direction": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "port_range": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "from": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "to": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "prefix_list_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "protocol": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "security_group_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "security_groups": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "source_vpc": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
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
            "subnet": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "subnet_route_table": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "vpc": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "vpc_endpoint": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "vpc_peering_connection": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "vpn_connection": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "vpn_gateway": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "list"
        }
      },
      "filter_in_arns": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "forward_path_components": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "acl_rule": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cidr": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "egress": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "port_range": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "from": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "to": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "protocol": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "rule_action": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "rule_number": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "component": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "destination_vpc": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "inbound_header": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "destination_addresses": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "destination_port_ranges": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "from": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "to": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "protocol": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "source_addresses": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "source_port_ranges": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "from": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "to": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            },
            "outbound_header": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "destination_addresses": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "destination_port_ranges": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "from": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "to": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "protocol": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "source_addresses": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "source_port_ranges": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "from": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "to": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            },
            "route_table_route": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "destination_cidr": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "destination_prefix_list_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "egress_only_internet_gateway_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "gateway_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "instance_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "nat_gateway_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "network_interface_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "origin": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "transit_gateway_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "vpc_peering_connection_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "security_group_rule": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cidr": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "direction": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "port_range": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "from": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "to": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "prefix_list_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "protocol": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "security_group_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "sequence_number": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "source_vpc": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "subnet": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "vpc": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "list"
        }
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "network_insights_analysis_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "network_insights_analysis_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "network_insights_path_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_path_found": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "return_path_components": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "acl_rule": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cidr": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "egress": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "port_range": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "from": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "to": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "protocol": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "rule_action": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "rule_number": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "component": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "destination_vpc": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "inbound_header": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "destination_addresses": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "destination_port_ranges": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "from": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "to": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "protocol": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "source_addresses": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "source_port_ranges": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "from": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "to": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            },
            "outbound_header": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "destination_addresses": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "destination_port_ranges": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "from": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "to": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "protocol": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "source_addresses": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "source_port_ranges": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "from": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "to": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            },
            "route_table_route": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "destination_cidr": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "destination_prefix_list_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "egress_only_internet_gateway_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "gateway_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "instance_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "nat_gateway_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "network_interface_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "origin": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "transit_gateway_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "vpc_peering_connection_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "security_group_rule": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cidr": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "direction": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "port_range": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "from": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "to": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "prefix_list_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "protocol": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "security_group_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "sequence_number": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "source_vpc": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "subnet": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "vpc": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "list"
        }
      },
      "start_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status_message": {
        "computed": true,
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
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource schema for AWS::EC2::NetworkInsightsAnalysis",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2NetworkInsightsAnalysisSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2NetworkInsightsAnalysis), &result)
	return &result
}
