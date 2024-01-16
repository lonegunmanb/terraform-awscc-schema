package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccElasticbeanstalkConfigurationTemplate = `{
  "block": {
    "attributes": {
      "application_name": {
        "computed": true,
        "description": "The name of the Elastic Beanstalk application to associate with this configuration template. ",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "An optional description for this configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "environment_id": {
        "computed": true,
        "description": "The ID of an environment whose settings you want to use to create the configuration template. You must specify EnvironmentId if you don't specify PlatformArn, SolutionStackName, or SourceConfiguration. ",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "option_settings": {
        "computed": true,
        "description": "Option values for the Elastic Beanstalk configuration, such as the instance type. If specified, these values override the values obtained from the solution stack or the source configuration template. For a complete list of Elastic Beanstalk configuration options, see [Option Values](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/command-options.html) in the AWS Elastic Beanstalk Developer Guide. ",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "namespace": {
              "computed": true,
              "description": "A unique namespace that identifies the option's associated AWS resource.",
              "description_kind": "plain",
              "type": "string"
            },
            "option_name": {
              "computed": true,
              "description": "The name of the configuration option.",
              "description_kind": "plain",
              "type": "string"
            },
            "resource_name": {
              "computed": true,
              "description": "A unique resource name for the option setting. Use it for a time–based scaling configuration option. ",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The current value for the configuration option.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "platform_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the custom platform. For more information, see [Custom Platforms](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/custom-platforms.html) in the AWS Elastic Beanstalk Developer Guide. ",
        "description_kind": "plain",
        "type": "string"
      },
      "solution_stack_name": {
        "computed": true,
        "description": "The name of an Elastic Beanstalk solution stack (platform version) that this configuration uses. For example, 64bit Amazon Linux 2013.09 running Tomcat 7 Java 7. A solution stack specifies the operating system, runtime, and application server for a configuration template. It also determines the set of configuration options as well as the possible and default values. For more information, see [Supported Platforms](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/concepts.platforms.html) in the AWS Elastic Beanstalk Developer Guide.\n\n You must specify SolutionStackName if you don't specify PlatformArn, EnvironmentId, or SourceConfiguration.\n\n Use the ListAvailableSolutionStacks API to obtain a list of available solution stacks. ",
        "description_kind": "plain",
        "type": "string"
      },
      "source_configuration": {
        "computed": true,
        "description": "An Elastic Beanstalk configuration template to base this one on. If specified, Elastic Beanstalk uses the configuration values from the specified configuration template to create a new configuration.\n\nValues specified in OptionSettings override any values obtained from the SourceConfiguration.\n\nYou must specify SourceConfiguration if you don't specify PlatformArn, EnvironmentId, or SolutionStackName.\n\nConstraint: If both solution stack name and source configuration are specified, the solution stack of the source configuration template must match the specified solution stack name. ",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "application_name": {
              "computed": true,
              "description": "The name of the application associated with the configuration.",
              "description_kind": "plain",
              "type": "string"
            },
            "template_name": {
              "computed": true,
              "description": "The name of the configuration template.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "template_name": {
        "computed": true,
        "description": "The name of the configuration template",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ElasticBeanstalk::ConfigurationTemplate",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccElasticbeanstalkConfigurationTemplateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccElasticbeanstalkConfigurationTemplate), &result)
	return &result
}
