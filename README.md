# HottoCoffee

HottoCoffee is an open-source software (OSS) service designed to manage batches.  
It provides an easy way to manage batch history by registering batches and execution schedules, and receiving execution results as HTTP requests from the server.

## Getting Started

### Start Application

To get started with HottoCoffee, follow these steps:

1. Clone the HottoCoffee repository to your local machine.
2. Start application with `docker compose up -d`
3. Begin managing your batches!

### Start development

1. Install following applications.
   |Application|version|url|
   |---|---|---|
   |Go|1.19|https://go.dev/doc/install|
   |Node.js|18|https://nodejs.org/en|
2. Run commands
```shell
cd frontend
npm install
npm run dev
```
```shell
cd back
```


## How it Works

HottoCoffee is designed to make managing the batch process simple and efficient.
Here's how it works:

1. Register batches and execution schedules with HottoCoffee. This can be done from the web interface.
2. when the batch is complete, HottoCoffee receives an HTTP request from your server informing it of the batch's completion.
3. you can then check the batch status and execution history in the HottoCoffee web interface.

## Contributions

HottoCoffee is an open source project and welcomes contributions from anyone!
If you would like to contribute to HottoCoffee, please see the guidelines in CONTRIBUTING.md.

## Problem and Bug Reports

If you encounter a problem or bug while using HottoCoffee, please create a new issue in the HottoCoffee repository.
We will respond as quickly as possible.

## License

HottoCoffee is released under the MIT License. Please see the LICENSE file for more information.
