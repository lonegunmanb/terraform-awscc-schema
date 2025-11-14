package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMediaconnectRouterNetworkInterfaces = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the data source.",
        "description_kind": "plain",
        "type": "string"
      },
      "ids": {
        "computed": true,
        "description": "Set of Resource Identifiers.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      }
    },
    "description": "Plural Data Source schema for AWS::MediaConnect::RouterNetworkInterface",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccMediaconnectRouterNetworkInterfacesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMediaconnectRouterNetworkInterfaces), &result)
	return &result
}
