package csv

import (
	"fmt"
	"testing"
	"time"

	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/customerr"
	"github.com/barbell-math/util/test"
)

func TestToStructsNonStruct(t *testing.T){
    cnt,err:=ToStructs[int](
        iter.SliceElems[[]string](VALID),
        NewOptions(),
    ).Count();
    test.ContainsError(customerr.IncorrectType,err,t)
    test.Eq(0, cnt,t)
}

func TestToStructsInvalidStruct(t *testing.T) {
    type Row struct {
        One int `csv:"Two"`
        Two int
    }
    cnt,err:=ToStructs[Row](
        iter.SliceElems[[]string](VALID),
        NewOptions().DateTimeFormat("01/02/2006"),
    ).Count()
    test.ContainsError(MalformedCSVStruct,err,t)
    test.ContainsError(DuplicateColName,err,t)
    test.Eq(0,cnt,t)
}

func TestToStructsValidStruct(t *testing.T){
    cntr:=0;
    baseTime,_:=time.Parse("01/02/2006","12/12/2012");
    err:=ToStructs[csvTest](
        iter.SliceElems[[]string](VALID),
        NewOptions().DateTimeFormat("01/02/2006"),
    ).ForEach(func(index int, val csvTest) (iter.IteratorFeedback, error) {
        test.Eq(index+1,val.I,t);
        test.Eq(int8(-index-2),val.I8,t)
        test.Eq(uint(index+100),val.Ui,t)
        test.Eq(uint8(index+101),val.Ui8,t)
        test.Eq(fmt.Sprintf("str%d",index+1),val.S,t)
        test.Eq(fmt.Sprintf("str%d",index+1),val.S1,t)
        test.Eq(index!=0, val.B,t)
        test.Eq(baseTime.Add(time.Hour*24*time.Duration(index)),val.T,t)
        cntr++;
        return iter.Continue, nil;
    });
    test.Nil(err,t)
    test.Eq(2,cntr,t);
}

func TestToStructsMissingColumns(t *testing.T){
    cntr:=0;
    baseTime,_:=time.Parse("01/02/2006","12/12/2012");
    err:=ToStructs[csvTest](
        iter.SliceElems[[]string](MISSING_COLUMNS),
        NewOptions().DateTimeFormat("01/02/2006"),
    ).ForEach(func(index int, val csvTest) (iter.IteratorFeedback, error) {
        test.Eq(0, val.I,t);
        test.Eq(int8(-index-2),val.I8,t)
        test.Eq(uint(index+100),val.Ui,t)
        test.Eq(uint8(index+101),val.Ui8,t)
        test.Eq(fmt.Sprintf("str%d",index+1),val.S,t)
        test.Eq("",val.S1,t)
        test.Eq(index!=0, val.B,t)
        test.Eq(baseTime.Add(time.Hour*24*(time.Duration(index))),val.T,t)
        cntr++;
        return iter.Continue,nil;
    });
    test.Nil(err,t)
    test.Eq(2,cntr,t);
}

func TestToStructsMissingValues(t *testing.T){
    cntr:=0;
    baseTime,_:=time.Parse("01/02/2006","00/00/0000");
    err:=ToStructs[csvTest](
        iter.SliceElems[[]string](MISSING_VALUES),
        NewOptions().DateTimeFormat("01/02/2006"),
    ).ForEach(func(index int, val csvTest) (iter.IteratorFeedback, error) {
        test.Eq(index+1, val.I,t);
        test.Eq(int8(0),val.I8,t)
        test.Eq(uint(index+100),val.Ui,t)
        test.Eq(uint8(index+101),val.Ui8,t)
        test.Eq(fmt.Sprintf("str%d",index+1),val.S,t)
        test.Eq("",val.S1,t)
        test.Eq(index!=0, val.B,t)
        test.Eq(baseTime,val.T,t)
        cntr++;
        return iter.Continue,nil;
    });
    test.Nil(err,t)
    test.Eq(2,cntr,t);
}

func TestToStructsMissingHeadersWhenNotSpecifedAsMissing(t *testing.T){
    cntr:=0;
    err:=ToStructs[csvTest](
        iter.SliceElems[[]string](MISSING_HEADERS),
        NewOptions(),
    ).ForEach(func(index int, val csvTest) (iter.IteratorFeedback, error) {
        cntr++;
        return iter.Continue,nil;
    });
    test.ContainsError(MalformedCSVFile,err,t)
    test.Eq(0, cntr,t)
}

func TestToStructsMalformedDifferingRowLengths(t *testing.T){
    cntr:=0;
    baseTime,_:=time.Parse("01/02/2006","12/12/2012");
    err:=ToStructs[csvTest](
        iter.SliceElems[[]string](MALFORMED_DIFFERING_ROW_LEN),
        NewOptions().DateTimeFormat("01/02/2006"),
    ).ForEach(func(index int, val csvTest) (iter.IteratorFeedback, error) {
        test.Eq(index+1,val.I,t);
        test.Eq(int8(-index-2),val.I8,t)
        test.Eq(uint(index+100),val.Ui,t)
        test.Eq(uint8(index+101),val.Ui8,t)
        test.Eq(fmt.Sprintf("str%d",index+1),val.S,t)
        test.Eq(fmt.Sprintf("str%d",index+1),val.S1,t)
        test.Eq(index!=0, val.B,t)
        test.Eq(baseTime.Add(time.Hour*24*time.Duration(index)),val.T,t)
        cntr++;
        return iter.Continue, nil;
    });
    test.ContainsError(MalformedCSVFile,err,t)
    test.Eq(1,cntr,t)
}

func TestToStructsMalformedDifferingRowLengthNoHeaders(t *testing.T){
    cntr:=0;
    baseTime,_:=time.Parse("01/02/2006","12/12/2012");
    err:=ToStructs[csvTest](
        iter.SliceElems[[]string](MALFORMED_DIFFERING_ROW_LEN_NO_HEADERS),
        NewOptions().DateTimeFormat("01/02/2006").HasHeaders(false),
    ).ForEach(func(index int, val csvTest) (iter.IteratorFeedback, error) {
        test.Eq(index+1,val.I,t);
        test.Eq(int8(-index-2),val.I8,t)
        test.Eq(uint(index+100),val.Ui,t)
        test.Eq(uint8(index+101),val.Ui8,t)
        test.Eq(fmt.Sprintf("str%d",index+1),val.S,t)
        test.Eq(fmt.Sprintf("str%d",index+1),val.S1,t)
        test.Eq(index!=0, val.B,t)
        test.Eq(baseTime.Add(time.Hour*24*time.Duration(index)),val.T,t)
        cntr++;
        return iter.Continue, nil;
    });
    test.ContainsError(MalformedCSVFile,err,t)
    test.Eq(1,cntr,t)
}

func TestToStructsMalformedTypes(t *testing.T){
    cntr:=0;
    err:=ToStructs[csvTest](
        iter.SliceElems[[]string](MALFORMED_INT),
        NewOptions(),
    ).ForEach(func(index int, val csvTest) (iter.IteratorFeedback, error) {
        cntr++;
        return iter.Continue,nil;
    });
    test.ContainsError(MalformedCSVFile,err,t)
    test.Eq(0, cntr,t)

    err=ToStructs[csvTest](
        iter.SliceElems[[]string](MALFORMED_UINT),
        NewOptions(),
    ).ForEach(func(index int, val csvTest) (iter.IteratorFeedback, error) {
        cntr++;
        return iter.Continue,nil;
    });
    test.ContainsError(MalformedCSVFile,err,t)
    test.Eq(0, cntr,t)

    err=ToStructs[csvTest](
        iter.SliceElems[[]string](MALFORMED_FLOAT),
        NewOptions(),
    ).ForEach(func(index int, val csvTest) (iter.IteratorFeedback, error) {
        cntr++;
        return iter.Continue,nil;
    });
    test.ContainsError(MalformedCSVFile,err,t)
    test.Eq(0, cntr,t)

    err=ToStructs[csvTest](
        iter.SliceElems[[]string](MALFORMED_TIME),
        NewOptions().DateTimeFormat("01/02/2006"),
    ).ForEach(func(index int, val csvTest) (iter.IteratorFeedback, error) {
        cntr++;
        return iter.Continue,nil;
    });
    test.ContainsError(MalformedCSVFile,err,t)
    test.Eq(0, cntr,t)
}
