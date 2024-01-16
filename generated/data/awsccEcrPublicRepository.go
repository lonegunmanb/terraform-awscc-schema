package data

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
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "repository_catalog_data": {
        "computed": true,
        "description": "The CatalogData property type specifies Catalog data for ECR Public Repository. For information about Catalog Data, see \u003clink\u003e",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "about_text": {
              "computed": true,
              "description": "Provide a detailed description of the repository. Identify what is included in the repository, any licensing details, or other relevant information.",
              "description_kind": "plain",
              "type": "string"
            },
            "architectures": {
              "computed": true,
              "description": "Select the system architectures that the images in your repository are compatible with.",
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            },
            "operating_systems": {
              "computed": true,
              "description": "Select the operating systems that the images in your repository are compatible with.",
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            },
            "repository_description": {
              "computed": true,
              "description": "The description of the public repository.",
              "description_kind": "plain",
              "type": "string"
            },
            "usage_text": {
              "computed": true,
              "description": "Provide detailed information about how to use the images in the repository. This provides context, support information, and additional usage details for users of the repository.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "repository_name": {
        "computed": true,
        "description": "The name to use for the repository. The repository name may be specified on its own (such as nginx-web-app) or it can be prepended with a namespace to group the repository into a category (such as project-a/nginx-web-app). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the repository name. For more information, see https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html.",
        "description_kind": "plain",
        "type": "string"
      },
      "repository_policy_text": {
        "computed": true,
        "description": "The JSON repository policy text to apply to the repository. For more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/RepositoryPolicyExamples.html in the Amazon Elastic Container Registry User Guide. ",
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
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::ECR::PublicRepository",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEcrPublicRepositorySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEcrPublicRepository), &result)
	return &result
}
