package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEcrPublicRepository = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "repository_catalog_data": {
        "computed": true,
        "description": "The details about the repository that are publicly visible in the Amazon ECR Public Gallery. For more information, see [Amazon ECR Public repository catalog data](https://docs.aws.amazon.com/AmazonECR/latest/public/public-repository-catalog-data.html) in the *Amazon ECR Public User Guide*.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "about_text": {
              "computed": true,
              "description": "Provide a detailed description of the repository. Identify what is included in the repository, any licensing details, or other relevant information.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "architectures": {
              "computed": true,
              "description": "Select the system architectures that the images in your repository are compatible with.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "operating_systems": {
              "computed": true,
              "description": "Select the operating systems that the images in your repository are compatible with.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "repository_description": {
              "computed": true,
              "description": "The description of the public repository.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "usage_text": {
              "computed": true,
              "description": "Provide detailed information about how to use the images in the repository. This provides context, support information, and additional usage details for users of the repository.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "repository_name": {
        "computed": true,
        "description": "The name to use for the public repository. The repository name may be specified on its own (such as ` + "`" + `` + "`" + `nginx-web-app` + "`" + `` + "`" + `) or it can be prepended with a namespace to group the repository into a category (such as ` + "`" + `` + "`" + `project-a/nginx-web-app` + "`" + `` + "`" + `). If you don't specify a name, CFNlong generates a unique physical ID and uses that ID for the repository name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).\n  If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "repository_policy_text": {
        "computed": true,
        "description": "The JSON repository policy text to apply to the public repository. For more information, see [Amazon ECR Public repository policies](https://docs.aws.amazon.com/AmazonECR/latest/public/public-repository-policies.html) in the *Amazon ECR Public User Guide*.",
        "description_kind": "plain",
        "optional": true,
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
              "description": "One part of a key-value pair that make up a tag. A ` + "`" + `` + "`" + `key` + "`" + `` + "`" + ` is a general label that acts like a category for more specific tag values.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "A ` + "`" + `` + "`" + `value` + "`" + `` + "`" + ` acts as a descriptor within a tag category (key).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "The ` + "`" + `` + "`" + `AWS::ECR::PublicRepository` + "`" + `` + "`" + ` resource specifies an Amazon Elastic Container Registry Public (Amazon ECR Public) repository, where users can push and pull Docker images, Open Container Initiative (OCI) images, and OCI compatible artifacts. For more information, see [Amazon ECR public repositories](https://docs.aws.amazon.com/AmazonECR/latest/public/public-repositories.html) in the *Amazon ECR Public User Guide*.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEcrPublicRepositorySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEcrPublicRepository), &result)
	return &result
}
