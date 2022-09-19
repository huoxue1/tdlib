package utils

import (
	"testing"
)

func TestQl_GetConfigFile(t *testing.T) {
	ql := InitQl("http://127.0.0.1:5706", "1241", "ff1S-13213-6Hv")
	file, err := ql.GetConfigFile("task_after.sh")
	if err != nil {
		t.Error(err)
		return
	}
	println(file)
	err = ql.SaveConfigFile("task_after.sh", file+"123")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestQl_GetCrons(t *testing.T) {
	ql := InitQl("http://127.0.0.1:5700", "1213", "ff1S-12414mhxDgQs-6Hv")
	crons, err := ql.GetCrons("")
	if err != nil {
		t.Error(err)
		return
	}
	for _, cron := range crons {
		println(cron.Name)
	}

	cron, err := ql.GetCron(276)
	if err != nil {
		t.Error(err)
		return
	}
	println(cron)
	err = ql.RunCrons(274)
	if err != nil {
		t.Error(err)
		return
	}
}
