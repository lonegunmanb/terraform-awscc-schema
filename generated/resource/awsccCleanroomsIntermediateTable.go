package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCleanroomsIntermediateTable = `{
  "block": {
    "attributes": {
      "analysis_rules": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "policy": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "v1": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "custom": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "additional_analyses": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "allowed_analyses": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "allowed_analysis_providers": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "allowed_result_receivers": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "differential_privacy": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "columns": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "name": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "list"
                                      },
                                      "optional": true
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              },
                              "disallowed_output_columns": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "collaboration_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "collaboration_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
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
      "intermediate_table_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "membership_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "membership_identifier": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "population_analysis_configuration": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "sql_parameters": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "analysis_template_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "query_string": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Represents an intermediate table that stores cached query results within a collaboration",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCleanroomsIntermediateTableSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCleanroomsIntermediateTable), &result)
	return &result
}
