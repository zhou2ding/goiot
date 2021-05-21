package logtemplate

import "fmt"

func LogProcess(start_time string, cur, cnt int, use_time, end_time string) string {
	ret := fmt.Sprintf("starttime:%v | current/count:%v/%v | use_time:%vs | endtime:%v", start_time, cur, cnt, use_time, end_time)
	return ret
}
