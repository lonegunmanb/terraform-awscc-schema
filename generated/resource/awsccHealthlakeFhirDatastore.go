package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccHealthlakeFhirDatastore = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "The time that a Data Store was created.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "nanos": {
              "computed": true,
              "description": "Nanoseconds.",
              "description_kind": "plain",
              "type": "number"
            },
            "seconds": {
              "computed": true,
              "description": "Seconds since epoch.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "datastore_arn": {
        "computed": true,
        "description": "The Amazon Resource Name used in the creation of the Data Store.",
        "description_kind": "plain",
        "type": "string"
      },
      "datastore_endpoint": {
        "computed": true,
        "description": "The AWS endpoint for the Data Store. Each Data Store will have it's own endpoint with Data Store ID in the endpoint URL.",
        "description_kind": "plain",
        "type": "string"
      },
      "datastore_id": {
        "computed": true,
        "description": "The AWS-generated ID number for the Data Store.",
        "description_kind": "plain",
        "type": "string"
      },
      "datastore_name": {
        "computed": true,
        "description": "The user-generated name for the Data Store.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "datastore_status": {
        "computed": true,
        "description": "The status of the Data Store. Possible statuses are 'CREATING', 'ACTIVE', 'DELETING', or 'DELETED'.",
        "description_kind": "plain",
        "type": "string"
      },
      "datastore_type_version": {
        "description": "The FHIR version. Only R4 version data is supported.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "identity_provider_configuration": {
        "computed": true,
        "description": "The identity provider configuration for the datastore",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "authorization_strategy": {
              "description": "Type of Authorization Strategy. The two types of supported Authorization strategies are SMART_ON_FHIR_V1 and AWS_AUTH.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "fine_grained_authorization_enabled": {
              "computed": true,
              "description": "Flag to indicate if fine-grained authorization will be enabled for the datastore",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "idp_lambda_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the Lambda function that will be used to decode the access token created by the authorization server.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "metadata": {
              "computed": true,
              "description": "The JSON metadata elements for identity provider configuration.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "preload_data_config": {
        "computed": true,
        "description": "The preloaded data configuration for the Data Store. Only data preloaded from Synthea is supported.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "preload_data_type": {
              "description": "The type of preloaded data. Only Synthea preloaded data is supported.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "sse_configuration": {
        "computed": true,
        "description": "The server-side encryption key configuration for a customer provided encryption key.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kms_encryption_config": {
              "description": "The customer-managed-key (CMK) used when creating a Data Store. If a customer owned key is not specified, an AWS owned key will be used for encryption.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cmk_type": {
                    "description": "The type of customer-managed-key (CMK) used for encryption. The two types of supported CMKs are customer owned CMKs and AWS owned CMKs.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "kms_key_id": {
                    "computed": true,
                    "description": "The KMS encryption key id/alias used to encrypt the Data Store contents at rest.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key of the tag.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value of the tag.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "HealthLake FHIR Datastore",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccHealthlakeFhirDatastoreSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccHealthlakeFhirDatastore), &result)
	return &result
}
