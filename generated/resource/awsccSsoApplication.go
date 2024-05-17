package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSsoApplication = `{
  "block": {
    "attributes": {
      "application_arn": {
        "computed": true,
        "description": "The Application ARN that is returned upon creation of the Identity Center (SSO) Application",
        "description_kind": "plain",
        "type": "string"
      },
      "application_provider_arn": {
        "description": "The ARN of the application provider under which the operation will run",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description information for the Identity Center (SSO) Application",
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
      "instance_arn": {
        "description": "The ARN of the instance of IAM Identity Center under which the operation will run",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The name you want to assign to this Identity Center (SSO) Application",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "portal_options": {
        "computed": true,
        "description": "A structure that describes the options for the portal associated with an application",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "sign_in_options": {
              "computed": true,
              "description": "A structure that describes the sign-in options for the access portal",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "application_url": {
                    "computed": true,
                    "description": "The URL that accepts authentication requests for an application, this is a required parameter if the Origin parameter is APPLICATION",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "origin": {
                    "description": "This determines how IAM Identity Center navigates the user to the target application",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "visibility": {
              "computed": true,
              "description": "Indicates whether this application is visible in the access portal",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "status": {
        "computed": true,
        "description": "Specifies whether the application is enabled or disabled",
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
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for Identity Center (SSO) Application",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSsoApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSsoApplication), &result)
	return &result
}
