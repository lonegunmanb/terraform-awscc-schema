package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatapipelinePipeline = `{
  "block": {
    "attributes": {
      "activate": {
        "computed": true,
        "description": "Indicates whether to validate and start the pipeline or stop an active pipeline. By default, the value is set to true.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "description": {
        "computed": true,
        "description": "A description of the pipeline.",
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
        "description": "The name of the pipeline.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parameter_objects": {
        "computed": true,
        "description": "The parameter objects used with the pipeline.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "attributes": {
              "computed": true,
              "description": "The attributes of the parameter object.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "key": {
                    "computed": true,
                    "description": "The field identifier.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "string_value": {
                    "computed": true,
                    "description": "The field value, expressed as a String.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "id": {
              "computed": true,
              "description": "The ID of the parameter object.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "parameter_values": {
        "computed": true,
        "description": "The parameter values used with the pipeline.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "id": {
              "computed": true,
              "description": "The ID of the parameter value.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "string_value": {
              "computed": true,
              "description": "The field value, expressed as a String.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "pipeline_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "pipeline_objects": {
        "computed": true,
        "description": "The objects that define the pipeline. These objects overwrite the existing pipeline definition. Not all objects, fields, and values can be updated. For information about restrictions, see Editing Your Pipeline in the AWS Data Pipeline Developer Guide.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "fields": {
              "computed": true,
              "description": "Key-value pairs that define the properties of the object.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "key": {
                    "computed": true,
                    "description": "Specifies the name of a field for a particular object. To view valid values for a particular field, see Pipeline Object Reference in the AWS Data Pipeline Developer Guide.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "ref_value": {
                    "computed": true,
                    "description": "A field value that you specify as an identifier of another object in the same pipeline definition.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "string_value": {
                    "computed": true,
                    "description": "A field value that you specify as a string. To view valid values for a particular field, see Pipeline Object Reference in the AWS Data Pipeline Developer Guide.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "id": {
              "computed": true,
              "description": "The ID of the object.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "The name of the object.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "pipeline_tags": {
        "computed": true,
        "description": "A list of arbitrary tags (key-value pairs) to associate with the pipeline, which you can use to control permissions. For more information, see Controlling Access to Pipelines and Resources in the AWS Data Pipeline Developer Guide.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of a tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value to associate with the key name.",
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
    "description": "An example resource schema demonstrating some basic constructs and validation rules.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDatapipelinePipelineSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatapipelinePipeline), &result)
	return &result
}
