package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSnsTopicPolicy = `{
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
      "topics": {
        "computed": true,
        "description": "The Amazon Resource Names (ARN) of the topics to which you want to add the policy. You can use the [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html)` + "`" + ` function to specify an [AWS::SNS::Topic](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html) resource.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::SNS::TopicPolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSnsTopicPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSnsTopicPolicy), &result)
	return &result
}
