package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSsmincidentsReplicationSet = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN of the ReplicationSet.",
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_protected": {
        "computed": true,
        "description": "Configures the ReplicationSet deletion protection.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "regions": {
        "description": "The ReplicationSet configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "region_configuration": {
              "computed": true,
              "description": "The ReplicationSet regional configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "sse_kms_key_id": {
                    "computed": true,
                    "description": "The ARN of the ReplicationSet.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "region_name": {
              "computed": true,
              "description": "The AWS region name.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "required": true
      },
      "tags": {
        "computed": true,
        "description": "The tags to apply to the replication set.",
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
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Resource type definition for AWS::SSMIncidents::ReplicationSet",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSsmincidentsReplicationSetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSsmincidentsReplicationSet), &result)
	return &result
}
