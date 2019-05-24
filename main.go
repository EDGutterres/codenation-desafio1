package main

import (
	"sort"
)

func main() {
	type State struct {
		name string
		size int
	}

	ac := State{
		name: "Acre",
		size: 152581,
	}

	al := State{
		name: "Alagoas",
		size: 27768,
	}

	ap := State{
		name: "Amapa",
		size: 142815,
	}

	am := State{
		name: "Amazonas",
		size: 1571000,
	}

	ba := State{
		name: "Bahia",
		size: 567295,
	}

	ce := State{
		name: "Ceara",
		size: 148826,
	}

	df := State{
		name: "Distrito Federal",
		size: 5802,
	}

	es := State{
		name: "Espirito Santo",
		size: 46095,
	}

	goi := State{
		name: "Goias",
		size: 340086,
	}

	ma := State{
		name: "Maranhao",
		size: 331983,
	}

	mt := State{
		name: "Mato Grosso",
		size: 903357,
	}

	ms := State{
		name: "Mato Grosso do Sul",
		size: 357125,
	}

	mg := State{
		name: "Minas Gerais",
		size: 586528,
	}

	pa := State{
		name: "Para",
		size: 1248000,
	}

	pb := State{
		name: "Paraiba",
		size: 56585,
	}

	pr := State{
		name: "Parana",
		size: 199315,
	}

	pe := State{
		name: "Pernambuco",
		size: 98312,
	}

	pi := State{
		name: "Piaui",
		size: 251529,
	}

	rj := State{
		name: "Rio de Janeiro",
		size: 43696,
	}

	rn := State{
		name: "Rio Grande do Norte",
		size: 52797,
	}

	rs := State{
		name: "Rio Grande do Sul",
		size: 281748,
	}

	ro := State{
		name: "Rondonia",
		size: 237576,
	}

	rr := State{
		name: "Roraima",
		size: 224299,
	}

	sc := State{
		name: "Santa Catarina",
		size: 95346,
	}

	sp := State{
		name: "Sao Paulo",
		size: 248209,
	}

	se := State{
		name: "Sergipe",
		size: 21910,
	}

	to := State{
		name: "Tocantins",
		size: 277621,
	}

	states := [27]State{
		ac,
		al,
		ap,
		am,
		ba,
		ce,
		df,
		es,
		goi,
		ma,
		mt,
		ms,
		mg,
		pa,
		pb,
		pr,
		pe,
		pi,
		rj,
		rn,
		rs,
		ro,
		rr,
		sc,
		sp,
		se,
		to,
	}

	bigger := make([]string, 10)
	sizes := make([]int, 27)
	biggerIndex := 0

	for i := 0; i < 27; i++ {
		sizes[i] = states[i].size
	}

	sort.Ints(sizes)

	for i := 26; i > 16; i-- {
		for j := 0; j < 27; j++ {
			if sizes[i] == states[j].size {
				bigger[biggerIndex] = states[j].name
				biggerIndex++
			}
		}

	}

	for key := range bigger {
		println(bigger[key])
	}
}
