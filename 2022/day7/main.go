package main

import (
	"fmt"
	"os"
	"sort"
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
			total_size += size
		}
	}

	return
}

func download_filespace(total_fs_size int, space_needed int, dir_sizes map[string]int) (size int) {

	total_fs_free := total_fs_size - dir_sizes["_/"]
	total_needed := space_needed - total_fs_free

	keys := make([]string, 0, len(dir_sizes))

	for key := range dir_sizes {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return dir_sizes[keys[i]] < dir_sizes[keys[j]]
	})

	for _, dir := range keys {
		if dir_sizes[dir] >= total_needed {
			size = dir_sizes[dir]
			break
		}
	}

	return
}

func dir_traversal(dir *Dir, pdirs []string, size int, dir_sizes map[string]int) map[string]int {

	index := dir.Parentdir + "_" + dir.Name
	if len(dir.Childfiles) > 0 {
		size = file_size(dir.Childfiles)
		dir_sizes[index] = size

		if len(pdirs) > 0 {
			for _, dir := range pdirs {
				dir_sizes[dir] = dir_sizes[dir] + size
			}
		}
	}

	if len(dir.Childdirs) > 0 {
		pdirs = append(pdirs, index)
		for _, child_dirs := range dir.Childdirs {
			dir_traversal(child_dirs, pdirs, size, dir_sizes)
		}
	} else {
		pdirs = []string{}
	}

	return dir_sizes
}

func dir_down(current_dir *Dir, dest string) *Dir {
	queue := make([]*Dir, 0)
	queue = append(queue, current_dir)
	for len(queue) > 0 {
		next_dir := queue[0]
		queue = queue[1:]

		if (next_dir.Name == dest) && (next_dir.Parentdir == current_dir.Name) {
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
	past_dirs := make([]*Dir, 0)
	past_dirs = append(past_dirs, current_dir)

	for i := 0; i < len(commands); i++ {
		line_parse := strings.Split(commands[i], " ")
		var command string
		if line_parse[0] == "$" {
			command = line_parse[1]

			switch command {
			case "cd":
				dest := line_parse[2]
				switch dest {
				case "/":
					break
				case "..":
					if current_dir.Name != "/" {
						n := len(past_dirs) - 1
						past_dirs = past_dirs[:n]
						current_dir = past_dirs[len(past_dirs)-1]
					}
				default:
					current_dir = dir_down(current_dir, dest)
					past_dirs = append(past_dirs, current_dir)
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

	var pdirs = []string{}
	dir_sizes = dir_traversal(&root, pdirs, 0, dir_sizes)

	total_file_size := calculate_total_size(dir_sizes)
	fmt.Println(total_file_size)

	fmt.Println(download_filespace(70000000, 30000000, dir_sizes))
}
