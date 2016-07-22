package trie

import "errors"

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	e        interface{}
	children [26]*TrieNode
}

func New() *Trie {
	return &Trie{&TrieNode{}}
}

func (t *Trie) Add(e string, data interface{}) {
	tn := t.root

	for _, i := range e {
		if tn.children[i-'a'] == nil {
			tn.children[i-'a'] = &TrieNode{e: nil}
		}
		tn = tn.children[i-'a']
	}

	tn.e = data
}

func (t *Trie) Get(e string) (interface{}, error) {
	tn := t.root
	for _, i := range e {
		if tn == nil {
			return nil, errors.New("No such element node")
		}
		tn = tn.children[i-'a']
	}

	if tn == nil || tn.e == nil {
		return nil, errors.New("No such element data")
	}

	return tn.e, nil
}

func (t *Trie) Traverse(handle func(e interface{})) {
	t.root.traverse(handle)
}

func (tn *TrieNode) traverse(handle func(e interface{})) {
	if tn == nil {
		return
	} else if tn.e != nil {
		handle(tn.e)
		//fmt.Println(tmp)
	}

	for _, c := range tn.children {
		//tmp += string(97 + i)
		c.traverse(handle)
		//tmp = tmp[:len(tmp)-1]
	}
}
