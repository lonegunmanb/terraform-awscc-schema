package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudfrontFieldLevelEncryptionConfig = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the field-level encryption configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "caller_reference": {
        "computed": true,
        "description": "A unique value that identifies the create request. Any string unique within the account works; CloudFront refuses to change it once set, so supplying a different value replaces the configuration. It is also the only value CloudFront returns that identifies a configuration as this stack's, so cleaning up after a failed create depends on the template keeping it.",
        "description_kind": "plain",
        "type": "string"
      },
      "comment": {
        "computed": true,
        "description": "An optional comment describing the field-level encryption configuration. CloudFront stores an omitted comment and an explicitly empty one identically as an empty string, and Read normalizes that empty string back to absent -- so omitting Comment round-trips, while asking for an empty string does not and would report perpetual drift. ` + "`" + `minLength: 1` + "`" + ` refuses the empty string at template validation; omission is the way to express no comment. The service documents a maximum of 128 characters, and ` + "`" + `maxLength` + "`" + ` moves that refusal from the API to template validation.",
        "description_kind": "plain",
        "type": "string"
      },
      "content_type_profile_config": {
        "computed": true,
        "description": "How CloudFront handles content types, including the profile paired with each one.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "content_type_profiles": {
              "computed": true,
              "description": "The content types paired with a profile. Required, and at least one item, because the service synthesizes one default item when the list is absent or empty, so neither state round-trips.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "content_type": {
                    "computed": true,
                    "description": "The content type CloudFront encrypts. The service accepts only application/x-www-form-urlencoded.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "format": {
                    "computed": true,
                    "description": "The encryption format CloudFront uses for this content type.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "profile_id": {
                    "computed": true,
                    "description": "The identifier of the field-level encryption profile CloudFront uses for this content type. Required, so that it cannot be omitted: the service accepts a content-type profile with the member absent on create and refuses one on update, and an update sends the whole profile list, so a configuration created that way could never be updated again. Supply an empty string to configure a content type and forward it WITHOUT field-level encryption. That mode is available through this resource and round-trips: create, update and read all accept and return the empty string, and it is the service's own encoding for the state -- asked for a content-type list with no items, CloudFront returns one item whose ProfileId is empty. Note that this is not the same as omitting the content type from the configuration, because ForwardWhenContentTypeIsUnknown governs only content types absent from it.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "forward_when_content_type_is_unknown": {
              "computed": true,
              "description": "Whether CloudFront forwards a request when the content type names a profile that does not exist.",
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "field_level_encryption_config_id": {
        "computed": true,
        "description": "The identifier CloudFront assigns to the field-level encryption configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description": "The time the field-level encryption configuration was last changed.",
        "description_kind": "plain",
        "type": "string"
      },
      "query_arg_profile_config": {
        "computed": true,
        "description": "How CloudFront handles query arguments, including the profile paired with each one.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "forward_when_query_arg_profile_is_unknown": {
              "computed": true,
              "description": "Whether CloudFront forwards a request when the query argument names a profile that does not exist.",
              "description_kind": "plain",
              "type": "bool"
            },
            "query_arg_profiles": {
              "computed": true,
              "description": "The query arguments paired with a profile. Required because the service returns an empty list rather than preserving absence, so an omitted value cannot round-trip. An empty list is accepted and is preserved.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "profile_id": {
                    "computed": true,
                    "description": "The identifier of the field-level encryption profile CloudFront uses for this query argument.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "query_arg": {
                    "computed": true,
                    "description": "The query argument that CloudFront encrypts with the paired profile.",
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
    "description": "Data Source schema for AWS::CloudFront::FieldLevelEncryptionConfig",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCloudfrontFieldLevelEncryptionConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudfrontFieldLevelEncryptionConfig), &result)
	return &result
}
