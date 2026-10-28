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
      "cloudwatch_logs_configuration": {
        "computed": true,
        "description": "Configuration for the CloudWatch Logs log group that X-Ray creates when this resource enables Transaction Search. It is used only at that time and is never returned",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "spans_log_group_kms_key_id": {
              "computed": true,
              "description": "The ARN of a symmetric customer managed AWS KMS key in the same account and Region, used to encrypt the aws/spans log group. Key IDs and alias ARNs are not supported, so for an AWS::KMS::Key use Fn::GetAtt Arn, not Ref. X-Ray applies the key only when it creates aws/spans while enabling Transaction Search. If aws/spans already exists, for example from an earlier enablement or an earlier create that failed partway, the create fails and X-Ray does not enable Transaction Search or change the log group. Remove this property to use the existing log group with its current encryption, and use the CloudWatch Logs AssociateKmsKey API to change its key. The caller needs kms:DescribeKey on the key, and the key policy must allow CloudWatch Logs to use it. Adding, changing or removing this property replaces the resource, which fails while Transaction Search is enabled",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
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
