package osexample

// Create example
func Create() {

	// envs := os.Environ()

	// for _, ev := range envs {
	// 	fmt.Println(ev)
	// }

	// str := os.ExpandEnv("${APPDATA}")
	// fmt.Println(str)

	// str := os.Getenv("APPDATA")
	// fmt.Println(str)

	// uid := os.Getuid()
	// euid := os.Geteuid()

	// fmt.Println(uid, "-", euid)

	// gid := os.Getgid()
	// fmt.Println(gid)

	// grps, err := os.Getgroups()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(grps)

	// pageSize := os.Getpagesize()
	// fmt.Println(pageSize)

	// ppid := os.Getppid()
	// pid := os.Getpid()
	// fmt.Println(ppid, pid)

	// dir, err := os.Getwd()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(dir)

	// hostname, err := os.Hostname()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(hostname)

	// fmt.Printf("%s", string(os.PathSeparator))
	// fmt.Printf("%t", os.IsPathSeparator('\\'))

	// err := os.Link("test1.txt", "text3.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%t", os.IsExist(err))

	// err := os.Mkdir("ttt", 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err := os.MkdirAll("a/b/c", 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// f, err := os.Create("./test1.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// f.Close()

	// f2, err := os.OpenFile("./test2.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// f2.Close()

	// d1, err := os.Open("test")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = d1.Chdir()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// f3, err := os.Create("Test123")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// f3.Close()

	// for {
	// 	dlist, err := d1.Readdir(2)
	// 	if err == io.EOF {
	// 		fmt.Println("io.EOF")
	// 		break

	// 	} else if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	for _, fi := range dlist {
	// 		fmt.Printf("%s\n", fi.Name())
	// 	}

	// }

	// d1.Close()

	// f2, err := os.Open("test")
	// finfo, err := f2.Stat()

	// fmt.Printf("Name: %s, Size: %d, Mode: %d, IsDir: %t", finfo.Name(), finfo.Size(), finfo.Mode(), finfo.IsDir())
	// f2.Close()
}
