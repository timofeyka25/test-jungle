### Prerequisites

Before you can start this service, you need to have the following software installed on your machine:

- Docker
- Docker Compose

### Installing

To install this service, follow these steps:

Clone the repository to your local machine:

    git clone https://github.com/timofeyka25/test-jungle.git

Navigate to the cloned repository:

    cd test-jungle

In your root directory, you have a file called `.env.example`. You need to fill all these environment variables and
rename the file to `.env`.

Then you need to download the json file with Google service account credentials to connect to the Google Storage and Google
Drive APIs. Guide for this step you can find in
this [link](https://developers.google.com/workspace/guides/create-credentials).

Set the CLOUD_CREDENTIALS_PATH variable in `.env` file to the path to the credentials.json file.

Finally, you can run this app using docker compose:
    
    docker-compose up