package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

type DotFile struct {
	ToInstall string
	Target    string
	CreateDir bool
	IsDir     bool
}

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg[0]) == 0 {
		log.Fatal("Please specify directory")
	}

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	workdir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	workdir = workdir + "/"

	directory := argsWithoutProg[0] + "/"
	directory = strings.Replace(directory, "//", "/", 1)

	base := workdir + "/" + argsWithoutProg[0]

	base, err = filepath.Abs(base)

	if _, err := os.Stat(directory); os.IsNotExist(err) {
		log.Fatal(err)
		os.Exit(1)
	}

	directory, err = filepath.Abs(directory)

	if err != nil {
		log.Fatal(err)
	}

	var files []DotFile

	err = filepath.Walk(directory,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}
			//fmt.Println(path, info.Size())
			files = append(files, DotFile{
				ToInstall: path,
				Target:    strings.Replace(path, base, usr.HomeDir, 1),
				CreateDir: false,
				IsDir:     false,
			})
			return nil
		})

	if err != nil {
		log.Fatal(err)
	}

	//adding scripts
	scriptFiles, err := GetScripts(workdir, usr.HomeDir+"/")
	if err != nil {
		log.Fatal(err)
	}

	files = append(files, scriptFiles...)

	toDo, errors := summary(files)
	if len(errors) > 0 {
		fmt.Println("###################################3")
		fmt.Println("Errors, resolve it before continuing")
		fmt.Println("###################################3")
		for _, summaryError := range errors {
			fmt.Println(summaryError)
		}

		os.Exit(1)

	}

	fmt.Println("###################################3")
	fmt.Println("SUMMARY")
	fmt.Println("###################################3")
	for _, task := range toDo {
		fmt.Println(task)
	}

	text := "N"
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Apply ? [y/N]: ")
	text, _ = reader.ReadString('\n')

	text = strings.TrimSpace(text)

	if text == "y" || text == "Y" {
		fmt.Println("Applying files")
		// if summary is ok, we can proceed
		for _, file := range files {
			Apply(file)
		}
	}

}

func GetScripts(workDir string, homeDir string) ([]DotFile, error) {

	var files []DotFile
	scriptFolder := workDir + "scripts"

	err := filepath.Walk(scriptFolder,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}
			//fmt.Println(path, info.Size())
			files = append(files, DotFile{
				ToInstall: path,
				Target:    strings.Replace(path, scriptFolder, homeDir+".local/bin", 1),
				CreateDir: false,
			})
			return nil
		})

	return files, err
}

//Apply the config
func Apply(file DotFile) {
	if file.IsDir {
		if file.CreateDir {
			fmt.Println("Creating folder " + file.Target)
			os.MkdirAll(file.Target, os.ModePerm)
		} else {
			fmt.Println("folder", file.Target, "already exists")
		}
	} else {
		fmt.Println("Creating symlink", file.ToInstall, "->", file.Target)
		os.Symlink(file.ToInstall, file.Target)
	}
}

func summary(files []DotFile) ([]string, []error) {

	var checkErrors []error
	var toDo []string

	for index, file := range files {

		source, sourceErr := os.Lstat(file.ToInstall)
		//checking if current file already exists
		fi, err := os.Lstat(file.Target)
		if err != nil {

			if sourceErr == nil {
				switch sourceMode := source.Mode(); {
				case sourceMode.IsDir():
					file.CreateDir = true
					file.IsDir = true
					files[index] = file

					toDo = append(toDo, file.Target+" will be created as a directory")
					break
				case sourceMode.IsRegular():
					toDo = append(toDo, file.Target+" doesn't exist, symlink will be created")
					break
				}
			}

			continue
		}

		if sourceErr == nil && source.IsDir() {
			file.IsDir = true
			files[index] = file
		}

		readlink, err2 := os.Readlink(file.Target)
		if err2 == nil {
			if file.ToInstall != readlink {
				checkErrors = append(checkErrors, errors.New(file.Target+" is already a symlink pointing to "+readlink))
			} else {
				toDo = append(toDo, file.Target+" is already pointing to "+file.ToInstall+" so it will be replaced")
			}
			continue
		}

		switch mode := fi.Mode(); {
		case mode.IsDir():
			//checkErrors = append(checkErrors, errors.New(file.Target+" is already a symlink pointing to "+readlink))
			break
		case mode.IsRegular():
			// do file stuff
			checkErrors = append(checkErrors, errors.New(file.Target+" is already an existing regular file"))
			break
		}

	}

	return toDo, checkErrors
}
