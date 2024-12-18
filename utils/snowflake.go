package utils

import (
	"time"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

func InitSnowfalke(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-04", startTime)
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano()
	node, err = sf.NewNode(machineID)
	return
}

func GenID() int64 {
	return node.Generate().Int64()
}
