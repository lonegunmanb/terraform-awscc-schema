package resource

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
        "optional": true,
        "type": "string"
      },
      "domain": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "security_group_ids": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "subnet_ids": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "vpc_endpoint_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vpc_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "endpoint_type": {
        "computed": true,
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
      "identity_provider_details": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "directory_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "function": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "invocation_role": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sftp_authentication_methods": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "url": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "identity_provider_type": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "logging_role": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "post_authentication_login_banner": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "pre_authentication_login_banner": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "passive_ip": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "set_stat_option": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "tls_session_resumption_mode": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "protocols": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "security_policy_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
        "optional": true,
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
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
                    "optional": true,
                    "type": "string"
                  },
                  "workflow_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "on_upload": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "execution_role": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "workflow_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      }
    },
    "description": "Definition of AWS::Transfer::Server Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccTransferServerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccTransferServer), &result)
	return &result
}
