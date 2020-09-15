package main

const ALPHABET_SIZE = 26

type Trie struct {
	children    []*Trie // Use hashMap if strings are long and some letters are rarely used or alphabet is much larger than 26 chars
	isEndOfWord bool
}

func Constructor() Trie {
	return Trie{
		children: make([]*Trie, ALPHABET_SIZE),
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'

		if this.children[index] == nil {
			t := Constructor()
			this.children[index] = &t
		}

		this = this.children[index]
	}

	this.isEndOfWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if this.children[index] == nil {
			return false
		}

		this = this.children[index]
	}

	if !this.isEndOfWord {
		return false
	}

	return true
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for i := 0; i < len(prefix); i++ {
		index := prefix[i] - 'a'
		if this.children[index] == nil {
			return false
		}

		this = this.children[index]
	}

	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
