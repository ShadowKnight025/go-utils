package go-utils

/*
   TODO:
   Replace print statements with logging.
   Add methods for RDWR mode, Append.
   Maybe add methods for symlinks as well.
   
*/

import(
	"fmt"
	"os"
)

var file_mode_map = map[string]int{
	"READ":      os.O_RDONLY,
	"WRITE":     os.O_WRONLY,
	"READWRITE": os.O_RDWR,
	"APPEND":    os.O_APPEND,
}


// create file
func create(file_name string){

	// use os.Create() to create a file if doesn't exist.

	file, err := os.Create(file_name)
	if err != nil{
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	fmt.Println("File Created Successfully!")
}

// open file utils module; add to utils package;
func write(file_name string){

	/*
             defer keyword allows a function to postpone
	     the execution of a statement until the
	     surrounding function has completed.

	     mode - READ, WRITE, READWRITE, APPEND
	*/

	// use os.Create() to create a file if doesn't exist.

	file, err := os.OpenFile(file_name, file_mode_map["WRITE"], os.ModePerm)
	if err != nil{
		fmt.Println("Error opening file:", err)
		return
	}

	_,err = file.WriteString("Hello World!\n")
	if err != nil{
		fmt.Println("Error writing file", err)
	}
	defer file.Close()
}

func read(file_name string, buffer []byte){

	file, err := os.OpenFile(file_name, file_mode_map["READ"], os.ModePerm)
	if err != nil{
		fmt.Println("Error opening file: ", err)
	}
	defer file.Close()
	file.Read(buffer)
}
