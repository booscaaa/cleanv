<p align="center">
  <h1 align="center">Cleanv - Golang SDK for Vue Projects</h1>
  <p align="center">
    <a href="https://pkg.go.dev/github.com/booscaaa/cleanv"><img alt="Reference" src="https://img.shields.io/badge/go-reference-purple?style=for-the-badge"></a>
    <a href="https://github.com/booscaaa/cleanv/releases/latest"><img alt="Release" src="https://img.shields.io/github/v/release/booscaaa/cleanv.svg?style=for-the-badge"></a>
    <a href="/LICENSE"><img alt="Software License" src="https://img.shields.io/badge/license-MIT-red.svg?style=for-the-badge"></a>
    <a href="https://github.com/booscaaa/cleanv/actions/workflows/test.yaml"><img alt="Test status" src="https://img.shields.io/github/workflow/status/booscaaa/cleanv/Test?label=TESTS&style=for-the-badge"></a>
    <a href="https://codecov.io/gh/booscaaa/cleanv"><img alt="Coverage" src="https://img.shields.io/codecov/c/github/booscaaa/cleanv/master.svg?style=for-the-badge"></a>
  </p>
</p>

<br>

## Why?

This project is part of my personal portfolio, so, I'll be happy if you could provide me any feedback about the project, code, structure or anything that you can report that could make me a better developer!

Email-me: boscardinvinicius@gmail.com

Connect with me at [LinkedIn](https://www.linkedin.com/in/booscaaa/).

<br>


## Funcionalidades

- Creating a standard development framework
- Creating separate files within existing code structures
- Automatically update

<br>

## Começando


### Instalação
<br>

**Download the library and configure it in the path/environment variables of your operating system**

<br>

Linux:
- https://github.com/booscaaa/cleanv/releases/latest/download/cleanv

```sh
mv ~/Downloads/cleanv /usr/local/cleanv/
nano ~/.bashrc

# paste the line below at the end of the file
PATH=$PATH:/usr/local/cleanv

# give general permissions on the folder
sudo chmod 777 /usr/local/cleanv

# reload bash info
source ~/.bashrc

# add autocomplete in bash
cleanv completion bash > /tmp/completion
source /tmp/completion
```

<br>

Windows:
- https://github.com/booscaaa/cleanv/releases/latest/download/cleanv.exe

- Put the executable in the C:/cleanv folder and set the path in the environment variables

<br>

<br>



The initial struct of vue app code like this:

![Struct1](../master/assets-readme/struct1.png?raw=true)

<br>
<br>

### With commands

# cleanv template

<br>

This command will generate the complete structure for developing a new program on the web.


Positioned in the root of the project with vue.js run the command:

```sh
cleanv template -s payment001 -m payment -r findPayment
```
![Prompt1](../master/assets-readme/struct2.png?raw=true)

<br>

# cleanv repository

<br>

This command will generate the structure of a new api call within a ready-made structure. Importing and injecting dependencies if necessary.


Positioned in the root of the project with vue.js run the command:

```sh
cleanv repository -s payment001 -m payment -n findTickets
```
![Prompt1](../master/assets-readme/struct3.png?raw=true)

<br>
<br>

**Caution!!!** Using the -d flag will delete the corresponding files.
```sh
cleanv repository -s screen_name -m module_name -n findSomething1 -d
```

<br>

# cleanv update

<br>

This command updates the sdk binary automatically;

```sh
cleanv update
```
<br>
<br>
<br>

## Contributing

You can send how many PR's do you want, I'll be glad to analyze and accept them! And if you have any question about the project...

Email-me: boscardinvinicius@gmail.com

Connect with me at [LinkedIn](https://www.linkedin.com/in/booscaaa/)

Thank you!

## License

This project is licensed under the MIT License - see the [LICENSE.md](https://github.com/booscaaa/cleanv/blob/master/LICENSE) file for details