# MinervaApi
**This is a simple Go and Firebase-based RESTful API designed to manage data stored in Firebase. The API allows users to create new topics and associated research. The API is designed to be easy to use and easy to set up.**

## Installation
To use this API, you will need to install **Go** and the **Firebase SDK**.
Create a project in the Firebase Console and add a service account. This will be used to communicate with Firebase.
Add the configuration file for the Firebase SDK to the project folder.
Create a Go module using the go mod init command.

## Usage
To use the API, clone this repo and run the **main.go** file.(After insert your key to the project) The API listens on port 8080 by default. To create a new topic, send a POST request to /topic with a JSON payload containing the topic title and author JWT. To create a new research, send a POST request to /topic/research with a JSON payload containing the research title, content, author jwt, contributor, and topic ID.
### Necessary Variables

To run this project, you will need to add your firebase key json file to project's api folder.

`key.json`

You can find your private key on Your firebase account -> Your firebase project -> Project Settings ->Service accounts -> Firebase Admin SDK -> (I choose Node.js) Generate new private key =>
your private key will be downloaded to your device.


#### Example Requests
```http
POST /user HTTP/1.1
Host: localhost:8080
Content-Type: application/json
```

| Parameter | Type     | Value                |
| :-------- | :------- | :------------------------- |
|`name` | `string` | "User's fullname" |
|`email` | `string` | "user@example.com" |
|`password` | `string` | "123456" |
|`photo_url` | `string` | "user.png" |


```http
POST /signin HTTP/1.1
Host: localhost:8080
Content-Type: application/json
```

| Parameter | Type     | Value                |
| :-------- | :------- | :------------------------- |
|`email` | `string` | "user@example.com" |
|`password` | `string` | "123456" |


```http
POST /topics HTTP/1.1
Host: localhost:8080
Content-Type: application/json
```

| Parameter | Type     | Value                |
| :-------- | :------- | :------------------------- |
| `title` | `string` | "Example Topic" |
| `author_jwt` | `string` | "[JWT]" |

(author_jwt (They are current user's jwt) will be sent by api if signin is succesfull)

```http
POST /topic/research HTTP/1.1
Host: localhost:8080
Content-Type: application/json
```
| Parameter | Type     | Value                |
| :-------- | :------- | :------------------------- |
| `"title"` | `string` | "Example Research" |
| `"content"` | `string` | "paragraph1" |
| `"author_jwt"` | `string` | "[JWT]" |
| `"contributor"` | `string` | "54321" |
| `"topic_id"` | `string` | "67890" |

#### Example Responses

```http
HTTP/1.1 200 OK
Content-Type: application/json
```

| Parameter | Type     | Value                |
| :-------- | :------- | :------------------------- |
| `data`    | `map[string]string` |`"name" `: "User's fullname" <br> `"email"` : "user@example.com" <br> `"title"` : "Research's title" <br> `"password"` : "" <br> `"photo_url"` : "user.png" |
| `error`    | `string` |"null"        |
| `status`    | `string` |"true"        |

```http
HTTP/1.1 401 Unauthorized
Content-Type: application/json
```
| Parameter | Type     | Value                |
| :-------- | :------- | :------------------------- |
| `"jwt"` | `string` | "[JWT]" |
| `"message"` | `string` | "Login access permitted" |
| `"author_id"` | `string` | "12345" |

```http
HTTP/1.1 200 OK
Content-Type: application/json
```
| Parameter | Type     | Value                |
| :-------- | :------- | :------------------------- |
| `data`    | `map[string]string` |` "id" `:      "xxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxx"  <br>  `"title"` :     "Research's title"  <br>  `"author_jwt"` :  "[JWT]" <br> `"topic_id"` :  "xxxxxxxxxxxxxxxxxxxx"    |  
| `error`    | `string` |"null"        |
| `status`    | `string` |"true"        |

```http
HTTP/1.1 200 OK
Content-Type: application/json
```
| Parameter | Type                |Key         |                       
| :-------- | :------------------ |:-----------|  
| `data`    | `map[string]string` |` "id" `:     "xxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxx"  <br>  `"title"` :     "Research's title"  <br>  `"author_jwt"` :  "[JWT]" <br> `"topic_id"` :  "xxxxxxxxxxxxxxxxxxxx"  <br> `"content"` : "paragraph1" <br> `"contributor"` : "Me" | 
| `error`    | `string` |"null"        ||
| `status`    | `string` |"true"        ||


## Contributing
Contributions are welcome! Please feel free to fork this repository and submit pull requests.

#### License
-> This project is licensed by me :).
