package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccXrayTransactionSearchConfig = `{
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
      },
      "indexing_percentage": {
        "computed": true,
        "description": "Determines the percentage of traces indexed from CloudWatch Logs to X-Ray",
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description": "Data Source schema for AWS::XRay::TransactionSearchConfig",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccXrayTransactionSearchConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccXrayTransactionSearchConfig), &result)
	return &result
}
