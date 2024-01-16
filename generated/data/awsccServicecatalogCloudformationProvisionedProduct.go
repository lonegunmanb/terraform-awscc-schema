package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccServicecatalogCloudformationProvisionedProduct = `{
  "block": {
    "attributes": {
      "accept_language": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cloudformation_stack_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "notification_arns": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "outputs": {
        "computed": true,
        "description": "List of key-value pair outputs.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "path_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "path_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "product_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "product_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "provisioned_product_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "provisioned_product_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "provisioning_artifact_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "provisioning_artifact_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "provisioning_parameters": {
        "computed": true,
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
      "provisioning_preferences": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "stack_set_accounts": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "stack_set_failure_tolerance_count": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "stack_set_failure_tolerance_percentage": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "stack_set_max_concurrency_count": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "stack_set_max_concurrency_percentage": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "stack_set_operation_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "stack_set_regions": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        }
      },
      "record_id": {
        "computed": true,
        "description_kind": "plain",
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
      }
    },
    "description": "Data Source schema for AWS::ServiceCatalog::CloudFormationProvisionedProduct",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccServicecatalogCloudformationProvisionedProductSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccServicecatalogCloudformationProvisionedProduct), &result)
	return &result
}
