package snowflake

import (
	"github.com/bwmarrin/snowflake"
	"time"
	"bluebell/settings"
)

var node *snowflake.Node

func Init() (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-01", settings.Config.SnowflakeConfig.StartTime)
	if err != nil {
		return
	}
	snowflake.Epoch = st.UnixMilli()
	node, err = snowflake.NewNode(settings.Config.SnowflakeConfig.MachineID)
	return
}

func GenID() int64 {
	return node.Generate().Int64()
}