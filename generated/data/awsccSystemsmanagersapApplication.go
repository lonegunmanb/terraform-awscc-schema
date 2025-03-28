package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSystemsmanagersapApplication = `{
  "block": {
    "attributes": {
      "application_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "application_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The ARN of the SSM-SAP application",
        "description_kind": "plain",
        "type": "string"
      },
      "components_info": {
        "computed": true,
        "description": "This is an optional parameter for component details to which the SAP ABAP application is attached, such as Web Dispatcher.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "component_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "ec_2_instance_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "sid": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "credentials": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "credential_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "database_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "secret_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "database_arn": {
        "computed": true,
        "description": "The ARN of the SAP HANA database",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instances": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "sap_instance_number": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "sid": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags of a SystemsManagerSAP application.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::SystemsManagerSAP::Application",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSystemsmanagersapApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSystemsmanagersapApplication), &result)
	return &result
}
