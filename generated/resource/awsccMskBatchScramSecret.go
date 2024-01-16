package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMskBatchScramSecret = `{
  "block": {
    "attributes": {
      "cluster_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "secret_arn_list": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "Resource Type definition for AWS::MSK::BatchScramSecret",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMskBatchScramSecretSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMskBatchScramSecret), &result)
	return &result
}
