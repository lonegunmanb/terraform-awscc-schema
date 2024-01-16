package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLakeformationTagAssociation = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "lf_tags": {
        "computed": true,
        "description": "List of Lake Formation Tags to associate with the Lake Formation Resource",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "catalog_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "tag_key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "tag_values": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "list"
        }
      },
      "resource": {
        "computed": true,
        "description": "Resource to tag with the Lake Formation Tags",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "catalog": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "database": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "catalog_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "table": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "catalog_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "database_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "table_wildcard": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "table_with_columns": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "catalog_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "column_names": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "database_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
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
      "resource_identifier": {
        "computed": true,
        "description": "Unique string identifying the resource. Used as primary identifier, which ideally should be a string",
        "description_kind": "plain",
        "type": "string"
      },
      "tags_identifier": {
        "computed": true,
        "description": "Unique string identifying the resource's tags. Used as primary identifier, which ideally should be a string",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::LakeFormation::TagAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLakeformationTagAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLakeformationTagAssociation), &result)
	return &result
}
