package main

import "fmt"

func main() {
	/*
		oprator aritmatika
			+ = tambah
			- = kurang
			* = kali
			/ = bagi
			% = sisa bagi
	*/

	//tambah
	a := 10        //pembuatan variabel a dengan isi 10
	b := 10        //pembuatan variabel b dengan isi 10
	c := a + b     //variable penampungan oprasi
	fmt.Println(c) //pemanggilan variable

	//kurang
	d := 10        //pembuatan variabel a dengan isi 10
	e := 10        //pembuatan variabel b dengan isi 10
	f := d - e     //variable penampungan oprasi
	fmt.Println(f) //pemanggilan variable

	//kali
	g := 10        //pembuatan variabel a dengan isi 10
	h := 10        //pembuatan variabel b dengan isi 10
	i := g * h     //variable penampungan oprasi
	fmt.Println(i) //pemanggilan variable

	//bagi
	j := 10        //pembuatan variabel a dengan isi 10
	k := 10        //pembuatan variabel b dengan isi 10
	l := j / k     //variable penampungan oprasi
	fmt.Println(l) //pemanggilan variable

	//sisa bagi
	m := 10        //pembuatan variabel a dengan isi 10
	n := 10        //pembuatan variabel b dengan isi 10
	o := m / n     //variable penampungan oprasi
	fmt.Println(o) //pemanggilan variable

	/*
		oprator augmented assigment
			a += 10
			a -= 10
			a *= 10
			a /= 10
			a %= 10
	*/

	var aa = 10
	aa += 10
	fmt.Println(aa)

	/*
		operatir unary
		++
		--
		-
		+
		!
	*/
	aa++
	fmt.Println(aa)
}
