package sort

// TrieNode 定义了Trie树的节点结构
type TrieNode struct {
	Children [26]*TrieNode // 假设只存储小写字母，使用26个字母
	IsEnd    bool          // 标记该节点是否为单词的结尾
}

// Trie 定义了Trie树的结构
type Trie struct {
	Root *TrieNode
}

// NewTrie 创建一个新的Trie树实例
func NewTrie() *Trie {
	return &Trie{
		Root: &TrieNode{},
	}
}

// Insert 向Trie树中插入一个单词
func (t *Trie) Insert(word string) {
	node := t.Root
	for _, r := range word {
		index := int(r - 'a')
		if node.Children[index] == nil {
			node.Children[index] = &TrieNode{}
		}
		node = node.Children[index]
	}
	node.IsEnd = true
}

// Search 在Trie树中搜索一个单词
func (t *Trie) Search(word string) bool {
	node := t.Root
	for _, r := range word {
		index := int(r - 'a')
		if node.Children[index] == nil {
			return false
		}
		node = node.Children[index]
	}
	return node.IsEnd
}

// StartsWith 检查Trie树中是否有以给定前缀开头的单词
func (t *Trie) StartsWith(prefix string) bool {
	node := t.Root
	for _, r := range prefix {
		index := int(r - 'a')
		if node.Children[index] == nil {
			return false
		}
		node = node.Children[index]
	}
	return true
}
