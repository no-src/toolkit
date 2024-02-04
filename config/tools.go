package config

var Tools = []Config{
	{
		[]string{"timestamp", "time", "ts"},
		[]string{"https://tool.lu/timestamp"},
	},
	{
		[]string{"dict", "d"},
		[]string{"https://www.youdao.com", "https://www.ldoceonline.com"},
	},
	{
		[]string{"translate", "tran", "t"},
		[]string{"https://fanyi.baidu.com", "https://translate.google.com", "https://fanyi.sogou.com"},
	},
	{
		[]string{"write", "w"},
		[]string{"https://write.youdao.com/#/index?from=select_free"},
	},
	{
		[]string{"json", "j"},
		[]string{"https://www.sojson.com"},
	},
	{
		[]string{"calendar", "cale"},
		[]string{"https://www.baidu.com/s?wd=万年历"},
	},
	{
		[]string{"github"},
		[]string{"https://github.com"},
	},
	{
		[]string{"trending"},
		[]string{"https://github.com/trending"},
	},
	{
		[]string{"chatgpt", "gpt", "chat", "ai"},
		[]string{"https://chat.openai.com"},
	},
	{
		[]string{"algorithms", "algorithm", "algo", "a"},
		[]string{"https://www.cs.usfca.edu/~galles/visualization/Algorithms.html", "https://algorithm-visualizer.org"},
	},
	{
		[]string{"search", "se"},
		[]string{"https://www.google.com", "https://www.baidu.com"},
	},
	{
		[]string{"go"},
		[]string{"https://go.dev", "https://github.com/golang/go"},
	},
	{
		[]string{"gofs"},
		[]string{"https://github.com/no-src/gofs"},
	},
	{
		[]string{"jwt"},
		[]string{"https://jwt.io"},
	},
}
