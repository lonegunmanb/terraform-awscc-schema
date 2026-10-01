package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConnectUseCase = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_id": {
        "description": "The identifier of the Amazon Connect instance.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "integration_association_id": {
        "description": "The identifier for the integration association.",
        "description_kind": "plain",
        "required": true,
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
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
        "description": "The type of use case to associate to the integration association.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for a use case associated with an Amazon Connect integration association.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccConnectUseCaseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectUseCase), &result)
	return &result
}
