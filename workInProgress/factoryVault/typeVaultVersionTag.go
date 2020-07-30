package factoryVault

type VaultVersionTag int

const (
	KVaultVersionTag_latest VaultVersionTag = iota
	KVaultVersionTag_1_5_0
	KVaultVersionTag_1_5_0_rc
	KVaultVersionTag_1_4_3
	KVaultVersionTag_1_3_7
	KVaultVersionTag_1_4_2
	KVaultVersionTag_1_3_6
	KVaultVersionTag_1_4_1
	KVaultVersionTag_1_3_5
	KVaultVersionTag_1_4_0
	KVaultVersionTag_1_4_0_rc1
	KVaultVersionTag_1_3_4
	KVaultVersionTag_1_3_3
	KVaultVersionTag_1_4_0_beta1
	KVaultVersionTag_1_3_2
	KVaultVersionTag_1_3_1
	KVaultVersionTag_1_3_0
	KVaultVersionTag_1_2_4
	KVaultVersionTag_1_3_0_beta1
	KVaultVersionTag_1_2_3
	KVaultVersionTag_1_2_2
	KVaultVersionTag_1_2_1
	KVaultVersionTag_1_2_0
	KVaultVersionTag_1_1_5
	KVaultVersionTag_1_2_0_beta2
	KVaultVersionTag_1_2_0_beta1
	KVaultVersionTag_1_1_3
	KVaultVersionTag_1_1_2
	KVaultVersionTag_1_1_1
	KVaultVersionTag_1_1_0
	KVaultVersionTag_1_1_0_beta2
	KVaultVersionTag_1_0_3
	KVaultVersionTag_0_11_6
	KVaultVersionTag_1_0_2
	KVaultVersionTag_1_0_1
	KVaultVersionTag_1_0_0
	KVaultVersionTag_1_0_0_beta2
	KVaultVersionTag_0_11_5
	KVaultVersionTag_1_0_0_beta1
	KVaultVersionTag_0_11_4
	KVaultVersionTag_0_11_3
	KVaultVersionTag_0_11_2
	KVaultVersionTag_0_11_1
	KVaultVersionTag_0_11_0
	KVaultVersionTag_0_11_0_beta1
	KVaultVersionTag_0_10_4
	KVaultVersionTag_0_10_3
	KVaultVersionTag_0_10_2
	KVaultVersionTag_0_10_1
	KVaultVersionTag_0_10_0
	KVaultVersionTag_0_10_0_rc1
	KVaultVersionTag_0_9_6
	KVaultVersionTag_0_10_0_beta1
	KVaultVersionTag_0_9_5
	KVaultVersionTag_0_9_4
	KVaultVersionTag_0_9_3
	KVaultVersionTag_0_9_2
	KVaultVersionTag_0_9_1
	KVaultVersionTag_0_9_0
	KVaultVersionTag_0_8_3
	KVaultVersionTag_0_8_2
	KVaultVersionTag_0_8_1
	KVaultVersionTag_0_8_0
	KVaultVersionTag_0_7_3
	KVaultVersionTag_0_7_2
	KVaultVersionTag_0_7_0
	KVaultVersionTag_0_6_5
	KVaultVersionTag_0_6_4
	KVaultVersionTag_0_8_0_rc1
	KVaultVersionTag_0_8_0_beta1
	KVaultVersionTag_0_6_3
	KVaultVersionTag_0_6_2
	KVaultVersionTag_0_6
	KVaultVersionTag_0_6_1
	KVaultVersionTag_0_6_0
	KVaultVersionTag_v0_6_0
)

func (el VaultVersionTag) String() string {
	return vaultVersionTags[el]
}

var vaultVersionTags = [...]string{
	"latest",
	"1.5.0",
	"1.5.0-rc",
	"1.4.3",
	"1.3.7",
	"1.4.2",
	"1.3.6",
	"1.4.1",
	"1.3.5",
	"1.4.0",
	"1.4.0-rc1",
	"1.3.4",
	"1.3.3",
	"1.4.0-beta1",
	"1.3.2",
	"1.3.1",
	"1.3.0",
	"1.2.4",
	"1.3.0-beta1",
	"1.2.3",
	"1.2.2",
	"1.2.1",
	"1.2.0",
	"1.1.5",
	"1.2.0-beta2",
	"1.2.0-beta1",
	"1.1.3",
	"1.1.2",
	"1.1.1",
	"1.1.0",
	"1.1.0-beta2",
	"1.0.3",
	"0.11.6",
	"1.0.2",
	"1.0.1",
	"1.0.0",
	"1.0.0-beta2",
	"0.11.5",
	"1.0.0-beta1",
	"0.11.4",
	"0.11.3",
	"0.11.2",
	"0.11.1",
	"0.11.0",
	"0.11.0-beta1",
	"0.10.4",
	"0.10.3",
	"0.10.2",
	"0.10.1",
	"0.10.0",
	"0.10.0-rc1",
	"0.9.6",
	"0.10.0-beta1",
	"0.9.5",
	"0.9.4",
	"0.9.3",
	"0.9.2",
	"0.9.1",
	"0.9.0",
	"0.8.3",
	"0.8.2",
	"0.8.1",
	"0.8.0",
	"0.7.3",
	"0.7.2",
	"0.7.0",
	"0.6.5",
	"0.6.4",
	"0.8.0-rc1",
	"0.8.0-beta1",
	"0.6.3",
	"0.6.2",
	"0.6",
	"0.6.1",
	"0.6.0",
	"v0.6.0",
}
