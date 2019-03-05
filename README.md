# About
This project is created to accumulate all possibilities to get data from SoundCloud API *without* API-key and registration of your app.

Currently it's kinda sketchy, **but** it works!

## To test this
* Retrieve client_id from your browser to .json file
* Get User-agent to .json file

```json
{
  "client_id": "Zju0N4zmk_____Bebep9nP703ozMTpx",
  "user_agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.109 Safari/537.36"
}
```

Currently it only can download tracks and sort them by genre, so just type
```bash
go run main.go
```

And that's it


## TO-DO
- [ ] Create startup flags and manual
- [ ] Retrieve client_id
- [ ] Multi-thread application
- [ ] Storage of ignored IDs (that return 404)

