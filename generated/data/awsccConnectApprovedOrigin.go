package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConnectApprovedOrigin = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_id": {
        "computed": true,
        "description": "Amazon Connect instance identifier",
        "description_kind": "plain",
        "type": "string"
      },
      "origin": {
        "computed": true,
        "description": "Domain name to be added to the allowlist of instance",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Connect::ApprovedOrigin",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccConnectApprovedOriginSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectApprovedOrigin), &result)
	return &result
}
