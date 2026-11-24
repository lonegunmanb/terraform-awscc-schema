package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApiportalPortal = `{
  "block": {
    "attributes": {
      "authorization": {
        "description": "Represents an authorization configuration for a portal.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "authorization_config": {
              "description": "Exactly one of the members is set: None for no authorization, or CognitoConfig for Amazon Cognito.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cognito_config": {
                    "computed": true,
                    "description": "The configuration for using Amazon Cognito user pools to control access to your portal.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "app_client_id": {
                          "computed": true,
                          "description": "The app client ID.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "user_pool_arn": {
                          "computed": true,
                          "description": "The user pool ARN.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "user_pool_domain": {
                          "computed": true,
                          "description": "The user pool domain.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "none": {
                    "computed": true,
                    "description": "Provide no authorization for your portal. This makes your portal publicly accessible on the web.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            },
            "fine_grained_access_control": {
              "description": "Whether fine-grained access control is enforced for the portal. No default: the template states the posture explicitly.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "endpoint_configuration": {
        "description": "Represents an endpoint configuration for a portal.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "acm_managed": {
              "computed": true,
              "description": "Represents a domain name and certificate for a portal.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "certificate_arn": {
                    "computed": true,
                    "description": "The certificate ARN.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "domain_name": {
                    "computed": true,
                    "description": "The domain name.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "certificate_arn": {
              "computed": true,
              "description": "The ARN of the ACM certificate.",
              "description_kind": "plain",
              "type": "string"
            },
            "domain_name": {
              "computed": true,
              "description": "The domain name.",
              "description_kind": "plain",
              "type": "string"
            },
            "none": {
              "computed": true,
              "description": "Provide no authorization for your portal. This makes your portal publicly accessible on the web.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "portal_default_domain_name": {
              "computed": true,
              "description": "The portal default domain name. Generated and managed by API Portal.",
              "description_kind": "plain",
              "type": "string"
            },
            "portal_domain_hosted_zone_id": {
              "computed": true,
              "description": "The portal domain hosted zone identifier.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "included_product_arns": {
        "computed": true,
        "description": "The ARNs of the products included in the portal.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "logo_uri": {
        "computed": true,
        "description": "The logo URI",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "portal_arn": {
        "computed": true,
        "description": "The ARN of the portal.",
        "description_kind": "plain",
        "type": "string"
      },
      "portal_content": {
        "description": "Contains the content that is visible to portal consumers including the themes, display names, and description.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "description": {
              "computed": true,
              "description": "A description of the portal.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "display_name": {
              "description": "The display name for the portal.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "theme": {
              "description": "Defines the theme for a portal.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "custom_colors": {
                    "description": "Represents custom colors for a published portal.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "accent_color": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "background_color": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "error_validation_color": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "header_color": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "navigation_color": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "text_color": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  },
                  "logo_uploaded_at": {
                    "computed": true,
                    "description": "The timestamp when the logo was last uploaded.",
                    "description_kind": "plain",
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
        "required": true
      },
      "portal_id": {
        "computed": true,
        "description": "The portal identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "publish_description": {
        "computed": true,
        "description": "A user-written description of the changes made in the last published version of the portal.",
        "description_kind": "plain",
        "type": "string"
      },
      "publish_status": {
        "computed": true,
        "description": "The current publishing status of the portal.",
        "description_kind": "plain",
        "type": "string"
      },
      "published_at": {
        "computed": true,
        "description": "The timestamp when the portal was last published.",
        "description_kind": "plain",
        "type": "string"
      },
      "rum_app_monitor_name": {
        "computed": true,
        "description": "The name of the Amazon CloudWatch RUM app monitor.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status_exception": {
        "computed": true,
        "description": "Failure detail for the portal's publish status.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "exception": {
              "computed": true,
              "description": "The reason the portal's last publish or disable failed.",
              "description_kind": "plain",
              "type": "string"
            },
            "message": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description": "The tags associated with a portal.",
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
      "updated_at": {
        "computed": true,
        "description": "The timestamp when the portal configuration was last modified.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::APIPortal::Portal",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApiportalPortalSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApiportalPortal), &result)
	return &result
}
