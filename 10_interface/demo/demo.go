package demo

import (
	"fmt"
)

type Langauge struct {
	ID   uint
	Name string
}

type Repository interface {
	QueryLang(uint) Langauge
}

type Handler struct {
	repo Repository
}

func NewHandler(repo Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) do(id uint) string {
	lang := h.repo.QueryLang(id)
	return fmt.Sprintf("%s language\n", lang.Name)
}

var data = map[uint]Langauge{
	1: {ID: 1, Name: "GO"},
	2: {ID: 2, Name: "JAVA"},
}

type StaticRepo struct {
	data map[uint]Langauge
}

func NewStaticRepository() StaticRepo {
	return StaticRepo{data: data}
}

func (s StaticRepo) QueryLang(index uint) Langauge {
	return s.data[index]
}

func Execute() {
	repo := NewStaticRepository()
	h := NewHandler(repo)
	s := h.do(2)
	fmt.Println(s)
}
