<h1> Query Master </h1>

<h3> Data Pipeline </h3>
<p>
This does not supports elastic update at the moment, with CUD operations made on Postgres database outside
the scope of this application. The sync works in fashion

- Every [X] Seconds, we make DB queries to check for any updates [Currently this is done on master instance, ideally we should direct it to a different replica instance]
- If there has been any of the CUD operations on three tables - hashtags, projects, and users, Elastic sync is done.
</p>

<h4> Optimal alternative [For future scope]</h4>
<p>
Instead of syncing on every [X] sec, on any chance, we are re-ingesting the entire data. But, we move ingestion triggers to func of insertions,
updations and deletions, we could save on the network payload size of ingestion call.
</p>
<hr>

<h3> Spin Up Instructions </h3>
<li> Run `make run`</li>
<li> Run `make migrate-up` </li>
<li>Ensure Ping is working by trying `curl -X GET localhost:8080/ping` </li>
<li> ingest the seed data by running query present in `ingest_seed_data.sql` in postgres shell of docker 
 use `make make docker-exec-psql ` for running postgres shell</li>

<hr>
<h3> All Set</h3>
<li> Now you can try curls for get by user IDs and Hashtags as </li>

```
curl --location --request GET 'localhost:8080/get_project_by_hashtags?hashtags=Tag1&hashtags=Tag2'
curl --location --request GET 'localhost:8080/get_project_by_user?user_id=1'
curl --location --request GET 'localhost:8080/get_project_by_search?query=for'
```

