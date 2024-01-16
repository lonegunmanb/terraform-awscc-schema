package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccElasticbeanstalkEnvironment = `{
  "block": {
    "attributes": {
      "application_name": {
        "description": "The name of the application that is associated with this environment.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cname_prefix": {
        "computed": true,
        "description": "If specified, the environment attempts to use this value as the prefix for the CNAME in your Elastic Beanstalk environment URL. If not specified, the CNAME is generated automatically by appending a random alphanumeric string to the environment name.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Your description for this environment.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "endpoint_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "environment_name": {
        "computed": true,
        "description": "A unique name for the environment.",
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
      "operations_role": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of an existing IAM role to be used as the environment's operations role.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "option_settings": {
        "computed": true,
        "description": "Key-value pairs defining configuration options for this environment, such as the instance type.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "namespace": {
              "description": "A unique namespace that identifies the option's associated AWS resource.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "option_name": {
              "description": "The name of the configuration option.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "resource_name": {
              "computed": true,
              "description": "A unique resource name for the option setting. Use it for a time–based scaling configuration option.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The current value for the configuration option.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "platform_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the custom platform to use with the environment.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "solution_stack_name": {
        "computed": true,
        "description": "The name of an Elastic Beanstalk solution stack (platform version) to use with the environment.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Specifies the tags applied to resources in the environment.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the tag.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "template_name": {
        "computed": true,
        "description": "The name of the Elastic Beanstalk configuration template to use with the environment.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tier": {
        "computed": true,
        "description": "Specifies the tier to use in creating this environment. The environment tier that you choose determines whether Elastic Beanstalk provisions resources to support a web application that handles HTTP(S) requests or a web application that handles background-processing tasks.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
              "computed": true,
              "description": "The name of this environment tier.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "computed": true,
              "description": "The type of this environment tier.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "version": {
              "computed": true,
              "description": "The version of this environment tier. When you don't set a value to it, Elastic Beanstalk uses the latest compatible worker tier version.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "version_label": {
        "computed": true,
        "description": "The name of the application version to deploy.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::ElasticBeanstalk::Environment",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccElasticbeanstalkEnvironmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccElasticbeanstalkEnvironment), &result)
	return &result
}
