package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2TransitGatewayMulticastDomain = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description": "The time the transit gateway multicast domain was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "options": {
        "computed": true,
        "description": "The options for the transit gateway multicast domain.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "auto_accept_shared_associations": {
              "computed": true,
              "description": "Indicates whether to automatically cross-account subnet associations that are associated with the transit gateway multicast domain. Valid Values: enable | disable",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "igmpv_2_support": {
              "computed": true,
              "description": "Indicates whether Internet Group Management Protocol (IGMP) version 2 is turned on for the transit gateway multicast domain. Valid Values: enable | disable",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "static_sources_support": {
              "computed": true,
              "description": "Indicates whether support for statically configuring transit gateway multicast group sources is turned on. Valid Values: enable | disable",
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
        "description": "The state of the transit gateway multicast domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags for the transit gateway multicast domain.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key of the tag. Constraints: Tag keys are case-sensitive and accept a maximum of 127 Unicode characters. May not begin with aws:.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the tag. Constraints: Tag values are case-sensitive and accept a maximum of 255 Unicode characters.",
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
        "description": "The ID of the transit gateway.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "transit_gateway_multicast_domain_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the transit gateway multicast domain.",
        "description_kind": "plain",
        "type": "string"
      },
      "transit_gateway_multicast_domain_id": {
        "computed": true,
        "description": "The ID of the transit gateway multicast domain.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "The AWS::EC2::TransitGatewayMulticastDomain type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2TransitGatewayMulticastDomainSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2TransitGatewayMulticastDomain), &result)
	return &result
}
