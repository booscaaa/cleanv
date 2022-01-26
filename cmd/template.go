/*
Copyright © 2022 Vinícius Boscardin

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"embed"
	"fmt"
	"os"
	"strings"

	"github.com/booscaaa/cleanv/util"
	"github.com/spf13/cobra"
)

type DefaultTemplate struct {
	Screen       string
	Repositories []string
	Controller   string
}

var TemplateFs embed.FS

// templateCmd represents the template command
var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "Create a default structure of a new screen.",
	Long:  `Create a default structure of a new screen.`,
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat("src"); os.IsNotExist(err) {
			fmt.Println("src folder not found, check the location where you are running the program!")
			return
		}

		screen, _ := cmd.Flags().GetString("screen")
		module, _ := cmd.Flags().GetString("module")
		repositories, _ := cmd.Flags().GetString("repositories")

		if screen == "" {
			fmt.Println("Enter the screen name")
			return
		}

		if module == "" {
			fmt.Println("Enter the module name")
			return
		}

		if repositories == "" {
			fmt.Println("Enter the repository name")
			return
		}

		defaultTemplate := DefaultTemplate{
			Screen:       screen,
			Repositories: strings.Split(repositories, ","),
			Controller:   strings.Title(strings.ToLower(screen)),
		}

		viewPath := "src/module/" + module + "/" + screen + "/view"
		controllerPath := "src/module/" + module + "/" + screen + "/controller"
		diPath := "src/module/" + module + "/" + screen + "/di"
		usecasePath := "src/module/" + module + "/" + screen + "/domain/usecase"
		modelPath := "src/module/" + module + "/" + screen + "/domain/model"
		repositoryPath := "src/module/" + module + "/" + screen + "/data/repository"

		if _, err := os.Stat(viewPath); os.IsNotExist(err) {
			os.MkdirAll(viewPath, os.ModePerm)
		}

		if _, err := os.Stat(controllerPath); os.IsNotExist(err) {
			os.MkdirAll(controllerPath, os.ModePerm)
		}

		if _, err := os.Stat(diPath); os.IsNotExist(err) {
			os.MkdirAll(diPath, os.ModePerm)
		}

		if _, err := os.Stat(usecasePath); os.IsNotExist(err) {
			os.MkdirAll(usecasePath, os.ModePerm)
		}

		if _, err := os.Stat(modelPath); os.IsNotExist(err) {
			os.MkdirAll(modelPath, os.ModePerm)
		}

		if _, err := os.Stat(repositoryPath); os.IsNotExist(err) {
			os.MkdirAll(repositoryPath, os.ModePerm)
		}

		//Files
		if _, err := os.Stat(viewPath + "/" + screen + ".vue"); os.IsNotExist(err) {
			util.PopulateFiles(viewPath+"/"+screen+".vue", "vue.embed", "default", defaultTemplate)
		}

		if _, err := os.Stat(controllerPath + "/" + screen + "Controller.js"); os.IsNotExist(err) {
			util.PopulateFiles(controllerPath+"/"+screen+"Controller.js", "controller.embed", "default", defaultTemplate)
		}

		if _, err := os.Stat(diPath + "/" + "di.js"); os.IsNotExist(err) {
			util.PopulateFiles(diPath+"/"+"di.js", "di.embed", "default", defaultTemplate)
		}

		if _, err := os.Stat(diPath + "/" + "axios.js"); os.IsNotExist(err) {
			util.PopulateFiles(diPath+"/"+"axios.js", "axios.embed", "default", defaultTemplate)
		}

		if _, err := os.Stat(usecasePath); os.IsNotExist(err) {
			os.MkdirAll(usecasePath, os.ModePerm)
		}

		if _, err := os.Stat(modelPath + "/" + screen + ".js"); os.IsNotExist(err) {
			util.PopulateFiles(modelPath+"/"+screen+".js", "model.embed", "default", defaultTemplate)
		}

		for _, repository := range defaultTemplate.Repositories {
			if _, err := os.Stat(repositoryPath + "/" + repository + "Repository.js"); os.IsNotExist(err) {
				util.PopulateFiles(repositoryPath+"/"+repository+"Repository.js", "repository.embed", "default", repository)
			}

			if _, err := os.Stat(usecasePath + "/" + repository + "UseCase.js"); os.IsNotExist(err) {
				util.PopulateFiles(usecasePath+"/"+repository+"UseCase.js", "usecase.embed", "default", repository)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(templateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// templateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	templateCmd.Flags().StringP("screen", "s", "", "The screen name to be created")
	templateCmd.Flags().StringP("module", "m", "", "The module name to be created")
	templateCmd.Flags().StringP("repositories", "r", "", "Name of repositories to be created separated by comma")
}
