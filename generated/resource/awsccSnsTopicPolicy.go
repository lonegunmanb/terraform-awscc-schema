package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSnsTopicPolicy = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_document": {
        "description": "A policy document that contains permissions to add to the specified SNS topics.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "topic_policy_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "topics": {
        "description": "The Amazon Resource Names (ARN) of the topics to which you want to add the policy. You can use the ` + "`" + `` + "`" + `Ref` + "`" + `` + "`" + ` function to specify an ` + "`" + `` + "`" + `AWS::SNS::Topic` + "`" + `` + "`" + ` resource.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "The ` + "`" + `` + "`" + `AWS::SNS::TopicPolicy` + "`" + `` + "`" + ` resource associates SNS topics with a policy. For an example snippet, see [Declaring an policy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/quickref-iam.html#scenario-sns-policy) in the *User Guide*.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSnsTopicPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSnsTopicPolicy), &result)
	return &result
}
