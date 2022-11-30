package base_demo

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(want, got){
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

func TestSplitWithComplexSep(t *testing.T){
	got := Split("abcd", "bc")
	want := []string{"a", "d"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected: %v, got: %v", want, got)
	}
}
//表格驱动测试
func TestSplitAll(t *testing.T){
	//1. 定义测试表格
	//2. 使用匿名结构体定义若干个测试用例
	//3. 为每个测试用例设置了一个名称
	tests := []struct {
		name string
		input string
		sep string
		want []string
	}{
		{"Base Case", "a:b:c:d", ":", []string{"a", "b", "c", "d"}},
		{"wrong Sep", "a:b:c", ",", []string{"a", "b", "c"}},
		{"More Sep", "abcd", "bc", []string{"a", "d"}},
		{"Leading Sep", "沙河有沙又有河","沙", []string{"", "河有", "又有河"}},
	}
	//遍历测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){ //
			got := Split(tt.input, tt.sep)
			if !reflect.DeepEqual(got, tt.want){
				t.Errorf("expected: %v, got:%v", tt.want, got)
			}
		})
	}
}