package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGrafanaWorkspace = `{
  "block": {
    "attributes": {
      "account_access_type": {
        "computed": true,
        "description": "These enums represent valid account access types. Specifically these enums determine whether the workspace can access AWS resources in the AWS account only, or whether it can also access resources in other accounts in the same organization. If the value CURRENT_ACCOUNT is used, a workspace role ARN must be provided. If the value is ORGANIZATION, a list of organizational units must be provided.",
        "description_kind": "plain",
        "type": "string"
      },
      "authentication_providers": {
        "computed": true,
        "description": "List of authentication providers to enable.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "client_token": {
        "computed": true,
        "description": "A unique, case-sensitive, user-provided identifier to ensure the idempotency of the request.",
        "description_kind": "plain",
        "type": "string"
      },
      "creation_timestamp": {
        "computed": true,
        "description": "Timestamp when the workspace was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_sources": {
        "computed": true,
        "description": "List of data sources on the service managed IAM role.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "description": {
        "computed": true,
        "description": "Description of a workspace.",
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint": {
        "computed": true,
        "description": "Endpoint for the Grafana workspace.",
        "description_kind": "plain",
        "type": "string"
      },
      "grafana_version": {
        "computed": true,
        "description": "The version of Grafana to support in your workspace.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "modification_timestamp": {
        "computed": true,
        "description": "Timestamp when the workspace was last modified",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The user friendly name of a workspace.",
        "description_kind": "plain",
        "type": "string"
      },
      "network_access_control": {
        "computed": true,
        "description": "The configuration settings for Network Access Control.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "prefix_list_ids": {
              "computed": true,
              "description": "The list of prefix list IDs. A prefix list is a list of CIDR ranges of IP addresses. The IP addresses specified are allowed to access your workspace. If the list is not included in the configuration then no IP addresses will be allowed to access the workspace.",
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            },
            "vpce_ids": {
              "computed": true,
              "description": "The list of Amazon VPC endpoint IDs for the workspace. If a NetworkAccessConfiguration is specified then only VPC endpoints specified here will be allowed to access the workspace.",
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
      "notification_destinations": {
        "computed": true,
        "description": "List of notification destinations on the customers service managed IAM role that the Grafana workspace can query.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "organization_role_name": {
        "computed": true,
        "description": "The name of an IAM role that already exists to use with AWS Organizations to access AWS data sources and notification channels in other accounts in an organization.",
        "description_kind": "plain",
        "type": "string"
      },
      "organizational_units": {
        "computed": true,
        "description": "List of Organizational Units containing AWS accounts the Grafana workspace can pull data from.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "permission_type": {
        "computed": true,
        "description": "These enums represent valid permission types to use when creating or configuring a Grafana workspace. The SERVICE_MANAGED permission type means the Managed Grafana service will create a workspace IAM role on your behalf. The CUSTOMER_MANAGED permission type means that the customer is expected to provide an IAM role that the Grafana workspace can use to query data sources.",
        "description_kind": "plain",
        "type": "string"
      },
      "plugin_admin_enabled": {
        "computed": true,
        "description": "Allow workspace admins to install plugins",
        "description_kind": "plain",
        "type": "bool"
      },
      "role_arn": {
        "computed": true,
        "description": "IAM Role that will be used to grant the Grafana workspace access to a customers AWS resources.",
        "description_kind": "plain",
        "type": "string"
      },
      "saml_configuration": {
        "computed": true,
        "description": "SAML configuration data associated with an AMG workspace.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "allowed_organizations": {
              "computed": true,
              "description": "List of SAML organizations allowed to access Grafana.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "assertion_attributes": {
              "computed": true,
              "description": "Maps Grafana friendly names to the IdPs SAML attributes.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "email": {
                    "computed": true,
                    "description": "Name of the attribute within the SAML assert to use as the users email in Grafana.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "groups": {
                    "computed": true,
                    "description": "Name of the attribute within the SAML assert to use as the users groups in Grafana.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "login": {
                    "computed": true,
                    "description": "Name of the attribute within the SAML assert to use as the users login handle in Grafana.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "Name of the attribute within the SAML assert to use as the users name in Grafana.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "org": {
                    "computed": true,
                    "description": "Name of the attribute within the SAML assert to use as the users organizations in Grafana.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "role": {
                    "computed": true,
                    "description": "Name of the attribute within the SAML assert to use as the users roles in Grafana.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "idp_metadata": {
              "computed": true,
              "description": "IdP Metadata used to configure SAML authentication in Grafana.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "url": {
                    "computed": true,
                    "description": "URL that vends the IdPs metadata.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "xml": {
                    "computed": true,
                    "description": "XML blob of the IdPs metadata.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "login_validity_duration": {
              "computed": true,
              "description": "The maximum lifetime an authenticated user can be logged in (in minutes) before being required to re-authenticate.",
              "description_kind": "plain",
              "type": "number"
            },
            "role_values": {
              "computed": true,
              "description": "Maps SAML roles to the Grafana Editor and Admin roles.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "admin": {
                    "computed": true,
                    "description": "List of SAML roles which will be mapped into the Grafana Admin role.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "editor": {
                    "computed": true,
                    "description": "List of SAML roles which will be mapped into the Grafana Editor role.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "saml_configuration_status": {
        "computed": true,
        "description": "Valid SAML configuration statuses.",
        "description_kind": "plain",
        "type": "string"
      },
      "sso_client_id": {
        "computed": true,
        "description": "The client ID of the AWS SSO Managed Application.",
        "description_kind": "plain",
        "type": "string"
      },
      "stack_set_name": {
        "computed": true,
        "description": "The name of the AWS CloudFormation stack set to use to generate IAM roles to be used for this workspace.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "These enums represent the status of a workspace.",
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_configuration": {
        "computed": true,
        "description": "The configuration settings for an Amazon VPC that contains data sources for your Grafana workspace to connect to.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "security_group_ids": {
              "computed": true,
              "description": "The list of Amazon EC2 security group IDs attached to the Amazon VPC for your Grafana workspace to connect.",
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            },
            "subnet_ids": {
              "computed": true,
              "description": "The list of Amazon EC2 subnet IDs created in the Amazon VPC for your Grafana workspace to connect.",
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
      "workspace_id": {
        "computed": true,
        "description": "The id that uniquely identifies a Grafana workspace.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Grafana::Workspace",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccGrafanaWorkspaceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGrafanaWorkspace), &result)
	return &result
}
