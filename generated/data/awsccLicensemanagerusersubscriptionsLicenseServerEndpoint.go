package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLicensemanagerusersubscriptionsLicenseServerEndpoint = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description": "The timestamp when License Manager created the license server endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "identity_provider_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) that identifies the registered identity provider the license server endpoint is created for. For Active Directory, that can be a self-managed Active Directory or an AWS Managed Microsoft AD.",
        "description_kind": "plain",
        "type": "string"
      },
      "license_server_endpoint_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the license server endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "license_server_endpoint_id": {
        "computed": true,
        "description": "The ID of the license server endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "license_server_endpoint_provisioning_status": {
        "computed": true,
        "description": "The current state of the provisioning process for the license server endpoint.",
        "description_kind": "plain",
        "type": "string"
      },
      "license_server_settings": {
        "computed": true,
        "description": "The settings of the RDS license server to create the endpoint for.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "server_settings": {
              "computed": true,
              "description": "The settings of a license server, by server type.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "rds_sal_settings": {
                    "computed": true,
                    "description": "The settings of a Remote Desktop Services Subscriber Access License server.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "rds_sal_credentials_provider": {
                          "computed": true,
                          "description": "The source of the credentials that License Manager uses to reach the license server.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "secrets_manager_credentials_provider": {
                                "computed": true,
                                "description": "Credentials held in AWS Secrets Manager.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "secret_id": {
                                      "computed": true,
                                      "description": "The identifier of the Secrets Manager secret that holds the license server administrator credentials. The name of the secret must begin with the prefix license-manager-user-.",
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
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            },
            "server_type": {
              "computed": true,
              "description": "The type of license server.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "license_servers": {
        "computed": true,
        "description": "The license servers that are accessed through this endpoint.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "health_status": {
              "computed": true,
              "description": "The health status of the license server.",
              "description_kind": "plain",
              "type": "string"
            },
            "ipv_4_address": {
              "computed": true,
              "description": "The domain IPv4 address of the license server.",
              "description_kind": "plain",
              "type": "string"
            },
            "ipv_6_address": {
              "computed": true,
              "description": "The domain IPv6 address of the license server.",
              "description_kind": "plain",
              "type": "string"
            },
            "provisioning_status": {
              "computed": true,
              "description": "The current state of the provisioning process for the license server.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "server_endpoint": {
        "computed": true,
        "description": "The network address of a license server endpoint.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "endpoint": {
              "computed": true,
              "description": "The network address of the endpoint.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "server_type": {
        "computed": true,
        "description": "The type of license server.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags that apply for the license server endpoint. A maximum of 50 tags.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key.",
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
    "description": "Data Source schema for AWS::LicenseManagerUserSubscriptions::LicenseServerEndpoint",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLicensemanagerusersubscriptionsLicenseServerEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLicensemanagerusersubscriptionsLicenseServerEndpoint), &result)
	return &result
}
