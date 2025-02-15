package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccTransferWebApp = `{
  "block": {
    "attributes": {
      "access_endpoint": {
        "computed": true,
        "description": "The AccessEndpoint is the URL that you provide to your users for them to interact with the Transfer Family web app. You can specify a custom URL or use the default value.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "Specifies the unique Amazon Resource Name (ARN) for the web app.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "identity_provider_details": {
        "computed": true,
        "description": "You can provide a structure that contains the details for the identity provider to use with your web app.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "application_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "instance_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) for the IAM Identity Center used for the web app.",
              "description_kind": "plain",
              "type": "string"
            },
            "role": {
              "computed": true,
              "description": "The IAM role in IAM Identity Center used for the web app.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description": "Key-value pairs that can be used to group and search for web apps.",
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
      "web_app_customization": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "favicon_file": {
              "computed": true,
              "description": "Specifies a favicon to display in the browser tab.",
              "description_kind": "plain",
              "type": "string"
            },
            "logo_file": {
              "computed": true,
              "description": "Specifies a logo to display on the web app.",
              "description_kind": "plain",
              "type": "string"
            },
            "title": {
              "computed": true,
              "description": "Specifies a title to display on the web app.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "web_app_id": {
        "computed": true,
        "description": "A unique identifier for the web app.",
        "description_kind": "plain",
        "type": "string"
      },
      "web_app_units": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "provisioned": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::Transfer::WebApp",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccTransferWebAppSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccTransferWebApp), &result)
	return &result
}
