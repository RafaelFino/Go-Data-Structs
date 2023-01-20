package main

import (
	"fmt"
	"github.com/google/uuid"
	"google-interview/internal/Tree"
	"math/rand"
	"os"
	"time"
)

func main() {
	RunStringTree()
	//RunIntTree()
}

func RunIntTree() {
	tree := Tree.NewIntTree()

	rdm := rand.New(rand.NewSource(time.Now().UnixNano()))
	qty := 10000
	arr := make([]int, qty)

	//Random Insert
	for i := 0; i < qty; i++ {
		arr[i] = rdm.Int()
		tree.Insert(arr[i])
	}

	//Test Print
	tree.Print(os.Stdout)

	//Test Find
	for i := 0; i < qty/100; i++ {
		fmt.Printf("Finding Item [%d] Heigth -> %d\n", arr[i], tree.Find(arr[i]))
	}
}

func RunStringTree() {
	tree := Tree.NewStringTree()

	qty := 1000
	arr := make([]string, qty)

	//Random Insert
	for i := 0; i < qty; i++ {
		arr[i] = uuid.New().String()
		tree.Insert(fmt.Sprintf("%s", arr[i]))
	}

	//Test Print
	tree.Print(os.Stdout)

	//Test Find
	for i := 0; i < qty/100; i++ {
		fmt.Printf("Finding Item [%s] Heigth -> %d\n", arr[i], tree.Find(fmt.Sprintf("%s", arr[i])))
	}
}
