package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccB2BiTransformer = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "edi_type": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "x12_details": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "transaction_set": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "version": {
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
      "file_format": {
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
      "input_conversion": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "format_options": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "x12": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "transaction_set": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "version": {
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
            "from_format": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "mapping": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "template": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "template_language": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "mapping_template": {
        "computed": true,
        "description": "This shape is deprecated: This is a legacy trait. Please use input-conversion or output-conversion.",
        "description_kind": "plain",
        "type": "string"
      },
      "modified_at": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "output_conversion": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "format_options": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "x12": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "transaction_set": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "version": {
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
            "to_format": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "sample_document": {
        "computed": true,
        "description": "This shape is deprecated: This is a legacy trait. Please use input-conversion or output-conversion.",
        "description_kind": "plain",
        "type": "string"
      },
      "sample_documents": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "bucket_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "keys": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "input": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "output": {
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
      "transformer_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "transformer_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::B2BI::Transformer",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccB2BiTransformerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccB2BiTransformer), &result)
	return &result
}
