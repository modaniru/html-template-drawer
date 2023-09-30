package entity

type Article struct {
	Name string
}

type Course struct {
	Id            string
	Title         string
	Image         string
	ArticlesCount int
}
