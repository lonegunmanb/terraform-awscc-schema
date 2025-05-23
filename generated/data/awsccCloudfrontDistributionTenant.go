package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudfrontDistributionTenant = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "connection_group_id": {
        "computed": true,
        "description": "The ID of the connection group for the distribution tenant. If you don't specify a connection group, CloudFront uses the default connection group.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "customizations": {
        "computed": true,
        "description": "Customizations for the distribution tenant. For each distribution tenant, you can specify the geographic restrictions, and the Amazon Resource Names (ARNs) for the ACM certificate and WAF web ACL. These are specific values that you can override or disable from the multi-tenant distribution that was used to create the distribution tenant.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "certificate": {
              "computed": true,
              "description": "The ACMlong (ACM) certificate.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of the ACM certificate.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "geo_restrictions": {
              "computed": true,
              "description": "The geographic restrictions.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "locations": {
                    "computed": true,
                    "description": "The locations for geographic restrictions.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "restriction_type": {
                    "computed": true,
                    "description": "The method that you want to use to restrict distribution of your content by country:\n  +  ` + "`" + `` + "`" + `none` + "`" + `` + "`" + `: No geographic restriction is enabled, meaning access to content is not restricted by client geo location.\n  +  ` + "`" + `` + "`" + `blacklist` + "`" + `` + "`" + `: The ` + "`" + `` + "`" + `Location` + "`" + `` + "`" + ` elements specify the countries in which you don't want CloudFront to distribute your content.\n  +  ` + "`" + `` + "`" + `whitelist` + "`" + `` + "`" + `: The ` + "`" + `` + "`" + `Location` + "`" + `` + "`" + ` elements specify the countries in which you want CloudFront to distribute your content.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "web_acl": {
              "computed": true,
              "description": "The WAF web ACL.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "action": {
                    "computed": true,
                    "description": "The action for the WAF web ACL customization. You can specify ` + "`" + `` + "`" + `override` + "`" + `` + "`" + ` to specify a separate WAF web ACL for the distribution tenant. If you specify ` + "`" + `` + "`" + `disable` + "`" + `` + "`" + `, the distribution tenant won't have WAF web ACL protections and won't inherit from the multi-tenant distribution.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of the WAF web ACL.",
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
      "distribution_id": {
        "computed": true,
        "description": "The ID of the multi-tenant distribution.",
        "description_kind": "plain",
        "type": "string"
      },
      "distribution_tenant_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_results": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "domain": {
              "computed": true,
              "description": "The specified domain.",
              "description_kind": "plain",
              "type": "string"
            },
            "status": {
              "computed": true,
              "description": "Whether the domain is active or inactive.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "domains": {
        "computed": true,
        "description": "The domains associated with the distribution tenant.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "e_tag": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enabled": {
        "computed": true,
        "description": "Indicates whether the distribution tenant is in an enabled state. If disabled, the distribution tenant won't serve traffic.",
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "managed_certificate_request": {
        "computed": true,
        "description": "An object that represents the request for the Amazon CloudFront managed ACM certificate.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "certificate_transparency_logging_preference": {
              "computed": true,
              "description": "You can opt out of certificate transparency logging by specifying the ` + "`" + `` + "`" + `disabled` + "`" + `` + "`" + ` option. Opt in by specifying ` + "`" + `` + "`" + `enabled` + "`" + `` + "`" + `. For more information, see [Certificate Transparency Logging](https://docs.aws.amazon.com/acm/latest/userguide/acm-concepts.html#concept-transparency) in the *User Guide*.",
              "description_kind": "plain",
              "type": "string"
            },
            "primary_domain_name": {
              "computed": true,
              "description": "The primary domain name associated with the CloudFront managed ACM certificate.",
              "description_kind": "plain",
              "type": "string"
            },
            "validation_token_host": {
              "computed": true,
              "description": "Specify how the HTTP validation token will be served when requesting the CloudFront managed ACM certificate.\n  +  For ` + "`" + `` + "`" + `cloudfront` + "`" + `` + "`" + `, CloudFront will automatically serve the validation token. Choose this mode if you can point the domain's DNS to CloudFront immediately.\n  +  For ` + "`" + `` + "`" + `self-hosted` + "`" + `` + "`" + `, you serve the validation token from your existing infrastructure. Choose this mode when you need to maintain current traffic flow while your certificate is being issued. You can place the validation token at the well-known path on your existing web server, wait for ACM to validate and issue the certificate, and then update your DNS to point to CloudFront.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "name": {
        "computed": true,
        "description": "The name of the distribution tenant.",
        "description_kind": "plain",
        "type": "string"
      },
      "parameters": {
        "computed": true,
        "description": "A list of parameter values to add to the resource. A parameter is specified as a key-value pair. A valid parameter value must exist for any parameter that is marked as required in the multi-tenant distribution.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "name": {
              "computed": true,
              "description": "The parameter name.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The parameter value.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A complex type that contains zero or more ` + "`" + `` + "`" + `Tag` + "`" + `` + "`" + ` elements.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "A string that contains ` + "`" + `` + "`" + `Tag` + "`" + `` + "`" + ` key.\n The string length should be between 1 and 128 characters. Valid characters include ` + "`" + `` + "`" + `a-z` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `A-Z` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `0-9` + "`" + `` + "`" + `, space, and the special characters ` + "`" + `` + "`" + `_ - . : / = + @` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "A string that contains an optional ` + "`" + `` + "`" + `Tag` + "`" + `` + "`" + ` value.\n The string length should be between 0 and 256 characters. Valid characters include ` + "`" + `` + "`" + `a-z` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `A-Z` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `0-9` + "`" + `` + "`" + `, space, and the special characters ` + "`" + `` + "`" + `_ - . : / = + @` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::CloudFront::DistributionTenant",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCloudfrontDistributionTenantSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudfrontDistributionTenant), &result)
	return &result
}
