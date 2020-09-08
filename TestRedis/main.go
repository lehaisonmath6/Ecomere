package main

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type Category struct {
	CategoryID       int64  `json:"category_id,omitempty"`
	ParentCategoryID int64  `json:"parent_category_id,omitempty"`
	CategoryName     string `json:"category_name,omitempty"`
}

type Image struct {
	URL    string `json:"url,omitempty"`
	Witdh  int64  `json:"witdh,omitempty"`
	Height int64  `json:"height,omitempty"`
}

type Location struct {
	Lat     float64 `json:"lat,omitempty"`
	Long    float64 `json:"long,omitempty"`
	Street  string  `json:"street,omitempty"`
	Number  int64   `json:"number,omitempty"`
	City    string  `json:"city,omitempty"`
	State   string  `json:"state,omitempty"`
	Country string  `json:"country,omitempty"`
}

// => Mashal chinh chuyen tu struct => bytes, C# object => bytes

type Item struct {
	ListImages []*Image  `json:"list_images,omitempty"`
	Title      string    `json:"title,omitempty"`
	Price      float64   `json:"price,omitempty"`
	Cat        *Category `json:"cat,omitempty"`
	Desciption string    `json:"desciption,omitempty"`
	Loc        *Location `json:"loc,omitempty"`
	Sold       bool      `json:"sold,omitempty"`
	UID        int64     `json:"uid,omitempty"`
	Timestamps int64     `json:"timestamps,omitempty"`
	ID         int64     `json:"id"`
}

func (m *Item) MarshalBinary() (data []byte, err error) {
	return json.Marshal(m)
}

func (m *Item) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}

func PutItem(rdb *redis.Client) {
	item1 := &Item{
		ID: 1,
		ListImages: []*Image{&Image{
			URL:    "https://surfacecu.com.vn/wp-content/uploads/2020/05/surface-laptop-3-a1-1.jpg",
			Witdh:  500,
			Height: 500,
		}, &Image{
			URL:    "https://surfacecu.com.vn/wp-content/uploads/2020/05/surface-laptop-3-a1-1.jpg",
			Witdh:  500,
			Height: 500,
		}},

		Title:      "Surface Laptop 3",
		Desciption: "Surface laptop 3 gía rẻ chip i7 10th ram 16Gb",
		Price:      25000000,
		Loc: &Location{
			Lat:  10.0005,
			Long: 20.1101,
		},
		UID:        5,
		Timestamps: time.Now().Unix(),
		Sold:       false,
		Cat: &Category{
			CategoryID:       2,
			ParentCategoryID: 1,
			CategoryName:     "Laptop",
		},
	}

	err := rdb.Set(context.Background(), strconv.FormatInt(item1.ID, 10), item1, time.Duration(0)).Err()
	if err != nil {
		log.Println("PUT err", err)
	} else {
		log.Println("PUT OKE")
	}
}

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	data, err := rdb.Get(context.Background(), "1").Result()
	if err != nil {
		log.Println("Read data err", err)
		return
	}
	log.Println("Data", data)
}
