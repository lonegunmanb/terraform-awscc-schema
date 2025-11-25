package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2TransitGatewayMeteringPolicyEntry = `{
  "block": {
    "attributes": {
      "destination_cidr_block": {
        "computed": true,
        "description": "The list of IP addresses of the instances receiving traffic from the transit gateway",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_port_range": {
        "computed": true,
        "description": "The list of ports on destination instances receiving traffic from the transit gateway",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_transit_gateway_attachment_id": {
        "computed": true,
        "description": "The ID of the source attachment through which traffic leaves a transit gateway",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_transit_gateway_attachment_type": {
        "computed": true,
        "description": "The type of the attachment through which traffic leaves a transit gateway",
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
      "metered_account": {
        "description": "The resource owner information responsible for paying default billable charges for the traffic flow",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_rule_number": {
        "description": "The rule number of the metering policy entry",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "protocol": {
        "computed": true,
        "description": "The protocol of the traffic",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_cidr_block": {
        "computed": true,
        "description": "The list of IP addresses of the instances sending traffic to the transit gateway for which the metering policy entry is applicable",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_port_range": {
        "computed": true,
        "description": "The list of ports on source instances sending traffic to the transit gateway",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_transit_gateway_attachment_id": {
        "computed": true,
        "description": "The ID of the source attachment through which traffic enters a transit gateway",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_transit_gateway_attachment_type": {
        "computed": true,
        "description": "The type of the attachment through which traffic enters a  transit gateway",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "State of the transit gateway metering policy",
        "description_kind": "plain",
        "type": "string"
      },
      "transit_gateway_metering_policy_id": {
        "description": "The ID of the transit gateway metering policy for which the entry is being created",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "update_effective_at": {
        "computed": true,
        "description": "The timestamp at which the latest action performed on the metering policy entry will become effective",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "AWS::EC2::TransitGatewayMeteringPolicyEntry Resource Definition",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2TransitGatewayMeteringPolicyEntrySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2TransitGatewayMeteringPolicyEntry), &result)
	return &result
}
