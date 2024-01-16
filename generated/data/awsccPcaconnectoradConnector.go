package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccPcaconnectoradConnector = `{
  "block": {
    "attributes": {
      "certificate_authority_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "connector_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "directory_id": {
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
        "type": [
          "map",
          "string"
        ]
      },
      "vpc_information": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "security_group_ids": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::PCAConnectorAD::Connector",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccPcaconnectoradConnectorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPcaconnectoradConnector), &result)
	return &result
}
