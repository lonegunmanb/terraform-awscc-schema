package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConnectUseCase = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_id": {
        "computed": true,
        "description": "The identifier of the Amazon Connect instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "integration_association_id": {
        "computed": true,
        "description": "The identifier for the integration association.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags used to organize, track, or control access for this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "use_case_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the use case.",
        "description_kind": "plain",
        "type": "string"
      },
      "use_case_id": {
        "computed": true,
        "description": "The identifier of the use case.",
        "description_kind": "plain",
        "type": "string"
      },
      "use_case_type": {
        "computed": true,
        "description": "The type of use case to associate to the integration association.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Connect::UseCase",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccConnectUseCaseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectUseCase), &result)
	return &result
}
