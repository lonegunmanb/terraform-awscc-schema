package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRedshiftClusterParameterGroup = `{
  "block": {
    "attributes": {
      "description": {
        "description": "A description of the parameter group.",
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
      "parameter_group_family": {
        "description": "The Amazon Redshift engine version to which the cluster parameter group applies. The cluster engine version determines the set of parameters.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parameter_group_name": {
        "computed": true,
        "description": "The name of the cluster parameter group.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "parameters": {
        "computed": true,
        "description": "An array of parameters to be modified. A maximum of 20 parameters can be modified in a single request.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "parameter_name": {
              "description": "The name of the parameter.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "parameter_value": {
              "description": "The value of the parameter. If ` + "`" + `ParameterName` + "`" + ` is ` + "`" + `wlm_json_configuration` + "`" + `, then the maximum size of ` + "`" + `ParameterValue` + "`" + ` is 8000 characters.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::Redshift::ClusterParameterGroup",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRedshiftClusterParameterGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRedshiftClusterParameterGroup), &result)
	return &result
}
