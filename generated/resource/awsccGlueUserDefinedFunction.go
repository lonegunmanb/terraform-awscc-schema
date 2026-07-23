package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGlueUserDefinedFunction = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the user-defined function.",
        "description_kind": "plain",
        "type": "string"
      },
      "class_name": {
        "computed": true,
        "description": "The Java class that contains the function code.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "database_name": {
        "description": "The name of the catalog database in which the function is located.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "function_name": {
        "description": "The name of the function.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "function_type": {
        "computed": true,
        "description": "The type of the function.",
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
      "owner_name": {
        "computed": true,
        "description": "The owner of the function.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "owner_type": {
        "computed": true,
        "description": "The owner type.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_uris": {
        "computed": true,
        "description": "The resource URIs for the function.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "resource_type": {
              "computed": true,
              "description": "The type of the resource.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "uri": {
              "computed": true,
              "description": "The URI for accessing the resource.",
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
    "description": "Represents a user-defined function (UDF) definition in the AWS Glue Data Catalog.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccGlueUserDefinedFunctionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGlueUserDefinedFunction), &result)
	return &result
}
