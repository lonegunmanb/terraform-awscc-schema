package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccPersonalizeDatasetGroup = `{
  "block": {
    "attributes": {
      "dataset_group_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the dataset group.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain": {
        "computed": true,
        "description": "The domain of a Domain dataset group.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kms_key_arn": {
        "computed": true,
        "description": "The Amazon Resource Name(ARN) of a AWS Key Management Service (KMS) key used to encrypt the datasets.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name for the new dataset group.",
        "description_kind": "plain",
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description": "The ARN of the AWS Identity and Access Management (IAM) role that has permissions to access the AWS Key Management Service (KMS) key. Supplying an IAM role is only valid when also specifying a KMS key.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Personalize::DatasetGroup",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccPersonalizeDatasetGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPersonalizeDatasetGroup), &result)
	return &result
}
