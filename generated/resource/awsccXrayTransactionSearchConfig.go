package resource

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
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "indexing_percentage": {
        "computed": true,
        "description": "Determines the percentage of traces indexed from CloudWatch Logs to X-Ray",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "description": "This schema provides construct and validation rules for AWS-XRay TransactionSearchConfig resource parameters.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccXrayTransactionSearchConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccXrayTransactionSearchConfig), &result)
	return &result
}
