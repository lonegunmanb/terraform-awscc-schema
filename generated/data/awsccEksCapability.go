package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEksCapability = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the capability.",
        "description_kind": "plain",
        "type": "string"
      },
      "capability_name": {
        "computed": true,
        "description": "A unique name for the capability. The name must be unique within your cluster and can contain alphanumeric characters, hyphens, and underscores.",
        "description_kind": "plain",
        "type": "string"
      },
      "cluster_name": {
        "computed": true,
        "description": "The name of the EKS cluster where you want to create the capability.",
        "description_kind": "plain",
        "type": "string"
      },
      "configuration": {
        "computed": true,
        "description": "The configuration settings for the capability. The structure of this object varies depending on the capability type. For Argo CD capabilities, you can configure IAM Identity Center integration, RBAC role mappings, and network access settings.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "argo_cd": {
              "computed": true,
              "description": "Configuration settings for an Argo CD capability. This includes the Kubernetes namespace, IAM Identity Center integration, RBAC role mappings, and network access configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "aws_idc": {
                    "computed": true,
                    "description": "Configuration for integrating Argo CD with IAM Identity Center. This allows you to use your organization's identity provider for authentication to Argo CD.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "idc_instance_arn": {
                          "computed": true,
                          "description": "The ARN of the IAM Identity Center instance to use for authentication.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "idc_managed_application_arn": {
                          "computed": true,
                          "description": "The ARN of the managed application created in IAM Identity Center for this Argo CD capability. This application is automatically created and managed by EKS.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "idc_region": {
                          "computed": true,
                          "description": "The Region where your IAM Identity Center instance is located.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "namespace": {
                    "computed": true,
                    "description": "The Kubernetes namespace where Argo CD resources will be created. If not specified, the default namespace is used.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "network_access": {
                    "computed": true,
                    "description": "Configuration for network access to the Argo CD capability's managed API server endpoint. By default, the Argo CD server is accessible via a public endpoint. You can optionally specify one or more VPC endpoint IDs to enable private connectivity from your VPCs.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "vpce_ids": {
                          "computed": true,
                          "description": "A list of VPC endpoint IDs to associate with the managed Argo CD API server endpoint. Each VPC endpoint provides private connectivity from a specific VPC to the Argo CD server. You can specify multiple VPC endpoint IDs to enable access from multiple VPCs.",
                          "description_kind": "plain",
                          "type": [
                            "list",
                            "string"
                          ]
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "rbac_role_mappings": {
                    "computed": true,
                    "description": "A list of role mappings that define which IAM Identity Center users or groups have which Argo CD roles. Each mapping associates an Argo CD role (ADMIN, EDITOR, or VIEWER) with one or more IAM Identity Center identities.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "identities": {
                          "computed": true,
                          "description": "A list of IAM Identity Center identities (users or groups) that should be assigned this Argo CD role.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "id": {
                                "computed": true,
                                "description": "The unique identifier of the IAM Identity Center user or group.",
                                "description_kind": "plain",
                                "type": "string"
                              },
                              "type": {
                                "computed": true,
                                "description": "The type of identity. Valid values are SSO_USER or SSO_GROUP.",
                                "description_kind": "plain",
                                "type": "string"
                              }
                            },
                            "nesting_mode": "list"
                          }
                        },
                        "role": {
                          "computed": true,
                          "description": "The Argo CD role to assign. Valid values are: ADMIN (full administrative access to Argo CD), EDITOR (edit access to Argo CD resources), or VIEWER (read-only access to Argo CD resources).",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "server_url": {
                    "computed": true,
                    "description": "The URL of the Argo CD server. Use this URL to access the Argo CD web interface and API.",
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
      "created_at": {
        "computed": true,
        "description": "The Unix epoch timestamp in seconds for when the capability was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "delete_propagation_policy": {
        "computed": true,
        "description": "Specifies how Kubernetes resources managed by the capability should be handled when the capability is deleted. Currently, the only supported value is RETAIN which retains all Kubernetes resources managed by the capability when the capability is deleted.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "modified_at": {
        "computed": true,
        "description": "The Unix epoch timestamp in seconds for when the capability was last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the IAM role that the capability uses to interact with AWS services. This role must have a trust policy that allows the EKS service principal to assume it, and it must have the necessary permissions for the capability type you're creating.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The current status of the capability. Valid values include: CREATING (the capability is being created), ACTIVE (the capability is running and available), UPDATING (the capability is being updated), DELETING (the capability is being deleted), CREATE_FAILED (the capability creation failed), UPDATE_FAILED (the capability update failed), or DELETE_FAILED (the capability deletion failed).",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "type": {
        "computed": true,
        "description": "The type of capability to create. Valid values are: ACK (AWS Controllers for Kubernetes, which lets you manage AWS resources directly from Kubernetes), ARGOCD (Argo CD for GitOps-based continuous delivery), or KRO (Kube Resource Orchestrator for composing and managing custom Kubernetes resources).",
        "description_kind": "plain",
        "type": "string"
      },
      "version": {
        "computed": true,
        "description": "The version of the capability software that is currently running.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::EKS::Capability",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEksCapabilitySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEksCapability), &result)
	return &result
}
