package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccHealthlakeDataTransformationProfile = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the data transformation profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description": "The identifier (key ID or ARN) of a customer-managed KMS key used to encrypt the profile's template content at rest. If omitted, an AWS owned key is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "profile_description": {
        "computed": true,
        "description": "A human-readable description of the profile's purpose.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "profile_id": {
        "computed": true,
        "description": "The unique, server-generated identifier of the profile (32-character lowercase hexadecimal).",
        "description_kind": "plain",
        "type": "string"
      },
      "profile_name": {
        "description": "The human-readable name of the profile.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source": {
        "computed": true,
        "description": "The source from which to create the profile's initial template content. Exactly one of the members must be specified. Use StarterProfile (C-CDA only), ProfileMapping (C-CDA or CSV), or ExistingVersionedProfileId to clone an existing profile. Each produces a published profile.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "existing_versioned_profile_id": {
              "computed": true,
              "description": "Create the profile by cloning a specific version of an existing profile.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "profile_id": {
                    "computed": true,
                    "description": "The unique identifier of the source profile to clone.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "version": {
                    "computed": true,
                    "description": "The version number of the source profile to clone.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "profile_mapping": {
              "computed": true,
              "description": "Create the profile from raw Velocity template mapping content.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "profile_mapping": {
                    "computed": true,
                    "description": "Map of template file paths to their Velocity template content.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "starter_profile": {
              "computed": true,
              "description": "Create the profile from a predefined starter profile of transformation templates.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "starter_profile_name": {
                    "computed": true,
                    "description": "The name of the starter profile to seed the profile from.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "source_format": {
        "description": "The source format that this profile converts from.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this profile.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "target_format": {
        "computed": true,
        "description": "The target format that this profile converts to. Always FHIR_R4.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Creates a Data Transformation Profile in AWS HealthLake that converts healthcare data from a source format (such as C-CDA or CSV) into FHIR R4. A profile is immutable once created; to change its template content, replace the resource. Only its tags can be updated in place.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccHealthlakeDataTransformationProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccHealthlakeDataTransformationProfile), &result)
	return &result
}
