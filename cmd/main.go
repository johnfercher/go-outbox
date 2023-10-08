package main

import (
	"fmt"
	"github.com/go-mysql-org/go-mysql/canal"
	"github.com/johnfercher/go-outbox/internal/binlogparser"
	"github.com/johnfercher/go-outbox/internal/config"
	"os"
)

func main() {
	cfg, err := config.Load(os.Args)
	if err != nil {
		panic(err)
	}

	binLogListener(cfg)
}

func binLogListener(cfg *config.Config) {
	c, err := getDefaultCanal(cfg)
	if err == nil {
		coords, err := c.GetMasterPos()
		if err == nil {
			c.SetEventHandler(&binlogHandler{})
			c.RunFrom(coords)
		}
	}
}
func getDefaultCanal(config *config.Config) (*canal.Canal, error) {
	cfg := canal.NewDefaultConfig()
	cfg.Addr = config.Mysql.Url
	cfg.User = config.Mysql.User
	cfg.Password = config.Mysql.Password
	cfg.Flavor = "mariadb"
	cfg.Dump.ExecutionPath = ""

	return canal.NewCanal(cfg)
}

type binlogHandler struct {
	canal.DummyEventHandler   // Dummy handler from external lib
	binlogparser.BinlogParser // Our custom helper
}

func (h *binlogHandler) OnRow(e *canal.RowsEvent) error {
	var n int //starting value
	var k int // step
	switch e.Action {
	case canal.DeleteAction:
		return nil // not covered in example
	case canal.UpdateAction:
		n = 1
		k = 2
	case canal.InsertAction:
		n = 0
		k = 1
	}
	for i := n; i < len(e.Rows); i += k {
		key := e.Table.Schema + "." + e.Table.Name
		fmt.Println(key)
	}
	return nil
}
func (h *binlogHandler) String() string { return "binlogHandler" }
