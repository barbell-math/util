package circularImportTests

import (
    "fmt"
    "testing"

    "github.com/barbell-math/util/algo/iter"
    "github.com/barbell-math/util/dataStruct"
    staticType "github.com/barbell-math/util/dataStruct/types/static"
    "github.com/barbell-math/util/test"
)

func TestZipBothNoElements(t *testing.T){
    count,err:=iter.Zip[int,string](
        iter.NoElem[int](),
        iter.NoElem[string](),
        func() staticType.Pair[int, string] { return &dataStruct.Pair[int,string]{} },
    ).Count()
    test.BasicTest(0,count,
        "Zip seemingly iterated over nothing.",t,
    )
    test.BasicTest(nil,err,
        "Zip returned an error when it should not have.",t,
    )
}

func TestZipEmptyLeftAndNonEmptyRight(t *testing.T){
    count,err:=iter.Zip[int,string](
        iter.SliceElems[int]([]int{0,1,2}),
        iter.NoElem[string](),
        func() staticType.Pair[int, string] { return &dataStruct.Pair[int,string]{} },
    ).Count()
    test.BasicTest(0,count,
        "Zip seemingly iterated over nothing.",t,
    )
    test.BasicTest(nil,err,
        "Zip returned an error when it should not have.",t,
    )
}

func TestZipEmptyRightAndNonEmptyLeft(t *testing.T){
    count,err:=iter.Zip[string,int](
        iter.NoElem[string](),
        iter.SliceElems[int]([]int{0,1,2}),
        func() staticType.Pair[string, int] { return &dataStruct.Pair[string,int]{} },
    ).Count()
    test.BasicTest(0,count,
        "Zip seemingly iterated over nothing.",t,
    )
    test.BasicTest(nil,err,
        "Zip returned an error when it should not have.",t,
    )
}

func TestZipRightLessThanLeft(t *testing.T){
    vals,err:=iter.Zip[string,int](
        iter.SliceElems[string]([]string{"0","1"}),
        iter.SliceElems[int]([]int{0,1,2}),
        func() staticType.Pair[string, int] { return &dataStruct.Pair[string,int]{} },
    ).Collect()
    test.BasicTest(2,len(vals),
        "Zip iterated over the wrong number of values.",t,
    )
    test.BasicTest(nil,err,
        "Zip returned an error when it should not have.",t,
    )
    for i:=0; i<2; i++ {
        test.BasicTest(fmt.Sprintf("%d",i),vals[i].GetA(),
            "Zip did not coalesce the values correctly.",t,
        )
        test.BasicTest(i,vals[i].GetB(),
            "Zip did not coalesce the values correctly.",t,
        )
    }
}

func TestZipLeftLessThanRight(t *testing.T){
    vals,err:=iter.Zip[string,int](
        iter.SliceElems[string]([]string{"0","1","2"}),
        iter.SliceElems[int]([]int{0,1}),
        func() staticType.Pair[string, int] { return &dataStruct.Pair[string,int]{} },
    ).Collect()
    test.BasicTest(2,len(vals),
        "Zip iterated over the wrong number of values.",t,
    )
    test.BasicTest(nil,err,
        "Zip returned an error when it should not have.",t,
    )
    for i:=0; i<2; i++ {
        test.BasicTest(fmt.Sprintf("%d",i),vals[i].GetA(),
            "Zip did not coalesce the values correctly.",t,
        )
        test.BasicTest(i,vals[i].GetB(),
            "Zip did not coalesce the values correctly.",t,
        )
    }
}

func TestZipLeftEqualsRight(t *testing.T){
    vals,err:=iter.Zip[string,int](
        iter.SliceElems[string]([]string{"0","1","2"}),
        iter.SliceElems[int]([]int{0,1,2}),
        func() staticType.Pair[string, int] { return &dataStruct.Pair[string,int]{} },
    ).Collect()
    test.BasicTest(3,len(vals),
        "Zip iterated over the wrong number of values.",t,
    )
    test.BasicTest(nil,err,
        "Zip returned an error when it should not have.",t,
    )
    for i:=0; i<3; i++ {
        test.BasicTest(fmt.Sprintf("%d",i),vals[i].GetA(),
            "Zip did not coalesce the values correctly.",t,
        )
        test.BasicTest(i,vals[i].GetB(),
            "Zip did not coalesce the values correctly.",t,
        )
    }
}
