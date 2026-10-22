package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNetworkfirewallFirewall = `{
  "block": {
    "attributes": {
      "availability_zone_change_protection": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "availability_zone_mappings": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "availability_zone": {
              "computed": true,
              "description": "A AvailabilityZone",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "delete_protection": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enabled_analysis_types": {
        "computed": true,
        "description": "The types of analysis to enable for the firewall. Can be TLS_SNI, HTTP_HOST, or both.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "endpoint_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "firewall_arn": {
        "computed": true,
        "description": "A resource ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "firewall_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "firewall_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "firewall_policy_arn": {
        "computed": true,
        "description": "A resource ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "firewall_policy_change_protection": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "nat_gateway_mappings": {
        "computed": true,
        "description": "The NAT gateways that the firewall proxies egress traffic through when NoSourcePreservation is true.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "nat_gateway_id": {
              "computed": true,
              "description": "The unique identifier of a NAT gateway.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "no_source_preservation": {
        "computed": true,
        "description": "Indicates whether the firewall operates in proxy mode, where the source IP address of traffic is not preserved. When true, specify NatGatewayMappings and VpcEndpoint instead of the top-level VpcId and SubnetMappings. You can't change this setting after you create the firewall.",
        "description_kind": "plain",
        "type": "bool"
      },
      "number_of_associations": {
        "computed": true,
        "description": "The number of VPC endpoint associations attached to the firewall.",
        "description_kind": "plain",
        "type": "number"
      },
      "proxy_settings": {
        "computed": true,
        "description": "The proxy listener configuration for a proxy-mode (no source preservation) firewall.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "listener_properties": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "port": {
                    "computed": true,
                    "description": "The port the proxy listener accepts traffic on.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "type": {
                    "computed": true,
                    "description": "The protocol the proxy listener uses.",
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
      },
      "subnet_change_protection": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "subnet_mappings": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ip_address_type": {
              "computed": true,
              "description": "A IPAddressType",
              "description_kind": "plain",
              "type": "string"
            },
            "subnet_id": {
              "computed": true,
              "description": "A SubnetId.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
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
      "transit_gateway_attachment_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "transit_gateway_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_endpoint": {
        "computed": true,
        "description": "The VPC and subnets for a proxy-mode firewall's endpoint. Provide this instead of the top-level VpcId and SubnetMappings when NoSourcePreservation is true.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "subnet_mappings": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "ip_address_type": {
                    "computed": true,
                    "description": "A IPAddressType",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "subnet_id": {
                    "computed": true,
                    "description": "A SubnetId.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "set"
              }
            },
            "vpc_id": {
              "computed": true,
              "description": "The unique identifier of the VPC for the firewall endpoint.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "vpc_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::NetworkFirewall::Firewall",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccNetworkfirewallFirewallSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworkfirewallFirewall), &result)
	return &result
}
