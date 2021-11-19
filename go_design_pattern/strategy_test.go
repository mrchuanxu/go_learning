package go_design_pattern_test

// 定义一族算法，将每个算法分别封装起来，让他们可以互相替换。
type FuncSort func(string)string

func QuickSort(string)string{
	return "quickSort"
}

func ExternalSort(string)string{
	return "ExternalSort"
}

func ConcurrentSort(string)string{
	return "ConcurrentSort"
}

var funcSortsM = map[string]FuncSort{
	"quickSort":QuickSort,"ExternalSort":ExternalSort,"ConcurrentSort":ConcurrentSort,
}

func sortFile(f string)string{
	sortFunc := funcSortsM["quickSort"](f)
	return sortFunc
}