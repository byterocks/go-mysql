package canal

import (
    "github.com/go-mysql-org/go-mysql/mysql"
    "github.com/go-mysql-org/go-mysql/replication"
)

type EventHandler interface {
    OnStart(c *Canal) error
    OnRotate(roateEvent *replication.RotateEvent) error
    // OnTableChanged is called when the table is created, altered, renamed or dropped.
    // You need to clear the associated data like cache with the table.
    // It will be called before OnDDL.
    OnTableChanged(schema string, table string) error
    OnDDL(nextPos mysql.Position, queryEvent *replication.QueryEvent) error
    OnRow(e *RowsEvent) error
    OnXID(nextPos mysql.Position, ev *replication.XIDEvent) error
    OnGTID(gtid mysql.GTIDSet, ev *replication.BinlogEvent) error
    // OnPosSynced Use your own way to sync position. When force is true, sync position immediately.
    OnPosSynced(pos mysql.Position, set mysql.GTIDSet, force bool) error
    String() string
}

type DummyEventHandler struct {
}

func (h *DummyEventHandler) OnStart(c *Canal) error                           { return nil }
func (h *DummyEventHandler) OnRotate(*replication.RotateEvent) error          { return nil }
func (h *DummyEventHandler) OnTableChanged(schema string, table string) error { return nil }
func (h *DummyEventHandler) OnDDL(nextPos mysql.Position, queryEvent *replication.QueryEvent) error {
    return nil
}
func (h *DummyEventHandler) OnRow(*RowsEvent) error                                { return nil }
func (h *DummyEventHandler) OnXID(mysql.Position, *replication.XIDEvent) error                            { return nil }
func (h *DummyEventHandler) OnGTID(mysql.GTIDSet, *replication.BinlogEvent) error                            { return nil }
func (h *DummyEventHandler) OnPosSynced(mysql.Position, mysql.GTIDSet, bool) error { return nil }

func (h *DummyEventHandler) String() string { return "DummyEventHandler" }

// `SetEventHandler` registers the sync handler, you must register your
// own handler before starting Canal.
func (c *Canal) SetEventHandler(h EventHandler) {
    c.eventHandler = h
}
