package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2VpnConnection = `{
  "block": {
    "attributes": {
      "customer_gateway_id": {
        "description": "The ID of the customer gateway at your end of the VPN connection.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "enable_acceleration": {
        "computed": true,
        "description": "Indicate whether to enable acceleration for the VPN connection.\n Default: ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "local_ipv_4_network_cidr": {
        "computed": true,
        "description": "The IPv4 CIDR on the customer gateway (on-premises) side of the VPN connection.\n Default: ` + "`" + `` + "`" + `0.0.0.0/0` + "`" + `` + "`" + `",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "local_ipv_6_network_cidr": {
        "computed": true,
        "description": "The IPv6 CIDR on the customer gateway (on-premises) side of the VPN connection.\n Default: ` + "`" + `` + "`" + `::/0` + "`" + `` + "`" + `",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "outside_ip_address_type": {
        "computed": true,
        "description": "The type of IPv4 address assigned to the outside interface of the customer gateway device.\n Valid values: ` + "`" + `` + "`" + `PrivateIpv4` + "`" + `` + "`" + ` | ` + "`" + `` + "`" + `PublicIpv4` + "`" + `` + "`" + ` \n Default: ` + "`" + `` + "`" + `PublicIpv4` + "`" + `` + "`" + `",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "remote_ipv_4_network_cidr": {
        "computed": true,
        "description": "The IPv4 CIDR on the AWS side of the VPN connection.\n Default: ` + "`" + `` + "`" + `0.0.0.0/0` + "`" + `` + "`" + `",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "remote_ipv_6_network_cidr": {
        "computed": true,
        "description": "The IPv6 CIDR on the AWS side of the VPN connection.\n Default: ` + "`" + `` + "`" + `::/0` + "`" + `` + "`" + `",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "static_routes_only": {
        "computed": true,
        "description": "Indicates whether the VPN connection uses static routes only. Static routes must be used for devices that don't support BGP.\n If you are creating a VPN connection for a device that does not support Border Gateway Protocol (BGP), you must specify ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "tags": {
        "computed": true,
        "description": "Any tags assigned to the VPN connection.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "transit_gateway_id": {
        "computed": true,
        "description": "The ID of the transit gateway associated with the VPN connection.\n You must specify either ` + "`" + `` + "`" + `TransitGatewayId` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `VpnGatewayId` + "`" + `` + "`" + `, but not both.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "transport_transit_gateway_attachment_id": {
        "computed": true,
        "description": "The transit gateway attachment ID to use for the VPN tunnel.\n Required if ` + "`" + `` + "`" + `OutsideIpAddressType` + "`" + `` + "`" + ` is set to ` + "`" + `` + "`" + `PrivateIpv4` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tunnel_inside_ip_version": {
        "computed": true,
        "description": "Indicate whether the VPN tunnels process IPv4 or IPv6 traffic.\n Default: ` + "`" + `` + "`" + `ipv4` + "`" + `` + "`" + `",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "type": {
        "description": "The type of VPN connection.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vpn_connection_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vpn_gateway_id": {
        "computed": true,
        "description": "The ID of the virtual private gateway at the AWS side of the VPN connection.\n You must specify either ` + "`" + `` + "`" + `TransitGatewayId` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `VpnGatewayId` + "`" + `` + "`" + `, but not both.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpn_tunnel_options_specifications": {
        "computed": true,
        "description": "The tunnel options for the VPN connection.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "dpd_timeout_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "dpd_timeout_seconds": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "enable_tunnel_lifecycle_control": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "ike_versions": {
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
            },
            "log_options": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cloudwatch_log_options": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "log_enabled": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "log_group_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "log_output_format": {
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
            "phase_1_dh_group_numbers": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "value": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "phase_1_encryption_algorithms": {
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
            },
            "phase_1_integrity_algorithms": {
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
            },
            "phase_1_lifetime_seconds": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "phase_2_dh_group_numbers": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "value": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "phase_2_encryption_algorithms": {
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
            },
            "phase_2_integrity_algorithms": {
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
            },
            "phase_2_lifetime_seconds": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "pre_shared_key": {
              "computed": true,
              "description": "The pre-shared key (PSK) to establish initial authentication between the virtual private gateway and customer gateway.\n Constraints: Allowed characters are alphanumeric characters, periods (.), and underscores (_). Must be between 8 and 64 characters in length and cannot start with zero (0).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "rekey_fuzz_percentage": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "rekey_margin_time_seconds": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "replay_window_size": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "startup_action": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "tunnel_inside_cidr": {
              "computed": true,
              "description": "The range of inside IP addresses for the tunnel. Any specified CIDR blocks must be unique across all VPN connections that use the same virtual private gateway. \n Constraints: A size /30 CIDR block from the ` + "`" + `` + "`" + `169.254.0.0/16` + "`" + `` + "`" + ` range. The following CIDR blocks are reserved and cannot be used:\n  +   ` + "`" + `` + "`" + `169.254.0.0/30` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `169.254.1.0/30` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `169.254.2.0/30` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `169.254.3.0/30` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `169.254.4.0/30` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `169.254.5.0/30` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `169.254.169.252/30` + "`" + `` + "`" + `",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "tunnel_inside_ipv_6_cidr": {
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
    "description": "Specifies a VPN connection between a virtual private gateway and a VPN customer gateway or a transit gateway and a VPN customer gateway.\n To specify a VPN connection between a transit gateway and customer gateway, use the ` + "`" + `` + "`" + `TransitGatewayId` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `CustomerGatewayId` + "`" + `` + "`" + ` properties.\n To specify a VPN connection between a virtual private gateway and customer gateway, use the ` + "`" + `` + "`" + `VpnGatewayId` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `CustomerGatewayId` + "`" + `` + "`" + ` properties.\n For more information, see [](https://docs.aws.amazon.com/vpn/latest/s2svpn/VPC_VPN.html) in the *User Guide*.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2VpnConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2VpnConnection), &result)
	return &result
}
