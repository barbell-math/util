package csv

import (
	"fmt"
	"testing"

	"github.com/barbell-math/util/algo/iter"
	"github.com/barbell-math/util/customerr"
	"github.com/barbell-math/util/test"
)

func TestFromStructsNonStruct(t *testing.T){
    cnt,err:=FromStructs[int](
        iter.SliceElems[int]([]int{0,1,2}),
        NewOptions(),
    ).Count();
    test.ContainsError(customerr.IncorrectType,err,t)
    test.Eq(0, cnt,t)
}

func TestFromStructsInvalidStruct(t *testing.T) {
    type Row struct {
        One int `csv:"Two"`
        Two int
    }
    cnt,err:=FromStructs[Row](
        iter.SliceElems[Row]([]Row{
            {One: 1, Two: 2},
        }),
        NewOptions().DateTimeFormat("01/02/2006"),
    ).Count()
    test.ContainsError(MalformedCSVStruct,err,t)
    test.ContainsError(DuplicateColName,err,t)
    test.Eq(0,cnt,t)
}

func TestFromStructsValidStruct(t *testing.T){
    res,err:=FromStructs[csvTest](
        iter.SliceElems[csvTest](VALID_STRUCT),
        NewOptions().DateTimeFormat("01/02/2006"),
    ).Collect()
    fmt.Println(res)
    test.Nil(err,t)
    test.Eq(len(VALID),len(res),t)
    for i:=0; i<len(res); i++ {
        test.Eq(len(VALID[i]),len(res[i]),t)
        for j:=0; j<len(res); j++ {
            test.Eq(VALID[i][j],res[i][j],t)
        }
    }
}

// func TestValidStructToCSV(t *testing.T) {
//     type testType struct {
//         V int;
//         priv int;
//     };
//     structs:=make([]testType,5);
//     for i,_:=range(structs) {
//         structs[i].V=i;
//         structs[i].priv=i+1;
//     }
//     res,err:=StructToCSV(iter.SliceElems(structs),true,"01/02/2006").Collect();
//     test.BasicTest(nil,err,
//         "StructToCSV returned an error when it should not have.",t,
//     );
//     test.BasicTest(len(structs)+1,len(res),
//         "StructToCSV did not produce the correct number of values.",t,
//     );
//     newStructs,err:=CSVToStruct[testType](
//         iter.SliceElems(res),"01/02/2006",
//     ).Collect();
//     test.BasicTest(len(structs),len(newStructs),
//         "StructToCSV -> CSVToStruct did not produce the correct number of values.",t,
//     );
//     for i,v:=range(structs) {
//         if i<len(newStructs) {
//             test.BasicTest(v.V,newStructs[i].V,
//                 "New structs public variable was not correctly set.",t,
//             );
//             test.BasicTest(0,newStructs[i].priv,
//                 "New structs private variable was modified when it should not have been.",t,
//             );
//         }
//     }
// }
// 
// func TestInvalidStructToCSV(t *testing.T) {
//     type testType struct {
//         V []int;
//         priv int;
//     };
//     structs:=make([]testType,5);
//     for i,_:=range(structs) {
//         structs[i].V=make([]int, 5);
//         structs[i].priv=i+1;
//     }
//     _,err:=StructToCSV(iter.SliceElems(structs),true,"01/02/2006").Collect();
//     if !IsUnsupportedType(err) {
//         test.FormatError(UnsupportedType(""),err,
//             "StructToCSV did not return the correct error.",t,
//         );
//     }
// }
// 
// func TestNonStructToCSV(t *testing.T) {
//     _,err:=StructToCSV(iter.SliceElems([]int{1,2,3,4}),true,"01/02/2006").Collect();
//     if !IsNonStructValue(err) {
//         test.FormatError(UnsupportedType(""),err,
//             "StructToCSV did not return the correct error.",t,
//         );
//     }
// }
// 
// func TestValidStructToCSVWithTime(t *testing.T) {
//     structs:=make([]csvTest,5);
//     for i,_:=range(structs) {
//         structs[i].B=true;
//         structs[i].T=time.Now();
//         structs[i].S="test string";
//         structs[i].Ui=uint(i);
//     }
//     res,err:=StructToCSV(iter.SliceElems(structs),true,"01/02/2006").Collect();
//     test.BasicTest(nil,err,
//         "StructToCSV returned an error when it should not have.",t,
//     );
//     test.BasicTest(len(structs)+1,len(res),
//         "StructToCSV did not produce the correct number of values.",t,
//     );
//     newStructs,err:=CSVToStruct[csvTest](
//         iter.SliceElems(res),"01/02/2006",
//     ).Collect();
//     test.BasicTest(len(structs),len(newStructs),
//         "StructToCSV -> CSVToStruct did not produce the correct number of values.",t,
//     );
//     for i,v:=range(structs) {
//         if i<len(newStructs) {
//             test.BasicTest(v.B,newStructs[i].B,
//                 "New structs public boolean variable was not correctly set.",t,
//             );
//             test.BasicTest(v.S,newStructs[i].S,
//                 "New structs public string variable was not correctly set.",t,
//             );
//             test.BasicTest(v.Ui,newStructs[i].Ui,
//                 "New structs public unsigned int variable was not correctly set.",t,
//             );
//             day,month,year:=v.T.Date();
//             day1,month1,year1:=newStructs[i].T.Date();
//             test.BasicTest(day,day1,
//                 "The day of the structs public time variable was not correctly set.",t,
//             );
//             test.BasicTest(month,month1,
//                 "The month of the structs public time variable was not correctly set.",t,
//             );
//             test.BasicTest(year,year1,
//                 "The year of the structs public time variable was not correctly set.",t,
//             );
//         }
//     }
// }
