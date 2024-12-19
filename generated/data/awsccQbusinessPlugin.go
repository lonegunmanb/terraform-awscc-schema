package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccQbusinessPlugin = `{
  "block": {
    "attributes": {
      "application_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "auth_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "basic_auth_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "role_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "secret_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "no_auth_configuration": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "o_auth_2_client_credential_configuration": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "authorization_url": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "role_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "secret_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "token_url": {
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
      "build_status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "custom_plugin_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "api_schema": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "payload": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "s3": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bucket": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "key": {
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
            "api_schema_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "description": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "display_name": {
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
      "plugin_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "plugin_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "server_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
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
      "type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "updated_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::QBusiness::Plugin",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccQbusinessPluginSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccQbusinessPlugin), &result)
	return &result
}
