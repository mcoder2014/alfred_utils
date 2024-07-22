package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/mcoder2014/alfred_utils/utils"

	aw "github.com/deanishe/awgo"
	"github.com/sirupsen/logrus"
)

func Trans(w *aw.Workflow, param string) error {
	param = strings.TrimSpace(param)

	// 当做 int64 处理
	err := transInt64(w, param)
	if err != nil {
		logrus.Infof("err: %v", err)
	}

	// 当做格式化日期处理
	err = transFormatTime(w, param)
	if err != nil {
		logrus.Infof("err: %v", err)
	}
	return nil
}

func transInt64(w *aw.Workflow, param string) error {
	val, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return fmt.Errorf("can not parse %s to int64", param)
	}

	// 秒时间戳 转换成时间
	// "2006-01-02 15:04:05.999999999 -0700 MST"
	t := time.Unix(val, 0)
	utils.NewCopyableItem(w, "秒时间戳", t.Format("2006-01-02 15:04:05"))
	utils.NewCopyableItem(w, "秒时间戳", t.Format("2006-01-02 15:04:05 MST"))
	utils.NewCopyableItem(w, "时间 RFC3339", t.Format(time.RFC3339))

	// 毫秒时间戳
	t = time.Unix(val/1e3, val%1e3)
	utils.NewCopyableItem(w, "毫秒时间戳", t.Format("2006-01-02 15:04:05.999 MST"))

	// 微秒时间戳
	t = time.Unix(val/1e6, val%1e6)
	utils.NewCopyableItem(w, "微秒时间戳", t.Format("2006-01-02 15:04:05.999999 MST"))

	// 纳秒时间戳
	t = time.Unix(val/1e9, val%1e9)
	utils.NewCopyableItem(w, "纳秒时间戳", t.Format("2006-01-02 15:04:05.999999999 MST"))

	// 作为 id 处理
	t = utils.GetTimeFromID(val)
	utils.NewCopyableItem(w, "id 提取时间前缀", t.Format("2006-01-02 15:04:05"))
	utils.NewCopyableItem(w, "id 提取时间戳前缀", strconv.FormatInt(t.Unix(), 10))
	return nil
}

func transFormatTime(w *aw.Workflow, param string) error {
	// 时区为"Asia/Shanghai"
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return fmt.Errorf("无法加载时区：%w", err)
	}

	var convert = func(layout string, t time.Time) {
		utils.NewCopyableItem(w, fmt.Sprintf("%s 转时间戳", layout), strconv.FormatInt(t.Unix(), 10))
		utils.NewCopyableItem(w, fmt.Sprintf("%s 转毫秒时间戳", layout), strconv.FormatInt(t.UnixMilli(), 10))
		utils.NewCopyableItem(w, fmt.Sprintf("%s 转微秒时间戳", layout), strconv.FormatInt(t.UnixMicro(), 10))
		utils.NewCopyableItem(w, fmt.Sprintf("%s 转纳秒时间戳", layout), strconv.FormatInt(t.UnixNano(), 10))
		utils.NewCopyableItem(w, fmt.Sprintf("%s 转游标", layout), strconv.FormatInt(utils.GetIDFromTime(t), 10))
		utils.NewCopyableItem(w, fmt.Sprintf("%s 转随机 ID", layout), strconv.FormatInt(utils.GenID(t), 10))
	}

	var executes = []func() error{
		func() error {
			// 2006-01-02 15:04:05
			const layout = "2006-01-02 15:04:05"
			t, err := time.ParseInLocation(layout, param, location)
			if err != nil {
				return err
			}
			convert(layout, t)
			return nil
		},
		func() error {
			// 2006/01/02 15:04:05
			const layout = "2006/01/02 15:04:05"
			t, err := time.ParseInLocation(layout, param, location)
			if err != nil {
				return err
			}
			convert(layout, t)
			return nil
		},
		func() error {
			// 2006-01-02 15:04:05.999
			const layout = "2006-01-02 15:04:05.999"
			t, err := time.ParseInLocation(layout, param, location)
			if err != nil {
				return err
			}
			convert(layout, t)
			return nil
		},
		func() error {
			// 2006-01-02 15:04:05.999999
			const layout = "2006-01-02 15:04:05.999999"
			t, err := time.ParseInLocation(layout, param, location)
			if err != nil {
				return err
			}
			convert(layout, t)
			return nil
		},
		func() error {
			// 2006-01-02 15:04:05.999999999
			const layout = "2006-01-02 15:04:05.999999999"
			t, err := time.ParseInLocation(layout, param, location)
			if err != nil {
				return err
			}
			convert(layout, t)
			return nil
		},
		func() error {
			// RFC 3339
			const layout = time.RFC3339
			t, err := time.ParseInLocation(layout, param, location)
			if err != nil {
				return err
			}
			convert(layout, t)
			return nil
		},
	}

	var success = false
	for _, f := range executes {
		if err := f(); err != nil {
			logrus.Infof("transFormatTime error: %v", err)
			continue
		}
		success = true
	}
	if !success {
		return fmt.Errorf("formart time: \"%v\"  failed", param)
	}
	return nil
}
