package service

import (
	"strconv"
	"time"
)

func GetTime() string {
	startTime, _ := time.Parse("2006-01-02", "2020-04-14")
	wenyi, _ := time.Parse("2006-01-02", "2022-09-28")
	fan, _ := time.Parse("2006-01-02", "2022-12-31")
	days := time.Since(startTime).Hours() / 24
	wenyiDays := wenyi.Sub(time.Now()).Hours() / 24
	fanDays := fan.Sub(time.Now()).Hours() / 24

	return time.Now().Format("2006-01-02 15:04:05") +
		"\n范范爱文艺的第" + strconv.Itoa(int(days)) + "天" +
		"\n距离文艺生日还有" + strconv.Itoa(int(wenyiDays)) + "天" +
		"\n距离范范生日还有" + strconv.Itoa(int(fanDays)) + "天"
}
