package csv

import (
	"fmt"
	"testing"
	"time"

	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/test"
)

type (
    csvTest struct {
        I int;
        I8 int8;
        Ui uint;
        Ui8 uint8;
        F32 float32;
        S string;
        S1 string;
        B bool;
        T time.Time;
    };
)

const (
    ValidCSV int=iota
    StringsCSV
    MissingColumns
    MissingValues
    MissingHeaders
    MalformedDifferingRowLen
    MalformedDifferingRowLenNoHeaders
    MalformedInt
    MalformedUint
    MalformedFloat
    MalformedTime
    ValidCSVTemplate
    ValidCSVTemplateStringVals
)

var (
    TEST_EXP_PARSE_RES map[int][][]string=map[int][][]string{
        ValidCSV: VALID_PARSE,
        StringsCSV: STRINGS_PARSE,
        MissingColumns: MISSING_COLUMNS_PARSE,
        MissingValues: MISSING_VALUES_PARSE,
        MissingHeaders: MISSING_HEADERS_PARSE,
        MalformedDifferingRowLen: MALFORMED_DIFFERING_ROW_LEN_PARSE,
        MalformedDifferingRowLenNoHeaders: MALFORMED_DIFFERING_ROW_LEN_NO_HEADERS_PARSE,
        MalformedInt: MALFORMED_INT,
        MalformedUint: MALFORMED_UINT,
        MalformedFloat: MALFORMED_FLOAT,
        MalformedTime: MALFORMED_TIME,
    }

    TEST_EXP_FROM_STRUCTS_RES map[int][][]string=map[int][][]string{
        ValidCSV: VALID_FROM_STRUCT,
        StringsCSV: STRINGS_FROM_STRUCT,
        MissingColumns: MISSING_COLUMNS_FROM_STRUCT,
        MissingValues: MISSING_VALUES_FROM_STRUCT,
        MissingHeaders: MISSING_HEADERS_FROM_STRUCT,
        MalformedDifferingRowLen: MALFORMED_DIFFERING_ROW_LEN_FROM_STRUCT,
        MalformedDifferingRowLenNoHeaders: MALFORMED_DIFFERING_ROW_LEN_NO_HEADERS_FROM_STRUCT,
        MalformedInt: MALFORMED_INT,
        MalformedUint: MALFORMED_UINT,
        MalformedFloat: MALFORMED_FLOAT,
        MalformedTime: MALFORMED_TIME,
    }

    TEST_STRUCTS map[int][]csvTest=map[int][]csvTest{
        ValidCSV: VALID_STRUCT,
        StringsCSV: STRINGS_STRUCT,
        MissingColumns: MISSING_COLUMNS_STRUCT,
        MissingValues: MISSING_VALUES_STRUCT,
        MissingHeaders: MISSING_HEADERS_STRUCT,
        MalformedDifferingRowLen: MALFORMED_DIFFERING_ROW_LEN_STRUCT,
        MalformedDifferingRowLenNoHeaders: MALFORMED_DIFFERING_ROW_LEN_NO_HEADERS_STRUCT,
        MalformedInt: MALFORMED_INT_STRUCT,
        MalformedUint: MALFORMED_UINT_STRUCT,
        MalformedFloat: MALFORMED_FLOAT_STRUCT,
        MalformedTime: MALFORMED_TIME_STRUCT,
    }

    TEST_FILES map[int]string=map[int]string{
        ValidCSV: "./testData/ValidCSV.csv",
        StringsCSV: "./testData/Strings.csv",
        MissingColumns: "./testData/MissingColumns.csv",
        MissingValues: "./testData/MissingValues.csv",
        MissingHeaders: "./testData/MissingHeaders.csv",
        MalformedDifferingRowLen: "./testData/MalformedDifferingRowLen.csv",
        MalformedDifferingRowLenNoHeaders: "./testData/MalformedDifferingRowLenNoHeaders.csv",
        MalformedInt: "./testData/MalformedInt.csv",
        MalformedUint: "./testData/MalformedUint.csv",
        MalformedFloat: "./testData/MalformedFloat.csv",
        MalformedTime: "./testData/MalformedTime.csv",
    }

    VALID_PARSE [][]string=[][]string{
        {"I","I8","Ui","Ui8","F32","S","S1","B","T"},
        {"1","-2","100","101","1.001","\"str1\"","str1","","12/12/2012"},
        {"2","-3","101","102","1.002","\"str2\"","str2","true","12/13/2012"},
    }
    VALID_FROM_STRUCT [][]string=[][]string{
        {"I","I8","Ui","Ui8","F32","S","S1","B","T"},
        {"1","-2","100","101","1.001","\"\"\"str1\"\"\"","str1","","12/12/2012"},
        {"2","-3","101","102","1.002","\"\"\"str2\"\"\"","str2","true","12/13/2012"},
    }
    VALID_STRUCT []csvTest=[]csvTest{
        {
            I: 1,I8: -2,Ui: 100,Ui8: 101,F32: 1.001,
            S: "\"str1\"",S1: "str1",B: false,
            T: time.Time{}.AddDate(2011,11,11),
        },
        {
            I: 2,I8: -3,Ui: 101,Ui8: 102,F32: 1.002,
            S: "\"str2\"",S1: "str2",B: true,
            T: time.Time{}.AddDate(2011,11,12),
        },
    }

    STRINGS_PARSE [][]string=[][]string{
        {"S","S1"},
        {"hello","world"},
        {"hello,world","hello,world"},
        {"hello\"world",""},
        {"\"hello world\"",""},
        {"hello\nworld",""},
    }
    STRINGS_FROM_STRUCT [][]string=[][]string{
        {"S","S1"},
        {"hello","world"},
        {"\"hello,world\"","\"hello,world\""},
        {"hello\"world",""},
        {"\"\"\"hello world\"\"\"",""},
        {"\"hello\nworld\"",""},
    }
    STRINGS_STRUCT []csvTest=[]csvTest{
        { S: "hello", S1: "world"},
        { S: "hello,world", S1: "hello,world"},
        { S: "hello\"world", S1: ""},
        { S: "\"hello world\"", S1: ""},
        { S: "hello\nworld", S1: ""},
    }

    MISSING_COLUMNS_PARSE [][]string=[][]string{
        {"I8","Ui","Ui8","F32","S","B","T"},
        {"-2","100","101","1.001","\"str1\"","","12/12/2012"},
        {"-3","101","102","1.002","\"str2\"","true","12/13/2012"},
    }
    MISSING_COLUMNS_FROM_STRUCT [][]string=[][]string{
        {"I8","Ui","Ui8","F32","S","B","T"},
        {"-2","100","101","1.001","\"\"\"str1\"\"\"","","12/12/2012"},
        {"-3","101","102","1.002","\"\"\"str2\"\"\"","true","12/13/2012"},
    }
    MISSING_COLUMNS_STRUCT []csvTest=[]csvTest{
        {
            I8: -2,Ui: 100,Ui8: 101,F32: 1.001,S: "\"str1\"",B: false,
            T: time.Time{}.AddDate(2011,11,11),
        },
        {
            I8: -3,Ui: 101,Ui8: 102,F32: 1.002,S: "\"str2\"",B: true,
            T: time.Time{}.AddDate(2011,11,12),
        },
    }

    MISSING_VALUES_PARSE [][]string=[][]string{
        {"I","I8","Ui","Ui8","F32","S","S1","B","T"},
        {"1","","100","101","1.001","\"str1\"","","",""},
        {"2","","101","102","1.002","\"str2\"","","true",""},
    }
    MISSING_VALUES_FROM_STRUCT [][]string=[][]string{
        {"I","I8","Ui","Ui8","F32","S","S1","B","T"},
        {"1","","100","101","1.001","\"\"\"str1\"\"\"","","",""},
        {"2","","101","102","1.002","\"\"\"str2\"\"\"","","true",""},
    }
    MISSING_VALUES_STRUCT []csvTest=[]csvTest{
        {
            I: 1,I8: 0,Ui: 100,Ui8: 101,F32: 1.001,
            S: "\"str1\"",S1: "",B: false,
            T: time.Time{},
        },
        {
            I: 2,I8: 0,Ui: 101,Ui8: 102,F32: 1.002,
            S: "\"str2\"",S1: "",B: true,
            T: time.Time{},
        },
    }

    MISSING_HEADERS_PARSE [][]string=[][]string{
        {"1","-2","100","101","1.001","\"str1\"","str1","","12/12/2012"},
        {"2","-3","101","102","1.002","\"str2\"","str2","true","12/13/2012"},
    }
    MISSING_HEADERS_FROM_STRUCT [][]string=[][]string{
        {"1","-2","100","101","1.001","\"\"\"str1\"\"\"","str1","","12/12/2012"},
        {"2","-3","101","102","1.002","\"\"\"str2\"\"\"","str2","true","12/13/2012"},
    }
    MISSING_HEADERS_STRUCT []csvTest=[]csvTest{
        {
            I: 1,I8: -2,Ui: 100,Ui8: 101,F32: 1.001,
            S: "\"str1\"",S1: "str1",B: false,
            T: time.Time{}.AddDate(2011,11,11),
        },
        {
            I: 2,I8: -3,Ui: 101,Ui8: 102,F32: 1.002,
            S: "\"str2\"",S1: "str2",B: true,
            T: time.Time{}.AddDate(2011,11,12),
        },
    }

    MALFORMED_DIFFERING_ROW_LEN_PARSE [][]string=[][]string{
        {"I","I8","Ui","Ui8","F32","S","S1","B","T"},
        {"1","-2","100","101","1.001","\"str1\"","str1","false","12/12/2012"},
        {"2","-3","101","102","1.002","\"str2\"","str2","true"},
    }
    MALFORMED_DIFFERING_ROW_LEN_FROM_STRUCT [][]string=[][]string{
        {"I","I8","Ui","Ui8","F32","S","S1","B","T"},
        {"1","-2","100","101","1.001","\"\"\"str1\"\"\"","str1","false","12/12/2012"},
        {"2","-3","101","102","1.002","\"\"\"str2\"\"\"","str2","true"},
    }
    MALFORMED_DIFFERING_ROW_LEN_STRUCT []csvTest=[]csvTest{
        {
            I: 1,I8: -2,Ui: 100,Ui8: 101,F32: 1.001,
            S: "\"str1\"",S1: "str1",B: false,
            T: time.Time{}.AddDate(2011,11,11),
        },
    }

    MALFORMED_DIFFERING_ROW_LEN_NO_HEADERS_PARSE [][]string=[][]string{
        {"1","-2","100","101","1.001","\"str1\"","str1","false","12/12/2012"},
        {"2","-3","101","102","1.002","\"str2\"","str2","true"},
    }
    MALFORMED_DIFFERING_ROW_LEN_NO_HEADERS_FROM_STRUCT [][]string=[][]string{
        {"1","-2","100","101","1.001","\"\"\"str1\"\"\"","str1","false","12/12/2012"},
        {"2","-3","101","102","1.002","\"\"\"str2\"\"\"","str2","true"},
    }
    MALFORMED_DIFFERING_ROW_LEN_NO_HEADERS_STRUCT []csvTest=[]csvTest{
        {
            I: 1,I8: -2,Ui: 100,Ui8: 101,F32: 1.001,
            S: "\"str1\"",S1: "str1",B: false,
            T: time.Time{}.AddDate(2011,11,11),
        },
    }

    MALFORMED_INT [][]string=[][]string{{"I"},{"str"}}
    MALFORMED_INT_STRUCT []csvTest=[]csvTest{{}}

    MALFORMED_UINT [][]string=[][]string{{"Ui"},{"-10"}}
    MALFORMED_UINT_STRUCT []csvTest=[]csvTest{{}}

    MALFORMED_FLOAT [][]string=[][]string{{"F32"},{"1.1.0"}}
    MALFORMED_FLOAT_STRUCT []csvTest=[]csvTest{{}}

    MALFORMED_TIME [][]string=[][]string{{"T"},{"13/12/2000"}}
    MALFORMED_TIME_STRUCT []csvTest=[]csvTest{{}}
)

func (c *csvTest)Eq(other *csvTest, t *testing.T) {
    test.Eq(c.I,other.I,t);
    test.Eq(int8(c.I8),other.I8,t)
    test.Eq(uint(c.Ui),other.Ui,t)
    test.Eq(uint8(c.Ui8),other.Ui8,t)
    test.Eq(c.S,other.S,t)
    test.Eq(c.S1,other.S1,t)
    test.Eq(c.B,other.B,t)
    test.Eq(c.T,other.T,t)
}

func TestFlatten(t *testing.T) {
    res,err:=Flatten(iter.SliceElems([][]string{}),NewOptions()).Collect();
    test.Nil(err,t)
    test.Eq(0,len(res),t)
    res,err=Flatten(iter.SliceElems([][]string{
        {"1"},
        {"2"},
        {"3"},
        {"4"},
    }),NewOptions()).Collect();
    test.Nil(err,t)
    for i,v:=range([]string{"1","2","3","4"}) {
        test.Eq(v,res[i],t)
    }
    res,err=Flatten(iter.SliceElems([][]string{
        {"1","2","3"},
        {"2","3","4"},
        {"3","4","5"},
        {"4","5","6"},
    }),NewOptions()).Collect();
    test.Nil(err,t)
    for i,v:=range([]string{"1,2,3","2,3,4","3,4,5","4,5,6"}) {
        test.Eq(v,res[i],t)
    }
}

func TestParse(t *testing.T) {
    helper:=func(fileName string, expRes [][]string) {
        res,err:=Parse(fileName,NewOptions()).Collect()
        test.Nil(err,t)
        test.Eq(len(expRes),len(res),t)
        for i:=0; i<len(res); i++ {
            test.Eq(len(expRes[i]),len(res[i]),t)
            for j:=0; j<len(res[i]); j++ {
                test.Eq(expRes[i][j],res[i][j],t)
            }
        }
    }
    for k,f:=range(TEST_FILES) {
        if k!=MalformedDifferingRowLen && k!=MalformedDifferingRowLenNoHeaders {
            helper(f,TEST_EXP_PARSE_RES[k])
        } else {
            _,err:=Parse(f,NewOptions()).Collect()
            test.NotNil(err,t)
        }
    }
}

func TestParseSkipHeaders(t *testing.T){
    err:=Parse("./testData/ValidCSVTemplate.csv",NewOptions()).
        Skip(1).
        ForEach(func(index int, val []string) (iter.IteratorFeedback, error) {
            for i,v:=range(val) {
                test.Eq(fmt.Sprintf("%d",index+i+1),v,t)
            }
            return iter.Continue,nil;
        },
    )
    test.Nil(err,t)
}

func TestParseWithStrings(t *testing.T){
    err:=Parse("./testData/ValidCSVTemplateStringVals.csv",NewOptions()).
        Take(1).
        ForEach(func(index int, val []string) (iter.IteratorFeedback, error) {
            for i,v:=range(val) {
                test.Eq(fmt.Sprintf("Column%d,\"Column%d\"",i+1,i+2),v,t)
            }
            return iter.Continue,nil;
        },
    );
    test.Nil(err,t)
}

func TestIntegrationTest(t *testing.T) {
    res,err:=Flatten(
        FromStructs[csvTest](
            ToStructs[csvTest](
                Parse(TEST_FILES[ValidCSV],NewOptions()),
                NewOptions().DateTimeFormat("01/02/2006"),
            ),
            NewOptions().DateTimeFormat("01/02/2006"),
        ),
        NewOptions().DateTimeFormat("01/02/2006"),
    ).Reduce(
        "",
        func(accum *string, iter string) error { *accum+=iter; return nil },
    )
    test.Nil(err,t)
    exp,err:=iter.FileLines(TEST_FILES[ValidCSV]).Reduce(
        "",
        func(accum *string, iter string) error { *accum+=iter; return nil },
    )
    test.Nil(err,t)
    test.Eq(exp,res,t)
}

// func TestIntegratedTestQuotes(t *testing.T) {
//     type Row struct {
//         Str string
//     }
//     o:=NewOptions()
//     res,err:=Flatten(
//         FromStructs[Row](
//             ToStructs[Row](
//                 iter.SliceElems[[]string](
//                     [][]string{
//                         {"Str"},
//                         {"hello world"},                // Normal str
//                         {"\"hello world\""},            // Quotes will be removed
//                         {"\"\"\"hello world\"\"\""},    // Quotes will be kept
// 
//                     },
//                     o,
//                 )
//             ),
//             o
//         ),
//         o
//     )
// }
