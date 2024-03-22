package csv

import (
	"testing"

	"github.com/barbell-math/util/customerr"
	"github.com/barbell-math/util/test"
)

func TestNewStructMappingNonStruct(t *testing.T) {
    var tmp int
    m,err:=newStructHeaderMapping[int](&tmp,NewOptions())
    test.ContainsError(customerr.IncorrectType,err,t)
    test.Eq(0,len(m),t)
}

func TestNewStructMappingNoTagsNominalCase(t *testing.T) {
    type Row struct {
        One int
        Two int
        Three int
    }
    var tmp Row
    m,err:=newStructHeaderMapping[Row](&tmp,NewOptions())
    test.Nil(err,t)
    test.Eq(3,len(m),t)
    test.Eq(structIndex(0),m["One"],t)
    test.Eq(structIndex(1),m["Two"],t)
    test.Eq(structIndex(2),m["Three"],t)
}

func TestNewStructMappingNoTagsUnexportedFields(t *testing.T) {
    type Row struct {
        One int
        Two int
        Three int
        four int
    }
    var tmp Row
    m,err:=newStructHeaderMapping[Row](&tmp,NewOptions())
    test.Nil(err,t)
    test.Eq(3,len(m),t)
    test.Eq(structIndex(0),m["One"],t)
    test.Eq(structIndex(1),m["Two"],t)
    test.Eq(structIndex(2),m["Three"],t)
}

func TestNewStructMappingWithTagsNominalCase(t *testing.T) {
    type Row struct {
        One int `csv:"one"`
        Two int `csv:"two"`
        Three int `csv:"three"`
    }
    var tmp Row
    m,err:=newStructHeaderMapping[Row](&tmp,NewOptions())
    test.Nil(err,t)
    test.Eq(3,len(m),t)
    test.Eq(structIndex(0),m["one"],t)
    test.Eq(structIndex(1),m["two"],t)
    test.Eq(structIndex(2),m["three"],t)
}

func TestNewStructMappingWithOtherTags(t *testing.T){
    type Row struct {
        One int `csv:"one" json:"foo"`
        Two int `csv:"two" json:"bar"`
        Three int `csv:"three" json:"blah"`
    }
    var tmp Row
    m,err:=newStructHeaderMapping[Row](&tmp,NewOptions())
    test.Nil(err,t)
    test.Eq(3,len(m),t)
    test.Eq(structIndex(0),m["one"],t)
    test.Eq(structIndex(1),m["two"],t)
    test.Eq(structIndex(2),m["three"],t)
}

func TestNewStructMappingWithSomeTags(t *testing.T){
    type Row struct {
        One int `csv:"one"`
        Two int
        Three int `csv:"three"`
    }
    var tmp Row
    m,err:=newStructHeaderMapping[Row](&tmp,NewOptions())
    test.Nil(err,t)
    test.Eq(3,len(m),t)
    test.Eq(structIndex(0),m["one"],t)
    test.Eq(structIndex(1),m["Two"],t)
    test.Eq(structIndex(2),m["three"],t)
}

func TestNewStructMappingWithSomeTagsAndOtherTags(t *testing.T){
    type Row struct {
        One int `csv:"one" json:"foo"`
        Two int `json:"bar"`
        Three int `csv:"three" json:"blah"`
    }
    var tmp Row
    m,err:=newStructHeaderMapping[Row](&tmp,NewOptions())
    test.Nil(err,t)
    test.Eq(3,len(m),t)
    test.Eq(structIndex(0),m["one"],t)
    test.Eq(structIndex(1),m["Two"],t)
    test.Eq(structIndex(2),m["three"],t)
}

func TestNewStructMappingWithDuplicateTags(t *testing.T){
    type Row struct {
        One int `csv:"one"`
        Two int `csv:"one"`
        Three int `csv:"three"`
    }
    var tmp Row
    m,err:=newStructHeaderMapping[Row](&tmp,NewOptions())
    test.ContainsError(DuplicateColName,err,t)
    test.Eq(1,len(m),t)
    test.Eq(structIndex(0),m["one"],t)
}

func TestNewStructMappingWithDuplicateTagsAndValuesForwardDecl(t *testing.T){
    type Row struct {
        One int `csv:"one"`
        Two int `csv:"Three"`
        Three int
    }
    var tmp Row
    m,err:=newStructHeaderMapping[Row](&tmp,NewOptions())
    test.ContainsError(DuplicateColName,err,t)
    test.Eq(2,len(m),t)
    test.Eq(structIndex(0),m["one"],t)
    test.Eq(structIndex(1),m["Three"],t)
}

func TestNewStructMappingWithDuplicateTagsAndValuesBackwardDecl(t *testing.T){
    type Row struct {
        One int `csv:"one"`
        Two int
        Three int `csv:"Two"`
    }
    var tmp Row
    m,err:=newStructHeaderMapping[Row](&tmp,NewOptions())
    test.ContainsError(DuplicateColName,err,t)
    test.Eq(2,len(m),t)
    test.Eq(structIndex(0),m["one"],t)
    test.Eq(structIndex(1),m["Two"],t)
}

func TestNewStructMappingNonDefaultTag(t *testing.T) {
    type Row struct {
        One int `tag:"one"`
        Two int `tag:"two"`
        Three int `tag:"three"`
    }
    var tmp Row
    m,err:=newStructHeaderMapping[Row](&tmp,NewOptions().StructTagName("tag"))
    test.Nil(err,t)
    test.Eq(3,len(m),t)
    test.Eq(structIndex(0),m["one"],t)
    test.Eq(structIndex(1),m["two"],t)
    test.Eq(structIndex(2),m["three"],t)
}

func TestNewStructMappingNonDefaultTagWithDefaultTag(t *testing.T) {
    type Row struct {
        One int `tag:"one" csv:"foo"`
        Two int `tag:"two" csv:"bar"`
        Three int `tag:"three" csv:"blah"`
    }
    var tmp Row
    m,err:=newStructHeaderMapping[Row](&tmp,NewOptions().StructTagName("tag"))
    test.Nil(err,t)
    test.Eq(3,len(m),t)
    test.Eq(structIndex(0),m["one"],t)
    test.Eq(structIndex(1),m["two"],t)
    test.Eq(structIndex(2),m["three"],t)
}

func TestNewStructMappingIgnoreTags(t *testing.T) {
    type Row struct {
        One int `csv:"one"`
        Two int `csv:"two"`
        Three int `csv:"three"`
    }
    var tmp Row
    m,err:=newStructHeaderMapping[Row](&tmp,NewOptions().UseStructTags(false))
    test.Nil(err,t)
    test.Eq(3,len(m),t)
    test.Eq(structIndex(0),m["One"],t)
    test.Eq(structIndex(1),m["Two"],t)
    test.Eq(structIndex(2),m["Three"],t)
}

func TestNewStructMappingIgnoreTagsAndNonDefaultTag(t *testing.T) {
    type Row struct {
        One int `tag:"one"`
        Two int `tag:"two"`
        Three int `tag:"three"`
    }
    var tmp Row
    m,err:=newStructHeaderMapping[Row](
        &tmp,
        NewOptions().StructTagName("tag").UseStructTags(false),
    )
    test.Nil(err,t)
    test.Eq(3,len(m),t)
    test.Eq(structIndex(0),m["One"],t)
    test.Eq(structIndex(1),m["Two"],t)
    test.Eq(structIndex(2),m["Three"],t)
}

