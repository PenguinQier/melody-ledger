package validation

func ErrorMessage() map[string]string {
	return map[string]string{
		"required": "这个空格等着你填呢，别让它空着哦！",
		"email":    "这里需要一个能收到邮件的邮箱地址，别搞错了哦",
		"min":      "太短了，再长一点，安全系数更高哦！",
		"max":      "太长了，记不住怎么办？短一点，但也别太简单哦！",
	}
}
