package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConnectUserHierarchyStructure = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_arn": {
        "computed": true,
        "description": "The identifier of the Amazon Connect instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "user_hierarchy_structure": {
        "computed": true,
        "description": "Information about the hierarchy structure.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "level_five": {
              "computed": true,
              "description": "Information about level five.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "hierarchy_level_arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of the hierarchy level.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "hierarchy_level_id": {
                    "computed": true,
                    "description": "The identifier of the hierarchy level.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "The name of the hierarchy level.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "level_four": {
              "computed": true,
              "description": "Information about level four.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "hierarchy_level_arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of the hierarchy level.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "hierarchy_level_id": {
                    "computed": true,
                    "description": "The identifier of the hierarchy level.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "The name of the hierarchy level.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "level_one": {
              "computed": true,
              "description": "Information about level one.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "hierarchy_level_arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of the hierarchy level.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "hierarchy_level_id": {
                    "computed": true,
                    "description": "The identifier of the hierarchy level.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "The name of the hierarchy level.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "level_three": {
              "computed": true,
              "description": "Information about level three.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "hierarchy_level_arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of the hierarchy level.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "hierarchy_level_id": {
                    "computed": true,
                    "description": "The identifier of the hierarchy level.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "The name of the hierarchy level.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "level_two": {
              "computed": true,
              "description": "Information about level two.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "hierarchy_level_arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of the hierarchy level.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "hierarchy_level_id": {
                    "computed": true,
                    "description": "The identifier of the hierarchy level.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "The name of the hierarchy level.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "user_hierarchy_structure_arn": {
        "computed": true,
        "description": "The identifier of the User Hierarchy Structure.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Connect::UserHierarchyStructure",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccConnectUserHierarchyStructureSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectUserHierarchyStructure), &result)
	return &result
}
