package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppconfigDeployment = `{
  "block": {
    "attributes": {
      "application_id": {
        "computed": true,
        "description": "The application ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "configuration_profile_id": {
        "computed": true,
        "description": "The configuration profile ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "configuration_version": {
        "computed": true,
        "description": "The configuration version to deploy. If deploying an AWS AppConfig hosted configuration version, you can specify either the version number or version label. For all other configurations, you must specify the version number.",
        "description_kind": "plain",
        "type": "string"
      },
      "deployment_number": {
        "computed": true,
        "description": "The sequence number of the deployment.",
        "description_kind": "plain",
        "type": "string"
      },
      "deployment_strategy_id": {
        "computed": true,
        "description": "The deployment strategy ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the deployment.",
        "description_kind": "plain",
        "type": "string"
      },
      "dynamic_extension_parameters": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "extension_reference": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "parameter_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "parameter_value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "environment_id": {
        "computed": true,
        "description": "The environment ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kms_key_identifier": {
        "computed": true,
        "description": "The AWS Key Management Service key identifier (key ID, key alias, or key ARN) provided when the resource was created or updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key-value string map. The valid character set is [a-zA-Z1-9+-=._:/]. The tag key can be up to 128 characters and must not start with aws:.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value can be up to 256 characters.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::AppConfig::Deployment",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAppconfigDeploymentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppconfigDeployment), &result)
	return &result
}
