package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccPcaconnectoradConnector = `{
  "block": {
    "attributes": {
      "certificate_authority_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "connector_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "directory_id": {
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
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "vpc_information": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "security_group_ids": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      }
    },
    "description": "Definition of AWS::PCAConnectorAD::Connector Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccPcaconnectoradConnectorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccPcaconnectoradConnector), &result)
	return &result
}
