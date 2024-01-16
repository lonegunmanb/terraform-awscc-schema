package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConnectTaskTemplate = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The identifier (arn) of the task template.",
        "description_kind": "plain",
        "type": "string"
      },
      "client_token": {
        "computed": true,
        "description": "the client token string in uuid format",
        "description_kind": "plain",
        "type": "string"
      },
      "constraints": {
        "computed": true,
        "description": "The constraints for the task template",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "invisible_fields": {
              "computed": true,
              "description": "The list of the task template's invisible fields",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "id": {
                    "computed": true,
                    "description": "the identifier (name) for the task template field",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "name": {
                          "computed": true,
                          "description": "The name of the task template field",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "list"
              }
            },
            "read_only_fields": {
              "computed": true,
              "description": "The list of the task template's read only fields",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "id": {
                    "computed": true,
                    "description": "the identifier (name) for the task template field",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "name": {
                          "computed": true,
                          "description": "The name of the task template field",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "list"
              }
            },
            "required_fields": {
              "computed": true,
              "description": "The list of the task template's required fields",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "id": {
                    "computed": true,
                    "description": "the identifier (name) for the task template field",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "name": {
                          "computed": true,
                          "description": "The name of the task template field",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "contact_flow_arn": {
        "computed": true,
        "description": "The identifier of the contact flow.",
        "description_kind": "plain",
        "type": "string"
      },
      "defaults": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "default_value": {
              "computed": true,
              "description": "the default value for the task template's field",
              "description_kind": "plain",
              "type": "string"
            },
            "id": {
              "computed": true,
              "description": "the identifier (name) for the task template field",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
                    "computed": true,
                    "description": "The name of the task template field",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "list"
        }
      },
      "description": {
        "computed": true,
        "description": "The description of the task template.",
        "description_kind": "plain",
        "type": "string"
      },
      "fields": {
        "computed": true,
        "description": "The list of task template's fields",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "description": {
              "computed": true,
              "description": "The description of the task template's field",
              "description_kind": "plain",
              "type": "string"
            },
            "id": {
              "computed": true,
              "description": "the identifier (name) for the task template field",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "name": {
                    "computed": true,
                    "description": "The name of the task template field",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "single_select_options": {
              "computed": true,
              "description": "list of field options to be used with single select",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "type": {
              "computed": true,
              "description": "The type of the task template's field",
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
      "instance_arn": {
        "computed": true,
        "description": "The identifier (arn) of the instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the task template.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the task template",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "One or more tags.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. . You can specify a value that is maximum of 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::Connect::TaskTemplate",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccConnectTaskTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectTaskTemplate), &result)
	return &result
}
