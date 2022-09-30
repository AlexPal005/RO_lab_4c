package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	readWriteLock sync.RWMutex
)

func main() {
	matrix := [][]int{}
	matrix = create_graph_matrix()
	changed_prices := [2]int{60, 70}
	numbers_of_vertex_from := [2]int{2, 1}
	numbers_of_vertex_to := [2]int{3, 0}
	go change_price(&matrix, numbers_of_vertex_from, numbers_of_vertex_to, changed_prices)

	number_for_delete_from := 2
	number_for_delete_to := 3
	number_for_add_from := 0
	number_for_add_to := 2
	cost := 60
	go delete_trip(&matrix, number_for_delete_from, number_for_delete_to, number_for_add_from, number_for_add_to, cost)

	go add_city(&matrix, 0)

	go delete_city(&matrix, 3)
	go search_trip(&matrix, 0, 3)
	time.Sleep(time.Second * 25)

	fmt.Println("Result:")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Print("\n")
	}

}

func create_graph_matrix() [][]int {
	matrix := [][]int{
		{0, 10, 0, 0, 0},  //0
		{10, 0, 20, 0, 0}, //1
		{0, 20, 0, 40, 0}, //2
		{0, 0, 40, 0, 30}, //3
		{0, 0, 0, 30, 0},
	}
	return matrix
}
func change_price(matrix *[][]int, numbers_vertex_from [2]int, numbers_vertex_to [2]int, changed_prices [2]int) {
	readWriteLock.Lock()
	for i := 0; i < 2; i++ {
		time.Sleep(time.Second * 2)
		(*matrix)[numbers_vertex_from[i]][numbers_vertex_to[i]] = changed_prices[i]
		(*matrix)[numbers_vertex_to[i]][numbers_vertex_from[i]] = changed_prices[i]
		fmt.Println("Change price from ", numbers_vertex_from[i], " to", numbers_vertex_to[i], " on", changed_prices[i], " !")
	}
	for i := 0; i < len((*matrix)); i++ {
		for j := 0; j < len((*matrix)[0]); j++ {
			fmt.Print((*matrix)[i][j], " ")
		}
		fmt.Print("\n")
	}
	readWriteLock.Unlock()
}
func delete_trip(matrix *[][]int, number_for_delete_from int, number_for_delete_to int,
	number_for_add_from int, number_for_add_to int, cost int) {

	readWriteLock.Lock()

	time.Sleep(time.Second * 2)
	(*matrix)[number_for_delete_from][number_for_delete_to] = 1
	(*matrix)[number_for_delete_to][number_for_delete_from] = 1
	fmt.Println("Trip from", number_for_delete_from, " to ", number_for_delete_to, " deleted!")

	time.Sleep(time.Second * 2)
	(*matrix)[number_for_add_from][number_for_add_to] = cost
	(*matrix)[number_for_add_to][number_for_add_from] = cost
	fmt.Println("Trip from", number_for_add_from, " to ", number_for_add_to, " Added!", " Cost = ", cost)
	for i := 0; i < len((*matrix)); i++ {
		for j := 0; j < len((*matrix)[0]); j++ {
			fmt.Print((*matrix)[i][j], " ")
		}
		fmt.Print("\n")
	}
	readWriteLock.Unlock()
}
func add_city(matrix *[][]int, add_near_city int) {
	readWriteLock.Lock()
	time.Sleep(time.Second * 2)
	arr := make([]int, len((*matrix)))
	(*matrix) = append((*matrix), arr)
	for i := 0; i < len((*matrix)); i++ {
		(*matrix)[i] = append((*matrix)[i], 0)
	}
	(*matrix)[len((*matrix))-1][add_near_city] = 1
	(*matrix)[add_near_city][len((*matrix))-1] = 1
	fmt.Println("City number", len((*matrix))-1, "Added near sity number ", add_near_city)
	for i := 0; i < len((*matrix)); i++ {
		for j := 0; j < len((*matrix)[0]); j++ {
			fmt.Print((*matrix)[i][j], " ")
		}
		fmt.Print("\n")
	}
	readWriteLock.Unlock()
}
func delete_city(matrix *[][]int, number_of_city int) {
	readWriteLock.Lock()
	time.Sleep(time.Second * 2)
	for i := 0; i < len(*matrix); i++ {
		(*matrix)[i][number_of_city] = 0
	}
	for i := 0; i < len((*matrix)[0]); i++ {
		(*matrix)[number_of_city][i] = (*matrix)[len(*matrix)-1][i]
	}
	(*matrix) = (*matrix)[:len((*matrix))-1]

	for i := 0; i < len((*matrix)); i++ {
		(*matrix)[i] = (*matrix)[i][:len((*matrix)[i])-1]
	}

	fmt.Println("Deleted number ", number_of_city)
	for i := 0; i < len((*matrix)); i++ {
		for j := 0; j < len((*matrix)[0]); j++ {
			fmt.Print((*matrix)[i][j], " ")
		}
		fmt.Print("\n")
	}
	readWriteLock.Unlock()
}
func search_trip(matrix *[][]int, city_one int, city_two int) {
	readWriteLock.RLock()
	time.Sleep(time.Second * 3)
	if (*matrix)[city_one][city_two] != 0 && (*matrix)[city_one][city_two] != 1 {
		fmt.Println("Price: ", (*matrix)[city_one][city_two], " grn")
	} else {
		cost := 0
		index := 0
		count := 0
		for i := city_one; i < len(*matrix); i++ {
			for j := 0; j < len(*matrix); j++ {
				if (*matrix)[i][j] != 0 && (*matrix)[i][j] != 1 {
					if count == 0 {
						cost += (*matrix)[i][j]
						count = 1
					} else {
						count = 0
					}
					index = j
					if index == city_two {
						break
					}
				}
			}
			if index == city_two {
				break
			}
		}
		fmt.Println("Price: ", cost, " grn")
	}
	readWriteLock.RUnlock()
}
