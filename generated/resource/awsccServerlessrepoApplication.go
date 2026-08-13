package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccServerlessrepoApplication = `{
  "block": {
    "attributes": {
      "application_id": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the application.",
        "description_kind": "plain",
        "type": "string"
      },
      "author": {
        "description": "The name of the author publishing the app.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description": "The date and time this resource was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "The description of the application.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "home_page_url": {
        "computed": true,
        "description": "A URL with more information about the application.",
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
      "is_verified_author": {
        "computed": true,
        "description": "Whether the author of this application has been verified.",
        "description_kind": "plain",
        "type": "bool"
      },
      "labels": {
        "computed": true,
        "description": "Labels to improve discovery of apps in search results.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "license_body": {
        "computed": true,
        "description": "A local text file that contains the license of the app.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the application.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "readme_body": {
        "computed": true,
        "description": "A text readme file in Markdown language that contains a more detailed description of the application.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "semantic_version": {
        "computed": true,
        "description": "The semantic version of the application.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_code_url": {
        "computed": true,
        "description": "A link to a public repository for the source code of your application.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "spdx_license_id": {
        "computed": true,
        "description": "A valid identifier from https://spdx.org/licenses/.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "template_body": {
        "computed": true,
        "description": "The local raw packaged AWS SAM template file of your application.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource type definition for an AWS Serverless Application Repository application.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccServerlessrepoApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccServerlessrepoApplication), &result)
	return &result
}
