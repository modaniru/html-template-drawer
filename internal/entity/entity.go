package entity

type Article struct {
	Id           string
	TemplateName string
	Title        string
	Course       string
}

type Course struct {
	Id            string
	Title         string
	TitleId       string
	Image         string
	ArticlesCount int
}

type CourseForm struct {
	Name  string
	Image string
}

type SaveCourse struct {
	Title   string
	Image   string
	TitleId string
}

type ArticleForm struct {
	Title        string
	TemplateName string
	CourseId     string
}
