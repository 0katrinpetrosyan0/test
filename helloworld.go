package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func createConfigFile() {
	f, err := os.Create(".pre-commit-config.yaml")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	words := []string{"repos:", "- repo: https://github.com/pre-commit/pre-commit-hooks", "  rev: v4.2.0", "  hooks:", "    - id: check-added-large-files", "    - id: check-merge-conflict", "    - id: check-vcs-permalinks", "    - id: end-of-file-fixer", "    - id: trailing-whitespace", "      args: [--markdown-linebreak-ext=md]", "      exclude: CHANGELOG.md", "    - id: check-yaml", "    - id: check-merge-conflict", "    - id: check-executables-have-shebangs", "    - id: check-case-conflict", "    - id: mixed-line-ending", "      args: [--fix=lf]", "    - id: detect-aws-credentials", "      args: ['--allow-missing-credentials']", "    - id: detect-private-key"}

	for _, word := range words {

		_, err := f.WriteString(word + "\n")

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("done")
}

func createDetectPreCommitFIle() {
	f, err := os.Create("pre-commit")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	words := []string{"#!/bin/bash", "echo 'Start git pre-commit hooks and checks... ';", "sudo pre-commit run -a;", "echo 'End git pre-commit hooks and checks... ';"}

	for _, word := range words {

		_, err := f.WriteString(word + "\n")

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("done")
}

func createFmtPreCommitFIle() {
	f, err := os.Create("pre-commit")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	words := []string{"#!/bin/bash", "echo 'Start terraform fmt';", "for y in $(find . -name '*.tf' -print)", "    do", "      if ! git diff --cached --quiet '${y}'", "      then", "        echo 'Formatting the file: ${y}';", "        terraform fmt -write=true '${y}'", "      fi", "done", "echo 'End terraform fmt';"}
	for _, word := range words {

		_, err := f.WriteString(word + "\n")

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("done")
}

func createDocsPreCommitFIle() {
	f, err := os.Create("pre-commit")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	words := []string{"#!/bin/bash", "echo 'Start terraform docs';", "for i in `ls -d modules/*/`;do", "    if ! git diff --cached --quiet '${i}'", "    then", "      echo 'Generating  doc file: ${i}';", "      terraform-docs markdown table --output-file README.md --output-mode inject '${i}';", "    fi", "done", "echo 'End terraform docs';"}

	for _, word := range words {

		_, err := f.WriteString(word + "\n")

		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("done")
}

func createSecretFolder() {
	if err := os.Mkdir("githooks-detect-secrets", os.ModePerm); err != nil {
		log.Fatal(err)
	}
	file := os.Chdir("githooks-detect-secrets")
	fmt.Println(file)
	createDetectPreCommitFIle()
	cmd := os.Chmod("pre-commit", 0777)
	fmt.Println(cmd)
	file1 := os.Chdir("../")
	fmt.Println(file1)

}

func createDocsFolder() {
	if err := os.Mkdir("githooks-terraform-docs", os.ModePerm); err != nil {
		log.Fatal(err)
	}
	file := os.Chdir("githooks-terraform-docs")
	fmt.Println(file)
	createDocsPreCommitFIle()
	cmd := os.Chmod("pre-commit", 0777)
	fmt.Println(cmd)
	file1 := os.Chdir("../")
	fmt.Println(file1)
}

func createFmtFolder() {
	if err := os.Mkdir("githooks-terraform-fmt", os.ModePerm); err != nil {
		log.Fatal(err)
	}
	file := os.Chdir("githooks-terraform-fmt")
	fmt.Println(file)
	createFmtPreCommitFIle()
	cmd := os.Chmod("pre-commit", 0777)
	fmt.Println(cmd)
	file1 := os.Chdir("../")
	fmt.Println(file1)
}

func installPreCommit() {
	if runtime.GOOS == "linux" {
		app := "pip"

		arg0 := "install"
		arg1 := "pre-commit"

		cmd := exec.Command(app, arg0, arg1)
		stdout, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println(string(stdout))
	} else if runtime.GOOS == "windows" {
		app := "conda"

		arg0 := "install"
		arg1 := "-c"
		arg2 := "conda-forge"
		arg3 := "pre-commit"

		cmd := exec.Command(app, arg0, arg1, arg2, arg3)
		stdout, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println(string(stdout))
	} else {
		app := "brew"

		arg0 := "install"
		arg1 := "pre-commit"

		cmd := exec.Command(app, arg0, arg1)
		stdout, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println(string(stdout))
	}
}

func configPreCommitGlobally() {
	app := "git"

	arg0 := "config"
	arg1 := "--get"
	arg2 := "core.hooksPath"
	arg3 := "githooks-detect-secrets"

	cmd := exec.Command(app, arg0, arg1,arg2,arg3)
	fmt.Println(cmd)
}

func main() {
	createConfigFile()
	createSecretFolder()
	createDocsFolder()
	createFmtFolder()
	configPreCommitGlobally()
	installPreCommit()
}
