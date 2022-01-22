# Starlu teaser cloud function

Cloud function + website for resizing, cropping and bordering images for pragalicious.com.

Depending on the lib and CLI tool in https://github.com/ardtieboy/starlu.

## Usage

Install the httpie tool

    http -f POST https://europe-west1-starlu.cloudfunctions.net/BorderImage myFile@sdf.jpg -o /tmp/blub.jpg

## Deploying

From the root of the repo:

    gcloud functions deploy BorderImage --runtime go116 --trigger-http --allow-unauthenticated 

