package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSsmDefaultHostManagementConfiguration = `{
  "block": {
    "attributes": {
      "customer_managed_role_arn": {
        "computed": true,
        "description": "The ARN of the customer-provided IAM role to register for DHMC. Required when PermissionModel is SELF_MANAGED and must not be set when PermissionModel is SERVICE_MANAGED.",
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
      "last_modified_date": {
        "computed": true,
        "description": "The date the DHMC service setting was last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified_user": {
        "computed": true,
        "description": "The principal that last modified the DHMC service setting.",
        "description_kind": "plain",
        "type": "string"
      },
      "permission_model": {
        "description": "The permission model used to enable DHMC. SERVICE_MANAGED lets SSM create and manage the DHMC IAM role on your behalf. SELF_MANAGED registers an IAM role that you provide.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description": "The ARN of the IAM role currently configured for DHMC. For SERVICE_MANAGED this is the SSM-created role in the IAM reserved path; for SELF_MANAGED this echoes the registered CustomerManagedRoleArn.",
        "description_kind": "plain",
        "type": "string"
      },
      "setting_arn": {
        "computed": true,
        "description": "The ARN of the DHMC service setting.",
        "description_kind": "plain",
        "type": "string"
      },
      "setting_id": {
        "computed": true,
        "description": "The ID of the DHMC service setting. This is a fixed value that uniquely identifies the DHMC configuration for the account and Region.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the DHMC service setting. 'Default' indicates DHMC is not configured; 'Customized' indicates a role is configured.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::SSM::DefaultHostManagementConfiguration. Enables or disables Default Host Management Configuration (DHMC) for an AWS account and Region by managing the '/ssm/managed-instance/default-ec2-instance-management-role' service setting. This is a singleton resource - there is at most one DHMC configuration per account/Region.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSsmDefaultHostManagementConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSsmDefaultHostManagementConfiguration), &result)
	return &result
}
