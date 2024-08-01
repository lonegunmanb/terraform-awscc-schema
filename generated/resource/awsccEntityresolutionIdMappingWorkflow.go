package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEntityresolutionIdMappingWorkflow = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "The time of this IdMappingWorkflow got created",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the IdMappingWorkflow",
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
      "id_mapping_techniques": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "id_mapping_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "provider_properties": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "intermediate_source_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "intermediate_s3_path": {
                          "description": "The s3 path that would be used to stage the intermediate data being generated during workflow execution.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "provider_configuration": {
                    "computed": true,
                    "description": "Additional Provider configuration that would be required for the provider service. The Configuration must be in JSON string format",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "provider_service_arn": {
                    "description": "Arn of the Provider Service being used.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "rule_based_properties": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "attribute_matching_model": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "record_matching_model": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "rule_definition_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "rules": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "matching_keys": {
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "rule_name": {
                          "description_kind": "plain",
                          "required": true,
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
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "input_source_config": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "input_source_arn": {
              "description": "An Glue table ARN for the input source table, MatchingWorkflow arn or IdNamespace ARN",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "schema_arn": {
              "computed": true,
              "description": "The SchemaMapping arn associated with the Schema",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
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
        "required": true
      },
      "output_source_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kms_arn": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "output_s3_path": {
              "description": "The S3 path to which Entity Resolution will write the output table",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "role_arn": {
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
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "updated_at": {
        "computed": true,
        "description": "The time of this IdMappingWorkflow got last updated at",
        "description_kind": "plain",
        "type": "string"
      },
      "workflow_arn": {
        "computed": true,
        "description": "The default IdMappingWorkflow arn",
        "description_kind": "plain",
        "type": "string"
      },
      "workflow_name": {
        "description": "The name of the IdMappingWorkflow",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "IdMappingWorkflow defined in AWS Entity Resolution service",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEntityresolutionIdMappingWorkflowSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEntityresolutionIdMappingWorkflow), &result)
	return &result
}
