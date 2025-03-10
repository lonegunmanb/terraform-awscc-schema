package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2VpcBlockPublicAccessOptions = `{
  "block": {
    "attributes": {
      "account_id": {
        "computed": true,
        "description": "The identifier for the specified AWS account.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "internet_gateway_block_mode": {
        "computed": true,
        "description": "The desired Block Public Access mode for Internet Gateways in your account. We do not allow to create in a off mode as this is the default value",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::VPCBlockPublicAccessOptions",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2VpcBlockPublicAccessOptionsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2VpcBlockPublicAccessOptions), &result)
	return &result
}
