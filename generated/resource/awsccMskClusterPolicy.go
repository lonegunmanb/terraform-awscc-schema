package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMskClusterPolicy = `{
  "block": {
    "attributes": {
      "cluster_arn": {
        "description": "The arn of the cluster for the resource policy.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "current_version": {
        "computed": true,
        "description": "The current version of the policy attached to the specified cluster",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy": {
        "description": "A policy document containing permissions to add to the specified cluster.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::MSK::ClusterPolicy",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMskClusterPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMskClusterPolicy), &result)
	return &result
}
