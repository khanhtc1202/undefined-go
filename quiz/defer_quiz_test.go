package quiz_test

import "testing"

// 問題1~8の関数の戻り値を答えてください。

/* * * * * * * * * * * *
  まずは例題
* * * * * * * * * * * */

func 例題() int {
	val := 1
	val = val + 4
	return val
}

var 例題の戻り値 = 5 // ここの値を変更。問題はデフォルト0。例題は5が正解。
// ※例題なので変更しないでください。

/* * * * * * * * * * * *
  ここから問題スタートです。
* * * * * * * * * * * */

func 問題1() int {
	val := 1
	func() { val++ }()
	return val
}

var 問題1の戻り値 = 2 // 右辺を変更してください

func 問題2() int {
	val := 1
	{
		val := 2
		val++
	}
	return val
}

var 問題2の戻り値 = 1 // 右辺を変更してください

func 問題3() int {
	val := 1
	defer func() { val++ }()
	return val
}

var 問題3の戻り値 = 1 // 右辺を変更してください

func 問題4() (val int) {
	defer func() { val++ }()
	return 1
}

var 問題4の戻り値 = 2 // 右辺を変更してください

func 問題5() int {
	val := 1
	func() {
		defer func() { val++ }()
	}()
	return val
}

var 問題5の戻り値 = 2 // 右辺を変更してください

func 問題6() (val int) {
	val = 1
	defer func() {
		val++
		defer func() { val++ }()
	}()
	return 3
}

var 問題6の戻り値 = 5 // 右辺を変更してください

func 問題7() (val int) {
	defer func() {
		val++
		val := 5
		val++
	}()
	return val
}

var 問題7の戻り値 = 1 // 右辺を変更してください

func 問題8() (val int) {
	defer func() {
		val = 5
		if err := recover(); err == nil {
			val++
		}
		val++
	}()
	panic("panic")
}

var 問題8の戻り値 = 6 // 右辺を変更してください

/* * * * * * * * * * * * * * * * * *
  お疲れ様です。
  ページトップの「Run」で結果をみてみましょう。
* * * * * * * * * * * * * * * * * */

func TestResult(t *testing.T) {
	tests := map[string]struct {
		fn     func() int
		expect int
	}{
		// "例題": {例題, 例題の戻り値},
		"問題1": {問題1, 問題1の戻り値},
		"問題2": {問題2, 問題2の戻り値},
		"問題3": {問題3, 問題3の戻り値},
		"問題4": {問題4, 問題4の戻り値},
		"問題5": {問題5, 問題5の戻り値},
		"問題6": {問題6, 問題6の戻り値},
		"問題7": {問題7, 問題7の戻り値},
		"問題8": {問題8, 問題8の戻り値},
	}
	const max = 8
	result := max
	for k, tc := range tests {
		if tc.fn() != tc.expect {
			t.Errorf("不正解!! %s: 間違った答え %d", k, tc.expect)
			result -= 1
		}
	}
	t.Logf("結果は.. %d問正解です!", result)
	if result == max {
		t.Log("全問正解おめでとうございます！")
	}
}
