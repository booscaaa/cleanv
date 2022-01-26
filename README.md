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


## Functionalities

- Creating a standard development framework
- Creating separate files within existing code structures
- Automatically update

<br>

## Starting


### Installation
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

### Examples of code struct generation
```sh
cleanv template -s payment001 -m payment -r findPayment
```
```
- module
  - payment
    - payment001
      - di
        - di.js
        - axios.js
      - controller
        - paymentController001.js
      - data
        - repository
          - findPaymentRepository.js
      - domain
        - model
          - payment001.js
        - usecase
          - findPaymentUseCase.js
      - view
        - payment001.vue
```   

```js
//Repository
const findPaymentRepository = (axios) => async () => {
  try {
    const response = await axios.get('/rest/TODO')
    return response //TODO
  } catch (error) {
    throw error
  }
}

export default findPaymentRepository
```

```js
//Usecase
const findPaymentUseCase = (repository) => async () => {
  try {
    //TODO
    return await repository()
  } catch (error) {
    throw error
  }
}

export default findPaymentUseCase
```
```js
//Model
class Payment001 {
    constructor() {}
}

export default Payment001
```
```js
//Dependencie Injection
import axiosInstance from './axios'

import findPaymentRepository from '../data/repository/findPaymentRepository'
import findPaymentUseCase from '../domain/usecase/findPaymentUseCase'

import Payment001Controller from '../controller/payment001Controller'

const instance = axiosInstance

const findPaymentRepositoryImpl = findPaymentRepository(instance)
const findPaymentUseCaseImpl = findPaymentUseCase(findPaymentRepositoryImpl)

const payment001Controller = (context) =>
  new Payment001Controller(
    context,
	  findPaymentUseCaseImpl,
  )

export { payment001Controller }
```

```js
//Axios instance for Dependencie Injection
import axios from 'axios'

const axiosInstace = axios.create({
  baseURL: process.env.VUE_APP_API_BASE_URL,
  headers: {
    Accept: 'application/json',
    'Content-Type': 'application/json',
    Access: 'application/json',
  },
})

export default axiosInstace
```

```js
//Controller
class Payment001Controller {

  constructor(
    context,
    findPaymentUseCase,
  ) {
    this.context = context
    this.findPaymentUseCase = findPaymentUseCase
  }

  async mounted() {
    try {
      //TODO
    } catch (error) {
      //HANDLER ERROR
    }
  }
}

export default Payment001Controller
```
```html
//Vue Screen
<template>
  <div>
    <!-- your code here -->
  </div>
</template>

<script>
import { payment001Controller } from "../di/di";
export default {
  data: (context) => ({
    controller: payment001Controller(context),
  }),
  mounted() {
    this.controller.mounted();
  },
};
</script>

```

### Adding new repository
```sh
cleanv repository -s payment001 -m payment -n findTickets
```
```
- module
  - payment
    - payment001
      - di
        - di.js
        - axios.js
      - controller
        - paymentController001.js
      - data
        - repository
          - findPaymentRepository.js
          - findTicketsRepository.js
      - domain
        - model
          - payment001.js
        - usecase
          - findPaymentUseCase.js
          - findTicketsUseCase.js
      - view
        - payment001.vue
```  
```js
//Repository
const findTicketsRepository = (axios) => async () => {
  try {
    const response = await axios.get('/rest/TODO')
    return response //TODO
  } catch (error) {
    throw error
  }
}

export default findTicketsRepository
```

```js
//Usecase
const findTicketsUseCase = (repository) => async () => {
  try {
    //TODO
    return await repository()
  } catch (error) {
    throw error
  }
}

export default findTicketsUseCase
```
```js
//Dependencie Injection
import axiosInstance from './axios'

import findTicketsRepository from '../data/repository/findTicketsRepository'
import findTicketsUseCase from '../domain/usecase/findTicketsUseCase'

import findPaymentRepository from '../data/repository/findPaymentRepository'
import findPaymentUseCase from '../domain/usecase/findPaymentUseCase'

import Payment001Controller from '../controller/payment001Controller'

const instance = axiosInstance

const findTicketsRepositoryImpl = findTicketsRepository(instance)
const findTicketsUseCaseImpl = findTicketsUseCase(findTicketsRepositoryImpl)


const findPaymentRepositoryImpl = findPaymentRepository(instance)
const findPaymentUseCaseImpl = findPaymentUseCase(findPaymentRepositoryImpl)


const payment001Controller = (context) =>
  new Payment001Controller(
    context,
		findTicketsUseCaseImpl,
    findPaymentUseCaseImpl,
  )

export { payment001Controller }
```
```js
//Controller
class Payment001Controller {

  constructor(
    context,
    findTicketsUseCase
    findPaymentUseCase,
  ) {
    this.context = context
    this.findTicketsUseCase = findTicketsUseCase
    this.findPaymentUseCase = findPaymentUseCase
  }

  async mounted() {
    try {
      //TODO
    } catch (error) {
      //HANDLER ERROR
    }
  }
}

export default Payment001Controller
```




## Contributing

You can send how many PR's do you want, I'll be glad to analyze and accept them! And if you have any question about the project...

Email-me: boscardinvinicius@gmail.com

Connect with me at [LinkedIn](https://www.linkedin.com/in/booscaaa/)

Thank you!

## License

This project is licensed under the MIT License - see the [LICENSE.md](https://github.com/booscaaa/cleanv/blob/master/LICENSE) file for details