package main

type Prod struct {
	production string
	size       int
}

type Productions struct {
	itens map[int]Prod
}

func (p *Productions) Start() {

	p.itens = make(map[int]Prod)
	productions := []Prod{
		{"P' -> P", 1},
		{"P -> inicio V A", 3},
		{"V -> varinicio LV", 2},
		{"LV -> D LV", 2},
		{"LV -> varfim pt_v", 2},
		{"D -> TIPO L pt_v", 3},
		{"L -> id vir L", 3},
		{"L -> id", 1},
		{"TIPO -> inteiro", 1},
		{"TIPO -> real", 1},
		{"TIPO -> literal", 1},
		{"A -> ES A", 2},
		{"ES -> leia id pt_v", 3},
		{"ES -> escreva ARG pt_v", 3},
		{"ARG -> lit", 1},
		{"ARG -> num", 1},
		{"ARG -> id", 1},
		{"A -> CMD A", 2},
		{"CMD -> id rcb LD pt_v", 4},
		{"LD -> OPRD opm OPRD", 3},
		{"LD -> OPRD", 1},
		{"OPRD -> id", 1},
		{"OPRD -> num", 1},
		{"A -> COND A", 2},
		{"COND -> CAB CP", 2},
		{"CAB -> se ab_p EXP_R fc_p entao", 5},
		{"EXP_R -> OPRD opr OPRD", 3},
		{"CP -> ES CP", 2},
		{"CP -> CMD CP", 2},
		{"CP -> COND CP", 2},
		{"CP -> fimse", 1},
		{"A -> R A", 2},
		{"R -> CABR CPR", 2},
		{"CABR -> repita ab_p EXP_R fc_p", 4},
		{"CPR -> ES CPR", 2},
		{"CPR -> CMD CPR", 2},
		{"CPR -> COND CPR", 2},
		{"CPR -> fimrepita", 1},
		{"A -> fim", 1},
	}
	loop := 1
	for _, reserved := range productions {
		p.Put(reserved, loop)
		loop++
	}
}

func (p *Productions) Put(pr Prod, loop int) (key int) {

	p.itens[loop] = pr
	return
}

func (p *Productions) Get(value int) (pr Prod) {
	pr = p.itens[value]
	return
}
