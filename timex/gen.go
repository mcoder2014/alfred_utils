package main

import (
	"strconv"
	"time"

	"github.com/mcoder2014/alfred_utils/utils"

	aw "github.com/deanishe/awgo"
)

func Gen(w *aw.Workflow) error {
	currentTime := time.Now()
	utils.NewCopyableItem(w, "当前时间（秒）", currentTime.Format("2006-01-02 15:04:05"))
	utils.NewCopyableItem(w, "当前秒时间戳", strconv.FormatInt(currentTime.Unix(), 10))
	utils.NewCopyableItem(w, "当前毫秒时间戳", strconv.FormatInt(currentTime.UnixMilli(), 10))
	utils.NewCopyableItem(w, "当前微秒时间戳", strconv.FormatInt(currentTime.UnixMicro(), 10))
	utils.NewCopyableItem(w, "当前纳秒时间戳", strconv.FormatInt(currentTime.UnixNano(), 10))
	utils.NewCopyableItem(w, "以当前时间生成游标", strconv.FormatInt(utils.GetIDFromTime(currentTime), 10))
	utils.NewCopyableItem(w, "以当前时间生成随机id", strconv.FormatInt(utils.GenID(currentTime), 10))
	return nil
}
