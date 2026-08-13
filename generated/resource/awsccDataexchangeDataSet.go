package resource

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
        "description": "The type of asset that is added to a data set.",
        "description_kind": "plain",
        "required": true,
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
        "description": "A description for the data set.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the data set.",
        "description_kind": "plain",
        "required": true,
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "updated_at": {
        "computed": true,
        "description": "The date and time that the data set was last updated, in ISO 8601 format.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Definition of AWS::DataExchange::DataSet Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDataexchangeDataSetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDataexchangeDataSet), &result)
	return &result
}
