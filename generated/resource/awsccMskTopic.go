package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMskTopic = `{
  "block": {
    "attributes": {
      "cluster_arn": {
        "description": "The Amazon Resource Name (ARN) of the MSK cluster",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "configs": {
        "computed": true,
        "description": "Base64 encoded configuration properties of the topic",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "partition_count": {
        "description": "The number of partitions for the topic",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "replication_factor": {
        "description": "The replication factor for the topic",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "topic_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the topic",
        "description_kind": "plain",
        "type": "string"
      },
      "topic_name": {
        "description": "The name of the topic",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::MSK::Topic",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMskTopicSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMskTopic), &result)
	return &result
}
