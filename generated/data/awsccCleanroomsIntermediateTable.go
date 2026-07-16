package data

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
                                "type": "string"
                              },
                              "allowed_analyses": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "allowed_analysis_providers": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "allowed_result_receivers": {
                                "computed": true,
                                "description_kind": "plain",
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
                                            "type": "string"
                                          }
                                        },
                                        "nesting_mode": "list"
                                      }
                                    }
                                  },
                                  "nesting_mode": "single"
                                }
                              },
                              "disallowed_output_columns": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": [
                                  "list",
                                  "string"
                                ]
                              }
                            },
                            "nesting_mode": "single"
                          }
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            },
            "type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
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
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
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
        "type": "string"
      },
      "membership_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "membership_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "population_analysis_configuration": {
        "computed": true,
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
                    "type": "string"
                  },
                  "query_string": {
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
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::CleanRooms::IntermediateTable",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCleanroomsIntermediateTableSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCleanroomsIntermediateTable), &result)
	return &result
}
