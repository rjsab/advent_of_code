package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Dir struct {
	Name       string
	Parentdir  string
	Childdirs  []*Dir
	Childfiles []File
}

type File struct {
	Name string
	Size int
}

func file_size(files []File) (size int) {
	for _, files := range files {
		size += files.Size
	}
	return
}

func calculate_total_size(dir_sizes map[string]int) (total_size int) {

	total_size = 0
	for _, size := range dir_sizes {
		if size <= 100000 {
			//fmt.Printf("%s: %d\n", dir, size)
			total_size += size
		}
	}
	return
}

func dir_traversal(dir *Dir, pdirs []string, size int, dir_sizes map[string]int) map[string]int {

	fmt.Println(dir.Name, pdirs)
	if len(dir.Childfiles) > 0 {
		size = file_size(dir.Childfiles)
		dir_sizes[dir.Name] = size

		if len(pdirs) > 0 {
			for _, dir := range pdirs {
				dir_sizes[dir] = dir_sizes[dir] + size
			}
		}
	}
	if dir.Parentdir == "" {
		dir_sizes["/"] = dir_sizes[dir.Parentdir] + size
	} else {
		dir_sizes[dir.Parentdir] = dir_sizes[dir.Parentdir] + size
	}

	if len(dir.Childdirs) > 0 {
		pdirs = append(pdirs, dir.Name)
		for _, child_dirs := range dir.Childdirs {
			dir_traversal(child_dirs, pdirs, size, dir_sizes)
		}
	} else {
		pdirs = []string{}
	}

	return dir_sizes
}

func dir_traverse(root *Dir, size int, dir_sizes map[string]int) map[string]int {
	queue := make([]*Dir, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		next_dir := queue[0]
		queue = queue[1:]
		fmt.Printf("Dir name: %s\n", next_dir.Name)
		if len(next_dir.Childfiles) > 0 {
			size = file_size(next_dir.Childfiles)
			//current_dir_size
			dir_sizes[next_dir.Name] = size
		}
		if len(next_dir.Childdirs) > 0 {
			queue = append(queue, next_dir.Childdirs...)
		}
		fmt.Println(queue)
	}
	return dir_sizes
}

func change_dir(root *Dir, dest, parentdir string) *Dir {
	queue := make([]*Dir, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		next_dir := queue[0]
		queue = queue[1:]
		if next_dir.Name == dest && next_dir.Parentdir == parentdir {
			return next_dir
		}
		if len(next_dir.Childdirs) > 0 {
			queue = append(queue, next_dir.Childdirs...)
		}
	}
	return nil
}

func command_parse(commands []string) Dir {
	//Iniitalize root dir
	root := Dir{
		Name: "/",
	}

	current_dir := &root
	for i := 0; i < len(commands); i++ {
		// Part 1
		line_parse := strings.Split(commands[i], " ")
		var command string
		if line_parse[0] == "$" {
			command = line_parse[1]

			switch command {
			case "cd":
				//change dir
				dest := line_parse[2]
				switch dest {
				case "/":
					break
				case "..":
					if current_dir.Parentdir == "" {
						current_dir = change_dir(&root, "/", "")
					} else {
						current_dir = change_dir(&root, current_dir.Parentdir, current_dir.Parentdir)
					}
				default:
					current_dir = change_dir(&root, dest, current_dir.Parentdir)
				}
			case "ls":
				//read next lines until line starts with $ again
				for j := i + 1; j < len(commands); j++ {
					ls_split := strings.Split(commands[j], " ")
					if ls_split[0] == "$" {
						i = j - 1
						break
					} else if ls_split[0] == "dir" {
						current_dir.Childdirs = append(current_dir.Childdirs, &Dir{Name: ls_split[1], Parentdir: current_dir.Name})
					} else {
						file_size, _ := strconv.Atoi(ls_split[0])
						file_name := ls_split[1]
						current_dir.Childfiles = append(current_dir.Childfiles, File{Name: file_name, Size: file_size})
					}
				}
			}
		}
	}

	return root
}

func main() {
	file, err := os.ReadFile("../inputs/day7.txt")
	if err != nil {
		fmt.Println(err)
	}

	var dir_sizes = make(map[string]int)
	lines := strings.Split(string(file), "\n")
	root := command_parse(lines)

	jsonF, _ := json.MarshalIndent(&root, "", " ")
	fmt.Println(string(jsonF))
	var pdirs = []string{}
	dir_sizes1 := dir_traversal(&root, pdirs, 0, dir_sizes)
	//dir_sizes2 := dir_traverse(&root, 0, dir_sizes)

	fmt.Println(calculate_total_size(dir_sizes1))
	//fmt.Println(calculate_total_size(dir_sizes2))
}
