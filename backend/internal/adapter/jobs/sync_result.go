package jobs

import (
	"backend/internal/core/domain"
	"backend/internal/core/util"
	"backend/internal/core/util/exception"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type IssueList struct {
	TurnNum       string
	OpenTime      string
	OpenNum       string
	Detail        string
	OpenTimeStamp int
	Status        int
}

type ExternalResult struct {
	TurnNum   string
	OpenTime  string
	Code      string
	IssueList []IssueList
}

type ExternalAPIResponse struct {
	Success bool
	Msg     string
	Code    int
	T       ExternalResult
}

var errorFormat string = "Fail parse result for region %v. Reason: %v"
var timeFormat string = "02-01-2006 15:04:05"

func (c *CronJob) saveOpenNumbs(result domain.Result, regionName string, numberToSave *[]domain.OpenNum) {
	var details []string
	if err := json.Unmarshal([]byte(result.Detail), &details); err != nil {
		c.logger.Fatal().Msgf("Parse Detail failed", regionName, err.Error())
	}

	if len(details) < 8 {
		if err := c.repository.Lottery().DeleteResult(result); err != nil {
			c.logger.Info().Msgf("Delete region %s due to not enough result", regionName, err.Error())
		}
	} else {
		for rank, detail := range details {
			*numberToSave = append(*numberToSave, domain.OpenNum{
				ResultID: result.ID,
				Numbs:    detail,
				Rank:     int8(rank),
				Result:   result,
			})
		}
	}

}

func (c *CronJob) retrieveLotteryResult(wg *sync.WaitGroup, region domain.Region) error {
	defer wg.Done()

	regionCode := region.Code

	requestUrl := fmt.Sprintf("%v/%v/%v", c.config.CronJob.LotteryDomain, c.config.CronJob.Limit, regionCode)

	c.logger.Info().Msgf("Retrieving result for %v at %v by %v", region.Name, time.Now().Format(timeFormat), requestUrl)

	resp, err := http.Get(requestUrl)

	if err != nil {
		c.logger.Fatal().Msgf("Failed to get result for region %v. Reason: %v", region.Name, err.Error())
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		c.logger.Fatal().Msgf(errorFormat, region.Name, err.Error())
	}

	externalResults := ExternalAPIResponse{}
	if err := json.Unmarshal(body, &externalResults); err != nil {
		c.logger.Fatal().Msgf("Parse response failed for region %s", region.Name, err.Error())
	}

	if !externalResults.Success {
		c.logger.Fatal().Msgf(errorFormat, region.Name, externalResults.Msg)
	}

	if len(externalResults.T.IssueList) == 0 {
		c.logger.Fatal().Msgf(errorFormat, region.Name, externalResults.Msg)
	}

	var resultToSave []domain.Result

	for _, issue := range externalResults.T.IssueList {
		if isOpenBeforeToday, _ := util.IsBeforeNow(issue.OpenTime, util.DateFormat); isOpenBeforeToday && !c.config.CronJob.SkipTimeCheck {
			c.logger.Info().Msgf("Skip saving %v because the open turn was not complete. Date: %s, Now: %s", region.Name, issue.OpenTime, util.GetToDay())
			continue
		}

		resultToSave = append(resultToSave,
			domain.Result{
				TurnNum:  issue.TurnNum,
				OpenNum:  issue.OpenNum,
				OpenTime: issue.OpenTime,
				Region:   regionCode,
				Detail:   issue.Detail,
			},
		)
	}

	if len(resultToSave) <= 0 {
		c.logger.Info().Msgf("Skip saving %v because no new open turn founded", region.Name)
		return exception.New(exception.TypeNotFound, "Empty value", nil)
	}

	savedResults, err := c.repository.Lottery().SyncResult(resultToSave)

	if err != nil {
		c.logger.Fatal().Msgf(errorFormat, region.Name, err.Error())
	}

	c.logger.Info().Msgf("Saved result for %v region", region.Name)

	var numberToSave []domain.OpenNum

	for _, v := range savedResults {
		c.saveOpenNumbs(v, region.Name, &numberToSave)
	}

	if err := c.repository.Lottery().SaveOpenNumb(numberToSave); err != nil {
		c.logger.Fatal().Msgf("Failed to update open numbs for %v. Reason: %v", region.Name, err.Error())
	}

	c.logger.Info().Msgf("Updated open numbs for %v region", region.Name)

	if _, err := c.repository.Region().UpdateRegionOpenTime(region.ID, externalResults.T.OpenTime); err != nil {
		c.logger.Fatal().Msgf("Failed to update turn num for %v. Reason: %v", region.Name, err.Error())
	}

	c.logger.Info().Msgf("Updated turn num for %v region", region.Name)
	return nil
}

func (c *CronJob) StartSyncResult() {
	c.logger.Info().Msgf("Running Sync Job at %v", time.Now().Format(timeFormat))

	regions, err := c.repository.Region().GetRegionHasTurnToday()

	if err != nil {
		c.logger.Fatal().Msgf("Can not get the region list %v", err.Error())
	}

	wg := sync.WaitGroup{}

	for _, region := range regions {
		wg.Add(1)
		go c.retrieveLotteryResult(&wg, region)
		time.Sleep(time.Duration(c.config.CronJob.Delay) * time.Millisecond)
	}

	wg.Wait()

	c.logger.Info().Msgf("Finished Sync Job at %v", time.Now().Format("2006-01-02 15:04:05"))
}
