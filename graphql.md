DelegateService
```
{
	GetBucketInfo(
		startEpoch: 24738
		epochCount: 1
		delegateName: "metanyx"
		pagination: { skip: 0, first: 30 }
	) {
		bucketInfo {
			exist
			bucketInfoList {
				epochNumber
				count
				bucketInfo {
					voterIotexAddress
					votes
					weightedVotes
					remainingDuration
					isNative
				}
			}
		}
	}
}

```
ChainService

```
query {
	Chain {
		mostRecentEpoch
		mostRecentBlockHeight
	}
}


```

## ActionService

### GetActionByVoter = v1 action.ByVoter
```
query {
	GetActionByVoter(
		address: "io19msajm9hv4u793jvnwcy23plkwzffywjh257sz"
		pagination: { skip: 10, first: 10 }
	) {
		exist
		count
		actionList {
			amount
			actHash
			actType
			timestamp
		}
	}
}
```
### GetEvmTransfersByAddress = v1 action.evmTransfersByAddress
```
query {
	GetEvmTransfersByAddress(
		address: "io19msajm9hv4u793jvnwcy23plkwzffywjh257sz"
		pagination: { skip: 0, first: 10 }
	) {
		exist
		count
		evmTransferList {
			quantity
			actHash
			from
			to
			blkHash
			blkHeight
			timestamp
		}
	}
}
```
### GetXrc20ByAddress = v1 xrc20.byAddress
```
query {
	GetXrc20ByAddress(
		address: "io19msajm9hv4u793jvnwcy23plkwzffywjh257sz"
		pagination: { skip: 0, first: 10 }
	) {
		exist
		count
		xrcList {
			actHash
			from
			to
			quantity
			timestamp
			contract
		}
	}
}

```