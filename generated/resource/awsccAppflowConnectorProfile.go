package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppflowConnectorProfile = `{
  "block": {
    "attributes": {
      "connection_mode": {
        "description": "Mode in which data transfer should be enabled. Private connection mode is currently enabled for Salesforce, Snowflake, Trendmicro and Singular",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "connector_label": {
        "computed": true,
        "description": "The label of the connector. The label is unique for each ConnectorRegistration in your AWS account. Only needed if calling for CUSTOMCONNECTOR connector type/.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "connector_profile_arn": {
        "computed": true,
        "description": "Unique identifier for connector profile resources",
        "description_kind": "plain",
        "type": "string"
      },
      "connector_profile_config": {
        "computed": true,
        "description": "Connector specific configurations needed to create connector profile",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "connector_profile_credentials": {
              "computed": true,
              "description": "Connector specific configuration needed to create connector profile based on Authentication mechanism",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "amplitude": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "api_key": {
                          "description": "A unique alphanumeric identi?er used to authenticate a user, developer, or calling program to your API.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "secret_key": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "custom_connector": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "api_key": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "api_key": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "api_secret_key": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "authentication_type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "basic": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "password": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "username": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "custom": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "credentials_map": {
                                "computed": true,
                                "description": "A map for properties for custom authentication.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              },
                              "custom_authentication_type": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "oauth_2": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "access_token": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "client_id": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "client_secret": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "o_auth_request": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "auth_code": {
                                      "computed": true,
                                      "description": "The code provided by the connector when it has been authenticated via the connected app.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "redirect_uri": {
                                      "computed": true,
                                      "description": "The URL to which the authentication server redirects the browser after authorization has been\ngranted.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              },
                              "refresh_token": {
                                "computed": true,
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
                  "datadog": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "api_key": {
                          "description": "A unique alphanumeric identi?er used to authenticate a user, developer, or calling program to your API.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "application_key": {
                          "description": "Application keys, in conjunction with your API key, give you full access to Datadog?s programmatic API. Application keys are associated with the user account that created them. The application key is used to log all requests made to the API.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "dynatrace": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "api_token": {
                          "description": "The API tokens used by Dynatrace API to authenticate various API calls.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "google_analytics": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "access_token": {
                          "computed": true,
                          "description": "The credentials used to access protected resources.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "client_id": {
                          "description": "The identi?er for the desired client.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "client_secret": {
                          "description": "The client secret used by the oauth client to authenticate to the authorization server.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "connector_o_auth_request": {
                          "computed": true,
                          "description": "The oauth needed to request security tokens from the connector endpoint.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "auth_code": {
                                "computed": true,
                                "description": "The code provided by the connector when it has been authenticated via the connected app.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "redirect_uri": {
                                "computed": true,
                                "description": "The URL to which the authentication server redirects the browser after authorization has been\ngranted.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "refresh_token": {
                          "computed": true,
                          "description": "The credentials used to acquire new access tokens.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "infor_nexus": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "access_key_id": {
                          "description": "The Access Key portion of the credentials.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "datakey": {
                          "description": "The encryption keys used to encrypt data.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "secret_access_key": {
                          "description": "The secret key used to sign requests.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "user_id": {
                          "description": "The identi?er for the user.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "marketo": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "access_token": {
                          "computed": true,
                          "description": "The credentials used to access protected resources.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "client_id": {
                          "description": "The identi?er for the desired client.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "client_secret": {
                          "description": "The client secret used by the oauth client to authenticate to the authorization server.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "connector_o_auth_request": {
                          "computed": true,
                          "description": "The oauth needed to request security tokens from the connector endpoint.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "auth_code": {
                                "computed": true,
                                "description": "The code provided by the connector when it has been authenticated via the connected app.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "redirect_uri": {
                                "computed": true,
                                "description": "The URL to which the authentication server redirects the browser after authorization has been\ngranted.",
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
                  "pardot": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "access_token": {
                          "computed": true,
                          "description": "The credentials used to access protected resources.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "client_credentials_arn": {
                          "computed": true,
                          "description": "The client credentials to fetch access token and refresh token.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "connector_o_auth_request": {
                          "computed": true,
                          "description": "The oauth needed to request security tokens from the connector endpoint.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "auth_code": {
                                "computed": true,
                                "description": "The code provided by the connector when it has been authenticated via the connected app.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "redirect_uri": {
                                "computed": true,
                                "description": "The URL to which the authentication server redirects the browser after authorization has been\ngranted.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "refresh_token": {
                          "computed": true,
                          "description": "The credentials used to acquire new access tokens.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "redshift": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "password": {
                          "computed": true,
                          "description": "The password that corresponds to the username.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "username": {
                          "computed": true,
                          "description": "The name of the user.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "salesforce": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "access_token": {
                          "computed": true,
                          "description": "The credentials used to access protected resources.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "client_credentials_arn": {
                          "computed": true,
                          "description": "The client credentials to fetch access token and refresh token.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "connector_o_auth_request": {
                          "computed": true,
                          "description": "The oauth needed to request security tokens from the connector endpoint.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "auth_code": {
                                "computed": true,
                                "description": "The code provided by the connector when it has been authenticated via the connected app.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "redirect_uri": {
                                "computed": true,
                                "description": "The URL to which the authentication server redirects the browser after authorization has been\ngranted.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "refresh_token": {
                          "computed": true,
                          "description": "The credentials used to acquire new access tokens.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "sapo_data": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "basic_auth_credentials": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "password": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "username": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "o_auth_credentials": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "access_token": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "client_id": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "client_secret": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "connector_o_auth_request": {
                                "computed": true,
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "auth_code": {
                                      "computed": true,
                                      "description": "The code provided by the connector when it has been authenticated via the connected app.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "redirect_uri": {
                                      "computed": true,
                                      "description": "The URL to which the authentication server redirects the browser after authorization has been\ngranted.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              },
                              "refresh_token": {
                                "computed": true,
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
                  "service_now": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "password": {
                          "description": "The password that corresponds to the username.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "username": {
                          "description": "The name of the user.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "singular": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "api_key": {
                          "description": "A unique alphanumeric identi?er used to authenticate a user, developer, or calling program to your API.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "slack": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "access_token": {
                          "computed": true,
                          "description": "The credentials used to access protected resources.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "client_id": {
                          "description": "The identi?er for the desired client.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "client_secret": {
                          "description": "The client secret used by the oauth client to authenticate to the authorization server.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "connector_o_auth_request": {
                          "computed": true,
                          "description": "The oauth needed to request security tokens from the connector endpoint.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "auth_code": {
                                "computed": true,
                                "description": "The code provided by the connector when it has been authenticated via the connected app.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "redirect_uri": {
                                "computed": true,
                                "description": "The URL to which the authentication server redirects the browser after authorization has been\ngranted.",
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
                  "snowflake": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "password": {
                          "description": "The password that corresponds to the username.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "username": {
                          "description": "The name of the user.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "trendmicro": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "api_secret_key": {
                          "description": "The Secret Access Key portion of the credentials.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "veeva": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "password": {
                          "description": "The password that corresponds to the username.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "username": {
                          "description": "The name of the user.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "zendesk": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "access_token": {
                          "computed": true,
                          "description": "The credentials used to access protected resources.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "client_id": {
                          "description": "The identi?er for the desired client.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "client_secret": {
                          "description": "The client secret used by the oauth client to authenticate to the authorization server.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "connector_o_auth_request": {
                          "computed": true,
                          "description": "The oauth needed to request security tokens from the connector endpoint.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "auth_code": {
                                "computed": true,
                                "description": "The code provided by the connector when it has been authenticated via the connected app.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "redirect_uri": {
                                "computed": true,
                                "description": "The URL to which the authentication server redirects the browser after authorization has been\ngranted.",
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
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "connector_profile_properties": {
              "computed": true,
              "description": "Connector specific properties needed to create connector profile - currently not needed for Amplitude, Trendmicro, Googleanalytics and Singular",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "custom_connector": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "o_auth_2_properties": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "o_auth_2_grant_type": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "token_url": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "token_url_custom_properties": {
                                "computed": true,
                                "description": "A map for properties for custom connector Token Url.",
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
                        "profile_properties": {
                          "computed": true,
                          "description": "A map for properties for custom connector.",
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
                  "datadog": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "instance_url": {
                          "description": "The location of the Datadog resource",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "dynatrace": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "instance_url": {
                          "description": "The location of the Dynatrace resource",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "infor_nexus": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "instance_url": {
                          "description": "The location of the InforNexus resource",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "marketo": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "instance_url": {
                          "description": "The location of the Marketo resource",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "pardot": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "business_unit_id": {
                          "description": "The Business unit id of Salesforce Pardot instance to be connected",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "instance_url": {
                          "computed": true,
                          "description": "The location of the Salesforce Pardot resource",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "is_sandbox_environment": {
                          "computed": true,
                          "description": "Indicates whether the connector profile applies to a demo or production environment",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "redshift": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bucket_name": {
                          "description": "The name of the Amazon S3 bucket associated with Redshift.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "bucket_prefix": {
                          "computed": true,
                          "description": "The object key for the destination bucket in which Amazon AppFlow will place the ?les.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "cluster_identifier": {
                          "computed": true,
                          "description": "The unique identifier of the Amazon Redshift cluster.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "data_api_role_arn": {
                          "computed": true,
                          "description": "The Amazon Resource Name (ARN) of the IAM role that grants Amazon AppFlow access to the data through the Amazon Redshift Data API.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "database_name": {
                          "computed": true,
                          "description": "The name of the Amazon Redshift database that will store the transferred data.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "database_url": {
                          "computed": true,
                          "description": "The JDBC URL of the Amazon Redshift cluster.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "is_redshift_serverless": {
                          "computed": true,
                          "description": "If Amazon AppFlow will connect to Amazon Redshift Serverless or Amazon Redshift cluster.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "role_arn": {
                          "description": "The Amazon Resource Name (ARN) of the IAM role.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "workgroup_name": {
                          "computed": true,
                          "description": "The name of the Amazon Redshift serverless workgroup",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "salesforce": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "instance_url": {
                          "computed": true,
                          "description": "The location of the Salesforce resource",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "is_sandbox_environment": {
                          "computed": true,
                          "description": "Indicates whether the connector profile applies to a sandbox or production environment",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "sapo_data": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "application_host_url": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "application_service_path": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "client_number": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "logon_language": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "o_auth_properties": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "auth_code_url": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "o_auth_scopes": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "token_url": {
                                "computed": true,
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "port_number": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "private_link_service_name": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "service_now": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "instance_url": {
                          "description": "The location of the ServiceNow resource",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "slack": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "instance_url": {
                          "description": "The location of the Slack resource",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "snowflake": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "account_name": {
                          "computed": true,
                          "description": "The name of the account.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "bucket_name": {
                          "description": "The name of the Amazon S3 bucket associated with Snow?ake.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "bucket_prefix": {
                          "computed": true,
                          "description": "The bucket prefix that refers to the Amazon S3 bucket associated with Snow?ake.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "private_link_service_name": {
                          "computed": true,
                          "description": "The Snow?ake Private Link service name to be used for private data transfers.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "region": {
                          "computed": true,
                          "description": "The region of the Snow?ake account.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "stage": {
                          "description": "The name of the Amazon S3 stage that was created while setting up an Amazon S3 stage in the\nSnow?ake account. This is written in the following format: \u003c Database\u003e\u003c Schema\u003e\u003cStage Name\u003e.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "warehouse": {
                          "description": "The name of the Snow?ake warehouse.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "veeva": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "instance_url": {
                          "description": "The location of the Veeva resource",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "zendesk": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "instance_url": {
                          "description": "The location of the Zendesk resource",
                          "description_kind": "plain",
                          "required": true,
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
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "connector_profile_name": {
        "description": "The maximum number of items to retrieve in a single batch.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "connector_type": {
        "description": "List of Saas providers that need connector profile to be created",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "credentials_arn": {
        "computed": true,
        "description": "A unique Arn for Connector-Profile resource",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "kms_arn": {
        "computed": true,
        "description": "The ARN of the AWS Key Management Service (AWS KMS) key that's used to encrypt your function's environment variables. If it's not provided, AWS Lambda uses a default service key.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::AppFlow::ConnectorProfile",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAppflowConnectorProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppflowConnectorProfile), &result)
	return &result
}
