package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSagemakerNotebookInstance = `{
  "block": {
    "attributes": {
      "accelerator_types": {
        "computed": true,
        "description": "A list of Amazon Elastic Inference (EI) instance types to associate with the notebook instance. Currently, only one instance type can be associated with a notebook instance.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "additional_code_repositories": {
        "computed": true,
        "description": "An array of up to three Git repositories associated with the notebook instance. These can be either the names of Git repositories stored as resources in your account, or the URL of Git repositories in AWS CodeCommit or in any other Git repository. These repositories are cloned at the same level as the default repository of your notebook instance.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "default_code_repository": {
        "computed": true,
        "description": "The Git repository associated with the notebook instance as its default code repository. This can be either the name of a Git repository stored as a resource in your account, or the URL of a Git repository in AWS CodeCommit or in any other Git repository. When you open a notebook instance, it opens in the directory that contains this repository.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "direct_internet_access": {
        "computed": true,
        "description": "Sets whether SageMaker AI provides internet access to the notebook instance. If you set this to Disabled this notebook instance is able to access resources only in your VPC, and is not be able to connect to SageMaker AI training and endpoint services unless you configure a NAT Gateway in your VPC. You can set the value of this parameter to Disabled only if you set a value for the SubnetId parameter.",
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
      "instance_metadata_service_configuration": {
        "computed": true,
        "description": "Information on the IMDS configuration of the notebook instance.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "minimum_instance_metadata_service_version": {
              "computed": true,
              "description": "Indicates the minimum IMDS version that the notebook instance supports. When passed as part of CreateNotebookInstance, if no value is selected, then it defaults to IMDSv1. This means that both IMDSv1 and IMDSv2 are supported. If passed as part of UpdateNotebookInstance, there is no default.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "instance_type": {
        "description": "The type of ML compute instance to launch for the notebook instance. Expect some interruption of service if this parameter is changed as CloudFormation stops a notebook instance and starts it up again to update it.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of a AWS Key Management Service key that SageMaker AI uses to encrypt data on the storage volume attached to your notebook instance. The KMS key you provide must be enabled.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "lifecycle_config_name": {
        "computed": true,
        "description": "The name of a lifecycle configuration to associate with the notebook instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "notebook_instance_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the notebook instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "notebook_instance_name": {
        "computed": true,
        "description": "The name of the new notebook instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "platform_identifier": {
        "computed": true,
        "description": "The platform identifier of the notebook instance runtime environment. The default value is notebook-al2023-v1.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "role_arn": {
        "description": "When you send any requests to AWS resources from the notebook instance, SageMaker AI assumes this role to perform tasks on your behalf. You must grant this role necessary permissions so SageMaker AI can perform these tasks. The policy must allow the SageMaker AI service principal (sagemaker.amazonaws.com) permissions to assume this role. To be able to pass this role to SageMaker AI, the caller of this API must have the iam:PassRole permission.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "root_access": {
        "computed": true,
        "description": "Whether root access is enabled or disabled for users of the notebook instance. The default value is Enabled. Lifecycle configurations need root access to be able to set up a notebook instance. Because of this, lifecycle configurations associated with a notebook instance always run with root access even if you disable root access for users.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_group_ids": {
        "computed": true,
        "description": "The VPC security group IDs, in the form sg-xxxxxxxx. The security groups must be for the same VPC as specified in the subnet.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "subnet_id": {
        "computed": true,
        "description": "The ID of the subnet in a VPC to which you would like to have a connectivity from your ML compute instance.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key. Tag keys must be unique per resource.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "volume_size_in_gb": {
        "computed": true,
        "description": "The size, in GB, of the ML storage volume to attach to the notebook instance. The default value is 5 GB. Expect some interruption of service if this parameter is changed as CloudFormation stops a notebook instance and starts it up again to update it.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "description": "Resource Type definition for AWS::SageMaker::NotebookInstance",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSagemakerNotebookInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSagemakerNotebookInstance), &result)
	return &result
}
