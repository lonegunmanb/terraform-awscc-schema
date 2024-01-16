package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLakeformationTagAssociation = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "lf_tags": {
        "description": "List of Lake Formation Tags to associate with the Lake Formation Resource",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "catalog_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "tag_key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "tag_values": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "resource": {
        "description": "Resource to tag with the Lake Formation Tags",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "catalog": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "database": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "catalog_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "table": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "catalog_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "database_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "table_wildcard": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "table_with_columns": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "catalog_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "column_names": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "database_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
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
    "description": "A resource schema representing a Lake Formation Tag Association. While tag associations are not explicit Lake Formation resources, this CloudFormation resource can be used to associate tags with Lake Formation entities.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLakeformationTagAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLakeformationTagAssociation), &result)
	return &result
}
