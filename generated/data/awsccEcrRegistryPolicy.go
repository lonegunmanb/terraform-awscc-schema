package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEcrRegistryPolicy = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_text": {
        "computed": true,
        "description": "The JSON policy text to apply to your registry. The policy text follows the same format as IAM policy text. For more information, see Registry permissions (https://docs.aws.amazon.com/AmazonECR/latest/userguide/registry-permissions.html) in the Amazon Elastic Container Registry User Guide.",
        "description_kind": "plain",
        "type": "string"
      },
      "registry_id": {
        "computed": true,
        "description": "The registry id.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ECR::RegistryPolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEcrRegistryPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEcrRegistryPolicy), &result)
	return &result
}
