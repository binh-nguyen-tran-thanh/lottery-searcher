package postgres

import (
	"backend/internal/adapter/repository/postgres/models"
	"backend/internal/core/domain"
	"backend/internal/core/port"
	"backend/internal/core/util"
	"backend/internal/core/util/exception"
	"fmt"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type lotteryRepo struct {
	db     *gorm.DB
	logger port.Logger
}

func NewLotteryRepo(db *gorm.DB, logger port.Logger) port.LotteryRepository {
	return lotteryRepo{
		db:     db,
		logger: logger,
	}
}

func (l lotteryRepo) SyncResult(result []domain.Result) (returnValues []domain.Result, err error) {
	if len(result) <= 0 {
		return nil, exception.New(exception.TypeInternal, "Empty result", nil)
	}

	var resultToSave []models.Result

	for idx := range result {
		resultToSave = append(resultToSave, models.AsResult(result[idx]))
	}

	res := l.db.Debug().Create(&resultToSave)

	if err := res.Error; err != nil {
		return nil, exception.New(exception.TypeInternal, "Fail to save result", err)
	}

	for _, v := range resultToSave {
		returnValues = append(returnValues, v.ToDomain())
	}

	return returnValues, nil
}

func (l lotteryRepo) SaveOpenNumb(data []domain.OpenNum) error {
	if len(data) <= 0 {
		return exception.New(exception.TypeNotFound, "Empty value", nil)
	}

	var valueToSave []models.OpenNumb

	for _, v := range data {
		valueToSave = append(valueToSave, models.AsOpenNumb(v))
	}

	result := l.db.Debug().Omit(clause.Associations).Create(valueToSave)

	if result.Error != nil {
		return exception.New(exception.TypeInternal, "Fail to save", result.Error)
	}

	return nil
}

func (l lotteryRepo) DeleteResult(result domain.Result) error {
	res := l.db.Delete(&result)

	if res.Error != nil {
		return exception.New(exception.TypeInternal, "Delete failed", res.Error)
	}

	return nil
}

func buildFilterOpenNumbsJoinStatement(payload port.FilterOpenNumbPayload) string {
	var joinConditionParts []string

	if payload.Region != "" {
		joinConditionParts = append(joinConditionParts, "results.region = @Region")
	}

	if !payload.FilterRangeFrom.IsZero() && payload.FilterRangeFrom.IsZero() {
		joinConditionParts = append(joinConditionParts, "results.open_time::date < @FilterRangeFrom")
	}

	if !payload.FilterRangeTo.IsZero() && payload.FilterRangeFrom.IsZero() {
		joinConditionParts = append(joinConditionParts, "results.open_time::date > @FilterRangeTo")
	}

	if !payload.FilterRangeTo.IsZero() && !payload.FilterRangeFrom.IsZero() {
		joinConditionParts = append(joinConditionParts, "results.open_time between @FilterRangeFrom and @FilterRangeTo")
	}

	if len(joinConditionParts) > 0 {
		return strings.Join(joinConditionParts, " and ")
	}

	return ""
}

func buildFilterOpenNumbsQueryStatement(payload port.FilterOpenNumbPayload) string {
	var queryParts []string

	if payload.ID != 0 {
		queryParts = append(queryParts, "id = @ID")
	}

	if payload.ResultId != 0 {
		queryParts = append(queryParts, "result_id = @ResultId")
	}

	if len(queryParts) > 0 {
		return strings.Join(queryParts, " and ")
	}

	return ""
}

func (l lotteryRepo) FilterOpenNumbs(payload port.FilterOpenNumbPayload) (returnValues []domain.OpenNum, err error) {
	limit, offset := util.GeneratePagingParamsWithDefaultValue(payload.Limit, payload.Offset)

	var result []models.OpenNumb

	res := l.db.Model(&models.OpenNumb{})

	joinCondition := buildFilterOpenNumbsJoinStatement(payload)
	queryString := buildFilterOpenNumbsQueryStatement(payload)

	if joinCondition != "" {
		joinStatement := fmt.Sprintf("join results on results.id = open_numbs.result_id and %s", joinCondition)
		joinParams := map[string]interface{}{
			"Region":          payload.Region,
			"FilterRangeFrom": util.ToDatabaseFormat(payload.FilterRangeFrom),
			"FilterRangeTo":   util.ToDatabaseFormat(payload.FilterRangeTo),
		}

		res.Joins(joinStatement, joinParams)
	} else {
		res.Joins("Result")
	}

	if queryString != "" {
		res.Where(queryString, payload)
	}

	res.Preload("Result").Order("rank DESC").Offset(int(offset)).Limit(int(limit)).Find(&result)

	if err = res.Error; err != nil {
		return nil, exception.New(exception.DataError, err.Error(), nil)
	}

	for _, v := range result {
		returnValues = append(returnValues, v.ToDomain())
	}

	return
}
func (l lotteryRepo) FilterResult(payload port.FilterResultPayload) (returnValues []domain.Result, err error) {
	var results []models.Result
	res := l.db.Debug().Model(&models.Result{})

	limit, offset := util.GeneratePagingParamsWithDefaultValue(payload.Limit, payload.Offset)

	if payload.Region != "" {
		res.Where("region = ?", payload.Region)
	}

	if payload.ID > 0 {
		res.Where("id = ?", payload.ID)
	}

	if !payload.FilterRangeFrom.IsZero() && payload.FilterRangeFrom.IsZero() {
		res.Where("open_time::date > ?", util.ToDatabaseFormat(payload.FilterRangeFrom))
	}

	if !payload.FilterRangeTo.IsZero() && payload.FilterRangeFrom.IsZero() {
		res.Where("open_time::date < ?", util.ToDatabaseFormat(payload.FilterRangeTo))
	}

	if !payload.FilterRangeTo.IsZero() && !payload.FilterRangeFrom.IsZero() {
		res.Where("open_time between ? and ?", util.ToDatabaseFormat(payload.FilterRangeFrom), util.ToDatabaseFormat(payload.FilterRangeTo))
	}

	res.Offset(int(offset)).Limit(int(limit)).Scan(&results)

	if err = res.Error; err != nil {
		return nil, exception.New(exception.DataError, err.Error(), nil)
	}

	for _, r := range results {
		returnValues = append(returnValues, r.ToDomain())
	}

	return
}

func (l lotteryRepo) FindTodayResult() (returnValues []domain.Result, err error) {
	var results []models.Result
	res := l.db.Debug().Model(&models.Result{}).Where("Date(open_time::date) = CURRENT_DATE").Find(&results)

	if err = res.Error; err != nil {
		return nil, exception.New(exception.DataError, "Fail to query", res.Error)
	}

	for _, result := range results {
		returnValues = append(returnValues, result.ToDomain())
	}

	return
}
