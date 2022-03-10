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
	chain {
		mostRecentEpoch
		mostRecentBlockHeight
	}
}


```