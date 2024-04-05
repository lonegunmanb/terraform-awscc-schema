package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCodepipelineCustomActionType = `{
  "block": {
    "attributes": {
      "category": {
        "computed": true,
        "description": "The category of the custom action, such as a build action or a test action.",
        "description_kind": "plain",
        "type": "string"
      },
      "configuration_properties": {
        "computed": true,
        "description": "The configuration properties for the custom action.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "description": {
              "computed": true,
              "description": "The description of the action configuration property that is displayed to users. ",
              "description_kind": "plain",
              "type": "string"
            },
            "key": {
              "computed": true,
              "description": "Whether the configuration property is a key.",
              "description_kind": "plain",
              "type": "bool"
            },
            "name": {
              "computed": true,
              "description": "The name of the action configuration property.",
              "description_kind": "plain",
              "type": "string"
            },
            "queryable": {
              "computed": true,
              "description": "Indicates that the property is used with PollForJobs. When creating a custom action, an action can have up to one queryable property. If it has one, that property must be both required and not secret.If you create a pipeline with a custom action type, and that custom action contains a queryable property, the value for that configuration property is subject to other restrictions. The value must be less than or equal to twenty (20) characters. The value can contain only alphanumeric characters, underscores, and hyphens. ",
              "description_kind": "plain",
              "type": "bool"
            },
            "required": {
              "computed": true,
              "description": "Whether the configuration property is a required value.",
              "description_kind": "plain",
              "type": "bool"
            },
            "secret": {
              "computed": true,
              "description": "Whether the configuration property is secret. Secrets are hidden from all calls except for GetJobDetails, GetThirdPartyJobDetails, PollForJobs, and PollForThirdPartyJobs.",
              "description_kind": "plain",
              "type": "bool"
            },
            "type": {
              "computed": true,
              "description": "The type of the configuration property.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "custom_action_type_id": {
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
      "input_artifact_details": {
        "computed": true,
        "description": "The details of the input artifact for the action, such as its commit ID.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "maximum_count": {
              "computed": true,
              "description": "The maximum number of artifacts allowed for the action type.",
              "description_kind": "plain",
              "type": "number"
            },
            "minimum_count": {
              "computed": true,
              "description": "The minimum number of artifacts allowed for the action type.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "output_artifact_details": {
        "computed": true,
        "description": "The details of the output artifact of the action, such as its commit ID.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "maximum_count": {
              "computed": true,
              "description": "The maximum number of artifacts allowed for the action type.",
              "description_kind": "plain",
              "type": "number"
            },
            "minimum_count": {
              "computed": true,
              "description": "The minimum number of artifacts allowed for the action type.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "provider_name": {
        "computed": true,
        "description": "The provider of the service used in the custom action, such as AWS CodeDeploy.",
        "description_kind": "plain",
        "type": "string"
      },
      "settings": {
        "computed": true,
        "description": "URLs that provide users information about this custom action.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "entity_url_template": {
              "computed": true,
              "description": "The URL returned to the AWS CodePipeline console that provides a deep link to the resources of the external system, such as the configuration page for an AWS CodeDeploy deployment group. This link is provided as part of the action display in the pipeline. ",
              "description_kind": "plain",
              "type": "string"
            },
            "execution_url_template": {
              "computed": true,
              "description": "The URL returned to the AWS CodePipeline console that contains a link to the top-level landing page for the external system, such as the console page for AWS CodeDeploy. This link is shown on the pipeline view page in the AWS CodePipeline console and provides a link to the execution entity of the external action. ",
              "description_kind": "plain",
              "type": "string"
            },
            "revision_url_template": {
              "computed": true,
              "description": "The URL returned to the AWS CodePipeline console that contains a link to the page where customers can update or change the configuration of the external action. ",
              "description_kind": "plain",
              "type": "string"
            },
            "third_party_configuration_url": {
              "computed": true,
              "description": "The URL of a sign-up page where users can sign up for an external service and perform initial configuration of the action provided by that service.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description": "Any tags assigned to the custom action.",
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
      "version": {
        "computed": true,
        "description": "The version identifier of the custom action.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::CodePipeline::CustomActionType",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCodepipelineCustomActionTypeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCodepipelineCustomActionType), &result)
	return &result
}
