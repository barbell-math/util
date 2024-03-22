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
    TEST_EXP_RES map[int][][]string=map[int][][]string{
        ValidCSV: VALID,
        MissingColumns: MISSING_COLUMNS,
        MissingValues: MISSING_VALUES,
        MissingHeaders: MISSING_HEADERS,
        MalformedDifferingRowLen: MALFORMED_DIFFERING_ROW_LEN,
        MalformedDifferingRowLenNoHeaders: MALFORMED_DIFFERING_ROW_LEN_NO_HEADERS,
        MalformedInt: MALFORMED_INT,
        MalformedUint: MALFORMED_UINT,
        MalformedFloat: MALFORMED_FLOAT,
        MalformedTime: MALFORMED_TIME,
        ValidCSVTemplate: VALID_CSV_TEMPLATE,
        ValidCSVTemplateStringVals: VALID_CSV_TEMPLATE_STRING_VALS,
    }

    TEST_FILES map[int]string=map[int]string{
        ValidCSV: "./testData/ValidCSV.csv",
        MissingColumns: "./testData/MissingColumns.csv",
        MissingValues: "./testData/MissingValues.csv",
        MissingHeaders: "./testData/MissingHeaders.csv",
        MalformedDifferingRowLen: "./testData/MalformedDifferingRowLen.csv",
        MalformedDifferingRowLenNoHeaders: "./testData/MalformedDifferingRowLenNoHeaders.csv",
        MalformedInt: "./testData/MalformedInt.csv",
        MalformedUint: "./testData/MalformedUint.csv",
        MalformedFloat: "./testData/MalformedFloat.csv",
        MalformedTime: "./testData/MalformedTime.csv",
        ValidCSVTemplate: "./testData/ValidCSVTemplate.csv",
        ValidCSVTemplateStringVals: "./testData/ValidCSVTemplateStringVals.csv",
    }

    TEST_STRUCTS map[int][]csvTest=map[int][]csvTest{
        ValidCSV: VALID_STRUCT,
    }

    VALID [][]string=[][]string{
        {"I","I8","Ui","Ui8","F32","S","S1","B","T"},
        {"1","-2","100","101","1.001","\"str1\"","str1","false","12/12/2012"},
        {"2","-3","101","102","1.002","\"str2\"","str2","true","12/13/2012"},
    }
    VALID_STRUCT []csvTest=[]csvTest{
        {
            I: 1,I8: -2,Ui: 100,Ui8: 101,F32: 1.001,
            S: "\"str1\"",S1: "str1",B: false,
            T: time.Time{}.AddDate(2012,12,12),
        },
        {
            I: 2,I8: -3,Ui: 101,Ui8: 102,F32: 1.002,
            S: "\"str2\"",S1: "str2",B: true,
            T: time.Time{}.AddDate(2012,12,13),
        },
    }

    MISSING_COLUMNS [][]string=[][]string{
        {"I8","Ui","Ui8","F32","S","B","T"},
        {"-2","100","101","1.001","\"str1\"","false","12/12/2012"},
        {"-3","101","102","1.002","\"str2\"","true","12/13/2012"},
    }

    MISSING_VALUES [][]string=[][]string{
        {"I","I8","Ui","Ui8","F32","S","S1","B","T"},
        {"1","","100","101","1.001","\"str1\"","","false",""},
        {"2","","101","102","1.002","\"str2\"","","true",""},
    }

    MISSING_HEADERS [][]string=[][]string{
        {"1","-2","100","101","1.001","\"str1\"","str1","false","12/12/2012"},
        {"2","-3","101","102","1.002","\"str2\"","str2","true","12/13/2012"},
    }

    MALFORMED_DIFFERING_ROW_LEN [][]string=[][]string{
        {"I","I8","Ui","Ui8","F32","S","S1","B","T"},
        {"1","-2","100","101","1.001","\"str1\"","str1","false","12/12/2012"},
        {"2","-3","101","102","1.002","\"str2\"","str2","true"},
    }

    MALFORMED_DIFFERING_ROW_LEN_NO_HEADERS [][]string=[][]string{
        {"1","-2","100","101","1.001","\"str1\"","str1","false","12/12/2012"},
        {"2","-3","101","102","1.002","\"str2\"","str2","true"},
    }

    MALFORMED_INT [][]string=[][]string{{"I"},{"str"}}
    MALFORMED_UINT [][]string=[][]string{{"Ui"},{"-10"}}
    MALFORMED_FLOAT [][]string=[][]string{{"F32"},{"1.1.0"}}
    MALFORMED_TIME [][]string=[][]string{{"T"},{"13/12/2000"}}

    VALID_CSV_TEMPLATE [][]string=[][]string{
        {"Column1","Column2","Column3","Column4","Column5"},
        {"1","2","3","4","5"},
        {"2","3","4","5","6"},
        {"3","4","5","6","7"},
        {"4","5","6","7","8"},
        {"5","6","7","8","9"},
    }

    VALID_CSV_TEMPLATE_STRING_VALS [][]string=[][]string{
        {"Column1,\"Column2\"","Column2,\"Column3\"","Column3,\"Column4\""},
        {"1","2","3"},
        {"2","3","4"},
        {"3","4","5"},
        {"4","5","6"},
        {"5","6","7"},
    }
)

// func TestFlatten(t *testing.T) {
//     res,err:=Flatten(iter.SliceElems([][]string{}),",").Collect();
//     test.BasicTest(nil,err,
//         "Flatten returned an error when it was not supposed to",t,
//     );
//     test.BasicTest(0,len(res),
//         "Flatten did not produce the correct value.",t,
//     );
//     test.BasicTest(nil,err,
//         "Flatten returned an error when it was not supposed to",t,
//     );
//     res,err=Flatten(iter.SliceElems([][]string{
//         {"1"},
//         {"2"},
//         {"3"},
//         {"4"},
//     }),",").Collect();
//     test.BasicTest(nil,err,
//         "Flatten returned an error when it was not supposed to",t,
//     );
//     for i,v:=range([]string{"1","2","3","4"}) {
//         test.BasicTest(v,res[i],
//             "Flatten did not produce the correct value.",t,
//         );
//     }
//     res,err=Flatten(iter.SliceElems([][]string{
//         {"1","2","3"},
//         {"2","3","4"},
//         {"3","4","5"},
//         {"4","5","6"},
//     }),",").Collect();
//     test.BasicTest(nil,err,
//         "Flatten returned an error when it was not supposed to",t,
//     );
//     for i,v:=range([]string{"1,2,3","2,3,4","3,4,5","4,5,6"}) {
//         test.BasicTest(v,res[i],
//             "Flatten did not produce the correct value.",t,
//         );
//     }
// }

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
            helper(f,TEST_EXP_RES[k])
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
