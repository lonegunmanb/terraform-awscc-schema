package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppstreamAppBlock = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created_time": {
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
      "packaging_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "post_setup_script_details": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "executable_parameters": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "executable_path": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "script_s3_location": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "s3_bucket": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "s3_key": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "timeout_in_seconds": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "setup_script_details": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "executable_parameters": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "executable_path": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "script_s3_location": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "s3_bucket": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "s3_key": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "timeout_in_seconds": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "source_s3_location": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "s3_bucket": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "s3_key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
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
            "tag_key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "tag_value": {
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
    "description": "Data Source schema for AWS::AppStream::AppBlock",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAppstreamAppBlockSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppstreamAppBlock), &result)
	return &result
}
