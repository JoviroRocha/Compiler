package main

type HashTable struct {
	itens map[int]Token
}

func (ht *HashTable) Start() {
	ht.itens = make(map[int]Token)

	reservedWords := []Token{
		{"inicio", "inicio", "inicio"},
		{"varinicio", "varinicio", "varinicio"},
		{"varfim", "varfim", "varfim"},
		{"escreva", "escreva", "escreva"},
		{"leia", "leia", "leia"},
		{"se", "se", "se"},
		{"entao", "entao", "entao"},
		{"fimse", "fimse", "fimse"},
		{"repita", "repita", "repita"},
		{"fimrepita", "fimrepita", "fimrepita"},
		{"fim", "fim", "fim"},
		{"inteiro", "inteiro", "inteiro"},
		{"literal", "literal", "literal"},
		{"real", "real", "real"},
	}
	for _, reserved := range reservedWords {
		ht.Put(reserved)
	}
}

func hash(lexema string) (key int) {
	key = 0
	for _, letra := range lexema {
		key = int(letra) + key*9
	}
	return
}

func (ht *HashTable) Put(token Token) (key int) {

	key = hash(token.Lexema)
	ht.itens[key] = token
	return
}

func (ht *HashTable) Att(token Token) int {
	return ht.Put(token)
}

func (ht *HashTable) Search(lexema string) (ok bool) {
	key := hash(lexema)
	_, ok = ht.itens[key]
	return
}

func (ht *HashTable) Get(lexema string) (token Token) {
	key := hash(lexema)
	token = ht.itens[key]
	return
}
