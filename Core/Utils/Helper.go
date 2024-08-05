package Utils

import "Taxonim/Token"

type Pool[T any] struct {
	Items    chan T
	Factory  func() T
	Close    func(T)
	AfterPut func(T)
}

/*
|========================================================
|	Get() -> T
|========================================================
|
|
|
|
|
*/
func (this *Pool[T]) Get() T {
	var Item T
	select {
	case Item = <-this.Items:
	default:
		Item = this.Factory()
	}
	return Item
}

/*
|========================================================
|	Put() -> error
|========================================================
|
|
|
|
|
*/
func (this *Pool[T]) Put(item T) error {
	if this.Items == nil {
		// pool is closed, close passed client
		this.Close(item)
		return nil
	}
	select {
	case this.Items <- item:
		if this.AfterPut != nil {
			this.AfterPut(item)
		}
		return nil
	default:
		this.Close(item)
		return nil
	}
}

/*
|========================================================
|	Len() -> int
|========================================================
|
|
|
|
|
*/
func (this *Pool[T]) Len() int {
	return len(this.Items)
}

/*
|========================================================
|	Done()
|========================================================
|
|
|
|
|
*/
func (this *Pool[T]) Done() {
	close(this.Items)
	for i := range this.Items {
		this.Close(i)
	}
}

/*
|========================================================
|	LookupIdent() -> Token
|========================================================
|
|
|
|
|
*/
func LookupIdent(ident string) Token.TokenType {
	if tok, ok := Token.Keywords[ident]; ok {
		return tok
	}
	return Token.IDENT
}
