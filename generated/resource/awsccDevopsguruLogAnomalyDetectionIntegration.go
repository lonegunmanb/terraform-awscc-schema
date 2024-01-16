package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDevopsguruLogAnomalyDetectionIntegration = `{
  "block": {
    "attributes": {
      "account_id": {
        "computed": true,
        "description": "User account id, used as the primary identifier for the resource",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "This resource schema represents the LogAnomalyDetectionIntegration resource in the Amazon DevOps Guru.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDevopsguruLogAnomalyDetectionIntegrationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDevopsguruLogAnomalyDetectionIntegration), &result)
	return &result
}
