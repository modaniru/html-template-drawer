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
