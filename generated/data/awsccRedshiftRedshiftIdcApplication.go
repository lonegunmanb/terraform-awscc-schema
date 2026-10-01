package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRedshiftRedshiftIdcApplication = `{
  "block": {
    "attributes": {
      "application_type": {
        "computed": true,
        "description": "The type of application being created.",
        "description_kind": "plain",
        "type": "string"
      },
      "authorized_token_issuer_list": {
        "computed": true,
        "description": "The token issuer list for the Amazon Redshift IAM Identity Center application instance.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "authorized_audiences_list": {
              "computed": true,
              "description": "The list of audiences for the authorized token issuer.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "trusted_token_issuer_arn": {
              "computed": true,
              "description": "The ARN for the authorized token issuer for integrating Amazon Redshift with IDC Identity Center.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "iam_role_arn": {
        "computed": true,
        "description": "The IAM role ARN for the Amazon Redshift IAM Identity Center application instance. It has the required permissions to be assumed and invoke the IDC Identity Center API.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "idc_display_name": {
        "computed": true,
        "description": "The display name for the Amazon Redshift IAM Identity Center application instance. It appears in the console.",
        "description_kind": "plain",
        "type": "string"
      },
      "idc_instance_arn": {
        "computed": true,
        "description": "The Amazon resource name (ARN) of the IAM Identity Center instance where Amazon Redshift creates a new managed application.",
        "description_kind": "plain",
        "type": "string"
      },
      "idc_managed_application_arn": {
        "computed": true,
        "description": "The ARN for the Amazon Redshift IAM Identity Center application.",
        "description_kind": "plain",
        "type": "string"
      },
      "idc_onboard_status": {
        "computed": true,
        "description": "The onboarding status for the Amazon Redshift IAM Identity Center application.",
        "description_kind": "plain",
        "type": "string"
      },
      "identity_namespace": {
        "computed": true,
        "description": "The namespace for the Amazon Redshift IAM Identity Center application instance. It determines which managed application verifies the connection token.",
        "description_kind": "plain",
        "type": "string"
      },
      "redshift_idc_application_arn": {
        "computed": true,
        "description": "The ARN for the Redshift application that integrates with IAM Identity Center.",
        "description_kind": "plain",
        "type": "string"
      },
      "redshift_idc_application_name": {
        "computed": true,
        "description": "The name of the Redshift application in IAM Identity Center.",
        "description_kind": "plain",
        "type": "string"
      },
      "service_integrations": {
        "computed": true,
        "description": "A collection of service integrations for the Redshift IAM Identity Center application.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "lake_formation": {
              "computed": true,
              "description": "A list of scopes set up for Lake Formation integration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "lake_formation_query": {
                    "computed": true,
                    "description": "The Lake Formation scope.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "authorization": {
                          "computed": true,
                          "description": "Determines whether the query scope is enabled or disabled.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "list"
              }
            },
            "redshift": {
              "computed": true,
              "description": "A list of scopes set up for Amazon Redshift integration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "connect": {
                    "computed": true,
                    "description": "The Amazon Redshift connect integration scope.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "authorization": {
                          "computed": true,
                          "description": "Determines whether the Amazon Redshift connect integration is enabled or disabled.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "list"
              }
            },
            "s3_access_grants": {
              "computed": true,
              "description": "A list of scopes set up for S3 Access Grants integration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "read_write_access": {
                    "computed": true,
                    "description": "The S3 Access Grants scope.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "authorization": {
                          "computed": true,
                          "description": "Determines whether the read/write scope is enabled or disabled.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "list"
        }
      },
      "sso_tag_keys": {
        "computed": true,
        "description": "A list of tag keys that Redshift Identity Center applications copy to IAM Identity Center.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::Redshift::RedshiftIdcApplication",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRedshiftRedshiftIdcApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRedshiftRedshiftIdcApplication), &result)
	return &result
}
