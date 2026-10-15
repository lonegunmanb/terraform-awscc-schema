package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLicensemanagerlinuxsubscriptionsSubscriptionProvider = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_successful_data_retrieval_time": {
        "computed": true,
        "description": "The timestamp from the last time that License Manager retrieved subscription details from the registered third-party Linux subscription provider.",
        "description_kind": "plain",
        "type": "string"
      },
      "secret_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the Secrets Manager secret that contains the access token for the third-party Linux subscription provider.",
        "description_kind": "plain",
        "type": "string"
      },
      "subscription_provider_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the Linux subscription provider resource that you registered.",
        "description_kind": "plain",
        "type": "string"
      },
      "subscription_provider_source": {
        "computed": true,
        "description": "The Linux subscription provider that the registration is for.",
        "description_kind": "plain",
        "type": "string"
      },
      "subscription_provider_status": {
        "computed": true,
        "description": "The status of the third-party subscription provider access token from the last subscription data request.",
        "description_kind": "plain",
        "type": "string"
      },
      "subscription_provider_status_message": {
        "computed": true,
        "description": "A message that describes the status of the third-party subscription provider access token.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The metadata tags assigned to the registered subscription provider.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::LicenseManagerLinuxSubscriptions::SubscriptionProvider",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLicensemanagerlinuxsubscriptionsSubscriptionProviderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLicensemanagerlinuxsubscriptionsSubscriptionProvider), &result)
	return &result
}
