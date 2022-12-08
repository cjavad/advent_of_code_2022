package main

import (
	"fmt"
	"strings"

	utils "github.com/cjavad/advent_of_code_2022"
)

type fs_entry struct {
	name    string
	size    int
	is_file bool
	entries []fs_entry
}

func main() {
	part1()
	part2()
}

func part1() {
	lines := utils.ReadInput("input.txt")
	terminal := parse_terminal(lines)
	filesystem := parse_filesystem(terminal)
	all_dirs_under_100k := get_all_directories(filesystem, 100_001)
	pretty_print_fs(all_dirs_under_100k, 0)
	fmt.Println("Part 1: ", sum_fs(all_dirs_under_100k))

	/*
		Not a part of part 1, but I wanted to see all combinations of directories that are less than 100k
	*/
	comb := find_combinations_lt([]fs_entry{}, all_dirs_under_100k, 100_001)
	sums_of_comb := make([]int, len(comb))
	for i, c := range comb {
		sums_of_comb[i] = sum_fs(c)
	}

	// fmt.Println(sums_of_comb)
}

func part2() {
	total_space := 70_000_000
	free_space_require := 30_000_000

	fs := parse_filesystem(parse_terminal(utils.ReadInput("input.txt")))
	all_dirs := get_all_directories(fs, 70_000_000)
	pretty_print_fs(all_dirs, 0)
	root_size := 0
	for _, entry := range fs {
		root_size += entry.size
	}
	free_space := total_space - root_size
	missing_space := free_space_require - free_space
	// Find directory that is closest to missing_space

	curr_index := 0
	space_diff := missing_space

	for i, entry := range all_dirs {
		if entry.size < missing_space {
			continue
		}

		if entry.size-missing_space < space_diff {
			space_diff = entry.size - missing_space
			curr_index = i
		}
	}

	fmt.Println("Part 2: ", all_dirs[curr_index])
}

func pretty_print_fs(fs []fs_entry, indent int) {
	// Calculate size of root directory
	if indent == 0 {
		root_size := 0
		for _, entry := range fs {
			root_size += entry.size
		}
		fmt.Printf("/ %d\n", root_size)
	}
	for _, entry := range fs {
		fmt.Printf("%s%s %d\n", strings.Repeat(" ", indent+2), entry.name, entry.size)
		if !entry.is_file {
			pretty_print_fs(entry.entries, indent+2)
		}
	}
}

/*
Not a part of the problem, finds all combinations of directories that are less than max_size
*/
func find_combinations_lt(curr []fs_entry, fs []fs_entry, max_size int) [][]fs_entry {
	max_size -= sum_fs(curr)
	r := make([][]fs_entry, 0)
	for _, entry := range fs {
		already_included := false
		for _, e := range curr {
			if e.name == entry.name {
				already_included = true
				break
			}
		}

		if !already_included {
			if entry.size < max_size {
				// Entry is valid combination
				// Copy curr
				c := make([]fs_entry, len(curr))
				copy(c, curr)
				r = append(r, append(c, entry))
				r = append(r, find_combinations_lt(append(c, entry), fs, max_size)...)
			}
		}
	}

	// Remove duplicates and single element combinations
	existing_lengths := make(map[int]bool)

	for i := len(r) - 1; i >= 0; i-- {
		comb := r[i]
		remove_comb := false
		if len(comb) == 1 {
			remove_comb = true
		} else if existing_lengths[sum_fs(comb)] {
			remove_comb = true
		} else {
			existing_lengths[sum_fs(comb)] = true
		}

		if remove_comb {
			r = append(r[:i], r[i+1:]...)
		}
	}

	return r
}

func get_all_directories(fs []fs_entry, max_size int) []fs_entry {
	r := make([]fs_entry, 0)
	for _, entry := range fs {
		if !entry.is_file {
			subdirectories := get_all_directories(entry.entries, max_size)
			entry.entries = make([]fs_entry, 0)
			if entry.size < max_size {
				r = append(r, entry)
			}
			r = append(r, subdirectories...)
		}
	}
	return r
}

func append_to_dir(cd []int, r []fs_entry, a []fs_entry) []fs_entry {
	if len(cd) == 0 {
		r = append(r, a...)
	} else {
		r[cd[0]].entries = append_to_dir(cd[1:], r[cd[0]].entries, a)
		// Calculate size of directory
		r[cd[0]].size = sum_fs(r[cd[0]].entries)
	}
	return r
}

func sum_fs(fs []fs_entry) int {
	r := 0
	for _, entry := range fs {
		r += entry.size
	}
	return r
}

func parse_filesystem(terminal []map[string][]string) []fs_entry {
	r := make([]fs_entry, 0)
	var cd []int // List of indices of the current directory in the filesystem
	for _, command := range terminal {
		switch command["command"][0] {
		case "cd":
			nd := command["command"][1]
			if nd == ".." {
				// Pop last element of cd
				if len(cd) > 0 {
					cd = cd[:len(cd)-1]
				}
			} else if nd == "/" {
				cd = make([]int, 0)
			} else {
				// Find index of directory in current directory
				tr := r
				for _, i := range cd {
					tr = tr[i].entries
				}

				for i, entry := range tr {
					if entry.name == nd {
						cd = append(cd, i)
						break
					}
				}
			}
		case "ls":
			files := parse_ls_output(command["output"])
			r = append_to_dir(cd, r, files)
		}
	}

	return r
}

func parse_ls_output(output []string) []fs_entry {
	// Returns a list of maps structured as follows:
	// [
	// 	{
	// 		"name": size",
	// 		...
	// 	},
	// ]

	r := make([]fs_entry, 0)

	// Use fmt.Sscanf to parse the output
	for _, line := range output {
		if line[0] == 'd' {
			var name string
			fmt.Sscanf(line, "dir %s", &name)
			r = append(r, fs_entry{name: name, size: 0, is_file: false, entries: make([]fs_entry, 0)})
		} else {
			var name string
			var size int
			fmt.Sscanf(line, "%d %s", &size, &name)
			r = append(r, fs_entry{name: name, size: size, is_file: true, entries: []fs_entry{}})
		}
	}

	return r
}

func parse_terminal(input []string) []map[string][]string {
	// Returns list of maps structured as follows:
	// {
	// 	"command": ["command", "arg1", "arg2", ...],
	//  "output": ["output_line1", "output_line2", ...]
	// }

	r := make([]map[string][]string, 0)
	was_command := false

	for _, line := range input {
		if line[0] == '$' {
			r = append(r, map[string][]string{"command": strings.Split(line, " ")[1:], "output": make([]string, 0)})
			was_command = true
		} else if was_command {
			// This is an output line
			r[len(r)-1]["output"] = append(r[len(r)-1]["output"], line)
		}
	}

	return r
}
