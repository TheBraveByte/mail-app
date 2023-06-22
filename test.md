# Introduction

# Overview of Concurrent Mail Server

- Concurrency and its importance
- Concurrency and Parallelism
- Goroutine and its basic building block
- Communicating Sequential Processes(CSP): The Model of Concurrency
- Explanation using the Universe: Planet Earth and Human
- Benefits of using Goroutines and Gmail Integration

# Setting Up the Environment

- Prerequisite Knowledge
- Installing necessary tools
- Creating a Gmail account and enabling API access
- Configuring credentials and environment variables

# Building the Mail Server API

- Designing the API architecture
- Understanding the Application Model
  - Code Snippet for the model Package
- Integrating MongoDB Database
  - Setting Up the Database Connection
  - Default Collection Function
  - Database Queries
  - Interface Implementation
- Creating API endpoints and handling requests
  - Implementing Functions in the Tools Package
  - Creating Handler Method to Process HTTP
  - Implementing an Interface for Handler Methods
  - Creating Routes Endpoints(URL)


- Integrating Gmail API for sending and receiving emails
- Implementing Goroutines and Channels for Concurrency
- Intialising the Application Server in the Main function

# Testing the application API

- Starting the API Server
- Test Run

# Building Subscriber and User Interface for Mail Sending

- Creating form for mail details and upload document
- Creating form for subscriber details

# Building a Concurrent Mail Server API with Goroutines and Gmail Integration

It is fascinating to take full advantage of the multicore processor for executing and implementing concurrent processes or programs, and it's interesting doing so using the simplicity of Go and its abstractions.

In this tutorial, you will learn how to implement a practical idea using concurrency and integrate it using Go functionality and its basic building blocks. So let's dive in at the right angle.

# Overview of Concurrent Mail Server

This tutorial will show you how to set up a concurrent mail server to efficiently send emails, compared to the traditional email method without concurrency. You will also discover the benefits of using concurrency in this process.
Furthermore, you learn to design a simple interface for subscribers to provide details and for users to upload documents of the mail to send.

But before you dive into that, Let me take you alongside the coast of concurrency and goroutine to get you started.

## Concurrency and its importance

It would help if you were anxious to know the full depth of concurrency. Let's get started.

Concurrency is the act of putting together multiple independent processes or task structures that are executed simultaneously and dealing with multiple requests simultaneously, improving performance and reducing delays in execution.

Some of the advantages of concurrency include:

- Performance optimisation
- Better resource utilisation
- Improved responsiveness of the server
- Better fault tolerance

Concurrency refers to running multiple independent processes or tasks simultaneously, in which it can process and manage numerous requests simultaneously at the same time. Simultaneously processing multiple tasks can optimise performance and significantly reduce execution delays.

To avoid confusion, I'll explain the differences and similarities between concurrency and parallelism in the following sections below.

## Concurrency and Parallelism

> _**Concurrency** simplifies Synchronization; **Concurrency** enables **Parallelism**._
> _**Concurrency** means **dealing** with multiple things at once while **Parallelism** means **doing** multiple things at the same time._
>
> - _Rob Pike_

![concurrent vs parallelism Image](https://i.imgur.com/eIId6C3.png)

Concurrency and parallelism often need clarification to be the same, even though they are not. However, they both involve the simultaneous execution of processes and are sometimes used interchangeably and based on similar principles.

The focus of concurrency is designing a program to handle multiple processes concurrently in an overlapping manner. It is possible to accomplish this through various mechanisms, such as multithreading and multiprocessing.

In contrast, parallelism means executing multiple processes simultaneously by utilising processing cores and taking advantage of the resources to perform tasks in parallel.

The main focus is to understand the implications of these two phenomena for the efficient design of software applications or systems.

## Goroutine and its basic building block

Goroutines achieve concurrency in Go, lightweight threads and the fundamental organisation unit in **Go**. Goroutines are more cost-effective than traditional threading and execute within the same address space, leveraging multiple dynamic OS threads.

In **Go**, Goroutines can be implemented as functions, anonymous functions, or methods, allowing them to execute concurrently alongside other code or tasks without requiring parallel execution. It is worth highlighting that even the `main` function in a **Go** program is considered a Goroutine.

One way to improve Goroutines' performance is to enhance them. It is crucial to understand their basic building blocks are briefly explained below:

- Channels
- Select statement
- Waitgroup
- Mutexes

**Channels** are synchronisation primitives that act as pipelines or transmission mediums for safe communication and data sharing between goroutines. There are two types of channels: buffered and unbuffered. Buffered channels have a predefined capacity to store a certain amount of data, while unbuffered channels have no capacity limit.

Refer to the code section below for insights on creating and using channels to communicate within your program.


```go
package main

import (
"fmt"
"time"
)


func main() {
// Declare a buffered channel of integers with a capacity of 2.
bufferedChan := make(chan int, 2)

// Unbuffered channel with no capacity
unbufferedChan := make(chan string)
```

In the main program, a variable is declared as a **Buffered** channel with a specific capacity, whereas the **Unbuffered** channels stated below do not have a set capacity.


```go

// Send two integers to the channel.
bufferedChan <- 10
bufferedChan <- 20

// Start a goroutine to receive data from the unbuffered channel.
go func() {
data := <-unbufferedChan
fmt.Println("Received from unbuffered channel:", data)
}()

// Wait for a short period before sending and receiving more data.

time.Sleep(time.Second)

```

Two values, 10 and 20, are sent to a **buffered** channel using the arrow syntax(`<-`). In the main program, an **unbuffered** channel utilises a goroutine to assign received value to a `data` variable and perform its process. A short delay introduces using the `time` package to ensure proper execution of the goroutine before the main function finishes.


```go

// Send some data to the unbuffered channel.
unbufferedChan <- "Hello, Gophers."

// Receive some data from the buffered channel.
data1 := <-bufferedChan
data2 := <-bufferedChan

// Print out the received data.
fmt.Println("Received from buffered channel:", data1, data2)

}

```

A `string` data is sent to the unbuffered channel, while the data sent to the buffered channels are received and stored in variables(`data1` & `data2`). If the buffered channel's capacity exceeds, it can cause an error and lead to a deadlock. Considering the channel when sending data. Buffered channels can store data but stay within their capacity. Unbuffered channels require immediate data reception.

Here's the complete code snippet below:


```go
package main

import "fmt"

func main() {
// Declare a buffered channel of integers with a capacity of 2.
bufferedChan := make(chan int, 2)

// Unbuffered channel with no capacity
unbufferedChan := make(chan string)
// Send two integers to the channel.
bufferedChan <- 10
bufferedChan <- 20

// Start a goroutine to receive data from the unbuffered channel.
go func() {
data := <-unbufferedChan
fmt.Println("Received from unbuffered channel:", data)
}()

// Wait for a short time before sending and receiving more data.
time.Sleep(time.Second)

// Send more data to the buffered channel.
bufferedChan <- 30

// Send some data to the unbuffered channel.
unbufferedChan <- "Hello, Gophers."

// Receive some data from the buffered channel.
data1 := <-bufferedChan
data2 := <-bufferedChan

// Print out the received data.
fmt.Println("Received from buffered channel:", data1, data2)

}

```

Here is the output of the code below:

![channels](https://i.imgur.com/Lla3u9E.png)

Go's **select** statement is similar to the `switch` statement. It is typically employed to execute tasks in the presence of channel data while concurrently reading from or writing to other goroutines. This feature enhances performance and minimises delays among goroutines during program execution.

A simple code explains how the `select` statement works.


```go
package main

import (
"fmt"
"time"
)

func main() {
// Declare an unbuffered channel of integer and string
intChan := make(chan int)
stringChan := make(chan string)

}
```

You create two-channel variables, one with the type "int" and the other with the type "string", both unbuffered.

---

```go
// Send an integer to the channel.
go func() {
intChan <- 299
}()

// Send a string to the channel.
go func() {
stringChan <- "Hello, Gophers."
}()

```

Goroutines are created and utilised to send data through their respective channels, allowing for simultaneous data transmission of different types (such as `int` and `string`).



```go
// Receive from the channels in a separate goroutine.
go func() {
select {
case i := <-intChan:
fmt.Println("Received an integer:", i)
case s := <-stringChan:
fmt.Println("Received a string:", s)
}
}()

// Wait for 1 second before ending the program.
time.Sleep(1 * time.Second)
```

An initialised goroutine receives data sent through the channels. It executes the `case` condition in the `select` statement block based on which goroutine receives the transmitted data first. The goroutine runs after a delay, just before the `main` program ends.

Below is the complete code snippet for the select statement.


```go
package main

import (
"fmt"
"time"
)

func main() {
// Declare an unbuffered channel of integers.
intChan := make(chan int)
stringChan := make(chan string)

// Send an integer to the channel.
go func() {
intChan <- 299
}()

// Send a string to the channel.
go func() {
stringChan <- "Hello, Gophers."
}()

// Receive from the channels in a separate goroutine.
go func() {
select {
case i := <-intChan:
fmt.Println("Received an integer:", i)
case s := <-stringChan:
fmt.Println("Received a string:", s)
}
}()

// Wait for 1 second before ending the program.
time.Sleep(1 * time.Second)
}
```

When you run the code above several times, the output you see should be similar to the image below:

![channels](https://i.imgur.com/eI2puXg.png)

This basic explanation is sufficient for your needs in this tutorial. For a deeper understanding of concurrency, Goroutines, and their building blocks, you can refer to this [resource](link for the article on Concurrency in Go @Earthly).

## Communicating Sequential Processes(CSP): The Model of Concurrency

Now that you have gained a basic understanding of concurrency and its practical implementation using Goroutines, it is essential to note that these concepts in Go are designed and built based on the theoretical model of [CSP](https://en.wikipedia.org/wiki/Communicating_sequential_processes).

![CSP](https://i.imgur.com/aZv1TX8.png)

CSP is a framework for describing patterns and understanding systems involving complex message exchanges. It helps explain how different components of a concurrent system interact. **Go's** approach to concurrency prioritises the principles of CSP and focuses on safety, efficiency, and ease of programming, building Goroutines upon them.

It's time to unleash the power of communication. Embrace CSP to conquer the complexities of concurrent systems and seamlessly orchestrate interactions. Up Next, You will learn how to relate concurrency and goroutines to the natural cycle of the earth and the universe at large.

## Explanation using the Universe: Planet Earth and Human

As you know, the universe comprises planets, galaxies, stars, moons, and asteroids. These celestial bodies coexist and engage in their natural processes concurrently in deep space.

Consider planet Earth as an example. It houses sophisticated systems with independent natural cycles. Weather changes, ocean currents, plant growth, ecosystem evolution, and human activities co-occur and influence one another.

Humans demonstrate concurrency. You have unique thoughts, feelings, and actions. Multitasking is natural for you, whether talking, reading this article, critical thinking or listening to music. Your brain adeptly manages these processes without becoming overwhelmed.

Concurrency proves beneficial in software programs and applications. Like humans, the software can benefit from executing multiple tasks simultaneously. This approach ensures a substantial increase in productivity, allows for the speedy completion of functions and promotes a well-rounded approach towards accomplishing your objectives.

You have covered the fundamentals of concurrency and goroutines sufficiently. It's time to prepare ourselves for the real challenges ahead.

## Benefits of Using Goroutines and Gmail Integration

Now that you are familiar with goroutines, you must understand the advantages and benefits of using them in conjunction with Gmail Integration for sending newsletter emails to users, which you will learn in this tutorial.

The benefits of utilising Goroutines and Gmail Integration include:

- Asynchronous operation and scalability.
- Efficient multitasking when sending emails.
- Fast email processing.
- Improved user and customer experience.

As you continue your journey, you will likely use goroutines to tackle various problem-solving tasks. In the following section, I will assist you in creating a necessary workspace.

# Setting Up the Environment

This section will teach you how to prepare your working environment for this tutorial, including creating a Gmail account and enabling third-party access.

## Prerequisite Knowledge

Before you start building things, ensure you have the following basic setup:

i. Make sure you have installed the Go compiler on your machine. The latest version available is Go v1.20.5. You can download it from [here](https://go.dev/dl/) based on your operating system if you still need to install it.

ii. Follow the installation instructions for your operating system (Mac, Linux, or Windows).

## Installing necessary tools

Now that you have set up your work environment, I will list all the necessary packages, tools, and libraries that you will use to build the application.

- Install the [go-mongodb-driver](https://www.mongodb.com/docs/drivers/go/current/), a tool designed to interact with a MongoDB cloud database. This project offers a range of tools to store user details and simplify the storage of emails intended for users.

- Install [gomail](https://github.com/go-gomail/gomail): a package that sends emails using an SMTP server to subscribed users for account notifications and communication.

- Install [chi](https://pkg.go.dev/github.com/go-chi/chi/v5): a package that builds HTTP services, initialises middleware, handles request paths, and serves static files.

- Install [godotenv](https://github.com/joho/godotenv): a package that loads environment variables from a `.env` file to access key variable values easily.

- Install [docconv](https://pkg.go.dev/code.sajari.com/docconv): package to converts uploaded `.docx` and `.doc` files to plain text for easier processing.

But before you install any of these packages, follow the instructions below:

1. Create a folder called "**mailapp**" or any name you like for the project.
2. Open the folder in your preferred IDE or text editor at the project's root level.
Use the terminal in the editor and run this command: `[go mod init github.com/username/project-folder-name]` in the given format 👇.

```go
go mod init github.com/akinbyte/mailapp
```

After initialising the `go mod init` command, the project folder will contain `go.mod` file.

![go-modules](https://i.imgur.com/X5qYNQv.png)

The `go.mod` file holds module details like name and version. Learn more about Go modules [here](https://go.dev/doc/modules/managing-dependencies).

Once you have completed that task, you can install the packages required for the application. Take it step by step.

To start, install the [MongoDB](https://www.mongodb.com/docs/drivers/go/current/) package for **Go** by executing the following command:

```go
go get go.mongodb.org/mongo-driver/mongo
```

Upon successful execution, you should observe a similar output in the terminal as shown below:

![mongo-install](https://i.imgur.com/DNkvNxJ.png)

The following package you need to install is the [Gomail](https://github.com/go-gomail/gomail) package. The installation process for this package is similar to the previous one. To install it, use the following command:

```go
go get gopkg.in/gomail.v2
```

After executing this command, you should see a similar output in your terminal, as shown below:

![go-mail-install](https://i.imgur.com/jqGQoYW.png)

Use the same approach to install [chi](https://pkg.go.dev/github.com/go-chi/chi/v5) for routing. Use the command below:

```go
go get github.com/go-chi/chi/v5
```

![go-chi-install](https://i.imgur.com/hp2a48a.png)

Follow the same procedure to install [godotenv](https://github.com/joho/godotenv) and [docconv](https://pkg.go.dev/code.sajari.com/docconv) using the commands below one at a time, respectively.

```go
go get github.com/joho/godotenv

go get -u code.sajari.com/docconv
```

![godotenv-install](https://i.imgur.com/InpfnYb.png)
![docconv-install](https://i.imgur.com/tktdrHV.png)

After installing all the libraries and their dependencies, you will notice that your project folder automatically creates a `go.sum` file. This file stores dependency checksums and versions, vital for managing dependencies in Go.

Create a `.env` file to store all the necessary environment variables to finalise the project setup. Once done, proceed to the next section, where you'll create a Gmail account and enable API access for email.

## Creating a Gmail account and enabling API access

I'm sure you already have a Gmail account, but you have two options for this project. You can either create a new account and connect it to this project or use your existing one.

To save time, let's assume you've already set up your Gmail account. Now, you need to configure it by enabling 2-Step Verification and allowing access to less secure apps, which is crucial.

When you create a Gmail account, by default, access by third parties is disabled. You'll need to follow specific procedures and configurations to enable access to other third-party apps, like the one you're going to build.
It is essential to configure and set up your account correctly.

i. Sign in to your Gmail account, go to the top right corner and click on your profile. Then, select "**Manage your Google Account**".
![manage-gmail](https://i.imgur.com/DfKUr6N.png)

ii. Once you've accessed your account page, you can modify settings and make configuration changes. Follow these steps:

- On the left menu, click on the "Security" section.
![account-page](https://i.imgur.com/06J03Wi.png)

- Scroll down to the bottom and click "**Less secure app**." You can find additional information about less secure apps if needed.
![less-secure-app](https://i.imgur.com/whVgwM6.png)
![less-secure-app-info](https://i.imgur.com/PLsMCSo.png)
![less-secure-app-info-1](https://i.imgur.com/9oKJLAQ.png)
- To enable access for third-party applications like this project, set up **2-Step Verification** in the **Security** section. Make sure to link your devices for verification and choose the appropriate prompts.
![verify-code](https://i.imgur.com/2gfDBID.png)
- Turn on 2-Step Verification by clicking on the corresponding option.
![2-step-turn-on](https://i.imgur.com/Nuc7F3y.png)
![turn-on](https://i.imgur.com/dOkw5MT.png)

iii. Once you have enabled the **2-Step Verification**, proceed to add an **App Password**.
![choose-create-app-password](https://i.imgur.com/3MP3q0U.png)

- Choose an app name and device from the dropdown menu, or create a custom name as shown in the image below:
![create-app-password](https://i.imgur.com/sfIOfUz.png)
- Click on the "Generate" button to create the **App password**.
![generate-app-password](https://i.imgur.com/jqAzM8U.png)
- Remember to write down or copy the generated App password for future use, as it will only be accessible once generated.

Once you have generated an App password, you can enable API access from other applications. In the next section, you will learn about using environment variables in the application.

## Configuring credentials and environment variables

Grant access to less secure apps using the generated App password to secure your Gmail account's integration into the project. To minimise vulnerabilities, storing sensitive credentials, including the generated password and other relevant details, in the `.env` file is recommended. You can implement enhanced security measures by referencing these credentials in the codebase instead of using them directly.

This approach will make it easy to configure and debug your codebase, enhance security, and prevent errors or mistakes during development. Below are the key-value pairs of the environment variable credentials stored in the `.env` file for this project:

```env
GMAIL_ACC=Yusufakinleye144@gmail.com

APP_PASSWORD=iqcepxuzsdumzslx

USER_NAME=bravebyte
```

As you progress, you can continue modifying the .env file by adding new key-value variables required for your application.

# Building the Mail Server API

This section teaches you to build a solid and reliable server-side mail application. But before you proceed, in the next section, I'll explain the application's structure for better comprehension.

## Designing the API architecture

The application structure you will be working with here relies on a diagrammatic representation below, illustrating the application's architecture.

![mail-app-arch](https://i.imgur.com/pXAUWGi.png)

Soon, you'll grasp the application's structure implemented using Package Oriented Design. Let's begin.

To structure the application for this project, you'll create the following directories (packages) and files. You'll learn about their features and usage.

- _Main and routes_: These are the leading root files of the project that use the main package for compilation.

- _db_: The _db_ package manages interactions with the **MongoDB** database, including reading and writing queries.

- _handlers_: The _handlers_ package contains methods to process user requests.

- _tools_: The _tools_ package consists of predefined reusable functions that facilitate user requests.

- _email_: The _email_ package is the project's main focus. It includes a function that receives emails through channels using goroutines and another function that sets up the **SMTP** server to send emails.

- _model_: The _model_ package contains structs that hold the payload or details received while processing requests.

It's important to note that the output method of the **query** from the **db** package to the **handlers** package implements an `interface`, which is the same as the one used by the **handlers** package for communication with the **main** package.

Now that you know the basics of each package's functionality, let's integrate the MongoDB database to store user information and send emails.

Before you proceed, here's a quick explanation of the application's `model`.

## Understanding the Application Model

The `model` package includes the `Subscriber struct`, which holds various fields of type `string` to store the details and `struct` tags of the **subscribers** when they submit their information.

```go
package model

import "time"
// Subscriber: information or details from subscribers
type Subscriber struct {
 FirstName string `bson:"first_name" json:"first_name"`
 LastName  string `bson:"last_name" json:"last_name"`
 Email     string `bson:"email" json:"email"`
 Interest  string `bson:"interest" json:"interest"`
}
```


The `MailUpload struct` stores the document uploaded by the user along with their corresponding types and `struct` tags.

```go
// MailUpload - holds the uploaded content and details for the mail
type MailUpload struct {
 DocxName    string    `bson:"docx_name" json:"docx_name"`
 DocxContent string    `bson:"docx" json:"docx"`
 Date        time.Time `bson:"date" json:"date"`
}
```

---

The `Mail struct` will be filled with information and seamlessly routed to subscribers via channels.

```go
// Mail: contains the field of what the mail entails
type Mail struct {
 Source      string
 Destination string
 Message     string
 Subject     string
 Name        string
}
```

---

### Code Snippet for the model Package

```go
package model

import "time"

// Subscriber: information or details from subscribers
type Subscriber struct {
 FirstName string `bson:"first_name" json:"first_name"`
 LastName  string `bson:"last_name" json:"last_name"`
 Email     string `bson:"email" json:"email"`
 Interest  string `bson:"interest" json:"interest"`
}

// Mail: contains the field of what the mail entails
type Mail struct {
 Source      string
 Destination string
 Message     string
 Subject     string
 Name        string
}

// MailUpload - holds the uploaded content and details for the mail
type MailUpload struct {
 DocxName    string    `bson:"docx_name" json:"docx_name"`
 DocxContent string    `bson:"docx" json:"docx"`
 Date        time.Time `bson:"date" json:"date"`
}

```

---

Once you understand this model for this application, you can begin implementing and setting up the MongoDB database.

## Integrating MongoDB Database

To integrate the MongoDB database for this project using MongoDB Atlas, a user-friendly cloud database.
Check this [**link**](https://www.mongodb.com/docs/atlas/getting-started/) for setup instructions in the **Get Started with Atlas** section, and remember to configure for easy access by changing the IP Address.

Follow the steps below to set up your database and get the connection string:

- Click on **Database** in the left menu and select **Browse Collection** to create a database for the application.
![atlas-dashboard](https://i.imgur.com/m7pZDFJ.png)

- Click **Create Database** to add a new database named **mail-app** and include two collections:
**mails** (stores sent mail) and **subscribers** (stores registered subscribers).
![create-db](https://i.imgur.com/rPVJCu7.png)

- Next, click "Connect" to choose a connection method.
![connect-method](https://i.imgur.com/Tayy1Km.png)

Select the required MongoDB Driver for **Go** and copy the connection string. Update it with your username and password.
`mongodb+srv://<username>:<password>@cluster0.opv1wfb.mongodb.net/?retryWrites=true&w=majority`
![connect-string](https://i.imgur.com/gUuJQBk.png)

Update the .env file with the URI connection string:


```env
APP_PASSWORD=iqcepxuzsdumzslx

GMAIL_ACC=Yusufakinleye144@gmail.com

USER_NAME=bravebyte

URI=mongodb+srv://<username>:<password>@cluster0.opv1wfb.mongodb.net/?retryWrites=true&w=majority

```

Now you can connect to the **mail-app** database using the provided `URI` and have it integrated into your code.

### Setting Up the Database Connection

You will learn how to connect to the **mail-app** database using the' URI' connection string. Let's begin!

Create these **Go** files: `db.go`, `query.go`, `collection.go`, and `service.go` in the _db_ package. In this section and the subsequent one, I will explain the implementation process in each file.

In `db.go`, import the necessary packages and set up the cloud database connection with the `SetConnect` function.

```go

package db

import (
"context"
"log"
"os"
"time"

"go.mongodb.org/mongo-driver/mongo"
"go.mongodb.org/mongo-driver/mongo/options"
)

func SetConnect(uri string) (*mongo.Client, error) {}
```

---

The `SetConnect` function sets a `context` timeout with a deadline and cancellation signal, ensuring the database connection process finishes within a specified time limit.

```go
dbCtx, dbCancelCtx := context.WithTimeout(context.Background(), 100*time.Second)
defer dbCancelCtx()
```

Connect the database using the `dbCtx` context and the `URI` string. Check for connection errors.

```go

client, err := mongo.Connect(dbCtx, options.Client().ApplyURI(uri))
if err != nil {
log.Panicln("Error while connecting to database: ", err)
}
```

Pinging the database verifies the connection, and the code returns the database client.

```go
if err := client.Ping(dbCtx, nil); err != nil {
log.Fatalln("cannot ping the database: ", err)
}

return client, nil
```

In the `db.go` file, the `OpenConnect` function keeps the database connection open to avoid program failure. It returns a `*mongo.Client` pointer, initialises `count` to zero, and sets up an infinite `for` loop.

```go
func OpenConnect() *mongo.Client {
 uri := os.Getenv("URI")
 count := 0
 log.Println("....... Setting up Connection to MongoDB .......")
 for{}
}
```

The infinite loop invokes `SetConnect` with the environment variable key `URI` using the `os` package. If there's a connection error, it logs the message for database disconnection and increments `count`. Otherwise, it confirms the database connection and returns the client.

```go
client, err := SetConnect(uri)
if err != nil {
 log.Println("Mail App Database not Connected")
 count++
 } else {
  log.Println("Mail App Database is Connected")
  return client
 }
```

In the loop, the code checks the value of `count`. If it exceeds the limit, it logs the error and returns nothing. If the `count` is less than 5, it logs retry information, pauses the program, and continues the connection process.

```go
if count >= 5 {
 log.Println(err)
 return nil
}

log.Println("Wait:.... Mail App Database Retrying to Connect ....")
time.Sleep(10 * time.Second)
continue
```

Please note that the `main` function of the program will invoke the `OpenConnect` function to retrieve the `*mongo.Client`. Afterwards, you will learn to access and work with the **mail-app** database and its collections.

### Default Collection Function

The `collection.go` file contains a `Default` function. It takes a `client` parameter of type `*mongo.Client` and a `collectionName` parameter of type `string`. This function returns a pointer to `*mongo.Collection` for accessing the specific **mail-app** database.

```go
package db

import "go.mongodb.org/mongo-driver/mongo"

func Default(client *mongo.Client, collectionName string)*mongo.Collection{
 return client.Database("mail-app").Collection(collectionName)
}
```

To interact with and access the **mail-app** database collections (mails & subscribers) by using this function to write queries for the application.

### Database Queries

You will incorporate the necessary query methods into this section's `query.go` file. These methods must effectively communicate with the **mail-app** database and its collections utilising the `Mongo struct`.

Import necessary packages and create a `Mongo struct` with a `MailDB` field (a `*mongo.Client` pointer type).

Lastly, a `NewMongo` constructor function is defined. It takes a `client` database pointer as a parameter and returns a `DataStore` interface for implementing the queries.

```go
package db

import (
"context"
"fmt"
"log"
"time"

"github.com/akinbyte/mailapp/model"
"go.mongodb.org/mongo-driver/bson"
"go.mongodb.org/mongo-driver/bson/primitive"
"go.mongodb.org/mongo-driver/mongo"
)

type Mongo struct {
MailDB *mongo.Client
}

func NewMongo(client *mongo.Client) DataStore {
return &Mongo{MailDB: client}
}
```

---

The application utilises three query methods to interact with the **mail-app** database:

- `AddSubscriber`: Adds and stores new subscriber details.
- `AddMail`: Stores the sent mail.
- `FindSubscribers`: Retrieves all registered subscribers.
  
---

To implement this, begin with the `AddSubscriber` query. This method requires the `subs` argument of type `model.Subscriber` and returns `bool`, `string`, and `error`.

```go
func (mg *Mongo) AddSubscriber(subs model.Subscriber) (bool, string, error) {}
```

---

This method sets a timeout in the `context` for query processing. The subscriber's `email` is used to check for their existence in the `subscribers` collection of the database while also verifying for errors.

```go
ctx, cancelCtx := context.WithTimeout(context.Background(), 10*time.Second)
defer cancelCtx()

var res bson.M
filter := bson.D{{Key: "email", Value: subs.Email}}
err := Default(mg.MailDB, "subscribers").FindOne(ctx, filter).Decode(&res)
```

---

If the error `mongo.ErrNoDocuments` occurs during error checking, indicating no matching document with the email filter, the subscriber details are added to the `subscribers` collection using the `insertOne` method of `mongo.Collection` type with the previously defined `Default` function.

Errors are checked and returned with relevant information. New subscribers are added to the `subscribers` collection without errors, while the function returns that the subscriber is registered already for existing subscribers.

```go
if err != nil {
 if err == mongo.ErrNoDocuments {
  _, err := Default(mg.MailDB, "subscribers").InsertOne(ctx, subs)
  if err != nil {
   return false, "", fmt.Errorf("AddSubscriber: cannot registered this account : %v", err)
  }
 return true, fmt.Sprintf("New Subscriber Added"), nil
 }
 log.Fatalln("AddSubscriber: cannot query database", err.Error())
}
return true, "", nil
```

---

The `AddMail` method stores mail details and content. It takes a parameter `mu` of type `model.MailUpload` and returns a `string` and `error`. This method facilitates efficient mail storage.

```go
func (mg *Mongo) AddMail(mu model.MailUpload) (string, error) {}
```

---

As usual, set the `context` timeout. The `mailUpload struct` saves the sent mail and its details in the `mails` collection of the `mail-app` database. When the mail details are successfully added, it returns a `nil` error. However, it displays returns an error message if there are any issues.

```go
ctx, cancelCtx := context.WithTimeout(context.Background(), 10*time.Second)
defer cancelCtx()
//All the uploaded mail to the database
_, err := Default(mg.MailDB, "mails").InsertOne(ctx, mu)
if err != nil {
 return "", fmt.Errorf("AddMail: unable to add new mail, %v", err)
}
return "New mail successfully added", nil
```

---

Finally, the `FindSubscribers` method retrieves subscriber details without any parameters. It returns a `slice` of type `Primitive.M`, an alias for `map[string]interface{}`, along with a possible `error`.

```go
func (mg *Mongo) FindSubscribers() ([]primitive.M, error) {}
```

---

A `context` timeout is added to process the query within a specified time. The `Find()` method queries the `subscribers` collection without any filter argument passed to it to obtain a `Cursor` that encompasses all the documents in the collection. The result of all the documents is decoded or wrapped in a variable called `res` of type `[]bson.M`. Finally, the `Cursor` is closed using the `defer Cursor.Close(ctx)`. Handled the errors and returned the outcomes of all subscriber's documents.

```go
ctx, cancelCtx := context.WithTimeout(context.Background(), 5*time.Second)
defer cancelCtx()

var res []bson.M
//Get all subscriber's data
cursor, err := Default(mg.MailDB, "subscribers").Find(ctx, bson.D{{}})
if err != nil {
 return []bson.M{}, err
}
//decode the result in `res`
if err = cursor.All(ctx, &res); err != nil {
 return []bson.M{}, fmt.Errorf("FindMail: Cannot get all mail: %v", err)
}
defer cursor.Close(ctx)

if err = cursor.Err(); err != nil {
 return []bson.M{}, fmt.Errorf("FindMail: Cursor Error : %v", err)
}
return res, nil
```

---

After completing this step, you've implemented all the required queries for this application. Next, You should have these methods added to an interface.

### Interface Implementation

In `service.go`, create the `DataStore interface` for a secure application, enhanced accessibility, and vulnerability prevention. Implement the interface by adding the queries method.

```go
package db

import (
"github.com/akinbyte/mailapp/model"
"go.mongodb.org/mongo-driver/bson/primitive"
)

type DataStore interface {
 AddSubscriber(subs model.Subscriber) (bool, string, error)
 AddMail(mu model.MailUpload) (string, error)
 FindSubscribers() ([]primitive.M, error)
}

```

---

You have completed all the tasks to implement in the `db` package. Now, let's create the API endpoints and learn how to handle requests.

## Creating API endpoints and handling requests

Now that you understand integrating the database, it's time to create the API endpoints for handling user requests. Before proceeding, examining the contents in the `tools` package is crucial, as the functions defined within it play a critical role in handling and processing each request.

The following explains the functionalities of the functions within the package:

### Implementing Functions in the Tools Package

First, import all the packages and built-ins needed here.

```go
package tools

import (
 "bufio"
 "encoding/json"
 "fmt"
 "html/template"
 "log"
 "net/http"
 "path/filepath"
 "strings"
 "time"

 "code.sajari.com/docconv"
 "github.com/akinbyte/mailapp/model"
)
```

---

The `ReadForm` function below ensures that the subscriber's request body is read safely and efficiently using `http.ParseForm()` to get the request body parameters. If any error comes up while parsing the form, the function prints out the error and returns an empty struct value with the error.

The subs of struct type of `model.Subscribers` fields are populated using the `rq.Form.Get("key")` to get the value associated with the given keys.

If successful, The function returns the filled `subs` struct and a `nil` error message.

```go
// ReadForm: This is a reusable function to help read the information or details submitted.
func ReadForm(rq *http.Request, subs model.Subscriber) (model.Subscriber, error) {
 if err := rq.ParseForm; err != nil {
 log.Println(err)
 return model.Subscriber{}, err }
 subs = model.Subscriber{
  FirstName: rq.Form.Get("first_name"),
  LastName:  rq.Form.Get("last_name"),
  Email:     rq.Form.Get("email"),
  Interest:  rq.Form.Get("interest"),
 }
 
 return subs, nil
}
```

---

It would help to use the `JSONWriter` function to write to the server. This function constructs an HTTP response using the parameters `wr` and `interface`. It also takes `msg` as the message to be sent with the response, `statusCode` as the response's status code, and returns an error.

This function sets the HTTP response header to JSON format and adds the status code. It encodes the `msg` value into JSON using `json.Marshal`. Finally, it writes the encoded data to the output response using the `Write` method of `http.ResponseWriter`. Throughout the process, it checks for errors.

```go
// JSONWriter: this will help send json response to the client page of this application
func JSONWriter(wr http.ResponseWriter, msg string, statusCode int) error {
wr.Header().Set("Content-Type", "application/json")
wr.WriteHeader(statusCode)

b, err := json.Marshal(msg)
if err != nil {
 return err
}
_, err = wr.Write(b)
if err != nil {
 return err
}
return nil
}
```

---

The `ReadMultiForm` function reads and processes a multipart form submitted via an HTTP request. It accepts three parameters: `wr http.ResponseWriter`, `rq *http.Request`, and `mail model.MailUpload`. The function returns a `model.MailUpload` object and an error, if any, for further processing.

```go
func ReadMultiForm(wr http.ResponseWriter, rq *http.Request, mail model.MailUpload) (model.MailUpload, error) {}
```

---

The function starts by parsing the multipart form data from the HTTP request using the `rq.ParseMultipartForm(10 << 20)` method, which can handle up to 10 megabytes (10MB) of data to ensure correct extraction of the form data. If there are any issues during the parsing process, the code logs the error and terminates the program.

```go
if err := rq.ParseMultipartForm(10 << 20); err != nil {
 log.Fatal(err)
}
```

---

The function parses the form data and retrieves the `MultipartForm` from the request using the `rq.MultipartForm` property to access the form's fields and files.

```go
form := rq.MultipartForm
```

---

Next, the function assigns the form's `docx_name` field value to the `mail.DocxName` variable, capturing the uploaded document's name from the submitted form.

```go
mail.DocxName = form.Value["docx_name"][0]
```

---

The function sets the `mail.Date` field to the current time using `time.Now()` to record the timestamp of the form submission.

```go
mail.Date = time.Now()
```

---

Using `form.File["docx"]`, the function checks for the existence of the `docx` file within the form, and if it cannot find the file, it returns an error indicating the inability to retrieve the uploaded document.

```go
file, ok := form.File["docx"]
if !ok {
 return model.MailUpload{}, fmt.Errorf("unable to get uploaded document")
}
```

---

If the file is present, the function verifies the file's extension using `filepath.Ext(file[0].Filename)`. The valid file extension is to access to read using `file[0].Open()`. If any error occurs during the opening, it returns an error indicating the inability to open the document, and the `defer` function closes up the file once it finishes reading.

```go
if file[0].Filename != "" {
fileExtension := filepath.Ext(file[0].Filename)

f, err := file[0].Open()
if err != nil {
 return model.MailUpload{}, fmt.Errorf("unable to open uploaded document")
}
defer f.Close()
}
```

---

The `switch-case` statement checks if the uploaded file extension matches the allowed extensions like ".doc", "docx", or ".txt".

```go
switch fileExtension {}
```

---

For `case ".txt"`, the function utilises the `bufio.Scanner` function to process the content of the opened file. It reads the file line by line, formats each string with an HTML line break (`<br>`), and appends it to the `mail.DocxContent` variable.

```go
scanner := bufio.NewScanner(f)

for scanner.Scan() {
line := fmt.Sprintf("%s<br>", scanner.Text())
mail.DocxContent += line
}

if err := scanner.Err(); err != nil {
 log.Fatal(err)
}
```

The code reads all the lines and checks for scanning errors using `scanner.Err()`. Logs and stops the program if an error occurs.

---

While for that of `case ".docx", ".doc"`, the function utilises a `docconv.ConvertDocx` function to convert an MS Word docx or doc file to text for easy readability. It reads the file line by line, formats each line with an HTML line break (`<br>`), appends it to the `content` variable, and then assigns its value to `mail.DocxContent`.

```go
// process .docx or .doc uploaded files
case ".docx", ".doc":
res, _, err := docconv.ConvertDocx(f)
if err != nil {
 log.Fatal(err)
}

lines := strings.Split(res, "\n")
var content string
// Add line breaks to each line
for _, line := range lines {
 content += line + "<br>"
}
mail.DocxContent = content
```

---

If the file extension is not allowed, the function returns an error message stating that only specific file extensions (like ".doc", "docx", or ".txt") are permitted using the `default` condition.

```go
default:
return model.MailUpload{}, fmt.Errorf("upload document not allow; try .txt .docx or .doc")
```

---

Finally, the function returns the `mail` object containing the extracted form information. It also bears a `nil` error value to indicate the successful process.

```go
return mail, nil
```

---

The `HTMLRender` function takes three parameters: `wr http.ResponseWriter`, `rq http.Request`, and `dt any` (alias for `interface`). It parses an HTML template file using `template.ParseFiles` and executes it with `wr` and `dt` using the `tmp.Execute` method. If an error occurs while parsing or during execution, the function returns an error message, and a `nil` value indicates successful rendering.

```go
func HTMLRender(wr http.ResponseWriter, rq *http.Request, dt any) error {
filePath := "./index.html"

tmp, err := template.ParseFiles(filePath)
if err != nil {
 return fmt.Errorf("HTMLRender Error: failed to parse file: %v", err)
}

err = tmp.Execute(wr, dt)
if err != nil {
 return fmt.Errorf("HTMLRender Error: failed to execute template: %v", err)
}

return nil
}
```

Now that you've understood the functionality implemented in the _tools_ package let's create the methods for handling HTTP requests.

### Creating Handler Method to Process HTTP

To handle requests and create endpoints, utilise the _handlers_ package. Implement the handler methods in `handlers.go` and define an `interface` in `service.go` that implements all the handler methods for HTTP requests; used a similar approach was in the _db_ package.

---

In `handlers.go`, import all the necessary packages.

```go
package handlers

import (
"fmt"
"log"
"net/http"
"time"
"os"

"github.com/akinbyte/mailapp/db"
"github.com/akinbyte/mailapp/model"
"github.com/akinbyte/mailapp/tools"

"go.mongodb.org/mongo-driver/mongo"
)
```

---

The `MailApp struct` contains two fields: `MailDB` of type `db.DataStore`, an interface implementing query methods in the _db_ package, and `MailChan` of type `chan model.Mail` channel to hold mail details to be sent to subscribers.

```go
type MailApp struct {
 MailDB db.DataStore
 MailChan chan model.Mail
}
```

---

The `NewMailApp` function accepts `client *mongo.Client` and `mailchan chan model.Mail` as parameters and returns a `Logic` interface implementing all the handler methods. It serves as a constructor function, enabling modularisation, dependency injection, and asynchronous communication for building the application.

```go
func NewMailApp(client *mongo.Client, mailchan chan model.Mail) Logic {
 return &MailApp{
  MailDB: db.NewMongo(client),
  MailChan: mailchan,
 }
}
```

---

The `MailApp struct` has a `Home` method that returns an anonymous function, `http.HandlerFunc`. This function takes `wr http.ResponseWriter` and `rq *http.Request` as parameters. It generates an HTML response for the route using `tools.HTMLRender()` from the _tools_ package. While rendering, any error is logged.

```go
func (ma *MailApp) Home() http.HandlerFunc {
 return func(wr http.ResponseWriter, rq *http.Request) {
  err := tools.HTMLRender(wr, rq, nil)
  if err != nil {
   log.Println(err)
   return
  }
 }
}
```

---

The following implemented method is `GetSubscriber`.

```go
func (ma *MailApp) GetSubscriber() http.HandlerFunc {
 return func(wr http.ResponseWriter, rq *http.Request) {}
}
```

---

Declare a variable named `subs` of type `model.Subscriber`.
Call the `tools.ReadForm` function from the _tools_ package to read the HTTP request by passing `wr`, `rq`, and `subs` as arguments. Respond to an error with `http.Error`, including the specified status code and error message.

```go
var subs model.Subscriber
subscriber, err := tools.ReadForm(rq, subs)
if err != nil {
 http.Error(wr, fmt.Sprintf("failed to read json : ",err), http.StatusBadRequest)
 return
}
```

---

The `AddSubscriber` method injects dependency through the `db.DataStore` interface. It receives the result of the `ReadForm` function as an argument. The `AddSubscriber` function returns three outputs: `ok` (to verify the successful addition of the subscriber), "msg" (a message to include in the HTTP response), and `err` (to indicate any errors encountered).

```go
ok, msg, err := ma.MailDB.AddSubscriber(subscriber)
if err != nil {
 http.Error(wr, msg, http.StatusInternalServerError)
 return
}
```

---

The `switch-case` statement utilises the `tools.JSONWriter` to generate an **HTTP** response by providing reasonable arguments for verifying subscriber registration.

```go
switch ok {
 case msg == "":
 tools.JSONWriter(wr, "You have already registered", http.StatusOK)
 case msg != "":
 tools.JSONWriter(wr, msg, http.StatusOK)
}
```

---

Lastly, the `SendMail` method is the primary handler for processing the HTTP request to send mail to all registered subscribers. I will explain the implementation of this handler step by step.

```go
func (ma *MailApp) SendMail() http.HandlerFunc {
 return func(wr http.ResponseWriter, rq *http.Request) {}
}
```

---

To retrieve input details and the uploaded document containing the mail to send, the `tools.ReadMultiForm` function accepts a `mailUpload` variable of type `model.MailUpload` as an argument, along with `wr` and `rq`. `http.Error` to obtain the corresponding HTTP response message and `http.StatusBadRequest` (400) status code.

```go
var mailUpload model.MailUpload
upload, err := tools.ReadMultiForm(wr, rq, mailUpload)
if err != nil {
 http.Error(wr, err.Error(), http.StatusBadRequest)
 return
}
```

---

The `tools.ReadMultiForm` function returns a value of type `model.MailUpload` and is then passed as an argument to the `AddMail` method, which is part of the `db.DataStore` interface. The purpose is to store the mail intended for subscribers in the database. Following this, an HTTP response with a status code of `http.StatusInternalServerError` (500), accompanied by an appropriate message.

```go
msg, err := ma.MailDB.AddMail(upload)
if err != nil {
http.Error(wr, msg, http.StatusInternalServerError)
return
}
```

---

The program logs information, notifies the following process and pauses for a few milliseconds using the `time` package.

```go
log.Println(msg)
log.Println("........ preparing to send mail to subscribers ........ ")
time.Sleep(time.Millisecond)
log.Println("........ Accessing the subscribers Database ........ ")
```

---

The next process calls the `FindSubscribers` method, which returns all registered subscribers in the form of a `[]primitive.M` slice of `map[string]interface{}` named `res`. The process also checks for any return error.

```go
res, err := ma.MailDB.FindSubscribers()
if err != nil {
 http.Error(wr, fmt.Sprintf("failed query: %v", err), http.StatusInternalServerError)
 return
}
```

---

Iterate through `res` to access each `map[string]interface{}` containing the database document of registered subscribers.

Assign subscriber details to variables `subEmail`, `firstName`, and `lastName` by extracting them using their respective keys. Fill in the fields of `model.Mail struct` with their corresponding values and assign it to the `mail` variable. Send `mail` through the `MailChan` channel field in the `MailApp struct` to a receiving Goroutine.

```go
for _, s := range res {
 subEmail := s["email"].(string)
 firstName := s["first_name"].(string)
 lastName := s["last_name"].(string)
 
 subName := fmt.Sprintf("%s %s", firstName, lastName)
 mail := model.Mail{
  Source: os.Getenv("GMAIL_ACC"),
  Destination: subEmail,
  Name: subName,
  Message: upload.DocxContent,
  Subject: upload.DocxName,
 }
 ma.MailChan <- mail
}
```

---

After successfully sending the mail to the subscriber through an efficient performance channel for processing, the `tool.JSONWriter` generates an HTTP response to the server with the status code `http.StatusOK` (200), confirming the successful delivery of the mail.

```go
err = tools.JSONWriter(wr, fmt.Sprintf("Mail Sent %v subscribers", len(res)), http.StatusOK)
if err != nil {
 http.Error(wr, err.Error(), http.StatusInternalServerError)
 return
}
```

---

Implement an interface after creating the handler methods to process the HTTP requests.

### Implementing an Interface for Handler Methods

Open the `service.go` file and implement the `Logic` interface to add all the handler methods.

---

```go
package handlers

import "net/http"

type Logic interface {
 Home() http.HandlerFunc
 GetSubscriber() http.HandlerFunc
 SendMail() http.HandlerFunc
}
```

---

Next, let's create endpoints for these handlers.

### Creating Routes Endpoints(URL)

You can now create endpoints(URL) which typically correspond to a specific function or method in the backend code that handles the request and returns the appropriate response.

The `Routes` function takes in the parameter of `lg handlers.Logic` to access all the handler methods in the `Logic interface` and returns `*chi.Mux` from the `chi` package.

A new HTTP multiplexer `mux` is initialised, which implements the `Router interface` of the `chi` package, then the `Logger` function from the `chi` middleware logs the start and end of each request call and then recovers from panics logs the panic (and a backtrace) using the `Recoverer` function.

```go
package main
import (
 "github.com/go-chi/chi/v5"
 "github.com/go-chi/chi/v5/middleware"
 "github.com/yusuf/mailapp/handlers"
)
func Routes(lg handlers.Logic) *chi.Mux {
 mux := chi.NewRouter()
 mux.Use(middleware.Logger)
 mux.Use(middleware.Recoverer)
}
```

---

Then, the HTTP method `GET` execute the `lg.Home` HandlerFunc attached with the URL `/` route pattern to render the homepage.

Also, there are the `POST` methods with the route pattern of `/api/submit` and `/api/send` to carry out the request for the Subscriber to register and for the user to send mail to their subscribers, respectively.

```go
mux.Get("/", lg.Home())
mux.Post("/api/submit", lg.GetSubscriber())
mux.Post("/api/send", lg.SendMail())
```

---

The server should serve and handle the static files (CSS and images) required by the application's client side. 

To achieve that, `http.Dir` allows access to the **static** directory path `./static` at the root level. This directory is passed as an argument to the `http.FileServer` function, which returns an `http.Handler` assigned to the `fileServer` variable. This `fileServer` serves HTTP requests with the contents of the directory.

```go
fileServer := http.FileServer(http.Dir("./static"))
```

---

The `Handle` method is invoked to execute the `fileServer` `http.Handler` by stripping off the prefix `/static` of the pattern `/static/*` with the `http.StripPrefix` function. `mux` Pointer to `chi.Mux` is returned to be passed as an argument to the `http.Server`.

```go
mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
return mux
```

---

The application router to execute all the HTTP requests is now correctly done. It would help if you looked into the next section that shows how the mail server is setup up to receive data from a channel.

## Integrating Gmail API for Sending and Receiving Emails

Using Goroutines to achieve concurrency in sending mail to various people all at once to achieve utmost performance and efficiency is the target goal of this application.

You will learn how to integrate **Google Gmail API** in this program in sending mail, utilising the advantage of goroutine using channels in retrieving data. To get this done, let's get to it.

---

Create an `email.go` file in the _email_ package. If done, first import the packages needed to implement the functionality of this package.

```go
package email

import (
 "log"
 "os"

 "github.com/akinbyte/mailapp/model"
 "gopkg.in/gomail.v2"
)
```

---

Create a `MailServer` function with a parameter of the `mailChan model.Mail` is the receiving channel object when the user requests to send mail.

```go
func MailServer(mailChan model.Mail) {}
```

---

The `gomail` package initialises a new SMTP Dialer to connect to the SMTP Server. It uses `smtp.gmail.com` as the host and 465 as the port and passes the credentials stored in the `.env` file as arguments to the `gomail.NewDialer` function.

The dial function returns a value that authenticates the user's access to the SMTP Server. It assigns this value to variables `s` and `err` for further processing. Next, it performs error checking, and if any errors are detected, the program panics.

```go
d := gomail.NewDialer("smtp.gmail.com", 465, os.Getenv("GMAIL_ACC"), os.Getenv("APP_PASSWORD"))
s, err := d.Dial()
if err != nil {
 log.Panicf("Error connecting to the Mail Server: ", err)
}
```

---

After the `Dial` function opens a connection to the server, a newly constructed message is now assigned to `msg` utilising the available methods such as `SetHeader`, `SetBody` and `SetAddressHeader` that `gomail.Message struct` implements

```go
msg := gomail.NewMessage()
msg.SetAddressHeader("From", mailChan.Source, os.Getenv("USER_NAME"))
msg.SetHeader("To", mailChan.Destination)
msg.SetHeader("Subject", mailChan.Subject)
msg.SetBody("text/html", mailChan.Message)
```

---

After connecting to the SMTP server and composing the message, the user's authentication access `s` and the composed `msg` are passed as arguments to the `Send` function, delivering the message to the subscriber. The program logs any errors that occur during the sending of the mail.
The `Reset` method maintains the message settings for future messages.

```go
if err := gomail.Send(s, msg); err != nil {
 log.Printf("Mail Sever : %s %v\n", mailChan.Destination, err)
}
msg.Reset()
```

---

The mail server, which helps send the mail to the respective subscriber, is now set up. Let's move on to the function where the mail server will receive the sent mail through a buffered channel.

## Implementing Goroutines and Channels for Concurrency

This application aims to utilise the power of concurrency using goroutines to deliver mail messages to multiple people to enhance efficiency and performance seamlessly.

There are various approaches to achieving synchronisations, using the `sync` or the `atomic` package to accomplish this application's goal: handle mail delivery by processing the mail objects received through the channel.

You will use the custom implementation with the Buffered channel, and I will provide a detailed explanation below.

---

The function `MailDelivery` uses two parameters, `mailChan` and `worker`, to receive a channel object and determine the number of goroutines (gophers) to generate.

```go
func MailDelivery(mailChan chan model.Mail, worker int) {}
```

---

The function starts by creating a buffered channel called `completionChan` of a boolean type with a capacity equal to the value of the `worker` parameter. This channel enables non-blocking data transmission when the buffer is not complete.

```go
completionChan := make(chan bool, worker)
```

---

A `for` loop generates goroutines (gophers) within the range of the specified number of `workers`. For each of the goroutines, a defined anonymous function helps in processing the mail message received from the `mailChan` that is continuously received using a `range` loop until the channel closes.

Remember that `defer` works in LIFO; It schedules the complete execution of the goroutine for processing the mail message and sends the completion signal to the `completeChan` when the goroutine finishes execution.

```go
for x := 0; x < worker; x += 1 {
go func() {
 // Signal completion
 defer func() {
  completionChan <- true
 }()
 for m := range mailChan {
  MailServer(m)
  } 
 }()
}
```

---

The program executes another `for` loop with the `worker` variable. This loop receives and discards completion signals from the `completeChan` channel using `<-completeChan`. The `defer` anonymous function manages this process. This loop effectively blocks the main goroutine, ensuring it waits for the completion signal from the `completeChan` channel until all the generated goroutines have finished processing.

```go
for x := 0; x < worker; x += 1 {
 <-completionChan
}
```

---

This function guarantees the completion of all worker goroutines' tasks before returning. The primary function, specifically the main goroutine, calls and executes it.

## Intialising the Application Server in the Main function

All this while, you must be wondering what the primary function entails or what it does. The `package main` function is the gateway function that helps execute all the code you have implemented. It is the only program recognised by the **Go** compiler.

The Mail Server API representation of the application server shows that the packages created and the code in each lead back to the **main**.

In the following few paragraphs, you will get a comprehensive understanding that will help you make sense of all you have learnt so far while initialising the primary function and application server.

---

Before you do that, update the `.env` file with the `PORT` address or the number of `8080` on which the HTTP server will listen and respond.

```env
PORT=8080

APP_PASSWORD=iqcepxuzsdumzslx

GMAIL_ACC=yusufakinleye144@gmail.com

USER_NAME=bravebyte

URI=mongodb+srv://ayaaakinleye:2701Akin2000@cluster0.opv1wfb.mongodb.net/?retryWrites=true&w=majority
```

---

In the `main` package, import all the needed builtins packages alongside the customs-created packages in the application.

```go
package main

import (
 "context"
 "log"
 "net/http"

 "github.com/akinbyte/mailapp/db"
 "github.com/akinbyte/mailapp/email"
 "github.com/akinbyte/mailapp/handlers"
 "github.com/akinbyte/mailapp/model"
 "github.com/joho/godotenv"
)
```

---

Three essential variables are employed: `MailChan`, a channel of type `chan model.Mail` for transmitting mail data between the main goroutine and other goroutines; The `BufferSize` variable specifies the capacity of the buffered channel, and the `Worker` variable handles the number of goroutines to create.

```go
var (
 MailChan   chan model.Mail
 BufferSize int
 Worker    int
)
```

---

In the `main` function, `MailChan` is assigned, allocated and initialised an object of type `chan model.Mail` with `BufferSize` a buffer capacity using the `make` built-in function. The `Worker` variable is assigned a value of 5.

```go
MailChan = make(chan model.Mail, BufferSize)
Worker = 5
```

---

The `godotenv.Load()` is used to read and access the value of the environment variable in the `.env` files using the key, and if an error occurs, the program `panic` logs the error.
Logs to show the database is about to connect.

```go
err := godotenv.Load()
if err != nil {
 log.Fatal(err)
}

log.Println("Starting the Mail App Server")

log.Println("Preparing Database Connection")
```

---

The application server establishes the database connection by calling the `OpenConnect` function before it starts. The `defer` function closes the link to the database after executing the main goroutine.

```go
client := db.OpenConnect()
defer func(ctx context.Context) {
 err := client.Disconnect(ctx)
 if err != nil {
  return
 }
}(context.TODO())

```

---

The `MailChan` and `Worker` is passed as an argument to the `MailDelivery` function from the `email` package to generate the goroutines, and the `MailChan` is closed once the main goroutine finishes executing.

```go
go email.MailDelivery(MailChan, Worker)
defer close(MailChan)
```

---

The `NewMailApp` function in the `handlers` package initialises the `MailApp struct` by accepting the `client` and `MailChan` as arguments. The `app` variable receives the assignment of the resulting instance.

```go
app := handlers.NewMailApp(client, MailChan)
```

---

The `Routes` function from the `handlers` package receives the `app` variable as an argument. It returns a value of type `*chi.Mux`, which you assign to the `handle` variable. The 'handle' variable initialises the application server on the port address `8080` .

```go
handle := Routes(app)

srv := http.Server{
Addr: fmt.Sprintf(":%s", os.Getenv("PORT")),
Handler: handle,
}
```

---

The application server is started and listens on the port address of `8080`; if an error occurs, the program panics and logs the error.

```go
if err := srv.ListenAndServe(); err != http.ErrServerClosed {
log.Fatalf("Shutting Down the Mail App Server ")
}
```

---

If the server shuts down using `CRTL + C`, the `main` function completes execution and the program exits.
The implementation and building of a concurrent mail server API has ended and is now ready to be tested.

## Testing the application API

As you have finished building the mail server API, it is time to test the application. To test run the application, you will need a REST Client API tool such as Postman, Insomnia or Thunder Client Extension in Vscode. I will be using the **ThunderClient** extension.

To follow along, make sure you have Thunder Client extension or Postman installed and set up the API as described below:

- Create an API collection named **MailAPP**, then add a new request as shown below
![create-collection-request](https://i.imgur.com/Z3WSRX7.png)

- Add a new request for the home page, which you will use to render the application interface later.
![home-page-request](https://i.imgur.com/g7NnLQg.png)

- Add the HTTP **GET** request method and configure it with the correct URL.
![home-request-config](https://i.imgur.com/Z3LeWnd.png)

- Create and configure new HTTP **POST** requests for subscribers to **subscribe** with their details and **send mail** content of the uploaded document with their URLs.
![subscribe-request](https://i.imgur.com/Rr3VKfY.png)
![send-mail](https://i.imgur.com/7cfPkSR.png)

Now that you've made configurations for the API. Let's move on to build and run the application to start the local server.

### Starting the API Server

To start the API server, Open your terminal or command prompt and ensure you are in the **mail app** project directory.

Execute this command below to build the application and start up the server.

For Linux or MacOS users (using the wildcard syntax):

```go
go run *.go
```

For Windows users:

```go
go run main.go routes.go
```

### Test Run

The Server is up and running and connected to the database. You can test to process each request created earlier one at a time.
![start-server](https://i.imgur.com/6Tclslv.png)

To test each request, click the **send** button to initialise the process. Start by requesting for the subscriber to submit their details.
![sub-details](https://i.imgur.com/Xlp8eaA.png)

Having the same output above would be best, indicating that the request is processed successfully. Now Go ahead and try sending the mail to all the subscribers.
![send-email](https://i.imgur.com/gArrNEQ.png)

The email sent to all registered subscribers without problems or mistakes indicates success.

To confirm that the registered subscriber received the sent message, I will check the email account's inbox used below.
![inbox-message](https://i.imgur.com/sB0g9zW.png)
![inbox-content](https://i.imgur.com/Lw1l1hm.png)
![inbox-header](https://i.imgur.com/J7dsaRY.png)

# Building User and Subscriber Interface for Mail Sending

Now that it is confirmed that the API is working as expected. You can now build an interface using HTML & CSS and integrate it with the Mail Server API.

## Creating a form for mail details and uploading a document

A brief explanation of the important parts of the interface structure of the application's client side.

---

It's described below how to create form inputs for the email title and its document, setting the form method to `post` and the encoding method to `multipart/form-data` to handle and process the uploaded document of the mail content by the  API and the `**submit**` button to be clicked to submit the form to the URL `/api/send` which triggers an `alert` message.

```html
<!-- Form for a user to upload the mail document to sent -->
      <form action="/api/send" method="post" enctype="multipart/form-data">
        <label for="">Title</label>
        <div class="user-input">
          <input type="text" name="docx_name" id="name" />
        </div>
        <label for="">Document</label>
        <div class="user-input">
          <input type="file" name="docx" id="" />
        </div>
        <button type="submit" value="" onclick="sendMailAlert()">Send</button>
      </form>
```

---

## Creating a form for subscriber details

Here is a form created for the subscriber to submit their details in the input fields named `first_name`, `last_name`, `email`, and `interest` and have it stored in the database by the  API. Also, the `**submit**` button triggers an 'alert' message when clicked to submit the form to this endpoint `/api/submit`.

```html
form action="/api/submit" method="post">
          <label for="">First Name</label>
          <div class="user-input">
            <input type="text" name="first_name" id="first_name" />
          </div>
          <label for="">Last Name</label>
          <div class="user-input">
            <input type="text" name="last_name" id="last_name" />
          </div>
          <label for="">Email</label>
          <div class="user-input">
            <input type="email" name="email" id="email" />
          </div>
          <label for="">Interest</label>
          <div class="user-input">
            <input type="text" name="interest" id="interest" />
          </div>
          <button type="submit" onclick="submitAlert()">Subscribe</button>
        </form>
```

---

check the full [HTML](https://github.com/akinbyte/mail-app/blob/main/index.html) design template and the [CSS](https://github.com/akinbyte/mail-app/blob/main/static/style.css) file.

# Conclusion

## Summary of this Article

The article comprehensively outlines the implementation of a concurrent mail server, covering topics such as concurrency models, API development, database integration, and Gmail API integration. It also discusses the setup process, testing, and ways to improve the user interface. This article is a practical guide for individuals looking to build a robust mail server with concurrency features.

## Next steps and reflections

The next steps involve enhancing functionality by adding features like email filtering, attachment handling, and advanced search capabilities, improving the user experience through intuitive forms and real-time validation, optimizing performance with load balancing and caching, and ensuring security measures.

Here is the GitHub [link](https://github.com/akinbyte/mail-app) to the repo.

