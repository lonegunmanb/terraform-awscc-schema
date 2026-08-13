package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDataexchangeDataSet = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN for the data set.",
        "description_kind": "plain",
        "type": "string"
      },
      "asset_type": {
        "computed": true,
        "description": "The type of asset that is added to a data set.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The date and time that the data set was created, in ISO 8601 format.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_set_id": {
        "computed": true,
        "description": "The unique identifier for the data set.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description for the data set.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the data set.",
        "description_kind": "plain",
        "type": "string"
      },
      "origin": {
        "computed": true,
        "description": "A property that defines the data set as OWNED by the account (for providers) or ENTITLED to the account (for subscribers).",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags for the data set.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "updated_at": {
        "computed": true,
        "description": "The date and time that the data set was last updated, in ISO 8601 format.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::DataExchange::DataSet",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDataexchangeDataSetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDataexchangeDataSet), &result)
	return &result
}
