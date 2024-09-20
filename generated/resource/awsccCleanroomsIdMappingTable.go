package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCleanroomsIdMappingTable = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "collaboration_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "collaboration_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "id_mapping_table_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "input_reference_config": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "input_reference_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "manage_resource_policies": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "input_reference_properties": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "id_mapping_table_input_source": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "id_namespace_association_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "kms_key_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "membership_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "membership_identifier": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
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
    "description": "Represents an association between an ID mapping workflow and a collaboration",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCleanroomsIdMappingTableSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCleanroomsIdMappingTable), &result)
	return &result
}
