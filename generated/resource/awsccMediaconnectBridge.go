package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMediaconnectBridge = `{
  "block": {
    "attributes": {
      "bridge_arn": {
        "computed": true,
        "description": "The Amazon Resource Number (ARN) of the bridge.",
        "description_kind": "plain",
        "type": "string"
      },
      "bridge_state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "egress_gateway_bridge": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "max_bitrate": {
              "computed": true,
              "description": "The maximum expected bitrate of the egress bridge.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "ingress_gateway_bridge": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "max_bitrate": {
              "computed": true,
              "description": "The maximum expected bitrate of the ingress bridge.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_outputs": {
              "computed": true,
              "description": "The maximum number of outputs on the ingress bridge.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "name": {
        "description": "The name of the bridge.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "outputs": {
        "computed": true,
        "description": "The outputs on this bridge.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "network_output": {
              "computed": true,
              "description": "The output of the bridge. A network output is delivered to your premises.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "ip_address": {
                    "computed": true,
                    "description": "The network output IP Address.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "The network output name.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "network_name": {
                    "computed": true,
                    "description": "The network output's gateway network name.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "port": {
                    "computed": true,
                    "description": "The network output port.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "protocol": {
                    "computed": true,
                    "description": "The network output protocol.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "ttl": {
                    "computed": true,
                    "description": "The network output TTL.",
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
          "nesting_mode": "list"
        },
        "optional": true
      },
      "placement_arn": {
        "description": "The placement Amazon Resource Number (ARN) of the bridge.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_failover_config": {
        "computed": true,
        "description": "The settings for source failover.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "failover_mode": {
              "computed": true,
              "description": "The type of failover you choose for this flow. FAILOVER allows switching between different streams.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "source_priority": {
              "computed": true,
              "description": "The priority you want to assign to a source. You can have a primary stream and a backup stream or two equally prioritized streams.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "primary_source": {
                    "computed": true,
                    "description": "The name of the source you choose as the primary source for this flow.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "state": {
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
      "sources": {
        "description": "The sources on this bridge.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "flow_source": {
              "computed": true,
              "description": "The source of the bridge. A flow source originates in MediaConnect as an existing cloud flow.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "flow_arn": {
                    "computed": true,
                    "description": "The ARN of the cloud flow used as a source of this bridge.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "flow_vpc_interface_attachment": {
                    "computed": true,
                    "description": "The name of the VPC interface attachment to use for this source.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "vpc_interface_name": {
                          "computed": true,
                          "description": "The name of the VPC interface to use for this resource.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "name": {
                    "computed": true,
                    "description": "The name of the flow source.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "network_source": {
              "computed": true,
              "description": "The source of the bridge. A network source originates at your premises.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "multicast_ip": {
                    "computed": true,
                    "description": "The network source multicast IP.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "multicast_source_settings": {
                    "computed": true,
                    "description": "The settings related to the multicast source.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "multicast_source_ip": {
                          "computed": true,
                          "description": "The IP address of the source for source-specific multicast (SSM).",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "name": {
                    "computed": true,
                    "description": "The name of the network source.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "network_name": {
                    "computed": true,
                    "description": "The network source's gateway network name.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "port": {
                    "computed": true,
                    "description": "The network source port.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "protocol": {
                    "computed": true,
                    "description": "The network source protocol.",
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
        "required": true
      }
    },
    "description": "Resource schema for AWS::MediaConnect::Bridge",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMediaconnectBridgeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMediaconnectBridge), &result)
	return &result
}
