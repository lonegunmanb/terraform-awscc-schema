package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGroundstationDataflowEndpointGroup = `{
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
      "dataflow_endpoint_group_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint_details": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "aws_ground_station_agent_endpoint": {
              "computed": true,
              "description": "Information about AwsGroundStationAgentEndpoint.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "agent_status": {
                    "computed": true,
                    "description": "The status of AgentEndpoint.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "audit_results": {
                    "computed": true,
                    "description": "The results of the audit.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "egress_address": {
                    "computed": true,
                    "description": "Egress address of AgentEndpoint with an optional mtu.",
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
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "port": {
                                "computed": true,
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
                  },
                  "ingress_address": {
                    "computed": true,
                    "description": "Ingress address of AgentEndpoint with a port range and an optional mtu.",
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
            "endpoint": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "address": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
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
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "mtu": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
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
            "security_details": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "role_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "security_group_ids": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "subnet_ids": {
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
            }
          },
          "nesting_mode": "list"
        },
        "required": true
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
    "description": "AWS Ground Station DataflowEndpointGroup schema for CloudFormation",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccGroundstationDataflowEndpointGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGroundstationDataflowEndpointGroup), &result)
	return &result
}
