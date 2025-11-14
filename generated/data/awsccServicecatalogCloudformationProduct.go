package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccServicecatalogCloudformationProduct = `{
  "block": {
    "attributes": {
      "accept_language": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cloudformation_product_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "distributor": {
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
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "owner": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "product_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "product_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "provisioning_artifact_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "provisioning_artifact_names": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "provisioning_artifact_parameters": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "description": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "disable_template_validation": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "info": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "replace_provisioning_artifacts": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "source_connection": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "connection_parameters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "code_star": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "artifact_path": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "branch": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "connection_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "repository": {
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
            "type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "support_description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "support_email": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "support_url": {
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
    "description": "Data Source schema for AWS::ServiceCatalog::CloudFormationProduct",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccServicecatalogCloudformationProductSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccServicecatalogCloudformationProduct), &result)
	return &result
}
