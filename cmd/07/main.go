package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const totalSize = 70_000_000
const updateSize = 30_000_000

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error reading file: ", err)
	}

	defer f.Close()

	fs := WalkInput(f)
	fs.calculateSize()

	fmt.Printf("Part 1 Solution: %d\n", Part1(fs))
	fmt.Printf("Part 2 Solution: %d\n", Part2(fs))
}

func Part1(fs *FileSystem) int {
	return fs.sumUnder100k()
}

func Part2(fs *FileSystem) int {
	needed := fs.sizeNeeded()
	return fs.minOverN(needed)
}

func WalkInput(r io.Reader) *FileSystem {
	s := bufio.NewScanner(r)
	fs := NewFileSystem()

	for s.Scan() {
		line := ParseLine(s.Text())
		line.consume(&fs)
	}

	return &fs
}

func ParseLine(s string) Line {
	inputs := strings.Split(s, " ")

	switch inputs[0] {
	case "$":
		return ParseCommand(inputs[1:])

	case "dir":
		dir := NewDir(inputs[1])
		return &dir

	default:
		f := NewFile(inputs[1], inputs[0])
		return &f
	}
}

func ParseCommand(inputs []string) Line {
	switch inputs[0] {
	case "ls":
		return &List{}

	case "cd":
		cd := NewChangeDir(inputs[1])
		return &cd

	default:
		log.Fatal("Error Parsing Command: ", inputs)
		return nil
	}
}

type FileSystem struct {
	root *Dir
	pwd  *Dir
}

func NewFileSystem() FileSystem {
	root := NewDir("/")

	return FileSystem{
		root: &root,
		pwd:  &root,
	}
}

func (fs *FileSystem) calculateSize() int {
	return fs.root.calculateSize()
}

func (fs *FileSystem) sumUnder100k() int {
	return fs.root.sumUnder100k()
}

func (fs *FileSystem) minOverN(n int) int {
	min := math.MaxInt
	dirs := fs.root.allDirs()

	for i := range dirs {
		s := dirs[i].size
		if s < min && s >= n {
			min = s
		}
	}

	return min
}

func (fs *FileSystem) sizeNeeded() int {
	size := fs.root.size
	unused := totalSize - size
	return updateSize - unused
}

type Line interface {
	consume(fs *FileSystem)
}

type ChangeDir struct {
	path string
}

func NewChangeDir(path string) ChangeDir {
	return ChangeDir{
		path: path,
	}
}

func (cd *ChangeDir) consume(fs *FileSystem) {
	switch cd.path {
	case "/":
		fs.pwd = fs.root

	case "..":
		fs.pwd = fs.pwd.parent

	default:
		for i := range fs.pwd.dirs {
			if fs.pwd.dirs[i].name == cd.path {
				fs.pwd = fs.pwd.dirs[i]
				break
			}
		}
	}
}

type List struct{}

func (ls *List) consume(fs *FileSystem) {
}

type File struct {
	name string
	size int
}

func NewFile(name string, size string) File {
	n, err := strconv.Atoi(size)
	if err != nil {
		log.Fatal("Error Parsing File size: ", size)
	}

	return File{name: name, size: n}
}

func (f *File) consume(fs *FileSystem) {
	fs.pwd.files = append(fs.pwd.files, f)
}

type Dir struct {
	name   string
	dirs   []*Dir
	files  []*File
	size   int
	parent *Dir
}

func NewDir(name string) Dir {
	return Dir{
		name:  name,
		dirs:  []*Dir{},
		files: []*File{},
	}
}

func (d *Dir) consume(fs *FileSystem) {
	d.parent = fs.pwd
	fs.pwd.dirs = append(fs.pwd.dirs, d)
}

func (d *Dir) calculateSize() int {
	sum := 0

	for i := range d.files {
		sum += d.files[i].size
	}

	for i := range d.dirs {
		sum += d.dirs[i].calculateSize()
	}

	d.size = sum
	return sum
}

func (d *Dir) sumUnder100k() int {
	sum := 0

	if d.size <= 100_000 {
		sum += d.size
	}

	for i := range d.dirs {
		sum += d.dirs[i].sumUnder100k()
	}

	return sum
}

func (d *Dir) allDirs() []*Dir {
	dirs := []*Dir{}

	dirs = append(dirs, d)
	for i := range d.dirs {
		dirs = append(dirs, d.dirs[i].allDirs()...)
	}

	return dirs
}
