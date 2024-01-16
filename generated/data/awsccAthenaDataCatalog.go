package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAthenaDataCatalog = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "A description of the data catalog to be created. ",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the data catalog to create. The catalog name must be unique for the AWS account and can use a maximum of 128 alphanumeric, underscore, at sign, or hyphen characters. ",
        "description_kind": "plain",
        "type": "string"
      },
      "parameters": {
        "computed": true,
        "description": "Specifies the Lambda function or functions to use for creating the data catalog. This is a mapping whose values depend on the catalog type. ",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "A list of comma separated tags to add to the data catalog that is created. ",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "type": {
        "computed": true,
        "description": "The type of data catalog to create: LAMBDA for a federated catalog, GLUE for AWS Glue Catalog, or HIVE for an external hive metastore. ",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Athena::DataCatalog",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAthenaDataCatalogSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAthenaDataCatalog), &result)
	return &result
}
