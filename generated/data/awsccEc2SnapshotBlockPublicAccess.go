package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2SnapshotBlockPublicAccess = `{
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
      "state": {
        "computed": true,
        "description": "The state of EBS Snapshot Block Public Access.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EC2::SnapshotBlockPublicAccess",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2SnapshotBlockPublicAccessSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2SnapshotBlockPublicAccess), &result)
	return &result
}
