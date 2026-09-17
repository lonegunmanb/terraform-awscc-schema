package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerAiWorkloadConfig = `{
  "block": {
    "attributes": {
      "ai_workload_config_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the AI workload configuration. The name segment is restricted to lowercase for the same reason as AIWorkloadConfigName: the engine derives that property from this ARN, so a permissive ARN would yield a derived name the schema itself rejects.",
        "description_kind": "plain",
        "type": "string"
      },
      "ai_workload_config_name": {
        "computed": true,
        "description": "The name of the AI workload configuration. The name must be unique within your AWS account in the current AWS Region. Only lowercase letters and digits are accepted: DeleteAIWorkloadConfig lowercases the name before looking it up, so a name containing an uppercase letter produces a configuration that can be created and read but never deleted.",
        "description_kind": "plain",
        "type": "string"
      },
      "ai_workload_configs": {
        "computed": true,
        "description": "The benchmark tool configuration and workload specification.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "workload_spec": {
              "computed": true,
              "description": "The workload specification that defines benchmark parameters.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "inline": {
                    "computed": true,
                    "description": "An inline YAML or JSON string that defines benchmark parameters. The service validates the document against its own benchmark schema: it must declare a benchmark object whose type member matches the pattern ^(aiperf)$.",
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
      "creation_time": {
        "computed": true,
        "description": "A timestamp that indicates when the AI workload configuration was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "dataset_config": {
        "computed": true,
        "description": "The dataset configuration for the workload. Specify input data channels with their data sources for benchmark workloads.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "input_data_config": {
              "computed": true,
              "description": "An array of input data channel configurations for the workload.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "channel_name": {
                    "computed": true,
                    "description": "The logical name for the data channel.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "data_source": {
                    "computed": true,
                    "description": "The data source for this channel.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "s3_data_source": {
                          "computed": true,
                          "description": "The Amazon S3 data source configuration.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "s3_uri": {
                                "computed": true,
                                "description": "The Amazon S3 URI of the data.",
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
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The metadata that you apply to the AI workload configuration to help you categorize and organize it.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key. Tag keys must be unique per resource.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::SageMaker::AIWorkloadConfig",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSagemakerAiWorkloadConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerAiWorkloadConfig), &result)
	return &result
}
