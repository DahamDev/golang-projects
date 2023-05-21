package main

func main(){
	myMap := map[string]string{
		"David":"teacher",
		"John":"engineer",
		"Peter":"plumber",
	}

	for name,employee := range myMap{
		println(name + " is a "+ employee)
	}

	println()

	davidsJob := myMap["David"]
	println("David job is "+ davidsJob)

	//david chagne the job
	myMap["David"] = "doctor"
	davidsNewJob := myMap["David"]
	println("David  new job job is "+ davidsNewJob)

}