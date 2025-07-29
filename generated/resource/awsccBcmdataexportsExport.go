package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBcmdataexportsExport = `{
  "block": {
    "attributes": {
      "export": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "data_query": {
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "query_statement": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "table_configurations": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
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
              },
              "required": true
            },
            "description": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "destination_configurations": {
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "s3_destination": {
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "s3_bucket": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "s3_output_configurations": {
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "compression": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "format": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "output_type": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "overwrite": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "required": true
                        },
                        "s3_prefix": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "s3_region": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            },
            "export_arn": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "refresh_cadence": {
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "frequency": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "export_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
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
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Definition of AWS::BCMDataExports::Export Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBcmdataexportsExportSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBcmdataexportsExport), &result)
	return &result
}
