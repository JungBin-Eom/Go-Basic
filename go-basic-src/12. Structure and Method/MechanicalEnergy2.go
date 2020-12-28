package main

import "fmt"

const g = 9.8

type object struct {
	m, v, h, ke, pe, me float32
}

func (o object) calKE() float32 {
	KE := 0.5 * o.m * o.v * o.v
	return KE
}

func (o object) calPE() float32 {
	PE := o.m * g * o.h
	return PE
}

func main() {
	var objNum int
	var m, v, h float32

	fmt.Scanln(&objNum)

	var objects []object

	for i := 0; i < objNum; i++ {
		fmt.Scanln(&m, &v, &h)
		newObj := object{}
		newObj.m = m
		newObj.v = v
		newObj.h = h
		newObj.ke = newObj.calKE()
		newObj.pe = newObj.calPE()
		newObj.me = newObj.ke + newObj.pe
		objects = append(objects, newObj)
	}

	for _, obj := range objects {
		fmt.Println(obj.ke, obj.pe, obj.me)
	}

}
