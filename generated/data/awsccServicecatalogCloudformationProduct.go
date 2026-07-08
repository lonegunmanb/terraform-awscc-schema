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
        "description": "The language code.",
        "description_kind": "plain",
        "type": "string"
      },
      "cloudformation_product_id": {
        "computed": true,
        "description": "The ID of the product, such as prod-tsjbmal34qvek",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the product.",
        "description_kind": "plain",
        "type": "string"
      },
      "distributor": {
        "computed": true,
        "description": "The distributor of the product.",
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
        "description": "The name of the product.",
        "description_kind": "plain",
        "type": "string"
      },
      "owner": {
        "computed": true,
        "description": "The owner of the product.",
        "description_kind": "plain",
        "type": "string"
      },
      "product_name": {
        "computed": true,
        "description": "The name of the product.",
        "description_kind": "plain",
        "type": "string"
      },
      "product_type": {
        "computed": true,
        "description": "The type of product.",
        "description_kind": "plain",
        "type": "string"
      },
      "provisioning_artifact_ids": {
        "computed": true,
        "description": "The IDs of the provisioning artifacts",
        "description_kind": "plain",
        "type": "string"
      },
      "provisioning_artifact_names": {
        "computed": true,
        "description": "The names of the provisioning artifacts",
        "description_kind": "plain",
        "type": "string"
      },
      "provisioning_artifact_parameters": {
        "computed": true,
        "description": "The configuration of the provisioning artifact (also known as a version).",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "description": {
              "computed": true,
              "description": "The description of the provisioning artifact, including how it differs from the previous provisioning artifact.",
              "description_kind": "plain",
              "type": "string"
            },
            "disable_template_validation": {
              "computed": true,
              "description": "If set to true, AWS Service Catalog stops validating the specified provisioning artifact even if it is invalid.",
              "description_kind": "plain",
              "type": "bool"
            },
            "info": {
              "computed": true,
              "description": "Specify the template source with one of the following options, but not both. Keys accepted: [ LoadTemplateFromURL, ImportFromPhysicalId ] The URL of the AWS CloudFormation template in Amazon S3 in JSON format. Specify the URL in JSON format as follows:\n\n\"LoadTemplateFromURL\": \"https://s3.amazonaws.com/cf-templates-ozkq9d3hgiq2-us-east-1/...\"\n\nImportFromPhysicalId: The physical id of the resource that contains the template. Currently only supports AWS CloudFormation stack arn. Specify the physical id in JSON format as follows: ImportFromPhysicalId: \"arn:aws:cloudformation:[us-east-1]:[accountId]:stack/[StackName]/[resourceId]",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "import_from_physical_id": {
                    "computed": true,
                    "description": "The physical id of the resource that contains the template. Currently only supports AWS CloudFormation stack arn",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "load_template_from_url": {
                    "computed": true,
                    "description": "The URL of the AWS CloudFormation template in Amazon S3 in JSON format.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "name": {
              "computed": true,
              "description": "The name of the provisioning artifact (for example, v1 v2beta). No spaces are allowed.",
              "description_kind": "plain",
              "type": "string"
            },
            "type": {
              "computed": true,
              "description": "The type of provisioning artifact. Valid values are CLOUD_FORMATION_TEMPLATE, TERRAFORM_OPEN_SOURCE, TERRAFORM_CLOUD, EXTERNAL",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "replace_provisioning_artifacts": {
        "computed": true,
        "description": "This property is turned off by default. If turned off, you can update provisioning artifacts or product attributes (such as description, distributor, name, owner, and more) and the associated provisioning artifacts will retain the same unique identifier. Provisioning artifacts are matched within the CloudFormationProduct resource, and only those that have been updated will be changed. Provisioning artifacts are matched by a combinaton of provisioning artifact template URL and name.",
        "description_kind": "plain",
        "type": "bool"
      },
      "source_connection": {
        "computed": true,
        "description": "A top level ProductViewDetail response containing details about the product's connection. AWS Service Catalog returns this field for the CreateProduct, UpdateProduct, DescribeProductAsAdmin, and SearchProductAsAdmin APIs. This response contains the same fields as the ConnectionParameters request, with the addition of the LastSync response.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "connection_parameters": {
              "computed": true,
              "description": "The connection details based on the connection Type.",
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
                          "description": "The absolute path where the artifact resides within the repo and branch, formatted as \"folder/file.json\".",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "branch": {
                          "computed": true,
                          "description": "The specific branch where the artifact resides.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "connection_arn": {
                          "computed": true,
                          "description": "The CodeStar ARN, which is the connection between AWS Service Catalog and the external repository.\n\n",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "repository": {
                          "computed": true,
                          "description": "The specific repository where the product's artifact-to-be-synced resides, formatted as \"Account/Repo.\"",
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
              "description": "The only supported SourceConnection type is Codestar.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "support_description": {
        "computed": true,
        "description": "The support information about the product.",
        "description_kind": "plain",
        "type": "string"
      },
      "support_email": {
        "computed": true,
        "description": "The contact email for product support.",
        "description_kind": "plain",
        "type": "string"
      },
      "support_url": {
        "computed": true,
        "description": "The contact URL for product support.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "One or more tags.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value",
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
