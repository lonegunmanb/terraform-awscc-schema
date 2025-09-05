package resource

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
      "exclusions_allowed": {
        "computed": true,
        "description": "Determines if exclusions are allowed. If you have enabled VPC BPA at the Organization level, exclusions may be not-allowed. Otherwise, they are allowed.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "internet_gateway_block_mode": {
        "description": "The desired Block Public Access mode for Internet Gateways in your account. We do not allow to create in a off mode as this is the default value",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::EC2::VPCBlockPublicAccessOptions",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2VpcBlockPublicAccessOptionsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2VpcBlockPublicAccessOptions), &result)
	return &result
}
