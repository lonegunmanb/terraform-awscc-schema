package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLakeformationDataCellsFilter = `{
  "block": {
    "attributes": {
      "column_names": {
        "computed": true,
        "description": "A list of columns to be included in this Data Cells Filter.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "column_wildcard": {
        "computed": true,
        "description": "An object representing the Data Cells Filter's Columns. Either Column Names or a Wildcard is required",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "excluded_column_names": {
              "computed": true,
              "description": "A list of column names to be excluded from the Data Cells Filter.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "database_name": {
        "description": "The name of the Database that the Table resides in.",
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
      "name": {
        "description": "The desired name of the Data Cells Filter.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "row_filter": {
        "computed": true,
        "description": "An object representing the Data Cells Filter's Row Filter. Either a Filter Expression or a Wildcard is required",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "all_rows_wildcard": {
              "computed": true,
              "description": "An empty object representing a row wildcard.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "filter_expression": {
              "computed": true,
              "description": "A PartiQL predicate.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "table_catalog_id": {
        "description": "The Catalog Id of the Table on which to create a Data Cells Filter.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "table_name": {
        "description": "The name of the Table to create a Data Cells Filter for.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "A resource schema representing a Lake Formation Data Cells Filter.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLakeformationDataCellsFilterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLakeformationDataCellsFilter), &result)
	return &result
}
