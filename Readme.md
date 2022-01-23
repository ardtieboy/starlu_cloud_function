# Starlu teaser cloud function

Cloud function + website for resizing, cropping and bordering images for pragalicious.com.

Depending on the lib and CLI tool in https://github.com/ardtieboy/starlu.

## Usage

Install the httpie tool

    http -f POST https://europe-west1-starlu.cloudfunctions.net/BorderImage myFile@sdf.jpg myFileName=WorldOfWarcraft -o /tmp/blub.jpg

## Deploying the cloud function

From the root of the repo:

    gcloud functions deploy BorderImage --runtime go116 --trigger-http --allow-unauthenticated 

## Deploying the website

    gcloud config set project starlu

    cd website

    gsutil cp * gs://starlu_cloud_function_website

Make the files public

    gsutil iam ch allUsers:objectViewer gs://starlu_cloud_function_website
