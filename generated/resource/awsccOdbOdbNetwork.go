package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccOdbOdbNetwork = `{
  "block": {
    "attributes": {
      "availability_zone": {
        "computed": true,
        "description": "The AWS Availability Zone (AZ) where the ODB network is located.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "availability_zone_id": {
        "computed": true,
        "description": "The AZ ID of the AZ where the ODB network is located.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "backup_subnet_cidr": {
        "computed": true,
        "description": "The CIDR range of the backup subnet in the ODB network.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "client_subnet_cidr": {
        "computed": true,
        "description": "The CIDR range of the client subnet in the ODB network.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "custom_domain_name": {
        "computed": true,
        "description": "The domain name to use for the resources in the ODB network.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "default_dns_prefix": {
        "computed": true,
        "description": "The DNS prefix to the default DNS domain name. The default DNS domain name is oraclevcn.com.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "delete_associated_resources": {
        "computed": true,
        "description": "Specifies whether to delete associated OCI networking resources along with the ODB network.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "display_name": {
        "computed": true,
        "description": "The user-friendly name of the ODB network.",
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
      "managed_services": {
        "computed": true,
        "description": "The managed services configuration for the ODB network.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "managed_s3_backup_access": {
              "computed": true,
              "description": "The managed Amazon S3 backup access configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "ipv_4_addresses": {
                    "computed": true,
                    "description": "The IPv4 addresses for the managed Amazon S3 backup access.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "status": {
                    "computed": true,
                    "description": "The status of the managed Amazon S3 backup access.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "managed_services_ipv_4_cidrs": {
              "computed": true,
              "description": "The IPv4 CIDR blocks for the managed services.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "resource_gateway_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the resource gateway.",
              "description_kind": "plain",
              "type": "string"
            },
            "s3_access": {
              "computed": true,
              "description": "The Amazon S3 access configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "domain_name": {
                    "computed": true,
                    "description": "The domain name for the Amazon S3 access.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "ipv_4_addresses": {
                    "computed": true,
                    "description": "The IPv4 addresses for the Amazon S3 access.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "s3_policy_document": {
                    "computed": true,
                    "description": "The endpoint policy for the Amazon S3 access.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "status": {
                    "computed": true,
                    "description": "The status of the Amazon S3 access.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "service_network_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the service network.",
              "description_kind": "plain",
              "type": "string"
            },
            "service_network_endpoint": {
              "computed": true,
              "description": "The service network endpoint configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "vpc_endpoint_id": {
                    "computed": true,
                    "description": "The identifier of the VPC endpoint.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "vpc_endpoint_type": {
                    "computed": true,
                    "description": "The type of the VPC endpoint.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "zero_etl_access": {
              "computed": true,
              "description": "The Zero-ETL access configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cidr": {
                    "computed": true,
                    "description": "The CIDR block for the Zero-ETL access.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "status": {
                    "computed": true,
                    "description": "The status of the Zero-ETL access.",
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
      "oci_network_anchor_id": {
        "computed": true,
        "description": "The unique identifier of the OCI network anchor for the ODB network.",
        "description_kind": "plain",
        "type": "string"
      },
      "oci_resource_anchor_name": {
        "computed": true,
        "description": "The name of the OCI resource anchor that's associated with the ODB network.",
        "description_kind": "plain",
        "type": "string"
      },
      "oci_vcn_url": {
        "computed": true,
        "description": "The URL for the VCN that's associated with the ODB network.",
        "description_kind": "plain",
        "type": "string"
      },
      "odb_network_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the ODB network.",
        "description_kind": "plain",
        "type": "string"
      },
      "odb_network_id": {
        "computed": true,
        "description": "The unique identifier of the ODB network.",
        "description_kind": "plain",
        "type": "string"
      },
      "s3_access": {
        "computed": true,
        "description": "Specifies the configuration for Amazon S3 access from the ODB network.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "s3_policy_document": {
        "computed": true,
        "description": "Specifies the endpoint policy for Amazon S3 access from the ODB network.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags to assign to the Odb Network.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that's 1 to 128 Unicode characters in length and can't be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., :, /, =, +, @, -, and \".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that's 1 to 256 characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "zero_etl_access": {
        "computed": true,
        "description": "Specifies the configuration for Zero-ETL access from the ODB network.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "The AWS::ODB::OdbNetwork resource creates an ODB Network",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccOdbOdbNetworkSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOdbOdbNetwork), &result)
	return &result
}
