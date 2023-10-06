package server

import (
	"fmt"
	"log"
	"sync"
)

type Log struct {
	// ミューテックスは相互排他ロックです。
	// ミューテックスのゼロ値は、ロックが解除されたミューテックスです。
	//
	// ミューテックスは最初に使用した後にコピーしてはなりません。
	mu      sync.Mutex
	records []Record
}

func NewLog() *Log {
	return &Log{}
}

func (c *Log) Append(record Record) (uint64, error) {
	// Lock は m *Mutex をロックします。
	// ロックがすでに使用されている場合、呼び出し元のゴルーチン
	// ミューテックスが使用可能になるまでブロックします。
	c.mu.Lock()

	// ロック解除は m のロックを解除します。
	// Ulock へのエントリ時に m がロックされていない場合は、実行時エラーになります。
	defer c.mu.Unlock()

	log.Println("c.records:", c.records)
	log.Println("len(c.records):", len(c.records))
	log.Println("record.Offset:", record.Offset)
	record.Offset = uint64(len(c.records))

	// ここで今回新しうく受け取ったリクエストのレコードを追加している
	c.records = append(c.records, record)
	log.Println("c.records:", c.records)
	log.Println("record.Offset:", record.Offset)
	return record.Offset, nil
}

func (c *Log) Read(offset uint64) (Record, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if offset >= uint64(len(c.records)) {
		return Record{}, ErrOffsetNotFound
	}
	return c.records[offset], nil
}

type Record struct {
	Value  []byte `json:"value"`
	Offset uint64 `json:"offset"`
}

var ErrOffsetNotFound = fmt.Errorf("offset not found")
