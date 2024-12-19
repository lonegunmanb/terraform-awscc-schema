package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccQbusinessApplication = `{
  "block": {
    "attributes": {
      "application_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "application_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "attachments_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "attachments_control_mode": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "auto_subscription_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "auto_subscribe": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "default_subscription_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "client_ids_for_oidc": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "encryption_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kms_key_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "iam_identity_provider_arn": {
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
      "identity_center_application_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "identity_center_instance_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "identity_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "personalization_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "personalization_control_mode": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "q_apps_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "q_apps_control_mode": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "quick_sight_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "client_namespace": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "role_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
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
      "updated_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::QBusiness::Application",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccQbusinessApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccQbusinessApplication), &result)
	return &result
}
