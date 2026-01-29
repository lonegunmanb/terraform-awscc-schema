package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatazoneProjectProfile = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created_by": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_unit_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_unit_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "environment_configurations": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "aws_account": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "aws_account_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "aws_region": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "region_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "configuration_parameters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "parameter_overrides": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "is_editable": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
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
                  "resolved_parameters": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "is_editable": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "bool"
                        },
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
                  "ssm_path": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "deployment_mode": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "deployment_order": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "description": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "environment_blueprint_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "environment_configuration_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "last_updated_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "project_profile_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "use_default_configurations": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      }
    },
    "description": "Data Source schema for AWS::DataZone::ProjectProfile",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDatazoneProjectProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatazoneProjectProfile), &result)
	return &result
}
