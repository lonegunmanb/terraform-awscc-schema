package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53CidrCollection = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon resource name (ARN) to uniquely identify the AWS resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "locations": {
        "computed": true,
        "description": "A complex type that contains information about the list of CIDR locations.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cidr_list": {
              "computed": true,
              "description": "A list of CIDR blocks.",
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            },
            "location_name": {
              "computed": true,
              "description": "The name of the location that is associated with the CIDR collection.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "name": {
        "computed": true,
        "description": "A unique name for the CIDR collection.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Route53::CidrCollection",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRoute53CidrCollectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53CidrCollection), &result)
	return &result
}
