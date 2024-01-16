package data

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
        "computed": true,
        "description": "The ID of a core network where you're creating a site-to-site VPN attachment.",
        "description_kind": "plain",
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
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "owner_account_id": {
        "computed": true,
        "description": "Owner account of the attachment.",
        "description_kind": "plain",
        "type": "string"
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
      },
      "updated_at": {
        "computed": true,
        "description": "Last update time of the attachment.",
        "description_kind": "plain",
        "type": "string"
      },
      "vpn_connection_arn": {
        "computed": true,
        "description": "The ARN of the site-to-site VPN attachment.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::NetworkManager::SiteToSiteVpnAttachment",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccNetworkmanagerSiteToSiteVpnAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworkmanagerSiteToSiteVpnAttachment), &result)
	return &result
}
