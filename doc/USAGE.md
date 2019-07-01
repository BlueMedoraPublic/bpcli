
# bpcli usage

This doc will walk you through how to create a mysql source.

Steps:
1. Create credential
2. Get collector ID
3. Create source
4. Check job status

`bpcli` integrates well with common unix utilities. This guide
relies on several UNIX utilities such as `tr`, `jq`, `grep`, and `watch`.

## Setup Environment
Export your api key:
```
export BINDPLANE_API_KEY='00000000-0000-0000-0000-000000000000'
```

## Create Credential

Find the credential id and set it to an environment variable:
```
CRED_TYPE_ID=`bpcli source type get --id mysql --json | jq '.credential_types[0].id' | tr -d '"'`
```

Retrieve a credential template using the previously set credential type id
```
bpcli credential type template --id $CRED_TYPE_ID > cred_mysql.json
```

Edit the credential `name` and `parameters` to reflect your
environment:
```
{
  "name": "app-db-0",
  "credential_type_id": "03e7424b-83c5-41ca-8365-2abf528881d5",
  "parameters": {
    "password": "securepassword12",
    "username": "app"
  }
}
```

Send the credential to BindPlane:
```
bpcli credential create --file cred_mysql.json
```

Verify that the credential was created:
```
bpcli credential list | grep app-db-0
```


## Get Collector ID
Retrieve a collector id and make a note of it:
```
bpcli collector list
```

## Configuring a Source

Retrieve a source template with:
```
bpcli source type list
bpcli source type template --id mysql > source_mysql.json
```

Edit the template with your favorite editor. You must include the
collector id and credential id from the previous steps, as well as
the source name and host. All other options can use default values.
```
{
  "name": "app-db-0",
  "source_type": "mysql",
  "collector_id": "a9702bff-a3c8-4a86-862c-a6d07dc8b2fb",
  "collection_interval": 60,
  "credentials": {
    "credentials": "741e28cf-637e-45bf-a02e-5ed415a78eea"
  },
  "configuration": {
    "collection_mode": "normal",
    "connection_timeout": 15,
    "host": "app-db-0.domain.net",
    "monitor_databases": "user",
    "monitor_queries": true,
    "monitor_tables": true,
    "order_queries_by": "avg_latency",
    "order_tablespaces_by": "file_size",
    "port": 3306,
    "query_count": 10,
    "query_history_interval": 24,
    "query_timeout": 5,
    "ssl_config": "No Verify",
    "table_space_count": 200
  }
}
```

Create the source and note the `job_id` that is returned
```
bpcli source create --file source_mysql.json
```

## Watch Job Status
```
bpcli job get --id <job id here>
```
