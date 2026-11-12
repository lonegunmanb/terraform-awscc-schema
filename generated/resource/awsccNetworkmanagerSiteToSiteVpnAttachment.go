package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNetworkmanagerSiteToSiteVpnAttachment = `{
  "block": {
    "attributes": {
      "attachment_id": {
        "computed": true,
        "description": "The ID of the attachment.",
        "description_kind": "plain",
        "type": "string"
      },
      "attachment_policy_rule_number": {
        "computed": true,
        "description": "The policy rule number associated with the attachment.",
        "description_kind": "plain",
        "type": "number"
      },
      "attachment_type": {
        "computed": true,
        "description": "The type of attachment.",
        "description_kind": "plain",
        "type": "string"
      },
      "core_network_arn": {
        "computed": true,
        "description": "The ARN of a core network for the VPC attachment.",
        "description_kind": "plain",
        "type": "string"
      },
      "core_network_id": {
        "description": "The ID of a core network where you're creating a site-to-site VPN attachment.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "Creation time of the attachment.",
        "description_kind": "plain",
        "type": "string"
      },
      "edge_location": {
        "computed": true,
        "description": "The Region where the edge is located.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_modification_errors": {
        "computed": true,
        "description": "Errors from the last modification of the attachment.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "nat_gateway_rules": {
        "computed": true,
        "description": "The complete set of NAT gateway rules on the attachment. Updating this property replaces the entire rule set, and rules left out of the new value are removed. To remove all rules, specify an empty list or remove this property. NAT gateway support can't be enabled without rules, so it is then disabled unless NatGatewaySupport is specified. Only the core network owner can configure NAT gateway rules. The rules aren't returned when the attachment is read, so drift detection doesn't check them. Not available in AWS GovCloud (US).",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "match_source_cidr": {
              "computed": true,
              "description": "The IPv4 CIDR that source traffic is matched against for translation.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "nat_cidr": {
              "computed": true,
              "description": "The IPv4 CIDR that matched source traffic is translated into. Must have the same prefix length as MatchSourceCidr.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "rule_number": {
              "computed": true,
              "description": "The rule number that uniquely identifies the NAT rule on the attachment.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "nat_gateway_support": {
        "computed": true,
        "description": "Whether NAT gateway support is enabled on the attachment. Required when NatGatewayRules lists rules. Set to DISABLE with NatGatewayRules to keep the rules on the attachment without translating traffic, and set it to ENABLE to activate them again. An attachment that is pending acceptance applies this setting once it is accepted. Only the core network owner can configure or read NAT gateway settings. Not available in AWS GovCloud (US).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "nat_system_route_support": {
        "computed": true,
        "description": "Whether system routes are installed for the NAT CIDRs of the attachment's NAT gateway rules. Configurable on BGP Site-to-Site VPN attachments. Always false and not configurable for static VPN attachments. If omitted, a new NAT gateway configuration uses the service default for the attachment and an existing configuration keeps its current value, so removing this property from a template doesn't reset the setting. Only the core network owner can configure this setting. Not available in AWS GovCloud (US).",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "network_function_group_name": {
        "computed": true,
        "description": "The name of the network function group attachment.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "owner_account_id": {
        "computed": true,
        "description": "Owner account of the attachment.",
        "description_kind": "plain",
        "type": "string"
      },
      "proposed_network_function_group_change": {
        "computed": true,
        "description": "The attachment to move from one network function group to another.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "attachment_policy_rule_number": {
              "computed": true,
              "description": "The rule number in the policy document that applies to this change.",
              "description_kind": "plain",
              "type": "number"
            },
            "network_function_group_name": {
              "computed": true,
              "description": "The name of the network function group to change.",
              "description_kind": "plain",
              "type": "string"
            },
            "tags": {
              "computed": true,
              "description": "The key-value tags that changed for the network function group.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "key": {
                    "computed": true,
                    "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "set"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "proposed_segment_change": {
        "computed": true,
        "description": "The attachment to move from one segment to another.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "attachment_policy_rule_number": {
              "computed": true,
              "description": "The rule number in the policy document that applies to this change.",
              "description_kind": "plain",
              "type": "number"
            },
            "segment_name": {
              "computed": true,
              "description": "The name of the segment to change.",
              "description_kind": "plain",
              "type": "string"
            },
            "tags": {
              "computed": true,
              "description": "The key-value tags that changed for the segment.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "key": {
                    "computed": true,
                    "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "set"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "resource_arn": {
        "computed": true,
        "description": "The ARN of the Resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "routing_policy_label": {
        "computed": true,
        "description": "Routing policy label",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "segment_name": {
        "computed": true,
        "description": "The name of the segment that attachment is in.",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The state of the attachment.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags for the attachment.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "updated_at": {
        "computed": true,
        "description": "Last update time of the attachment.",
        "description_kind": "plain",
        "type": "string"
      },
      "vpn_connection_arn": {
        "description": "The ARN of the site-to-site VPN attachment.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "AWS::NetworkManager::SiteToSiteVpnAttachment Resource Type definition.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccNetworkmanagerSiteToSiteVpnAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworkmanagerSiteToSiteVpnAttachment), &result)
	return &result
}
