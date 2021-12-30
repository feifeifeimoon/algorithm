package main

import (
	"container/list"
	"fmt"
	"reflect"
)

var letters = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	var ret [][]string
	wordMap := make(map[string][]string, len(wordList)+1)
	for i := range wordList {
		wordMap[wordList[i]] = []string{}
	}
	wordMap[beginWord] = []string{}

	if _, ok := wordMap[endWord]; !ok {
		return ret
	}

	queue := list.New()
	queue.PushBack(beginWord)

	visited := make(map[string]struct{})
	visited[beginWord] = struct{}{}

	wordLevel := make(map[string]int)

	level := 0
	found := false
	for queue.Len() > 0 {
		level++
		for size := queue.Len(); size > 0; size-- {
			elem := queue.Front()
			word := elem.Value.(string)

			newWord := []byte(word)
			for i := range word {
				tmp := word[i]
				// 从 a-z 一个一个尝试 看有没有在wordList里的
				for j := range letters {
					if word[i] == letters[j] {
						continue
					}

					newWord[i] = letters[j]
					nw := string(newWord)
					if _, ok := wordMap[nw]; !ok {
						continue
					}

					wordMap[word] = append(wordMap[word], nw)
					if nw == endWord {
						found = true
					}
					// visited
					if _, ok := visited[nw]; !ok {
						visited[nw] = struct{}{}
						wordLevel[nw] = level
						queue.PushBack(nw)
					}
				}
				newWord[i] = tmp
			}
			queue.Remove(elem)
		}
	}

	if !found {
		return ret
	}

	for k, v := range wordMap {
		fmt.Println(k, v)
	}
	for k, v := range wordLevel {
		fmt.Println(k, v)
	}

	var dfs func(word string, level int, path []string)
	dfs = func(word string, level int, path []string) {
		path = append(path, word)
		fmt.Printf("dfs %v %v %v %p\n", word, level, path, path)
		if word == endWord {
			deepCopy := make([]string, len(path))
			copy(deepCopy, path)
			ret = append(ret, deepCopy)
			return
		}
		for _, v := range wordMap[word] {
			if wordLevel[v] == level+1 {
				dfs(v, level+1, path)
			}

		}
	}
	dfs(beginWord, 0, []string{})
	return ret
}

func main() {
	type args struct {
		beginWord, endWord string
		wordList           []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			"test1",
			args{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}},
			[][]string{{"hit", "hot", "dot", "dog", "cog"}, {"hit", "hot", "lot", "log", "cog"}},
		},
		{
			"test2",
			args{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}},
			[][]string{},
		},
		{
			"test3",
			args{"teach", "place", []string{"peale", "wilts", "place", "fetch", "purer", "pooch", "peace", "poach", "berra", "teach", "rheum", "peach"}},
			[][]string{{"teach", "peach", "peace", "place"}},
		},
	}
	for _, tt := range tests {
		if got := findLadders(tt.args.beginWord, tt.args.endWord, tt.args.wordList); !reflect.DeepEqual(got, tt.want) {
			fmt.Printf("%s findLadders() = %v, want %v\n", tt.name, got, tt.want)
		}
	}
}
