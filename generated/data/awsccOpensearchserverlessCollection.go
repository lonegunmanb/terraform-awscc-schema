package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccOpensearchserverlessCollection = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the collection.",
        "description_kind": "plain",
        "type": "string"
      },
      "collection_endpoint": {
        "computed": true,
        "description": "The endpoint for the collection.",
        "description_kind": "plain",
        "type": "string"
      },
      "collection_id": {
        "computed": true,
        "description": "The identifier of the collection",
        "description_kind": "plain",
        "type": "string"
      },
      "dashboard_endpoint": {
        "computed": true,
        "description": "The OpenSearch Dashboards endpoint for the collection.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the collection",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kms_key_arn": {
        "computed": true,
        "description": "The ARN of the AWS KMS key used to encrypt the collection.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the collection.\n\nThe name must meet the following criteria:\nUnique to your account and AWS Region\nStarts with a lowercase letter\nContains only lowercase letters a-z, the numbers 0-9 and the hyphen (-)\nContains between 3 and 32 characters\n",
        "description_kind": "plain",
        "type": "string"
      },
      "standby_replicas": {
        "computed": true,
        "description": "The possible standby replicas for the collection",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "List of tags to be added to the resource",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key in the key-value pair",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value in the key-value pair",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "type": {
        "computed": true,
        "description": "The possible types for the collection",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::OpenSearchServerless::Collection",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccOpensearchserverlessCollectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOpensearchserverlessCollection), &result)
	return &result
}
