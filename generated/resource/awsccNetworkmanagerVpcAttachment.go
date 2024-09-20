package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNetworkmanagerVpcAttachment = `{
  "block": {
    "attributes": {
      "attachment_id": {
        "computed": true,
        "description": "Id of the attachment.",
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
        "description": "Attachment type.",
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
        "description": "The ID of a core network for the VPC attachment.",
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
      "network_function_group_name": {
        "computed": true,
        "description": "The name of the network function group attachment.",
        "description_kind": "plain",
        "type": "string"
      },
      "options": {
        "computed": true,
        "description": "Vpc options of the attachment.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "appliance_mode_support": {
              "computed": true,
              "description": "Indicates whether to enable ApplianceModeSupport Support for Vpc Attachment. Valid Values: true | false",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "ipv_6_support": {
              "computed": true,
              "description": "Indicates whether to enable Ipv6 Support for Vpc Attachment. Valid Values: enable | disable",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
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
              "optional": true,
              "type": "number"
            },
            "network_function_group_name": {
              "computed": true,
              "description": "The name of the network function group to change.",
              "description_kind": "plain",
              "optional": true,
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
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
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
              "optional": true,
              "type": "number"
            },
            "segment_name": {
              "computed": true,
              "description": "The name of the segment to change.",
              "description_kind": "plain",
              "optional": true,
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
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "resource_arn": {
        "computed": true,
        "description": "The ARN of the Resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "segment_name": {
        "computed": true,
        "description": "The name of the segment attachment..",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "State of the attachment.",
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_arns": {
        "description": "Subnet Arn list",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
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
      "vpc_arn": {
        "description": "The ARN of the VPC.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "AWS::NetworkManager::VpcAttachment Resoruce Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccNetworkmanagerVpcAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNetworkmanagerVpcAttachment), &result)
	return &result
}
