package data

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
        "type": "string"
      },
      "iam_federation_options": {
        "computed": true,
        "description": "Describe IAM federation options in form of key value map",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "group_attribute": {
              "computed": true,
              "description": "Group attribute for this IAM federation integration",
              "description_kind": "plain",
              "type": "string"
            },
            "user_attribute": {
              "computed": true,
              "description": "User attribute for this IAM federation integration",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "iam_identity_center_options": {
        "computed": true,
        "description": "Describes IAM Identity Center options for an OpenSearch Serverless security configuration in the form of a key-value map",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "application_arn": {
              "computed": true,
              "description": "The ARN of the IAM Identity Center application used to integrate with OpenSearch Serverless",
              "description_kind": "plain",
              "type": "string"
            },
            "application_description": {
              "computed": true,
              "description": "The description of the IAM Identity Center application used to integrate with OpenSearch Serverless",
              "description_kind": "plain",
              "type": "string"
            },
            "application_name": {
              "computed": true,
              "description": "The name of the IAM Identity Center application used to integrate with OpenSearch Serverless",
              "description_kind": "plain",
              "type": "string"
            },
            "group_attribute": {
              "computed": true,
              "description": "Group attribute for this IAM Identity Center integration",
              "description_kind": "plain",
              "type": "string"
            },
            "instance_arn": {
              "computed": true,
              "description": "The ARN of the IAM Identity Center instance used to integrate with OpenSearch Serverless",
              "description_kind": "plain",
              "type": "string"
            },
            "user_attribute": {
              "computed": true,
              "description": "User attribute for this IAM Identity Center integration",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The friendly name of the security config",
        "description_kind": "plain",
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
              "type": "string"
            },
            "metadata": {
              "computed": true,
              "description": "The XML saml provider metadata document that you want to use",
              "description_kind": "plain",
              "type": "string"
            },
            "open_search_serverless_entity_id": {
              "computed": true,
              "description": "Custom entity id attribute to override default entity id for this saml integration",
              "description_kind": "plain",
              "type": "string"
            },
            "session_timeout": {
              "computed": true,
              "description": "Defines the session timeout in minutes",
              "description_kind": "plain",
              "type": "number"
            },
            "user_attribute": {
              "computed": true,
              "description": "Custom attribute for this saml integration",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
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
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::OpenSearchServerless::SecurityConfig",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccOpensearchserverlessSecurityConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOpensearchserverlessSecurityConfig), &result)
	return &result
}
