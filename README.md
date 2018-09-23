# quoteboard_slack

if challenge, respond with challenge response

if event, 
    check if app_mention - https://api.slack.com/events/app_mention
        see if a vote has already been started? otherwise send a message to start a vote
    check if reaction_added - https://api.slack.com/events/reaction_added
        if reaction = thumbsup, item_user = <bot's item_user id>, and item { type = message }, get the timestamp
            get the content of the message, make sure it was a message to start a vote (TODO: figure out how to - https://api.slack.com/methods/conversations.replies)
                edge case: make sure that this hasnt already been added to the quote board
