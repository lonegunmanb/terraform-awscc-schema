package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEmrserverlessApplication = `{
  "block": {
    "attributes": {
      "application_id": {
        "computed": true,
        "description": "The ID of the EMR Serverless Application.",
        "description_kind": "plain",
        "type": "string"
      },
      "architecture": {
        "computed": true,
        "description": "The cpu architecture of an application.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the EMR Serverless Application.",
        "description_kind": "plain",
        "type": "string"
      },
      "auto_start_configuration": {
        "computed": true,
        "description": "Configuration for Auto Start of Application.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enabled": {
              "computed": true,
              "description": "If set to true, the Application will automatically start. Defaults to true.",
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "auto_stop_configuration": {
        "computed": true,
        "description": "Configuration for Auto Stop of Application.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enabled": {
              "computed": true,
              "description": "If set to true, the Application will automatically stop after being idle. Defaults to true.",
              "description_kind": "plain",
              "type": "bool"
            },
            "idle_timeout_minutes": {
              "computed": true,
              "description": "The amount of time [in minutes] to wait before auto stopping the Application when idle. Defaults to 15 minutes.",
              "description_kind": "plain",
              "type": "number"
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
      "image_configuration": {
        "computed": true,
        "description": "The image configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "image_uri": {
              "computed": true,
              "description": "The URI of an image in the Amazon ECR registry. This field is required when you create a new application. If you leave this field blank in an update, Amazon EMR will remove the image configuration.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "initial_capacity": {
        "computed": true,
        "description": "Initial capacity initialized when an Application is started.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "Worker type for an analytics framework.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "worker_configuration": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "cpu": {
                          "computed": true,
                          "description": "Per worker CPU resource. vCPU is the only supported unit and specifying vCPU is optional.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "disk": {
                          "computed": true,
                          "description": "Per worker Disk resource. GB is the only supported unit and specifying GB is optional",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "memory": {
                          "computed": true,
                          "description": "Per worker memory resource. GB is the only supported unit and specifying GB is optional.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "worker_count": {
                    "computed": true,
                    "description": "Initial count of workers to be initialized when an Application is started. This count will be continued to be maintained until the Application is stopped",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "set"
        }
      },
      "maximum_capacity": {
        "computed": true,
        "description": "Maximum allowed cumulative resources for an Application. No new resources will be created once the limit is hit.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cpu": {
              "computed": true,
              "description": "Per worker CPU resource. vCPU is the only supported unit and specifying vCPU is optional.",
              "description_kind": "plain",
              "type": "string"
            },
            "disk": {
              "computed": true,
              "description": "Per worker Disk resource. GB is the only supported unit and specifying GB is optional",
              "description_kind": "plain",
              "type": "string"
            },
            "memory": {
              "computed": true,
              "description": "Per worker memory resource. GB is the only supported unit and specifying GB is optional.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "name": {
        "computed": true,
        "description": "User friendly Application name.",
        "description_kind": "plain",
        "type": "string"
      },
      "network_configuration": {
        "computed": true,
        "description": "Network Configuration for customer VPC connectivity.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "security_group_ids": {
              "computed": true,
              "description": "The ID of the security groups in the VPC to which you want to connect your job or application.",
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            },
            "subnet_ids": {
              "computed": true,
              "description": "The ID of the subnets in the VPC to which you want to connect your job or application.",
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        }
      },
      "release_label": {
        "computed": true,
        "description": "EMR release label.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tag map with key and value",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 1 to 128 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "type": {
        "computed": true,
        "description": "The type of the application",
        "description_kind": "plain",
        "type": "string"
      },
      "worker_type_specifications": {
        "computed": true,
        "description": "The key-value pairs that specify worker type to WorkerTypeSpecificationInput. This parameter must contain all valid worker types for a Spark or Hive application. Valid worker types include Driver and Executor for Spark applications and HiveDriver and TezTask for Hive applications. You can either set image details in this parameter for each worker type, or in imageConfiguration for all worker types.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "image_configuration": {
              "computed": true,
              "description": "The image configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "image_uri": {
                    "computed": true,
                    "description": "The URI of an image in the Amazon ECR registry. This field is required when you create a new application. If you leave this field blank in an update, Amazon EMR will remove the image configuration.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "map"
        }
      }
    },
    "description": "Data Source schema for AWS::EMRServerless::Application",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEmrserverlessApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEmrserverlessApplication), &result)
	return &result
}
