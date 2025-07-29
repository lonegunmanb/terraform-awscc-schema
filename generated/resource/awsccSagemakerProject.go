package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerProject = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description": "The time at which the project was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "project_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the Project.",
        "description_kind": "plain",
        "type": "string"
      },
      "project_description": {
        "computed": true,
        "description": "The description of the project.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project_id": {
        "computed": true,
        "description": "Project Id.",
        "description_kind": "plain",
        "type": "string"
      },
      "project_name": {
        "description": "The name of the project.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project_status": {
        "computed": true,
        "description": "The status of a project.",
        "description_kind": "plain",
        "type": "string"
      },
      "service_catalog_provisioned_product_details": {
        "computed": true,
        "description": "Provisioned ServiceCatalog  Details",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "provisioned_product_id": {
              "computed": true,
              "description": "The identifier of the provisioning artifact (also known as a version).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "provisioned_product_status_message": {
              "computed": true,
              "description": "Provisioned Product Status Message",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "service_catalog_provisioning_details": {
        "computed": true,
        "description": "Input ServiceCatalog Provisioning Details",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "path_id": {
              "computed": true,
              "description": "The path identifier of the product.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "product_id": {
              "computed": true,
              "description": "Service Catalog product identifier.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "provisioning_artifact_id": {
              "computed": true,
              "description": "The identifier of the provisioning artifact (also known as a version).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "provisioning_parameters": {
              "computed": true,
              "description": "Parameters specified by the administrator that are required for provisioning the product.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "key": {
                    "computed": true,
                    "description": "The parameter key.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description": "The parameter value.",
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
          "nesting_mode": "single"
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
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "template_provider_details": {
        "computed": true,
        "description": "An array of template providers associated with the project.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cfn_template_provider_detail": {
              "computed": true,
              "description": "CloudFormation template provider details for a SageMaker project.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "parameters": {
                    "computed": true,
                    "description": "A list of parameters used in the CloudFormation template.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "key": {
                          "computed": true,
                          "description": "The key of the parameter.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
                          "computed": true,
                          "description": "The value of the parameter.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "role_arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of the IAM role used by the template provider.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "template_name": {
                    "computed": true,
                    "description": "The name of the template used for the project.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "template_url": {
                    "computed": true,
                    "description": "The URL of the CloudFormation template.",
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
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::SageMaker::Project",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSagemakerProjectSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerProject), &result)
	return &result
}
