package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccKmsKey = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "bypass_policy_lockout_safety_check": {
        "computed": true,
        "description": "Skips (\"bypasses\") the key policy lockout safety check. The default value is false.\n  Setting this value to true increases the risk that the KMS key becomes unmanageable. Do not set this value to true indiscriminately.\n For more information, see [Default key policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policy-default.html#prevent-unmanageable-key) in the *Developer Guide*.\n  Use this parameter only when you intend to prevent the principal that is making the request from making a subsequent [PutKeyPolicy](https://docs.aws.amazon.com/kms/latest/APIReference/API_PutKeyPolicy.html) request on the KMS key.",
        "description_kind": "plain",
        "type": "bool"
      },
      "description": {
        "computed": true,
        "description": "A description of the KMS key. Use a description that helps you to distinguish this KMS key from others in the account, such as its intended use.",
        "description_kind": "plain",
        "type": "string"
      },
      "enable_key_rotation": {
        "computed": true,
        "description": "Enables automatic rotation of the key material for the specified KMS key. By default, automatic key rotation is not enabled.\n  KMS supports automatic rotation only for symmetric encryption KMS keys (` + "`" + `` + "`" + `KeySpec` + "`" + `` + "`" + ` = ` + "`" + `` + "`" + `SYMMETRIC_DEFAULT` + "`" + `` + "`" + `). For asymmetric KMS keys, HMAC KMS keys, and KMS keys with Origin ` + "`" + `` + "`" + `EXTERNAL` + "`" + `` + "`" + `, omit the ` + "`" + `` + "`" + `EnableKeyRotation` + "`" + `` + "`" + ` property or set it to ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `.\n To enable automatic key rotation of the key material for a multi-Region KMS key, set ` + "`" + `` + "`" + `EnableKeyRotation` + "`" + `` + "`" + ` to ` + "`" + `` + "`" + `true` + "`" + `` + "`" + ` on the primary key (created by using ` + "`" + `` + "`" + `AWS::KMS::Key` + "`" + `` + "`" + `). KMS copies the rotation status to all replica keys. For details, see [Rotating multi-Region keys](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-manage.html#multi-region-rotate) in the *Developer Guide*.\n When you enable automatic rotation, KMS automatically creates new key material for the KMS key one year after the enable date and every year thereafter. KMS retains all key material until you delete the KMS key. For detailed information about automatic key rotation, see [Rotating KMS keys](https://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html) in the *Developer Guide*.",
        "description_kind": "plain",
        "type": "bool"
      },
      "enabled": {
        "computed": true,
        "description": "Specifies whether the KMS key is enabled. Disabled KMS keys cannot be used in cryptographic operations.\n When ` + "`" + `` + "`" + `Enabled` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `, the *key state* of the KMS key is ` + "`" + `` + "`" + `Enabled` + "`" + `` + "`" + `. When ` + "`" + `` + "`" + `Enabled` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `, the key state of the KMS key is ` + "`" + `` + "`" + `Disabled` + "`" + `` + "`" + `. The default value is ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `.\n The actual key state of the KMS key might be affected by actions taken outside of CloudFormation, such as running the [EnableKey](https://docs.aws.amazon.com/kms/latest/APIReference/API_EnableKey.html), [DisableKey](https://docs.aws.amazon.com/kms/latest/APIReference/API_DisableKey.html), or [ScheduleKeyDeletion](https://docs.aws.amazon.com/kms/latest/APIReference/API_ScheduleKeyDeletion.html) operations.\n For information about the key states of a KMS key, see [Key state: Effect on your KMS key](https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the *Developer Guide*.",
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "key_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "key_policy": {
        "computed": true,
        "description": "The key policy to attach to the KMS key.\n If you provide a key policy, it must meet the following criteria:\n  +  The key policy must allow the caller to make a subsequent [PutKeyPolicy](https://docs.aws.amazon.com/kms/latest/APIReference/API_PutKeyPolicy.html) request on the KMS key. This reduces the risk that the KMS key becomes unmanageable. For more information, see [Default key policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam) in the *Developer Guide*. (To omit this condition, set ` + "`" + `` + "`" + `BypassPolicyLockoutSafetyCheck` + "`" + `` + "`" + ` to true.)\n  +  Each statement in the key policy must contain one or more principals. The principals in the key policy must exist and be visible to KMS. When you create a new AWS principal (for example, an IAM user or role), you might need to enforce a delay before including the new principal in a key policy because the new principal might not be immediately visible to KMS. For more information, see [Changes that I make are not always immediately visible](https://docs.aws.amazon.com/IAM/latest/UserGuide/troubleshoot_general.html#troubleshoot_general_eventual-consistency) in the *User Guide*.\n  \n If you do not provide a key policy, KMS attaches a default key policy to the KMS key. For more information, see [Default key policy](https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default) in the *Developer Guide*.\n A key policy document can include only the following characters:\n  +  Printable ASCII characters\n  +  Printable characters in the Basic Latin and Latin-1 Supplement character set\n  +  The tab (` + "`" + `` + "`" + `\\u0009` + "`" + `` + "`" + `), line feed (` + "`" + `` + "`" + `\\u000A` + "`" + `` + "`" + `), and carriage return (` + "`" + `` + "`" + `\\u000D` + "`" + `` + "`" + `) special characters\n  \n  *Minimum*: ` + "`" + `` + "`" + `1` + "`" + `` + "`" + ` \n  *Maximum*: ` + "`" + `` + "`" + `32768` + "`" + `` + "`" + `",
        "description_kind": "plain",
        "type": "string"
      },
      "key_spec": {
        "computed": true,
        "description": "Specifies the type of KMS key to create. The default value, ` + "`" + `` + "`" + `SYMMETRIC_DEFAULT` + "`" + `` + "`" + `, creates a KMS key with a 256-bit symmetric key for encryption and decryption. In China Regions, ` + "`" + `` + "`" + `SYMMETRIC_DEFAULT` + "`" + `` + "`" + ` creates a 128-bit symmetric key that uses SM4 encryption. You can't change the ` + "`" + `` + "`" + `KeySpec` + "`" + `` + "`" + ` value after the KMS key is created. For help choosing a key spec for your KMS key, see [Choosing a KMS key type](https://docs.aws.amazon.com/kms/latest/developerguide/symm-asymm-choose.html) in the *Developer Guide*.\n The ` + "`" + `` + "`" + `KeySpec` + "`" + `` + "`" + ` property determines the type of key material in the KMS key and the algorithms that the KMS key supports. To further restrict the algorithms that can be used with the KMS key, use a condition key in its key policy or IAM policy. For more information, see [condition keys](https://docs.aws.amazon.com/kms/latest/developerguide/policy-conditions.html#conditions-kms) in the *Developer Guide*.\n  If you change the value of the ` + "`" + `` + "`" + `KeySpec` + "`" + `` + "`" + ` property on an existing KMS key, the update request fails, regardless of the value of the [UpdateReplacePolicy attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatereplacepolicy.html). This prevents you from accidentally deleting a KMS key by changing an immutable property value.\n    [services that are integrated with](https://docs.aws.amazon.com/kms/features/#AWS_Service_Integration) use symmetric encryption KMS keys to protect your data. These services do not support encryption with asymmetric KMS keys. For help determining whether a KMS key is asymmetric, see [Identifying asymmetric KMS keys](https://docs.aws.amazon.com/kms/latest/developerguide/find-symm-asymm.html) in the *Developer Guide*.\n   KMS supports the following key specs for KMS keys:\n  +  Symmetric encryption key (default)\n  +   ` + "`" + `` + "`" + `SYMMETRIC_DEFAULT` + "`" + `` + "`" + ` (AES-256-GCM)\n  \n  +  HMAC keys (symmetric)\n  +   ` + "`" + `` + "`" + `HMAC_224` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `HMAC_256` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `HMAC_384` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `HMAC_512` + "`" + `` + "`" + ` \n  \n  +  Asymmetric RSA key pairs\n  +   ` + "`" + `` + "`" + `RSA_2048` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `RSA_3072` + "`" + `` + "`" + ` \n  +   ` + "`" + `` + "`" + `RSA_4096` + "`" + `` + "`" + ` \n  \n  +  Asymmetric NIST-recommended elliptic curve key pairs\n  +   ` + "`" + `` + "`" + `ECC_NIST_P256` + "`" + `` + "`" + ` (secp256r1)\n  +   ` + "`" + `` + "`" + `ECC_NIST_P384` + "`" + `` + "`" + ` (secp384r1)\n  +   ` + "`" + `` + "`" + `ECC_NIST_P521` + "`" + `` + "`" + ` (secp521r1)\n  \n  +  Other asymmetric elliptic curve key pairs\n  +   ` + "`" + `` + "`" + `ECC_SECG_P256K1` + "`" + `` + "`" + ` (secp256k1), commonly used for cryptocurrencies.\n  \n  +  SM2 key pairs (China Regions only)\n  +   ` + "`" + `` + "`" + `SM2` + "`" + `` + "`" + `",
        "description_kind": "plain",
        "type": "string"
      },
      "key_usage": {
        "computed": true,
        "description": "Determines the [cryptographic operations](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations) for which you can use the KMS key. The default value is ` + "`" + `` + "`" + `ENCRYPT_DECRYPT` + "`" + `` + "`" + `. This property is required for asymmetric KMS keys and HMAC KMS keys. You can't change the ` + "`" + `` + "`" + `KeyUsage` + "`" + `` + "`" + ` value after the KMS key is created.\n  If you change the value of the ` + "`" + `` + "`" + `KeyUsage` + "`" + `` + "`" + ` property on an existing KMS key, the update request fails, regardless of the value of the [UpdateReplacePolicy attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatereplacepolicy.html). This prevents you from accidentally deleting a KMS key by changing an immutable property value.\n  Select only one valid value.\n  +  For symmetric encryption KMS keys, omit the property or specify ` + "`" + `` + "`" + `ENCRYPT_DECRYPT` + "`" + `` + "`" + `.\n  +  For asymmetric KMS keys with RSA key material, specify ` + "`" + `` + "`" + `ENCRYPT_DECRYPT` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `SIGN_VERIFY` + "`" + `` + "`" + `.\n  +  For asymmetric KMS keys with ECC key material, specify ` + "`" + `` + "`" + `SIGN_VERIFY` + "`" + `` + "`" + `.\n  +  For asymmetric KMS keys with SM2 (China Regions only) key material, specify ` + "`" + `` + "`" + `ENCRYPT_DECRYPT` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `SIGN_VERIFY` + "`" + `` + "`" + `.\n  +  For HMAC KMS keys, specify ` + "`" + `` + "`" + `GENERATE_VERIFY_MAC` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "type": "string"
      },
      "multi_region": {
        "computed": true,
        "description": "Creates a multi-Region primary key that you can replicate in other AWS-Regions. You can't change the ` + "`" + `` + "`" + `MultiRegion` + "`" + `` + "`" + ` value after the KMS key is created.\n For a list of AWS-Regions in which multi-Region keys are supported, see [Multi-Region keys in](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html) in the **.\n  If you change the value of the ` + "`" + `` + "`" + `MultiRegion` + "`" + `` + "`" + ` property on an existing KMS key, the update request fails, regardless of the value of the [UpdateReplacePolicy attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatereplacepolicy.html). This prevents you from accidentally deleting a KMS key by changing an immutable property value.\n  For a multi-Region key, set to this property to ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `. For a single-Region key, omit this property or set it to ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `. The default value is ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `.\n  *Multi-Region keys* are an KMS feature that lets you create multiple interoperable KMS keys in different AWS-Regions. Because these KMS keys have the same key ID, key material, and other metadata, you can use them to encrypt data in one AWS-Region and decrypt it in a different AWS-Region without making a cross-Region call or exposing the plaintext data. For more information, see [Multi-Region keys](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html) in the *Developer Guide*.\n You can create a symmetric encryption, HMAC, or asymmetric multi-Region KMS key, and you can create a multi-Region key with imported key material. However, you cannot create a multi-Region key in a custom key store.\n To create a replica of this primary key in a different AWS-Region , create an [AWS::KMS::ReplicaKey](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-replicakey.html) resource in a CloudFormation stack in the replica Region. Specify the key ARN of this primary key.",
        "description_kind": "plain",
        "type": "bool"
      },
      "origin": {
        "computed": true,
        "description": "The source of the key material for the KMS key. You cannot change the origin after you create the KMS key. The default is ` + "`" + `` + "`" + `AWS_KMS` + "`" + `` + "`" + `, which means that KMS creates the key material.\n To [create a KMS key with no key material](https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys-create-cmk.html) (for imported key material), set this value to ` + "`" + `` + "`" + `EXTERNAL` + "`" + `` + "`" + `. For more information about importing key material into KMS, see [Importing Key Material](https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html) in the *Developer Guide*.\n You can ignore ` + "`" + `` + "`" + `ENABLED` + "`" + `` + "`" + ` when Origin is ` + "`" + `` + "`" + `EXTERNAL` + "`" + `` + "`" + `. When a KMS key with Origin ` + "`" + `` + "`" + `EXTERNAL` + "`" + `` + "`" + ` is created, the key state is ` + "`" + `` + "`" + `PENDING_IMPORT` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `ENABLED` + "`" + `` + "`" + ` is ` + "`" + `` + "`" + `false` + "`" + `` + "`" + `. After you import the key material, ` + "`" + `` + "`" + `ENABLED` + "`" + `` + "`" + ` updated to ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `. The KMS key can then be used for Cryptographic Operations. \n   CFN doesn't support creating an ` + "`" + `` + "`" + `Origin` + "`" + `` + "`" + ` parameter of the ` + "`" + `` + "`" + `AWS_CLOUDHSM` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `EXTERNAL_KEY_STORE` + "`" + `` + "`" + ` values.",
        "description_kind": "plain",
        "type": "string"
      },
      "pending_window_in_days": {
        "computed": true,
        "description": "Specifies the number of days in the waiting period before KMS deletes a KMS key that has been removed from a CloudFormation stack. Enter a value between 7 and 30 days. The default value is 30 days.\n When you remove a KMS key from a CloudFormation stack, KMS schedules the KMS key for deletion and starts the mandatory waiting period. The ` + "`" + `` + "`" + `PendingWindowInDays` + "`" + `` + "`" + ` property determines the length of waiting period. During the waiting period, the key state of KMS key is ` + "`" + `` + "`" + `Pending Deletion` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `Pending Replica Deletion` + "`" + `` + "`" + `, which prevents the KMS key from being used in cryptographic operations. When the waiting period expires, KMS permanently deletes the KMS key.\n  KMS will not delete a [multi-Region primary key](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html) that has replica keys. If you remove a multi-Region primary key from a CloudFormation stack, its key state changes to ` + "`" + `` + "`" + `PendingReplicaDeletion` + "`" + `` + "`" + ` so it cannot be replicated or used in cryptographic operations. This state can persist indefinitely. When the last of its replica keys is deleted, the key state of the primary key changes to ` + "`" + `` + "`" + `PendingDeletion` + "`" + `` + "`" + ` and the waiting period specified by ` + "`" + `` + "`" + `PendingWindowInDays` + "`" + `` + "`" + ` begins. When this waiting period expires, KMS deletes the primary key. For details, see [Deleting multi-Region keys](https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-delete.html) in the *Developer Guide*.\n You cannot use a CloudFormation template to cancel deletion of the KMS key after you remove it from the stack, regardless of the waiting period. If you specify a KMS key in your template, even one with the same name, CloudFormation creates a new KMS key. To cancel deletion of a KMS key, use the KMS console or the [CancelKeyDeletion](https://docs.aws.amazon.com/kms/latest/APIReference/API_CancelKeyDeletion.html) operation.\n For information about the ` + "`" + `` + "`" + `Pending Deletion` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `Pending Replica Deletion` + "`" + `` + "`" + ` key states, see [Key state: Effect on your KMS key](https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the *Developer Guide*. For more information about deleting KMS keys, see the [ScheduleKeyDeletion](https://docs.aws.amazon.com/kms/latest/APIReference/API_ScheduleKeyDeletion.html) operation in the *API Reference* and [Deleting KMS keys](https://docs.aws.amazon.com/kms/latest/developerguide/deleting-keys.html) in the *Developer Guide*.",
        "description_kind": "plain",
        "type": "number"
      },
      "rotation_period_in_days": {
        "computed": true,
        "description": "Specifies a custom period of time between each rotation date. If no value is specified, the default value is 365 days.\n The rotation period defines the number of days after you enable automatic key rotation that KMS will rotate your key material, and the number of days between each automatic rotation thereafter.\n You can use the [kms:RotationPeriodInDays](https://docs.aws.amazon.com/kms/latest/developerguide/conditions-kms.html#conditions-kms-rotation-period-in-days) condition key to further constrain the values that principals can specify in the ` + "`" + `` + "`" + `RotationPeriodInDays` + "`" + `` + "`" + ` parameter.\n For more information about rotating KMS keys and automatic rotation, see [Rotating keys](https://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html) in the *Developer Guide*.",
        "description_kind": "plain",
        "type": "number"
      },
      "tags": {
        "computed": true,
        "description": "Assigns one or more tags to the replica key.\n  Tagging or untagging a KMS key can allow or deny permission to the KMS key. For details, see [ABAC for](https://docs.aws.amazon.com/kms/latest/developerguide/abac.html) in the *Developer Guide*.\n  For information about tags in KMS, see [Tagging keys](https://docs.aws.amazon.com/kms/latest/developerguide/tagging-keys.html) in the *Developer Guide*. For information about tags in CloudFormation, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::KMS::Key",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccKmsKeySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccKmsKey), &result)
	return &result
}
