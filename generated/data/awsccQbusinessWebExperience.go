package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccQbusinessWebExperience = `{
  "block": {
    "attributes": {
      "application_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "customization_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "custom_css_url": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "favicon_url": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "font_url": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "logo_url": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "default_endpoint": {
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
      "identity_provider_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "open_id_connect_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "secrets_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "secrets_role": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "saml_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "authentication_url": {
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
      "origins": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "role_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "sample_prompts_control_mode": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "subtitle": {
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
      },
      "title": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "updated_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "web_experience_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "web_experience_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "welcome_message": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::QBusiness::WebExperience",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccQbusinessWebExperienceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccQbusinessWebExperience), &result)
	return &result
}
