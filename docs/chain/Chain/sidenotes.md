#! GraphQL Request
```json
query {
	Chain {
		mostRecentEpoch
		mostRecentBlockHeight
	}
}

```
#! GraphQL Response
```json
{
	"data": {
		"Chain": {
			"mostRecentBlockHeight": 16754877,
			"mostRecentEpoch": 25794
		}
	}
}
```
#! RESTful Request
```json
POST /api.ChainService.Chain

{}
```
#! RESTful Response
```json
{
	"mostRecentEpoch": "25794",
	"mostRecentBlockHeight": "16755041"
}
```