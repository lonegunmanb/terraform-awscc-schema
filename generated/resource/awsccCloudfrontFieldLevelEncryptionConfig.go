package resource

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
        "description": "A unique value that identifies the create request. Any string unique within the account works; CloudFront refuses to change it once set, so supplying a different value replaces the configuration. It is also the only value CloudFront returns that identifies a configuration as this stack's, so cleaning up after a failed create depends on the template keeping it.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "comment": {
        "computed": true,
        "description": "An optional comment describing the field-level encryption configuration. CloudFront stores an omitted comment and an explicitly empty one identically as an empty string, and Read normalizes that empty string back to absent -- so omitting Comment round-trips, while asking for an empty string does not and would report perpetual drift. ` + "`" + `minLength: 1` + "`" + ` refuses the empty string at template validation; omission is the way to express no comment. The service documents a maximum of 128 characters, and ` + "`" + `maxLength` + "`" + ` moves that refusal from the API to template validation.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "content_type_profile_config": {
        "description": "How CloudFront handles content types, including the profile paired with each one.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "content_type_profiles": {
              "description": "The content types paired with a profile. Required, and at least one item, because the service synthesizes one default item when the list is absent or empty, so neither state round-trips.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "content_type": {
                    "description": "The content type CloudFront encrypts. The service accepts only application/x-www-form-urlencoded.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "format": {
                    "description": "The encryption format CloudFront uses for this content type.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "profile_id": {
                    "description": "The identifier of the field-level encryption profile CloudFront uses for this content type. Required, so that it cannot be omitted: the service accepts a content-type profile with the member absent on create and refuses one on update, and an update sends the whole profile list, so a configuration created that way could never be updated again. Supply an empty string to configure a content type and forward it WITHOUT field-level encryption. That mode is available through this resource and round-trips: create, update and read all accept and return the empty string, and it is the service's own encoding for the state -- asked for a content-type list with no items, CloudFront returns one item whose ProfileId is empty. Note that this is not the same as omitting the content type from the configuration, because ForwardWhenContentTypeIsUnknown governs only content types absent from it.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "required": true
            },
            "forward_when_content_type_is_unknown": {
              "description": "Whether CloudFront forwards a request when the content type names a profile that does not exist.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "field_level_encryption_config_id": {
        "computed": true,
        "description": "The identifier CloudFront assigns to the field-level encryption configuration.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description": "The time the field-level encryption configuration was last changed.",
        "description_kind": "plain",
        "type": "string"
      },
      "query_arg_profile_config": {
        "description": "How CloudFront handles query arguments, including the profile paired with each one.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "forward_when_query_arg_profile_is_unknown": {
              "description": "Whether CloudFront forwards a request when the query argument names a profile that does not exist.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "query_arg_profiles": {
              "description": "The query arguments paired with a profile. Required because the service returns an empty list rather than preserving absence, so an omitted value cannot round-trip. An empty list is accepted and is preserved.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "profile_id": {
                    "description": "The identifier of the field-level encryption profile CloudFront uses for this query argument.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "query_arg": {
                    "description": "The query argument that CloudFront encrypts with the paired profile.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "required": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      }
    },
    "description": "Resource Type definition for AWS::CloudFront::FieldLevelEncryptionConfig. A field-level encryption configuration, which specifies the query arg profiles and content type profiles CloudFront uses to encrypt sensitive fields in viewer requests.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudfrontFieldLevelEncryptionConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudfrontFieldLevelEncryptionConfig), &result)
	return &result
}
