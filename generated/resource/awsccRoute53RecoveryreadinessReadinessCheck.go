package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRoute53RecoveryreadinessReadinessCheck = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "readiness_check_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the readiness check.",
        "description_kind": "plain",
        "type": "string"
      },
      "readiness_check_name": {
        "computed": true,
        "description": "Name of the ReadinessCheck to create.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_set_name": {
        "computed": true,
        "description": "The name of the resource set to check.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A collection of tags associated with a resource.",
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
      }
    },
    "description": "Aws Route53 Recovery Readiness Check Schema and API specification.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRoute53RecoveryreadinessReadinessCheckSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRoute53RecoveryreadinessReadinessCheck), &result)
	return &result
}
