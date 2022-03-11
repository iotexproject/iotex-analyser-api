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

ActionService

GetActionByVoter = v1 actionByVoter
```
query{
	GetActionByVoter(
		address: "io19msajm9hv4u793jvnwcy23plkwzffywjh257sz"
		pagination: { skip: 0, first: 300 }
	) {
		actionList {
			exist
			count
			actions {
          amount
          actHash
          actType
          timeStamp
			}
		}
	}
}
```