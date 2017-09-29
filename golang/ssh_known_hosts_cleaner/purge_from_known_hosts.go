package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"bufio"
	"strings"
)

func main() {
	//variables
	target_search := "10.252.186" // the default target is to remove your openstack ssh spew.
	//default_file := "/Users/pwhite/.ssh/known_hosts2"  // fix file name first for live, then change it to any users .ssh/known_hosts
	new_file_array := []string{}
	tmp_known_hosts := "tmp_known_hosts"
	change_count := 0

	// flags
	custom_target := flag.String("target", "", "Define a target for an IP or network prefix to clean out of known_hosts.")
	verbose := flag.Bool("verbose", false, "turn on verbose output for test stuffs.")
	flag.Parse()

	// check for custom flag or use default
	verbose_value := *verbose
	custom_target_value := *custom_target
	if custom_target_value != "" {
		target_search = custom_target_value
	}

	usr, err := user.Current()
	if err != nil {
		fmt.Print("Err user is broken.")
		os.Exit(1)
	}
	home := usr.HomeDir
	def_file := "/.ssh/known_hosts"
	target_file := home + def_file

	if verbose_value { // just a verbose output for dev testing.
		fmt.Println("Target search: ", target_search)
		fmt.Println("Target File:", target_file)
		fmt.Println("New File Array:", new_file_array)
		fmt.Println("Change Count:", change_count)
		fmt.Println("Custom Target Search:", custom_target_value)
	}

	// read in file ERR and exit if issues.

	readme, err := os.Open(target_file)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	defer readme.Close()

	// iterate through read file and store new_file content lines in new_file_array.
	// increment change_count for each row that matches the content search.
	scanner := bufio.NewScanner(readme)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), target_search) {
			new_file_array = append(new_file_array, scanner.Text())
		} else {
			change_count = change_count + 1
		}
	}


	// check chnage_count : if count == 0 exit with message ; if count > 0 write new file.
	if change_count == 0 {
		fmt.Printf("No instances of [%s] were found. Exiting without making changes.\n", target_search)
		os.Exit(0)
	}

    temp_out, err := os.Create(tmp_known_hosts)
    if err != nil {
    	fmt.Println(err)
    	os.Exit(2)
	}

    for write_line, err := range new_file_array {
    	temp_out.WriteString(write_line)
    	if err != nil {
    		fmt.Println(err)
		}
	}

	// todo: add second block of tests here.

	// rename newfile to old filename and set permissions.
	// todo: use os.Rename(oldpath, newpath) error
	file_rename, err := os.Rename(temp_out, target_file)
	if err != nil {
		fmt.Println(err)
	}

	// todo: use os.Chmod(name string, mode FileMode) error

	file_remode, err := os.Chmod(target_file, 0644, )
	if err != nil {
		fmt.Println(err)
	}

	// todo: use os.Chown(name string, uid, gid int) error
	file_reown, err := os.Chown(target_file, , )
	if err != nil {
		fmt.Println(err)
	}
	// announce what you did and exit.
	// todo: report output

	fmt.Print("FIN")
}


