package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccOpensearchserverlessSecurityConfig = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "Security config description",
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
        "computed": true,
        "description": "The friendly name of the security config",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "saml_options": {
        "computed": true,
        "description": "Describes saml options in form of key value map",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "group_attribute": {
              "computed": true,
              "description": "Group attribute for this saml integration",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "metadata": {
              "description": "The XML saml provider metadata document that you want to use",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "session_timeout": {
              "computed": true,
              "description": "Defines the session timeout in minutes",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "user_attribute": {
              "computed": true,
              "description": "Custom attribute for this saml integration",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "security_config_id": {
        "computed": true,
        "description": "The identifier of the security config",
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "computed": true,
        "description": "Config type for security config",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Amazon OpenSearchServerless security config resource",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccOpensearchserverlessSecurityConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOpensearchserverlessSecurityConfig), &result)
	return &result
}
