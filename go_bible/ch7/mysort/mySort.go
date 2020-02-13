package mysort

// Interface 排序接口
type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}
