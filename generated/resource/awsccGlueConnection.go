package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGlueConnection = `{
  "block": {
    "attributes": {
      "catalog_id": {
        "description": "The ID of the data catalog to create the catalog object in. Currently, this should be the AWS account ID.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "connection_input": {
        "description": "The connection properties used for this connection.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "athena_properties": {
              "computed": true,
              "description": "Connection properties specific to the Athena compute environment.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "authentication_configuration": {
              "computed": true,
              "description": "The authentication configuration used to connect to the connection.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "authentication_type": {
                    "computed": true,
                    "description": "A structure containing the authentication configuration in the CreateConnection request.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "basic_authentication_credentials": {
                    "computed": true,
                    "description": "For supplying basic auth credentials when not providing a SecretArn value",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "password": {
                          "computed": true,
                          "description": "The password used in the authentication configuration.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "username": {
                          "computed": true,
                          "description": "The username used in the authentication configuration.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "custom_authentication_credentials": {
                    "computed": true,
                    "description": "A structure containing the authentication credentials in the CreateConnection request.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "kms_key_arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of the KMS key used in the authentication configuration.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "o_auth_2_properties": {
                    "computed": true,
                    "description": "A structure containing properties for OAuth2 in the CreateConnection request.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "authorization_code_properties": {
                          "computed": true,
                          "description": "The set of properties required for the the OAuth2 AUTHORIZATION_CODE grant type workflow.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "authorization_code": {
                                "computed": true,
                                "description": "The authorization code used in the authentication configuration.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "redirect_uri": {
                                "computed": true,
                                "description": "The redirect URI where the user gets redirected to by authorization server when issuing an authorization code.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "o_auth_2_client_application": {
                          "computed": true,
                          "description": "The OAuth2 client app used for the connection.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "aws_managed_client_application_reference": {
                                "computed": true,
                                "description": "The reference to the SaaS-side client app that is AWS managed.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "user_managed_client_application_client_id": {
                                "computed": true,
                                "description": "The client application clientID if the ClientAppType is USER_MANAGED.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "o_auth_2_credentials": {
                          "computed": true,
                          "description": "A structure containing the OAuth2 credentials used in the authentication configuration.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "access_token": {
                                "computed": true,
                                "description": "The access token used in the authentication configuration.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "jwt_token": {
                                "computed": true,
                                "description": "The JSON Web Token (JWT) used when the authentication type is OAuth2.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "refresh_token": {
                                "computed": true,
                                "description": "The refresh token used when the authentication type is OAuth2.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "user_managed_client_application_client_secret": {
                                "computed": true,
                                "description": "The client application client secret if the client application is user managed.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "o_auth_2_grant_type": {
                          "computed": true,
                          "description": "The grant type used in the authentication configuration.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "token_url": {
                          "computed": true,
                          "description": "The URL used in the authentication configuration.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "token_url_parameters_map": {
                          "computed": true,
                          "description": "A map of key-value pairs used in the authentication configuration.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "secret_arn": {
                    "computed": true,
                    "description": "The secret manager ARN to store credentials in the CreateConnection request.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "connection_properties": {
              "computed": true,
              "description": "A map of key-value pairs used as parameters for this connection.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "connection_type": {
              "description": "The type of the connection that needs to be created.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "description": {
              "computed": true,
              "description": "A description of the connection.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "match_criteria": {
              "computed": true,
              "description": "A list of criteria that can be used in selecting this connection.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "name": {
              "computed": true,
              "description": "The name of the connection.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "physical_connection_requirements": {
              "computed": true,
              "description": "The physical connection requirements.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "availability_zone": {
                    "computed": true,
                    "description": "The availability zone where the connection is located.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "security_group_id_list": {
                    "computed": true,
                    "description": "The security group ID list used by the connection.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "subnet_id": {
                    "computed": true,
                    "description": "The subnet ID used by the connection.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "python_properties": {
              "computed": true,
              "description": "Connection properties specific to the Python compute environment.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "spark_properties": {
              "computed": true,
              "description": "Connection properties specific to the Spark compute environment.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "validate_credentials": {
              "computed": true,
              "description": "A flag to validate the credentials during create connection. Default is true.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "validate_for_compute_environments": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
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
      "name": {
        "computed": true,
        "description": "The name of the connection.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The collection of tags. Each tag element is associated with a given resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Glue::Connection",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccGlueConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGlueConnection), &result)
	return &result
}
