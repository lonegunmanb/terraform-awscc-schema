package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConnectApprovedOrigin = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_id": {
        "description": "Amazon Connect instance identifier",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "origin": {
        "description": "Domain name to be added to the allowlist of instance",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Connect::ApprovedOrigin",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccConnectApprovedOriginSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectApprovedOrigin), &result)
	return &result
}
