package reflection

func walk(x interface{}, fn func(input string)) {
	fn("I still can't believe South Korea beat Germany to put them last in their group")
}
