package resource

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
        "optional": true,
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
        "optional": true,
        "type": "string"
      },
      "distributor": {
        "computed": true,
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "owner": {
        "description_kind": "plain",
        "required": true,
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
        "optional": true,
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
              "optional": true,
              "type": "string"
            },
            "disable_template_validation": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "info": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "replace_provisioning_artifacts": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
                          "optional": true,
                          "type": "string"
                        },
                        "branch": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "connection_arn": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "repository": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
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
              "optional": true
            },
            "type": {
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
      "support_description": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "support_email": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "support_url": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
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
    "description": "Resource Type definition for AWS::ServiceCatalog::CloudFormationProduct",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccServicecatalogCloudformationProductSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccServicecatalogCloudformationProduct), &result)
	return &result
}
