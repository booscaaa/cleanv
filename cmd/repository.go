/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/booscaaa/cleanv/util"
	"github.com/spf13/cobra"
)

// repositoryCmd represents the repository command
var repositoryCmd = &cobra.Command{
	Use:   "repository",
	Short: "Creates a repository with linked usecase. Create a reference of both in DI and inject in controller",
	Long:  `Creates a repository with linked usecase. Create a reference of both in DI and inject in controller`,
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat("src"); os.IsNotExist(err) {
			fmt.Println("src folder not found, check the location where you are running the program")
			return
		}

		screen, _ := cmd.Flags().GetString("screen")
		module, _ := cmd.Flags().GetString("module")
		name, _ := cmd.Flags().GetString("name")
		delete, _ := cmd.Flags().GetBool("delete")

		if screen == "" {
			fmt.Println("Enter the screen name")
			return
		}

		if module == "" {
			fmt.Println("Enter the module name")
			return
		}

		if name == "" {
			fmt.Println("Enter the repository name")
			return
		}

		if delete {
			deleteData(screen, module, name)
		} else {
			createData(screen, module, name)
		}

	},
}

func createData(screen, module, name string) {
	usecasePath := "src/module/" + module + "/" + screen + "/domain/usecase"
	repositoryPath := "src/module/" + module + "/" + screen + "/data/repository"
	diPath := "src/module/" + module + "/" + screen + "/di"
	controllerPath := "src/module/" + module + "/" + screen + "/controller"

	if _, err := os.Stat(usecasePath); os.IsNotExist(err) {
		os.MkdirAll(usecasePath, os.ModePerm)
	}

	if _, err := os.Stat(repositoryPath); os.IsNotExist(err) {
		os.MkdirAll(repositoryPath, os.ModePerm)
	}

	if _, err := os.Stat(repositoryPath + "/" + name + "Repository.js"); os.IsNotExist(err) {
		util.PopulateFiles(repositoryPath+"/"+name+"Repository.js", "repository.embed", "single-command-repository", name)
	}

	if _, err := os.Stat(usecasePath + "/" + name + "UseCase.js"); os.IsNotExist(err) {
		util.PopulateFiles(usecasePath+"/"+name+"UseCase.js", "usecase.embed", "single-command-repository", name)
	}

	tmpl := template.New("import-di.embed").Delims("[[", "]]")
	tmpl, _ = tmpl.ParseFS(TemplateFs, "templates/single-command-repository/import-di.embed")

	var importDiText bytes.Buffer
	tmpl.Execute(&importDiText, name)

	importDiTextResult := importDiText.String()

	tmpl = template.New("config-inject-di.embed").Delims("[[", "]]")
	tmpl, _ = tmpl.ParseFS(TemplateFs, "templates/single-command-repository/config-inject-di.embed")

	var configinjectDiText bytes.Buffer
	tmpl.Execute(&configinjectDiText, name)

	configinjectDiTextResult := configinjectDiText.String()

	newDI, _ := os.Create(diPath + "/newdi.js")

	file, _ := os.OpenFile(diPath+"/di.js", os.O_RDWR, 0644)

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		text := sc.Text()
		newDI.WriteString(text + "\n")

		if strings.Contains(text, "import axios") {
			newDI.WriteString("\n")
			newDI.WriteString(importDiTextResult)
		}

		if strings.Contains(text, "const baseUrl") {
			newDI.WriteString("\n")
			newDI.WriteString(configinjectDiTextResult + "\n")
		}

		if strings.Contains(text, "context,") {
			newDI.WriteString("\t\t")
			newDI.WriteString(name + "UseCaseImpl," + "\n")
		}

	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return
	}

	os.Rename(diPath+"/newdi.js", diPath+"/di.js")

	file.Close()

	newController, _ := os.Create(controllerPath + "/newController.js")

	file, _ = os.OpenFile(controllerPath+"/"+screen+"Controller.js", os.O_RDWR, 0644)

	sc = bufio.NewScanner(file)

	for sc.Scan() {
		text := sc.Text()
		newController.WriteString(text + "\n")
		if strings.Contains(text, "context,") {
			newController.WriteString("\t\t")
			newController.WriteString(name + "UseCase," + "\n")
		}

		if strings.Contains(text, "this.context = context") {
			newController.WriteString("\t\t")
			newController.WriteString("this." + name + "UseCase = " + name + "UseCase" + "\n")
		}
	}

	os.Rename(controllerPath+"/newController.js", controllerPath+"/"+screen+"Controller.js")

	file.Close()
}

func deleteData(screen, module, name string) {
	usecasePath := "src/module/" + module + "/" + screen + "/domain/usecase"
	repositoryPath := "src/module/" + module + "/" + screen + "/data/repository"
	diPath := "src/module/" + module + "/" + screen + "/di"
	controllerPath := "src/module/" + module + "/" + screen + "/controller"

	if _, err := os.Stat(usecasePath + "/" + name + "UseCase.js"); !os.IsNotExist(err) {
		os.Remove(usecasePath + "/" + name + "UseCase.js")
	}

	if _, err := os.Stat(repositoryPath + "/" + name + "Repository.js"); !os.IsNotExist(err) {
		os.Remove(repositoryPath + "/" + name + "Repository.js")
	}

	newDI, _ := os.Create(diPath + "/newdi.js")

	file, _ := os.OpenFile(diPath+"/di.js", os.O_RDWR, 0644)

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		text := sc.Text()

		if !strings.Contains(text, name+"UseCase") && !strings.Contains(text, name+"Repository") {
			newDI.WriteString(text + "\n")
		}
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return
	}

	os.Rename(diPath+"/newdi.js", diPath+"/di.js")

	file.Close()

	newController, _ := os.Create(controllerPath + "/newController.js")

	file, _ = os.OpenFile(controllerPath+"/"+screen+"Controller.js", os.O_RDWR, 0644)

	sc = bufio.NewScanner(file)

	for sc.Scan() {
		text := sc.Text()

		if !strings.Contains(text, name+"UseCase") {
			newController.WriteString(text + "\n")
		}
	}

	os.Rename(controllerPath+"/newController.js", controllerPath+"/"+screen+"Controller.js")

	file.Close()
}

func init() {
	rootCmd.AddCommand(repositoryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// repositoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	repositoryCmd.Flags().StringP("screen", "p", "", "Nome do screen a ser criado")
	repositoryCmd.Flags().StringP("module", "m", "", "Nome do module a ser criado")
	repositoryCmd.Flags().StringP("name", "n", "", "Nome do repository a ser criado")
	repositoryCmd.Flags().BoolP("delete", "d", false, "Deletar repository")
}
