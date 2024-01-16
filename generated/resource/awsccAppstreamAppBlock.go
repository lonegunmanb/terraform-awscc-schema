package resource

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
        "optional": true,
        "type": "string"
      },
      "display_name": {
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "packaging_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
              "optional": true,
              "type": "string"
            },
            "executable_path": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "script_s3_location": {
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "s3_bucket": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "s3_key": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            },
            "timeout_in_seconds": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "setup_script_details": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "executable_parameters": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "executable_path": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "script_s3_location": {
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "s3_bucket": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "s3_key": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            },
            "timeout_in_seconds": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "source_s3_location": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "s3_bucket": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "s3_key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
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
            "tag_key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "tag_value": {
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
    "description": "Resource Type definition for AWS::AppStream::AppBlock",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAppstreamAppBlockSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppstreamAppBlock), &result)
	return &result
}
