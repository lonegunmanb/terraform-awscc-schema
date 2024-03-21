package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSecretsmanagerSecret = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "The description of the secret.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "generate_secret_string": {
        "computed": true,
        "description": "A structure that specifies how to generate a password to encrypt and store in the secret. To include a specific string in the secret, use ` + "`" + `` + "`" + `SecretString` + "`" + `` + "`" + ` instead. If you omit both ` + "`" + `` + "`" + `GenerateSecretString` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `SecretString` + "`" + `` + "`" + `, you create an empty secret. When you make a change to this property, a new secret version is created.\n We recommend that you specify the maximum length and include every character type that the system you are generating a password for can support.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "exclude_characters": {
              "computed": true,
              "description": "A string of the characters that you don't want in the password.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "exclude_lowercase": {
              "computed": true,
              "description": "Specifies whether to exclude lowercase letters from the password. If you don't include this switch, the password can contain lowercase letters.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "exclude_numbers": {
              "computed": true,
              "description": "Specifies whether to exclude numbers from the password. If you don't include this switch, the password can contain numbers.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "exclude_punctuation": {
              "computed": true,
              "description": "Specifies whether to exclude the following punctuation characters from the password: ` + "`" + `` + "`" + `! \" # $ % \u0026 ' ( ) * + , - . / : ; \u003c = \u003e ? @ [ \\ ] ^ _ ` + "`" + ` { | } ~` + "`" + `` + "`" + `. If you don't include this switch, the password can contain punctuation.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "exclude_uppercase": {
              "computed": true,
              "description": "Specifies whether to exclude uppercase letters from the password. If you don't include this switch, the password can contain uppercase letters.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "generate_string_key": {
              "computed": true,
              "description": "The JSON key name for the key/value pair, where the value is the generated password. This pair is added to the JSON structure specified by the ` + "`" + `` + "`" + `SecretStringTemplate` + "`" + `` + "`" + ` parameter. If you specify this parameter, then you must also specify ` + "`" + `` + "`" + `SecretStringTemplate` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "include_space": {
              "computed": true,
              "description": "Specifies whether to include the space character. If you include this switch, the password can contain space characters.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "password_length": {
              "computed": true,
              "description": "The length of the password. If you don't include this parameter, the default length is 32 characters.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "require_each_included_type": {
              "computed": true,
              "description": "Specifies whether to include at least one upper and lowercase letter, one number, and one punctuation. If you don't include this switch, the password contains at least one of every character type.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "secret_string_template": {
              "computed": true,
              "description": "A template that the generated string must match. When you make a change to this property, a new secret version is created.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description": "The ARN, key ID, or alias of the KMS key that Secrets Manager uses to encrypt the secret value in the secret. An alias is always prefixed by ` + "`" + `` + "`" + `alias/` + "`" + `` + "`" + `, for example ` + "`" + `` + "`" + `alias/aws/secretsmanager` + "`" + `` + "`" + `. For more information, see [About aliases](https://docs.aws.amazon.com/kms/latest/developerguide/alias-about.html).\n To use a KMS key in a different account, use the key ARN or the alias ARN.\n If you don't specify this value, then Secrets Manager uses the key ` + "`" + `` + "`" + `aws/secretsmanager` + "`" + `` + "`" + `. If that key doesn't yet exist, then Secrets Manager creates it for you automatically the first time it encrypts the secret value.\n If the secret is in a different AWS account from the credentials calling the API, then you can't use ` + "`" + `` + "`" + `aws/secretsmanager` + "`" + `` + "`" + ` to encrypt the secret, and you must create and use a customer managed KMS key.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the new secret.\n The secret name can contain ASCII letters, numbers, and the following characters: /_+=.@-\n Do not end your secret name with a hyphen followed by six characters. If you do so, you risk confusion and unexpected results when searching for a secret by partial ARN. Secrets Manager automatically adds a hyphen and six random characters after the secret name at the end of the ARN.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "replica_regions": {
        "computed": true,
        "description": "A custom type that specifies a ` + "`" + `` + "`" + `Region` + "`" + `` + "`" + ` and the ` + "`" + `` + "`" + `KmsKeyId` + "`" + `` + "`" + ` for a replica secret.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kms_key_id": {
              "computed": true,
              "description": "The ARN, key ID, or alias of the KMS key to encrypt the secret. If you don't include this field, Secrets Manager uses ` + "`" + `` + "`" + `aws/secretsmanager` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "region": {
              "description": "A string that represents a ` + "`" + `` + "`" + `Region` + "`" + `` + "`" + `, for example \"us-east-1\".",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "secret_string": {
        "computed": true,
        "description": "The text to encrypt and store in the secret. We recommend you use a JSON structure of key/value pairs for your secret value. To generate a random password, use ` + "`" + `` + "`" + `GenerateSecretString` + "`" + `` + "`" + ` instead. If you omit both ` + "`" + `` + "`" + `GenerateSecretString` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `SecretString` + "`" + `` + "`" + `, you create an empty secret. When you make a change to this property, a new secret version is created.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of tags to attach to the secret. Each tag is a key and value pair of strings in a JSON text string, for example:\n  ` + "`" + `` + "`" + `[{\"Key\":\"CostCenter\",\"Value\":\"12345\"},{\"Key\":\"environment\",\"Value\":\"production\"}]` + "`" + `` + "`" + ` \n Secrets Manager tag key names are case sensitive. A tag with the key \"ABC\" is a different tag from one with key \"abc\".\n Stack-level tags, tags you apply to the CloudFormation stack, are also attached to the secret. \n If you check tags in permissions policies as part of your security strategy, then adding or removing a tag can change permissions. If the completion of this operation would result in you losing your permissions for this secret, then Secrets Manager blocks the operation and returns an ` + "`" + `` + "`" + `Access Denied` + "`" + `` + "`" + ` error. For more information, see [Control access to secrets using tags](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_examples.html#tag-secrets-abac) and [Limit access to identities with tags that match secrets' tags](https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_examples.html#auth-and-access_tags2).\n For information about how to format a JSON parameter for the various command line tool environments, see [Using JSON for Parameters](https://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#cli-using-param-json). If your command-line tool or SDK requires quotation marks around the parameter, you should use single quotes to avoid confusion with the double quotes required in the JSON text.\n The following restrictions apply to tags:\n  +  Maximum number of tags per secret: 50\n  +  Maximum key length: 127 Unicode characters in UTF-8\n  +  Maximum value length: 255 Unicode characters in UTF-8\n  +  Tag keys and values are case sensitive.\n  +  Do not use the ` + "`" + `` + "`" + `aws:` + "`" + `` + "`" + ` prefix in your tag names or values because AWS reserves it for AWS use. You can't edit or delete tag names or values with this prefix. Tags with this prefix do not count against your tags per secret limit.\n  +  If you use your tagging schema across multiple services and resources, other services might have restrictions on allowed characters. Generally allowed characters: letters, spaces, and numbers representable in UTF-8, plus the following special characters: + - = . _ : / @.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key identifier, or name, of the tag.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The string value associated with the key of the tag.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Creates a new secret. A *secret* can be a password, a set of credentials such as a user name and password, an OAuth token, or other secret information that you store in an encrypted form in Secrets Manager.\n For RDS master user credentials, see [AWS::RDS::DBCluster MasterUserSecret](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbcluster-masterusersecret.html).\n To retrieve a secret in a CFNshort template, use a *dynamic reference*. For more information, see [Retrieve a secret in an resource](https://docs.aws.amazon.com/secretsmanager/latest/userguide/cfn-example_reference-secret.html).\n A common scenario is to first create a secret with ` + "`" + `` + "`" + `GenerateSecretString` + "`" + `` + "`" + `, which generates a password, and then use a dynamic reference to retrieve the username and password from the secret to use as credentials for a new database. See the example *Creating a Redshift cluster and a secret for the admin credentials*.\n For information about creating a secret in the console, see [Create a secret](https://docs.aws.amazon.com/secretsmanager/latest/userguide/manage_create-basic-secret.html). For information about creating a secret using the CLI or SDK, see [CreateSecret](https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_CreateSecret.html).\n For information about retrieving a secret in code, see [Retrieve secrets from Secrets Manager](https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieving-secrets.html).",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSecretsmanagerSecretSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecretsmanagerSecret), &result)
	return &result
}
