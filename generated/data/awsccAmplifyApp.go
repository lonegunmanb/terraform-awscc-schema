package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAmplifyApp = `{
  "block": {
    "attributes": {
      "access_token": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "app_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "app_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "auto_branch_creation_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "auto_branch_creation_patterns": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "basic_auth_config": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enable_basic_auth": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "password": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "username": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "build_spec": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "enable_auto_branch_creation": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "enable_auto_build": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "enable_performance_mode": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "enable_pull_request_preview": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "environment_variables": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
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
            "framework": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "pull_request_environment_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "stage": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "basic_auth_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enable_basic_auth": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "password": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "username": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "build_spec": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "custom_headers": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "custom_rules": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "condition": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "source": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "status": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "target": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "default_domain": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enable_branch_auto_deletion": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "environment_variables": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
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
      "iam_service_role": {
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
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "oauth_token": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "platform": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "repository": {
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
      }
    },
    "description": "Data Source schema for AWS::Amplify::App",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAmplifyAppSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAmplifyApp), &result)
	return &result
}
