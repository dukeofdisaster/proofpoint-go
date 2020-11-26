# proofpoint-go
Client for proofpoint POD logging. 

- lib/models: ProofpointMessage, empty fields not omitted
- lib/modelsomitempty: ProofpointMessage, empty fields omitted

## TODO
- [WIP]: flesh out the client; from what I can tell websockets should keep
connection open and stream events, however this has not been the case with the
proofpoint POD api; in a python version of this using asyncio, the server will
often term the session with OK status, so implementing a polling model is
unavoidabable (or at least has been for me, suppose I could handle connection
close by just re-authing in the code in an infinite loop or something,
but that's effectively polling).
- Deduping: used sqlite3 db to dedup GUIDs in the python version; will see how
this scales. Realistically should only ever need to store a day or two worth
of seen GUIDs, so if size of DB is a concern this can be handled by just
rolling the DB with a bash script every X days.  
