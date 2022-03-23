DelegateService

BucketInfo
```
{
	BucketInfo(
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
BookKeeping
```
query {
	BookKeeping(
		startEpoch: 23328
		epochCount: 1201
		delegateName: "iotexlab"
		percentage: 90
		includeFoundationBonus: false
		includeBlockReward: false
		pagination: { skip: 0, first: 66666 }
	) {
		count
		rewardDistribution {
			voterEthAddress
			amount
		}
	}
}

```

Productivity
```
query {
	DelegateProductivity(
		startEpoch: 25020, epochCount: 10, delegateName: "iotexlab"
	) {
		productivity {
			exist
			production
			expectedProduction
		}
	}
}

```

Reward
```
query {
	DelegateReward(startEpoch: 25000, epochCount: 120, delegateName: "iotexlab") {
		reward {
			exist
			blockReward
			foundationBonus
			epochReward
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