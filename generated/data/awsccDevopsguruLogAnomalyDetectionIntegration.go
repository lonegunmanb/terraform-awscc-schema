package data

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
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::DevOpsGuru::LogAnomalyDetectionIntegration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDevopsguruLogAnomalyDetectionIntegrationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDevopsguruLogAnomalyDetectionIntegration), &result)
	return &result
}
