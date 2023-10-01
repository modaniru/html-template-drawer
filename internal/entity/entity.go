package entity

type Article struct {
	TemplateName string
	Title        string
}

type Course struct {
	Id            string
	Title         string
	Image         string
	ArticlesCount int
}

type CourseForm struct {
	Name  string
	Image string
}

type ArticleForm struct {
	Title        string
	TemplateName string
	CourseId     string
}
