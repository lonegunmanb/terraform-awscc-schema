package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2TransitGatewayAttachmentNatConfiguration = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Synthetic ARN encoding the NAT configuration's identity (the transit gateway attachment ID). Transit Gateway does not issue an ARN for a NAT configuration; this value exists so the configuration has a primary identifier, and it is not accepted by any EC2 API.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "nat_gateway_rules": {
        "description": "The complete set of NAT rules for the attachment. Updating this property replaces the entire rule set: rules omitted from the new value are removed.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "match_source_cidr": {
              "description": "The source CIDR that the rule matches on.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "nat_cidr": {
              "description": "The CIDR that the matched source addresses are translated to.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "rule_number": {
              "description": "The rule number.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "nesting_mode": "set"
        },
        "required": true
      },
      "nat_gateway_support": {
        "description": "Whether NAT gateway support is enabled on the attachment. Set to disable to stage NAT rules without activating them; the rules are retained and become active when this is later set to enable.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "nat_system_route_support": {
        "computed": true,
        "description": "Whether routes for the NAT CIDRs are installed into the transit gateway route tables associated with the attachment. Whether this is configurable depends on the attachment; where the value is forced, supplying the opposite value is rejected. Omitting this property leaves the setting to the service: on a new configuration the service applies the attachment's default, but on an attachment where the value is configurable it keeps the value already in effect rather than reverting to that default. Removing this property from a template, or deleting this resource, therefore leaves the setting at its last value rather than restoring the default, and a resource later created on the same attachment without this property inherits that value.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "transit_gateway_attachment_id": {
        "description": "The ID of the transit gateway attachment to configure NAT on.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::EC2::TransitGatewayAttachmentNatConfiguration. Configures static 1:1 NAT on a transit gateway attachment. Each rule translates a source CIDR to a NAT CIDR. NatGatewaySupport controls whether creating this resource activates the rules or stages them; deleting this resource disables NAT gateway support and removes all NAT rules. This resource must be managed from the account that owns the transit gateway. NAT support depends on the attachment type, and the transit gateway rejects a request to configure NAT on an attachment type that does not support it.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2TransitGatewayAttachmentNatConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2TransitGatewayAttachmentNatConfiguration), &result)
	return &result
}
