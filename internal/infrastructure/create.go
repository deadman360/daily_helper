package infrastructure

import "fmt"

func (r *repository) Create(name string) {
	fmt.Println("chegamos no repository", name)
}
