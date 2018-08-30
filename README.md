# slack_image_cmd
This bot will reply the image URL to the slash command

## How to
/image search key word

### Ex: /image twitter
![slash_command_aaa.gif](https://qiita-image-store.s3.amazonaws.com/0/81341/e11d9792-19b4-f391-554b-30db0169ebfc.gif)

## Set up
### 1. Google custom search
- Get google custom search API Key
<br>[Google Custom search](https://developers.google.com/custom-search/json-api/v1/overview?hl=en)
- Plese add search engine and get engine key
<br>[Add search engine](https://cse.google.com/manage/all)
### 2. Slack bot
- Create slack bot and get slack token
<br>[Create slack bot](https://slack.com/customize/slackbot)
### 3. Add environment value
- SLACK_API_TOKEN
- CUSTOM_SEARCH_KEY
- CUSTOM_SEARCH_ENGINE_ID

## Local build
- install ngrok
    - ngork  http 8080

- port :8080
-