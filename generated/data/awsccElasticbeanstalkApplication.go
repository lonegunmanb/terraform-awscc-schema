package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccElasticbeanstalkApplication = `{
  "block": {
    "attributes": {
      "application_name": {
        "computed": true,
        "description": "A name for the Elastic Beanstalk application. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the application name.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Your description of the application.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_lifecycle_config": {
        "computed": true,
        "description": "Specifies an application resource lifecycle configuration to prevent your application from accumulating too many versions.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "service_role": {
              "computed": true,
              "description": "The ARN of an IAM service role that Elastic Beanstalk has permission to assume. The ServiceRole property is required the first time that you provide a ResourceLifecycleConfig for the application. After you provide it once, Elastic Beanstalk persists the Service Role with the application, and you don't need to specify it again. You can, however, specify it in subsequent updates to change the Service Role to another value.",
              "description_kind": "plain",
              "type": "string"
            },
            "version_lifecycle_config": {
              "computed": true,
              "description": "Defines lifecycle settings for application versions.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "max_age_rule": {
                    "computed": true,
                    "description": "Specify a max age rule to restrict the length of time that application versions are retained for an application.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "delete_source_from_s3": {
                          "computed": true,
                          "description": "Set to true to delete a version's source bundle from Amazon S3 when Elastic Beanstalk deletes the application version.",
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "enabled": {
                          "computed": true,
                          "description": "Specify true to apply the rule, or false to disable it.",
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "max_age_in_days": {
                          "computed": true,
                          "description": "Specify the number of days to retain an application versions.",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "max_count_rule": {
                    "computed": true,
                    "description": "Specify a max count rule to restrict the number of application versions that are retained for an application.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "delete_source_from_s3": {
                          "computed": true,
                          "description": "Set to true to delete a version's source bundle from Amazon S3 when Elastic Beanstalk deletes the application version.",
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "enabled": {
                          "computed": true,
                          "description": "Specify true to apply the rule, or false to disable it.",
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "max_count": {
                          "computed": true,
                          "description": "Specify the maximum number of application versions to retain.",
                          "description_kind": "plain",
                          "type": "number"
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
      }
    },
    "description": "Data Source schema for AWS::ElasticBeanstalk::Application",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccElasticbeanstalkApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccElasticbeanstalkApplication), &result)
	return &result
}
