package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLightsailStaticIp = `{
  "block": {
    "attributes": {
      "attached_to": {
        "computed": true,
        "description": "The instance where the static IP is attached.",
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
      "ip_address": {
        "computed": true,
        "description": "The static IP address.",
        "description_kind": "plain",
        "type": "string"
      },
      "is_attached": {
        "computed": true,
        "description": "A Boolean value indicating whether the static IP is attached.",
        "description_kind": "plain",
        "type": "bool"
      },
      "static_ip_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "static_ip_name": {
        "description": "The name of the static IP address.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Lightsail::StaticIp",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLightsailStaticIpSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLightsailStaticIp), &result)
	return &result
}
