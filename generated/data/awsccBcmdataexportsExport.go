package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBcmdataexportsExport = `{
  "block": {
    "attributes": {
      "export": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "data_query": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "query_statement": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "table_configurations": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "map",
                      [
                        "map",
                        "string"
                      ]
                    ]
                  }
                },
                "nesting_mode": "single"
              }
            },
            "description": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "destination_configurations": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "s3_destination": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "s3_bucket": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "s3_output_configurations": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "compression": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "format": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "output_type": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "overwrite": {
                                "computed": true,
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          }
                        },
                        "s3_prefix": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "s3_region": {
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
            "export_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "refresh_cadence": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "frequency": {
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
      "export_arn": {
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
    "description": "Data Source schema for AWS::BCMDataExports::Export",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBcmdataexportsExportSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBcmdataexportsExport), &result)
	return &result
}
