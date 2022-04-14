---
Title: BucketInfo
---

BucketInfo provides voting bucket detail information for candidates within a range of epochs

#- Attributes
- startEpoch `uint64`
- Epoch number to start from
- epochCount `uint64`
- Number of epochs to query
- delegateName `string`
- Name of the delegate
- pagination 
- Pagination info
  - skip `uint64`
  - starting index of results
  - first `uint64`
  - number of records per page 

#- Response Values
- exist `bool`
- Whether the delegate has voting bucket information within the specified epoch range
- count `uint64`
- Total number of buckets in the given epoch for the given delegate
- bucketInfoList
- bucket info list
  - epochNumber `uint64`
  - Epoch number
  - count `uint64`
  - Count for epoch
  - bucketInfo
  - bucket infomation
    - voterEthAddress `string`
    - voter’s ERC20 address
    - voterIotexAddress `string`
    - voter's IoTeX address
    - isNative `bool`
    - whether the bucket is native
    - votes `string`
    - voter's votes
    - weightedVotes `string`
    - voter’s weighted votes
    - remainingDuration `string`
    - bucket remaining duration
    - startTime `string`
    - bucket start time
    - decay `bool`
    - whether the vote weight decays
    - bucketID `uint64`
    - Bucket ID
