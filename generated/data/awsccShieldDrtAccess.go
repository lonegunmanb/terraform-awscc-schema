package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccShieldDrtAccess = `{
  "block": {
    "attributes": {
      "account_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "log_bucket_list": {
        "computed": true,
        "description": "Authorizes the Shield Response Team (SRT) to access the specified Amazon S3 bucket containing log data such as Application Load Balancer access logs, CloudFront logs, or logs from third party sources. You can associate up to 10 Amazon S3 buckets with your subscription.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "role_arn": {
        "computed": true,
        "description": "Authorizes the Shield Response Team (SRT) using the specified role, to access your AWS account to assist with DDoS attack mitigation during potential attacks. This enables the SRT to inspect your AWS WAF configuration and create or update AWS WAF rules and web ACLs.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Shield::DRTAccess",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccShieldDrtAccessSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccShieldDrtAccess), &result)
	return &result
}
