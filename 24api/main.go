package main

// Model for course file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// nakli db using slice

var courses []Course

// middleware, helper-file
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName ==""
}

func main() {
}
