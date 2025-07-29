package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccTransferServer = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "as_2_service_managed_egress_ip_addresses": {
        "computed": true,
        "description": "The list of egress IP addresses of this server. These IP addresses are only relevant for servers that use the AS2 protocol. They are used for sending asynchronous MDNs. These IP addresses are assigned automatically when you create an AS2 server. Additionally, if you update an existing server and add the AS2 protocol, static IP addresses are assigned as well.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "certificate": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "endpoint_details": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "address_allocation_ids": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "security_group_ids": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "subnet_ids": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "vpc_endpoint_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "vpc_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "endpoint_type": {
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
      "identity_provider_details": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "directory_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "function": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "invocation_role": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "sftp_authentication_methods": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "url": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "identity_provider_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ip_address_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "logging_role": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "post_authentication_login_banner": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "pre_authentication_login_banner": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "protocol_details": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "as_2_transports": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "passive_ip": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "set_stat_option": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "tls_session_resumption_mode": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "protocols": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "s3_storage_options": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "directory_listing_optimization": {
              "computed": true,
              "description": "Indicates whether optimization to directory listing on S3 servers is used. Disabled by default for compatibility.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "security_policy_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "server_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "structured_log_destinations": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "tags": {
        "computed": true,
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
      "workflow_details": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "on_partial_upload": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "execution_role": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "workflow_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "on_upload": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "execution_role": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "workflow_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::Transfer::Server",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccTransferServerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccTransferServer), &result)
	return &result
}
