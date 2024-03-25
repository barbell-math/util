package log

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/algo/widgets"
	"github.com/barbell-math/util/container/basic"
	"github.com/barbell-math/util/container/containers"
	"github.com/barbell-math/util/container/staticContainers"
	"github.com/barbell-math/util/test"
)

func generateLog(l *ValueLogger[int], numLines int) {
	for i := 0; i < numLines; i++ {
		l.Log(i, "Line %d", i)
	}
}

func generateIntertwinedLogs(
	l1 *ValueLogger[int],
	l2 *ValueLogger[int],
	numLines int,
) {
	cntr := 0
	for i := 0; i < numLines; i++ {
		cntr++
		l1.Log(cntr, "L1 Line %d", cntr)
		cntr++
		l2.Log(cntr, "L2 Line %d", cntr)
	}
}

func TestNewValueLoggerBadPath(t *testing.T) {
	_, err := NewValueLogger[int](Error, "./non/existant/path/to/file.txt", NewOptions())
	test.ContainsError(os.ErrNotExist, err, t)
}

func TestLogIteration(t *testing.T) {
	l, _ := NewValueLogger[int](Error, "./testData/generateLog.log", NewOptions())
	generateLog(&l, 1000)
	l.LogElems().ForEach(
		func(index int, val LogEntry[int]) (iter.IteratorFeedback, error) {
			test.Eq(Error, val.Status, t)
			test.Eq(fmt.Sprintf("Line %d", index), val.Message, t)
			test.Eq(index, val.Val, t)
			return iter.Continue, nil
		},
	)
	l.Close()
}

func TestLogIterationTime(t *testing.T) {
	l, _ := NewValueLogger[int](
		Error,
		"./testData/generateLog.log",
		NewOptions().LogFlags(
			log.Ldate|log.Ltime|log.Lmicroseconds,
		).DateTimeFormat(
			"2006/01/02 15:04:05.000000",
		),
	)
	generateLog(&l, 1000)
	cntr := 0
	c, _ := containers.NewCircularBuffer[
		LogEntry[int],
		widgets.NilWidget[LogEntry[int]],
	](2)
	err := containers.Window[LogEntry[int]](
		l.LogElems(),
		&c,
		false,
	).ForEach(func(
		index int,
		q staticContainers.Vector[LogEntry[int]],
	) (iter.IteratorFeedback, error) {
		cntr++
		f, _ := q.Get(0)
		s, _ := q.Get(1)
		test.True(f.Time.Before(s.Time), t)
		return iter.Continue, nil
	})
	test.Nil(err, t)
	test.Eq(999, cntr, t)
	l.Close()
}

func TestLogAppend(t *testing.T) {
	l, _ := NewValueLogger[int](Error, "./testData/generateLog.log", NewOptions())
	generateLog(&l, 1000)
	l.Close()
	l, _ = NewValueLogger[int](Error, "./testData/generateLog.log", NewOptions().Append(true))
	generateLog(&l, 1000)
	cntr := 0
	err := l.LogElems().ForEach(
		func(index int, val LogEntry[int]) (iter.IteratorFeedback, error) {
			test.Eq(Error, val.Status, t)
			test.Eq(fmt.Sprintf("Line %d", index%1000), val.Message, t)
			test.Eq(index%1000, val.Val, t)
			cntr++
			return iter.Continue, nil
		},
	)
	test.Nil(err, t)
	test.Eq(2000, cntr, t)
	l.Close()
}

func TestLogJoin(t *testing.T) {
	cntr := 0
	l1, _ := NewValueLogger[int](
		Error,
		"./testData/generateLog.part1.log",
		NewOptions().LogFlags(
			log.Ldate|log.Ltime|log.Lmicroseconds,
		).DateTimeFormat(
			"2006/01/02 15:04:05.000000",
		),
	)
	l2, _ := NewValueLogger[int](
		Error,
		"./testData/generateLog.part2.log",
		NewOptions().LogFlags(
			log.Ldate|log.Ltime|log.Lmicroseconds,
		).DateTimeFormat(
			"2006/01/02 15:04:05.000000",
		),
	)
	generateIntertwinedLogs(&l1, &l2, 1000)
	l1S, _ := NewValueLogger[int](
		Error,
		"./testData/generateLog.part1.log",
		NewOptions().Append(true).LogFlags(
			log.Ldate|log.Ltime|log.Lmicroseconds,
		).DateTimeFormat(
			"2006/01/02 15:04:05.000000",
		),
	)
	l2S, _ := NewValueLogger[int](
		Error,
		"./testData/generateLog.part2.log",
		NewOptions().Append(true).LogFlags(
			log.Ldate|log.Ltime|log.Lmicroseconds,
		).DateTimeFormat(
			"2006/01/02 15:04:05.000000",
		),
	)
	err := iter.JoinSame[LogEntry[int]](
		l1S.LogElems(),
		l2S.LogElems(),
		basic.NewVariant[LogEntry[int], LogEntry[int]],
		JoinLogByTimeInc[int, int],
	).ForEach(
		func(index int, val LogEntry[int]) (iter.IteratorFeedback, error) {
			test.Eq(Error, val.Status, t)
			test.Eq(fmt.Sprintf("L%d Line %d", index%2+1, index+1), val.Message, t)
			test.Eq(index+1, val.Val, t)
			cntr++
			return iter.Continue, nil
		},
	)
	test.Nil(err, t)
	test.Eq(2000, cntr, t)
	l1.Close()
	l2.Close()
}

func BenchmarkJoinLog(b *testing.B) {
	l1, _ := NewValueLogger[int](
		Error,
		"./testData/generateLog.part1.log",
		NewOptions().LogFlags(
			log.Ldate|log.Ltime|log.Lmicroseconds,
		).DateTimeFormat(
			"2006/01/02 15:04:05.000000",
		),
	)
	l2, _ := NewValueLogger[int](
		Error,
		"./testData/generateLog.part2.log",
		NewOptions().LogFlags(
			log.Ldate|log.Ltime|log.Lmicroseconds,
		).DateTimeFormat(
			"2006/01/02 15:04:05.000000",
		),
	)
	generateIntertwinedLogs(&l1, &l2, 1000)
	l1S, _ := NewValueLogger[int](
		Error,
		"./testData/generateLog.part1.log",
		NewOptions().Append(true).LogFlags(
			log.Ldate|log.Ltime|log.Lmicroseconds,
		).DateTimeFormat(
			"2006/01/02 15:04:05.000000",
		),
	)
	l2S, _ := NewValueLogger[int](
		Error,
		"./testData/generateLog.part2.log",
		NewOptions().Append(true).LogFlags(
			log.Ldate|log.Ltime|log.Lmicroseconds,
		).DateTimeFormat(
			"2006/01/02 15:04:05.000000",
		),
	)
	for i := 0; i < b.N; i++ {
		iter.JoinSame[LogEntry[int]](
			l1S.LogElems(),
			l2S.LogElems(),
			basic.NewVariant[LogEntry[int], LogEntry[int]],
			JoinLogByTimeInc[int, int],
		).ForEach(
			func(index int, val LogEntry[int]) (iter.IteratorFeedback, error) {
				return iter.Continue, nil
			})
	}
	l1.Close()
	l2.Close()
}

func BenchmarkBigLog(b *testing.B) {
	cntr := 0
	l, _ := NewValueLogger[int](Error, "./testData/big.log", NewOptions())
	fmt.Println("Generating log file with 500K lines (this may take a while)...")
	generateLog(&l, 500000)
	fmt.Println("Running tests...")
	for i := 0; i < b.N; i++ {
		l.LogElems().ForEach(
			func(index int, val LogEntry[int]) (iter.IteratorFeedback, error) {
				cntr++
				return iter.Continue, nil
			})
		fmt.Printf("Processed %d lines.\n", cntr)
	}
	l.Close()
}
