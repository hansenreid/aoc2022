package main

import (
	"math"
	"reflect"
	"strings"
	"testing"
)

const input = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

func TestParseLine(t *testing.T) {
	cases := map[string]Line{
		"$ ls":          &List{},
		"$ cd ..":       &ChangeDir{path: ".."},
		"$ cd /":        &ChangeDir{path: "/"},
		"dir a":         &Dir{name: "a", dirs: []*Dir{}, files: []*File{}, size: 0},
		"1234 file.txt": &File{name: "file.txt", size: 1234},
	}

	for line, expected := range cases {
		result := ParseLine(line)

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected: %v, Got: %v", expected, result)
		}

	}
}

func TestConsumeChangeDir(t *testing.T) {
	fs := NewFileSystem()
	d := NewDir("test")
	d.consume(&fs)

	cd := NewChangeDir("test")
	cd.consume(&fs)

	d2 := NewDir("test2")
	d2.consume(&fs)

	cd = NewChangeDir("test2")
	cd.consume(&fs)

	if fs.pwd != &d2 {
		t.Errorf("ChangeDir did not change into nested dir")
	}

	cd = NewChangeDir("..")
	cd.consume(&fs)

	if fs.pwd != &d {
		t.Errorf("ChangeDir did not change into parent dir")
	}

	cd = NewChangeDir("/")
	cd.consume(&fs)

	if fs.pwd != fs.root {
		t.Errorf("ChangeDir did not change into root dir")
	}
}

func TestPart1(t *testing.T) {
	r := strings.NewReader(input)
	fs := WalkInput(r)
	fs.calculateSize()

	sum := Part1(fs)

	if sum != 95437 {
		t.Errorf("Expected 95437, Got: %d", sum)
	}
}

func TestPart2(t *testing.T) {
	r := strings.NewReader(input)
	fs := WalkInput(r)
	fs.calculateSize()

	sum := Part2(fs)

	if sum != 24933642 {
		t.Errorf("Expected 24933642, Got: %d", sum)
	}
}

func TestSizeNeeded(t *testing.T) {
	r := strings.NewReader(input)
	fs := WalkInput(r)
	fs.calculateSize()

	result := fs.sizeNeeded()

	if result != 8381165 {
		t.Errorf("GOT: %d", result)
	}

}
