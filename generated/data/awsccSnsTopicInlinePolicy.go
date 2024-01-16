package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSnsTopicInlinePolicy = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_document": {
        "computed": true,
        "description": "A policy document that contains permissions to add to the specified SNS topics.",
        "description_kind": "plain",
        "type": "string"
      },
      "topic_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the topic to which you want to add the policy.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SNS::TopicInlinePolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSnsTopicInlinePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSnsTopicInlinePolicy), &result)
	return &result
}
