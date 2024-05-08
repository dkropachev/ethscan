package types_test

import (
	"encoding/json"
	"github.com/dkropachev/ethscan/pkg/types"
	"testing"

	"github.com/nsf/jsondiff"
)

var jsonCMPOptions = &jsondiff.Options{
	Added:            jsondiff.Tag{Begin: "++++ ", End: " ===="},
	Removed:          jsondiff.Tag{Begin: "---- ", End: " ===="},
	Changed:          jsondiff.Tag{Begin: "---- ", End: " ===="},
	ChangedSeparator: " ++++ ",
	Indent:           "    ",
}

func TestBlock(t *testing.T) {
	b := "{\"jsonrpc\":\"2.0\",\"id\":1,\"result\":{\"baseFeePerGas\":\"0x151b9fc1b\",\"blobGasUsed\":\"0x20000\",\"difficulty\":\"0x0\",\"excessBlobGas\":\"0x0\",\"extraData\":\"0x546974616e2028746974616e6275696c6465722e78797a29\",\"gasLimit\":\"0x1c9c380\",\"gasUsed\":\"0x1637178\",\"hash\":\"0xb4ac5e3d870d4d4535c69e7a22dbcc83d1cf238608b3126c0b7c65fce37acaf4\",\"logsBloom\":\"0xc1b1abdaef807fd1579bf9f79e423a5b3c5f80d3fdbbfa2f3659115cde510b8665d9c5c7f9ef08fbab88d830dede75d42bb3fd1c9b19bc5a976db220bebf0a84fd61b7bf599dadefeeefcdbad2786bfe94a7fc2184ef7fc85725de6fa8a6de83cfedad03077f2d25b6afd3f9a338ecf746df4df7bafc0f6bf0f6ef1f806ea77d1a4a1ed983cb2b6e3b5a075e6b62290d1be9cd49bba9d6bb536748477f3c7a3eaf9d7dc2cfb9ebed6e9f52dffa7c17f585fdc798fafefb5e0affc2efd5d80b605f56f9be4cf6afcddb4262fecd4fd6efe4dfdc9bd76f02f57caf738bd1f2fb6c8a7bfe3b9f272464dc65bdd7706989b7cabcfe38fffde27f0805e17c1b27dec7\",\"miner\":\"0x4838b106fce9647bdf1e7877bf73ce8b0bad5f97\",\"mixHash\":\"0x465b99acf806ce4f00ca94f899054df56c8ae75ad59aa0705b7d5a9cbf9254b2\",\"nonce\":\"0x0000000000000000\",\"number\":\"0x12d3a0b\",\"parentBeaconBlockRoot\":\"0x21e17a53e5936f31baf5086bc50b3aed56c90bdf101c0572e96c803c19509481\",\"parentHash\":\"0xedb4b61747d8b193a07c5a54c8c73370c8ca9b29206e9c8e63f6eba24960f77a\",\"receiptsRoot\":\"0x1bf97f892fc7b5d6f14df852e413155e1ae031d80c8f5112255ff4850e5c970b\",\"sha3Uncles\":\"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347\",\"size\":\"0x441fc\",\"stateRoot\":\"0x973a51a33f75ae3ad0e5d60ba351628b9d0b83dd5f1c000fa750a8d5177e0f3e\",\"timestamp\":\"0x662becd7\",\"totalDifficulty\":\"0xc70d815d562d3cfa955\",\"transactions\":[{\"blockHash\":\"0xb4ac5e3d870d4d4535c69e7a22dbcc83d1cf238608b3126c0b7c65fce37acaf4\",\"blockNumber\":\"0x12d3a0b\",\"chainId\":\"0x1\",\"from\":\"0x75e89d5979e4f6fba9f97c104c2f0afb3f1dcb88\",\"gas\":\"0xc350\",\"gasPrice\":\"0x306dc4200\",\"hash\":\"0x5fc0fd88da12e3900d7614e29830568cd33133c00dbf70fd3d7c8cc525bec853\",\"input\":\"0x\",\"nonce\":\"0x657b2a\",\"r\":\"0x8d48619719e0b301bb142624883e2159e24ceee560b488ecd8abe8a07b9ce3ac\",\"s\":\"0x7defde8319ae0e626c2bdb44054f225104994e19752dd566b602325b9fd86ece\",\"to\":\"0xf0408039e030547b90b77475cd54c3e2c9410e21\",\"transactionIndex\":\"0x0\",\"type\":\"0x0\",\"v\":\"0x25\",\"value\":\"0x14f604cc2cc000\"},{\"blockHash\":\"0xb4ac5e3d870d4d4535c69e7a22dbcc83d1cf238608b3126c0b7c65fce37acaf4\",\"blockNumber\":\"0x12d3a0b\",\"chainId\":\"0x1\",\"from\":\"0x75e89d5979e4f6fba9f97c104c2f0afb3f1dcb88\",\"gas\":\"0x14e29\",\"gasPrice\":\"0x306dc4200\",\"hash\":\"0xfc3b7f0fa5f88662d5161dbef8608d6956d1e25399dbb5e3669aebdd9de05c5d\",\"input\":\"0xa9059cbb000000000000000000000000b9aa69c21a8a360dfe9effb079955f789d7c6715000000000000000000000000000000000000000000000013c86bcd849b4d0000\",\"nonce\":\"0x657b2b\",\"r\":\"0xf21c09a71b9cbe3e22c92e65b75459238f89c18fa2d1482810d85470c8e091c4\",\"s\":\"0x366905345785d2387cbf333d35bbe4f627fafb2100a5aea5c1546f6a666c2805\",\"to\":\"0x9813037ee2218799597d83d4a5b6f3b6778218d9\",\"transactionIndex\":\"0x1\",\"type\":\"0x0\",\"v\":\"0x26\",\"value\":\"0x0\"}],\"transactionsRoot\":\"0x2609fbba0d100c19c0e233df04f374748c09d06d90714061af248421b557193c\",\"uncles\":[],\"withdrawals\":[{\"address\":\"0x7addee2a2540e0ae72d1e95936f792b736dd9908\",\"amount\":\"0x11af587\",\"index\":\"0x294b39b\",\"validatorIndex\":\"0x2dba5\"},{\"address\":\"0x7addee2a2540e0ae72d1e95936f792b736dd9908\",\"amount\":\"0x11a5680\",\"index\":\"0x294b39c\",\"validatorIndex\":\"0x2dba6\"},{\"address\":\"0x7addee2a2540e0ae72d1e95936f792b736dd9908\",\"amount\":\"0x38acc9a\",\"index\":\"0x294b39d\",\"validatorIndex\":\"0x2dba7\"},{\"address\":\"0x7addee2a2540e0ae72d1e95936f792b736dd9908\",\"amount\":\"0x118f447\",\"index\":\"0x294b39e\",\"validatorIndex\":\"0x2dba8\"},{\"address\":\"0x7addee2a2540e0ae72d1e95936f792b736dd9908\",\"amount\":\"0x11a7c90\",\"index\":\"0x294b39f\",\"validatorIndex\":\"0x2dba9\"},{\"address\":\"0x7addee2a2540e0ae72d1e95936f792b736dd9908\",\"amount\":\"0x11a08da\",\"index\":\"0x294b3a0\",\"validatorIndex\":\"0x2dbaa\"},{\"address\":\"0x7addee2a2540e0ae72d1e95936f792b736dd9908\",\"amount\":\"0x11a2e5b\",\"index\":\"0x294b3a1\",\"validatorIndex\":\"0x2dbab\"},{\"address\":\"0x7addee2a2540e0ae72d1e95936f792b736dd9908\",\"amount\":\"0x119fc56\",\"index\":\"0x294b3a2\",\"validatorIndex\":\"0x2dbac\"},{\"address\":\"0x7addee2a2540e0ae72d1e95936f792b736dd9908\",\"amount\":\"0x11ac1b4\",\"index\":\"0x294b3a3\",\"validatorIndex\":\"0x2dbad\"},{\"address\":\"0x7addee2a2540e0ae72d1e95936f792b736dd9908\",\"amount\":\"0x11aa7f9\",\"index\":\"0x294b3a4\",\"validatorIndex\":\"0x2dbae\"},{\"address\":\"0x7addee2a2540e0ae72d1e95936f792b736dd9908\",\"amount\":\"0x11a8bda\",\"index\":\"0x294b3a5\",\"validatorIndex\":\"0x2dbaf\"},{\"address\":\"0x7addee2a2540e0ae72d1e95936f792b736dd9908\",\"amount\":\"0x11a94ed\",\"index\":\"0x294b3a6\",\"validatorIndex\":\"0x2dbb0\"},{\"address\":\"0x7addee2a2540e0ae72d1e95936f792b736dd9908\",\"amount\":\"0x118952d\",\"index\":\"0x294b3a7\",\"validatorIndex\":\"0x2dbb1\"},{\"address\":\"0x7addee2a2540e0ae72d1e95936f792b736dd9908\",\"amount\":\"0x11a66f7\",\"index\":\"0x294b3a8\",\"validatorIndex\":\"0x2dbb2\"},{\"address\":\"0x7addee2a2540e0ae72d1e95936f792b736dd9908\",\"amount\":\"0x11b3af9\",\"index\":\"0x294b3a9\",\"validatorIndex\":\"0x2dbb3\"},{\"address\":\"0x7addee2a2540e0ae72d1e95936f792b736dd9908\",\"amount\":\"0x11b287c\",\"index\":\"0x294b3aa\",\"validatorIndex\":\"0x2dbb4\"}],\"withdrawalsRoot\":\"0xfdda9cb1baeddb1ec4ed6362a68f6f5fed69c82ab7479f3db2f8485a30dbac0f\"}}"
	t.Run("UnmarshalJSON", func(t *testing.T) {
		actual := struct {
			Result types.Block
		}{}
		err := json.Unmarshal([]byte(b), &actual)
		if err != nil {
			t.Fatal(err)
		}
		//
		// b := "{\"jsonrpc\":\"2.0\",\"id\":1,\"result\":{\"baseFeePerGas\":\"0x151b9fc1b\",\"blobGasUsed\":\"0x20000\",\"difficulty\":\"0x0\",\"excessBlobGas\":\"0x0\",\"extraData\":\"0x546974616e2028746974616e6275696c6465722e78797a29\",\"gasLimit\":\"0x1c9c380\",\"gasUsed\":\"0x1637178\",\"hash\":\"0xb4ac5e3d870d4d4535c69e7a22dbcc83d1cf238608b3126c0b7c65fce37acaf4\",\"logsBloom\":\"0xc1b1abdaef807fd1579bf9f79e423a5b3c5f80d3fdbbfa2f3659115cde510b8665d9c5c7f9ef08fbab88d830dede75d42bb3fd1c9b19bc5a976db220bebf0a84fd61b7bf599dadefeeefcdbad2786bfe94a7fc2184ef7fc85725de6fa8a6de83cfedad03077f2d25b6afd3f9a338ecf746df4df7bafc0f6bf0f6ef1f806ea77d1a4a1ed983cb2b6e3b5a075e6b62290d1be9cd49bba9d6bb536748477f3c7a3eaf9d7dc2cfb9ebed6e9f52dffa7c17f585fdc798fafefb5e0affc2efd5d80b605f56f9be4cf6afcddb4262fecd4fd6efe4dfdc9bd76f02f57caf738bd1f2fb6c8a7bfe3b9f272464dc65bdd7706989b7cabcfe38fffde27f0805e17c1b27dec7\",\"miner\":\"0x4838b106fce9647bdf1e7877bf73ce8b0bad5f97\",\"mixHash\":\"0x465b99acf806ce4f00ca94f899054df56c8ae75ad59aa0705b7d5a9cbf9254b2\",\"nonce\":\"0x0000000000000000\",\"number\":\"0x12d3a0b\",\"parentBeaconBlockRoot\":\"0x21e17a53e5936f31baf5086bc50b3aed56c90bdf101c0572e96c803c19509481\",\"parentHash\":\"0xedb4b61747d8b193a07c5a54c8c73370c8ca9b29206e9c8e63f6eba24960f77a\",\"receiptsRoot\":\"0x1bf97f892fc7b5d6f14df852e413155e1ae031d80c8f5112255ff4850e5c970b\",\"sha3Uncles\":\"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347\",\"size\":\"0x441fc\",\"stateRoot\":\"0x973a51a33f75ae3ad0e5d60ba351628b9d0b83dd5f1c000fa750a8d5177e0f3e\",\"timestamp\":\"0x662becd7\",\"totalDifficulty\":\"0xc70d815d562d3cfa955\",
		out, err := json.Marshal(actual.Result)
		if err != nil {
			t.Fatal(err)
		}
		expected := "{\n        \t\"baseFeePerGas1\": 5666110491,\n        \t\"blobGasUsed\": 131071,\n        \t\"difficulty\": 0,\n        \t\"excessBlobGas\": 0,\n        \t\"extraData\": \"0x546974616e2028746974616e6275696c6465722e78797a2900\",\n        \t\"gasLimit\": 30000000,\n        \t\"gasUsed\": 23294328,\n        \t\"hash\": \"0xb4ac5e3d870d4d4535c69e7a22dbcc83d1cf238608b3126c0b7c65fce37acaf4\",\n        \t\"miner\": \"0x4838b106fce9647bdf1e7877bf73ce8b0bad5f97\",\n        \t\"mixHash\": \"0x465b99acf806ce4f00ca94f899054df56c8ae75ad59aa0705b7d5a9cbf9254b2\",\n        \t\"nonce\": 0,\n        \t\"number\": 19741195,\n        \t\"parentHash\": \"0xedb4b61747d8b193a07c5a54c8c73370c8ca9b29206e9c8e63f6eba24960f77a\",\n        \t\"receiptsRoot\": \"0x1bf97f892fc7b5d6f14df852e413155e1ae031d80c8f5112255ff4850e5c970b\",\n        \t\"sha3Uncles\": \"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347\",\n        \t\"size\": 279036,\n        \t\"stateRoot\": \"0x973a51a33f75ae3ad0e5d60ba351628b9d0b83dd5f1c000fa750a8d5177e0f3e\",\n        \t\"timestamp\": 1714154711,\n        \t\"totalDifficulty\": 58750003716598352816469,\n        \t\"transactions\": [\n        \t\t{\n        \t\t\t\"blockHash\": \"0xb4ac5e3d870d4d4535c69e7a22dbcc83d1cf238608b3126c0b7c65fce37acaf4\",\n        \t\t\t\"blockNumber\": 19741195,\n        \t\t\t\"from\": \"0x75e89d5979e4f6fba9f97c104c2f0afb3f1dcb88\",\n        \t\t\t\"gas\": 50000,\n        \t\t\t\"gasPrice\": 13000000000,\n        \t\t\t\"hash\": \"0x5fc0fd88da12e3900d7614e29830568cd33133c00dbf70fd3d7c8cc525bec853\",\n        \t\t\t\"input\": \"0x00\",\n        \t\t\t\"nonce\": 6650666,\n        \t\t\t\"to\": \"0xf0408039e030547b90b77475cd54c3e2c9410e21\",\n        \t\t\t\"transactionIndex\": 0,\n        \t\t\t\"value\": 5900000000000000\n        \t\t},\n        \t\t{\n        \t\t\t\"blockHash\": \"0xb4ac5e3d870d4d4535c69e7a22dbcc83d1cf238608b3126c0b7c65fce37acaf4\",\n        \t\t\t\"blockNumber\": 19741195,\n        \t\t\t\"from\": \"0x75e89d5979e4f6fba9f97c104c2f0afb3f1dcb88\",\n        \t\t\t\"gas\": 85545,\n        \t\t\t\"gasPrice\": 13000000000,\n        \t\t\t\"hash\": \"0xfc3b7f0fa5f88662d5161dbef8608d6956d1e25399dbb5e3669aebdd9de05c5d\",\n        \t\t\t\"input\": \"0xa9059cbb000000000000000000000000b9aa69c21a8a360dfe9effb079955f789d7c6715000000000000000000000000000000000000000000000013c86bcd849b4d000000\",\n        \t\t\t\"nonce\": 6650667,\n        \t\t\t\"to\": \"0x9813037ee2218799597d83d4a5b6f3b6778218d9\",\n        \t\t\t\"transactionIndex\": 1,\n        \t\t\t\"value\": 0\n        \t\t}\n        \t],\n        \t\"transactionsRoot\": \"0x2609fbba0d100c19c0e233df04f374748c09d06d90714061af248421b557193c\",\n        \t\"uncles\": [],\n        \t\"withdrawals\": [\n        \t\t{\n        \t\t\t\"address\": \"0x7addee2a2540e0ae72d1e95936f792b736dd9908\",\n        \t\t\t\"amount\": 18544007,\n        \t\t\t\"index\": 43299739,\n        \t\t\t\"validatorIndex\": 187301\n        \t\t},\n        \t\t{\n        \t\t\t\"address\": \"0x7addee2a2540e0ae72d1e95936f792b736dd9908\",\n        \t\t\t\"amount\": 18503296,\n        \t\t\t\"index\": 43299740,\n        \t\t\t\"validatorIndex\": 187302\n        \t\t},\n        \t\t{\n        \t\t\t\"address\": \"0x7addee2a2540e0ae72d1e95936f792b736dd9908\",\n        \t\t\t\"amount\": 59427994,\n        \t\t\t\"index\": 43299741,\n        \t\t\t\"validatorIndex\": 187303\n        \t\t},\n        \t\t{\n        \t\t\t\"address\": \"0x7addee2a2540e0ae72d1e95936f792b736dd9908\",\n        \t\t\t\"amount\": 18412615,\n        \t\t\t\"index\": 43299742,\n        \t\t\t\"validatorIndex\": 187304\n        \t\t},\n        \t\t{\n        \t\t\t\"address\": \"0x7addee2a2540e0ae72d1e95936f792b736dd9908\",\n        \t\t\t\"amount\": 18513040,\n        \t\t\t\"index\": 43299743,\n        \t\t\t\"validatorIndex\": 187305\n        \t\t},\n        \t\t{\n        \t\t\t\"address\": \"0x7addee2a2540e0ae72d1e95936f792b736dd9908\",\n        \t\t\t\"amount\": 18483418,\n        \t\t\t\"index\": 43299744,\n        \t\t\t\"validatorIndex\": 187306\n        \t\t},\n        \t\t{\n        \t\t\t\"address\": \"0x7addee2a2540e0ae72d1e95936f792b736dd9908\",\n        \t\t\t\"amount\": 18493019,\n        \t\t\t\"index\": 43299745,\n        \t\t\t\"validatorIndex\": 187307\n        \t\t},\n        \t\t{\n        \t\t\t\"address\": \"0x7addee2a2540e0ae72d1e95936f792b736dd9908\",\n        \t\t\t\"amount\": 18480214,\n        \t\t\t\"index\": 43299746,\n        \t\t\t\"validatorIndex\": 187308\n        \t\t},\n        \t\t{\n        \t\t\t\"address\": \"0x7addee2a2540e0ae72d1e95936f792b736dd9908\",\n        \t\t\t\"amount\": 18530740,\n        \t\t\t\"index\": 43299747,\n        \t\t\t\"validatorIndex\": 187309\n        \t\t},\n        \t\t{\n        \t\t\t\"address\": \"0x7addee2a2540e0ae72d1e95936f792b736dd9908\",\n        \t\t\t\"amount\": 18524153,\n        \t\t\t\"index\": 43299748,\n        \t\t\t\"validatorIndex\": 187310\n        \t\t},\n        \t\t{\n        \t\t\t\"address\": \"0x7addee2a2540e0ae72d1e95936f792b736dd9908\",\n        \t\t\t\"amount\": 18516954,\n        \t\t\t\"index\": 43299749,\n        \t\t\t\"validatorIndex\": 187311\n        \t\t},\n        \t\t{\n        \t\t\t\"address\": \"0x7addee2a2540e0ae72d1e95936f792b736dd9908\",\n        \t\t\t\"amount\": 18519277,\n        \t\t\t\"index\": 43299750,\n        \t\t\t\"validatorIndex\": 187312\n        \t\t},\n        \t\t{\n        \t\t\t\"address\": \"0x7addee2a2540e0ae72d1e95936f792b736dd9908\",\n        \t\t\t\"amount\": 18388269,\n        \t\t\t\"index\": 43299751,\n        \t\t\t\"validatorIndex\": 187313\n        \t\t},\n        \t\t{\n        \t\t\t\"address\": \"0x7addee2a2540e0ae72d1e95936f792b736dd9908\",\n        \t\t\t\"amount\": 18507511,\n        \t\t\t\"index\": 43299752,\n        \t\t\t\"validatorIndex\": 187314\n        \t\t},\n        \t\t{\n        \t\t\t\"address\": \"0x7addee2a2540e0ae72d1e95936f792b736dd9908\",\n        \t\t\t\"amount\": 18561785,\n        \t\t\t\"index\": 43299753,\n        \t\t\t\"validatorIndex\": 187315\n        \t\t},\n        \t\t{\n        \t\t\t\"address\": \"0x7addee2a2540e0ae72d1e95936f792b736dd9908\",\n        \t\t\t\"amount\": 18557052,\n        \t\t\t\"index\": 43299754,\n        \t\t\t\"validatorIndex\": 187316\n        \t\t}\n        \t]\n        }"
		status, diff := jsondiff.Compare([]byte(expected), out, jsonCMPOptions)
		if status != jsondiff.FullMatch {
			t.Error(diff)
		}
	})
}