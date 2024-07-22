package main

import (
	"strconv"
	"time"

	"github.com/mcoder2014/alfred_utils/utils"

	aw "github.com/deanishe/awgo"
)

func Gen(w *aw.Workflow) error {
	currentTime := time.Now()
	utils.NewCopyableItem(w, "当前时间", currentTime.Format("2006-01-02 15:04:05"))
	utils.NewCopyableItem(w, "时间 RFC3339", currentTime.Format(time.RFC3339))
	utils.NewCopyableItem(w, "当前秒时间戳", strconv.FormatInt(currentTime.Unix(), 10))
	utils.NewCopyableItem(w, "当前毫秒时间戳", strconv.FormatInt(currentTime.UnixMilli(), 10))
	utils.NewCopyableItem(w, "当前微秒时间戳", strconv.FormatInt(currentTime.UnixMicro(), 10))
	utils.NewCopyableItem(w, "当前纳秒时间戳", strconv.FormatInt(currentTime.UnixNano(), 10))
	utils.NewCopyableItem(w, "以当前时间生成游标", strconv.FormatInt(utils.GetIDFromTime(currentTime), 10))
	utils.NewCopyableItem(w, "以当前时间生成随机id", strconv.FormatInt(utils.GenID(currentTime), 10))

	todayZeroTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
	utils.NewCopyableItem(w, "以今日0时生成随机id", strconv.FormatInt(utils.GenID(todayZeroTime), 10))
	utils.NewCopyableItem(w, "以昨日0时生成随机id", strconv.FormatInt(utils.GenID(todayZeroTime.Add(-24*time.Hour)), 10))
	utils.NewCopyableItem(w, "以前天0时生成随机id", strconv.FormatInt(utils.GenID(todayZeroTime.Add(-2*24*time.Hour)), 10))
	utils.NewCopyableItem(w, "以一周前0时生成随机id", strconv.FormatInt(utils.GenID(todayZeroTime.Add(-7*24*time.Hour)), 10))
	return nil
}
