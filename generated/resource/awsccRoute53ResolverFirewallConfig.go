package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53ResolverFirewallConfig = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN of the firewall configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "firewall_config_id": {
        "computed": true,
        "description": "The ID of the firewall configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "firewall_fail_open": {
        "description": "Determines how DNS Firewall operates during failures. By default, fail open is disabled, which means the failure mode is closed. If you enable this option, the failure mode is open. This behavior is only enforced for VPCs that have at least one DNS Firewall rule group association.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "owner_id": {
        "computed": true,
        "description": "The Amazon Web Services account ID of the owner of the VPC that this firewall configuration applies to.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_id": {
        "description": "The ID of the VPC that this firewall configuration applies to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Route53Resolver::FirewallConfig. Configuration of the firewall behavior provided by DNS Firewall for a single VPC from Amazon Virtual Private Cloud (Amazon VPC).",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRoute53ResolverFirewallConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53ResolverFirewallConfig), &result)
	return &result
}
