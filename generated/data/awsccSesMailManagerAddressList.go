package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSesMailManagerAddressList = `{
  "block": {
    "attributes": {
      "address_list_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "address_list_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "address_list_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
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
      }
    },
    "description": "Data Source schema for AWS::SES::MailManagerAddressList",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSesMailManagerAddressListSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSesMailManagerAddressList), &result)
	return &result
}
