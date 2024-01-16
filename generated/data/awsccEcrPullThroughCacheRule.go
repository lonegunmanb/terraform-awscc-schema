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
        "description": "The Amazon Resource Name (ARN) of the AWS Secrets Manager secret that identifies the credentials to authenticate to the upstream registry.",
        "description_kind": "plain",
        "type": "string"
      },
      "ecr_repository_prefix": {
        "computed": true,
        "description": "The ECRRepositoryPrefix is a custom alias for upstream registry url.",
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
        "description": "The name of the upstream registry.",
        "description_kind": "plain",
        "type": "string"
      },
      "upstream_registry_url": {
        "computed": true,
        "description": "The upstreamRegistryUrl is the endpoint of upstream registry url of the public repository to be cached",
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
