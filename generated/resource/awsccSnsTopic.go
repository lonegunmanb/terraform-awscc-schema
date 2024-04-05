package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSnsTopic = `{
  "block": {
    "attributes": {
      "archive_policy": {
        "computed": true,
        "description": "The archive policy determines the number of days SNS retains messages. You can set a retention period from 1 to 365 days.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "content_based_deduplication": {
        "computed": true,
        "description": "Enables content-based deduplication for FIFO topics.\n  +  By default, ` + "`" + `` + "`" + `ContentBasedDeduplication` + "`" + `` + "`" + ` is set to ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `. If you create a FIFO topic and this attribute is ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `, you must specify a value for the ` + "`" + `` + "`" + `MessageDeduplicationId` + "`" + `` + "`" + ` parameter for the [Publish](https://docs.aws.amazon.com/sns/latest/api/API_Publish.html) action. \n  +  When you set ` + "`" + `` + "`" + `ContentBasedDeduplication` + "`" + `` + "`" + ` to ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `, SNS uses a SHA-256 hash to generate the ` + "`" + `` + "`" + `MessageDeduplicationId` + "`" + `` + "`" + ` using the body of the message (but not the attributes of the message).\n (Optional) To override the generated value, you can specify a value for the the ` + "`" + `` + "`" + `MessageDeduplicationId` + "`" + `` + "`" + ` parameter for the ` + "`" + `` + "`" + `Publish` + "`" + `` + "`" + ` action.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "data_protection_policy": {
        "computed": true,
        "description": "The body of the policy document you want to use for this topic.\n You can only add one policy per topic.\n The policy must be in JSON string format.\n Length Constraints: Maximum length of 30,720.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "delivery_status_logging": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "failure_feedback_role_arn": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "protocol": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "success_feedback_role_arn": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "success_feedback_sample_rate": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "display_name": {
        "computed": true,
        "description": "The display name to use for an SNS topic with SMS subscriptions. The display name must be maximum 100 characters long, including hyphens (-), underscores (_), spaces, and tabs.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "fifo_topic": {
        "computed": true,
        "description": "Set to true to create a FIFO topic.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "kms_master_key_id": {
        "computed": true,
        "description": "The ID of an AWS managed customer master key (CMK) for SNS or a custom CMK. For more information, see [Key terms](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html#sse-key-terms). For more examples, see ` + "`" + `` + "`" + `KeyId` + "`" + `` + "`" + ` in the *API Reference*.\n This property applies only to [server-side-encryption](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "signature_version": {
        "computed": true,
        "description": "The signature version corresponds to the hashing algorithm used while creating the signature of the notifications, subscription confirmations, or unsubscribe confirmation messages sent by Amazon SNS. By default, ` + "`" + `` + "`" + `SignatureVersion` + "`" + `` + "`" + ` is set to ` + "`" + `` + "`" + `1` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "subscription": {
        "computed": true,
        "description": "The SNS subscriptions (endpoints) for this topic.\n  If you specify the ` + "`" + `` + "`" + `Subscription` + "`" + `` + "`" + ` property in the ` + "`" + `` + "`" + `AWS::SNS::Topic` + "`" + `` + "`" + ` resource and it creates an associated subscription resource, the associated subscription is not deleted when the ` + "`" + `` + "`" + `AWS::SNS::Topic` + "`" + `` + "`" + ` resource is deleted.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "endpoint": {
              "description": "The endpoint that receives notifications from the SNS topic. The endpoint value depends on the protocol that you specify. For more information, see the ` + "`" + `` + "`" + `Endpoint` + "`" + `` + "`" + ` parameter of the ` + "`" + `` + "`" + `Subscribe` + "`" + `` + "`" + ` action in the *API Reference*.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "protocol": {
              "description": "The subscription's protocol. For more information, see the ` + "`" + `` + "`" + `Protocol` + "`" + `` + "`" + ` parameter of the ` + "`" + `` + "`" + `Subscribe` + "`" + `` + "`" + ` action in the *API Reference*.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "tags": {
        "computed": true,
        "description": "The list of tags to add to a new topic.\n  To be able to tag a topic on creation, you must have the ` + "`" + `` + "`" + `sns:CreateTopic` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `sns:TagResource` + "`" + `` + "`" + ` permissions.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The required key portion of the tag.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The optional value portion of the tag.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "topic_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "topic_name": {
        "computed": true,
        "description": "The name of the topic you want to create. Topic names must include only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and must be between 1 and 256 characters long. FIFO topic names must end with ` + "`" + `` + "`" + `.fifo` + "`" + `` + "`" + `.\n If you don't specify a name, CFN generates a unique physical ID and uses that ID for the topic name. For more information, see [Name type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).\n  If you specify a name, you can't perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tracing_config": {
        "computed": true,
        "description": "Tracing mode of an SNS topic. By default ` + "`" + `` + "`" + `TracingConfig` + "`" + `` + "`" + ` is set to ` + "`" + `` + "`" + `PassThrough` + "`" + `` + "`" + `, and the topic passes through the tracing header it receives from an SNS publisher to its subscriptions. If set to ` + "`" + `` + "`" + `Active` + "`" + `` + "`" + `, SNS will vend X-Ray segment data to topic owner account if the sampled flag in the tracing header is true.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "The ` + "`" + `` + "`" + `AWS::SNS::Topic` + "`" + `` + "`" + ` resource creates a topic to which notifications can be published.\n  One account can create a maximum of 100,000 standard topics and 1,000 FIFO topics. For more information, see [endpoints and quotas](https://docs.aws.amazon.com/general/latest/gr/sns.html) in the *General Reference*.\n   The structure of ` + "`" + `` + "`" + `AUTHPARAMS` + "`" + `` + "`" + ` depends on the .signature of the API request. For more information, see [Examples of the complete Signature Version 4 signing process](https://docs.aws.amazon.com/general/latest/gr/sigv4-signed-request-examples.html) in the *General Reference*.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSnsTopicSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSnsTopic), &result)
	return &result
}
