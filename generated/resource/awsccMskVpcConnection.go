package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMskVpcConnection = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "authentication": {
        "description": "The type of private link authentication",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "client_subnets": {
        "description_kind": "plain",
        "required": true,
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
      "security_groups": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "A key-value pair to associate with a resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "target_cluster_arn": {
        "description": "The Amazon Resource Name (ARN) of the target cluster",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vpc_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::MSK::VpcConnection",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMskVpcConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMskVpcConnection), &result)
	return &result
}
