package resource

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
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "service_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "transit_gateway_attachment_allowed_principals": {
        "computed": true,
        "description": "The AWS account root ARNs (arn:aws:iam::account-id:root) allowed to create transit gateway attachments that connect to this Gateway Load Balancer endpoint service.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "Resource Type definition for AWS::EC2::VPCEndpointServicePermissions",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2VpcEndpointServicePermissionsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2VpcEndpointServicePermissions), &result)
	return &result
}
