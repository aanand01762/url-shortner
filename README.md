#  Application to create short url

#### How to run the application
```
git clone https://github.com/aanand01762/url-shortner.git
cd url-shortner
go build
./url-shortner
```

#### How to run the application inside container
* Build conatiner image using Dockerfile
```
git clone https://github.com/aanand01762/url-shortner.git
cd url-shortner
docker build .
```
* Tag the image
```
url-shortner % docker images
REPOSITORY   TAG       IMAGE ID       CREATED        SIZE
<none>       <none>    0e67cb6971c4   16 hours ago   310MB
url-shortner % docker tag  0e67cb6971c4 url-shortner:v1
url-shortner % docker images
REPOSITORY     TAG       IMAGE ID       CREATED        SIZE
url-shortner   v1        0e67cb6971c4   16 hours ago   310MB
```
* Run the container using bind mount, below command will mount the outputs directory inside the container to the current working directory. Thus user can access the test.json (output file) where records are stored. 
```
docker run -d -it --name <conatiner_name> -p <localhost port>:8080 --mount type=bind,source="$(pwd)",target=/app/outputs  url-shortner:v1
```


## API Contracts

#### Create Record
---
  Craete short url and update output file with latest, and return the newly craeted record.
* **URL**
  /records/
* **Method:**
  `POST`
*  **URL Params**
   None
* **Data Params**
   **Required:**
    `{
        url: string
     }`    
  
#### Delete Record
---
  Delete url entry from outputs/json file and return the updated list of records.
* **URL**
  /records/{id}
* **Method:**
  `DELETE`
  
*  **URL Params**
   **Required:**
   `id=integer`
* **Data Params**
  None
  
#### Get Records
---
  Returns json data of all url entries.
* **URL**
  /records
* **Method:**
  `GET`
* **Data Params**
  None

