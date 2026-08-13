package data

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
        "computed": true,
        "description": "The name of the author publishing the app.",
        "description_kind": "plain",
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description": "The date and time this resource was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the application.",
        "description_kind": "plain",
        "type": "string"
      },
      "home_page_url": {
        "computed": true,
        "description": "A URL with more information about the application.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
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
        "type": [
          "list",
          "string"
        ]
      },
      "license_body": {
        "computed": true,
        "description": "A local text file that contains the license of the app.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the application.",
        "description_kind": "plain",
        "type": "string"
      },
      "readme_body": {
        "computed": true,
        "description": "A text readme file in Markdown language that contains a more detailed description of the application.",
        "description_kind": "plain",
        "type": "string"
      },
      "semantic_version": {
        "computed": true,
        "description": "The semantic version of the application.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_code_url": {
        "computed": true,
        "description": "A link to a public repository for the source code of your application.",
        "description_kind": "plain",
        "type": "string"
      },
      "spdx_license_id": {
        "computed": true,
        "description": "A valid identifier from https://spdx.org/licenses/.",
        "description_kind": "plain",
        "type": "string"
      },
      "template_body": {
        "computed": true,
        "description": "The local raw packaged AWS SAM template file of your application.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ServerlessRepo::Application",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccServerlessrepoApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccServerlessrepoApplication), &result)
	return &result
}
