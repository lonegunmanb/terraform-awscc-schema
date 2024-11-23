package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudfrontOriginAccessControl = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "origin_access_control_config": {
        "description": "The origin access control.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "description": {
              "computed": true,
              "description": "A description of the origin access control.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "description": "A name to identify the origin access control. You can specify up to 64 characters.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "origin_access_control_origin_type": {
              "description": "The type of origin that this origin access control is for.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "signing_behavior": {
              "description": "Specifies which requests CloudFront signs (adds authentication information to). Specify ` + "`" + `` + "`" + `always` + "`" + `` + "`" + ` for the most common use case. For more information, see [origin access control advanced settings](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-restricting-access-to-s3.html#oac-advanced-settings) in the *Amazon CloudFront Developer Guide*.\n This field can have one of the following values:\n  +   ` + "`" + `` + "`" + `always` + "`" + `` + "`" + ` ? CloudFront signs all origin requests, overwriting the ` + "`" + `` + "`" + `Authorization` + "`" + `` + "`" + ` header from the viewer request if one exists.\n  +   ` + "`" + `` + "`" + `never` + "`" + `` + "`" + ` ? CloudFront doesn't sign any origin requests. This value turns off origin access control for all origins in all distributions that use this origin access control.\n  +   ` + "`" + `` + "`" + `no-override` + "`" + `` + "`" + ` ? If the viewer request doesn't contain the ` + "`" + `` + "`" + `Authorization` + "`" + `` + "`" + ` header, then CloudFront signs the origin request. If the viewer request contains the ` + "`" + `` + "`" + `Authorization` + "`" + `` + "`" + ` header, then CloudFront doesn't sign the origin request and instead passes along the ` + "`" + `` + "`" + `Authorization` + "`" + `` + "`" + ` header from the viewer request. *WARNING: To pass along the Authorization header from the viewer request, you must add the Authorization header to a cache policy for all cache behaviors that use origins associated with this origin access control.*",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "signing_protocol": {
              "description": "The signing protocol of the origin access control, which determines how CloudFront signs (authenticates) requests. The only valid value is ` + "`" + `` + "`" + `sigv4` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "origin_access_control_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Creates a new origin access control in CloudFront. After you create an origin access control, you can add it to an origin in a CloudFront distribution so that CloudFront sends authenticated (signed) requests to the origin.\n This makes it possible to block public access to the origin, allowing viewers (users) to access the origin's content only through CloudFront.\n For more information about using a CloudFront origin access control, see [Restricting access to an origin](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-restricting-access-to-origin.html) in the *Amazon CloudFront Developer Guide*.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudfrontOriginAccessControlSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudfrontOriginAccessControl), &result)
	return &result
}
