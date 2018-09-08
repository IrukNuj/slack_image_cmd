# slack_image_cmd
This bot will reply the image URL to the slash command.
<div align="center">
  <img src="https://qiita-image-store.s3.amazonaws.com/0/81341/e11d9792-19b4-f391-554b-30db0169ebfc.gif" height=464 wigth=720>
</div>

## Set up
### 1. Google custom search
- Get google custom search API Key -> CUSTOM_SEARCH_KEY
<br>[Google Custom search](https://developers.google.com/custom-search/json-api/v1/overview?hl=en)
- Plese add search engine and get engine key -> CUSTOM_SEARCH_ENGINE_ID
<br>[Add search engine](https://cse.google.com/manage/all)

### 2. Slack bot
- Create slack bot and get slack token
<br>[Create slack bot](https://slack.com/customize/slackbot)

### 3. Create slack slash command
- [Create new app](https://api.slack.com/apps?new_app=1) and select workspace
- select "Slash Command"
- create new slash command
- select "Basic informaition"
- get "Verification Token"

### 4. Environment
- `PORT`: Server port, Default `8080`
- `VERIFICATION_TOKEN`: Slack vertification token
- `CUSTOM_SEARCH_KEY`: Google custom search api key
- `CUSTOM_SEARCH_ENGINE_ID`: Google custom search engine ID

## Local build
### Execution
1. installed ngrok
2. innstalled go v1.10.1

### Run
```
# Start ngrok port 8080
$ ngrok 8080
# Start go server
$ go run main.go
# Stop ngrok OR go server
$ ctrl + c
```
