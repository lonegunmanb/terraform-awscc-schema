package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMskTopic = `{
  "block": {
    "attributes": {
      "cluster_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the MSK cluster",
        "description_kind": "plain",
        "type": "string"
      },
      "configs": {
        "computed": true,
        "description": "Base64 encoded configuration properties of the topic",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "partition_count": {
        "computed": true,
        "description": "The number of partitions for the topic",
        "description_kind": "plain",
        "type": "number"
      },
      "replication_factor": {
        "computed": true,
        "description": "The replication factor for the topic",
        "description_kind": "plain",
        "type": "number"
      },
      "topic_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the topic",
        "description_kind": "plain",
        "type": "string"
      },
      "topic_name": {
        "computed": true,
        "description": "The name of the topic",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::MSK::Topic",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccMskTopicSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMskTopic), &result)
	return &result
}
