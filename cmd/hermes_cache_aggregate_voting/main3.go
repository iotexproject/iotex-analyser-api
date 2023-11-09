package main

import (
	"encoding/csv"
	"fmt"
	"math/big"
	"os"

	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-core/ioctl/util"
	"github.com/pkg/errors"
)

var (
	lsdVoters = []string{
		"io1uzy4hf7r23wyzvjwxczlnw4etqtw3wn2larrcw",
		"io1n5wlsg5udu3msn9dr40rfvztkm69vn8cca8jg5",
		"io1yyxej7yqq4psuv4vhnly9d6wfmyj4pcp4ga4en",
		"io19un9p5jx2r5x85m4uh8ylmphj22ee0h3myxvl3",
		"io1hwugvvl2aumf9jk64vrm98l0vglr923ntq5rfz",
		"io1eafyw2cwcclt7e6j859vydvsz9rq8fep5jyfye",
		"io1pr774gjrn8ecaug749288mrr3077eh54qfej9e",
		"io1779h5nq80ley2wzpzqkaljdpc7tdd6sh2mfhxp",
		"io1k8rmm9jfq0p802qaa426g6gm00nmlsdr5zfa4g",
		"io1w4jxyy9vrlwzz0ahp2p3n2227a3hxwuna54adp",
		"io1zaz6224zjw2z94srldgrtm9smd5x5eruzmulnv",
		"io17axg9h32tnjyk65vculc8u9jkegzjl4maevj84",
		"io14m3erqjsv87t4mzz7m0lpysghjcmg7e7paejg4",
		"io1p7nm6s9ny6l4euykd2nn29unj3vpw4pr53mq6n",
		"io18pzy2xrrwzerpujsl3js43e87kxdrykt5tpdmq",
		"io1yvvlkfcrzlkewye0dc6uaqc4gnvn5wfqkq0ctn",
		"io1hrkvamljtr6h7s5kr2fxq8rr8x0d7v52xaag9w",
		"io19snrtfm6742akugdqg84rc5mcntc5e3dkupaad",
		"io1qhxnu7ama7t7enf9v7jqyqlpg04cx28qpafd0l",
		"io10jvhasjlt5pwxpqjksxmws8tje37ljehhm5ksw",
		"io1ycf88gd9y7uukhjfumzxuv7a0xqff0zkz2k503",
		"io18eltrhe3acrsdm772m42zkmf596j587jsa02lx",
		"io1v8k5lgt7hf9vk7qffwm6uv4uj5j6r3pfdf60wf",
		"io19msajm9hv4u793jvnwcy23plkwzffywjh257sz",
		"io1dcuhfumkk8uznqery04t5cq6edamensask7ef8",
		"io15d0f4yzxg6ka6r4x6prdt3mflkcahlaqx7gtf6",
		"io14cntt59fx92erdm0h4x7jjmlu0lucf2umelpq2",
		"io1cdv420ksuerpgkh6wkdete7yq02ulnsp6usfdw",
		"io1e2rfj7mh4ejawp9dy2ns035305xyahlwmhc0c3",
		"io10dmh3t9uthk7w5kuc823fn96ldwg0m8wnfwy2t",
		"io1e933nrdazcv55dwac6mvrn472qcnfvtvx3uuk6",
		"io1n6wzmpjxn4xpafrx0hcregv0ljt9jzd7zf99m3",
		"io1ujrhg2h94u5elmglysyqjydghlxctrv65skk63",
		"io15ehqtwyu2qw23u0mgq6lhxsg37spmrquf9ljdp",
		"io1tmn3hhhddtew505zd38qy5qzdcgznup3qgtt6n",
		"io1phrtrec42qwwwf8amhmhv89slhy2l7uuesxk2t",
		"io1cew5hztn9c5y3d26yuuwdwdl0wxpm3l42a3eta",
		"io16w526wlurq338h3jxr8a5gw7s7pwepq2lk8l7h",
		"io19jg5h2r5m9qfpwswdat8jzacadk5cljl9j4m98",
		"io1ar4uhyg8l576vp2jl5wzz9wdvg73jhlglwtjad",
		"io1wmk6anemradhzuv2fd27u8cw9jp9msf5crjz7q",
		"io1ftu2zrzl9gd3fxuxg806dylsaa2qjtudsafse7",
		"io16q3ulsvcwx9lzye7mxd74hpz0c6ql4kn3wcnj4",
		"io19s6f55ux5kuamsyfm7gn93y55tt5yn3c8tm6kj",
		"io1q77yjjytau2t0yrs656djzu6xe63umuj9afq6j",
		"io1q226ssx0tsw2z4uthtzcd62alq9fmm32ym72sd",
		"io1v7ts37jua2py826er9t3hr9tru4nne6csck2zl",
		"io1pg84e0mkuzvtnvmhgtpdgldvd5c8g5twmla86y",
		"io1pddye59lnhg00s8md4gf689sw3y4rrgzl48u5p",
		"io1k7e79ff6dlnc796sp8xlut76as89fet58zf3vl",
		"io1x0d7ju9xmvrcj8u2g6glcpk0vqavcs5968d7j9",
		"io1wkflcfuupysc00a73h8590377azht4wh3x6lu5",
		"io1egmsv7ylpl2jz3pfmtxpqdsksgzs9fxscsrn7f",
		"io1l90jmdvwr2lpdke35tr4cpz54acha3gfzxj9zm",
		"io1jslfpm04ady9ht7vm9jl5fy930wa5t8utpwd5d",
		"io1leaukdnk42lf56eeev3l8f06g8hd0tgmep8z96",
	}
)

type VoterAmount struct {
	FixAmount *big.Int
	OriAmount *big.Int
}

func voterInList(voter string) bool {
	for _, v := range lsdVoters {
		if v == voter {
			return true
		}
	}
	return false
}
func main() {
	filelist := []string{
		"hermes-37021-38612.csv",
		"hermes-38613-39612.csv",
	}
	result := make(map[string]VoterAmount)
	for _, fileName := range filelist {
		votesArr, err := readCSV(fileName)
		if err != nil {
			panic(err)
		}
		for _, row := range votesArr {
			voter := row[0]
			if !voterInList(voter) {
				continue
			}
			fixAmount, _ := big.NewInt(0).SetString(row[1], 10)
			oriAmount, _ := big.NewInt(0).SetString(row[2], 10)
			if _, ok := result[voter]; !ok {
				result[voter] = VoterAmount{
					FixAmount: big.NewInt(0),
					OriAmount: big.NewInt(0),
				}
			}
			result[voter] = VoterAmount{
				FixAmount: result[voter].FixAmount.Add(result[voter].FixAmount, fixAmount),
				OriAmount: result[voter].OriAmount.Add(result[voter].OriAmount, oriAmount),
			}
		}
	}
	for voter, amount := range result {
		diffAmount := big.NewInt(0).Sub(amount.FixAmount, amount.OriAmount)
		diffPrice := util.RauToString(diffAmount, util.IotxDecimalNum)
		voterAddr, err := address.FromString(voter)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s,%s,%s\n", voter, voterAddr.Hex(), diffPrice)
	}
}

func readCSV(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	csvReader := csv.NewReader(file)
	csvReader.FieldsPerRecord = -1
	votesArr, err := csvReader.ReadAll()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return votesArr, nil
}
