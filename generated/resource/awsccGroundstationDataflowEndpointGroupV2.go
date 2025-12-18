package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGroundstationDataflowEndpointGroupV2 = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "contact_post_pass_duration_seconds": {
        "computed": true,
        "description": "Amount of time, in seconds, after a contact ends that the Ground Station Dataflow Endpoint Group will be in a POSTPASS state. A Ground Station Dataflow Endpoint Group State Change event will be emitted when the Dataflow Endpoint Group enters and exits the POSTPASS state.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "contact_pre_pass_duration_seconds": {
        "computed": true,
        "description": "Amount of time, in seconds, before a contact starts that the Ground Station Dataflow Endpoint Group will be in a PREPASS state. A Ground Station Dataflow Endpoint Group State Change event will be emitted when the Dataflow Endpoint Group enters and exits the PREPASS state.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "dataflow_endpoint_group_v2_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint_details": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "downlink_aws_ground_station_agent_endpoint": {
              "computed": true,
              "description": "Information about DownlinkAwsGroundStationAgentEndpoint",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "agent_status": {
                    "computed": true,
                    "description": "The status of AgentEndpoint.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "audit_results": {
                    "computed": true,
                    "description": "The results of the audit.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "dataflow_details": {
                    "computed": true,
                    "description": "Dataflow details for downlink",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "agent_connection_details": {
                          "computed": true,
                          "description": "Connection details for downlink, from ground station to agent, and customer to agent",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "agent_ip_and_port_address": {
                                "computed": true,
                                "description": "Socket address of an uplink or downlink agent endpoint with a port range and an optional mtu.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "mtu": {
                                      "computed": true,
                                      "description": "Maximum transmission unit (MTU) size in bytes of a dataflow endpoint.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "socket_address": {
                                      "computed": true,
                                      "description": "A socket address with a port range.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "name": {
                                            "computed": true,
                                            "description": "IPv4 socket address.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "port_range": {
                                            "computed": true,
                                            "description": "Port range of a socket address.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "maximum": {
                                                  "computed": true,
                                                  "description": "A maximum value.",
                                                  "description_kind": "plain",
                                                  "type": "number"
                                                },
                                                "minimum": {
                                                  "computed": true,
                                                  "description": "A minimum value.",
                                                  "description_kind": "plain",
                                                  "type": "number"
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
                              "egress_address_and_port": {
                                "computed": true,
                                "description": "Socket address of an uplink or downlink agent endpoint with an optional mtu.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "mtu": {
                                      "computed": true,
                                      "description": "Maximum transmission unit (MTU) size in bytes of a dataflow endpoint.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "socket_address": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "name": {
                                            "computed": true,
                                            "description": "IPv4 socket address.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "port": {
                                            "computed": true,
                                            "description": "Port of a socket address.",
                                            "description_kind": "plain",
                                            "type": "number"
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
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "uplink_aws_ground_station_agent_endpoint": {
              "computed": true,
              "description": "Information about UplinkAwsGroundStationAgentEndpoint",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "agent_status": {
                    "computed": true,
                    "description": "The status of AgentEndpoint.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "audit_results": {
                    "computed": true,
                    "description": "The results of the audit.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "dataflow_details": {
                    "computed": true,
                    "description": "Dataflow details for uplink",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "agent_connection_details": {
                          "computed": true,
                          "description": "Connection details for uplink, from ground station to agent, and customer to agent",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "agent_ip_and_port_address": {
                                "computed": true,
                                "description": "Socket address of an uplink or downlink agent endpoint with a port range and an optional mtu.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "mtu": {
                                      "computed": true,
                                      "description": "Maximum transmission unit (MTU) size in bytes of a dataflow endpoint.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "socket_address": {
                                      "computed": true,
                                      "description": "A socket address with a port range.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "name": {
                                            "computed": true,
                                            "description": "IPv4 socket address.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "port_range": {
                                            "computed": true,
                                            "description": "Port range of a socket address.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "maximum": {
                                                  "computed": true,
                                                  "description": "A maximum value.",
                                                  "description_kind": "plain",
                                                  "type": "number"
                                                },
                                                "minimum": {
                                                  "computed": true,
                                                  "description": "A minimum value.",
                                                  "description_kind": "plain",
                                                  "type": "number"
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
                              "ingress_address_and_port": {
                                "computed": true,
                                "description": "Socket address of an uplink or downlink agent endpoint with an optional mtu.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "mtu": {
                                      "computed": true,
                                      "description": "Maximum transmission unit (MTU) size in bytes of a dataflow endpoint.",
                                      "description_kind": "plain",
                                      "type": "number"
                                    },
                                    "socket_address": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "name": {
                                            "computed": true,
                                            "description": "IPv4 socket address.",
                                            "description_kind": "plain",
                                            "type": "string"
                                          },
                                          "port": {
                                            "computed": true,
                                            "description": "Port of a socket address.",
                                            "description_kind": "plain",
                                            "type": "number"
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
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "name": {
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
      "endpoints": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "downlink_aws_ground_station_agent_endpoint": {
              "computed": true,
              "description": "Information about DownlinkAwsGroundStationAgentEndpoint used for create",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "dataflow_details": {
                    "computed": true,
                    "description": "Dataflow details for downlink",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "agent_connection_details": {
                          "computed": true,
                          "description": "Connection details for downlink, from ground station to agent, and customer to agent",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "agent_ip_and_port_address": {
                                "computed": true,
                                "description": "Socket address of an uplink or downlink agent endpoint with a port range and an optional mtu.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "mtu": {
                                      "computed": true,
                                      "description": "Maximum transmission unit (MTU) size in bytes of a dataflow endpoint.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "socket_address": {
                                      "computed": true,
                                      "description": "A socket address with a port range.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "name": {
                                            "computed": true,
                                            "description": "IPv4 socket address.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "port_range": {
                                            "computed": true,
                                            "description": "Port range of a socket address.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "maximum": {
                                                  "computed": true,
                                                  "description": "A maximum value.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                },
                                                "minimum": {
                                                  "computed": true,
                                                  "description": "A minimum value.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
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
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              },
                              "egress_address_and_port": {
                                "computed": true,
                                "description": "Socket address of an uplink or downlink agent endpoint with an optional mtu.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "mtu": {
                                      "computed": true,
                                      "description": "Maximum transmission unit (MTU) size in bytes of a dataflow endpoint.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "socket_address": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "name": {
                                            "computed": true,
                                            "description": "IPv4 socket address.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "port": {
                                            "computed": true,
                                            "description": "Port of a socket address.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
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
                  "name": {
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
            "uplink_aws_ground_station_agent_endpoint": {
              "computed": true,
              "description": "Information about UplinkAwsGroundStationAgentEndpoint used for create",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "dataflow_details": {
                    "computed": true,
                    "description": "Dataflow details for uplink",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "agent_connection_details": {
                          "computed": true,
                          "description": "Connection details for uplink, from ground station to agent, and customer to agent",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "agent_ip_and_port_address": {
                                "computed": true,
                                "description": "Socket address of an uplink or downlink agent endpoint with a port range and an optional mtu.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "mtu": {
                                      "computed": true,
                                      "description": "Maximum transmission unit (MTU) size in bytes of a dataflow endpoint.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "socket_address": {
                                      "computed": true,
                                      "description": "A socket address with a port range.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "name": {
                                            "computed": true,
                                            "description": "IPv4 socket address.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "port_range": {
                                            "computed": true,
                                            "description": "Port range of a socket address.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "maximum": {
                                                  "computed": true,
                                                  "description": "A maximum value.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
                                                },
                                                "minimum": {
                                                  "computed": true,
                                                  "description": "A minimum value.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "number"
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
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              },
                              "ingress_address_and_port": {
                                "computed": true,
                                "description": "Socket address of an uplink or downlink agent endpoint with an optional mtu.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "mtu": {
                                      "computed": true,
                                      "description": "Maximum transmission unit (MTU) size in bytes of a dataflow endpoint.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "socket_address": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "name": {
                                            "computed": true,
                                            "description": "IPv4 socket address.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "port": {
                                            "computed": true,
                                            "description": "Port of a socket address.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "number"
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
                  "name": {
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
          "nesting_mode": "list"
        },
        "optional": true
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
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS Ground Station DataflowEndpointGroupV2",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccGroundstationDataflowEndpointGroupV2Schema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGroundstationDataflowEndpointGroupV2), &result)
	return &result
}
