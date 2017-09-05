package dgorm_test

import (
	"testing"

	"log"

	"github.com/akshaydeo/dgorm"
)

type Dog struct {
	Id    int    `dgraph:"uid"`
	Name  string `dgraph:"name"`
	Color string `json:"color" dgraph:"color"`
}

func TestDgraph_Add(t *testing.T) {
	dg, err := dgorm.Connect([]string{"127.0.0.1:9080"})
	if err != nil {
		t.Fail()
	}
	d := new(Dog)
	d.Id = 1
	d.Name = "jarvis"
	d.Color = "white"
	err = dg.Add(d)
	if err != nil {
		log.Println(err.Error())
		t.Fail()
	}
}

type Dog2 struct {
	Id         int    `dgraph:"uid"`
	Name       string `dgraph:"name"`
	Color      string `json:"color" dgraph:"color"`
	LikesPlace Place  `dgraph:"likes_place"`
}

type Place struct {
	Id   int    `dgraph:"uid"`
	Name string `json:"name" dgraph:"name"`
}

func TestDgraph_AddWithRelation(t *testing.T) {
	dg, err := dgorm.Connect([]string{"127.0.0.1:9080"})
	if err != nil {
		t.Fail()
	}
	d := new(Dog2)
	d.Id = 1
	d.Name = "jarvis"
	d.Color = "white"
	d.LikesPlace = Place{1, "Pune"}
	err = dg.Add(d)
	if err != nil {
		log.Println(err.Error())
		t.Fail()
	}
}

type Dog3 struct {
	Id         int    `dgraph:"uid"`
	Name       string `dgraph:"name"`
	Color      string `json:"color" dgraph:"color"`
	LikesPlace Place2 `dgraph:"likes_place"`
}

type Place2 struct {
	Name string `json:"name" dgraph:"name"`
}

func TestDgraph_AddWithRelationWithoutId(t *testing.T) {
	dg, err := dgorm.Connect([]string{"127.0.0.1:9080"})
	if err != nil {
		t.Fail()
	}
	d := new(Dog3)
	d.Id = 1
	d.Name = "jarvis"
	d.Color = "white"
	d.LikesPlace = Place2{"Pune"}
	err = dg.Add(d)
	if err != nil {
		log.Println(err.Error())
		t.Fail()
	}
}

type Dog4 struct {
	Id         int    `dgraph:"uid"`
	Name       string `dgraph:"name"`
	Color      string `json:"color" dgraph:"color"`
	LikesPlace *Place `dgraph:"likes_place"`
}

func TestDgraph_AddWithRelationPointer(t *testing.T) {
	dg, err := dgorm.Connect([]string{"127.0.0.1:9080"})
	if err != nil {
		t.Fail()
	}
	d := new(Dog4)
	d.Id = 1
	d.Name = "jarvis"
	d.Color = "white"
	d.LikesPlace = &Place{1, "Pune"}
	err = dg.Add(d)
	if err != nil {
		log.Println(err.Error())
		t.Fail()
	}
}

type Dog5 struct {
	Id         int     `dgraph:"uid"`
	Name       string  `dgraph:"name"`
	Color      string  `json:"color" dgraph:"color"`
	LikesPlace []Place `dgraph:"likes_place"`
}

func TestDgraph_AddWithRelationSliceWithStruct(t *testing.T) {
	dg, err := dgorm.Connect([]string{"127.0.0.1:9080"})
	if err != nil {
		t.Fail()
	}
	d := new(Dog5)
	d.Id = 1
	d.Name = "jarvis"
	d.Color = "white"
	d.LikesPlace = []Place{Place{1, "Pune"}, Place{2, "Mumbai"}}
	err = dg.Add(d)
	if err != nil {
		log.Println(err.Error())
		t.Fail()
	}
}

type Dog6 struct {
	Id         int      `dgraph:"uid"`
	Name       string   `dgraph:"name"`
	Color      string   `json:"color" dgraph:"color"`
	LikesPlace []*Place `dgraph:"likes_place"`
}

func TestDgraph_AddWithRelationSliceWithPointer(t *testing.T) {
	dg, err := dgorm.Connect([]string{"127.0.0.1:9080"})
	if err != nil {
		t.Fail()
	}
	d := new(Dog6)
	d.Id = 1
	d.Name = "jarvis"
	d.Color = "white"
	d.LikesPlace = []*Place{&Place{1, "Pune"}, &Place{2, "Mumbai"}}
	err = dg.Add(d)
	if err != nil {
		log.Println(err.Error())
		t.Fail()
	}
}

type Dog7 struct {
	Id        int      `dgraph:"uid"`
	Name      string   `dgraph:"name"`
	NickNames []string `dgraph:"nicknames"`
	Color     string   `json:"color" dgraph:"color"`
}

func TestDgraph_AddWithRelationSliceWithPrimitiveDt(t *testing.T) {
	dg, err := dgorm.Connect([]string{"127.0.0.1:9080"})
	if err != nil {
		t.Fail()
	}
	d := new(Dog7)
	d.Id = 1
	d.Name = "jarvis"
	d.Color = "white"
	d.NickNames = []string{"chotu", "motu"}
	err = dg.Add(d)
	if err != nil {
		log.Println(err.Error())
		t.Fail()
	}
}

type Dog8 struct {
	Id        int       `dgraph:"uid"`
	Name      string    `dgraph:"name"`
	NickNames []*string `dgraph:"nicknames"`
	Color     string    `json:"color" dgraph:"color"`
}

func TestDgraph_AddWithRelationSliceWithPrimitiveDtPointer(t *testing.T) {
	dg, err := dgorm.Connect([]string{"127.0.0.1:9080"})
	if err != nil {
		t.Fail()
	}
	d := new(Dog8)
	d.Id = 1
	d.Name = "jarvis"
	d.Color = "white"
	d.NickNames = []*string{dgorm.StrPtr("chotu"), dgorm.StrPtr("motu")}
	err = dg.Add(d)
	if err != nil {
		log.Println(err.Error())
		t.Fail()
	}
}
