package main

func Valid(name string) string {
	if name != "123456" {
		panic("不通过")
	}
	return "Hello world"
}
