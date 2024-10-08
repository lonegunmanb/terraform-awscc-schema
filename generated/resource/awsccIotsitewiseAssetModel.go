package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotsitewiseAssetModel = `{
  "block": {
    "attributes": {
      "asset_model_arn": {
        "computed": true,
        "description": "The ARN of the asset model, which has the following format.",
        "description_kind": "plain",
        "type": "string"
      },
      "asset_model_composite_models": {
        "computed": true,
        "description": "The composite asset models that are part of this asset model. Composite asset models are asset models that contain specific properties.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "composed_asset_model_id": {
              "computed": true,
              "description": "The component model ID for which the composite model is composed of",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "composite_model_properties": {
              "computed": true,
              "description": "The property definitions of the asset model. You can specify up to 200 properties per asset model.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "data_type": {
                    "computed": true,
                    "description": "The data type of the asset model property.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "data_type_spec": {
                    "computed": true,
                    "description": "The data type of the structure for this property.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "external_id": {
                    "computed": true,
                    "description": "The External ID of the Asset Model Property",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description": "The ID of the Asset Model Property",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "logical_id": {
                    "computed": true,
                    "description": "Customer provided Logical ID for property.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "The name of the asset model property.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "type": {
                    "computed": true,
                    "description": "The property type",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "attribute": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "default_value": {
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
                        "metric": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "expression": {
                                "computed": true,
                                "description": "The mathematical expression that defines the metric aggregation function. You can specify up to 10 functions per expression.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "variables": {
                                "computed": true,
                                "description": "The list of variables used in the expression.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "name": {
                                      "computed": true,
                                      "description": "The friendly name of the variable to be used in the expression.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "value": {
                                      "computed": true,
                                      "description": "The variable that identifies an asset property from which to use values.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "hierarchy_external_id": {
                                            "computed": true,
                                            "description": "The External ID of the hierarchy that is trying to be referenced",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "hierarchy_id": {
                                            "computed": true,
                                            "description": "The ID of the hierarchy that is trying to be referenced",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "hierarchy_logical_id": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "property_external_id": {
                                            "computed": true,
                                            "description": "The External ID of the property that is trying to be referenced",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "property_id": {
                                            "computed": true,
                                            "description": "The ID of the property that is trying to be referenced",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "property_logical_id": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "property_path": {
                                            "computed": true,
                                            "description": "The path of the property that is trying to be referenced",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "name": {
                                                  "computed": true,
                                                  "description": "The name of the property",
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
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
                                    }
                                  },
                                  "nesting_mode": "list"
                                },
                                "optional": true
                              },
                              "window": {
                                "computed": true,
                                "description": "The window (time interval) over which AWS IoT SiteWise computes the metric's aggregation expression",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "tumbling": {
                                      "computed": true,
                                      "description": "Contains a tumbling window, which is a repeating fixed-sized, non-overlapping, and contiguous time interval. This window is used in metric and aggregation computations.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "interval": {
                                            "computed": true,
                                            "description": "The time interval for the tumbling window.",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "offset": {
                                            "computed": true,
                                            "description": "The shift or reference point on timeline for the contiguous time intervals.",
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
                        "transform": {
                          "computed": true,
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "expression": {
                                "computed": true,
                                "description": "The mathematical expression that defines the transformation function. You can specify up to 10 functions per expression.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "variables": {
                                "computed": true,
                                "description": "The list of variables used in the expression.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "name": {
                                      "computed": true,
                                      "description": "The friendly name of the variable to be used in the expression.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "value": {
                                      "computed": true,
                                      "description": "The variable that identifies an asset property from which to use values.",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "hierarchy_external_id": {
                                            "computed": true,
                                            "description": "The External ID of the hierarchy that is trying to be referenced",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "hierarchy_id": {
                                            "computed": true,
                                            "description": "The ID of the hierarchy that is trying to be referenced",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "hierarchy_logical_id": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "property_external_id": {
                                            "computed": true,
                                            "description": "The External ID of the property that is trying to be referenced",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "property_id": {
                                            "computed": true,
                                            "description": "The ID of the property that is trying to be referenced",
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "property_logical_id": {
                                            "computed": true,
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "property_path": {
                                            "computed": true,
                                            "description": "The path of the property that is trying to be referenced",
                                            "description_kind": "plain",
                                            "nested_type": {
                                              "attributes": {
                                                "name": {
                                                  "computed": true,
                                                  "description": "The name of the property",
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
                                        "nesting_mode": "single"
                                      },
                                      "optional": true
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
                        "type_name": {
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
                  "unit": {
                    "computed": true,
                    "description": "The unit of the asset model property, such as Newtons or RPM.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "description": {
              "computed": true,
              "description": "A description for the asset composite model.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "external_id": {
              "computed": true,
              "description": "The External ID of the composite model",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "id": {
              "computed": true,
              "description": "The Actual ID of the composite model",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "A unique, friendly name for the asset composite model.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "parent_asset_model_composite_model_external_id": {
              "computed": true,
              "description": "The parent composite model External ID",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "path": {
              "computed": true,
              "description": "The path of the composite model. This is only for derived composite models",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "type": {
              "computed": true,
              "description": "The type of the composite model. For alarm composite models, this type is AWS/ALARM",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "asset_model_description": {
        "computed": true,
        "description": "A description for the asset model.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "asset_model_external_id": {
        "computed": true,
        "description": "The external ID of the asset model.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "asset_model_hierarchies": {
        "computed": true,
        "description": "The hierarchy definitions of the asset model. Each hierarchy specifies an asset model whose assets can be children of any other assets created from this asset model. You can specify up to 10 hierarchies per asset model.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "child_asset_model_id": {
              "computed": true,
              "description": "The ID of the asset model. All assets in this hierarchy must be instances of the child AssetModelId asset model.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "external_id": {
              "computed": true,
              "description": "Customer provided external ID for hierarchy",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "id": {
              "computed": true,
              "description": "Customer provided actual ID for hierarchy",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "logical_id": {
              "computed": true,
              "description": "Customer provided logical ID for hierarchy.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "The name of the asset model hierarchy.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "asset_model_id": {
        "computed": true,
        "description": "The ID of the asset model.",
        "description_kind": "plain",
        "type": "string"
      },
      "asset_model_name": {
        "description": "A unique, friendly name for the asset model.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "asset_model_properties": {
        "computed": true,
        "description": "The property definitions of the asset model. You can specify up to 200 properties per asset model.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "data_type": {
              "computed": true,
              "description": "The data type of the asset model property.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "data_type_spec": {
              "computed": true,
              "description": "The data type of the structure for this property.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "external_id": {
              "computed": true,
              "description": "The External ID of the Asset Model Property",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "id": {
              "computed": true,
              "description": "The ID of the Asset Model Property",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "logical_id": {
              "computed": true,
              "description": "Customer provided Logical ID for property.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "The name of the asset model property.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "computed": true,
              "description": "The property type",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "attribute": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "default_value": {
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
                  "metric": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "expression": {
                          "computed": true,
                          "description": "The mathematical expression that defines the metric aggregation function. You can specify up to 10 functions per expression.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "variables": {
                          "computed": true,
                          "description": "The list of variables used in the expression.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "name": {
                                "computed": true,
                                "description": "The friendly name of the variable to be used in the expression.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "value": {
                                "computed": true,
                                "description": "The variable that identifies an asset property from which to use values.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "hierarchy_external_id": {
                                      "computed": true,
                                      "description": "The External ID of the hierarchy that is trying to be referenced",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "hierarchy_id": {
                                      "computed": true,
                                      "description": "The ID of the hierarchy that is trying to be referenced",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "hierarchy_logical_id": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "property_external_id": {
                                      "computed": true,
                                      "description": "The External ID of the property that is trying to be referenced",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "property_id": {
                                      "computed": true,
                                      "description": "The ID of the property that is trying to be referenced",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "property_logical_id": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "property_path": {
                                      "computed": true,
                                      "description": "The path of the property that is trying to be referenced",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "name": {
                                            "computed": true,
                                            "description": "The name of the property",
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
                                  "nesting_mode": "single"
                                },
                                "optional": true
                              }
                            },
                            "nesting_mode": "list"
                          },
                          "optional": true
                        },
                        "window": {
                          "computed": true,
                          "description": "The window (time interval) over which AWS IoT SiteWise computes the metric's aggregation expression",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "tumbling": {
                                "computed": true,
                                "description": "Contains a tumbling window, which is a repeating fixed-sized, non-overlapping, and contiguous time interval. This window is used in metric and aggregation computations.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "interval": {
                                      "computed": true,
                                      "description": "The time interval for the tumbling window.",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "offset": {
                                      "computed": true,
                                      "description": "The shift or reference point on timeline for the contiguous time intervals.",
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
                  "transform": {
                    "computed": true,
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "expression": {
                          "computed": true,
                          "description": "The mathematical expression that defines the transformation function. You can specify up to 10 functions per expression.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "variables": {
                          "computed": true,
                          "description": "The list of variables used in the expression.",
                          "description_kind": "plain",
                          "nested_type": {
                            "attributes": {
                              "name": {
                                "computed": true,
                                "description": "The friendly name of the variable to be used in the expression.",
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "value": {
                                "computed": true,
                                "description": "The variable that identifies an asset property from which to use values.",
                                "description_kind": "plain",
                                "nested_type": {
                                  "attributes": {
                                    "hierarchy_external_id": {
                                      "computed": true,
                                      "description": "The External ID of the hierarchy that is trying to be referenced",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "hierarchy_id": {
                                      "computed": true,
                                      "description": "The ID of the hierarchy that is trying to be referenced",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "hierarchy_logical_id": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "property_external_id": {
                                      "computed": true,
                                      "description": "The External ID of the property that is trying to be referenced",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "property_id": {
                                      "computed": true,
                                      "description": "The ID of the property that is trying to be referenced",
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "property_logical_id": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "property_path": {
                                      "computed": true,
                                      "description": "The path of the property that is trying to be referenced",
                                      "description_kind": "plain",
                                      "nested_type": {
                                        "attributes": {
                                          "name": {
                                            "computed": true,
                                            "description": "The name of the property",
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
                                  "nesting_mode": "single"
                                },
                                "optional": true
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
                  "type_name": {
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
            "unit": {
              "computed": true,
              "description": "The unit of the asset model property, such as Newtons or RPM.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "asset_model_type": {
        "computed": true,
        "description": "The type of the asset model (ASSET_MODEL OR COMPONENT_MODEL)",
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
      "tags": {
        "computed": true,
        "description": "A list of key-value pairs that contain metadata for the asset model.",
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
    "description": "Resource schema for AWS::IoTSiteWise::AssetModel",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIotsitewiseAssetModelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotsitewiseAssetModel), &result)
	return &result
}
