package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMediatailorChannelPolicy = `{
  "block": {
    "attributes": {
      "channel_name": {
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
      "policy": {
        "description": "\u003cp\u003eThe IAM policy for the channel. IAM policies are used to control access to your channel.\u003c/p\u003e",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Definition of AWS::MediaTailor::ChannelPolicy Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMediatailorChannelPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMediatailorChannelPolicy), &result)
	return &result
}
