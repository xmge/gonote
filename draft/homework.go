package main

import "time"

const (
	_ = iota
	Before
	Process
	After
	Closed
)

// 一个完整的活动的周期
// 预热 - (0点)进行中 - 冷却 - 空闲
type ActivityConfigList []*ActivityConfigItem

type ActivityConfigItem struct {
	Id string `json:"id"`
	StartDate time.Time `json:"start_date"`
	EndDate time.Time `json:"end_date"`
	DurationDays int `json:"duration_days"` 
	ShowBeforeHours int `json:"show_before_hours"`
	HideAfterHours int `json:"hide_after_hours"`
	IntervalDays int `json:"interval_days"`
	OtherConfigItem map[string]interface{} `json:"other_config_item"` // 其他配置项省略
}

type ActivityProcessDetail struct {
	StartDate time.Time `json:"start_date"`
	EndDate time.Time `json:"end_date"`
	CurrentYureTime time.Time `json:"current_yure_time"`
	CurrentStartTime time.Time `json:"current_start_time"`
	CurrentCollDownTime time.Time `json:"current_coll_down_time"`
	CurrentStopTime time.Time `json:"current_stop_time"`
}

type ActivityService struct {
	activityConfigList ActivityConfigList
}

func NewActivityService(jsonConfig string) *ActivityService {

}



func (service *ActivityService)GetCurrentActivities() []*ActivityProcessDetail {
	now := time.Now()

	// 判断是否在日期内
	activityProcessDetails := make([]*ActivityProcessDetail,0)
	for _, item := range service.activityConfigList {
		// 活动启止日期是否包含当前时间，不包含则 continue

		realStartTime := item.StartDate.Add(-1 * time.Hour * time.Duration(item.ShowBeforeHours))
		realEndTime := item.EndDate.Add(-1 * time.Hour * time.Duration(item.HideAfterHours))
		if !(now.After(realStartTime) && now.Before(realEndTime)) {
			
		}

		//判断是否是进行时间
		if !(now.After(item.StartDate) && now.Before(item.EndDate)) {
			continue
		}

		// 判断是否是预热时间
		if !(now.After(activityConfigItem.StartDate) && now.Before(activityConfigItem.EndDate)) {
			continue
		}

		// 判断是否是冷却时间

		//

		// 生成 ActivityProcessDetail
		detail := &ActivityProcessDetail{}
		detail.StartDate=item.StartDate
		detail.EndDate=item.EndDate
		detail.StartDate=item.StartDate
		detail.StartDate=item.StartDate
		activityProcessDetails = append(activityProcessDetails, detail)
	}
	return activityProcessDetails
}