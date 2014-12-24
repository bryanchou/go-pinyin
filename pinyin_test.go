package pinyin

import (
	"reflect"
	"testing"
)

type pinyinFunc func(string, *Args) [][]string

func testPinyin(s string, d map[*Args][][]string, f pinyinFunc) (t *testing.T) {
	for a := range d {
		value, _ := d[a]
		v := f(s, a)
		if !reflect.DeepEqual(v, value) {
			t.Errorf("Expected %s, got %s", value, v)
		}
	}
	return t
}

func TestPinyin(t *testing.T) {
	hans := "中国人"
	test_data := map[*Args][][]string{
		&Args{}: [][]string{
			[]string{"zhong"},
			[]string{"guo"},
			[]string{"ren"},
		},
		&Args{Style: NORMAL}: [][]string{
			[]string{"zhong"},
			[]string{"guo"},
			[]string{"ren"},
		},
		&Args{Style: TONE}: [][]string{
			[]string{"zhōng"},
			[]string{"guó"},
			[]string{"rén"},
		},
		&Args{Style: TONE2}: [][]string{
			[]string{"zho1ng"},
			[]string{"guo2"},
			[]string{"re2n"},
		},
		&Args{Style: INITIALS}: [][]string{
			[]string{"zh"},
			[]string{"g"},
			[]string{"r"},
		},
		&Args{Style: FIRST_LETTER}: [][]string{
			[]string{"z"},
			[]string{"g"},
			[]string{"r"},
		},
		&Args{Style: FINALS}: [][]string{
			[]string{"ong"},
			[]string{"uo"},
			[]string{"en"},
		},
		&Args{Style: FINALS_TONE}: [][]string{
			[]string{"ōng"},
			[]string{"uó"},
			[]string{"én"},
		},
		&Args{Style: FINALS_TONE2}: [][]string{
			[]string{"o1ng"},
			[]string{"uo2"},
			[]string{"e2n"},
		},
		&Args{Heteronym: true}: [][]string{
			[]string{"zhong", "zhong"},
			[]string{"guo"},
			[]string{"ren"},
		},
	}

	testPinyin(hans, test_data, Pinyin)
}

func TestNoneHans(t *testing.T) {
	s := "abc"
	v := Pinyin(s, NewArgs())
	value := [][]string{[]string{}, []string{}, []string{}}
	if !reflect.DeepEqual(v, value) {
		t.Errorf("Expected %s, got %s", value, v)
	}
}

func TestLazyPinyin(t *testing.T) {
	s := "中国人"
	v := LazyPinyin(s, &Args{})
	value := []string{"zhong", "guo", "ren"}
	if !reflect.DeepEqual(v, value) {
		t.Errorf("Expected %s, got %s", value, v)
	}
}

func TestSlug(t *testing.T) {
	s := "中国人"
	v := Slug(s, &Args{})
	value := "zhongguoren"
	if v != value {
		t.Errorf("Expected %s, got %s", value, v)
	}

	v = Slug(s, &Args{Separator: ","})
	value = "zhong,guo,ren"
	if v != value {
		t.Errorf("Expected %s, got %s", value, v)
	}

	a := NewArgs()
	a.Separator = "-"
	v = Slug(s, a)
	value = "zhong-guo-ren"
	if v != value {
		t.Errorf("Expected %s, got %s", value, v)
	}
}

func TestFinal(t *testing.T) {
	value := "an"
	v := final("an")
	if v != value {
		t.Errorf("Expected %s, got %s", value, v)
	}
}
