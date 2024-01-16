package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2VpcEndpointServicePermissions = `{
  "block": {
    "attributes": {
      "allowed_principals": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "service_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::VPCEndpointServicePermissions",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2VpcEndpointServicePermissionsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2VpcEndpointServicePermissions), &result)
	return &result
}
