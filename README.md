<h1> Query Master </h1>

<h3> Spin Up Instructions </h3>
<li> Run `make run`</li>
<li> Run `make migrate-up` </li>
<li>Ensure Ping is working by trying `curl -X GET localhost:8080/ping` </li>
<li> ingest the seed data by running query present in `ingest_seed_data.sql` in postgres shell of docker 
 use `make make docker-exec-psql ` for running postgres shell</li>

<hr>
<h3> All Set</h3>
<li> Now you can try curls for get by user IDs and Hashtags as </li>

```asdas
curl --location --request GET 'localhost:8080/get_project_by_hashtags?hashtags=Tag1&hashtags=Tag2'
curl --location --request GET 'localhost:8080/get_project_by_user?user_id=1'
```

