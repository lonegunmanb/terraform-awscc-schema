package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMediatailorChannelPolicy = `{
  "block": {
    "attributes": {
      "channel_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy": {
        "computed": true,
        "description": "\u003cp\u003eThe IAM policy for the channel. IAM policies are used to control access to your channel.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::MediaTailor::ChannelPolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccMediatailorChannelPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMediatailorChannelPolicy), &result)
	return &result
}
