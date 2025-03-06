# Simnple Encore.dev App 

## Quick Start

- encore app create
- edit app


```
encore run
  ✔ Building Encore application graph... Done!
  ✔ Analyzing service topology... Done!
  ✔ Generating boilerplate code... Done!
  ✔ Compiling application source code... Done!
  ✔ Starting Encore application... Done!

  Encore development server running!

  Your API is running at:     http://127.0.0.1:4000
  Development Dashboard URL:  http://127.0.0.1:9400/mapservices-fps2

  New Encore release available: v1.46.6 (you have v1.46.5)
  Update with: encore version update
```


## Local Run URL: 

../index.html?apiUrl=http://127.0.0.1:4000/getPoints

## Docker Image 

```
encore build docker myencoredemo1
4:22PM INF compiling Encore application for linux/amd64
4:22PM INF saving image to local docker daemon
4:22PM INF successfully saved local docker image
4:22PM INF successfully exported app as docker image
```

- docker images

- docker save -o encoredemo.tar encoredemo:latest 

- docker load -i encoredemo.tar

- docker tag myencoredemo1 us-east1-docker.pkg.dev/encoredemo/demorepo/myencoredemo1:latest
- docker push us-east1-docker.pkg.dev/encoredemo/demorepo/myencoredemo1:latest
- docker pull us-east1-docker.pkg.dev/encoredemo/demorepo/myencoredemo1:latest
- gcloud run deploy my-app --image us-east1-docker.pkg.dev/encoredemo/demorepo/myencoredemo1 --region us-east1 --platform managed --allow-unauthenticated
- https://console.cloud.google.com/run?referrer=search&project=encoredemo
```
Deploying container to Cloud Run service [my-app] in project [encoredemo] region [us-east1]
OK Deploying new service... Done.                                                                               
  OK Creating Revision...                                                                                       
  OK Routing traffic...                                                                                         
  OK Setting IAM Policy...                                                                                      
Done.                                                                                                           
Service [my-app] revision [my-app-00001-fhb] has been deployed and is serving 100 percent of traffic.
Service URL: https://my-app-412519109003.us-east1.run.app
```

https://myencoredemo1-412519109003.us-central1.run.app/getPoints

Learn More: https://encore.dev/docs/how-to/self-host

https://console.cloud.google.com/artifacts?referrer=search&project=encoredemo