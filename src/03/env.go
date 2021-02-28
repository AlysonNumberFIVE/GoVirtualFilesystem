
package main


func newPwd(dir string) {
	newPath := env["PWD"] + "/" + dir
	env["OLDPWD"] = env["PWD"]
	env["PWD"] = newPath
}

func backTrackPwd(path string) {

	var newPath bytes.Buffer

	steps := strings.Split(path, "/") 
	currentDir := strings.Split(env["PWD"], "/")
	totalDepth := len(currentDir) - len(steps) + 1
	newPath.WriteString(currentDir[0])
	counter := 1
	for counter < totalDepth {
		newPath.WriteString("/" + currentDir[counter])
		counter++
	}
	env["OLDPWD"] = env["PWD"]
	env["PWD"] = newPath.String()
}

func swapPwd() {
	tmp := env["PWD"]
	env["PWD"] = env["OLDPWD"]
	env["OLDPWD"] = tmp
}
