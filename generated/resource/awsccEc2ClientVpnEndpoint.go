package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2ClientVpnEndpoint = `{
  "block": {
    "attributes": {
      "authentication_options": {
        "description": "Information about the authentication method to be used to authenticate clients.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "active_directory": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "directory_id": {
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
            "federated_authentication": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "saml_provider_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "self_service_saml_provider_arn": {
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
            "mutual_authentication": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "client_root_certificate_chain_arn": {
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
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "authorization_policy": {
        "computed": true,
        "description": "The Cedar authorization policy to associate with the Client VPN endpoint.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "description": {
              "computed": true,
              "description": "A brief description of the authorization policy.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "policy_document": {
              "computed": true,
              "description": "The Cedar policy document.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "shadow_mode": {
              "computed": true,
              "description": "Whether the authorization policy runs in shadow (observe-only) mode.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "client_cidr_block": {
        "computed": true,
        "description": "The IPv4 address range, in CIDR notation, from which to assign client IP addresses. The address range cannot overlap with the local CIDR of the VPC in which the associated subnet is located, or the routes that you add manually. The address range cannot be changed after the Client VPN endpoint has been created. Client CIDR range must have a size of at least /22 and must not be greater than /12.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "client_connect_options": {
        "computed": true,
        "description": "The options for managing connection authorization for new client connections.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "lambda_function_arn": {
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
      "client_login_banner_options": {
        "computed": true,
        "description": "Options for enabling a customizable text banner that will be displayed on AWS provided clients when a VPN session is established.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "banner_text": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "client_route_enforcement_options": {
        "computed": true,
        "description": "Client route enforcement is a feature of the Client VPN service that helps enforce administrator defined routes on devices connected through the VPN. This feature helps improve your security posture by ensuring that network traffic originating from a connected client is not inadvertently sent outside the VPN tunnel.\nClient route enforcement works by monitoring the route table of a connected device for routing policy changes to the VPN connection. If the feature detects any VPN routing policy modifications, it will automatically force an update to the route table, reverting it back to the expected route configurations.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enforced": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "client_vpn_endpoint_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "connection_log_options": {
        "description": "Information about the client connection logging options. If you enable client connection logging, data about client connections is sent to a Cloudwatch Logs log stream. The following information is logged: client connection requests, client connection results (successful and unsuccessful), reasons for unsuccessful client connection requests, client connection termination time.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cloudwatch_log_group": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "cloudwatch_log_stream": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "include_authorization_policy_context": {
              "computed": true,
              "description": "Indicates whether authorization-policy evaluation context is included in the connection logs.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "description": {
        "computed": true,
        "description": "A brief description of the Client VPN endpoint.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "device_posture_options": {
        "computed": true,
        "description": "The device-posture trust providers to associate with the Client VPN endpoint.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enabled": {
              "computed": true,
              "description": "Whether device posture is enabled for the Client VPN endpoint.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "trust_providers": {
              "computed": true,
              "description": "The device-posture trust providers.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "public_signing_key_url": {
                    "computed": true,
                    "description": "The URL of the public signing key. Conditionally required by the provider type and enforced server-side (crowdstrike/jumpcloud require it, jamf rejects it).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "tenant_id": {
                    "computed": true,
                    "description": "The tenant identifier registered with the trust provider.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "trust_provider_type": {
                    "computed": true,
                    "description": "The device-posture trust provider type.",
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
          "nesting_mode": "single"
        },
        "optional": true
      },
      "disconnect_on_session_timeout": {
        "computed": true,
        "description": "Indicates whether the client VPN session is disconnected after the maximum sessionTimeoutHours is reached. If true, users are prompted to reconnect client VPN. If false, client VPN attempts to reconnect automatically.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "dns_servers": {
        "computed": true,
        "description": "Information about the DNS servers to be used for DNS resolution. A Client VPN endpoint can have up to two DNS servers. If no DNS server is specified, the DNS address configured on the device is used for the DNS server.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "endpoint_ip_address_type": {
        "computed": true,
        "description": "The IP address type of the Client VPN endpoint.",
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
      "security_group_ids": {
        "computed": true,
        "description": "The IDs of one or more security groups to apply to the target network. You must also specify the ID of the VPC that contains the security groups.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "self_service_portal": {
        "computed": true,
        "description": "Specify whether to enable the self-service portal for the Client VPN endpoint",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "server_certificate_arn": {
        "description": "The ARN of the server certificate. For more information, see the AWS Certificate Manager User Guide.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "session_timeout_hours": {
        "computed": true,
        "description": "The maximum VPN session duration time in hours.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "split_tunnel": {
        "computed": true,
        "description": "Indicates whether split-tunnel is enabled on the AWS Client VPN endpoint.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "tag_specifications": {
        "computed": true,
        "description": "The tags to apply to the Client VPN endpoint during creation.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "resource_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
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
          "nesting_mode": "list"
        },
        "optional": true
      },
      "traffic_ip_address_type": {
        "computed": true,
        "description": "The IP address type of the Client VPN endpoint.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "transit_gateway_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "availability_zone_ids": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "availability_zones": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "transit_gateway_id": {
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
      "transport_protocol": {
        "computed": true,
        "description": "The transport protocol to be used by the VPN session.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpc_id": {
        "computed": true,
        "description": "The ID of the VPC to associate with the Client VPN endpoint. If no security group IDs are specified in the request, the default security group for the VPC is applied.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpn_port": {
        "computed": true,
        "description": "The port number to assign to the Client VPN endpoint for TCP and UDP traffic.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "description": "Resource Type definition for AWS::EC2::ClientVpnEndpoint",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2ClientVpnEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2ClientVpnEndpoint), &result)
	return &result
}
