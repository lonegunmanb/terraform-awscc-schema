package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEcrPullThroughCacheRule = `{
  "block": {
    "attributes": {
      "credential_arn": {
        "computed": true,
        "description": "The ARN of the Secrets Manager secret associated with the pull through cache rule.",
        "description_kind": "plain",
        "type": "string"
      },
      "custom_role_arn": {
        "computed": true,
        "description": "The ARN of the IAM role to be assumed by Amazon ECR to authenticate to ECR upstream registry. This role must be in the same account as the registry that you are configuring.",
        "description_kind": "plain",
        "type": "string"
      },
      "ecr_repository_prefix": {
        "computed": true,
        "description": "The Amazon ECR repository prefix associated with the pull through cache rule.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "upstream_registry": {
        "computed": true,
        "description": "The name of the upstream source registry associated with the pull through cache rule.",
        "description_kind": "plain",
        "type": "string"
      },
      "upstream_registry_url": {
        "computed": true,
        "description": "The upstream registry URL associated with the pull through cache rule.",
        "description_kind": "plain",
        "type": "string"
      },
      "upstream_repository_prefix": {
        "computed": true,
        "description": "The repository name prefix of upstream registry to match with the upstream repository name. When this field isn't specified, Amazon ECR will use the ` + "`" + `ROOT` + "`" + `.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ECR::PullThroughCacheRule",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEcrPullThroughCacheRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEcrPullThroughCacheRule), &result)
	return &result
}
