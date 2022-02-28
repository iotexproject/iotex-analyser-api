package votings

import "github.com/iotexproject/iotex-analyser-api/db"

func GetProbationList(epochNumber uint64) ([]*ProbationList, error) {
	var probationListAll []*ProbationList
	if err := db.DB().Table("probation").Select("epoch_number,intensity_rate,address,count").Where("epoch_number = ?", epochNumber).Find(&probationListAll).Error; err != nil {
		return nil, err
	}
	return probationListAll, nil
}
