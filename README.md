# MinervaApi
**This is a simple Go and Firebase-based RESTful API designed to manage data stored in Firebase. The API allows users to create new topics and associated research. The API is designed to be easy to use and easy to set up.**

## Installation
To use this API, you will need to install **Go** and the **Firebase SDK**.
Create a project in the Firebase Console and add a service account. This will be used to communicate with Firebase.
Add the configuration file for the Firebase SDK to the project folder.
Create a Go module using the go mod init command.

## Usage
To use the API, clone this repo and run the **main.go** file.(After insert your key to the project) The API listens on port 7334 by default. To create a new topic, send a POST request to /topic with a JSON payload containing the topic name and creator ID. To create a new research, send a POST request to /research with a JSON payload containing the research header, content, creator ID, contributor, and topic ID.
### Necessary Variables

To run this project, you will need to add your firebase key json file to project's FireBase folder.

`key.json`

You can find your private key on Your firebase account -> Your firebase project -> Project Settings ->Service accounts -> Firebase Admin SDK -> (I choose Node.js) Generate new private key
your private key will be downloaded to your device.


#### Example Requests
```http
POST /topic HTTP/1.1
Host: localhost:7334
Content-Type: application/json
```

| Parameter | Type     | Value                |
| :-------- | :------- | :------------------------- |
| `topic_name` | `string` | "Example Topic" |
| `topic_creator_id` | `string` | "12345" |

(topic_creator_id and research_creator_id (They are user's id) will be sent by api if signin is succesfull)

```http
POST /topic/research HTTP/1.1
Host: localhost:7334
Content-Type: application/json
```
| Parameter | Type     | Value                |
| :-------- | :------- | :------------------------- |
| `research_header` | `string` | "Example Research" |
| `"research_content` | `string` | "paragraph1" |
| `"research_creator_id` | `string` | "12345" |
| `"research_contributor` | `string` | "54321" |
| `"research_topic_id` | `string` | "67890" |


#### Example Responses
```http
HTTP/1.1 201 Created
Content-Type: application/json
```
| Parameter | Type     | Value                |
| :-------- | :------- | :------------------------- |
| `message` | `string` | "Topic created successfully" |

```http
HTTP/1.1 201 Created
Content-Type: application/json
```
| Parameter | Type     | Value                |
| :-------- | :------- | :------------------------- |
| `message` | `string` | "Research created successfully" |

## Contributing
Contributions are welcome! Please feel free to fork this repository and submit pull requests.

#### License
-> This project is licensed by me :).
