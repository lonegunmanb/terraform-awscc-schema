package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSecurityagentIntegration = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the integration.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "description": "The display name for the integration.",
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
      "installation_id": {
        "computed": true,
        "description": "The installation identifier from the integration provider.",
        "description_kind": "plain",
        "type": "string"
      },
      "integration_id": {
        "computed": true,
        "description": "The unique identifier of the integration.",
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description": "The identifier of the AWS KMS key to use for encrypting data associated with the integration. Can be a key ID, key ARN, alias name, or alias ARN.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "private_connection_name": {
        "computed": true,
        "description": "The name of an active private connection used to reach a self-hosted provider instance over private networking.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "provider_input": {
        "computed": true,
        "description": "The provider-specific input required to create the integration. Exactly one of the provider input objects must be specified.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "bitbucket": {
              "computed": true,
              "description": "The input required to create a Bitbucket integration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "code": {
                    "computed": true,
                    "description": "The OAuth 2.0 authorization code returned from the consent redirect.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "installation_id": {
                    "computed": true,
                    "description": "The Atlassian installation identifier, available from the Atlassian administration console.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "state": {
                    "computed": true,
                    "description": "The CSRF state token echoed back from the OAuth redirect.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "workspace": {
                    "computed": true,
                    "description": "The Bitbucket workspace slug that identifies the workspace to integrate.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "confluence": {
              "computed": true,
              "description": "The input required to create a Confluence integration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "code": {
                    "computed": true,
                    "description": "The OAuth 2.0 authorization code returned from the consent redirect.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "installation_id": {
                    "computed": true,
                    "description": "The Atlassian installation identifier, available from the Atlassian administration console.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "site_url": {
                    "computed": true,
                    "description": "The Confluence Cloud site URL, for example https://mysite.atlassian.net.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "state": {
                    "computed": true,
                    "description": "The CSRF state token echoed back from the OAuth redirect.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "github": {
              "computed": true,
              "description": "The input required to create a GitHub integration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "code": {
                    "computed": true,
                    "description": "The OAuth authorization code received from GitHub.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "installation_id": {
                    "computed": true,
                    "description": "The installation identifier provided by GitHub Enterprise Server on the install callback. Required for GitHub Enterprise Server integrations and ignored for GitHub.com.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "organization_name": {
                    "computed": true,
                    "description": "The name of the GitHub organization to integrate with.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "state": {
                    "computed": true,
                    "description": "The CSRF state token for validating the OAuth flow.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "target_url": {
                    "computed": true,
                    "description": "The HTTPS URL of a self-hosted GitHub Enterprise Server instance. Omit this value for GitHub.com.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "gitlab": {
              "computed": true,
              "description": "The input required to create a GitLab integration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "access_token": {
                    "computed": true,
                    "description": "The GitLab access token used to authenticate. This can be a personal access token or a group access token.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "group_id": {
                    "computed": true,
                    "description": "The identifier of the GitLab group. Required when TokenType is GROUP and ignored for personal tokens.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "target_url": {
                    "computed": true,
                    "description": "The HTTPS URL of a self-managed GitLab instance. Omit this value for GitLab SaaS (gitlab.com).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "token_type": {
                    "computed": true,
                    "description": "The type of GitLab access token.",
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
      "provider_name": {
        "description": "The integration provider.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "provider_type": {
        "computed": true,
        "description": "The type of the integration provider.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags to associate with the integration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "target_url": {
        "computed": true,
        "description": "The HTTPS URL of the customer self-hosted instance, such as a GitHub Enterprise Server or self-managed GitLab instance. This value is absent for SaaS integrations.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::SecurityAgent::Integration. Creates a connection between Security Agent and a third-party provider for code review and remediation.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSecurityagentIntegrationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecurityagentIntegration), &result)
	return &result
}
