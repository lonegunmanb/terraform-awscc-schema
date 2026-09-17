package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGlueConnectionType = `{
  "block": {
    "attributes": {
      "connection_properties": {
        "computed": true,
        "description": "Configuration that defines the base URL and additional request parameters needed during connection creation.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "additional_request_parameters": {
              "computed": true,
              "description": "Key-value pairs of additional request parameters.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "allowed_values": {
                    "computed": true,
                    "description": "A list of allowed values for the property.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "default_value": {
                    "computed": true,
                    "description": "The default value for the property.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "key_override": {
                    "computed": true,
                    "description": "A key name to use when sending this property in API requests.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "The name of the property.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "property_location": {
                    "computed": true,
                    "description": "Specifies where this property should be included in REST requests.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "property_type": {
                    "computed": true,
                    "description": "The data type of this property.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "required": {
                    "computed": true,
                    "description": "Indicates whether the property is required.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "url": {
              "computed": true,
              "description": "Defines a property configuration for connection types.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "allowed_values": {
                    "computed": true,
                    "description": "A list of allowed values for the property.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "default_value": {
                    "computed": true,
                    "description": "The default value for the property.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "key_override": {
                    "computed": true,
                    "description": "A key name to use when sending this property in API requests.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "The name of the property.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "property_location": {
                    "computed": true,
                    "description": "Specifies where this property should be included in REST requests.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "property_type": {
                    "computed": true,
                    "description": "The data type of this property.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "required": {
                    "computed": true,
                    "description": "Indicates whether the property is required.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
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
      "connection_type": {
        "description": "The name of the connection type. Must be prefixed with REST-.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "connection_type_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the registered connection type.",
        "description_kind": "plain",
        "type": "string"
      },
      "connector_authentication_configuration": {
        "computed": true,
        "description": "Configuration that defines supported authentication types and required properties.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "authentication_types": {
              "computed": true,
              "description": "A list of authentication types supported.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "basic_authentication_properties": {
              "computed": true,
              "description": "Basic authentication configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "password": {
                    "computed": true,
                    "description": "Defines a secret property configuration. SECRET-type properties cannot have DefaultValue or AllowedValues.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "key_override": {
                          "computed": true,
                          "description": "A key name to use when sending this property in API requests.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "name": {
                          "computed": true,
                          "description": "The name of the property.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "property_location": {
                          "computed": true,
                          "description": "Specifies where this property should be included in REST requests.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "property_type": {
                          "computed": true,
                          "description": "The data type of this property. Must be SECRET for secret properties.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "required": {
                          "computed": true,
                          "description": "Indicates whether the property is required.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "username": {
                    "computed": true,
                    "description": "Defines a secret property configuration. SECRET-type properties cannot have DefaultValue or AllowedValues.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "key_override": {
                          "computed": true,
                          "description": "A key name to use when sending this property in API requests.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "name": {
                          "computed": true,
                          "description": "The name of the property.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "property_location": {
                          "computed": true,
                          "description": "Specifies where this property should be included in REST requests.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "property_type": {
                          "computed": true,
                          "description": "The data type of this property. Must be SECRET for secret properties.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "required": {
                          "computed": true,
                          "description": "Indicates whether the property is required.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
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
            "custom_authentication_properties": {
              "computed": true,
              "description": "Custom authentication configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "authentication_parameters": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "key_override": {
                          "computed": true,
                          "description": "A key name to use when sending this property in API requests.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "name": {
                          "computed": true,
                          "description": "The name of the property.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "property_location": {
                          "computed": true,
                          "description": "Specifies where this property should be included in REST requests.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "property_type": {
                          "computed": true,
                          "description": "The data type of this property. Must be SECRET for secret properties.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "required": {
                          "computed": true,
                          "description": "Indicates whether the property is required.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "o_auth_2_properties": {
              "computed": true,
              "description": "OAuth2 configuration container.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "authorization_code_properties": {
                    "computed": true,
                    "description": "OAuth2 authorization code configuration.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "authorization_code": {
                          "computed": true,
                          "description": "Defines a property configuration for connection types.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "allowed_values": {
                                "computed": true,
                                "description": "A list of allowed values for the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "default_value": {
                                "computed": true,
                                "description": "The default value for the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "key_override": {
                                "computed": true,
                                "description": "A key name to use when sending this property in API requests.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "name": {
                                "computed": true,
                                "description": "The name of the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "property_location": {
                                "computed": true,
                                "description": "Specifies where this property should be included in REST requests.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "property_type": {
                                "computed": true,
                                "description": "The data type of this property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "required": {
                                "computed": true,
                                "description": "Indicates whether the property is required.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "authorization_code_url": {
                          "computed": true,
                          "description": "Defines a property configuration for connection types.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "allowed_values": {
                                "computed": true,
                                "description": "A list of allowed values for the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "default_value": {
                                "computed": true,
                                "description": "The default value for the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "key_override": {
                                "computed": true,
                                "description": "A key name to use when sending this property in API requests.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "name": {
                                "computed": true,
                                "description": "The name of the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "property_location": {
                                "computed": true,
                                "description": "Specifies where this property should be included in REST requests.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "property_type": {
                                "computed": true,
                                "description": "The data type of this property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "required": {
                                "computed": true,
                                "description": "Indicates whether the property is required.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "client_id": {
                          "computed": true,
                          "description": "Defines a secret property configuration. SECRET-type properties cannot have DefaultValue or AllowedValues.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "key_override": {
                                "computed": true,
                                "description": "A key name to use when sending this property in API requests.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "name": {
                                "computed": true,
                                "description": "The name of the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "property_location": {
                                "computed": true,
                                "description": "Specifies where this property should be included in REST requests.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "property_type": {
                                "computed": true,
                                "description": "The data type of this property. Must be SECRET for secret properties.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "required": {
                                "computed": true,
                                "description": "Indicates whether the property is required.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "client_secret": {
                          "computed": true,
                          "description": "Defines a secret property configuration. SECRET-type properties cannot have DefaultValue or AllowedValues.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "key_override": {
                                "computed": true,
                                "description": "A key name to use when sending this property in API requests.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "name": {
                                "computed": true,
                                "description": "The name of the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "property_location": {
                                "computed": true,
                                "description": "Specifies where this property should be included in REST requests.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "property_type": {
                                "computed": true,
                                "description": "The data type of this property. Must be SECRET for secret properties.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "required": {
                                "computed": true,
                                "description": "Indicates whether the property is required.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "content_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "prompt": {
                          "computed": true,
                          "description": "Defines a property configuration for connection types.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "allowed_values": {
                                "computed": true,
                                "description": "A list of allowed values for the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "default_value": {
                                "computed": true,
                                "description": "The default value for the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "key_override": {
                                "computed": true,
                                "description": "A key name to use when sending this property in API requests.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "name": {
                                "computed": true,
                                "description": "The name of the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "property_location": {
                                "computed": true,
                                "description": "Specifies where this property should be included in REST requests.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "property_type": {
                                "computed": true,
                                "description": "The data type of this property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "required": {
                                "computed": true,
                                "description": "Indicates whether the property is required.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "redirect_uri": {
                          "computed": true,
                          "description": "Defines a property configuration for connection types.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "allowed_values": {
                                "computed": true,
                                "description": "A list of allowed values for the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "default_value": {
                                "computed": true,
                                "description": "The default value for the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "key_override": {
                                "computed": true,
                                "description": "A key name to use when sending this property in API requests.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "name": {
                                "computed": true,
                                "description": "The name of the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "property_location": {
                                "computed": true,
                                "description": "Specifies where this property should be included in REST requests.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "property_type": {
                                "computed": true,
                                "description": "The data type of this property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "required": {
                                "computed": true,
                                "description": "Indicates whether the property is required.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "request_method": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "scope": {
                          "computed": true,
                          "description": "Defines a property configuration for connection types.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "allowed_values": {
                                "computed": true,
                                "description": "A list of allowed values for the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "default_value": {
                                "computed": true,
                                "description": "The default value for the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "key_override": {
                                "computed": true,
                                "description": "A key name to use when sending this property in API requests.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "name": {
                                "computed": true,
                                "description": "The name of the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "property_location": {
                                "computed": true,
                                "description": "Specifies where this property should be included in REST requests.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "property_type": {
                                "computed": true,
                                "description": "The data type of this property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "required": {
                                "computed": true,
                                "description": "Indicates whether the property is required.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "token_url": {
                          "computed": true,
                          "description": "Defines a property configuration for connection types.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "allowed_values": {
                                "computed": true,
                                "description": "A list of allowed values for the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "default_value": {
                                "computed": true,
                                "description": "The default value for the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "key_override": {
                                "computed": true,
                                "description": "A key name to use when sending this property in API requests.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "name": {
                                "computed": true,
                                "description": "The name of the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "property_location": {
                                "computed": true,
                                "description": "Specifies where this property should be included in REST requests.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "property_type": {
                                "computed": true,
                                "description": "The data type of this property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "required": {
                                "computed": true,
                                "description": "Indicates whether the property is required.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "token_url_parameters": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "allowed_values": {
                                "computed": true,
                                "description": "A list of allowed values for the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "default_value": {
                                "computed": true,
                                "description": "The default value for the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "key_override": {
                                "computed": true,
                                "description": "A key name to use when sending this property in API requests.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "name": {
                                "computed": true,
                                "description": "The name of the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "property_location": {
                                "computed": true,
                                "description": "Specifies where this property should be included in REST requests.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "property_type": {
                                "computed": true,
                                "description": "The data type of this property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "required": {
                                "computed": true,
                                "description": "Indicates whether the property is required.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "optional": true
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "client_credentials_properties": {
                    "computed": true,
                    "description": "OAuth2 client credentials configuration.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "client_id": {
                          "computed": true,
                          "description": "Defines a secret property configuration. SECRET-type properties cannot have DefaultValue or AllowedValues.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "key_override": {
                                "computed": true,
                                "description": "A key name to use when sending this property in API requests.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "name": {
                                "computed": true,
                                "description": "The name of the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "property_location": {
                                "computed": true,
                                "description": "Specifies where this property should be included in REST requests.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "property_type": {
                                "computed": true,
                                "description": "The data type of this property. Must be SECRET for secret properties.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "required": {
                                "computed": true,
                                "description": "Indicates whether the property is required.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "client_secret": {
                          "computed": true,
                          "description": "Defines a secret property configuration. SECRET-type properties cannot have DefaultValue or AllowedValues.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "key_override": {
                                "computed": true,
                                "description": "A key name to use when sending this property in API requests.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "name": {
                                "computed": true,
                                "description": "The name of the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "property_location": {
                                "computed": true,
                                "description": "Specifies where this property should be included in REST requests.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "property_type": {
                                "computed": true,
                                "description": "The data type of this property. Must be SECRET for secret properties.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "required": {
                                "computed": true,
                                "description": "Indicates whether the property is required.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "content_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "request_method": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "scope": {
                          "computed": true,
                          "description": "Defines a property configuration for connection types.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "allowed_values": {
                                "computed": true,
                                "description": "A list of allowed values for the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "default_value": {
                                "computed": true,
                                "description": "The default value for the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "key_override": {
                                "computed": true,
                                "description": "A key name to use when sending this property in API requests.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "name": {
                                "computed": true,
                                "description": "The name of the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "property_location": {
                                "computed": true,
                                "description": "Specifies where this property should be included in REST requests.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "property_type": {
                                "computed": true,
                                "description": "The data type of this property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "required": {
                                "computed": true,
                                "description": "Indicates whether the property is required.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "token_url": {
                          "computed": true,
                          "description": "Defines a property configuration for connection types.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "allowed_values": {
                                "computed": true,
                                "description": "A list of allowed values for the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "default_value": {
                                "computed": true,
                                "description": "The default value for the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "key_override": {
                                "computed": true,
                                "description": "A key name to use when sending this property in API requests.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "name": {
                                "computed": true,
                                "description": "The name of the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "property_location": {
                                "computed": true,
                                "description": "Specifies where this property should be included in REST requests.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "property_type": {
                                "computed": true,
                                "description": "The data type of this property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "required": {
                                "computed": true,
                                "description": "Indicates whether the property is required.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "token_url_parameters": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "allowed_values": {
                                "computed": true,
                                "description": "A list of allowed values for the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "default_value": {
                                "computed": true,
                                "description": "The default value for the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "key_override": {
                                "computed": true,
                                "description": "A key name to use when sending this property in API requests.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "name": {
                                "computed": true,
                                "description": "The name of the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "property_location": {
                                "computed": true,
                                "description": "Specifies where this property should be included in REST requests.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "property_type": {
                                "computed": true,
                                "description": "The data type of this property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "required": {
                                "computed": true,
                                "description": "Indicates whether the property is required.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "optional": true
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "jwt_bearer_properties": {
                    "computed": true,
                    "description": "JWT bearer token configuration.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "content_type": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "jwt_token": {
                          "computed": true,
                          "description": "Defines a secret property configuration. SECRET-type properties cannot have DefaultValue or AllowedValues.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "key_override": {
                                "computed": true,
                                "description": "A key name to use when sending this property in API requests.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "name": {
                                "computed": true,
                                "description": "The name of the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "property_location": {
                                "computed": true,
                                "description": "Specifies where this property should be included in REST requests.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "property_type": {
                                "computed": true,
                                "description": "The data type of this property. Must be SECRET for secret properties.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "required": {
                                "computed": true,
                                "description": "Indicates whether the property is required.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "request_method": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "token_url": {
                          "computed": true,
                          "description": "Defines a property configuration for connection types.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "allowed_values": {
                                "computed": true,
                                "description": "A list of allowed values for the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "default_value": {
                                "computed": true,
                                "description": "The default value for the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "key_override": {
                                "computed": true,
                                "description": "A key name to use when sending this property in API requests.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "name": {
                                "computed": true,
                                "description": "The name of the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "property_location": {
                                "computed": true,
                                "description": "Specifies where this property should be included in REST requests.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "property_type": {
                                "computed": true,
                                "description": "The data type of this property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "required": {
                                "computed": true,
                                "description": "Indicates whether the property is required.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "token_url_parameters": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "allowed_values": {
                                "computed": true,
                                "description": "A list of allowed values for the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "default_value": {
                                "computed": true,
                                "description": "The default value for the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "key_override": {
                                "computed": true,
                                "description": "A key name to use when sending this property in API requests.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "name": {
                                "computed": true,
                                "description": "The name of the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "property_location": {
                                "computed": true,
                                "description": "Specifies where this property should be included in REST requests.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "property_type": {
                                "computed": true,
                                "description": "The data type of this property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "required": {
                                "computed": true,
                                "description": "Indicates whether the property is required.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "optional": true
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "o_auth_2_grant_type": {
                    "computed": true,
                    "description": "The OAuth2 grant type to use.",
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
      "description": {
        "computed": true,
        "description": "A description of the connection type.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "integration_type": {
        "computed": true,
        "description": "The integration type for the connection. Currently only REST is supported.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rest_configuration": {
        "description": "Configuration for HTTP request and response handling.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "entity_configurations": {
              "computed": true,
              "description": "A map of entity configurations.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "schema": {
                    "computed": true,
                    "description": "The schema definition for this entity.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "field_data_type": {
                          "computed": true,
                          "description": "The data type of the field.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "filter_overrides": {
                          "computed": true,
                          "description": "Configuration that defines per-field overrides for filter behavior, allowing individual fields to customize how filter operations are applied.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "between_configuration": {
                                "computed": true,
                                "description": "Configuration that defines how BETWEEN range filter operations are translated into REST API request parameters.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "high_bound_key": {
                                      "computed": true,
                                      "description": "The parameter name used for the upper bound value in a BETWEEN filter operation.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "low_bound_key": {
                                      "computed": true,
                                      "description": "The parameter name used for the lower bound value in a BETWEEN filter operation.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "template": {
                                      "computed": true,
                                      "description": "A template string for constructing the BETWEEN filter expression.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              },
                              "date_time_format": {
                                "computed": true,
                                "description": "The date and time format for filter expressions on this field, overriding the global DateTimeFormat.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "field_name": {
                                "computed": true,
                                "description": "An override for the field name to use in filter expressions, if different from the schema field name.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "operator_mappings": {
                                "computed": true,
                                "description": "A map of logical filter operators to their field-specific API representations, overriding the global operator mappings.",
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
                        "is_nullable": {
                          "computed": true,
                          "description": "Indicates whether this field can contain null values.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "is_orderable": {
                          "computed": true,
                          "description": "Indicates whether this field can be used for ordering results.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "is_partitionable": {
                          "computed": true,
                          "description": "Indicates whether this field can be used for partitioning queries to the data source.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "is_queryable": {
                          "computed": true,
                          "description": "Indicates whether this field can be used in filter predicates when querying data.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "name": {
                          "computed": true,
                          "description": "The name of the field.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "response_date_format": {
                          "computed": true,
                          "description": "The format pattern for parsing date values from API responses. Accepts Java DateTimeFormatter patterns, EPOCH_SECONDS, or EPOCH_MILLIS.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "map"
                    },
                    "optional": true
                  },
                  "source_configuration": {
                    "computed": true,
                    "description": "Configuration that defines how to make requests to endpoints.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "filter_configuration": {
                          "computed": true,
                          "description": "Configuration that defines how filter predicates are applied to REST API requests, supporting both query parameter and filter string strategies.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "between_configuration": {
                                "computed": true,
                                "description": "Configuration that defines how BETWEEN range filter operations are translated into REST API request parameters.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "high_bound_key": {
                                      "computed": true,
                                      "description": "The parameter name used for the upper bound value in a BETWEEN filter operation.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "low_bound_key": {
                                      "computed": true,
                                      "description": "The parameter name used for the lower bound value in a BETWEEN filter operation.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "template": {
                                      "computed": true,
                                      "description": "A template string for constructing the BETWEEN filter expression.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              },
                              "date_time_format": {
                                "computed": true,
                                "description": "The global date and time format for filter expressions.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "filter_mode": {
                                "computed": true,
                                "description": "The strategy for applying filters to requests.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "filter_string_configuration": {
                                "computed": true,
                                "description": "Configuration for constructing filter expression strings when using the FILTER_STRING filter mode.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "query_parameter_name": {
                                      "computed": true,
                                      "description": "The query parameter name used to send the constructed filter expression string in API requests.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "quote_character": {
                                      "computed": true,
                                      "description": "The character used to quote values when QuoteStringValues is true. Defaults to double quotes if not specified.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "quote_string_values": {
                                      "computed": true,
                                      "description": "Indicates whether string and date values should be wrapped with a quote character in the filter expression.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    }
                                  },
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              },
                              "operator_mappings": {
                                "computed": true,
                                "description": "A map of logical filter operators to their API-specific string representations.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "map",
                                  "string"
                                ]
                              },
                              "strip_quotes": {
                                "computed": true,
                                "description": "Indicates whether surrounding double quotes should be stripped from filter values before processing.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "pagination_configuration": {
                          "computed": true,
                          "description": "Configuration for handling paginated responses.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "cursor_configuration": {
                                "computed": true,
                                "description": "Cursor-based pagination configuration.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "limit_parameter": {
                                      "computed": true,
                                      "description": "Parameter extraction configuration.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "default_value": {
                                            "computed": true,
                                            "description": "The default value.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "key": {
                                            "computed": true,
                                            "description": "The parameter key name.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "property_location": {
                                            "computed": true,
                                            "description": "Specifies where to place the parameter in requests.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "value": {
                                            "computed": true,
                                            "description": "Defines how to extract values from HTTP responses.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_path": {
                                                  "computed": true,
                                                  "description": "A JSON path expression to extract a value from response body.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "header_key": {
                                                  "computed": true,
                                                  "description": "The name of an HTTP response header from which to extract the value.",
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
                                    "next_page": {
                                      "computed": true,
                                      "description": "Parameter extraction configuration.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "default_value": {
                                            "computed": true,
                                            "description": "The default value.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "key": {
                                            "computed": true,
                                            "description": "The parameter key name.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "property_location": {
                                            "computed": true,
                                            "description": "Specifies where to place the parameter in requests.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "value": {
                                            "computed": true,
                                            "description": "Defines how to extract values from HTTP responses.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_path": {
                                                  "computed": true,
                                                  "description": "A JSON path expression to extract a value from response body.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "header_key": {
                                                  "computed": true,
                                                  "description": "The name of an HTTP response header from which to extract the value.",
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
                              "offset_configuration": {
                                "computed": true,
                                "description": "Offset-based pagination configuration.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "limit_parameter": {
                                      "computed": true,
                                      "description": "Parameter extraction configuration.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "default_value": {
                                            "computed": true,
                                            "description": "The default value.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "key": {
                                            "computed": true,
                                            "description": "The parameter key name.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "property_location": {
                                            "computed": true,
                                            "description": "Specifies where to place the parameter in requests.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "value": {
                                            "computed": true,
                                            "description": "Defines how to extract values from HTTP responses.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_path": {
                                                  "computed": true,
                                                  "description": "A JSON path expression to extract a value from response body.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "header_key": {
                                                  "computed": true,
                                                  "description": "The name of an HTTP response header from which to extract the value.",
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
                                    "offset_parameter": {
                                      "computed": true,
                                      "description": "Parameter extraction configuration.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "default_value": {
                                            "computed": true,
                                            "description": "The default value.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "key": {
                                            "computed": true,
                                            "description": "The parameter key name.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "property_location": {
                                            "computed": true,
                                            "description": "Specifies where to place the parameter in requests.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "value": {
                                            "computed": true,
                                            "description": "Defines how to extract values from HTTP responses.",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "content_path": {
                                                  "computed": true,
                                                  "description": "A JSON path expression to extract a value from response body.",
                                                  "description_kind": "plain",
                                                  "optional": true,
                                                  "type": "string"
                                                },
                                                "header_key": {
                                                  "computed": true,
                                                  "description": "The name of an HTTP response header from which to extract the value.",
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
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "request_method": {
                          "computed": true,
                          "description": "The HTTP method to use.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "request_parameters": {
                          "computed": true,
                          "description": "Request parameters configuration.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "allowed_values": {
                                "computed": true,
                                "description": "A list of allowed values for the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "default_value": {
                                "computed": true,
                                "description": "The default value for the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "key_override": {
                                "computed": true,
                                "description": "A key name to use when sending this property in API requests.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "name": {
                                "computed": true,
                                "description": "The name of the property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "property_location": {
                                "computed": true,
                                "description": "Specifies where this property should be included in REST requests.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "property_type": {
                                "computed": true,
                                "description": "The data type of this property.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "required": {
                                "computed": true,
                                "description": "Indicates whether the property is required.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "optional": true
                        },
                        "request_path": {
                          "computed": true,
                          "description": "The URL path for the REST endpoint.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "response_configuration": {
                          "computed": true,
                          "description": "Configuration for parsing JSON responses from REST API calls.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "error_path": {
                                "computed": true,
                                "description": "JSON path expression for error information location.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "result_path": {
                                "computed": true,
                                "description": "JSON path expression for result data location.",
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
                "nesting_mode": "map"
              },
              "optional": true
            },
            "global_source_configuration": {
              "computed": true,
              "description": "Configuration that defines how to make requests to endpoints.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "filter_configuration": {
                    "computed": true,
                    "description": "Configuration that defines how filter predicates are applied to REST API requests, supporting both query parameter and filter string strategies.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "between_configuration": {
                          "computed": true,
                          "description": "Configuration that defines how BETWEEN range filter operations are translated into REST API request parameters.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "high_bound_key": {
                                "computed": true,
                                "description": "The parameter name used for the upper bound value in a BETWEEN filter operation.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "low_bound_key": {
                                "computed": true,
                                "description": "The parameter name used for the lower bound value in a BETWEEN filter operation.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "template": {
                                "computed": true,
                                "description": "A template string for constructing the BETWEEN filter expression.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "date_time_format": {
                          "computed": true,
                          "description": "The global date and time format for filter expressions.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "filter_mode": {
                          "computed": true,
                          "description": "The strategy for applying filters to requests.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "filter_string_configuration": {
                          "computed": true,
                          "description": "Configuration for constructing filter expression strings when using the FILTER_STRING filter mode.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "query_parameter_name": {
                                "computed": true,
                                "description": "The query parameter name used to send the constructed filter expression string in API requests.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "quote_character": {
                                "computed": true,
                                "description": "The character used to quote values when QuoteStringValues is true. Defaults to double quotes if not specified.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "quote_string_values": {
                                "computed": true,
                                "description": "Indicates whether string and date values should be wrapped with a quote character in the filter expression.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              }
                            },
                            "nesting_mode": "single"
                          },
                          "optional": true
                        },
                        "operator_mappings": {
                          "computed": true,
                          "description": "A map of logical filter operators to their API-specific string representations.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "strip_quotes": {
                          "computed": true,
                          "description": "Indicates whether surrounding double quotes should be stripped from filter values before processing.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "pagination_configuration": {
                    "computed": true,
                    "description": "Configuration for handling paginated responses.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "cursor_configuration": {
                          "computed": true,
                          "description": "Cursor-based pagination configuration.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "limit_parameter": {
                                "computed": true,
                                "description": "Parameter extraction configuration.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "default_value": {
                                      "computed": true,
                                      "description": "The default value.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "key": {
                                      "computed": true,
                                      "description": "The parameter key name.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "property_location": {
                                      "computed": true,
                                      "description": "Specifies where to place the parameter in requests.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "value": {
                                      "computed": true,
                                      "description": "Defines how to extract values from HTTP responses.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "content_path": {
                                            "computed": true,
                                            "description": "A JSON path expression to extract a value from response body.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "header_key": {
                                            "computed": true,
                                            "description": "The name of an HTTP response header from which to extract the value.",
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
                              "next_page": {
                                "computed": true,
                                "description": "Parameter extraction configuration.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "default_value": {
                                      "computed": true,
                                      "description": "The default value.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "key": {
                                      "computed": true,
                                      "description": "The parameter key name.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "property_location": {
                                      "computed": true,
                                      "description": "Specifies where to place the parameter in requests.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "value": {
                                      "computed": true,
                                      "description": "Defines how to extract values from HTTP responses.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "content_path": {
                                            "computed": true,
                                            "description": "A JSON path expression to extract a value from response body.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "header_key": {
                                            "computed": true,
                                            "description": "The name of an HTTP response header from which to extract the value.",
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
                        "offset_configuration": {
                          "computed": true,
                          "description": "Offset-based pagination configuration.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "limit_parameter": {
                                "computed": true,
                                "description": "Parameter extraction configuration.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "default_value": {
                                      "computed": true,
                                      "description": "The default value.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "key": {
                                      "computed": true,
                                      "description": "The parameter key name.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "property_location": {
                                      "computed": true,
                                      "description": "Specifies where to place the parameter in requests.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "value": {
                                      "computed": true,
                                      "description": "Defines how to extract values from HTTP responses.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "content_path": {
                                            "computed": true,
                                            "description": "A JSON path expression to extract a value from response body.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "header_key": {
                                            "computed": true,
                                            "description": "The name of an HTTP response header from which to extract the value.",
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
                              "offset_parameter": {
                                "computed": true,
                                "description": "Parameter extraction configuration.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "default_value": {
                                      "computed": true,
                                      "description": "The default value.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "key": {
                                      "computed": true,
                                      "description": "The parameter key name.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "property_location": {
                                      "computed": true,
                                      "description": "Specifies where to place the parameter in requests.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "value": {
                                      "computed": true,
                                      "description": "Defines how to extract values from HTTP responses.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "content_path": {
                                            "computed": true,
                                            "description": "A JSON path expression to extract a value from response body.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "header_key": {
                                            "computed": true,
                                            "description": "The name of an HTTP response header from which to extract the value.",
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
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "request_method": {
                    "computed": true,
                    "description": "The HTTP method to use.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "request_parameters": {
                    "computed": true,
                    "description": "Request parameters configuration.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "allowed_values": {
                          "computed": true,
                          "description": "A list of allowed values for the property.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "default_value": {
                          "computed": true,
                          "description": "The default value for the property.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "key_override": {
                          "computed": true,
                          "description": "A key name to use when sending this property in API requests.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "name": {
                          "computed": true,
                          "description": "The name of the property.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "property_location": {
                          "computed": true,
                          "description": "Specifies where this property should be included in REST requests.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "property_type": {
                          "computed": true,
                          "description": "The data type of this property.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "required": {
                          "computed": true,
                          "description": "Indicates whether the property is required.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "optional": true
                  },
                  "request_path": {
                    "computed": true,
                    "description": "The URL path for the REST endpoint.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "response_configuration": {
                    "computed": true,
                    "description": "Configuration for parsing JSON responses from REST API calls.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "error_path": {
                          "computed": true,
                          "description": "JSON path expression for error information location.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "result_path": {
                          "computed": true,
                          "description": "JSON path expression for result data location.",
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
            "validation_endpoint_configuration": {
              "computed": true,
              "description": "Configuration for the validation endpoint. Only supports RequestMethod and RequestPath.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "request_method": {
                    "computed": true,
                    "description": "The HTTP method to use.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "request_path": {
                    "computed": true,
                    "description": "The URL path for the REST endpoint.",
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
        "required": true
      },
      "tags": {
        "computed": true,
        "description": "Tags to assign to the connection type.",
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
      }
    },
    "description": "Registers a custom connection type in Glue based on the configuration provided. This enables customers to configure custom connectors for any data source with REST-based APIs.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccGlueConnectionTypeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGlueConnectionType), &result)
	return &result
}
